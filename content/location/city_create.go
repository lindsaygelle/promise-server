package location

import (
	"encoding/json"
	"io"
)

type CityCreate struct {
	CountryID uint `json:"country_id"`
}

func DecodeCityCreate(readCloser io.ReadCloser) (cityCreate CityCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&cityCreate)
	return
}
