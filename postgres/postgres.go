package postgres

import (
	"fmt"
	"os"
)

const (
	driver       = "postgres"
	driverSource = "dbname=%s host=%s password=%s port=%s sslmode=%s user=%s"
)

const (
	envDatabase = "POSTGRES_DB"
	envHost     = "POSTGRES_HOST"
	envPassword = "POSTGRES_PASSWORD"
	envPort     = "POSTGRES_PORT"
	envSSLMode  = "POSTGRES_SSL"
	envUser     = "POSTGRES_USER"
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

// getEnvDatabase returns the envDatabase value.
func getEnvDatabase() (string, error) {
	return getEnv(envDatabase)
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

// getEnvSSLMode returns the envSSLMode value.
func getEnvSSLMode() (string, error) {
	return getEnv(envSSLMode)
}

// getEnvUser returns the envUser value.
func getEnvUser() (string, error) {
	return getEnv(envUser)
}

// newDriverSource returns a new database driver source string.
func newDriverSource(c Config) string {
	var (
		db       = c.Db()
		host     = c.Host()
		password = c.Password()
		port     = c.Port()
		sslmode  = c.SSLMode()
		user     = c.User()
	)
	return fmt.Sprintf(driverSource, db, host, password, port, sslmode, user)
}
