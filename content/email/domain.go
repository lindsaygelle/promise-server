package email

import (
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
