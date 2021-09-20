package account

import (
	"database/sql"

	"github.com/lindsaygelle/promise/promise-server/content/account"
)

type ProfileService interface {
	Get(profileID string) (account.Profile, error)
	GetAll() (account.Profiles, error)
}

type profileService struct {
	*sql.DB
}

func NewProfileService(database *sql.DB) ProfileService {
	return &profileService{database}
}
