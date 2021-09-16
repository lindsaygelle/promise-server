package promise

type TaskValidator func(*Task) error

func validateTask(*Task) error {
	return nil
}
