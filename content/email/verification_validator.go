package email

type VerificationValidator func(*Verification) error

func validateVerification(*Verification) error {
	return nil
}
