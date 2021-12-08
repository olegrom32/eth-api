package config

import (
	"github.com/kelseyhightower/envconfig"
)

type RESTServer struct {
	Addr string `envconfig:"REST_ADDR"`
}

type Ropsten struct {
	URL      string `envconfig:"ROPSTEN_URL"`
	Contract string `envconfig:"ROPSTEN_CONTRACT"`
}

type Config struct {
	RESTServer RESTServer
	Ropsten    Ropsten
}

func NewConfig() (*Config, error) {
	cfg := new(Config)

	return cfg, envconfig.Process("", cfg)
}
