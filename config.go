package traefik_vault_auth

// Config for the plugin configuration.
type Config struct {
	Vault      Vault `yaml:"vault"`      // Vault remote server configuration
	CustomRealm string `yaml:"customRealm"` // CustomRealm can be used to personalize Basic Auth window message
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	c := &Config{}
	return c.addMissingFields()
}

func (c *Config) addMissingFields() *Config {
	if c.CustomRealm == "" {
		c.CustomRealm = "Use a valid Vault user to authenticate"
	}
	return c
}
