package account

import (
	"encoding/json"
	"io"
	"time"
)

type Profile struct {
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
}

func DecodeProfile(readCloser io.ReadCloser) (profile Profile, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&profile)
	return
}
