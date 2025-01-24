package service

import (
	"errors"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/stretchr/testify/assert"
	"listener/storage"
	"testing"

	"github.com/khrulsergey/chain-processor/pkg/model"
	"github.com/khrulsergey/chain-processor/pkg/types"
)

func TestCreateTransaction(t *testing.T) {
	mockStorage := new(storage.MockTransactionStorage)
	mockLogger := new(logger.MockLogger)
	transaction := &model.Transaction{ID: 1}

	mockStorage.
		On("Save", transaction).
		Return(nil)

	service := NewTransactionService(mockStorage, mockLogger)

	createdTransaction, err := service.Create(transaction)

	assert.NoError(t, err)
	assert.Equal(t, transaction, createdTransaction)
	mockStorage.AssertCalled(t, "Save", transaction)
}

func TestCreateTransactionError(t *testing.T) {
	mockStorage := new(storage.MockTransactionStorage)
	mockLogger := new(logger.MockLogger)
	transaction := &model.Transaction{ID: 1}

	mockStorage.
		On("Save", transaction).
		Return(errors.New("failed to save transaction"))

	service := NewTransactionService(mockStorage, mockLogger)

	createdTransaction, err := service.Create(transaction)

	assert.Error(t, err)
	assert.Nil(t, createdTransaction)
	assert.Equal(t, "failed to save transaction", err.Error())
	mockStorage.AssertCalled(t, "Save", transaction)
}

func TestNotExistsByTxId(t *testing.T) {
	mockStorage := new(storage.MockTransactionStorage)
	mockLogger := new(logger.MockLogger)
	chainId := 1
	txHash := "mockTxHash"

	mockStorage.
		On("NotExistsByTx", txHash, chainId).
		Return(true)

	service := NewTransactionService(mockStorage, mockLogger)

	result := service.NotExistsByTxId(chainId, txHash)

	assert.True(t, result)
	mockStorage.AssertCalled(t, "NotExistsByTx", txHash, chainId)
}

func TestFindLastProcessedBlockNumber(t *testing.T) {
	mockStorage := new(storage.MockTransactionStorage)
	mockLogger := new(logger.MockLogger)
	chainId := 1
	contractAddress := "mockAddress"

	mockStorage.
		On("FindLastProcessedBlock", chainId, contractAddress).
		Return(int64(100), nil)

	service := NewTransactionService(mockStorage, mockLogger)

	result, err := service.FindLastProcessedBlockNumber(chainId, nil, contractAddress)

	assert.NoError(t, err)
	assert.Equal(t, int64(100), result)
	mockStorage.AssertCalled(t, "FindLastProcessedBlock", chainId, contractAddress)
}

func TestFindLastProcessedBlockNumberForEventType(t *testing.T) {
	mockStorage := new(storage.MockTransactionStorage)
	mockLogger := new(logger.MockLogger)
	chainId := 1
	contractAddress := "mockAddress"
	txType := types.TransactionType("someType")

	mockStorage.
		On("FindLastProcessedBlockForEventType", chainId, txType, contractAddress).
		Return(int64(200), nil)

	service := NewTransactionService(mockStorage, mockLogger)

	result, err := service.FindLastProcessedBlockNumber(chainId, &txType, contractAddress)

	assert.NoError(t, err)
	assert.Equal(t, int64(200), result)
	mockStorage.AssertCalled(t, "FindLastProcessedBlockForEventType", chainId, txType, contractAddress)
}
