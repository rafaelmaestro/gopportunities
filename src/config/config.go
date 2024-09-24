package config

type httpConfig struct {
	AppPort string `mapstructure:"APP_PORT"`
}

type dbConfig struct {
	Driver            string `mapstructure:"DB_DRIVER"`
	Host              string `mapstructure:"DB_URI"`
	Port              string `mapstructure:"DB_PORT"`
	Name              string `mapstructure:"DB_NAME"`
	User              string `mapstructure:"DB_USER"`
	Pass              string `mapstructure:"DB_PASS"`
	ConnectionRetries int    `mapstructure:"DB_CONNECTION_RETRIES"`
}

type Config struct {
	Http httpConfig
	Db   dbConfig
}
