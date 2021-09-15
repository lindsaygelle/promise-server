package account

type AddressCreateValidator func(AddressCreate) error

var addressCreateValidators = [...]AddressCreateValidator{
	validateAddressCreateCityID,
	validateAddressCreateCountryID,
	validateAddressCreateProfileID}

func validateAddressCreate(addressCreate AddressCreate) error {
	for _, validator := range addressCreateValidators {
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
