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
	err = json.NewDecoder(readCloser).Decode(&taskCreate)
	return
}
