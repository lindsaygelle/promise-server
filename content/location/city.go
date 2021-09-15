package location

import (
	"encoding/json"
	"io"
	"time"
)

type City struct {
	CountryID   uint      `json:"country_id"`
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	TimeCreated time.Time `json:"time_created"`
	TimeEdited  time.Time `json:"time_edited"`
}

func DecodeCity(readCloser io.ReadCloser) (city City, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&city)
	return
}
