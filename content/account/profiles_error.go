package account

type ProfilesError int

func (e ProfilesError) Error() (s string) {
	return
}

const (
	ErrProfiles ProfilesError = iota + 1
	ErrProfilesNotFound
)
