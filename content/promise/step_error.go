package promise

type StepError int

func (e StepError) Error() (s string) {
	return
}

const (
	ErrStep StepError = iota + 1
	ErrStepName
	ErrStepNotFound
	ErrStepTaskID
	ErrStepTimeDue
)
