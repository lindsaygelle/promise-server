package promise

import "github.com/lindsaygelle/promise/promise-server/content/promise"

func (c *categoryService) Get(categoryID string) (promise.Category, error) {
	return promise.ScanCategory(nil)
}

func (c *categoryService) GetAllByProfile(profileID string) (promise.Categories, error) {
	return promise.ScanCategories(nil)
}
