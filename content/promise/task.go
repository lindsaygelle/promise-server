package promise

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"time"
)

type Task struct {
	TaskID        uint       `json:"category_id"`
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

func DecodeTask(readCloser io.ReadCloser) (task Task, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&task)
	if err != nil {
		err = ErrTask
		return
	}
	return
}

func ScanTask(scanner interface{ Scan(...interface{}) error }) (Task, error) {
	var b []byte
	err := scanner.Scan(&b)
	if err == sql.ErrNoRows {
		return Task{}, ErrTaskNotFound
	}
	return DecodeTask(io.NopCloser(bytes.NewReader(b)))
}
