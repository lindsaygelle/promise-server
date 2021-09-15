package account

import (
	"encoding/json"
	"io"
)

type AddressCreate struct {
	CityID    uint `json:"city_id"`
	CountryID uint `json:"country_id"`
	ProfileID uint `json:"profile_id"`
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
