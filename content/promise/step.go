package promise

import "time"

type Step struct {
	CreatedAt   time.Time `json:"created_at"`
	Description *string   `json:"description"`
	EditedAt    time.Time `json:"edited_at"`
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	TaskID      uint      `json:"task_id"`
}
