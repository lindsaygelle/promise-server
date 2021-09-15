package account

type ProfileError int

func (e ProfileError) Error() (s string) {
	return
}

const (
	ErrProfile ProfileError = iota + 1
	ErrProfileEmail
	ErrProfileName
	ErrProfileNotFound
)
