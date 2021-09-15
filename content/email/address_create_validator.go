package email

import "net/mail"

type AddressCreateValidator func(AddressCreate) error

var addressCreateValidators = [...]AddressCreateValidator{
	validateAddressCreateEmail,
	validateAddressCreateProfileID}

func validateAddressCreate(addressCreate AddressCreate) error {
	for _, validator := range addressCreateValidators {
		if err := validator(addressCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateAddressCreateEmail(addressCreate AddressCreate) error {
	if _, err := mail.ParseAddress(addressCreate.Email); err != nil {
		return ErrAddressEmail
	}
	return nil
}

func validateAddressCreateProfileID(addressCreate AddressCreate) error {
	if addressCreate.ProfileID == 0 {
		return ErrAddressProfileID
	}
	return nil
}
