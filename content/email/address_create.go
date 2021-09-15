package email

import (
	"encoding/json"
	"io"
)

type AddressCreate struct {
	Email     string `json:"email"`
	ProfileID uint   `json:"profile_id"`
}

func DecodeAddressCreate(readCloser io.ReadCloser) (addressCreate AddressCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&addressCreate)
	if err != nil {
		err = ErrAddress
		return
	}
	err = validateAddressCreate(addressCreate)
	return
}
