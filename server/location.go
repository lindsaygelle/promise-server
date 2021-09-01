package server

import (
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/location"
)

// GetCountry returns a location.Country.
func GetCountry(client database.Client, id string) (location.Country, error) {
	return location.GetCountry(client, id)
}

// GetCountries returns a slice of location.Country.
func GetCountries(client database.Client) ([]location.Country, error) {
	return location.GetCountries(client)
}
