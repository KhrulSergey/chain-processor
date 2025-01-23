package service

import (
	"errors"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/khrulsergey/chain-processor/pkg/types"
	mapper "github.com/khrulsergey/chain-processor/util"
	web3model "listener/web3/model"
)

type EventProcessor interface {
	Handle(event web3model.ContractEvent)
	onNewPair(pairEvent *web3model.PairEvent) error
	onNewTrade(tradeEvent *web3model.TradeEvent) error
}

type eventProcessor struct {
	transactionService TransactionService
	kafkaService       KafkaService
	logger             logger.Logger
}

func NewEventProcessor(transactionService TransactionService, kafkaService KafkaService,
	logger logger.Logger) EventProcessor {
	return &eventProcessor{
		transactionService: transactionService,
		kafkaService:       kafkaService,
		logger:             logger,
	}
}

func (ep *eventProcessor) Handle(event web3model.ContractEvent) {
	var err error
	switch event.GetType() {
	case types.TradeEvent:
		err = ep.onNewTrade(event.(*web3model.TradeEvent))
	case types.PairCreatedEvent:
		err = ep.onNewPair(event.(*web3model.PairEvent))
	default:
		return
	}
	if err != nil {
		ep.logger.Warningf("Error processing new event type:%s, hash: %s, chain: %d. Error: %v",
			event.GetType(), event.GetTxHash(), event.GetChainId(), err)
	}
}

func (ep *eventProcessor) onNewPair(pairEvent *web3model.PairEvent) error {
	return errors.New("not implemented")
}

func (ep *eventProcessor) onNewTrade(tradeEvent *web3model.TradeEvent) error {
	tx := mapper.MapToTransaction(tradeEvent)
	err := ep.kafkaService.ProduceMessage(mapper.MapToDto(tradeEvent))
	if err != nil {
		ep.logger.Warningf("Can't send data to event with hash: %s to broker: %v ",
			tx.TxHash, err)
		tx.Status = types.Failed
	}
	tx, err = ep.transactionService.Create(tx)
	if err != nil && tx != nil {
		ep.logger.Debugf("New trade event successfully processed : %s", tx.TxHash)
	}
	return err
}
