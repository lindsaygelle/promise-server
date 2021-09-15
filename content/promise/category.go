package promise

import (
	"encoding/json"
	"io"
	"time"
)

type Category struct {
	Description *string   `json:"description"`
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	ProfileID   uint      `json:"profile_id"`
	TimeCreated time.Time `json:"time_created"`
	TimeEdited  time.Time `json:"time_edited"`
}

type CategoryValidator func(Category) error

func DecodeCategory(reader io.ReadCloser) (category Category, err error) {
	defer reader.Close()
	err = json.NewDecoder(reader).Decode(&category)
	if err != nil {
		err = ErrCategory
		return
	}
	return
}
