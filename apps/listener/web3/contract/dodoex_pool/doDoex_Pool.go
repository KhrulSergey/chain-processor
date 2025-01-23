// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dodoex_pool

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

// PMMPricingPMMState is an auto generated low-level Go binding around an user-defined struct.
type PMMPricingPMMState struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  uint8
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"DODOFlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"DODOSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLpFeeRate\",\"type\":\"uint256\"}],\"name\":\"LpFeeRateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumPMMPricing.RState\",\"name\":\"newRState\",\"type\":\"uint8\"}],\"name\":\"RChange\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_BASE_PRICE_CUMULATIVE_LAST_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_RESERVE_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_TARGET_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_TOKEN_\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BLOCK_TIMESTAMP_LAST_\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_IS_OPEN_TWAP_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_I_\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_K_\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_LP_FEE_RATE_\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MAINTAINER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MT_FEE_RATE_MODEL_\",\"outputs\":[{\"internalType\":\"contractIFeeRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_NEW_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_RESERVE_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_TARGET_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_TOKEN_\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_RState_\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetTo\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"input\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"midPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPMMState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q0\",\"type\":\"uint256\"},{\"internalType\":\"enumPMMPricing.RState\",\"name\":\"R\",\"type\":\"uint8\"}],\"internalType\":\"structPMMPricing.PMMState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPMMStateForCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuoteInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"input\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteReserve\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maintainer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mtFeeRateModel\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpenTWAP\",\"type\":\"bool\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payBaseAmount\",\"type\":\"uint256\"}],\"name\":\"querySellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFee\",\"type\":\"uint256\"},{\"internalType\":\"enumPMMPricing.RState\",\"name\":\"newRState\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newBaseTarget\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payQuoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFee\",\"type\":\"uint256\"},{\"internalType\":\"enumPMMPricing.RState\",\"name\":\"newRState\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newQuoteTarget\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratioSync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newLpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newK\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteReserve\",\"type\":\"uint256\"}],\"name\":\"reset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"retrieve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveQuoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveBaseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newK\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteReserve\",\"type\":\"uint256\"}],\"name\":\"tuneParameters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteReserve\",\"type\":\"uint256\"}],\"name\":\"tunePrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// BASEPRICECUMULATIVELAST is a free data retrieval call binding the contract method 0xfe24cb7f.
//
// Solidity: function _BASE_PRICE_CUMULATIVE_LAST_() view returns(uint256)
func (_Contract *ContractCaller) BASEPRICECUMULATIVELAST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_BASE_PRICE_CUMULATIVE_LAST_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEPRICECUMULATIVELAST is a free data retrieval call binding the contract method 0xfe24cb7f.
//
// Solidity: function _BASE_PRICE_CUMULATIVE_LAST_() view returns(uint256)
func (_Contract *ContractSession) BASEPRICECUMULATIVELAST() (*big.Int, error) {
	return _Contract.Contract.BASEPRICECUMULATIVELAST(&_Contract.CallOpts)
}

// BASEPRICECUMULATIVELAST is a free data retrieval call binding the contract method 0xfe24cb7f.
//
// Solidity: function _BASE_PRICE_CUMULATIVE_LAST_() view returns(uint256)
func (_Contract *ContractCallerSession) BASEPRICECUMULATIVELAST() (*big.Int, error) {
	return _Contract.Contract.BASEPRICECUMULATIVELAST(&_Contract.CallOpts)
}

// BASERESERVE is a free data retrieval call binding the contract method 0x7d721504.
//
// Solidity: function _BASE_RESERVE_() view returns(uint112)
func (_Contract *ContractCaller) BASERESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_BASE_RESERVE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASERESERVE is a free data retrieval call binding the contract method 0x7d721504.
//
// Solidity: function _BASE_RESERVE_() view returns(uint112)
func (_Contract *ContractSession) BASERESERVE() (*big.Int, error) {
	return _Contract.Contract.BASERESERVE(&_Contract.CallOpts)
}

// BASERESERVE is a free data retrieval call binding the contract method 0x7d721504.
//
// Solidity: function _BASE_RESERVE_() view returns(uint112)
func (_Contract *ContractCallerSession) BASERESERVE() (*big.Int, error) {
	return _Contract.Contract.BASERESERVE(&_Contract.CallOpts)
}

// BASETARGET is a free data retrieval call binding the contract method 0xe539ef49.
//
// Solidity: function _BASE_TARGET_() view returns(uint112)
func (_Contract *ContractCaller) BASETARGET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_BASE_TARGET_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASETARGET is a free data retrieval call binding the contract method 0xe539ef49.
//
// Solidity: function _BASE_TARGET_() view returns(uint112)
func (_Contract *ContractSession) BASETARGET() (*big.Int, error) {
	return _Contract.Contract.BASETARGET(&_Contract.CallOpts)
}

// BASETARGET is a free data retrieval call binding the contract method 0xe539ef49.
//
// Solidity: function _BASE_TARGET_() view returns(uint112)
func (_Contract *ContractCallerSession) BASETARGET() (*big.Int, error) {
	return _Contract.Contract.BASETARGET(&_Contract.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_Contract *ContractCaller) BASETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_BASE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_Contract *ContractSession) BASETOKEN() (common.Address, error) {
	return _Contract.Contract.BASETOKEN(&_Contract.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_Contract *ContractCallerSession) BASETOKEN() (common.Address, error) {
	return _Contract.Contract.BASETOKEN(&_Contract.CallOpts)
}

// BLOCKTIMESTAMPLAST is a free data retrieval call binding the contract method 0x880a4d87.
//
// Solidity: function _BLOCK_TIMESTAMP_LAST_() view returns(uint32)
func (_Contract *ContractCaller) BLOCKTIMESTAMPLAST(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_BLOCK_TIMESTAMP_LAST_")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BLOCKTIMESTAMPLAST is a free data retrieval call binding the contract method 0x880a4d87.
//
// Solidity: function _BLOCK_TIMESTAMP_LAST_() view returns(uint32)
func (_Contract *ContractSession) BLOCKTIMESTAMPLAST() (uint32, error) {
	return _Contract.Contract.BLOCKTIMESTAMPLAST(&_Contract.CallOpts)
}

// BLOCKTIMESTAMPLAST is a free data retrieval call binding the contract method 0x880a4d87.
//
// Solidity: function _BLOCK_TIMESTAMP_LAST_() view returns(uint32)
func (_Contract *ContractCallerSession) BLOCKTIMESTAMPLAST() (uint32, error) {
	return _Contract.Contract.BLOCKTIMESTAMPLAST(&_Contract.CallOpts)
}

// ISOPENTWAP is a free data retrieval call binding the contract method 0x2df6cb48.
//
// Solidity: function _IS_OPEN_TWAP_() view returns(bool)
func (_Contract *ContractCaller) ISOPENTWAP(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_IS_OPEN_TWAP_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISOPENTWAP is a free data retrieval call binding the contract method 0x2df6cb48.
//
// Solidity: function _IS_OPEN_TWAP_() view returns(bool)
func (_Contract *ContractSession) ISOPENTWAP() (bool, error) {
	return _Contract.Contract.ISOPENTWAP(&_Contract.CallOpts)
}

// ISOPENTWAP is a free data retrieval call binding the contract method 0x2df6cb48.
//
// Solidity: function _IS_OPEN_TWAP_() view returns(bool)
func (_Contract *ContractCallerSession) ISOPENTWAP() (bool, error) {
	return _Contract.Contract.ISOPENTWAP(&_Contract.CallOpts)
}

// I is a free data retrieval call binding the contract method 0xf811d692.
//
// Solidity: function _I_() view returns(uint128)
func (_Contract *ContractCaller) I(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_I_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// I is a free data retrieval call binding the contract method 0xf811d692.
//
// Solidity: function _I_() view returns(uint128)
func (_Contract *ContractSession) I() (*big.Int, error) {
	return _Contract.Contract.I(&_Contract.CallOpts)
}

// I is a free data retrieval call binding the contract method 0xf811d692.
//
// Solidity: function _I_() view returns(uint128)
func (_Contract *ContractCallerSession) I() (*big.Int, error) {
	return _Contract.Contract.I(&_Contract.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint64)
func (_Contract *ContractCaller) K(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_K_")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint64)
func (_Contract *ContractSession) K() (uint64, error) {
	return _Contract.Contract.K(&_Contract.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint64)
func (_Contract *ContractCallerSession) K() (uint64, error) {
	return _Contract.Contract.K(&_Contract.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint64)
func (_Contract *ContractCaller) LPFEERATE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_LP_FEE_RATE_")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint64)
func (_Contract *ContractSession) LPFEERATE() (uint64, error) {
	return _Contract.Contract.LPFEERATE(&_Contract.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint64)
func (_Contract *ContractCallerSession) LPFEERATE() (uint64, error) {
	return _Contract.Contract.LPFEERATE(&_Contract.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_Contract *ContractCaller) MAINTAINER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_MAINTAINER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_Contract *ContractSession) MAINTAINER() (common.Address, error) {
	return _Contract.Contract.MAINTAINER(&_Contract.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_Contract *ContractCallerSession) MAINTAINER() (common.Address, error) {
	return _Contract.Contract.MAINTAINER(&_Contract.CallOpts)
}

// MTFEERATEMODEL is a free data retrieval call binding the contract method 0xf6b06e70.
//
// Solidity: function _MT_FEE_RATE_MODEL_() view returns(address)
func (_Contract *ContractCaller) MTFEERATEMODEL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_MT_FEE_RATE_MODEL_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MTFEERATEMODEL is a free data retrieval call binding the contract method 0xf6b06e70.
//
// Solidity: function _MT_FEE_RATE_MODEL_() view returns(address)
func (_Contract *ContractSession) MTFEERATEMODEL() (common.Address, error) {
	return _Contract.Contract.MTFEERATEMODEL(&_Contract.CallOpts)
}

// MTFEERATEMODEL is a free data retrieval call binding the contract method 0xf6b06e70.
//
// Solidity: function _MT_FEE_RATE_MODEL_() view returns(address)
func (_Contract *ContractCallerSession) MTFEERATEMODEL() (common.Address, error) {
	return _Contract.Contract.MTFEERATEMODEL(&_Contract.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_Contract *ContractCaller) NEWOWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_NEW_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_Contract *ContractSession) NEWOWNER() (common.Address, error) {
	return _Contract.Contract.NEWOWNER(&_Contract.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_Contract *ContractCallerSession) NEWOWNER() (common.Address, error) {
	return _Contract.Contract.NEWOWNER(&_Contract.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_Contract *ContractCaller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_Contract *ContractSession) OWNER() (common.Address, error) {
	return _Contract.Contract.OWNER(&_Contract.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_Contract *ContractCallerSession) OWNER() (common.Address, error) {
	return _Contract.Contract.OWNER(&_Contract.CallOpts)
}

// QUOTERESERVE is a free data retrieval call binding the contract method 0xbbf5ce78.
//
// Solidity: function _QUOTE_RESERVE_() view returns(uint112)
func (_Contract *ContractCaller) QUOTERESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_QUOTE_RESERVE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTERESERVE is a free data retrieval call binding the contract method 0xbbf5ce78.
//
// Solidity: function _QUOTE_RESERVE_() view returns(uint112)
func (_Contract *ContractSession) QUOTERESERVE() (*big.Int, error) {
	return _Contract.Contract.QUOTERESERVE(&_Contract.CallOpts)
}

// QUOTERESERVE is a free data retrieval call binding the contract method 0xbbf5ce78.
//
// Solidity: function _QUOTE_RESERVE_() view returns(uint112)
func (_Contract *ContractCallerSession) QUOTERESERVE() (*big.Int, error) {
	return _Contract.Contract.QUOTERESERVE(&_Contract.CallOpts)
}

// QUOTETARGET is a free data retrieval call binding the contract method 0x77f58657.
//
// Solidity: function _QUOTE_TARGET_() view returns(uint112)
func (_Contract *ContractCaller) QUOTETARGET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_QUOTE_TARGET_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTETARGET is a free data retrieval call binding the contract method 0x77f58657.
//
// Solidity: function _QUOTE_TARGET_() view returns(uint112)
func (_Contract *ContractSession) QUOTETARGET() (*big.Int, error) {
	return _Contract.Contract.QUOTETARGET(&_Contract.CallOpts)
}

// QUOTETARGET is a free data retrieval call binding the contract method 0x77f58657.
//
// Solidity: function _QUOTE_TARGET_() view returns(uint112)
func (_Contract *ContractCallerSession) QUOTETARGET() (*big.Int, error) {
	return _Contract.Contract.QUOTETARGET(&_Contract.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_Contract *ContractCaller) QUOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_QUOTE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_Contract *ContractSession) QUOTETOKEN() (common.Address, error) {
	return _Contract.Contract.QUOTETOKEN(&_Contract.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_Contract *ContractCallerSession) QUOTETOKEN() (common.Address, error) {
	return _Contract.Contract.QUOTETOKEN(&_Contract.CallOpts)
}

// RState is a free data retrieval call binding the contract method 0xbf357dae.
//
// Solidity: function _RState_() view returns(uint32)
func (_Contract *ContractCaller) RState(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_RState_")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RState is a free data retrieval call binding the contract method 0xbf357dae.
//
// Solidity: function _RState_() view returns(uint32)
func (_Contract *ContractSession) RState() (uint32, error) {
	return _Contract.Contract.RState(&_Contract.CallOpts)
}

// RState is a free data retrieval call binding the contract method 0xbf357dae.
//
// Solidity: function _RState_() view returns(uint32)
func (_Contract *ContractCallerSession) RState() (uint32, error) {
	return _Contract.Contract.RState(&_Contract.CallOpts)
}

// GetBaseInput is a free data retrieval call binding the contract method 0x65f6fcbb.
//
// Solidity: function getBaseInput() view returns(uint256 input)
func (_Contract *ContractCaller) GetBaseInput(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBaseInput")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseInput is a free data retrieval call binding the contract method 0x65f6fcbb.
//
// Solidity: function getBaseInput() view returns(uint256 input)
func (_Contract *ContractSession) GetBaseInput() (*big.Int, error) {
	return _Contract.Contract.GetBaseInput(&_Contract.CallOpts)
}

// GetBaseInput is a free data retrieval call binding the contract method 0x65f6fcbb.
//
// Solidity: function getBaseInput() view returns(uint256 input)
func (_Contract *ContractCallerSession) GetBaseInput() (*big.Int, error) {
	return _Contract.Contract.GetBaseInput(&_Contract.CallOpts)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_Contract *ContractCaller) GetMidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_Contract *ContractSession) GetMidPrice() (*big.Int, error) {
	return _Contract.Contract.GetMidPrice(&_Contract.CallOpts)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_Contract *ContractCallerSession) GetMidPrice() (*big.Int, error) {
	return _Contract.Contract.GetMidPrice(&_Contract.CallOpts)
}

// GetPMMState is a free data retrieval call binding the contract method 0xa382d1b9.
//
// Solidity: function getPMMState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint8) state)
func (_Contract *ContractCaller) GetPMMState(opts *bind.CallOpts) (PMMPricingPMMState, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPMMState")

	if err != nil {
		return *new(PMMPricingPMMState), err
	}

	out0 := *abi.ConvertType(out[0], new(PMMPricingPMMState)).(*PMMPricingPMMState)

	return out0, err

}

// GetPMMState is a free data retrieval call binding the contract method 0xa382d1b9.
//
// Solidity: function getPMMState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint8) state)
func (_Contract *ContractSession) GetPMMState() (PMMPricingPMMState, error) {
	return _Contract.Contract.GetPMMState(&_Contract.CallOpts)
}

// GetPMMState is a free data retrieval call binding the contract method 0xa382d1b9.
//
// Solidity: function getPMMState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint8) state)
func (_Contract *ContractCallerSession) GetPMMState() (PMMPricingPMMState, error) {
	return _Contract.Contract.GetPMMState(&_Contract.CallOpts)
}

// GetPMMStateForCall is a free data retrieval call binding the contract method 0xfd1ed7e9.
//
// Solidity: function getPMMStateForCall() view returns(uint256 i, uint256 K, uint256 B, uint256 Q, uint256 B0, uint256 Q0, uint256 R)
func (_Contract *ContractCaller) GetPMMStateForCall(opts *bind.CallOpts) (struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPMMStateForCall")

	outstruct := new(struct {
		I  *big.Int
		K  *big.Int
		B  *big.Int
		Q  *big.Int
		B0 *big.Int
		Q0 *big.Int
		R  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.I = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.K = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.B = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Q = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.B0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Q0 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.R = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPMMStateForCall is a free data retrieval call binding the contract method 0xfd1ed7e9.
//
// Solidity: function getPMMStateForCall() view returns(uint256 i, uint256 K, uint256 B, uint256 Q, uint256 B0, uint256 Q0, uint256 R)
func (_Contract *ContractSession) GetPMMStateForCall() (struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  *big.Int
}, error) {
	return _Contract.Contract.GetPMMStateForCall(&_Contract.CallOpts)
}

// GetPMMStateForCall is a free data retrieval call binding the contract method 0xfd1ed7e9.
//
// Solidity: function getPMMStateForCall() view returns(uint256 i, uint256 K, uint256 B, uint256 Q, uint256 B0, uint256 Q0, uint256 R)
func (_Contract *ContractCallerSession) GetPMMStateForCall() (struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  *big.Int
}, error) {
	return _Contract.Contract.GetPMMStateForCall(&_Contract.CallOpts)
}

// GetQuoteInput is a free data retrieval call binding the contract method 0x71f9100c.
//
// Solidity: function getQuoteInput() view returns(uint256 input)
func (_Contract *ContractCaller) GetQuoteInput(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getQuoteInput")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteInput is a free data retrieval call binding the contract method 0x71f9100c.
//
// Solidity: function getQuoteInput() view returns(uint256 input)
func (_Contract *ContractSession) GetQuoteInput() (*big.Int, error) {
	return _Contract.Contract.GetQuoteInput(&_Contract.CallOpts)
}

// GetQuoteInput is a free data retrieval call binding the contract method 0x71f9100c.
//
// Solidity: function getQuoteInput() view returns(uint256 input)
func (_Contract *ContractCallerSession) GetQuoteInput() (*big.Int, error) {
	return _Contract.Contract.GetQuoteInput(&_Contract.CallOpts)
}

// GetUserFeeRate is a free data retrieval call binding the contract method 0x44096609.
//
// Solidity: function getUserFeeRate(address user) view returns(uint256 lpFeeRate, uint256 mtFeeRate)
func (_Contract *ContractCaller) GetUserFeeRate(opts *bind.CallOpts, user common.Address) (struct {
	LpFeeRate *big.Int
	MtFeeRate *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUserFeeRate", user)

	outstruct := new(struct {
		LpFeeRate *big.Int
		MtFeeRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpFeeRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MtFeeRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserFeeRate is a free data retrieval call binding the contract method 0x44096609.
//
// Solidity: function getUserFeeRate(address user) view returns(uint256 lpFeeRate, uint256 mtFeeRate)
func (_Contract *ContractSession) GetUserFeeRate(user common.Address) (struct {
	LpFeeRate *big.Int
	MtFeeRate *big.Int
}, error) {
	return _Contract.Contract.GetUserFeeRate(&_Contract.CallOpts, user)
}

// GetUserFeeRate is a free data retrieval call binding the contract method 0x44096609.
//
// Solidity: function getUserFeeRate(address user) view returns(uint256 lpFeeRate, uint256 mtFeeRate)
func (_Contract *ContractCallerSession) GetUserFeeRate(user common.Address) (struct {
	LpFeeRate *big.Int
	MtFeeRate *big.Int
}, error) {
	return _Contract.Contract.GetUserFeeRate(&_Contract.CallOpts, user)
}

// GetVaultReserve is a free data retrieval call binding the contract method 0x36223ce9.
//
// Solidity: function getVaultReserve() view returns(uint256 baseReserve, uint256 quoteReserve)
func (_Contract *ContractCaller) GetVaultReserve(opts *bind.CallOpts) (struct {
	BaseReserve  *big.Int
	QuoteReserve *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVaultReserve")

	outstruct := new(struct {
		BaseReserve  *big.Int
		QuoteReserve *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseReserve = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QuoteReserve = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVaultReserve is a free data retrieval call binding the contract method 0x36223ce9.
//
// Solidity: function getVaultReserve() view returns(uint256 baseReserve, uint256 quoteReserve)
func (_Contract *ContractSession) GetVaultReserve() (struct {
	BaseReserve  *big.Int
	QuoteReserve *big.Int
}, error) {
	return _Contract.Contract.GetVaultReserve(&_Contract.CallOpts)
}

// GetVaultReserve is a free data retrieval call binding the contract method 0x36223ce9.
//
// Solidity: function getVaultReserve() view returns(uint256 baseReserve, uint256 quoteReserve)
func (_Contract *ContractCallerSession) GetVaultReserve() (struct {
	BaseReserve  *big.Int
	QuoteReserve *big.Int
}, error) {
	return _Contract.Contract.GetVaultReserve(&_Contract.CallOpts)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address trader, uint256 payBaseAmount) view returns(uint256 receiveQuoteAmount, uint256 mtFee, uint8 newRState, uint256 newBaseTarget)
func (_Contract *ContractCaller) QuerySellBase(opts *bind.CallOpts, trader common.Address, payBaseAmount *big.Int) (struct {
	ReceiveQuoteAmount *big.Int
	MtFee              *big.Int
	NewRState          uint8
	NewBaseTarget      *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "querySellBase", trader, payBaseAmount)

	outstruct := new(struct {
		ReceiveQuoteAmount *big.Int
		MtFee              *big.Int
		NewRState          uint8
		NewBaseTarget      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceiveQuoteAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MtFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NewRState = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.NewBaseTarget = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address trader, uint256 payBaseAmount) view returns(uint256 receiveQuoteAmount, uint256 mtFee, uint8 newRState, uint256 newBaseTarget)
func (_Contract *ContractSession) QuerySellBase(trader common.Address, payBaseAmount *big.Int) (struct {
	ReceiveQuoteAmount *big.Int
	MtFee              *big.Int
	NewRState          uint8
	NewBaseTarget      *big.Int
}, error) {
	return _Contract.Contract.QuerySellBase(&_Contract.CallOpts, trader, payBaseAmount)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address trader, uint256 payBaseAmount) view returns(uint256 receiveQuoteAmount, uint256 mtFee, uint8 newRState, uint256 newBaseTarget)
func (_Contract *ContractCallerSession) QuerySellBase(trader common.Address, payBaseAmount *big.Int) (struct {
	ReceiveQuoteAmount *big.Int
	MtFee              *big.Int
	NewRState          uint8
	NewBaseTarget      *big.Int
}, error) {
	return _Contract.Contract.QuerySellBase(&_Contract.CallOpts, trader, payBaseAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address trader, uint256 payQuoteAmount) view returns(uint256 receiveBaseAmount, uint256 mtFee, uint8 newRState, uint256 newQuoteTarget)
func (_Contract *ContractCaller) QuerySellQuote(opts *bind.CallOpts, trader common.Address, payQuoteAmount *big.Int) (struct {
	ReceiveBaseAmount *big.Int
	MtFee             *big.Int
	NewRState         uint8
	NewQuoteTarget    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "querySellQuote", trader, payQuoteAmount)

	outstruct := new(struct {
		ReceiveBaseAmount *big.Int
		MtFee             *big.Int
		NewRState         uint8
		NewQuoteTarget    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceiveBaseAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MtFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NewRState = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.NewQuoteTarget = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address trader, uint256 payQuoteAmount) view returns(uint256 receiveBaseAmount, uint256 mtFee, uint8 newRState, uint256 newQuoteTarget)
func (_Contract *ContractSession) QuerySellQuote(trader common.Address, payQuoteAmount *big.Int) (struct {
	ReceiveBaseAmount *big.Int
	MtFee             *big.Int
	NewRState         uint8
	NewQuoteTarget    *big.Int
}, error) {
	return _Contract.Contract.QuerySellQuote(&_Contract.CallOpts, trader, payQuoteAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address trader, uint256 payQuoteAmount) view returns(uint256 receiveBaseAmount, uint256 mtFee, uint8 newRState, uint256 newQuoteTarget)
func (_Contract *ContractCallerSession) QuerySellQuote(trader common.Address, payQuoteAmount *big.Int) (struct {
	ReceiveBaseAmount *big.Int
	MtFee             *big.Int
	NewRState         uint8
	NewQuoteTarget    *big.Int
}, error) {
	return _Contract.Contract.QuerySellQuote(&_Contract.CallOpts, trader, payQuoteAmount)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Contract *ContractSession) Version() (string, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Contract *ContractCallerSession) Version() (string, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Contract *ContractTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Contract *ContractSession) ClaimOwnership() (*types.Transaction, error) {
	return _Contract.Contract.ClaimOwnership(&_Contract.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Contract *ContractTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Contract.Contract.ClaimOwnership(&_Contract.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xd0a494e4.
//
// Solidity: function flashLoan(uint256 baseAmount, uint256 quoteAmount, address assetTo, bytes data) returns()
func (_Contract *ContractTransactor) FlashLoan(opts *bind.TransactOpts, baseAmount *big.Int, quoteAmount *big.Int, assetTo common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "flashLoan", baseAmount, quoteAmount, assetTo, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xd0a494e4.
//
// Solidity: function flashLoan(uint256 baseAmount, uint256 quoteAmount, address assetTo, bytes data) returns()
func (_Contract *ContractSession) FlashLoan(baseAmount *big.Int, quoteAmount *big.Int, assetTo common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.FlashLoan(&_Contract.TransactOpts, baseAmount, quoteAmount, assetTo, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xd0a494e4.
//
// Solidity: function flashLoan(uint256 baseAmount, uint256 quoteAmount, address assetTo, bytes data) returns()
func (_Contract *ContractTransactorSession) FlashLoan(baseAmount *big.Int, quoteAmount *big.Int, assetTo common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.FlashLoan(&_Contract.TransactOpts, baseAmount, quoteAmount, assetTo, data)
}

// Init is a paid mutator transaction binding the contract method 0x01a3c30b.
//
// Solidity: function init(address owner, address maintainer, address baseTokenAddress, address quoteTokenAddress, uint256 lpFeeRate, address mtFeeRateModel, uint256 k, uint256 i, bool isOpenTWAP) returns()
func (_Contract *ContractTransactor) Init(opts *bind.TransactOpts, owner common.Address, maintainer common.Address, baseTokenAddress common.Address, quoteTokenAddress common.Address, lpFeeRate *big.Int, mtFeeRateModel common.Address, k *big.Int, i *big.Int, isOpenTWAP bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "init", owner, maintainer, baseTokenAddress, quoteTokenAddress, lpFeeRate, mtFeeRateModel, k, i, isOpenTWAP)
}

// Init is a paid mutator transaction binding the contract method 0x01a3c30b.
//
// Solidity: function init(address owner, address maintainer, address baseTokenAddress, address quoteTokenAddress, uint256 lpFeeRate, address mtFeeRateModel, uint256 k, uint256 i, bool isOpenTWAP) returns()
func (_Contract *ContractSession) Init(owner common.Address, maintainer common.Address, baseTokenAddress common.Address, quoteTokenAddress common.Address, lpFeeRate *big.Int, mtFeeRateModel common.Address, k *big.Int, i *big.Int, isOpenTWAP bool) (*types.Transaction, error) {
	return _Contract.Contract.Init(&_Contract.TransactOpts, owner, maintainer, baseTokenAddress, quoteTokenAddress, lpFeeRate, mtFeeRateModel, k, i, isOpenTWAP)
}

// Init is a paid mutator transaction binding the contract method 0x01a3c30b.
//
// Solidity: function init(address owner, address maintainer, address baseTokenAddress, address quoteTokenAddress, uint256 lpFeeRate, address mtFeeRateModel, uint256 k, uint256 i, bool isOpenTWAP) returns()
func (_Contract *ContractTransactorSession) Init(owner common.Address, maintainer common.Address, baseTokenAddress common.Address, quoteTokenAddress common.Address, lpFeeRate *big.Int, mtFeeRateModel common.Address, k *big.Int, i *big.Int, isOpenTWAP bool) (*types.Transaction, error) {
	return _Contract.Contract.Init(&_Contract.TransactOpts, owner, maintainer, baseTokenAddress, quoteTokenAddress, lpFeeRate, mtFeeRateModel, k, i, isOpenTWAP)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_Contract *ContractTransactor) InitOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initOwner", newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_Contract *ContractSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.InitOwner(&_Contract.TransactOpts, newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_Contract *ContractTransactorSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.InitOwner(&_Contract.TransactOpts, newOwner)
}

// RatioSync is a paid mutator transaction binding the contract method 0xc57a5d03.
//
// Solidity: function ratioSync() returns()
func (_Contract *ContractTransactor) RatioSync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "ratioSync")
}

// RatioSync is a paid mutator transaction binding the contract method 0xc57a5d03.
//
// Solidity: function ratioSync() returns()
func (_Contract *ContractSession) RatioSync() (*types.Transaction, error) {
	return _Contract.Contract.RatioSync(&_Contract.TransactOpts)
}

// RatioSync is a paid mutator transaction binding the contract method 0xc57a5d03.
//
// Solidity: function ratioSync() returns()
func (_Contract *ContractTransactorSession) RatioSync() (*types.Transaction, error) {
	return _Contract.Contract.RatioSync(&_Contract.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0x8ff3928c.
//
// Solidity: function reset(address assetTo, uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 baseOutAmount, uint256 quoteOutAmount, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_Contract *ContractTransactor) Reset(opts *bind.TransactOpts, assetTo common.Address, newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, baseOutAmount *big.Int, quoteOutAmount *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reset", assetTo, newLpFeeRate, newI, newK, baseOutAmount, quoteOutAmount, minBaseReserve, minQuoteReserve)
}

// Reset is a paid mutator transaction binding the contract method 0x8ff3928c.
//
// Solidity: function reset(address assetTo, uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 baseOutAmount, uint256 quoteOutAmount, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_Contract *ContractSession) Reset(assetTo common.Address, newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, baseOutAmount *big.Int, quoteOutAmount *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reset(&_Contract.TransactOpts, assetTo, newLpFeeRate, newI, newK, baseOutAmount, quoteOutAmount, minBaseReserve, minQuoteReserve)
}

// Reset is a paid mutator transaction binding the contract method 0x8ff3928c.
//
// Solidity: function reset(address assetTo, uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 baseOutAmount, uint256 quoteOutAmount, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_Contract *ContractTransactorSession) Reset(assetTo common.Address, newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, baseOutAmount *big.Int, quoteOutAmount *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reset(&_Contract.TransactOpts, assetTo, newLpFeeRate, newI, newK, baseOutAmount, quoteOutAmount, minBaseReserve, minQuoteReserve)
}

// Retrieve is a paid mutator transaction binding the contract method 0x28c4e24c.
//
// Solidity: function retrieve(address to, address token, uint256 amount) returns()
func (_Contract *ContractTransactor) Retrieve(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "retrieve", to, token, amount)
}

// Retrieve is a paid mutator transaction binding the contract method 0x28c4e24c.
//
// Solidity: function retrieve(address to, address token, uint256 amount) returns()
func (_Contract *ContractSession) Retrieve(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Retrieve(&_Contract.TransactOpts, to, token, amount)
}

// Retrieve is a paid mutator transaction binding the contract method 0x28c4e24c.
//
// Solidity: function retrieve(address to, address token, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Retrieve(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Retrieve(&_Contract.TransactOpts, to, token, amount)
}

// SellBase is a paid mutator transaction binding the contract method 0xbd6015b4.
//
// Solidity: function sellBase(address to) returns(uint256 receiveQuoteAmount)
func (_Contract *ContractTransactor) SellBase(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sellBase", to)
}

// SellBase is a paid mutator transaction binding the contract method 0xbd6015b4.
//
// Solidity: function sellBase(address to) returns(uint256 receiveQuoteAmount)
func (_Contract *ContractSession) SellBase(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SellBase(&_Contract.TransactOpts, to)
}

// SellBase is a paid mutator transaction binding the contract method 0xbd6015b4.
//
// Solidity: function sellBase(address to) returns(uint256 receiveQuoteAmount)
func (_Contract *ContractTransactorSession) SellBase(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SellBase(&_Contract.TransactOpts, to)
}

// SellQuote is a paid mutator transaction binding the contract method 0xdd93f59a.
//
// Solidity: function sellQuote(address to) returns(uint256 receiveBaseAmount)
func (_Contract *ContractTransactor) SellQuote(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sellQuote", to)
}

// SellQuote is a paid mutator transaction binding the contract method 0xdd93f59a.
//
// Solidity: function sellQuote(address to) returns(uint256 receiveBaseAmount)
func (_Contract *ContractSession) SellQuote(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SellQuote(&_Contract.TransactOpts, to)
}

// SellQuote is a paid mutator transaction binding the contract method 0xdd93f59a.
//
// Solidity: function sellQuote(address to) returns(uint256 receiveBaseAmount)
func (_Contract *ContractTransactorSession) SellQuote(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SellQuote(&_Contract.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x3b20884a.
//
// Solidity: function tuneParameters(uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_Contract *ContractTransactor) TuneParameters(opts *bind.TransactOpts, newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "tuneParameters", newLpFeeRate, newI, newK, minBaseReserve, minQuoteReserve)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x3b20884a.
//
// Solidity: function tuneParameters(uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_Contract *ContractSession) TuneParameters(newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TuneParameters(&_Contract.TransactOpts, newLpFeeRate, newI, newK, minBaseReserve, minQuoteReserve)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x3b20884a.
//
// Solidity: function tuneParameters(uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_Contract *ContractTransactorSession) TuneParameters(newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TuneParameters(&_Contract.TransactOpts, newLpFeeRate, newI, newK, minBaseReserve, minQuoteReserve)
}

// TunePrice is a paid mutator transaction binding the contract method 0x10d76460.
//
// Solidity: function tunePrice(uint256 newI, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_Contract *ContractTransactor) TunePrice(opts *bind.TransactOpts, newI *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "tunePrice", newI, minBaseReserve, minQuoteReserve)
}

// TunePrice is a paid mutator transaction binding the contract method 0x10d76460.
//
// Solidity: function tunePrice(uint256 newI, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_Contract *ContractSession) TunePrice(newI *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TunePrice(&_Contract.TransactOpts, newI, minBaseReserve, minQuoteReserve)
}

// TunePrice is a paid mutator transaction binding the contract method 0x10d76460.
//
// Solidity: function tunePrice(uint256 newI, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_Contract *ContractTransactorSession) TunePrice(newI *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TunePrice(&_Contract.TransactOpts, newI, minBaseReserve, minQuoteReserve)
}

// ContractDODOFlashLoanIterator is returned from FilterDODOFlashLoan and is used to iterate over the raw logs and unpacked data for DODOFlashLoan events raised by the Contract contract.
type ContractDODOFlashLoanIterator struct {
	Event *ContractDODOFlashLoan // Event containing the contract specifics and raw log

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
func (it *ContractDODOFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDODOFlashLoan)
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
		it.Event = new(ContractDODOFlashLoan)
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
func (it *ContractDODOFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDODOFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDODOFlashLoan represents a DODOFlashLoan event raised by the Contract contract.
type ContractDODOFlashLoan struct {
	Borrower    common.Address
	AssetTo     common.Address
	BaseAmount  *big.Int
	QuoteAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDODOFlashLoan is a free log retrieval operation binding the contract event 0x0b82e93068db15abd9fbb2682c65462ea8a0a10582dce93a5664818e296f54eb.
//
// Solidity: event DODOFlashLoan(address borrower, address assetTo, uint256 baseAmount, uint256 quoteAmount)
func (_Contract *ContractFilterer) FilterDODOFlashLoan(opts *bind.FilterOpts) (*ContractDODOFlashLoanIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DODOFlashLoan")
	if err != nil {
		return nil, err
	}
	return &ContractDODOFlashLoanIterator{contract: _Contract.contract, event: "DODOFlashLoan", logs: logs, sub: sub}, nil
}

// WatchDODOFlashLoan is a free log subscription operation binding the contract event 0x0b82e93068db15abd9fbb2682c65462ea8a0a10582dce93a5664818e296f54eb.
//
// Solidity: event DODOFlashLoan(address borrower, address assetTo, uint256 baseAmount, uint256 quoteAmount)
func (_Contract *ContractFilterer) WatchDODOFlashLoan(opts *bind.WatchOpts, sink chan<- *ContractDODOFlashLoan) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DODOFlashLoan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDODOFlashLoan)
				if err := _Contract.contract.UnpackLog(event, "DODOFlashLoan", log); err != nil {
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

// ParseDODOFlashLoan is a log parse operation binding the contract event 0x0b82e93068db15abd9fbb2682c65462ea8a0a10582dce93a5664818e296f54eb.
//
// Solidity: event DODOFlashLoan(address borrower, address assetTo, uint256 baseAmount, uint256 quoteAmount)
func (_Contract *ContractFilterer) ParseDODOFlashLoan(log types.Log) (*ContractDODOFlashLoan, error) {
	event := new(ContractDODOFlashLoan)
	if err := _Contract.contract.UnpackLog(event, "DODOFlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDODOSwapIterator is returned from FilterDODOSwap and is used to iterate over the raw logs and unpacked data for DODOSwap events raised by the Contract contract.
type ContractDODOSwapIterator struct {
	Event *ContractDODOSwap // Event containing the contract specifics and raw log

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
func (it *ContractDODOSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDODOSwap)
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
		it.Event = new(ContractDODOSwap)
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
func (it *ContractDODOSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDODOSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDODOSwap represents a DODOSwap event raised by the Contract contract.
type ContractDODOSwap struct {
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	Trader     common.Address
	Receiver   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDODOSwap is a free log retrieval operation binding the contract event 0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f.
//
// Solidity: event DODOSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address trader, address receiver)
func (_Contract *ContractFilterer) FilterDODOSwap(opts *bind.FilterOpts) (*ContractDODOSwapIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DODOSwap")
	if err != nil {
		return nil, err
	}
	return &ContractDODOSwapIterator{contract: _Contract.contract, event: "DODOSwap", logs: logs, sub: sub}, nil
}

// WatchDODOSwap is a free log subscription operation binding the contract event 0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f.
//
// Solidity: event DODOSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address trader, address receiver)
func (_Contract *ContractFilterer) WatchDODOSwap(opts *bind.WatchOpts, sink chan<- *ContractDODOSwap) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DODOSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDODOSwap)
				if err := _Contract.contract.UnpackLog(event, "DODOSwap", log); err != nil {
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

// ParseDODOSwap is a log parse operation binding the contract event 0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f.
//
// Solidity: event DODOSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address trader, address receiver)
func (_Contract *ContractFilterer) ParseDODOSwap(log types.Log) (*ContractDODOSwap, error) {
	event := new(ContractDODOSwap)
	if err := _Contract.contract.UnpackLog(event, "DODOSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLpFeeRateChangeIterator is returned from FilterLpFeeRateChange and is used to iterate over the raw logs and unpacked data for LpFeeRateChange events raised by the Contract contract.
type ContractLpFeeRateChangeIterator struct {
	Event *ContractLpFeeRateChange // Event containing the contract specifics and raw log

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
func (it *ContractLpFeeRateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLpFeeRateChange)
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
		it.Event = new(ContractLpFeeRateChange)
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
func (it *ContractLpFeeRateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLpFeeRateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLpFeeRateChange represents a LpFeeRateChange event raised by the Contract contract.
type ContractLpFeeRateChange struct {
	NewLpFeeRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLpFeeRateChange is a free log retrieval operation binding the contract event 0x9950d5a2f2c7264863d40100bf993f0cdbc4711806caba6284d07e80fd500879.
//
// Solidity: event LpFeeRateChange(uint256 newLpFeeRate)
func (_Contract *ContractFilterer) FilterLpFeeRateChange(opts *bind.FilterOpts) (*ContractLpFeeRateChangeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LpFeeRateChange")
	if err != nil {
		return nil, err
	}
	return &ContractLpFeeRateChangeIterator{contract: _Contract.contract, event: "LpFeeRateChange", logs: logs, sub: sub}, nil
}

// WatchLpFeeRateChange is a free log subscription operation binding the contract event 0x9950d5a2f2c7264863d40100bf993f0cdbc4711806caba6284d07e80fd500879.
//
// Solidity: event LpFeeRateChange(uint256 newLpFeeRate)
func (_Contract *ContractFilterer) WatchLpFeeRateChange(opts *bind.WatchOpts, sink chan<- *ContractLpFeeRateChange) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LpFeeRateChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLpFeeRateChange)
				if err := _Contract.contract.UnpackLog(event, "LpFeeRateChange", log); err != nil {
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

// ParseLpFeeRateChange is a log parse operation binding the contract event 0x9950d5a2f2c7264863d40100bf993f0cdbc4711806caba6284d07e80fd500879.
//
// Solidity: event LpFeeRateChange(uint256 newLpFeeRate)
func (_Contract *ContractFilterer) ParseLpFeeRateChange(log types.Log) (*ContractLpFeeRateChange, error) {
	event := new(ContractLpFeeRateChange)
	if err := _Contract.contract.UnpackLog(event, "LpFeeRateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferPreparedIterator is returned from FilterOwnershipTransferPrepared and is used to iterate over the raw logs and unpacked data for OwnershipTransferPrepared events raised by the Contract contract.
type ContractOwnershipTransferPreparedIterator struct {
	Event *ContractOwnershipTransferPrepared // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferPrepared)
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
		it.Event = new(ContractOwnershipTransferPrepared)
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
func (it *ContractOwnershipTransferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferPrepared represents a OwnershipTransferPrepared event raised by the Contract contract.
type ContractOwnershipTransferPrepared struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferPrepared is a free log retrieval operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferPrepared(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferPreparedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferPreparedIterator{contract: _Contract.contract, event: "OwnershipTransferPrepared", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferPrepared is a free log subscription operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferPrepared(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferPrepared, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferPrepared)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
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

// ParseOwnershipTransferPrepared is a log parse operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferPrepared(log types.Log) (*ContractOwnershipTransferPrepared, error) {
	event := new(ContractOwnershipTransferPrepared)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRChangeIterator is returned from FilterRChange and is used to iterate over the raw logs and unpacked data for RChange events raised by the Contract contract.
type ContractRChangeIterator struct {
	Event *ContractRChange // Event containing the contract specifics and raw log

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
func (it *ContractRChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRChange)
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
		it.Event = new(ContractRChange)
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
func (it *ContractRChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRChange represents a RChange event raised by the Contract contract.
type ContractRChange struct {
	NewRState uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRChange is a free log retrieval operation binding the contract event 0xdf176ad18be4f9f32efaa32f06e9d1175476504739a745f1399a6d3fa4b75917.
//
// Solidity: event RChange(uint8 newRState)
func (_Contract *ContractFilterer) FilterRChange(opts *bind.FilterOpts) (*ContractRChangeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RChange")
	if err != nil {
		return nil, err
	}
	return &ContractRChangeIterator{contract: _Contract.contract, event: "RChange", logs: logs, sub: sub}, nil
}

// WatchRChange is a free log subscription operation binding the contract event 0xdf176ad18be4f9f32efaa32f06e9d1175476504739a745f1399a6d3fa4b75917.
//
// Solidity: event RChange(uint8 newRState)
func (_Contract *ContractFilterer) WatchRChange(opts *bind.WatchOpts, sink chan<- *ContractRChange) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRChange)
				if err := _Contract.contract.UnpackLog(event, "RChange", log); err != nil {
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

// ParseRChange is a log parse operation binding the contract event 0xdf176ad18be4f9f32efaa32f06e9d1175476504739a745f1399a6d3fa4b75917.
//
// Solidity: event RChange(uint8 newRState)
func (_Contract *ContractFilterer) ParseRChange(log types.Log) (*ContractRChange, error) {
	event := new(ContractRChange)
	if err := _Contract.contract.UnpackLog(event, "RChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
