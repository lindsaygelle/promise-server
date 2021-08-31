package account

import (
	"time"

	"github.com/lindsaygelle/promise/promise-server/database"
)

type Account struct {
	Created time.Time `json:"created"`
	ID      uint      `json:"id"`
	Name    string    `json:"name"`
}

// GetAccounts returns a slice of account.Account.
func GetAccounts(client database.Client) ([]Account, error) {
	var accounts []Account
	rows, err := client.Query(`SELECT created, id, name FROM account.account`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = addAccount(&accounts, rows)
		if err != nil {
			return nil, err
		}
	}
	return accounts, nil
}

// NewAccount returns a new account.Account.
//
// NewAccount returns an error on the condition it cannot correctly scan the database row.
func NewAccount(rows database.Rows) (account Account, err error) {
	err = rows.Scan(&account.Created, &account.ID, &account.Name)
	return account, err
}

// addAccount scans a account.account record from the database rows and adds it to the collection.
//
// addAccount returns an error on the condition the account cannot be scanned.
func addAccount(v *[]Account, rows database.Rows) error {
	account, err := NewAccount(rows)
	if err != nil {
		return err
	}
	*v = append(*v, account)
	return err
}
