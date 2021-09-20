package promise

import "github.com/lindsaygelle/promise/promise-server/content/promise"

func (t *taskService) Get(taskID string) (promise.Task, error) {
	return promise.ScanTask(nil)
}

func (t *taskService) GetAllByProfile(profileID string) (promise.Tasks, error) {
	return promise.ScanTasks(nil)
}
