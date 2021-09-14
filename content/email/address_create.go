package email

import (
	"encoding/json"
	"io"
)

type AddressCreate struct{}

func DecodeAddressCreate(readCloser io.ReadCloser) (addressCreate AddressCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&addressCreate)
	return
}
