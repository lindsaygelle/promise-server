package promise

type CategoryValidator func(*Category) error

func validateCategory(*Category) error {
	return nil
}
