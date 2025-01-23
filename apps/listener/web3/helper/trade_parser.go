package helper

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/khrulsergey/chain-processor/pkg/types"
	"listener/web3/contract/bentobox_pool"
	"listener/web3/contract/dmm_pool"
	"listener/web3/contract/dodoex_pool"
	"listener/web3/contract/swaap_vault"
	"listener/web3/contract/uni_v2_pair"
	"listener/web3/contract/uni_v3_pool"
	"listener/web3/model"
	"math/big"
)

func ParseSwapEvent(eventLog *ethTypes.Log, tx *ethTypes.Transaction, block *ethTypes.Block,
	sender *common.Address, poolType types.ContractType, goalContractAddress string, client *ethclient.Client) (*model.TradeEvent, error) {
	switch poolType {
	case types.OneInchRouter:
		return parseOneInchRouterEvent(goalContractAddress, eventLog, tx, block, sender, client)
	case types.UniV3Pool:
		return parseUniV3PoolEvent(goalContractAddress, eventLog, tx, block, sender, client)
	case types.UniV2Pair:
		return parseUniV2PairEvent(goalContractAddress, eventLog, tx, block, sender, client)
	case types.DmmPool:
		return parseDmmPoolEvent(goalContractAddress, eventLog, tx, block, sender, client)
	case types.SwaapVault:
		return parseSwaapVaultEvent(goalContractAddress, eventLog, tx, block, sender, client)
	case types.DodoexPool:
		return parseDodoexPoolEvent(goalContractAddress, eventLog, tx, block, sender, client)
	case types.BentoBoxV1Pool:
		return parseBentoBoxV1Event(goalContractAddress, eventLog, tx, block, sender, client)
	}
	return nil, errors.New("unsupported pool type")
}

func parseOneInchRouterEvent(goalContractAddress string, eventLog *ethTypes.Log, tx *ethTypes.Transaction,
	block *ethTypes.Block, sender *common.Address, client *ethclient.Client) (*model.TradeEvent, error) {
	return &model.TradeEvent{
		PoolType: types.OneInchRouter,
	}, nil
}

func parseUniV3PoolEvent(goalContractAddress string, eventLog *ethTypes.Log, tx *ethTypes.Transaction,
	block *ethTypes.Block, sender *common.Address, client *ethclient.Client) (*model.TradeEvent, error) {
	chainId := int(tx.ChainId().Int64())
	contractUniV3Pair, err := uni_v3_pool.NewContract(eventLog.Address, client)
	if err != nil {
		return nil, err
	}
	poolData, err := contractUniV3Pair.ParseSwap(*eventLog)
	if err != nil {
		return nil, err
	}
	token0Addr, _ := contractUniV3Pair.Token0(nil)
	token1Addr, _ := contractUniV3Pair.Token1(nil)
	var tokenIn, tokenOut common.Address
	var amountIn, amountOut *big.Int
	if poolData.Amount1.Cmp(big.NewInt(0)) < 0 {
		amountIn = poolData.Amount0
		amountOut = new(big.Int).Abs(poolData.Amount1)
		tokenIn = token0Addr
		tokenOut = token1Addr
	} else {
		amountIn = poolData.Amount1
		amountOut = new(big.Int).Abs(poolData.Amount0)
		tokenIn = token1Addr
		tokenOut = token0Addr
	}

	return model.NewTradeEvent(
		eventLog.TxHash.Hex(),
		tx.Time(),
		block.Number(),
		goalContractAddress,
		sender.Hex(),
		poolData.Recipient.Hex(),
		tokenIn.Hex(),
		tokenOut.Hex(),
		amountIn,
		amountOut,
		tx.Value(),
		nil,
		nil,
		//SqrtPriceX96:     poolData.SqrtPriceX96,
		//Liquidity:        poolData.Liquidity,
		//Tick:             poolData.Tick,
		types.TradeEvent,
		chainId,
		types.UniV3Pool,
		eventLog.Address.Hex(),
	), nil
}

func parseUniV2PairEvent(goalContractAddress string, eventLog *ethTypes.Log, tx *ethTypes.Transaction,
	block *ethTypes.Block, sender *common.Address, client *ethclient.Client) (*model.TradeEvent, error) {
	chainId := int(tx.ChainId().Int64())
	contractPair, err := uni_v2_pair.NewContract(eventLog.Address, client)
	if err != nil {
		return nil, err
	}
	poolData, err := contractPair.ParseSwap(*eventLog)
	if err != nil {
		return nil, err
	}
	token0Addr, _ := contractPair.Token0(nil)
	token1Addr, _ := contractPair.Token1(nil)
	var tokenIn, tokenOut common.Address
	var amountIn, amountOut *big.Int
	if poolData.Amount0In.Cmp(big.NewInt(0)) > 0 {
		amountIn = poolData.Amount0In
		amountOut = poolData.Amount1Out
		tokenIn = token0Addr
		tokenOut = token1Addr
	} else {
		amountIn = poolData.Amount1In
		amountOut = poolData.Amount0Out
		tokenIn = token1Addr
		tokenOut = token0Addr
	}
	return model.NewTradeEvent(
		eventLog.TxHash.Hex(),
		tx.Time(),
		block.Number(),
		goalContractAddress,
		sender.Hex(),
		poolData.To.Hex(),
		tokenIn.Hex(),
		tokenOut.Hex(),
		amountIn,
		amountOut,
		tx.Value(),
		nil,
		nil,
		types.TradeEvent,
		chainId,
		types.UniV2Pair,
		eventLog.Address.Hex(),
	), nil
}

func parseDmmPoolEvent(goalContractAddress string, eventLog *ethTypes.Log, tx *ethTypes.Transaction,
	block *ethTypes.Block, sender *common.Address, client *ethclient.Client) (*model.TradeEvent, error) {
	chainId := int(tx.ChainId().Int64())
	contractPair, err := dmm_pool.NewContract(eventLog.Address, client)
	if err != nil {
		return nil, err
	}
	poolData, err := contractPair.ParseSwap(*eventLog)
	if err != nil {
		return nil, err
	}
	token0Addr, _ := contractPair.Token0(nil)
	token1Addr, _ := contractPair.Token1(nil)
	var tokenIn, tokenOut common.Address
	var amountIn, amountOut *big.Int
	if poolData.Amount0In.Cmp(big.NewInt(0)) > 0 {
		amountIn = poolData.Amount0In
		amountOut = poolData.Amount1Out
		tokenIn = token0Addr
		tokenOut = token1Addr
	} else {
		amountIn = poolData.Amount1In
		amountOut = poolData.Amount0Out
		tokenIn = token1Addr
		tokenOut = token0Addr
	}
	return model.NewTradeEvent(
		eventLog.TxHash.Hex(),
		tx.Time(),
		block.Number(),
		goalContractAddress,
		sender.Hex(),
		poolData.To.Hex(),
		tokenIn.Hex(),
		tokenOut.Hex(),
		amountIn,
		amountOut,
		tx.Value(),
		nil,
		nil,
		types.TradeEvent,
		chainId,
		types.DmmPool,
		eventLog.Address.Hex(),
	), nil
}

func parseSwaapVaultEvent(goalContractAddress string, eventLog *ethTypes.Log, tx *ethTypes.Transaction,
	block *ethTypes.Block, sender *common.Address, client *ethclient.Client) (*model.TradeEvent, error) {
	chainId := int(tx.ChainId().Int64())
	contractVault, err := swaap_vault.NewContract(eventLog.Address, client)
	if err != nil {
		return nil, err
	}
	poolData, err := contractVault.ParseSwap(*eventLog)
	if err != nil {
		return nil, err
	}
	return model.NewTradeEvent(
		eventLog.TxHash.Hex(),
		tx.Time(),
		block.Number(),
		goalContractAddress,
		sender.Hex(),
		eventLog.Address.Hex(),
		poolData.TokenIn.Hex(),
		poolData.TokenOut.Hex(),
		poolData.AmountIn,
		poolData.AmountOut,
		tx.Value(),
		nil,
		nil,
		types.TradeEvent,
		chainId,
		types.SwaapVault,
		eventLog.Address.Hex(),
	), nil
}

func parseDodoexPoolEvent(goalContractAddress string, eventLog *ethTypes.Log, tx *ethTypes.Transaction,
	block *ethTypes.Block, sender *common.Address, client *ethclient.Client) (*model.TradeEvent, error) {
	chainId := int(tx.ChainId().Int64())
	contractDoDoexPool, err := dodoex_pool.NewContract(eventLog.Address, client)
	if err != nil {
		return nil, err
	}
	poolData, err := contractDoDoexPool.ParseDODOSwap(*eventLog)
	if err != nil {
		return nil, err
	}
	return model.NewTradeEvent(
		eventLog.TxHash.Hex(),
		tx.Time(),
		block.Number(),
		goalContractAddress,
		sender.Hex(),
		poolData.Receiver.Hex(),
		poolData.FromToken.Hex(),
		poolData.ToToken.Hex(),
		poolData.FromAmount,
		poolData.ToAmount,
		tx.Value(),
		nil,
		nil,
		types.TradeEvent,
		chainId,
		types.DodoexPool,
		eventLog.Address.Hex(),
	), nil
}

func parseBentoBoxV1Event(goalContractAddress string, eventLog *ethTypes.Log, tx *ethTypes.Transaction,
	block *ethTypes.Block, sender *common.Address, client *ethclient.Client) (*model.TradeEvent, error) {
	chainId := int(tx.ChainId().Int64())
	contractBentoBoxPool, err := bentobox_pool.NewContract(eventLog.Address, client)
	if err != nil {
		return nil, err
	}
	poolData, err := contractBentoBoxPool.ParseLogWithdraw(*eventLog)
	if err != nil {
		return nil, err
	}
	return model.NewTradeEvent(
		eventLog.TxHash.Hex(),
		tx.Time(),
		block.Number(),
		goalContractAddress,
		sender.Hex(),
		poolData.From.Hex(),
		"", //don't know the tokenIn
		poolData.Token.Hex(),
		nil, //don't know the amountIn
		poolData.Amount,
		tx.Value(),
		nil,
		nil,
		types.TradeEvent,
		chainId,
		types.DodoexPool,
		eventLog.Address.Hex(),
	), nil
}
