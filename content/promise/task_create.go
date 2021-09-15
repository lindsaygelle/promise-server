package promise

import (
	"encoding/json"
	"io"
)

type TaskCreate struct {
	CategoryID  *uint   `json:"category_id"`
	Description *string `json:"description"`
	ProfileID   uint    `json:"profile_id"`
	Name        string  `json:"name"`
}

type TaskCreateValidator func(TaskCreate) error

func DecodeTaskCreate(readCloser io.ReadCloser) (taskCreate TaskCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&taskCreate)
	if err != nil {
		err = ErrTask
		return
	}
	err = verifyTaskCreate(taskCreate)
	return
}

func verifyTaskCreate(taskCreate TaskCreate) error {
	for _, validator := range [...]TaskCreateValidator{
		verifyTaskCreateCategory,
		verifyTaskCreateName} {
		if err := validator(taskCreate); err != nil {
			return err
		}
	}
	return nil
}

func verifyTaskCreateCategory(taskCreate TaskCreate) error {
	if taskCreate.CategoryID != nil && *taskCreate.CategoryID == 0 {
		return ErrTaskCategoryID
	}
	return nil
}

func verifyTaskCreateName(taskCreate TaskCreate) error {
	if len(taskCreate.Name) == 0 {
		return ErrTaskName
	}
	return nil
}
