package promise

import (
	"database/sql"

	"github.com/lindsaygelle/promise/promise-server/content/promise"
)

type TaskService interface {
	Get(taskID string) (promise.Task, error)
	GetAllByProfile(profileID string) (promise.Tasks, error)
}

type taskService struct {
	*sql.DB
}

func NewTaskService(database *sql.DB) TaskService {
	return &taskService{database}
}
