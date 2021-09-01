package language

import (
	"time"

	"github.com/lindsaygelle/promise/promise-server/database"
)

type Language struct {
	Created time.Time `json:"created"`
	ID      uint      `json:"id"`
	Name    string    `json:"name"`
}

// GetLanguages returns a slice of language.Language.
func GetLanguages(client database.Client) ([]Language, error) {
	var languages []Language
	rows, err := client.Query(`SELECT created, id, name FROM language.language`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = addLanguage(&languages, rows)
		if err != nil {
			return nil, err
		}
	}
	return languages, nil
}

// NewLanguage returns a new language.Language.
//
// NewLanguage returns an error on the condition it cannot correctly scan the database row.
func NewLanguage(v interface{ Scan(...interface{}) error }) (language Language, err error) {
	err = v.Scan(&language.Created, &language.ID, &language.Name)
	return language, err
}

// addLanguage scans a language.language record from the database rows and adds it to the collection.
//
// addLanguage returns an error on the condition the language cannot be scanned.
func addLanguage(v *[]Language, rows database.Rows) error {
	langauge, err := NewLanguage(rows)
	if err != nil {
		return err
	}
	*v = append(*v, langauge)
	return err
}
