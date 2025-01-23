package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/khrulsergey/chain-processor/pkg/types"
	"listener/web3/contract/bentobox_pool"
	"listener/web3/contract/dmm_pool"
	"listener/web3/contract/dodoex_pool"
	"listener/web3/contract/oneinch_router"
	"listener/web3/contract/swaap_vault"
	"listener/web3/contract/uni_v2_pair"
	"listener/web3/contract/uni_v3_pool"
	"math/big"
	"time"
)

// EvmContract defines the common methods and events for the contracts
type EvmContract[T any] interface {
	NewContract(address common.Address, backend bind.ContractBackend) (*T, error)
	//NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error)
}

var AbiRefMap = map[types.ContractType]string{
	types.OneInchRouter:  oneinch_router.ContractABI,
	types.UniV2Pair:      uni_v2_pair.ContractABI,
	types.UniV3Pool:      uni_v3_pool.ContractABI,
	types.DmmPool:        dmm_pool.ContractABI,
	types.SwaapVault:     swaap_vault.ContractABI,
	types.DodoexPool:     dodoex_pool.ContractABI,
	types.BentoBoxV1Pool: bentobox_pool.ContractABI,
}

var TradeMethodNameMap = map[types.ContractType]string{
	types.OneInchRouter:  "",
	types.UniV2Pair:      "Swap",
	types.UniV3Pool:      "Swap",
	types.DmmPool:        "Swap",
	types.SwaapVault:     "Swap",
	types.DodoexPool:     "DODOSwap",
	types.BentoBoxV1Pool: "withdraw",
}

type TradeEvent struct {
	TxHash           string
	Timestamp        time.Time
	BlockNumber      *big.Int
	ChainId          int
	InitiatorAddress string
	AgentAddress     string
	FromTokenAddress string
	ToTokenAddress   string
	FromTokenAmount  *big.Int
	ToTokenAmount    *big.Int

	//SqrtPriceX96    *big.Int `json:"sqrtPriceX96,omitempty"`
	//Liquidity       *big.Int `json:"liquidity,omitempty"`
	//Tick            *big.Int `json:"tick,omitempty"`
	PoolType        types.ContractType
	ContractAddress string
}
