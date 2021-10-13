package account

import (
	"net/http"

	"github.com/lindsaygelle/promise/promise-server/content/account"
)

func (p *profileService) Get(profileID string) (profile account.Profile, err error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return
	}
	row := tx.QueryRow("SELECT ROW_TO_JSON(T.*) FROM account.profile WHERE id = ?", profileID)
	profile, err = account.ScanProfile(row)
	return
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
