package util

import (
	"github.com/google/uuid"
	"github.com/khrulsergey/chain-processor/pkg/dto"
	"github.com/khrulsergey/chain-processor/pkg/model"
	"github.com/khrulsergey/chain-processor/pkg/types"
	"github.com/shopspring/decimal"
	web3model "listener/web3/model"
	"time"
)

func MapToTransaction(tradeEvent *web3model.TradeEvent) *model.Transaction {
	return &model.Transaction{
		ChainId:         tradeEvent.GetChainId(),
		TxHash:          tradeEvent.GetTxHash(),
		BlockNumber:     tradeEvent.GetBlockNumber().Int64(),
		Type:            tradeEvent.GetType(),
		ContractAddress: tradeEvent.GetContractAddress(),
		TxAt:            tradeEvent.GetTimestamp(),
		Status:          types.Processed,
		ProcessedAt:     time.Now(),
	}
}

func MapToDto(tradeEvent *web3model.TradeEvent) *dto.TradeDto {
	return &dto.TradeDto{
		Type:             tradeEvent.GetType(),
		ChainId:          tradeEvent.GetChainId(),
		BlockNumber:      tradeEvent.GetBlockNumber(),
		ContractAddress:  tradeEvent.GetContractAddress(),
		TxHash:           tradeEvent.GetTxHash(),
		Timestamp:        tradeEvent.GetTimestamp(),
		PoolType:         tradeEvent.PoolType,
		InitiatorAddress: tradeEvent.InitiatorAddress,
		AgentAddress:     tradeEvent.AgentAddress,
		FromTokenAddress: tradeEvent.FromTokenAddress,
		ToTokenAddress:   tradeEvent.ToTokenAddress,
		FromTokenAmount:  tradeEvent.FromTokenAmount,
		ToTokenAmount:    tradeEvent.ToTokenAmount,
		EthValue:         tradeEvent.EthValue,
		Price:            tradeEvent.Price,
	}
}

func MapToEntity(tradeDto *dto.TradeDto) *model.Trade {
	var fromTokenAmount, toTokenAmount, ethValue, price decimal.Decimal
	var id uuid.UUID
	if tradeDto.FromTokenAmount != nil {
		fromTokenAmount = decimal.NewFromBigInt(tradeDto.FromTokenAmount, -18)
	}
	if tradeDto.ToTokenAmount != nil {
		toTokenAmount = decimal.NewFromBigInt(tradeDto.ToTokenAmount, -18)
	}
	if tradeDto.EthValue != nil {
		ethValue = decimal.NewFromBigInt(tradeDto.EthValue, -18)
	}
	if tradeDto.Price != nil {
		price = decimal.NewFromBigInt(tradeDto.Price, -18)
	}
	if tradeDto.ExternalId != "" {
		id, _ = uuid.Parse(tradeDto.ExternalId)
	}
	return &model.Trade{
		ID:                    id,
		ChainId:               tradeDto.ChainId,
		BlockNumber:           tradeDto.BlockNumber.Int64(),
		ContractAddress:       tradeDto.ContractAddress,
		TxHash:                tradeDto.TxHash,
		TxAt:                  tradeDto.Timestamp,
		PoolType:              tradeDto.PoolType,
		InitiatorAddress:      tradeDto.InitiatorAddress,
		AgentAddress:          tradeDto.AgentAddress,
		FromTokenAddress:      tradeDto.FromTokenAddress,
		ToTokenAddress:        tradeDto.ToTokenAddress,
		FromTokenAmount:       &fromTokenAmount,
		ToTokenAmount:         &toTokenAmount,
		EthValue:              &ethValue,
		Price:                 &price,
		CurrentBasePoolAmount: nil,
	}
}
