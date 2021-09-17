package promise

type StepsError int

func (e StepsError) Error() (s string) {
	return
}

const (
	ErrSteps StepsError = iota + 1
	ErrStepsNotFound
)
