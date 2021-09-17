package account

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
)

type Profiles []Profiles

func DecodeProfiles(readCloser io.ReadCloser) (profiles Profiles, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&profiles)
	if err != nil {
		err = ErrProfiles
		return
	}
	return
}

func ScanProfiles(scanner interface{ Scan(...interface{}) error }) (profiles Profiles, err error) {
	var b []byte
	err = scanner.Scan(&b)
	if err == sql.ErrNoRows {
		err = ErrProfilesNotFound
		return
	}
	profiles, err = DecodeProfiles(io.NopCloser(bytes.NewReader(b)))
	return
}
