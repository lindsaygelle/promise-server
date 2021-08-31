package postgres

import (
	_ "github.com/lib/pq"

	"database/sql"

	"github.com/lindsaygelle/promise/promise-server/database"
)

type client struct {
	*sql.DB
}

func (c *client) Query(query string, arguments ...interface{}) (database.Rows, error) {
	v, err := c.DB.Query(query, arguments...)
	return &rows{v}, err
}

func (c *client) QueryRow(query string, arguments ...interface{}) database.Row {
	return &row{c.DB.QueryRow(query, arguments...)}
}

// NewClient returns a new Client.
func NewClient(c Config) database.Client {
	v, err := sql.Open(driver, newDriverSource(c))
	if err != nil {
		panic(err)
	}
	return &client{v}
}
