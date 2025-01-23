package storage

import (
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/khrulsergey/chain-processor/pkg/model"

	"gorm.io/gorm"
)

type TradeStorage interface {
	Save(transaction *model.Trade) error
}

var (
	_ TradeStorage = &tradeStorageGorm{}
)

type tradeStorageGorm struct {
	db  *gorm.DB
	log logger.Logger
}

func NewTradeStorageGorm(db *gorm.DB, log logger.Logger) TradeStorage {
	return &tradeStorageGorm{
		db:  db,
		log: log,
	}
}

func (r *tradeStorageGorm) Save(entity *model.Trade) error {
	return r.db.Create(entity).Error
}
