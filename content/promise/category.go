package promise

import (
	"encoding/json"
	"io"
	"time"
)

type Category struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	EditedAt    time.Time `json:"edited_at"`
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	ProfileID   uint      `json:"profile_id"`
}

func DecodeCategory(reader io.ReadCloser) (category Category, err error) {
	err = json.NewDecoder(reader).Decode(&category)
	return
}
