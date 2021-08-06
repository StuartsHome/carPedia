package config

type Config struct {
}

var _config *Config

func initDefault() *Config {
	return &Config{}
}
