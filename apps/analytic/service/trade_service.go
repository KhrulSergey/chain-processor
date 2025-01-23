package service

import (
	"analytic/storage"
	"github.com/google/uuid"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/khrulsergey/chain-processor/pkg/model"
)

type TradeService interface {
	Create(trade *model.Trade) (*model.Trade, error)
}

var (
	_ TradeService = &tradeService{}
)

type tradeService struct {
	tradeStorage storage.TradeStorage
	log          logger.Logger
}

func NewTradeService(tradeStorage storage.TradeStorage, log logger.Logger) TradeService {
	return &tradeService{
		tradeStorage: tradeStorage,
		log:          log,
	}
}

func (s *tradeService) Create(trade *model.Trade) (*model.Trade, error) {
	if trade.ID == uuid.Nil {
		trade.ID = uuid.New()
	}
	err := s.tradeStorage.Save(trade)
	if err != nil {
		return nil, err
	}
	return trade, nil
}
