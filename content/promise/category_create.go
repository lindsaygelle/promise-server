package promise

import (
	"encoding/json"
	"io"
)

type CategoryCreate struct {
	Description *string `json:"description"`
	Name        string  `json:"name"`
	ProfileID   string  `json:"profile_id"`
}

func DecodeCategoryCreate(reader io.ReadCloser) (categoryCreate CategoryCreate, err error) {
	err = json.NewDecoder(reader).Decode(&categoryCreate)
	return
}
