package server

import (
	"github.com/lindsaygelle/promise/promise-server/account"
	"github.com/lindsaygelle/promise/promise-server/database"
)

// AccountSetting returns a account.Setting.
func GetAccountSetting(client database.Client, id string) (account.Setting, error) {
	return account.GetSetting(client, id)
}

// AccountSettings returns a slice of account.Setting.
func GetAccountSettings(client database.Client) ([]account.Setting, error) {
	return account.GetSettings(client)
}

// GetAccount returns a account.Account.
func GetAccount(client database.Client, id string) (account.Account, error) {
	return account.GetAccount(client, id)
}

// Accounts returns a slice of account.Account.
func GetAccounts(client database.Client) ([]account.Account, error) {
	return account.GetAccounts(client)
}
