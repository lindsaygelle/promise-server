package promise

import (
	"database/sql"

	"github.com/lindsaygelle/promise/promise-server/content/promise"
)

type CategoryService interface {
	Get(categoryID string) (promise.Category, error)
	GetAllByProfile(profileID string) (promise.Categories, error)
}

type categoryService struct {
	*sql.DB
}

func NewCategoryService(database *sql.DB) CategoryService {
	return &categoryService{database}
}
