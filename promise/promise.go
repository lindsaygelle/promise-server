package promise

import (
	"time"

	"github.com/lindsaygelle/promise/promise-server/database"
)

type Promise struct {
	Completed   database.NullTime `json:"completed"`
	Created     time.Time         `json:"created"`
	Description string            `json:"description"`
	Edited      time.Time         `json:"edited"`
	ID          uint              `json:"id"`
	Locked      bool              `json:"locked"`
	Maker       uint              `json:"maker"`
	Name        string            `json:"name"`
	Owner       uint              `json:"owner"`
}

// GetPromise returns a Promise.
func GetPromise(client database.Client, id string) (Promise, error) {
	row, err := client.QueryRow(`SELECT * FROM promise.promise WHERE id=$1`, id)
	if err != nil {
		return Promise{}, err
	}
	return NewPromise(row)
}

// GetPromises returns a slice of Promise.
func GetPromises(client database.Client) ([]Promise, error) {
	rows, err := client.Query(`SELECT * FROM promise.promise`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return getPromises(rows)
}

func GetPromisesMaker(client database.Client, id string) ([]Promise, error) {
	rows, err := client.Query(`SELECT * FROM promise.promise WHERE maker=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return getPromises(rows)
}

func GetPromisesOwner(client database.Client, id string) ([]Promise, error) {
	rows, err := client.Query(`SELECT * FROM promise.promise WHERE owner=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return getPromises(rows)
}

// NewPromise returns a new Promise.
//
// NewPromise returns an error on the condition it cannot correctly scan the database row.
func NewPromise(v database.Scanner) (promise Promise, err error) {
	err = v.Scan(&promise.Completed, &promise.Created, &promise.Description, &promise.Edited, &promise.ID, &promise.Locked, &promise.Maker, &promise.Name, &promise.Owner)
	return promise, err
}

// addPromise scans a promise.Promise record from the database rows and adds it to the collection.
//
// addPromise returns an error on the condition the account cannot be scanned.
func addPromise(v *[]Promise, rows database.Rows) error {
	Promise, err := NewPromise(rows)
	if err != nil {
		return err
	}
	*v = append(*v, Promise)
	return err
}

func getPromises(rows database.Rows) ([]Promise, error) {
	var promises []Promise
	err := processPromises(&promises, rows)
	if err != nil {
		return nil, err
	}
	return promises, nil
}

func processPromises(promises *[]Promise, rows database.Rows) error {
	for rows.Next() {
		err := addPromise(promises, rows)
		if err != nil {
			return err
		}
	}
	return nil
}
