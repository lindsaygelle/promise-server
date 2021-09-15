package account

import (
	"encoding/json"
	"io"
)

type ProfileCreate struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func DecodeProfileCreate(readCloser io.ReadCloser) (profileCreate ProfileCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&profileCreate)
	if err != nil {
		err = ErrProfile
		return
	}
	err = validateProfileCreate(&profileCreate)
	return
}
