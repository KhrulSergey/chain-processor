package model

import (
	"github.com/google/uuid"
	"github.com/khrulsergey/chain-processor/pkg/types"
	"github.com/shopspring/decimal"
	"time"
)

type Trade struct {
	ID                    uuid.UUID          `gorm:"column:id;primarykey"`
	ChainId               int                `gorm:"column:chain_id"`
	PoolType              types.ContractType `gorm:"column:pool_type"`
	InitiatorAddress      string             `gorm:"column:initiator_address"`
	AgentAddress          string             `gorm:"column:agent_address"`
	FromTokenAddress      string             `gorm:"column:from_token_address"`
	ToTokenAddress        string             `gorm:"column:to_token_address"`
	FromTokenAmount       *decimal.Decimal   `gorm:"column:from_token_amount"`
	ToTokenAmount         *decimal.Decimal   `gorm:"column:to_token_amount"`
	EthValue              *decimal.Decimal   `gorm:"column:eth_value"`
	Price                 *decimal.Decimal   `gorm:"column:price"`
	CurrentBasePoolAmount *decimal.Decimal   `gorm:"column:current_base_pool_amount"`
	TxHash                string             `gorm:"column:tx_hash"`
	BlockNumber           int64              `gorm:"column:block_number"`
	ContractAddress       string             `gorm:"column:contract_address"`
	TxAt                  time.Time          `gorm:"column:tx_at"`
	CreatedAt             time.Time          `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt             time.Time          `gorm:"column:updated_at;autoUpdateTime:true"`
}
