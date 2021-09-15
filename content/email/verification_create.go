package email

import (
	"encoding/json"
	"io"
)

type VerificationCreate struct {
	AddressID uint `json:"address_id"`
}

func DecodeVerificationCreate(readCloser io.ReadCloser) (verificationCreate VerificationCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&verificationCreate)
	if err != nil {
		err = ErrVerification
		return
	}
	err = validateVerificationCreate(&verificationCreate)
	return
}
