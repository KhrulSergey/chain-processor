package model

import (
	"github.com/khrulsergey/chain-processor/pkg/types"
	"time"
)

type Transaction struct {
	ID              uint                      `gorm:"column:id;primarykey"`
	ChainId         int                       `gorm:"column:chain_id"`
	TxHash          string                    `gorm:"column:tx_hash"`
	BlockNumber     int64                     `gorm:"column:block_number"`
	Type            types.TransactionType     `gorm:"column:type"`
	ContractAddress string                    `gorm:"column:contract_address"`
	Status          types.ProcessedStatusType `gorm:"column:status"`
	TxAt            time.Time                 `gorm:"column:tx_at"`
	ProcessedAt     time.Time                 `gorm:"column:processed_at"`
	CreatedAt       time.Time                 `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt       time.Time                 `gorm:"column:updated_at;autoUpdateTime:true"`
}
