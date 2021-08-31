package account

import (
	"time"

	"github.com/lindsaygelle/promise/promise-server/database"
)

type Setting struct {
	Account   uint      `json:"account"`
	Biography string    `json:"biography"`
	Country   uint      `json:"country"`
	Edited    time.Time `json:"edited"`
	Language  uint      `json:"language"`
}

// GetSettings returns a slice of account.Setting.
func GetSettings(client database.Client) ([]Setting, error) {
	var settings []Setting
	rows, err := client.Query(`SELECT account, biography, country, edited, language FROM account.setting`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = addSettings(&settings, rows)
		if err != nil {
			return nil, err
		}
	}
	return settings, nil
}

// NewSetting returns a new account.Setting.
//
// NewSetting returns an error on the condition it cannot correctly scan the database row.
func NewSetting(rows database.Rows) (setting Setting, err error) {
	err = rows.Scan(&setting.Account, &setting.Biography, &setting.Country, &setting.Edited, &setting.Language)
	return setting, err
}

// addSettings scans a account.setting record from the database rows and adds it to the collection.
//
// addSettings returns an error on the condition the settings cannot be scanned.
func addSettings(v *[]Setting, rows database.Rows) error {
	settings, err := NewSetting(rows)
	if err != nil {
		return err
	}
	*v = append(*v, settings)
	return err
}
