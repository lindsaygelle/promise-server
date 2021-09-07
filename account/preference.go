package account

import (
	"database/sql"
	"encoding/json"

	"github.com/lindsaygelle/promise/promise-server/database"
)

const readPreference = `
SELECT
	ROW_TO_JSON(T.*)
FROM
	account.preference AS T
WHERE T.profile=$1`

type Preference struct {
	Profile uint `json:"profile"`
}

// ReadPreference reads an Preference from the database by the profile ID.
func ReadPreference(database *sql.DB, ID string) (preference Preference, err error) {
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	row := transaction.QueryRow(readPreference, ID)
	preference, err = scanPreference(row)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

// scanProfile returns a new Profile from the database.
//
// scanProfile expects the scanner to contain JSON content.
func scanPreference(scanner database.Scanner) (preference Preference, err error) {
	var content []byte
	err = scanner.Scan(&content)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &preference)
	return
}
