package model

import (
	"github.com/khrulsergey/chain-processor/pkg/types"
	"math/big"
	"time"
)

type ContractEvent interface {
	GetType() types.TransactionType
	GetChainId() int
	GetBlockNumber() *big.Int
	GetContractAddress() string
	GetTxHash() string
	GetTimestamp() time.Time
}

type BaseContractEvent struct {
	Type            types.TransactionType `json:"type"`
	ChainId         int                   `json:"chainId"`
	BlockNumber     *big.Int              `json:"blockNumber"`
	ContractAddress string                `json:"contractAddress"`
	TxHash          string                `json:"txHash"`
	Timestamp       time.Time             `json:"timestamp"`
}

func (e *BaseContractEvent) GetType() types.TransactionType {
	return e.Type
}

func (e *BaseContractEvent) GetChainId() int {
	return e.ChainId
}

func (e *BaseContractEvent) GetBlockNumber() *big.Int {
	return e.BlockNumber
}

func (e *BaseContractEvent) GetContractAddress() string {
	return e.ContractAddress
}

func (e *BaseContractEvent) GetTxHash() string {
	return e.TxHash
}

func (e *BaseContractEvent) GetTimestamp() time.Time {
	return e.Timestamp
}
