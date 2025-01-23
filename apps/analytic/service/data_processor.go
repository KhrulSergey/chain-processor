package service

import (
	"errors"
	"fmt"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/khrulsergey/chain-processor/pkg/dto"
	mapper "github.com/khrulsergey/chain-processor/util"
	"reflect"
)

type DataProcessor interface {
	Handle(data interface{}) error
	onNewTrade(pairEvent *dto.TradeDto) error
	onNewPair(tradeEvent *dto.TokenPairDto) error
}

type dataProcessor struct {
	tradeService TradeService
	logger       logger.Logger
}

func NewDataProcessor(tradeService TradeService, logger logger.Logger) DataProcessor {
	return &dataProcessor{
		tradeService: tradeService,
		logger:       logger,
	}
}

func (ep *dataProcessor) Handle(data interface{}) error {
	var err error
	dataType := reflect.TypeOf(data)
	switch dataType {
	case reflect.TypeOf(&dto.TradeDto{}):
		err = ep.onNewTrade(data.(*dto.TradeDto))
	case reflect.TypeOf(&dto.TokenPairDto{}):
		err = ep.onNewPair(data.(*dto.TokenPairDto))
	default:
		{
			err = fmt.Errorf("can't process unknown data type: %s on dataProcessor", dataType)
		}
	}
	if err != nil {
		ep.logger.Warningf("Error processing data: %v. Error: %v", data, err)
	}
	return err
}

func (ep *dataProcessor) onNewPair(pairEvent *dto.TokenPairDto) error {
	return errors.New("not implemented")
}

func (ep *dataProcessor) onNewTrade(tradeDto *dto.TradeDto) error {
	ep.logger.Debugf("Try process data for chain: %d and hash: %s", tradeDto.ChainId, tradeDto.TxHash)
	trade := mapper.MapToEntity(tradeDto)
	trade, err := ep.tradeService.Create(trade)
	if err != nil {
		ep.logger.Warningf("Error processing trade event: %v. Error: %v", tradeDto, err)
		return err
	}
	ep.logger.Debugf("New trade successfully processed for chain: %d and hash: %s",
		trade.ChainId, trade.TxHash)
	return nil
}
