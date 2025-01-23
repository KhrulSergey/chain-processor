package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"strconv"
)

type ChainConfigs struct {
	chains map[int]ChainConfig
}

type ChainConfigNumber struct {
	ChainCount int `env:"CHAIN_SETTINGS_COUNT" envDefault:"1"`
}

type ChainType string

const (
	ETHEREUM ChainType = "ETHEREUM"
	LINEA    ChainType = "LINEA"
	POLYGON  ChainType = "POLYGON"
)

type ChainConfig struct {
	Enabled                 bool      `env:"ENABLED" envDefault:"false"`
	Type                    ChainType `env:"TYPE" envDefault:"ETHEREUM"`
	ChainId                 int       `env:"CHAINID" envDefault:"1"`
	RpcUrl                  string    `env:"RPC_URL" envDefault:"https://ethereum-rpc.publicnode.com"`
	WssUri                  string    `env:"WSS_URI" envDefault:"wss://ethereum-rpc.publicnode.com"`
	GasTokenMarketSymbol    string    `env:"GAS_TOKEN_MARKET_SYMBOL" envDefault:"ETH"`
	GasTokenDecimals        int       `env:"GAS_TOKEN_DECIMALS" envDefault:"18"`
	GasTokenWrapAddress     string    `env:"GAS_TOKEN_WRAP_ADDRESS" envDefault:"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"`
	GasLimitMax             int       `env:"GAS_LIMIT_MAX" envDefault:"0"`
	BlocksCountPerMinute    int64     `env:"BLOCKS_COUNT_PER_MINUTE" envDefault:"30"`
	BlockLimitByRequest     int64     `env:"BLOCK_LIMIT_BY_REQUEST" envDefault:"15000"`
	MaxRequestTimeoutSec    int       `env:"MAX_REQUEST_TIMEOUT_SEC" envDefault:"15"`
	MaxRequestsAttempt      int       `env:"MAX_REQUESTS_ATTEMPT" envDefault:"5"`
	MaxRequestsLimitsPerSec int       `env:"MAX_REQUESTS_LIMIT_PER_SEC" envDefault:"25"`
	AccessibilityTimeoutSec int       `env:"ACCESSIBILITY_TIMEOUT_SEC" envDefault:"20"`
	RestoreTimeoutSec       int       `env:"RESTORE_TIMEOUT_SEC" envDefault:"5"`

	PrivateKey string `env:"PRIVATE_KEY" envDefault:"0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"` //0xa0Ee7A142d267C1f36714E4a8F75612F20a79720

	DexContractAddress        string `env:"DEX_CONTRACT_ADDRESS" envDefault:""`
	DexFactoryContractAddress string `env:"DEX_FACTORY_CONTRACT_ADDRESS" envDefault:"0x11111112542D85B3EF69AE05771c2dCCff4fAa26"`
}

func ParseChainConfig() (map[int]ChainConfig, error) {
	chainConfigs := make(map[int]ChainConfig)

	chainConfigNumber := &ChainConfigNumber{}
	if err := env.Parse(chainConfigNumber); err != nil {
		return nil, err
	}

	// Load configurations for each network
	for i := 0; i < chainConfigNumber.ChainCount; i++ {
		var cfg ChainConfig
		// Set environment variable prefixes based on the index
		prefix := "CHAIN_SETTINGS_" + strconv.Itoa(i) + "_"
		// Parse the environment variables with the specific prefix
		if err := env.ParseWithOptions(&cfg, env.Options{Prefix: prefix}); err != nil {
			return nil, err
		}

		// Store the loaded configuration in the map
		chainConfigs[cfg.ChainId] = cfg
	}

	return chainConfigs, nil
}

func NewChainConfigs() (*ChainConfigs, error) {
	chainConfigManager := &ChainConfigs{
		chains: make(map[int]ChainConfig),
	}
	chains, err := ParseChainConfig()
	if err != nil {
		return nil, err
	}
	chainConfigManager.chains = chains
	return chainConfigManager, nil
}

func (m *ChainConfigs) GetConfigByChainId(chainId int) (*ChainConfig, error) {
	value, exists := m.chains[chainId]
	if !exists {
		return nil, fmt.Errorf("chain config not found for chain id %d", chainId)
	}
	return &value, nil
}

func (m *ChainConfigs) GetAllEnabledChains() []*ChainConfig {
	var enabledChains []*ChainConfig
	for _, config := range m.chains {
		if config.Enabled {
			enabledChains = append(enabledChains, &config)
		}
	}
	return enabledChains
}
