package promise

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"time"
)

type Category struct {
	Description *string   `json:"description"`
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	ProfileID   uint      `json:"profile_id"`
	TimeCreated time.Time `json:"time_created"`
	TimeEdited  time.Time `json:"time_edited"`
}

func DecodeCategory(reader io.ReadCloser) (category Category, err error) {
	defer reader.Close()
	err = json.NewDecoder(reader).Decode(&category)
	if err != nil {
		err = ErrCategory
		return
	}
	return
}

func ScanCategory(scanner interface{ Scan(...interface{}) error }) (Category, error) {
	var b []byte
	err := scanner.Scan(&b)
	if err == sql.ErrNoRows {
		return Category{}, ErrCategoryNotFound
	}
	return DecodeCategory(io.NopCloser(bytes.NewReader(b)))
}
