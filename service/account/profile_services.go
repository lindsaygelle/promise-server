package account

import (
	"net/http"

	"github.com/lindsaygelle/promise/promise-server/content/account"
)

func (p *profileService) Get(profileID string) (account.Profile, error) {
	return account.ScanProfile(nil)
}

func (p *profileService) GetAll() (account.Profiles, error) {
	return account.ScanProfiles(nil)
}

func (p *profileService) Make(r *http.Request) (profile account.Profile, err error) {
	profileCreate, err := account.DecodeProfileCreate(r.Body)
	if err != nil {
		return
	}
	tx, err := p.Begin()
	if err != nil {
		return
	}
	row := tx.QueryRow("SELECT * FROM account.profile_insert(?)", profileCreate.Name)
	profile, err = account.ScanProfile(row)
	return
}
