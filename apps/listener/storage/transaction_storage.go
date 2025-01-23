package storage

import (
	"github.com/khrulsergey/chain-processor/pkg/model"
	"github.com/khrulsergey/chain-processor/pkg/types"
	"gorm.io/gorm"
)

type TransactionStorage interface {
	GetLatestBlockNumber() (int64, error)
	NotExistsByTx(txId string, chainId int) bool
	FindLastProcessedBlock(chainId int, contractAddress string) (int64, error)
	FindLastProcessedBlockForEventType(chainId int, txType types.TransactionType, contractAddress string) (int64, error)
	Save(transaction *model.Transaction) error
}

type TransactionStorageGorm struct {
	db *gorm.DB
}

func NewTransactionStorageGorm(db *gorm.DB) TransactionStorage {
	return &TransactionStorageGorm{
		db: db,
	}
}

func (r *TransactionStorageGorm) GetLatestBlockNumber() (int64, error) {
	var blockNumber int64
	err := r.db.Table("transactions").Select("MAX(block_number)").Scan(&blockNumber).Error
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}

func (r *TransactionStorageGorm) NotExistsByTx(txId string, chainId int) bool {
	var exists bool
	err := r.db.Raw("SELECT NOT EXISTS (SELECT id FROM transactions WHERE lower(tx_hash) = lower(?) AND chain_id = ?)", txId, chainId).Scan(&exists).Error
	if err != nil {
		return false
	}
	return exists
}

func (r *TransactionStorageGorm) FindLastProcessedBlock(chainId int, contractAddress string) (int64, error) {
	var blockNumber int64
	err := r.db.Raw("SELECT coalesce(max_block, 0) FROM (SELECT max(t.block_number) AS max_block FROM transactions t WHERE t.chain_id = ? AND lower(t.contract_address) = lower(?)) as q", chainId, contractAddress).Scan(&blockNumber).Error
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}

func (r *TransactionStorageGorm) FindLastProcessedBlockForEventType(chainId int, txType types.TransactionType, contractAddress string) (int64, error) {
	var blockNumber int64
	err := r.db.Raw("SELECT coalesce(max_block, 0) FROM (SELECT max(t.block_number) AS max_block FROM transactions t WHERE t.chain_id = ? AND t.type = ? AND lower(t.contract_address) = lower(?)) as q", chainId, txType, contractAddress).Scan(&blockNumber).Error
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}

func (r *TransactionStorageGorm) Save(transaction *model.Transaction) error {
	return r.db.Create(transaction).Error
}
