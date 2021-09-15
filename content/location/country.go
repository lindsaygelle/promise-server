package location

import (
	"encoding/json"
	"io"
	"time"
)

type Country struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	TimeCreated time.Time `json:"time_created"`
	TimeEdited  time.Time `json:"time_edited"`
}

func DecodeCountry(readCloser io.ReadCloser) (country Country, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&country)
	return
}
