package email

import "net/mail"

type AddressCreateValidator func(AddressCreate) error

var addressCreateValidators = [...]AddressCreateValidator{
	validateAddressCreateEmail,
	validateAddressCreateProfileID}

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
