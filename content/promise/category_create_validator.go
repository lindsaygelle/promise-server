package promise

type CategoryCreateValidator func(*CategoryCreate) error

var categoryCreateValidators = [...]CategoryCreateValidator{
	validateCategoryCreateName,
	validateCategoryCreateProfileID}

func validateCategoryCreate(categoryCreate *CategoryCreate) error {
	for _, validator := range categoryCreateValidators {
		if err := validator(categoryCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateCategoryCreateName(categoryCreate *CategoryCreate) error {
	if len(categoryCreate.Name) == 0 {
		return ErrCategoryName
	}
	return nil
}

func validateCategoryCreateProfileID(categoryCreate *CategoryCreate) error {
	if categoryCreate.ProfileID == 0 {
		return ErrCategoryProfileID
	}
	return nil
}
