package promise

import (
	"encoding/json"
	"io"
	"time"
)

type StepCreate struct {
	Description *string    `json:"description"`
	Name        string     `json:"name"`
	TaskID      uint       `json:"task_id"`
	TimeDue     *time.Time `json:"time_due"`
}

func DecodeStepCreate(readCloser io.ReadCloser) (stepCreate StepCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&stepCreate)
	if err != nil {
		err = ErrStep
		return
	}
	err = validateStepCreate(stepCreate)
	return
}
