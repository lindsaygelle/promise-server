package email

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"time"
)

type Domain struct {
	ID          uint      `json:"id"`
	Value       string    `json:"value"`
	TimeCreated time.Time `json:"time_created"`
	TimeEdited  time.Time `json:"time_edited"`
}

func DecodeDomain(readCloser io.ReadCloser) (domain Domain, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&domain)
	return
}

func ScanDomain(scanner interface{ Scan(...interface{}) error }) (Domain, error) {
	var b []byte
	err := scanner.Scan(&b)
	if err == sql.ErrNoRows {
		return Domain{}, ErrDomainNotFound
	}
	return DecodeDomain(io.NopCloser(bytes.NewReader(b)))
}
