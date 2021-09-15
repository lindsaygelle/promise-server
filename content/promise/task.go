package promise

import (
	"encoding/json"
	"io"
	"time"
)

type Task struct {
	CategoryID    uint       `json:"category_id"`
	Description   *string    `json:"description"`
	ID            uint       `json:"id"`
	IsCompleted   bool       `json:"is_completed"`
	IsLocked      bool       `json:"is_locked"`
	ProfileID     uint       `json:"profile_id"`
	Name          string     `json:"name"`
	TimeCompleted *time.Time `json:"time_completed"`
	TimeCreated   time.Time  `json:"time_created"`
	TimeEdited    time.Time  `json:"time_edited"`
}

type TaskValidator func(Task) error

func DecodeTask(readCloser io.ReadCloser) (task Task, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&task)
	if err != nil {
		err = ErrTask
		return
	}
	return
}
