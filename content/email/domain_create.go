package email

import (
	"encoding/json"
	"io"
)

type DomainCreate struct {
	Value string `json:"value"`
}

func DecodeDomainCreate(readCloser io.ReadCloser) (domainCreate DomainCreate, err error) {
	defer readCloser.Close()
	err = json.NewDecoder(readCloser).Decode(&domainCreate)
	if err != nil {
		err = ErrDomain
		return
	}
	err = validateDomainCreate(domainCreate)
	return
}
