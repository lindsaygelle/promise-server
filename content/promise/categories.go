package promise

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
)

type Categories []Categories

func DecodeCategories(reader io.ReadCloser) (categories Categories, err error) {
	defer reader.Close()
	err = json.NewDecoder(reader).Decode(&categories)
	if err != nil {
		err = ErrCategories
		return
	}
	return
}

func ScanCategories(scanner interface{ Scan(...interface{}) error }) (categories Categories, err error) {
	var b []byte
	err = scanner.Scan(&b)
	if err == sql.ErrNoRows {
		err = ErrCategoriesNotFound
		return
	}
	categories, err = DecodeCategories(io.NopCloser(bytes.NewReader(b)))
	return
}
