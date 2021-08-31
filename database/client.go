package database

type Client interface {
	Close() error
	Ping() error
	Query(string, ...interface{}) (Rows, error)
}
