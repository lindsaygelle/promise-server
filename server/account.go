package server

import (
	"github.com/lindsaygelle/promise/promise-server/account"
	"github.com/lindsaygelle/promise/promise-server/database"
)

// Accounts returns a slice of account.Account.
func Accounts(v database.Client) ([]account.Account, error) {
	var accounts []account.Account
	rows, err := v.Query(`SELECT created, id, name FROM account.account`)
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

// addAccount scans a account record from the database rows and adds it to the collection.
//
// addAccount returns an error on the condition the account cannot be scanned.
func addAccount(v *[]account.Account, rows database.Rows) error {
	var account account.Account
	err := scanAccount(&account, rows)
	if err != nil {
		return err
	}
	*v = append(*v, account)
	return err
}

// scanAccount scans a row from the database and sets the field.
func scanAccount(v *account.Account, rows database.Rows) error {
	return rows.Scan(&v.Created, &v.ID, &v.Name)
}
