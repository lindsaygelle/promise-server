package email

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
)

type Verification struct {
	AddressID   uint      `json:"address_id"`
	Code        uuid.UUID `json:"code"`
	TimeCreated time.Time `json:"time_created"`
	TimeDue     time.Time `json:"time_due"`
}

func DecodeVerification(readCloser io.ReadCloser) (verification Verification, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&verification)
	return
}

func ScanVerification(scanner interface{ Scan(...interface{}) error }) (Verification, error) {
	var b []byte
	err := scanner.Scan(&b)
	if err == sql.ErrNoRows {
		return Verification{}, ErrVerificationNotFound
	}
	return DecodeVerification(io.NopCloser(bytes.NewReader(b)))
}
