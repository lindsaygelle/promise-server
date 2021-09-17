package promise

type CategoriesError int

func (e CategoriesError) Error() (s string) {
	return
}

const (
	ErrCategories CategoriesError = iota + 1
	ErrCategoriesNotFound
)
