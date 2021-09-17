package promise

type TasksError int

func (e TasksError) Error() (s string) {
	return
}

const (
	ErrTasks TasksError = iota + 1
	ErrTasksNotFound
)
