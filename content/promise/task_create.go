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

func DecodeTaskCreate(readCloser io.ReadCloser) (taskCreate TaskCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&taskCreate)
	if err != nil {
		err = ErrTask
		return
	}
	err = validateTaskCreate(taskCreate)
	return
}
