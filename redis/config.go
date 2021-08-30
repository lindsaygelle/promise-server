package redis

type Config interface {
	// Db returns the database value.
	Db() int
	// Host returns hostname value.
	Host() string
	// Password returns the password value.
	Password() string
	// Port returns the port value.
	Port() string
}

type config struct {
	host,
	password,
	port string
}

func (v config) Db() int          { return 0 }
func (v config) Host() string     { return v.host }
func (v config) Password() string { return v.password }
func (v config) Port() string     { return v.port }

// NewConfig returns a new Config.
func NewConfig() Config {
	c := &config{}
	setConfig(c)
	return c
}

// setConfig sets all fields for a config instance.
func setConfig(c *config) {
	for _, fn := range [](func(*config) error){
		setConfigHost,
		setConfigPassword,
		setConfigPort} {
		err := fn(c)
		if err != nil {
			panic(err)
		}
	}
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
