package promise

import (
	"encoding/json"
	"io"
)

type CategoryCreate struct {
	Description *string `json:"description"`
	Name        string  `json:"name"`
	ProfileID   uint    `json:"profile_id"`
}

func DecodeCategoryCreate(readCloser io.ReadCloser) (categoryCreate CategoryCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&categoryCreate)
	if err != nil {
		err = ErrCategory
		return
	}
	err = validateCategoryCreate(&categoryCreate)
	return
}
