package email

type DomainError int

func (e DomainError) Error() (s string) {
	return
}

const (
	ErrDomain DomainError = iota + 1
	ErrDomainID
	ErrDomainValue
)
