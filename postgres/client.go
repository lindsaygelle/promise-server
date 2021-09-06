package postgres

import (
	_ "github.com/lib/pq"

	"database/sql"
)

func NewClient(c Config) *sql.DB {
	database, err := sql.Open(driver, newDriverSource(c))
	if err != nil {
		panic(err)
	}
	return database
}
