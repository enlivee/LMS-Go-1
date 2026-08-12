package config

type Config struct {
	DefaultBalance float64
}

func NewConfig() *Config {
	c := Config{DefaultBalance: 1000}
	return &c
}