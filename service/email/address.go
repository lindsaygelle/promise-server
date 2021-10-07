package email

import "database/sql"

type AddressService interface {
	Get(profileID string)
}

type addressService struct {
	*sql.DB
}

func NewAddressService(database *sql.DB) AddressService {
	return &addressService{database}
}
