package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	General GeneralConfig
}

type GeneralConfig struct {
	Debug bool `envconfig:"DEBUG"`
}

func Get() *Config {
	cfg := Config{}
	cfgs := map[string]interface{}{
		"general": &cfg.General,
	}

	for p, c := range cfgs {
		if err := envconfig.Process(p, c); err != nil {
			log.Panicf("Config parse error: %v", err)
		}
	}

	return &cfg
}
