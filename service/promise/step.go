package promise

import (
	"database/sql"

	"github.com/lindsaygelle/promise/promise-server/content/promise"
)

type StepService interface {
	Get(stepID string) (promise.Step, error)
	GetAll() (promise.Steps, error)
	GetAllByTask(taskID string) (promise.Steps, error)
}

type stepService struct {
	*sql.DB
}
