package server

import (
	"github.com/lindsaygelle/promise/promise-server/database"
	"github.com/lindsaygelle/promise/promise-server/location"
)

// Countries returns a slice of location.Country.
func Countries(v database.Client) ([]location.Country, error) {
	var countries []location.Country
	rows, err := v.Query(`SELECT alpha_2, alpha_3, id, name, numeric FROM location.country`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = addCountry(&countries, rows)
		if err != nil {
			return nil, err
		}
	}
	return countries, nil
}

// addCountry scans a country record from the database rows and adds it to the collection.
//
// addCountry returns an error on the condition the country cannot be scanned.
func addCountry(v *[]location.Country, rows database.Rows) error {
	var country location.Country
	err := scanCountry(&country, rows)
	if err != nil {
		return err
	}
	*v = append(*v, country)
	return err
}

// scanCountry scans a row from the database and sets the field.
func scanCountry(v *location.Country, rows database.Rows) error {
	return rows.Scan(&v.Alpha2, &v.Alpha3, &v.ID, &v.Name, &v.Numeric)
}
