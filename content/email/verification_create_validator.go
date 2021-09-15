package email

type VerificationCreateValidator func(*VerificationCreate) error

var verificationCreateValidator = [...]VerificationCreateValidator{
	validateVerificationCreateAddressID}

func validateVerificationCreate(verificationCreate *VerificationCreate) error {
	for _, validator := range verificationCreateValidator {
		if err := validator(verificationCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateVerificationCreateAddressID(verificationCreate *VerificationCreate) error {
	if verificationCreate.AddressID == 0 {
		return ErrVerificationAddressID
	}
	return nil
}
