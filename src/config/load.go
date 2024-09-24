package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func getViper() *viper.Viper {
	v := viper.New()
    viper.SetConfigName("app")
    viper.SetConfigType("env")
	viper.AutomaticEnv()
	return v
}

func NewConfig() (*Config, error) {
	fmt.Println("Loading config module")
	v := getViper()
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
