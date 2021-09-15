package promise

import (
	"encoding/json"
	"io"
	"time"
)

type Step struct {
	Description   *string    `json:"description"`
	ID            uint       `json:"id"`
	IsCompleted   bool       `json:"is_completed"`
	IsLocked      bool       `json:"is_locked"`
	Name          string     `json:"name"`
	StatusID      uint       `json:"status_id"`
	TaskID        uint       `json:"task_id"`
	TimeCompleted *time.Time `json:"time_completed"`
	TimeCreated   time.Time  `json:"time_created"`
	TimeDue       time.Time  `json:"time_due"`
	TimeEdited    time.Time  `json:"time_edited"`
}

func DecodeStep(readCloser io.ReadCloser) (step Step, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&step)
	if err != nil {
		err = ErrStep
		return
	}
	return
}
