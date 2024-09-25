package config

import (
	"log"
	"os"
	"sync/atomic"

	"github.com/caarlos0/env/v6"
)

type Option func(*env.Options)

// These constants help create conditions to specific environments
const (
	Production  = "production"
	Staging     = "staging"
	Development = "development"
	Test        = "test"
)

var (
	_global atomic.Value
)

// Get return the config loaded on the start of the application
func Get() Config {
	return _global.Load().(Config)
}

// Init will start all the environments variable into the Config struct
func Init() *Config {
	cfg := Config{}
	load("api-test", &cfg)
	_global.Store(cfg)
	return &cfg
}

// Load will map all the envs to the reference struct
func load(appName string, reference *Config) {
	if len(appName) > 0 {
		os.Setenv("APP_NAME", appName)
	}

	if err := env.Parse(reference); err != nil {
		log.Panicf("error parsing configs - %+v", err)
	}
}
