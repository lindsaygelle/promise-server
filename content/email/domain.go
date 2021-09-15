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
	if err != nil {
		err = ErrDomain
		return
	}
	err = validateDomain(&domain)
	return
}

func ScanDomain(scanner interface{ Scan(...interface{}) error }) (domain Domain, err error) {
	var b []byte
	err = scanner.Scan(&b)
	if err == sql.ErrNoRows {
		err = ErrDomainNotFound
		return
	}
	domain, err = DecodeDomain(io.NopCloser(bytes.NewReader(b)))
	return
}
