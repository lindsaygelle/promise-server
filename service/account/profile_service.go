package account

import (
	"github.com/lindsaygelle/promise/promise-server/content/account"
)

func (p *profileService) Get(id string) (account.Profile, error) {
	return account.ScanProfile(nil)
}

func (p *profileService) GetAll() (account.Profiles, error) {
	return account.ScanProfiles(nil)
}
