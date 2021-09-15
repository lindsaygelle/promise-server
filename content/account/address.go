package account

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"time"
)

type Address struct {
	CityID      uint      `json:"city_id"`
	CountryID   uint      `json:"country_id"`
	ID          uint      `json:"id"`
	ProfileID   uint      `json:"profile_id"`
	TimeCreated time.Time `json:"time_created"`
	TimeEdited  time.Time `json:"time_edited"`
}

func DecodeAddress(readCloser io.ReadCloser) (address Address, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&address)
	if err != nil {
		err = ErrAddress
		return
	}
	err = validateAddress(address)
	return
}

func ScanAddress(scanner interface{ Scan(...interface{}) error }) (Address, error) {
	var b []byte
	err := scanner.Scan(&b)
	if err == sql.ErrNoRows {
		return Address{}, ErrAddressNotFound
	}
	return DecodeAddress(io.NopCloser(bytes.NewReader(b)))
}
