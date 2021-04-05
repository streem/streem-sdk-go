package config

// config holds key and environment used for creating and signing jws token.
type Config struct {
	ApiKeyId       string
	ApiKeySecret   string
	ApiEnvironment string
}

var instance *Config

// NewConfig generates a config struct and returns the instance of it
func NewConfig(apiKeyId, apiKeySecret, apiEnvironment string) *Config {
	instance = &Config{
		apiKeyId,
		apiKeySecret,
		apiEnvironment,
	}

	return instance
}

// Get returns the current config
func Get() *Config {
	return instance
}
