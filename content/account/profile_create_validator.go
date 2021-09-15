package account

import "net/mail"

type ProfileCreateValidator func(ProfileCreate) error

var profileCreateValidators = [...]ProfileCreateValidator{
	validateProfileCreateEmail,
	validateProfileCreateName}

func validateProfileCreate(profileCreate ProfileCreate) error {
	for _, validator := range profileCreateValidators {
		if err := validator(profileCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateProfileCreateEmail(profileCreate ProfileCreate) error {
	_, err := mail.ParseAddress(profileCreate.Email)
	if err != nil {
		return ErrProfileEmail
	}
	return nil
}

func validateProfileCreateName(profileCreate ProfileCreate) error {
	if len(profileCreate.Name) == 0 {
		return ErrProfileName
	}
	return nil
}
