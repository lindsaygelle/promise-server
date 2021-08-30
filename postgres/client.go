package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Client interface {
	Ping() error
}

type client struct {
	*sql.DB
}

// NewClient returns a new Client.
func NewClient(c Config) Client {
	db, err := sql.Open(driver, newDriverSource(c))
	if err != nil {
		panic(err)
	}
	return &client{db}
}
