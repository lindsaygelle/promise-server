package account

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/mail"
	"strings"

	"github.com/lindsaygelle/promise/promise-server/database"
)

const readProfile = `
SELECT
	ROW_TO_JSON(T.*)
FROM
	account.profile AS T
WHERE id=$1`

const readProfiles = `
SELECT
	ROW_TO_JSON(T.*)
FROM
	account.profile AS T`

const writeProfile = `
WITH
email_address AS (
	INSERT INTO email.address (address, domain) VALUES ($1, $2)
	RETURNING id
),
email_verification AS
(
	INSERT INTO email.verification (address)
	SELECT id FROM email_address
    RETURNING address
),
account_profile AS 
(
	INSERT INTO account.profile (email)
	SELECT id FROM email_address
	RETURNING *
),
account_preference AS
(
    INSERT INTO account.preference (profile)
    SELECT id FROM account_profile
),
account_setting AS (
	INSERT INTO account.setting (profile)
	SELECT id FROM account_profile
)
SELECT ROW_TO_JSON(T.*) FROM account_profile AS T;`

type Profile struct {
	Created string `json:"created_at"`
	Email   uint   `json:"email"`
	ID      uint   `json:"id"`
}

type ProfileMake struct {
	Email string `json:"email"`
}

// ReadProfile reads an profile from the database by its ID.
func ReadProfile(database *sql.DB, ID string) (profile Profile, err error) {
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	row := transaction.QueryRow(readProfile, ID)
	profile, err = scanProfile(row)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

// ReadProfiles reads all profiles from the database.
func ReadProfiles(database *sql.DB) (profiles []Profile, err error) {
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	rows, err := transaction.Query(readProfiles)
	if err != nil {
		transaction.Rollback()
		return
	}
	err = processProfiles(&profiles, rows)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

func WriteProfile(database *sql.DB, ID string, reader io.Reader) (profile Profile, err error) {
	profileMake, err := newProfileMake(reader)
	if err != nil {
		return
	}
	err = verifyEmail(profileMake.Email)
	if err != nil {
		return
	}
	address, domain := processEmailAddress(profileMake.Email)
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	profile, err = scanProfile(transaction.QueryRow(writeProfile, address, domain))
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

// addProfile scans a account.profile record from the database rows and adds it to the collection.
//
// addProfile returns an error on the condition the profile cannot be scanned.
func addProfile(profiles *[]Profile, rows *sql.Rows) (err error) {
	profile, err := scanProfile(rows)
	if err != nil {
		return err
	}
	*profiles = append(*profiles, profile)
	return err
}

func newProfileMake(reader io.Reader) (profileMake ProfileMake, err error) {
	err = json.NewDecoder(reader).Decode(&profileMake)
	return
}

func processEmailAddress(email string) (address string, domain string) {
	i := strings.LastIndex(email, "@")
	address, domain = email[:i], email[i+1:]
	return
}

// processProfiles processes all rows from the database and adds them to the collection.
//
// processProfiles returns an error on the condition the row cannot be added.
func processProfiles(profiles *[]Profile, rows *sql.Rows) (err error) {
	for rows.Next() {
		err = addProfile(profiles, rows)
		if err != nil {
			return
		}
	}
	return
}

// scanProfile returns a new Profile from the database.
//
// scanProfile expects the scanner to contain JSON content.
func scanProfile(scanner database.Scanner) (profile Profile, err error) {
	var content []byte
	err = scanner.Scan(&content)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &profile)
	return
}

func verifyEmail(email string) (err error) {
	_, err = mail.ParseAddress(email)
	return
}
