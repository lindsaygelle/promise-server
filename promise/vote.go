package promise

import (
	"time"

	"github.com/lindsaygelle/promise/promise-server/database"
)

type Vote struct {
	Account uint      `json:"account"`
	Created time.Time `json:"created"`
	Promise uint      `json:"promise"`
	Value   bool      `json:"value"`
}

// GetVote returns a Vote.
func GetVote(client database.Client, id string) (Vote, error) {
	row, err := client.QueryRow(`SELECT account, created, promise, value promise.vote WHERE promise=$1`, id)
	if err != nil {
		return Vote{}, err
	}
	return NewVote(row)
}

// GetVotes returns a slice of Vote.
func GetVotes(client database.Client) ([]Vote, error) {
	var votes []Vote
	rows, err := client.Query(`SELECT account, created, promise, value FROM promise.vote`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	err = processVotes(&votes, rows)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

// NewVote returns a new Vote.
//
// NewVote returns an error on the condition it cannot correctly scan the database row.
func NewVote(v database.Scanner) (vote Vote, err error) {
	err = v.Scan(&vote.Account, &vote.Created, &vote.Promise, &vote.Value)
	return vote, err
}

// addVote scans a promise.vote record from the database rows and adds it to the collection.
//
// addVote returns an error on the condition the account cannot be scanned.
func addVote(v *[]Vote, rows database.Rows) error {
	vote, err := NewVote(rows)
	if err != nil {
		return err
	}
	*v = append(*v, vote)
	return err
}

func processVotes(votes *[]Vote, rows database.Rows) error {
	for rows.Next() {
		err := addVote(votes, rows)
		if err != nil {
			return err
		}
	}
	return nil
}
