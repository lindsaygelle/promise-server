package postgres

type Config interface {
	// Db returns the name of the configured database.
	Db() string
	// Host returns the hostname of the configured database.
	Host() string
	// Port returns the port address of the configured database.
	Port() string
	// SSLMode returns the SSL connection setting for the configured database.
	SSLMode() string
	// User returns the configured database user.
	User() string
}

// config handles the internal configuration for postgres connection.
type config struct {
	db,
	host,
	password,
	port,
	sslmode,
	user string
}

func (v config) Db() string      { return v.db }      // Db returns config.db
func (v config) Host() string    { return v.host }    // Host returns config.host
func (v config) Port() string    { return v.port }    // Port returns config.port
func (v config) SSLMode() string { return v.sslmode } // SSLMode returns config.sslmode
func (v config) User() string    { return v.user }    // User returns config.user

// NewConfig returns a new Config.
func NewConfig() Config {
	c := config{}
	setConfig(&c)
	return c
}

// setConfig sets all fields for a config instance.
func setConfig(c *config) {
	for _, fn := range [](func(*config) error){
		setConfigDatabase,
		setConfigHost,
		setConfigPassword,
		setConfigPort,
		setConfigSSLMode,
		setConfigUser} {
		err := fn(c)
		if err != nil {
			panic(err)
		}
	}
}

// setConfigDatabase sets the config.db value.
func setConfigDatabase(c *config) (err error) {
	c.db, err = getEnvDatabase()
	return err
}

// setConfigHost sets the config.host value.
func setConfigHost(c *config) (err error) {
	c.host, err = getEnvHost()
	return err
}

// setConfigPassword sets the config.password value.
func setConfigPassword(c *config) (err error) {
	c.password, err = getEnvPassword()
	return err
}

// setConfigPort sets the config.port value.
func setConfigPort(c *config) (err error) {
	c.port, err = getEnvPort()
	return err
}

// setConfigSSLMode sets the config.sslmode value.
func setConfigSSLMode(c *config) (err error) {
	c.sslmode, err = getEnvSSLMode()
	return err
}

// setConfigUser sets the config.user value.
func setConfigUser(c *config) (err error) {
	c.user, err = getEnvUser()
	return err
}
