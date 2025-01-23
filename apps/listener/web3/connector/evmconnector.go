package connector

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	"time"
)

type EVMConnector struct {
	config    *config.ChainConfig
	rpcClient *ethclient.Client
	wssClient *ethclient.Client
	logger    logger.Logger
}

func NewEVMConnector(cfg *config.ChainConfig, logger logger.Logger) *EVMConnector {
	return &EVMConnector{config: cfg, logger: logger}
}

func (e *EVMConnector) Connect() error {
	var err error
	e.rpcClient, err = ethclient.Dial(e.config.RpcUrl)
	if err != nil {
		e.logger.Debugf("Failed to connect to RPC on chain %d: %v", e.config.ChainId, err)
		return err
	}
	e.wssClient, err = ethclient.Dial(e.config.WssUri)
	if err != nil {
		e.logger.Debugf("Failed to connect to Wss on chain %d: %v", e.config.ChainId, err)
		return err
	}
	e.logger.Debug("Connected to chain id: ", e.config.ChainId)
	return nil
}

func (e *EVMConnector) Dispose() {
	if e.rpcClient != nil {
		e.rpcClient.Close()
	}
	if e.wssClient != nil {
		e.wssClient.Close()
	}
	e.logger.Debug("Disposed connector for chain id: ", e.config.ChainId)
}

func (e *EVMConnector) CheckAccessibility() error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(e.config.MaxRequestTimeoutSec)*time.Second)
	defer cancel()
	blockNumber, err := e.wssClient.BlockNumber(ctx)
	e.logger.Debug("Block number over WSS:", blockNumber)
	if err != nil {
		return err
	}

	blockNumber, err = e.rpcClient.BlockNumber(ctx)
	e.logger.Debug("Block number over RPC:", blockNumber)
	if err != nil {
		return err
	}

	e.logger.Debug("Accessibility check passed for chain id: ", e.config.ChainId)
	return nil
}

func (e *EVMConnector) GetHttpClient() *ethclient.Client {
	return e.rpcClient
}

func (e *EVMConnector) GetWssClient() *ethclient.Client {
	return e.wssClient
}

func (e *EVMConnector) GetCurrentBlockNumber() (int64, error) {
	var block uint64
	var err error

	for i := 0; i < e.config.MaxRequestsAttempt; i++ {
		block, err = e.rpcClient.BlockNumber(context.Background())
		if err == nil {
			return int64(block), nil
		}
		e.logger.Errorf("Error on sending request: getCurrentBlock, on chain id: %d. Error: %v", e.config.ChainId, err)
		time.Sleep(time.Duration(e.config.MaxRequestTimeoutSec) * time.Second)
	}

	return 0, err
}
