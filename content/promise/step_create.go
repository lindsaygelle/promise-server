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

type StepCreateValidator func(StepCreate) error

func DecodeStepCreate(readCloser io.ReadCloser) (stepCreate StepCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&stepCreate)
	if err != nil {
		return
	}
	err = verifyStepCreate(stepCreate)
	return
}

func verifyStepCreate(stepCreate StepCreate) error {
	for _, validator := range [...]StepCreateValidator{
		verifyStepCreateTaskID,
		verifyStepCreateTaskTimeDue} {
		if err := validator(stepCreate); err != nil {
			return err
		}
	}
	return nil
}

func verifyStepCreateTaskID(stepCreate StepCreate) error {
	if stepCreate.TaskID == 0 {
		return ErrStepTaskID
	}
	return nil
}

func verifyStepCreateTaskTimeDue(stepCreate StepCreate) error {
	if stepCreate.TimeDue != nil && stepCreate.TimeDue.Before(time.Now()) {
		return ErrStepTimeDue
	}
	return nil
}
