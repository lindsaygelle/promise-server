package promise

type TaskCreateValidator func(*TaskCreate) error

var taskCreateValidators = [...]TaskCreateValidator{
	validateTaskCreateCategory,
	validateTaskCreateName}

func validateTaskCreate(taskCreate *TaskCreate) error {
	for _, validator := range taskCreateValidators {
		if err := validator(taskCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateTaskCreateCategory(taskCreate *TaskCreate) error {
	if taskCreate.CategoryID != nil && *taskCreate.CategoryID == 0 {
		return ErrTaskCategoryID
	}
	return nil
}

func validateTaskCreateName(taskCreate *TaskCreate) error {
	if len(taskCreate.Name) == 0 {
		return ErrTaskName
	}
	return nil
}
