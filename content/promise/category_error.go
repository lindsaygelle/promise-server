package promise

type CategoryError int

func (e CategoryError) Error() (s string) {
	return
}

const (
	ErrCategory CategoryError = iota + 1
	ErrCategoryName
	ErrCategoryNotFound
	ErrCategoryProfileID
)
