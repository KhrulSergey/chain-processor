package service

import (
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/khrulsergey/chain-processor/pkg/model"
	"github.com/khrulsergey/chain-processor/pkg/types"
	"listener/storage"
)

type TransactionService interface {
	Create(transaction *model.Transaction) (*model.Transaction, error)
	NotExistsByTxId(chainId int, txHash string) bool
	FindLastProcessedBlockNumber(chainId int, txType *types.TransactionType, contractAddress string) (int64, error)
}

var (
	_ TransactionService = &transactionService{}
)

type transactionService struct {
	transactionRepository storage.TransactionStorage
	log                   logger.Logger
}

func NewTransactionService(transactionStorage storage.TransactionStorage, log logger.Logger) TransactionService {
	return &transactionService{
		transactionRepository: transactionStorage,
		log:                   log,
	}
}

func (s *transactionService) Create(transaction *model.Transaction) (*model.Transaction, error) {
	err := s.transactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) NotExistsByTxId(chainId int, txHash string) bool {
	return s.transactionRepository.NotExistsByTx(txHash, chainId)
}

func (s *transactionService) FindLastProcessedBlockNumber(chainId int, txType *types.TransactionType, contractAddress string) (int64, error) {
	if txType == nil {
		return s.transactionRepository.FindLastProcessedBlock(chainId, contractAddress)
	}
	return s.transactionRepository.FindLastProcessedBlockForEventType(chainId, *txType, contractAddress)
}
