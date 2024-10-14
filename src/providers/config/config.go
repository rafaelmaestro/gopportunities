package config

import "time"

type HttpConfig struct {
	HttpPort string `env:"HTTP_PORT"`
}

type DbConfig struct {
	Driver                string        `env:"DRIVER"`
	Host                  string        `env:"HOST"`
	Port                  string        `env:"PORT"`
	Name                  string        `env:"NAME"`
	User                  string        `env:"USER"`
	Pass                  string        `env:"PASSWORD"`
	ConnectionRetries     int           `env:"CONNECTION_RETRIES"`
	MaxIdleConnections    int           `env:"MAX_IDLE_CONNECTIONS"`
	MaxOpenConnections    int           `env:"MAX_OPEN_CONNECTIONS"`
	ConnectionMaxLifetime time.Duration `env:"CONNECTION_MAX_LIFETIME"`
}

type KafkaConfig struct {
	Brokers                  string            `env:"BROKERS"`
	ConsumerMinBytes         string            `env:"CONSUMER_MIN_BYTES"`
	ConsumerMaxBytes         string            `env:"CONSUMER_MAX_BYTES"`
	HeartbeatInterval        int               `env:"HEARTBEAT_INTERVAL"`
	ConcurrentReaders        int               `env:"CONCURRENT_READERS"`
	GroupID                  string            `env:"GROUP_ID"`
	SessionTimeoutMultiplier int               `env:"SESSION_TIMEOUT_MULTIPLIER"`
	Topics                   map[string]string // Mapa de eventos e tópicos
}

type AppConfig struct {
	AppLogLevel  string `env:"LOG_LEVEL"`
	AppLogOutput string `env:"LOG_OUTPUT"`
}

type Config struct {
	Http  HttpConfig
	App   AppConfig   `envPrefix:"APP_"`
	Db    DbConfig    `envPrefix:"DB_"`
	Kafka KafkaConfig `envPrefix:"KAFKA_"`
}
