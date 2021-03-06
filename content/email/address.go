package email

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"time"
)

type Address struct {
	Address      string     `json:"address"`
	DomainID     string     `json:"domain_id"`
	ID           uint       `json:"id"`
	IsVerified   bool       `json:"is_verified"`
	ProfileID    uint       `json:"profile_id"`
	TimeCreated  time.Time  `json:"time_created"`
	TimeVerified *time.Time `json:"time_verified"`
}

func DecodeAddress(readCloser io.ReadCloser) (address Address, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&address)
	if err != nil {
		err = ErrAddress
		return
	}
	err = validateAddress(&address)
	return
}

func ScanAddress(scanner interface{ Scan(...interface{}) error }) (address Address, err error) {
	var b []byte
	err = scanner.Scan(&b)
	if err == sql.ErrNoRows {
		err = ErrAddressNotFound
		return
	}
	address, err = DecodeAddress(io.NopCloser(bytes.NewReader(b)))
	return
}
