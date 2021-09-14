package location

import (
	"encoding/json"
	"io"
)

type CountryCreate struct {
	Name string `json:"name"`
}

func DecodeCountryCreate(readCloser io.ReadCloser) (countryCreate CountryCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&countryCreate)
	return
}
