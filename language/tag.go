package language

import (
	"time"

	"github.com/lindsaygelle/promise/promise-server/database"
)

type Tag struct {
	Created  time.Time `json:"created"`
	ID       uint      `json:"id"`
	Language uint      `json:"language"`
	Name     string    `json:"name"`
	Tag      string    `json:"tag"`
}

// GetTag returns a language.Tag.
func GetTag(client database.Client, id string) (Tag, error) {
	row, err := client.QueryRow(`SELECT created, id, language, name, tag FROM language.tag WHERE id=$1`, id)
	if err != nil {
		return Tag{}, err
	}
	return NewTag(row)
}

// GetTags returns a slice of language.Tag.
func GetTags(v database.Client) ([]Tag, error) {
	var tags []Tag
	rows, err := v.Query(`SELECT created, id, language, name, tag FROM language.tag`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = addTag(&tags, rows)
		if err != nil {
			return nil, err
		}
	}
	return tags, nil
}

// NewTag returns a new language.Tag.
//
// NewTag returns an error on the condition it cannot correctly scan the database row.
func NewTag(v database.Scanner) (tag Tag, err error) {
	err = v.Scan(&tag.Created, &tag.ID, &tag.Language, &tag.Name, &tag.Tag)
	return tag, err
}

// addTag scans a language.tag record from the database rows and adds it to the collection.
//
// addTag returns an error on the condition the account cannot be scanned.
func addTag(v *[]Tag, rows database.Rows) error {
	tag, err := NewTag(rows)
	if err != nil {
		return err
	}
	*v = append(*v, tag)
	return err
}
