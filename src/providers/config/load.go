package config

import (
	"os"
	"sync/atomic"

	"github.com/caarlos0/env/v6"
	"github.com/fnunezzz/go-logger"
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
// Usage: change the value below to the correct app name
func Init() *Config {
	cfg := Config{}
	load("gopportunities", &cfg)
	_global.Store(cfg)
	return &cfg
}

// Load will map all the envs to the reference struct
func load(appName string, reference *Config) {
	sLog := logger.Get()

	if len(appName) > 0 {
		os.Setenv("APP_NAME", appName)
	}

	if err := env.Parse(reference); err != nil {
		sLog.Errorf("error parsing configs - %+v", err)
	}
}
