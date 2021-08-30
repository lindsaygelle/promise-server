package redis

import (
	"fmt"
	"os"
)

const (
	driverAddr = "%s:%s"
)

const (
	envHost     = "REDIS_HOST"
	envPassword = "REDIS_PASSWORD"
	envPort     = "REDIS_PORT"
)

const (
	errEnv = "%s"
)

// getEnv returns an variable from the environment.
// On the condition the value cannot be found an error is returned.
func getEnv(key string) (string, error) {
	env := os.Getenv(key)
	if len(env) == 0 {
		return env, fmt.Errorf(errEnv, key)
	}
	return env, nil
}

// getEnvHost returns the envHost value.
func getEnvHost() (string, error) {
	return getEnv(envHost)
}

// getEnvPassword returns the envPassword value.
func getEnvPassword() (string, error) {
	return getEnv(envPassword)
}

// getEnvPort returns the envPort value.
func getEnvPort() (string, error) {
	return getEnv(envPort)
}

// newDriverAddr returns a new driver address source string.
func newDriverAddr(c Config) string {
	var (
		host = c.Host()
		port = c.Port()
	)
	return fmt.Sprintf(driverAddr, host, port)
}
