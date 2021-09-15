package account

import (
	"encoding/json"
	"io"
	"time"
)

type Profile struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	TimeCreated time.Time `json:"time_created"`
	TimeEdited  time.Time `json:"time_edited"`
	Verified    bool      `json:"verified"`
}

func DecodeProfile(readCloser io.ReadCloser) (profile Profile, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&profile)
	return
}
