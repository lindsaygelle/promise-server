package server

import (
	"github.com/lindsaygelle/promise/promise-server/account"
	"github.com/lindsaygelle/promise/promise-server/database"
)

// AccountSettings returns a slice of account.Setting.
func AccountSettings(client database.Client) ([]account.Setting, error) {
	return account.GetSettings(client)
}

// Accounts returns a slice of account.Account.
func Accounts(client database.Client) ([]account.Account, error) {
	return account.GetAccounts(client)
}
