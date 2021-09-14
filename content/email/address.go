package email

import (
	"encoding/json"
	"io"
	"time"
)

type Address struct {
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	DomainID  string    `json:"domain_id"`
	ID        uint      `json:"id"`
	Verified  bool      `json:"verified"`
}

func DecodeAddress(readCloser io.ReadCloser) (address Address, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&address)
	return
}
