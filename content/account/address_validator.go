package account

type AddressValidator func(Address) error

var addressValidators = [...]AddressValidator{}

func validateAddress(address Address) error {
	for _, validator := range addressValidators {
		if err := validator(address); err != nil {
			return err
		}
	}
	return nil
}
