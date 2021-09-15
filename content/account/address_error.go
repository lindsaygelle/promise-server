package account

type AddressError int

func (e AddressError) Error() (s string) {
	return
}

const (
	ErrAddress AddressError = iota + 1
	ErrAddressCityID
	ErrAddressCountryID
	ErrAddressProfileID
)
