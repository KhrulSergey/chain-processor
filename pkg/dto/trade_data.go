package dto

import (
	"github.com/khrulsergey/chain-processor/pkg/types"
	"math/big"
	"time"
)

type TradeDto struct {
	ExternalId       string                `json:"-"`
	Type             types.TransactionType `json:"type"`
	ChainId          int                   `json:"chainId"`
	BlockNumber      *big.Int              `json:"blockNumber"`
	ContractAddress  string                `json:"contractAddress"`
	TxHash           string                `json:"txHash"`
	Timestamp        time.Time             `json:"timestamp"`
	PoolType         types.ContractType    `json:"poolType"`
	InitiatorAddress string                `json:"initiatorAddress"`
	AgentAddress     string                `json:"agentAddress"`
	FromTokenAddress string                `json:"fromTokenAddress"`
	ToTokenAddress   string                `json:"toTokenAddress"`
	FromTokenAmount  *big.Int              `json:"fromTokenAmount"`
	ToTokenAmount    *big.Int              `json:"toTokenAmount"`
	EthValue         *big.Int              `json:"ethValue"`
	Price            *big.Int              `json:"price"`
}

func (e *TradeDto) GetExternalId() string {
	return e.ExternalId
}

func (e *TradeDto) SetExternalId(externalId string) {
	e.ExternalId = externalId
}
