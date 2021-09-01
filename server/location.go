package server

import (
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/location"
)

// Country returns a location.Country.
func Country(client database.Client, id string) (location.Country, error) {
	return location.GetCountry(client, id)
}

// Countries returns a slice of location.Country.
func Countries(client database.Client) ([]location.Country, error) {
	return location.GetCountries(client)
}
