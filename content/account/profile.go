package account

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"time"
)

type Profile struct {
	ID          uint      `json:"id"`
	IsVerified  bool      `json:"is_verified"`
	Name        string    `json:"name"`
	TimeCreated time.Time `json:"time_created"`
	TimeEdited  time.Time `json:"time_edited"`
}

func DecodeProfile(readCloser io.ReadCloser) (profile Profile, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&profile)
	if err != nil {
		err = ErrProfile
		return
	}
	return
}

func ScanProfile(scanner interface{ Scan(...interface{}) error }) (Profile, error) {
	var b []byte
	err := scanner.Scan(&b)
	if err == sql.ErrNoRows {
		err = ErrProfileNotFound
		return Profile{}, err
	}
	return DecodeProfile(io.NopCloser(bytes.NewReader(b)))
}
