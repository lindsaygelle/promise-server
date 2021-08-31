package database

type Row interface {
	Scan(...interface{}) error
}
