package config

import "github.com/caarlos0/env/v11"

type ChDBConfig struct {
	DatabaseHost     string `env:"CLICKHOUSE_HOST" envDefault:"127.0.0.1"`
	DatabaseSSL      bool   `env:"CLICKHOUSE_HOST_SSL" envDefault:"false"`
	DatabasePort     string `env:"CLICKHOUSE_PORT" envDefault:"5123"`
	DatabaseUser     string `env:"CLICKHOUSE_USER" envDefault:"clickhouse"`
	DatabasePassword string `env:"CLICKHOUSE_PASSWORD" envDefault:"clickhouse"`
	DatabaseName     string `env:"CLICKHOUSE_DB" envDefault:"evm_data"`
	DatabaseRootCA   string `env:"CLICKHOUSE_ROOT_CA" envDefault:""`
}

func NewChDBConfig() (*ChDBConfig, error) {
	config := &ChDBConfig{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	return config, nil
}
