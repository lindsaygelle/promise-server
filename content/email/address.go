package email

import (
	"encoding/json"
	"io"
	"time"
)

type Address struct {
	Address      string     `json:"address"`
	DomainID     string     `json:"domain_id"`
	ID           uint       `json:"id"`
	TimeCreated  time.Time  `json:"time_created"`
	TimeVerified *time.Time `json:"time_verified"`
	Verified     bool       `json:"verified"`
}

func DecodeAddress(readCloser io.ReadCloser) (address Address, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&address)
	return
}
