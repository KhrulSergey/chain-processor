package config

import "github.com/caarlos0/env/v11"

// AppConfig contains needed envs to run main app services
type AppConfig struct {
	AppVersion string `env:"APP_VERSION" envDefault:"0.1-beta"`

	ServiceExternalUrl string `env:"SERVICE_EXTERNAL_URL" envDefault:"localhost:4005"`
	ServiceProtocol    string `env:"SERVICE_PROTOCOL" envDefault:"http"`

	Host string `env:"SERVICE_HOST" envDefault:"localhost"`
	Port string `env:"SERVICE_PORT" envDefault:"4005"`

	GrpcHost string `env:"SERVICE_GRPC_HOST" envDefault:"localhost"`
	GrpcPort string `env:"SERVICE_GRPC_PORT" envDefault:"7105"`

	TokenSecret string `env:"TOKEN_SECRET" envDefault:"24ed47b3b99144c2817fc788bad7b003"`
	LoggerMode  string `env:"LOG_LEVEL" envDefault:"DEBUG"`
}

var MaxWorkersPerService = 200

// InitAppConfig parses envs and constructs the config for main app services
func InitAppConfig() (*AppConfig, error) {
	config := &AppConfig{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	return config, nil
}
