package config

import "github.com/caarlos0/env/v11"

type PgDBConfig struct {
	DatabaseHost     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	DatabasePort     string `env:"POSTGRES_PORT" envDefault:"5432"`
	DatabaseUser     string `env:"POSTGRES_USER" envDefault:"postgres"`
	DatabasePassword string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	DatabaseName     string `env:"POSTGRES_DB" envDefault:"evm_events"`
	DatabaseRootCA   string `env:"POSTGRES_ROOT_CA" envDefault:""`
	ConnMaxLifeTime  int    `env:"POSTGRES_CONN_MAX_LIFETIME_S" envDefault:"3600"`
}

func NewPgDBConfig() (*PgDBConfig, error) {
	config := &PgDBConfig{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	return config, nil
}
