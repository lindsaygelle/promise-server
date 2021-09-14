package email

import (
	"encoding/json"
	"io"
	"time"
)

type Domain struct {
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
	ID        uint      `json:"id"`
	Value     string    `json:"value"`
}

func DecodeDomain(readCloser io.ReadCloser) (domain Domain, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&domain)
	return
}
