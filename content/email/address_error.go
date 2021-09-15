package email

type AddressError int

func (e AddressError) Error() (s string) {
	return
}

const (
	ErrAddress AddressError = iota + 1
	ErrAddressEmail
	ErrAddressNotFound
	ErrAddressProfileID
)
