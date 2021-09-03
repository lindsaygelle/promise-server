package location

import "github.com/lindsaygelle/promise/promise-server/database"

type Country struct {
	Alpha2  string `json:"alpha_2"`
	Alpha3  string `json:"alpha_3"`
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Numeric uint8  `json:"numeric"`
}

// GetCountry returns a location.Country.
func GetCountry(client database.Client, id string) (Country, error) {
	row, err := client.QueryRow(`SELECT alpha_2, alpha_3, id, name, numeric FROM location.country WHERE id=$1`, id)
	if err != nil {
		return Country{}, err
	}
	return NewCountry(row)
}

// GetCountries returns a slice of location.Country.
func GetCountries(client database.Client) ([]Country, error) {
	var countries []Country
	rows, err := client.Query(`SELECT alpha_2, alpha_3, id, name, numeric FROM location.country`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	err = processCountries(&countries, rows)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

// NewCountry returns a new country.Country.
//
// NewCountry returns an error on the condition it cannot correctly scan the database row.
func NewCountry(v database.Scanner) (country Country, err error) {
	err = v.Scan(&country.Alpha2, &country.Alpha3, &country.ID, &country.Name, &country.Numeric)
	return country, err
}

// addCountry scans a location.country record from the database rows and adds it to the collection.
//
// addCountry returns an error on the condition the country cannot be scanned.
func addCountry(v *[]Country, rows database.Rows) error {
	country, err := NewCountry(rows)
	if err != nil {
		return err
	}
	*v = append(*v, country)
	return nil
}

func processCountries(countries *[]Country, rows database.Rows) error {
	for rows.Next() {
		err := addCountry(countries, rows)
		if err != nil {
			return err
		}
	}
	return nil
}
