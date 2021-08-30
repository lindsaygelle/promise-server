package redis

import (
	"github.com/go-redis/redis"
)

type Client interface {
	Ping() *redis.StatusCmd
}

type client struct {
	*redis.Client
}

// NewClient returns a new Client.
func NewClient(c Config) Client {
	var (
		db       = c.Db()
		password = c.Password()
	)
	v := redis.NewClient(&redis.Options{
		Addr:     newDriverAddr(c),
		DB:       db,
		Password: password})
	return client{v}
}
