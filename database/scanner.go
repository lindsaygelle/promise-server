package database

type Scanner interface {
	Scan(...interface{}) error
}
