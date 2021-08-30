package postgres

import "database/sql"

type Client interface {
	Ping() error
}

type client struct {
	*sql.DB
}

func NewClient(c Config) {

}
