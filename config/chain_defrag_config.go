package config

import (
	"github.com/caarlos0/env/v11"
)

// ChainEventsDefragmentationConfig contains needed envs to run defragmentation chain events tasks
type ChainEventsDefragmentationConfig struct {
	FirstTask  DefragmentationTaskConfig `envPrefix:"FIRST_TASK_"`
	SecondTask DefragmentationTaskConfig `envPrefix:"SECOND_TASK_"`
}

type DefragmentationTaskConfig struct {
	DelayMin            int `env:"DELAY_MIN"`
	TimeoutMin          int `env:"TIMEOUT_MIN"`
	EventsAgeScopeHours int `env:"EVENTS_AGE_SCOPE_HOURS"`
}

// NewChainEventsDefragmentationConfig parses envs and constructs the config
func NewChainEventsDefragmentationConfig() (*ChainEventsDefragmentationConfig, error) {
	config := &ChainEventsDefragmentationConfig{}

	//chainConfigs := make(map[int]ChainConfig)
	//chainConfigNumber := &ChainConfigNumber{}
	//if err := env.Parse(chainConfigNumber); err != nil {
	//	return nil, err
	//}
	//
	//// Load configurations for each network
	//for i := 0; i < chainConfigNumber.ChainCount; i++ {
	//	var cfg ChainConfig
	//	// Set environment variable prefixes based on the index
	//	prefix := "CHAIN_SETTINGS_" + strconv.Itoa(i) + "_"
	//	// Parse the environment variables with the specific prefix
	//	if err := env.ParseWithOptions(&cfg, env.Options{Prefix: prefix}); err != nil {
	//		return nil, err
	//	}
	//
	//	// Store the loaded configuration in the map
	//	chainConfigs[cfg.ChainId] = cfg
	//}
	//
	//return chainConfigs, nil

	if err := env.Parse(config); err != nil {
		return nil, err
	}
	return config, nil
}
