package promise

import "github.com/lindsaygelle/promise/promise-server/content/promise"

func (c *categoryService) Get(id string) (promise.Category, error) {
	return promise.ScanCategory(nil)
}

func (c *categoryService) GetAll() (promise.Categories, error) {
	return promise.ScanCategories(nil)
}

func (c *categoryService) GetAllByProfile(profileID string) (promise.Categories, error) {
	return promise.ScanCategories(nil)
}
