package email

type DomainValidator func(*Domain) error

func validateDomain(*Domain) error {
	return nil
}
