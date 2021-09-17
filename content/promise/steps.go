package promise

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
)

type Steps []Step

func DecodeSteps(readCloser io.ReadCloser) (steps Steps, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&steps)
	if err != nil {
		err = ErrSteps
		return
	}
	return
}

func ScanSteps(scanner interface{ Scan(...interface{}) error }) (steps Steps, err error) {
	var b []byte
	err = scanner.Scan(&b)
	if err == sql.ErrNoRows {
		err = ErrStepsNotFound
		return
	}
	steps, err = DecodeSteps(io.NopCloser(bytes.NewReader(b)))
	return
}
