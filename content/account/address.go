package account

import (
	"encoding/json"
	"io"
	"time"
)

type Address struct {
	CityID    uint      `json:"city_id"`
	CountryID uint      `json:"country_id"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
	ID        uint      `json:"id"`
	ProfileID uint      `json:"profile_id"`
}

func DecodeAddress(readCloser io.ReadCloser) (address Address, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&address)
	return
}
