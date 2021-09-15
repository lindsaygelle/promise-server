package account

import (
	"encoding/json"
	"io"
	"net/mail"
)

type ProfileCreate struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ProfileCreateValidator func(ProfileCreate) error

func DecodeProfileCreate(readCloser io.ReadCloser) (profileCreate ProfileCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&profileCreate)
	if err != nil {
		err = ErrProfile
		return
	}
	err = validateProfileCreate(profileCreate)
	return
}

func validateProfileCreate(profileCreate ProfileCreate) error {
	for _, validator := range [...]ProfileCreateValidator{
		validateProfileCreateEmail,
		validateProfileCreateName} {
		if err := validator(profileCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateProfileCreateEmail(profileCreate ProfileCreate) error {
	_, err := mail.ParseAddress(profileCreate.Email)
	if err != nil {
		return ErrProfileEmail
	}
	return nil
}

func validateProfileCreateName(profileCreate ProfileCreate) error {
	if len(profileCreate.Name) == 0 {
		return ErrProfileName
	}
	return nil
}
