package email

import (
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
)

type Verification struct {
	AddressID uint      `json:"address_id"`
	Code      uuid.UUID `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func DecodeVerification(readCloser io.ReadCloser) (verification Verification, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&verification)
	return
}
