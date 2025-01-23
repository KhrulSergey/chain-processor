package model

import (
	"github.com/khrulsergey/chain-processor/pkg/types"
	"math/big"
	"time"
)

type TradeEvent struct {
	ContractEvent
	PoolType              types.ContractType
	PoolContractAddress   string
	InitiatorAddress      string
	AgentAddress          string
	FromTokenAddress      string
	ToTokenAddress        string
	FromTokenAmount       *big.Int
	ToTokenAmount         *big.Int
	EthValue              *big.Int
	Price                 *big.Int
	CurrentBasePoolAmount *big.Int
}

func NewTradeEvent(txHash string, timestamp time.Time, blockNumber *big.Int, contractAddress string,
	initiatorAddress string, agentAddress string, fromTokenAddress string, toTokenAddress string,
	fromTokenAmount *big.Int, toTokenAmount *big.Int, ethValue *big.Int,
	price *big.Int, currentBasePoolAmount *big.Int, eventType types.TransactionType, chainId int,
	poolType types.ContractType, poolContractAddress string) *TradeEvent {
	return &TradeEvent{
		ContractEvent: &BaseContractEvent{
			Type:            eventType,
			ChainId:         chainId,
			BlockNumber:     blockNumber,
			ContractAddress: contractAddress,
			TxHash:          txHash,
			Timestamp:       timestamp,
		},
		PoolType:              poolType,
		PoolContractAddress:   poolContractAddress,
		InitiatorAddress:      initiatorAddress,
		AgentAddress:          agentAddress,
		FromTokenAddress:      fromTokenAddress,
		ToTokenAddress:        toTokenAddress,
		FromTokenAmount:       fromTokenAmount,
		ToTokenAmount:         toTokenAmount,
		EthValue:              ethValue,
		Price:                 price,
		CurrentBasePoolAmount: currentBasePoolAmount,
	}
}
