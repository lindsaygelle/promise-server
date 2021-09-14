package location

import (
	"encoding/json"
	"io"
	"time"
)

type City struct {
	CountryID uint      `json:"country_id"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
}

func DecodeCity(readCloser io.ReadCloser) (city City, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&city)
	return
}
