package account

import (
	"database/sql"
	"encoding/json"

	"github.com/lindsaygelle/promise/promise-server/database"
)

const readSetting = `
SELECT
	ROW_TO_JSON(T.*)
FROM
	account.setting AS T
WHERE T.profile=$1`

type Setting struct {
	Profile uint `json:"profile"`
}

// ReadSetting reads a Setting from the database by the profile ID.
func ReadSetting(database *sql.DB, ID string) (setting Setting, err error) {
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	row := transaction.QueryRow(readSetting, ID)
	setting, err = scanSetting(row)
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
func scanSetting(scanner database.Scanner) (setting Setting, err error) {
	var content []byte
	err = scanner.Scan(&content)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &setting)
	return
}
