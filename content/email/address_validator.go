package email

type AddressValidator func(*Address) error

func validateAddress(*Address) error {
	return nil
}
