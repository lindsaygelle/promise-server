package email

import "net/url"

type DomainCreateValidator func(DomainCreate) error

var domainCreateValidators = [...]DomainCreateValidator{
	validateProfileCreateValue}

func validateDomainCreate(domainCreate DomainCreate) error {
	for _, validator := range domainCreateValidators {
		if err := validator(domainCreate); err != nil {
			return err
		}
	}
	return nil
}

func validateProfileCreateValue(domainCreate DomainCreate) error {
	_, err := url.Parse(domainCreate.Value)
	if err != nil {
		return ErrDomainValue
	}
	return nil
}
