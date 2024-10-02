package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

type AppCfg struct {
	Env      string   `json:"env" yaml:"env" env:"ENV" env-default:"production" `
	Logger   Logger   `json:"logger" yaml:"logger" env-prefix:"LOGGER_"`
	HTTP     HTTP     `json:"http" yaml:"http" env-prefix:"HTTP_"`
	TLS      TLS      `json:"tls" yaml:"tls" env-prefix:"TLS_"`
	Postgres Postgres `json:"postgres" yaml:"postgres" env-prefix:"POSTGRES_"`
}

func (config AppCfg) IsLocal() bool {
	return config.Env == "local"
}

func NewAppCfg(filepath string) (*AppCfg, error) {
	var err error
	var c AppCfg
	if filepath == "" {
		err = cleanenv.ReadEnv(&c)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read env")
		}
	} else {
		err = cleanenv.ReadConfig(filepath, &c)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read config")
		}
	}
	return &c, err
}

// TODO: need a shareable way to define config patch between commands that is not under this module
func LoadConfig() *AppCfg {
	filepath := pflag.StringP("config", "c", "", "configuration filepath (default: None)")
	pflag.Parse()
	// ________________________________________________________________________
	// Load config
	cfg, err := NewAppCfg(*filepath)
	if err != nil {
		log.Fatalf("cannot load config: %s", err)
		panic("Could not load config. Shutting down")
	}
	return cfg
}
