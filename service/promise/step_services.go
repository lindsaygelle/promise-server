package promise

import "github.com/lindsaygelle/promise/promise-server/content/promise"

func (s *stepService) Get(stepID string) (promise.Step, error) {
	return promise.ScanStep(nil)
}

func (s *stepService) GetAllByTask(taskID string) (promise.Steps, error) {
	return promise.ScanSteps(nil)
}
