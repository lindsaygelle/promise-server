package promise

import (
	"database/sql"

	"github.com/lindsaygelle/promise/promise-server/content/promise"
)

type StepService interface {
	Get(stepID string) (promise.Step, error)
	GetAllByTask(taskID string) (promise.Steps, error)
}

type stepService struct {
	*sql.DB
}

func NewStepService(database *sql.DB) StepService {
	return &stepService{database}
}
