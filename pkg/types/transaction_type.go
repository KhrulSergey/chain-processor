package types

type TransactionType string

const (
	PairCreatedEvent  TransactionType = "PAIR_CREATED"
	TradeEvent        TransactionType = "TOKEN_TRADE"
	Migrate           TransactionType = "MIGRATE"
	LiquidityTransfer TransactionType = "LIQUIDITY_TRANSFER"
)

type ProcessedStatusType string

const (
	Processed ProcessedStatusType = "PROCESSED"
	Failed    ProcessedStatusType = "FAILED"
)
