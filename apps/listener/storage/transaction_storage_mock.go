package storage

import (
	"github.com/khrulsergey/chain-processor/pkg/model"
	"github.com/khrulsergey/chain-processor/pkg/types"
	"github.com/stretchr/testify/mock"
)

// MockTransactionStorage for mocking TransactionStorage dependency
type MockTransactionStorage struct {
	mock.Mock
}

func (m *MockTransactionStorage) GetLatestBlockNumber() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTransactionStorage) Save(transaction *model.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionStorage) NotExistsByTx(txHash string, chainId int) bool {
	args := m.Called(txHash, chainId)
	return args.Bool(0)
}

func (m *MockTransactionStorage) FindLastProcessedBlock(chainId int, contractAddress string) (int64, error) {
	args := m.Called(chainId, contractAddress)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTransactionStorage) FindLastProcessedBlockForEventType(chainId int, txType types.TransactionType, contractAddress string) (int64, error) {
	args := m.Called(chainId, txType, contractAddress)
	return args.Get(0).(int64), args.Error(1)
}
