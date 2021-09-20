package promise

import (
	"database/sql"

	"github.com/lindsaygelle/promise/promise-server/content/promise"
)

type CategoryService interface {
	Get(categoryID string) (promise.Category, error)
	GetAll() (promise.Categories, error)
	GetAllByProfile(profileID string) (promise.Categories, error)
}

type categoryService struct {
	*sql.DB
}
