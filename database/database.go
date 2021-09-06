package database

import "database/sql"

const (
	null = `null`
)

type Database interface {
	Begin() *sql.Tx
}
