package account

import (
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

type ProfileValidator func(Profile) error

func DecodeProfile(readCloser io.ReadCloser) (profile Profile, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&profile)
	if err != nil {
		err = ErrProfile
		return
	}
	return
}
