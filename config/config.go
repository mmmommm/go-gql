package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env string `envconfig:"ENV" default:"local"`
	ServerHost string `envconfig:"SERVER_HOST" default:"0.0.0.0"`
	ServerPort int    `envconfig:"SERVER_PORT" default:"9090"`
	DBEngine                string `envconfig:"DB_ENGINE" default:"mysql"`
	DBUser                  string `envconfig:"DB_USER" default:"sampleuser"`
	DBPass                 string `envconfig:"DB_PASSWORD" default:"password"`
	DBHost                  string `envconfig:"DB_ADDR" default:"dev_db"`
	DBPort                  int    `envconfig:"DB_PORT" default:"3306"`
	DBName                  string `envconfig:"DB_NAME" default:"dev_db"`
}

func ProvideConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("go-gql", &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c Config) ServerAddr() string {
	return fmt.Sprintf("%s:%d", c.ServerHost, c.ServerPort)
}

func (c Config) DBAddr() string {
	return fmt.Sprintf("%s:%d", c.DBHost, c.DBPort)
}