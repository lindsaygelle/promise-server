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
	stmt, err := c.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	v, err := stmt.Query(arguments...)
	if err != nil {
		return nil, err
	}
	return &rows{v}, err
}

func (c *client) QueryRow(query string, arguments ...interface{}) (database.Row, error) {
	stmt, err := c.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	v := stmt.QueryRow(arguments...)
	return &row{v}, nil
}

// NewClient returns a new Client.
func NewClient(c Config) database.Client {
	v, err := sql.Open(driver, newDriverSource(c))
	if err != nil {
		panic(err)
	}
	return &client{v}
}
