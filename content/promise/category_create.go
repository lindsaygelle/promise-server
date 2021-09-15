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

type CategoryCreateValidator func(CategoryCreate) error

func DecodeCategoryCreate(readCloser io.ReadCloser) (categoryCreate CategoryCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&categoryCreate)
	if err != nil {
		return
	}
	err = validateCategoryCreate(categoryCreate)
	return
}

func validateCategoryCreate(categoryCreate CategoryCreate) error {
	for _, validator := range []CategoryCreateValidator{
		validateCategoryCreateName,
		validateCategoryCreateProfileID} {
		if err := validator(categoryCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateCategoryCreateName(categoryCreate CategoryCreate) error {
	if len(categoryCreate.Name) == 0 {
		return ErrCategoryName
	}
	return nil
}

func validateCategoryCreateProfileID(categoryCreate CategoryCreate) error {
	if categoryCreate.ProfileID == 0 {
		return ErrCategoryProfileID
	}
	return nil
}
