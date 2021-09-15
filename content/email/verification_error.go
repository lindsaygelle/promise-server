package email

type VerificationError int

func (e VerificationError) Error() (s string) {
	return
}

const (
	ErrVerification VerificationError = iota + 1
	ErrVerificationAddressID
	ErrVerificationNotFound
)
