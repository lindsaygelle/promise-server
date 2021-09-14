package location

import (
	"encoding/json"
	"io"
	"time"
)

type Country struct {
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
}

func DecodeCountry(readCloser io.ReadCloser) (country Country, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&country)
	return
}
