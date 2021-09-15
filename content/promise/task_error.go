package promise

type TaskError int

func (e TaskError) Error() (s string) {
	return
}

const (
	ErrTask TaskError = iota + 1
	ErrTaskID
	ErrTaskCategoryID
	ErrTaskName
	ErrTaskNotFound
)
