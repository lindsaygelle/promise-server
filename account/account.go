package account

import (
	"database/sql"
	"encoding/json"

	"github.com/lindsaygelle/promise/promise-server/database"
)

type Profile struct {
	Created string `json:"created_at"`
	Email   uint   `json:"email"`
	ID      uint   `json:"id"`
}

// ReadProfile reads an profile from the database by its ID.
func ReadProfile(database *sql.DB, ID string) (profile Profile, err error) {
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	row := transaction.QueryRow(`SELECT ROW_TO_JSON(T.*) FROM account.profile AS T WHERE id=$1`, ID)
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
	rows, err := transaction.Query(`SELECT ROW_TO_JSON(T.*) FROM account.profile AS T`)
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
