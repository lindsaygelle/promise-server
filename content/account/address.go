package account

import (
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
	return
}
