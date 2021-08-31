package database

type Rows interface {
	Close() error
	Next() bool
	Scan(v ...interface{}) error
}
