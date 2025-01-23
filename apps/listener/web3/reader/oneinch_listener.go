package reader

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/khrulsergey/chain-processor/pkg/types"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/pool"
	"listener/service"
	"listener/web3/connector"
	"listener/web3/contract"
	"listener/web3/helper"
	"listener/web3/model"
	"math/big"
	"strings"
	"sync/atomic"
	"time"
)

type ContractListener struct {
	ctx                  context.Context
	contractChainId      int
	factoryAddress       string
	supportedContractAbi map[types.ContractType]abi.ABI
	chainConnector       connector.ChainConnector
	eventProcessor       service.EventProcessor
	transactionService   service.TransactionService
	requestsLimitPerSec  int
	requestsLimitTotal   int64
	lastProcessedBlock   atomic.Int64
	log                  logger.Logger
}

func NewOneInchContractListener(ctx context.Context, cfg *config.ChainConfig, chainConnector *connector.ChainConnectorManager,
	chainEventProcessor service.EventProcessor, transactionService service.TransactionService, log logger.Logger) *ContractListener {
	ccm, err := chainConnector.GetChainConnector(cfg.ChainId)
	if err != nil {
		log.Panicf("can't initialize listener for chain %d with no chainConnector", cfg.ChainId)
	}
	cl := &ContractListener{
		ctx:                 ctx,
		contractChainId:     cfg.ChainId,
		factoryAddress:      cfg.DexFactoryContractAddress,
		chainConnector:      ccm,
		eventProcessor:      chainEventProcessor,
		transactionService:  transactionService,
		requestsLimitPerSec: cfg.MaxRequestsLimitsPerSec,
		requestsLimitTotal:  cfg.BlockLimitByRequest,
		log:                 log,
	}

	cl.supportedContractAbi = make(map[types.ContractType]abi.ABI)
	for name, abiStr := range contract.AbiRefMap {
		cl.supportedContractAbi[name], err = abi.JSON(strings.NewReader(abiStr))
		if err != nil {
			cl.log.Warning(err)
		}
	}

	//start continuously listen contract
	go func() {
		for {
			cl.Listen()
			cl.log.Info("Sleep listening for 5 sec for next iteration")
			time.Sleep(5 * time.Second)
		}
	}()
	return cl
}

func (cl *ContractListener) Listen() {
	contractAddress := common.HexToAddress(cl.factoryAddress)
	cl.log.Debug("Start listen contract: ", contractAddress.Hex())

	httpConnector := cl.chainConnector.GetHttpClient()
	latestBlock, err := httpConnector.BlockNumber(cl.ctx)
	if err != nil {
		cl.log.Warningf("Failed to get latest block number: %v", err)
	}
	cl.log.Debugf("Latest block: %d", latestBlock)

	startBlock := cl.lastProcessedBlock.Load()
	if startBlock == 0 {
		startBlock, err = cl.transactionService.FindLastProcessedBlockNumber(cl.contractChainId, nil, contractAddress.Hex())
	}
	if startBlock == 0 {
		startBlock = int64(latestBlock) - cl.requestsLimitTotal
	}

	blockChan := make(chan *ethTypes.Block, 100)
	errChan := make(chan error, 100)
	eventChan := make(chan model.ContractEvent, 100)

	var wg conc.WaitGroup
	//Load blocks in separate goroutine
	wg.Go(func() { cl.blockFetcher(blockChan, errChan, startBlock, int64(latestBlock)) })

	//Process fetched blocks in separate goroutine
	go func() {
		for block := range blockChan {
			cl.processBlock(block, contractAddress, eventChan, errChan)
		}
	}()

	go func() {
		for event := range eventChan {
			if event != nil {
				cl.eventProcessor.Handle(event)
				cl.log.Infof("Catch event chain: %d, hash: %s, type: %v",
					event.GetChainId(), event.GetTxHash(), event.GetType())
				cl.log.Info("---------------------------------------------------")
			}
		}
	}()

	// Check for errors
	go func() {
		for err := range errChan {
			if err != nil {
				cl.log.Warningf("Error: %v", err)
			}
		}
	}()

	// Wait for all block fetchers to complete
	wg.Wait()
	close(blockChan)
	close(errChan)
}

func (cl *ContractListener) blockFetcher(blockChan chan<- *ethTypes.Block, errChan chan<- error,
	startBlock, latestBlock int64) {
	httpConnector := cl.chainConnector.GetHttpClient()
	var blockWorkerPool = pool.New().WithMaxGoroutines(cl.requestsLimitPerSec)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	counter := atomic.Int64{}
	for blockNumber := startBlock; blockNumber <= latestBlock; blockNumber++ {
		blockWorkerPool.Go(func() {
			counter.Add(1)
			cl.blockFetcherWorker(blockNumber, httpConnector, blockChan, errChan)
			counter.Add(-1)
		})
		if counter.Load() >= int64(cl.requestsLimitPerSec) {
			<-ticker.C
		}
		cl.lastProcessedBlock.Store(blockNumber)
	}
	blockWorkerPool.Wait()
}

func (cl *ContractListener) blockFetcherWorker(blockNumber int64, httpConnector *ethclient.Client,
	blockChan chan<- *ethTypes.Block, errChan chan<- error) {
	block, err := httpConnector.BlockByNumber(cl.ctx, big.NewInt(blockNumber))
	if block == nil {
		errChan <- fmt.Errorf("can't fetch block by number: %d. Original error: %v", blockNumber, err)
		time.Sleep(5 * time.Second) //sleep worker for 5 sec in case of connector error
		return
	}
	//cl.log.Debugf("Fetched block by number: %d", block.Number().Uint64())
	blockChan <- block
	if err != nil {
		errChan <- err
		return
	}
}

func (cl *ContractListener) processBlock(block *ethTypes.Block, contractAddress common.Address,
	eventChan chan<- model.ContractEvent, errChan chan<- error) {
	var txWorkerPool = pool.New().WithMaxGoroutines(config.MaxWorkersPerService)
	for _, tx := range block.Transactions() {
		txWorkerPool.Go(func() {
			cl.processBlockWorker(tx, contractAddress, block, eventChan, errChan)
		})
	}
	txWorkerPool.Wait()
}

func (cl *ContractListener) processBlockWorker(tx *ethTypes.Transaction, contractAddress common.Address,
	block *ethTypes.Block, eventChan chan<- model.ContractEvent, errChan chan<- error) {
	if tx.To() != nil && *tx.To() == contractAddress {
		cl.log.Debugf("Start parse TX hash: %v", tx.Hash().Hex())
		httpConnector := cl.chainConnector.GetHttpClient()
		txReceipt, err := httpConnector.TransactionReceipt(cl.ctx, tx.Hash())
		if err != nil {
			errChan <- err
			return
		} else if txReceipt.Status != 1 {
			cl.log.Debugf("Skip rejected transaction with hash: %s, its status: %v", tx.Hash().Hex(), txReceipt.Status)
			cl.log.Info("---------------------------------------------------")
			return
		}
		sender, _ := calcSender(tx)

		//Parse TX's events and find Trades in it
		var eventStart *model.TradeEvent
		var eventLast *model.TradeEvent
		var wgEvent conc.WaitGroup

		wgEvent.Go(func() { cl.runFindEventLog(txReceipt.Logs, tx, block, sender, false, &eventStart, errChan) })
		wgEvent.Go(func() { cl.runFindEventLog(txReceipt.Logs, tx, block, sender, true, &eventLast, errChan) })
		wgEvent.Wait()

		if eventStart == nil {
			cl.log.Debugf("Can't parse Transaction with Hash: %s", tx.Hash().Hex())
			return
		} else {
			if eventStart.PoolType == types.OneInchRouter {
				cl.log.Debugf("Skip. Transaction is full of OneInchRouter events: %s", tx.Hash().Hex())
				return
			} else if eventLast == nil {
				cl.log.Warning("Can't parse eventLast: %s", tx.Hash().Hex())
			}
		}

		var event *model.TradeEvent = eventStart
		event.ToTokenAddress = eventLast.ToTokenAddress
		event.ToTokenAmount = eventLast.ToTokenAmount
		//send data to channel
		eventChan <- event
	}
}

func (cl *ContractListener) runFindEventLog(txLogs []*ethTypes.Log, tx *ethTypes.Transaction, block *ethTypes.Block,
	sender *common.Address, reverse bool, event **model.TradeEvent, errChan chan<- error) {
	var err error
	*event, err = cl.findEventLog(txLogs, tx, block, sender, reverse)
	if err != nil {
		errChan <- err
	}
}

func (cl *ContractListener) findEventLog(logs []*ethTypes.Log, tx *ethTypes.Transaction, block *ethTypes.Block,
	sender *common.Address, reverse bool) (*model.TradeEvent, error) {
	var event *model.TradeEvent
	var err error
	parseEventLog := func(log *ethTypes.Log) {
		event, err = cl.parseEventLog(log, tx, block, sender)
	}
	if reverse {
		for i := len(logs) - 1; i >= 0 && event == nil && err == nil; i-- {
			parseEventLog(logs[i])
		}
	} else {
		for i := 0; i < len(logs) && event == nil && err == nil; i++ {
			parseEventLog(logs[i])
		}
	}
	return event, err
}

func (cl *ContractListener) parseEventLog(eventLog *ethTypes.Log, tx *ethTypes.Transaction, block *ethTypes.Block,
	sender *common.Address) (*model.TradeEvent, error) {
	for name, contractAbi := range cl.supportedContractAbi {
		eventMeta, _ := contractAbi.EventByID(eventLog.Topics[0])
		if eventMeta != nil &&
			(name == types.OneInchRouter || eventMeta.Name == contract.TradeMethodNameMap[name]) {
			event, err := helper.ParseSwapEvent(eventLog, tx, block, sender, name, cl.factoryAddress,
				cl.chainConnector.GetHttpClient())
			if err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, nil
}

func calcSender(tx *ethTypes.Transaction) (*common.Address, error) {
	msg, err := core.TransactionToMessage(tx, ethTypes.LatestSignerForChainID(tx.ChainId()), nil)
	if msg == nil {
		return nil, fmt.Errorf("can't parse the message from the transaction")
	}
	return &msg.From, err
}
