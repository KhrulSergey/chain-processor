// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swaap_vault

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IVaultBatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
type IVaultBatchSwapStep struct {
	PoolId        [32]byte
	AssetInIndex  *big.Int
	AssetOutIndex *big.Int
	Amount        *big.Int
	UserData      []byte
}

// IVaultExitPoolRequest is an auto generated low-level Go binding around an user-defined struct.
type IVaultExitPoolRequest struct {
	Assets            []common.Address
	MinAmountsOut     []*big.Int
	UserData          []byte
	ToInternalBalance bool
}

// IVaultFundManagement is an auto generated low-level Go binding around an user-defined struct.
type IVaultFundManagement struct {
	Sender              common.Address
	FromInternalBalance bool
	Recipient           common.Address
	ToInternalBalance   bool
}

// IVaultJoinPoolRequest is an auto generated low-level Go binding around an user-defined struct.
type IVaultJoinPoolRequest struct {
	Assets              []common.Address
	MaxAmountsIn        []*big.Int
	UserData            []byte
	FromInternalBalance bool
}

// IVaultPoolBalanceOp is an auto generated low-level Go binding around an user-defined struct.
type IVaultPoolBalanceOp struct {
	Kind   uint8
	PoolId [32]byte
	Token  common.Address
	Amount *big.Int
}

// IVaultSingleSwap is an auto generated low-level Go binding around an user-defined struct.
type IVaultSingleSwap struct {
	PoolId   [32]byte
	Kind     uint8
	AssetIn  common.Address
	AssetOut common.Address
	Amount   *big.Int
	UserData []byte
}

// IVaultUserBalanceOp is an auto generated low-level Go binding around an user-defined struct.
type IVaultUserBalanceOp struct {
	Kind      uint8
	Asset     common.Address
	Amount    *big.Int
	Sender    common.Address
	Recipient common.Address
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"AuthorizerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExternalBalanceTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"delta\",\"type\":\"int256\"}],\"name\":\"InternalBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"deltas\",\"type\":\"int256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"protocolFeeAmounts\",\"type\":\"uint256[]\"}],\"name\":\"PoolBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"cashDelta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"managedDelta\",\"type\":\"int256\"}],\"name\":\"PoolBalanceManaged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"specialization\",\"type\":\"uint8\"}],\"name\":\"PoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"RelayerApprovalChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"}],\"name\":\"TokensRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"int256[]\",\"name\":\"limits\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"batchSwap\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"assetDeltas\",\"type\":\"int256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"deregisterTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.ExitPoolRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"exitPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getInternalBalance\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPoolTokenInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"managed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"}],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"contractIProtocolFeesCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"hasApprovedRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.JoinPoolRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"joinPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.PoolBalanceOpKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIVault.PoolBalanceOp[]\",\"name\":\"ops\",\"type\":\"tuple[]\"}],\"name\":\"managePoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.UserBalanceOpKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIAsset\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIVault.UserBalanceOp[]\",\"name\":\"ops\",\"type\":\"tuple[]\"}],\"name\":\"manageUserBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAsset[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"}],\"name\":\"queryBatchSwap\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIVault.PoolSpecialization\",\"name\":\"specialization\",\"type\":\"uint8\"}],\"name\":\"registerPool\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"}],\"name\":\"registerTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"setAuthorizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setRelayerApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractIAsset\",\"name\":\"assetOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.SingleSwap\",\"name\":\"singleSwap\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIVault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountCalculated\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contract *ContractCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contract *ContractSession) WETH() (common.Address, error) {
	return _Contract.Contract.WETH(&_Contract.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contract *ContractCallerSession) WETH() (common.Address, error) {
	return _Contract.Contract.WETH(&_Contract.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Contract *ContractCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Contract *ContractSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Contract.Contract.GetActionId(&_Contract.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_Contract *ContractCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _Contract.Contract.GetActionId(&_Contract.CallOpts, selector)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Contract *ContractCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Contract *ContractSession) GetAuthorizer() (common.Address, error) {
	return _Contract.Contract.GetAuthorizer(&_Contract.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Contract *ContractCallerSession) GetAuthorizer() (common.Address, error) {
	return _Contract.Contract.GetAuthorizer(&_Contract.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Contract *ContractCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Contract *ContractSession) GetDomainSeparator() ([32]byte, error) {
	return _Contract.Contract.GetDomainSeparator(&_Contract.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Contract *ContractCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _Contract.Contract.GetDomainSeparator(&_Contract.CallOpts)
}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Contract *ContractCaller) GetInternalBalance(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getInternalBalance", user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Contract *ContractSession) GetInternalBalance(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetInternalBalance(&_Contract.CallOpts, user, tokens)
}

// GetInternalBalance is a free data retrieval call binding the contract method 0x0f5a6efa.
//
// Solidity: function getInternalBalance(address user, address[] tokens) view returns(uint256[] balances)
func (_Contract *ContractCallerSession) GetInternalBalance(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetInternalBalance(&_Contract.CallOpts, user, tokens)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_Contract *ContractCaller) GetNextNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNextNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_Contract *ContractSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetNextNonce(&_Contract.CallOpts, account)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_Contract *ContractCallerSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetNextNonce(&_Contract.CallOpts, account)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Contract *ContractCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPausedState")

	outstruct := new(struct {
		Paused              bool
		PauseWindowEndTime  *big.Int
		BufferPeriodEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PauseWindowEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodEndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Contract *ContractSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Contract.Contract.GetPausedState(&_Contract.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_Contract *ContractCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _Contract.Contract.GetPausedState(&_Contract.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Contract *ContractCaller) GetPool(opts *bind.CallOpts, poolId [32]byte) (common.Address, uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPool", poolId)

	if err != nil {
		return *new(common.Address), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Contract *ContractSession) GetPool(poolId [32]byte) (common.Address, uint8, error) {
	return _Contract.Contract.GetPool(&_Contract.CallOpts, poolId)
}

// GetPool is a free data retrieval call binding the contract method 0xf6c00927.
//
// Solidity: function getPool(bytes32 poolId) view returns(address, uint8)
func (_Contract *ContractCallerSession) GetPool(poolId [32]byte) (common.Address, uint8, error) {
	return _Contract.Contract.GetPool(&_Contract.CallOpts, poolId)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Contract *ContractCaller) GetPoolTokenInfo(opts *bind.CallOpts, poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPoolTokenInfo", poolId, token)

	outstruct := new(struct {
		Cash            *big.Int
		Managed         *big.Int
		LastChangeBlock *big.Int
		AssetManager    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cash = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Managed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastChangeBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AssetManager = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Contract *ContractSession) GetPoolTokenInfo(poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	return _Contract.Contract.GetPoolTokenInfo(&_Contract.CallOpts, poolId, token)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0xb05f8e48.
//
// Solidity: function getPoolTokenInfo(bytes32 poolId, address token) view returns(uint256 cash, uint256 managed, uint256 lastChangeBlock, address assetManager)
func (_Contract *ContractCallerSession) GetPoolTokenInfo(poolId [32]byte, token common.Address) (struct {
	Cash            *big.Int
	Managed         *big.Int
	LastChangeBlock *big.Int
	AssetManager    common.Address
}, error) {
	return _Contract.Contract.GetPoolTokenInfo(&_Contract.CallOpts, poolId, token)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Contract *ContractCaller) GetPoolTokens(opts *bind.CallOpts, poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPoolTokens", poolId)

	outstruct := new(struct {
		Tokens          []common.Address
		Balances        []*big.Int
		LastChangeBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Balances = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LastChangeBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Contract *ContractSession) GetPoolTokens(poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	return _Contract.Contract.GetPoolTokens(&_Contract.CallOpts, poolId)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xf94d4668.
//
// Solidity: function getPoolTokens(bytes32 poolId) view returns(address[] tokens, uint256[] balances, uint256 lastChangeBlock)
func (_Contract *ContractCallerSession) GetPoolTokens(poolId [32]byte) (struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}, error) {
	return _Contract.Contract.GetPoolTokens(&_Contract.CallOpts, poolId)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Contract *ContractCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Contract *ContractSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Contract.Contract.GetProtocolFeesCollector(&_Contract.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_Contract *ContractCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _Contract.Contract.GetProtocolFeesCollector(&_Contract.CallOpts)
}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Contract *ContractCaller) HasApprovedRelayer(opts *bind.CallOpts, user common.Address, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hasApprovedRelayer", user, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Contract *ContractSession) HasApprovedRelayer(user common.Address, relayer common.Address) (bool, error) {
	return _Contract.Contract.HasApprovedRelayer(&_Contract.CallOpts, user, relayer)
}

// HasApprovedRelayer is a free data retrieval call binding the contract method 0xfec90d72.
//
// Solidity: function hasApprovedRelayer(address user, address relayer) view returns(bool)
func (_Contract *ContractCallerSession) HasApprovedRelayer(user common.Address, relayer common.Address) (bool, error) {
	return _Contract.Contract.HasApprovedRelayer(&_Contract.CallOpts, user, relayer)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Contract *ContractTransactor) BatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "batchSwap", kind, swaps, assets, funds, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Contract *ContractSession) BatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BatchSwap(&_Contract.TransactOpts, kind, swaps, assets, funds, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds, int256[] limits, uint256 deadline) payable returns(int256[] assetDeltas)
func (_Contract *ContractTransactorSession) BatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BatchSwap(&_Contract.TransactOpts, kind, swaps, assets, funds, limits, deadline)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Contract *ContractTransactor) DeregisterTokens(opts *bind.TransactOpts, poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deregisterTokens", poolId, tokens)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Contract *ContractSession) DeregisterTokens(poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DeregisterTokens(&_Contract.TransactOpts, poolId, tokens)
}

// DeregisterTokens is a paid mutator transaction binding the contract method 0x7d3aeb96.
//
// Solidity: function deregisterTokens(bytes32 poolId, address[] tokens) returns()
func (_Contract *ContractTransactorSession) DeregisterTokens(poolId [32]byte, tokens []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DeregisterTokens(&_Contract.TransactOpts, poolId, tokens)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Contract *ContractTransactor) ExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "exitPool", poolId, sender, recipient, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Contract *ContractSession) ExitPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Contract.Contract.ExitPool(&_Contract.TransactOpts, poolId, sender, recipient, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) returns()
func (_Contract *ContractTransactorSession) ExitPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultExitPoolRequest) (*types.Transaction, error) {
	return _Contract.Contract.ExitPool(&_Contract.TransactOpts, poolId, sender, recipient, request)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Contract *ContractTransactor) FlashLoan(opts *bind.TransactOpts, recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "flashLoan", recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Contract *ContractSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Contract.Contract.FlashLoan(&_Contract.TransactOpts, recipient, tokens, amounts, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5c38449e.
//
// Solidity: function flashLoan(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Contract *ContractTransactorSession) FlashLoan(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Contract.Contract.FlashLoan(&_Contract.TransactOpts, recipient, tokens, amounts, userData)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Contract *ContractTransactor) JoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "joinPool", poolId, sender, recipient, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Contract *ContractSession) JoinPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Contract.Contract.JoinPool(&_Contract.TransactOpts, poolId, sender, recipient, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address sender, address recipient, (address[],uint256[],bytes,bool) request) payable returns()
func (_Contract *ContractTransactorSession) JoinPool(poolId [32]byte, sender common.Address, recipient common.Address, request IVaultJoinPoolRequest) (*types.Transaction, error) {
	return _Contract.Contract.JoinPool(&_Contract.TransactOpts, poolId, sender, recipient, request)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Contract *ContractTransactor) ManagePoolBalance(opts *bind.TransactOpts, ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "managePoolBalance", ops)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Contract *ContractSession) ManagePoolBalance(ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Contract.Contract.ManagePoolBalance(&_Contract.TransactOpts, ops)
}

// ManagePoolBalance is a paid mutator transaction binding the contract method 0xe6c46092.
//
// Solidity: function managePoolBalance((uint8,bytes32,address,uint256)[] ops) returns()
func (_Contract *ContractTransactorSession) ManagePoolBalance(ops []IVaultPoolBalanceOp) (*types.Transaction, error) {
	return _Contract.Contract.ManagePoolBalance(&_Contract.TransactOpts, ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Contract *ContractTransactor) ManageUserBalance(opts *bind.TransactOpts, ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "manageUserBalance", ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Contract *ContractSession) ManageUserBalance(ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Contract.Contract.ManageUserBalance(&_Contract.TransactOpts, ops)
}

// ManageUserBalance is a paid mutator transaction binding the contract method 0x0e8e3e84.
//
// Solidity: function manageUserBalance((uint8,address,uint256,address,address)[] ops) payable returns()
func (_Contract *ContractTransactorSession) ManageUserBalance(ops []IVaultUserBalanceOp) (*types.Transaction, error) {
	return _Contract.Contract.ManageUserBalance(&_Contract.TransactOpts, ops)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Contract *ContractTransactor) QueryBatchSwap(opts *bind.TransactOpts, kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "queryBatchSwap", kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Contract *ContractSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Contract.Contract.QueryBatchSwap(&_Contract.TransactOpts, kind, swaps, assets, funds)
}

// QueryBatchSwap is a paid mutator transaction binding the contract method 0xf84d066e.
//
// Solidity: function queryBatchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) funds) returns(int256[])
func (_Contract *ContractTransactorSession) QueryBatchSwap(kind uint8, swaps []IVaultBatchSwapStep, assets []common.Address, funds IVaultFundManagement) (*types.Transaction, error) {
	return _Contract.Contract.QueryBatchSwap(&_Contract.TransactOpts, kind, swaps, assets, funds)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Contract *ContractTransactor) RegisterPool(opts *bind.TransactOpts, specialization uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerPool", specialization)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Contract *ContractSession) RegisterPool(specialization uint8) (*types.Transaction, error) {
	return _Contract.Contract.RegisterPool(&_Contract.TransactOpts, specialization)
}

// RegisterPool is a paid mutator transaction binding the contract method 0x09b2760f.
//
// Solidity: function registerPool(uint8 specialization) returns(bytes32)
func (_Contract *ContractTransactorSession) RegisterPool(specialization uint8) (*types.Transaction, error) {
	return _Contract.Contract.RegisterPool(&_Contract.TransactOpts, specialization)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Contract *ContractTransactor) RegisterTokens(opts *bind.TransactOpts, poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerTokens", poolId, tokens, assetManagers)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Contract *ContractSession) RegisterTokens(poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RegisterTokens(&_Contract.TransactOpts, poolId, tokens, assetManagers)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0x66a9c7d2.
//
// Solidity: function registerTokens(bytes32 poolId, address[] tokens, address[] assetManagers) returns()
func (_Contract *ContractTransactorSession) RegisterTokens(poolId [32]byte, tokens []common.Address, assetManagers []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RegisterTokens(&_Contract.TransactOpts, poolId, tokens, assetManagers)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Contract *ContractTransactor) SetAuthorizer(opts *bind.TransactOpts, newAuthorizer common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAuthorizer", newAuthorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Contract *ContractSession) SetAuthorizer(newAuthorizer common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAuthorizer(&_Contract.TransactOpts, newAuthorizer)
}

// SetAuthorizer is a paid mutator transaction binding the contract method 0x058a628f.
//
// Solidity: function setAuthorizer(address newAuthorizer) returns()
func (_Contract *ContractTransactorSession) SetAuthorizer(newAuthorizer common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAuthorizer(&_Contract.TransactOpts, newAuthorizer)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Contract *ContractTransactor) SetPaused(opts *bind.TransactOpts, paused bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPaused", paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Contract *ContractSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPaused(&_Contract.TransactOpts, paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Contract *ContractTransactorSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPaused(&_Contract.TransactOpts, paused)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Contract *ContractTransactor) SetRelayerApproval(opts *bind.TransactOpts, sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRelayerApproval", sender, relayer, approved)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Contract *ContractSession) SetRelayerApproval(sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetRelayerApproval(&_Contract.TransactOpts, sender, relayer, approved)
}

// SetRelayerApproval is a paid mutator transaction binding the contract method 0xfa6e671d.
//
// Solidity: function setRelayerApproval(address sender, address relayer, bool approved) returns()
func (_Contract *ContractTransactorSession) SetRelayerApproval(sender common.Address, relayer common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetRelayerApproval(&_Contract.TransactOpts, sender, relayer, approved)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Contract *ContractTransactor) Swap(opts *bind.TransactOpts, singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "swap", singleSwap, funds, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Contract *ContractSession) Swap(singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Swap(&_Contract.TransactOpts, singleSwap, funds, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) funds, uint256 limit, uint256 deadline) payable returns(uint256 amountCalculated)
func (_Contract *ContractTransactorSession) Swap(singleSwap IVaultSingleSwap, funds IVaultFundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Swap(&_Contract.TransactOpts, singleSwap, funds, limit, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractAuthorizerChangedIterator is returned from FilterAuthorizerChanged and is used to iterate over the raw logs and unpacked data for AuthorizerChanged events raised by the Contract contract.
type ContractAuthorizerChangedIterator struct {
	Event *ContractAuthorizerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAuthorizerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAuthorizerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAuthorizerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAuthorizerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAuthorizerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAuthorizerChanged represents a AuthorizerChanged event raised by the Contract contract.
type ContractAuthorizerChanged struct {
	NewAuthorizer common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuthorizerChanged is a free log retrieval operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Contract *ContractFilterer) FilterAuthorizerChanged(opts *bind.FilterOpts, newAuthorizer []common.Address) (*ContractAuthorizerChangedIterator, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAuthorizerChangedIterator{contract: _Contract.contract, event: "AuthorizerChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizerChanged is a free log subscription operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Contract *ContractFilterer) WatchAuthorizerChanged(opts *bind.WatchOpts, sink chan<- *ContractAuthorizerChanged, newAuthorizer []common.Address) (event.Subscription, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAuthorizerChanged)
				if err := _Contract.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizerChanged is a log parse operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Contract *ContractFilterer) ParseAuthorizerChanged(log types.Log) (*ContractAuthorizerChanged, error) {
	event := new(ContractAuthorizerChanged)
	if err := _Contract.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractExternalBalanceTransferIterator is returned from FilterExternalBalanceTransfer and is used to iterate over the raw logs and unpacked data for ExternalBalanceTransfer events raised by the Contract contract.
type ContractExternalBalanceTransferIterator struct {
	Event *ContractExternalBalanceTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractExternalBalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExternalBalanceTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractExternalBalanceTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractExternalBalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExternalBalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExternalBalanceTransfer represents a ExternalBalanceTransfer event raised by the Contract contract.
type ContractExternalBalanceTransfer struct {
	Token     common.Address
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExternalBalanceTransfer is a free log retrieval operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Contract *ContractFilterer) FilterExternalBalanceTransfer(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*ContractExternalBalanceTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ExternalBalanceTransfer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractExternalBalanceTransferIterator{contract: _Contract.contract, event: "ExternalBalanceTransfer", logs: logs, sub: sub}, nil
}

// WatchExternalBalanceTransfer is a free log subscription operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Contract *ContractFilterer) WatchExternalBalanceTransfer(opts *bind.WatchOpts, sink chan<- *ContractExternalBalanceTransfer, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ExternalBalanceTransfer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExternalBalanceTransfer)
				if err := _Contract.contract.UnpackLog(event, "ExternalBalanceTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExternalBalanceTransfer is a log parse operation binding the contract event 0x540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c.
//
// Solidity: event ExternalBalanceTransfer(address indexed token, address indexed sender, address recipient, uint256 amount)
func (_Contract *ContractFilterer) ParseExternalBalanceTransfer(log types.Log) (*ContractExternalBalanceTransfer, error) {
	event := new(ContractExternalBalanceTransfer)
	if err := _Contract.contract.UnpackLog(event, "ExternalBalanceTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Contract contract.
type ContractFlashLoanIterator struct {
	Event *ContractFlashLoan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFlashLoan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFlashLoan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFlashLoan represents a FlashLoan event raised by the Contract contract.
type ContractFlashLoan struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Contract *ContractFilterer) FilterFlashLoan(opts *bind.FilterOpts, recipient []common.Address, token []common.Address) (*ContractFlashLoanIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractFlashLoanIterator{contract: _Contract.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Contract *ContractFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *ContractFlashLoan, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFlashLoan)
				if err := _Contract.contract.UnpackLog(event, "FlashLoan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashLoan is a log parse operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Contract *ContractFilterer) ParseFlashLoan(log types.Log) (*ContractFlashLoan, error) {
	event := new(ContractFlashLoan)
	if err := _Contract.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInternalBalanceChangedIterator is returned from FilterInternalBalanceChanged and is used to iterate over the raw logs and unpacked data for InternalBalanceChanged events raised by the Contract contract.
type ContractInternalBalanceChangedIterator struct {
	Event *ContractInternalBalanceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInternalBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInternalBalanceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInternalBalanceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInternalBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInternalBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInternalBalanceChanged represents a InternalBalanceChanged event raised by the Contract contract.
type ContractInternalBalanceChanged struct {
	User  common.Address
	Token common.Address
	Delta *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInternalBalanceChanged is a free log retrieval operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Contract *ContractFilterer) FilterInternalBalanceChanged(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*ContractInternalBalanceChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "InternalBalanceChanged", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractInternalBalanceChangedIterator{contract: _Contract.contract, event: "InternalBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchInternalBalanceChanged is a free log subscription operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Contract *ContractFilterer) WatchInternalBalanceChanged(opts *bind.WatchOpts, sink chan<- *ContractInternalBalanceChanged, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "InternalBalanceChanged", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInternalBalanceChanged)
				if err := _Contract.contract.UnpackLog(event, "InternalBalanceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInternalBalanceChanged is a log parse operation binding the contract event 0x18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42.
//
// Solidity: event InternalBalanceChanged(address indexed user, address indexed token, int256 delta)
func (_Contract *ContractFilterer) ParseInternalBalanceChanged(log types.Log) (*ContractInternalBalanceChanged, error) {
	event := new(ContractInternalBalanceChanged)
	if err := _Contract.contract.UnpackLog(event, "InternalBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the Contract contract.
type ContractPausedStateChangedIterator struct {
	Event *ContractPausedStateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPausedStateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPausedStateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPausedStateChanged represents a PausedStateChanged event raised by the Contract contract.
type ContractPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Contract *ContractFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*ContractPausedStateChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &ContractPausedStateChangedIterator{contract: _Contract.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Contract *ContractFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *ContractPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPausedStateChanged)
				if err := _Contract.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePausedStateChanged is a log parse operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_Contract *ContractFilterer) ParsePausedStateChanged(log types.Log) (*ContractPausedStateChanged, error) {
	event := new(ContractPausedStateChanged)
	if err := _Contract.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPoolBalanceChangedIterator is returned from FilterPoolBalanceChanged and is used to iterate over the raw logs and unpacked data for PoolBalanceChanged events raised by the Contract contract.
type ContractPoolBalanceChangedIterator struct {
	Event *ContractPoolBalanceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPoolBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPoolBalanceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPoolBalanceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPoolBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPoolBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPoolBalanceChanged represents a PoolBalanceChanged event raised by the Contract contract.
type ContractPoolBalanceChanged struct {
	PoolId             [32]byte
	LiquidityProvider  common.Address
	Tokens             []common.Address
	Deltas             []*big.Int
	ProtocolFeeAmounts []*big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoolBalanceChanged is a free log retrieval operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Contract *ContractFilterer) FilterPoolBalanceChanged(opts *bind.FilterOpts, poolId [][32]byte, liquidityProvider []common.Address) (*ContractPoolBalanceChangedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PoolBalanceChanged", poolIdRule, liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return &ContractPoolBalanceChangedIterator{contract: _Contract.contract, event: "PoolBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchPoolBalanceChanged is a free log subscription operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Contract *ContractFilterer) WatchPoolBalanceChanged(opts *bind.WatchOpts, sink chan<- *ContractPoolBalanceChanged, poolId [][32]byte, liquidityProvider []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PoolBalanceChanged", poolIdRule, liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPoolBalanceChanged)
				if err := _Contract.contract.UnpackLog(event, "PoolBalanceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolBalanceChanged is a log parse operation binding the contract event 0xe5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78.
//
// Solidity: event PoolBalanceChanged(bytes32 indexed poolId, address indexed liquidityProvider, address[] tokens, int256[] deltas, uint256[] protocolFeeAmounts)
func (_Contract *ContractFilterer) ParsePoolBalanceChanged(log types.Log) (*ContractPoolBalanceChanged, error) {
	event := new(ContractPoolBalanceChanged)
	if err := _Contract.contract.UnpackLog(event, "PoolBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPoolBalanceManagedIterator is returned from FilterPoolBalanceManaged and is used to iterate over the raw logs and unpacked data for PoolBalanceManaged events raised by the Contract contract.
type ContractPoolBalanceManagedIterator struct {
	Event *ContractPoolBalanceManaged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPoolBalanceManagedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPoolBalanceManaged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPoolBalanceManaged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPoolBalanceManagedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPoolBalanceManagedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPoolBalanceManaged represents a PoolBalanceManaged event raised by the Contract contract.
type ContractPoolBalanceManaged struct {
	PoolId       [32]byte
	AssetManager common.Address
	Token        common.Address
	CashDelta    *big.Int
	ManagedDelta *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolBalanceManaged is a free log retrieval operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Contract *ContractFilterer) FilterPoolBalanceManaged(opts *bind.FilterOpts, poolId [][32]byte, assetManager []common.Address, token []common.Address) (*ContractPoolBalanceManagedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var assetManagerRule []interface{}
	for _, assetManagerItem := range assetManager {
		assetManagerRule = append(assetManagerRule, assetManagerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PoolBalanceManaged", poolIdRule, assetManagerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractPoolBalanceManagedIterator{contract: _Contract.contract, event: "PoolBalanceManaged", logs: logs, sub: sub}, nil
}

// WatchPoolBalanceManaged is a free log subscription operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Contract *ContractFilterer) WatchPoolBalanceManaged(opts *bind.WatchOpts, sink chan<- *ContractPoolBalanceManaged, poolId [][32]byte, assetManager []common.Address, token []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var assetManagerRule []interface{}
	for _, assetManagerItem := range assetManager {
		assetManagerRule = append(assetManagerRule, assetManagerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PoolBalanceManaged", poolIdRule, assetManagerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPoolBalanceManaged)
				if err := _Contract.contract.UnpackLog(event, "PoolBalanceManaged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolBalanceManaged is a log parse operation binding the contract event 0x6edcaf6241105b4c94c2efdbf3a6b12458eb3d07be3a0e81d24b13c44045fe7a.
//
// Solidity: event PoolBalanceManaged(bytes32 indexed poolId, address indexed assetManager, address indexed token, int256 cashDelta, int256 managedDelta)
func (_Contract *ContractFilterer) ParsePoolBalanceManaged(log types.Log) (*ContractPoolBalanceManaged, error) {
	event := new(ContractPoolBalanceManaged)
	if err := _Contract.contract.UnpackLog(event, "PoolBalanceManaged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPoolRegisteredIterator is returned from FilterPoolRegistered and is used to iterate over the raw logs and unpacked data for PoolRegistered events raised by the Contract contract.
type ContractPoolRegisteredIterator struct {
	Event *ContractPoolRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPoolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPoolRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPoolRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPoolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPoolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPoolRegistered represents a PoolRegistered event raised by the Contract contract.
type ContractPoolRegistered struct {
	PoolId         [32]byte
	PoolAddress    common.Address
	Specialization uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolRegistered is a free log retrieval operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Contract *ContractFilterer) FilterPoolRegistered(opts *bind.FilterOpts, poolId [][32]byte, poolAddress []common.Address) (*ContractPoolRegisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PoolRegistered", poolIdRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractPoolRegisteredIterator{contract: _Contract.contract, event: "PoolRegistered", logs: logs, sub: sub}, nil
}

// WatchPoolRegistered is a free log subscription operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Contract *ContractFilterer) WatchPoolRegistered(opts *bind.WatchOpts, sink chan<- *ContractPoolRegistered, poolId [][32]byte, poolAddress []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PoolRegistered", poolIdRule, poolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPoolRegistered)
				if err := _Contract.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolRegistered is a log parse operation binding the contract event 0x3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e.
//
// Solidity: event PoolRegistered(bytes32 indexed poolId, address indexed poolAddress, uint8 specialization)
func (_Contract *ContractFilterer) ParsePoolRegistered(log types.Log) (*ContractPoolRegistered, error) {
	event := new(ContractPoolRegistered)
	if err := _Contract.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRelayerApprovalChangedIterator is returned from FilterRelayerApprovalChanged and is used to iterate over the raw logs and unpacked data for RelayerApprovalChanged events raised by the Contract contract.
type ContractRelayerApprovalChangedIterator struct {
	Event *ContractRelayerApprovalChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRelayerApprovalChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRelayerApprovalChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRelayerApprovalChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRelayerApprovalChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRelayerApprovalChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRelayerApprovalChanged represents a RelayerApprovalChanged event raised by the Contract contract.
type ContractRelayerApprovalChanged struct {
	Relayer  common.Address
	Sender   common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerApprovalChanged is a free log retrieval operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Contract *ContractFilterer) FilterRelayerApprovalChanged(opts *bind.FilterOpts, relayer []common.Address, sender []common.Address) (*ContractRelayerApprovalChangedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RelayerApprovalChanged", relayerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractRelayerApprovalChangedIterator{contract: _Contract.contract, event: "RelayerApprovalChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerApprovalChanged is a free log subscription operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Contract *ContractFilterer) WatchRelayerApprovalChanged(opts *bind.WatchOpts, sink chan<- *ContractRelayerApprovalChanged, relayer []common.Address, sender []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RelayerApprovalChanged", relayerRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRelayerApprovalChanged)
				if err := _Contract.contract.UnpackLog(event, "RelayerApprovalChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerApprovalChanged is a log parse operation binding the contract event 0x46961fdb4502b646d5095fba7600486a8ac05041d55cdf0f16ed677180b5cad8.
//
// Solidity: event RelayerApprovalChanged(address indexed relayer, address indexed sender, bool approved)
func (_Contract *ContractFilterer) ParseRelayerApprovalChanged(log types.Log) (*ContractRelayerApprovalChanged, error) {
	event := new(ContractRelayerApprovalChanged)
	if err := _Contract.contract.UnpackLog(event, "RelayerApprovalChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Contract contract.
type ContractSwapIterator struct {
	Event *ContractSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSwap represents a Swap event raised by the Contract contract.
type ContractSwap struct {
	PoolId    [32]byte
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Contract *ContractFilterer) FilterSwap(opts *bind.FilterOpts, poolId [][32]byte, tokenIn []common.Address, tokenOut []common.Address) (*ContractSwapIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Swap", poolIdRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &ContractSwapIterator{contract: _Contract.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Contract *ContractFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ContractSwap, poolId [][32]byte, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Swap", poolIdRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSwap)
				if err := _Contract.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0x2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b.
//
// Solidity: event Swap(bytes32 indexed poolId, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut)
func (_Contract *ContractFilterer) ParseSwap(log types.Log) (*ContractSwap, error) {
	event := new(ContractSwap)
	if err := _Contract.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTokensDeregisteredIterator is returned from FilterTokensDeregistered and is used to iterate over the raw logs and unpacked data for TokensDeregistered events raised by the Contract contract.
type ContractTokensDeregisteredIterator struct {
	Event *ContractTokensDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTokensDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTokensDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTokensDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTokensDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTokensDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTokensDeregistered represents a TokensDeregistered event raised by the Contract contract.
type ContractTokensDeregistered struct {
	PoolId [32]byte
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensDeregistered is a free log retrieval operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Contract *ContractFilterer) FilterTokensDeregistered(opts *bind.FilterOpts, poolId [][32]byte) (*ContractTokensDeregisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TokensDeregistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTokensDeregisteredIterator{contract: _Contract.contract, event: "TokensDeregistered", logs: logs, sub: sub}, nil
}

// WatchTokensDeregistered is a free log subscription operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Contract *ContractFilterer) WatchTokensDeregistered(opts *bind.WatchOpts, sink chan<- *ContractTokensDeregistered, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TokensDeregistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTokensDeregistered)
				if err := _Contract.contract.UnpackLog(event, "TokensDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensDeregistered is a log parse operation binding the contract event 0x7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610.
//
// Solidity: event TokensDeregistered(bytes32 indexed poolId, address[] tokens)
func (_Contract *ContractFilterer) ParseTokensDeregistered(log types.Log) (*ContractTokensDeregistered, error) {
	event := new(ContractTokensDeregistered)
	if err := _Contract.contract.UnpackLog(event, "TokensDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTokensRegisteredIterator is returned from FilterTokensRegistered and is used to iterate over the raw logs and unpacked data for TokensRegistered events raised by the Contract contract.
type ContractTokensRegisteredIterator struct {
	Event *ContractTokensRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTokensRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTokensRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTokensRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTokensRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTokensRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTokensRegistered represents a TokensRegistered event raised by the Contract contract.
type ContractTokensRegistered struct {
	PoolId        [32]byte
	Tokens        []common.Address
	AssetManagers []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensRegistered is a free log retrieval operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Contract *ContractFilterer) FilterTokensRegistered(opts *bind.FilterOpts, poolId [][32]byte) (*ContractTokensRegisteredIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TokensRegistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTokensRegisteredIterator{contract: _Contract.contract, event: "TokensRegistered", logs: logs, sub: sub}, nil
}

// WatchTokensRegistered is a free log subscription operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Contract *ContractFilterer) WatchTokensRegistered(opts *bind.WatchOpts, sink chan<- *ContractTokensRegistered, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TokensRegistered", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTokensRegistered)
				if err := _Contract.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensRegistered is a log parse operation binding the contract event 0xf5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4.
//
// Solidity: event TokensRegistered(bytes32 indexed poolId, address[] tokens, address[] assetManagers)
func (_Contract *ContractFilterer) ParseTokensRegistered(log types.Log) (*ContractTokensRegistered, error) {
	event := new(ContractTokensRegistered)
	if err := _Contract.contract.UnpackLog(event, "TokensRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
