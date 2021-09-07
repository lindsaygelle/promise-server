package redis

import (
	"github.com/go-redis/redis"
)

// NewClient returns a new Client.
func NewClient(c Config) *redis.Client {
	var (
		db       = c.Db()
		password = c.Password()
	)
	v := redis.NewClient(&redis.Options{
		Addr:     newDriverAddr(c),
		DB:       db,
		Password: password})
	return v
}
