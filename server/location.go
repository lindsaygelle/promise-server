package server

import (
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/location"
)

// Countries returns a slice of location.Country.
func Countries(c database.Client) ([]location.Country, error) {
	return location.GetCountries(c)
}
