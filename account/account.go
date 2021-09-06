package account

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/mail"
	"time"

	"github.com/lindsaygelle/promise/promise-server/database"
)

type Account struct {
	AccountName

	Created time.Time `json:"created"`
	Edited  time.Time `json:"edited"`
	ID      uint      `json:"id"`
}

type AccountMake struct {
	AccountName

	Email string `json:"email"`
}

type AccountName struct {
	Name string `json:"name"`
}

type AccountMakeValidator func(AccountMake) error

var (
	ErrNoEmail = errors.New("email empty")
	ErrNoName  = errors.New("name empty")
)

var (
	accountMakeValidators = [...]AccountMakeValidator{
		validateAccountMakeEmail,
		validateAccountMakeName}
)

// NewAccount returns a new Account from the database.
func NewAccount(scanner database.Scanner) (account Account, err error) {
	err = scanner.Scan(&account.Created, &account.Edited, &account.ID, &account.Name)
	return
}

// NewAccountMake returns a new AccountMake from an io.ReadCloser.
func NewAccountMake(reader io.ReadCloser) (accountMake AccountMake, err error) {
	err = json.NewDecoder(reader).Decode(&accountMake)
	return
}

// NewAccountName returns a new AccountName from an io.ReadCloser.
func NewAccountName(reader io.ReadCloser) (accountName AccountName, err error) {
	err = json.NewDecoder(reader).Decode(&accountName)
	return
}

// EditAccountName edits an account name by its ID.
func EditAccountName(database *sql.DB, ID string, request *http.Request) (err error) {
	defer request.Body.Close()
	accountName, err := NewAccountName(request.Body)
	if err != nil {
		return
	}
	err = validateAccountName(accountName)
	if err != nil {
		return
	}
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	_, err = transaction.Exec("UPDATE account.account SET edited=CURRENT_TIME, name=$1 WHERE account.id=$2", accountName.Name, ID)
	if err != nil {
		transaction.Rollback()
	}
	transaction.Commit()
	return
}

// ReadAccount reads an account from the database by its ID.
func ReadAccount(database *sql.DB, ID string) (account Account, err error) {
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	row := transaction.QueryRow(`SELECT created, edited, id, name FROM account.account WHERE id=$1`, ID)
	account, err = NewAccount(row)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

// ReadAccounts reads all accounts from the database.
func ReadAccounts(database *sql.DB) (accounts []Account, err error) {
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	rows, err := transaction.Query(`SELECT created, id, name FROM account.account`)
	if err != nil {
		transaction.Rollback()
		return
	}
	err = processAccounts(&accounts, rows)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

// WriteAccount writes a new Account to the database.
func WriteAccount(database *sql.DB, request *http.Request) (account Account, err error) {
	defer request.Body.Close()
	accountMake, err := NewAccountMake(request.Body)
	if err != nil {
		return
	}
	err = validateAccountMake(accountMake)
	if err != nil {
		return
	}
	transaction, err := database.Begin()
	if err != nil {
		return
	}
	row := transaction.QueryRow(`INSERT INTO account.account (email, name) VALUES ($1, $2) RETURNING created, edited, id, name`, accountMake.Email, accountMake.Name)
	account, err = NewAccount(row)
	if err != nil {
		transaction.Rollback()
		return
	}
	transaction.Commit()
	return
}

// addAccount scans a account.account record from the database rows and adds it to the collection.
//
// addAccount returns an error on the condition the account cannot be scanned.
func addAccount(accounts *[]Account, rows *sql.Rows) (err error) {
	account, err := NewAccount(rows)
	if err != nil {
		return err
	}
	*accounts = append(*accounts, account)
	return err
}

// processAccounts processes all rows from the database and adds them to the collection.
//
// processAccounts returns an error on the condition the row cannot be added.
func processAccounts(accounts *[]Account, rows *sql.Rows) (err error) {
	for rows.Next() {
		err = addAccount(accounts, rows)
		if err != nil {
			return
		}
	}
	return
}

// validateAccountMake validates an AccountMake.
func validateAccountMake(accountMake AccountMake) (err error) {
	for _, fn := range accountMakeValidators {
		err = fn(accountMake)
		if err != nil {
			break
		}
	}
	return err
}

// validateAccountMakeEmail validates AccountMake.Email.
func validateAccountMakeEmail(accountMake AccountMake) (err error) {
	if len(accountMake.Email) == 0 {
		return ErrNoEmail
	}
	_, err = mail.ParseAddress(accountMake.Email)
	return
}

// validateAccountMakeName validates Accountmake.Name.
func validateAccountMakeName(accountMake AccountMake) (err error) {
	if len(accountMake.Name) == 0 {
		err = ErrNoName
	}
	return
}

func validateAccountName(accountName AccountName) (err error) {
	if len(accountName.Name) == 0 {
		err = ErrNoName
	}
	return
}
