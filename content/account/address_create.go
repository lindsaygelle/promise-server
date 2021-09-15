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

type AddressCreateValidator func(AddressCreate) error

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

func validateAddressCreate(addressCreate AddressCreate) error {
	for _, validator := range [...]AddressCreateValidator{
		validateAddressCreateCityID,
		validateAddressCreateCountryID,
		validateAddressCreateProfileID} {
		if err := validator(addressCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateAddressCreateCityID(addressCreate AddressCreate) error {
	if addressCreate.CityID == 0 {
		return ErrAddressCityID
	}
	return nil
}

func validateAddressCreateCountryID(addressCreate AddressCreate) error {
	if addressCreate.CountryID == 0 {
		return ErrAddressCountryID
	}
	return nil
}

func validateAddressCreateProfileID(addressCreate AddressCreate) error {
	if addressCreate.CountryID == 0 {
		return ErrAddressProfileID
	}
	return nil
}
