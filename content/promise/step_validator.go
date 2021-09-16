package promise

type StepValidator func(Step) error

func validateStep(*Step) error {
	return nil
}
