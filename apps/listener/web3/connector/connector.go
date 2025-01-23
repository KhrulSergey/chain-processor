package connector

import "github.com/ethereum/go-ethereum/ethclient"

type ChainConnector interface {
	Connect() error
	Dispose()
	CheckAccessibility() error

	GetHttpClient() *ethclient.Client
	GetWssClient() *ethclient.Client

	GetCurrentBlockNumber() (int64, error)
}
