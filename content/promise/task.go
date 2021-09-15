package promise

import (
	"encoding/json"
	"io"
	"time"
)

type Task struct {
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	Description *string   `json:"description"`
	EditedAt    time.Time `json:"edited_at"`
	ID          uint      `json:"id"`
	ProfileID   uint      `json:"profile_id"`
	Name        string    `json:"name"`
}

func DecodeTask(readCloser io.ReadCloser) (task Task, err error) {
	err = json.NewDecoder(readCloser).Decode(&task)
	return
}
