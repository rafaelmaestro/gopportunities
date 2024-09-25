package config

type HttpConfig struct {
	AppPort string `env:"APP_PORT"`
}

type DbConfig struct {
	Driver            string `env:"DRIVER"`
	Host              string `env:"HOST"`
	Port              string `env:"PORT"`
	Name              string `env:"NAME"`
	User              string `env:"USER"`
	Pass              string `env:"PASSWORD"`
	ConnectionRetries int    `env:"CONNECTION_RETRIES"`
}

type KafkaConfig struct {
	Brokers           string `env:"BROKERS"`
	ConnectionRetries int    `env:"CONNECTION_RETRIES"`
	ProducerRetries   int    `env:"PRODUCER_RETRIES"`
}

type Config struct {
	Http  HttpConfig
	Db    DbConfig    `envPrefix:"DB_"`
	Kafka KafkaConfig `envPrefix:"KAFKA_"`
}
