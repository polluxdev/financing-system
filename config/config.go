package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Tz             string `yaml:"tz"`
		AppConfig      `yaml:"app"`
		HTTPConfig     `yaml:"http"`
		LogConfig      `yaml:"log"`
		PostgresConfig `yaml:"postgres"`
	}

	AppConfig struct {
		Name    string `yaml:"name" env:"APP_NAME" env-required:"true"`
		Version string `yaml:"version" env:"APP_VERSION" env-required:"true"`
		Env     string `yaml:"env" env:"APP_ENV" env-required:"true"`
		Host    string `yaml:"host" env:"APP_HOST" env-required:"true"`
	}

	HTTPConfig struct {
		Port            string `yaml:"port" env:"HTTP_PORT" env-required:"true"`
		ReadTimeout     int    `yaml:"read_timeout" env:"HTTP_READ_TIMEOUT" env-required:"true"`
		WriteTimeout    int    `yaml:"write_timeout" env:"HTTP_WRITE_TIMEOUT" env-required:"true"`
		ShutdownTimeout int    `yaml:"shutdown_timeout" env:"HTTP_SHUTDOWN_TIMEOUT" env-required:"true"`
	}

	LogConfig struct {
		Level string `yaml:"log_level" env:"LOG_LEVEL" env-required:"true"`
	}

	PostgresConfig struct {
		User         string `yaml:"user" env:"POSTGRES_USER" env-required:"true"`
		Password     string `yaml:"password" env:"POSTGRES_PASSWORD" env-required:"true"`
		Host         string `yaml:"host" env:"POSTGRES_HOST" env-required:"true"`
		Port         int    `yaml:"port" env:"POSTGRES_PORT" env-required:"true"`
		Database     string `yaml:"database" env:"POSTGRES_DATABASE" env-required:"true"`
		SSLMode      string `yaml:"ssl_mode" env:"POSTGRES_SSL_MODE" env-required:"true"`
		TimeZone     string `yaml:"timezone" env:"POSTGRES_TIMEZONE" env-required:"true"`
		MaxPoolSize  int    `yaml:"max_pool_size" env:"POSTGRES_MAX_POOL_SIZE" env-required:"true"`
		ConnAttempts int    `yaml:"conn_attempts" env:"POSTGRES_CONN_ATTEMPTS" env-required:"true"`
		ConnTimeout  int    `yaml:"conn_timeout" env:"POSTGRES_CONN_TIMEOUT" env-required:"true"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
