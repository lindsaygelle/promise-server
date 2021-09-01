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

// GetPromise returns a promise.Promise.
func GetPromise(client database.Client, id string) (Promise, error) {
	row, err := client.QueryRow(`SELECT * FROM promise.promise WHERE id=$1`, id)
	if err != nil {
		return Promise{}, err
	}
	return NewPromise(row)
}

// GetPromises returns a slice of promise.Promise.
func GetPromises(v database.Client) ([]Promise, error) {
	var Promises []Promise
	rows, err := v.Query(`SELECT * FROM promise.promise`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = addPromise(&Promises, rows)
		if err != nil {
			return nil, err
		}
	}
	return Promises, nil
}

// NewPromise returns a new promise.promise.
//
// NewPromise returns an error on the condition it cannot correctly scan the database row.
func NewPromise(v interface{ Scan(...interface{}) error }) (promise Promise, err error) {
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