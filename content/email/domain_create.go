package email

import (
	"encoding/json"
	"io"
)

type DomainCreate struct{}

func DecodeDomainCreate(readCloser io.ReadCloser) (domainCreate DomainCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&domainCreate)
	return
}
