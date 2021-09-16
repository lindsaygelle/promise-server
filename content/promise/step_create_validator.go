package promise

import "time"

type StepCreateValidator func(*StepCreate) error

var stepCreateValidators = [...]StepCreateValidator{
	validateStepCreateTaskID,
	validateStepCreateTaskTimeDue}

func validateStepCreate(stepCreate *StepCreate) error {
	for _, validator := range stepCreateValidators {
		if err := validator(stepCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateStepCreateTaskID(stepCreate *StepCreate) error {
	if stepCreate.TaskID == 0 {
		return ErrStepTaskID
	}
	return nil
}

func validateStepCreateTaskTimeDue(stepCreate *StepCreate) error {
	if stepCreate.TimeDue != nil && stepCreate.TimeDue.Before(time.Now()) {
		return ErrStepTimeDue
	}
	return nil
}
