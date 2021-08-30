package postgres

type Config interface {
	// Db returns the database value.
	Db() string
	// Host returns hostname value.
	Host() string
	// Password returns the password value.
	Password() string
	// Port returns the port value.
	Port() string
	// SSLMode returns SSL value.
	SSLMode() string
	// User returns user value.
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

func (v config) Db() string       { return v.db }
func (v config) Host() string     { return v.host }
func (v config) Password() string { return v.password }
func (v config) Port() string     { return v.port }
func (v config) SSLMode() string  { return v.sslmode }
func (v config) User() string     { return v.user }

// NewConfig returns a new Config.
func NewConfig() Config {
	c := &config{}
	setConfig(c)
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
