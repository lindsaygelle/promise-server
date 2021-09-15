package account

type ProfileValidator func(*Profile) error

func validateProfile(*Profile) error {
	return nil
}
