// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bentobox_pool

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"wethToken_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cloneAddress\",\"type\":\"address\"}],\"name\":\"LogDeploy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"LogDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"LogFlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"}],\"name\":\"LogRegisterProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"LogSetMasterContractApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogStrategyDivest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogStrategyInvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogStrategyLoss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogStrategyProfit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"LogStrategyQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"LogStrategySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetPercentage\",\"type\":\"uint256\"}],\"name\":\"LogStrategyTargetPercentage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"LogTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"LogWhiteListMasterContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"LogWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"calls\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"revertOnFail\",\"type\":\"bool\"}],\"name\":\"batch\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"},{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBatchFlashBorrower\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"batchFlashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"useCreate2\",\"type\":\"bool\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cloneAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shareOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFlashBorrower\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"balance\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeAmount\",\"type\":\"uint256\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"masterContractApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"masterContractOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingStrategy\",\"outputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permitToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"masterContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"setMasterContractApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIStrategy\",\"name\":\"newStrategy\",\"type\":\"address\"}],\"name\":\"setStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetPercentage_\",\"type\":\"uint64\"}],\"name\":\"setStrategyTargetPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"strategy\",\"outputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"strategyData\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"strategyStartDate\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"targetPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"balance\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"roundUp\",\"type\":\"bool\"}],\"name\":\"toAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"roundUp\",\"type\":\"bool\"}],\"name\":\"toShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totals\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"elastic\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"base\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares\",\"type\":\"uint256[]\"}],\"name\":\"transferMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"direct\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"renounce\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"whitelistMasterContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedMasterContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shareOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contract.Contract.DOMAINSEPARATOR(&_Contract.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contract.Contract.DOMAINSEPARATOR(&_Contract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address , address ) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address , address ) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address , address ) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, arg0, arg1)
}

// MasterContractApproved is a free data retrieval call binding the contract method 0x91e0eab5.
//
// Solidity: function masterContractApproved(address , address ) view returns(bool)
func (_Contract *ContractCaller) MasterContractApproved(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "masterContractApproved", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MasterContractApproved is a free data retrieval call binding the contract method 0x91e0eab5.
//
// Solidity: function masterContractApproved(address , address ) view returns(bool)
func (_Contract *ContractSession) MasterContractApproved(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.MasterContractApproved(&_Contract.CallOpts, arg0, arg1)
}

// MasterContractApproved is a free data retrieval call binding the contract method 0x91e0eab5.
//
// Solidity: function masterContractApproved(address , address ) view returns(bool)
func (_Contract *ContractCallerSession) MasterContractApproved(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.MasterContractApproved(&_Contract.CallOpts, arg0, arg1)
}

// MasterContractOf is a free data retrieval call binding the contract method 0xbafe4f14.
//
// Solidity: function masterContractOf(address ) view returns(address)
func (_Contract *ContractCaller) MasterContractOf(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "masterContractOf", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterContractOf is a free data retrieval call binding the contract method 0xbafe4f14.
//
// Solidity: function masterContractOf(address ) view returns(address)
func (_Contract *ContractSession) MasterContractOf(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.MasterContractOf(&_Contract.CallOpts, arg0)
}

// MasterContractOf is a free data retrieval call binding the contract method 0xbafe4f14.
//
// Solidity: function masterContractOf(address ) view returns(address)
func (_Contract *ContractCallerSession) MasterContractOf(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.MasterContractOf(&_Contract.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Contract *ContractCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Contract *ContractSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonces(&_Contract.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonces(&_Contract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Contract *ContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Contract *ContractSession) PendingOwner() (common.Address, error) {
	return _Contract.Contract.PendingOwner(&_Contract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Contract *ContractCallerSession) PendingOwner() (common.Address, error) {
	return _Contract.Contract.PendingOwner(&_Contract.CallOpts)
}

// PendingStrategy is a free data retrieval call binding the contract method 0x5108a558.
//
// Solidity: function pendingStrategy(address ) view returns(address)
func (_Contract *ContractCaller) PendingStrategy(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pendingStrategy", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingStrategy is a free data retrieval call binding the contract method 0x5108a558.
//
// Solidity: function pendingStrategy(address ) view returns(address)
func (_Contract *ContractSession) PendingStrategy(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.PendingStrategy(&_Contract.CallOpts, arg0)
}

// PendingStrategy is a free data retrieval call binding the contract method 0x5108a558.
//
// Solidity: function pendingStrategy(address ) view returns(address)
func (_Contract *ContractCallerSession) PendingStrategy(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.PendingStrategy(&_Contract.CallOpts, arg0)
}

// Strategy is a free data retrieval call binding the contract method 0x228bfd9f.
//
// Solidity: function strategy(address ) view returns(address)
func (_Contract *ContractCaller) Strategy(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "strategy", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategy is a free data retrieval call binding the contract method 0x228bfd9f.
//
// Solidity: function strategy(address ) view returns(address)
func (_Contract *ContractSession) Strategy(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.Strategy(&_Contract.CallOpts, arg0)
}

// Strategy is a free data retrieval call binding the contract method 0x228bfd9f.
//
// Solidity: function strategy(address ) view returns(address)
func (_Contract *ContractCallerSession) Strategy(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.Strategy(&_Contract.CallOpts, arg0)
}

// StrategyData is a free data retrieval call binding the contract method 0xdf23b45b.
//
// Solidity: function strategyData(address ) view returns(uint64 strategyStartDate, uint64 targetPercentage, uint128 balance)
func (_Contract *ContractCaller) StrategyData(opts *bind.CallOpts, arg0 common.Address) (struct {
	StrategyStartDate uint64
	TargetPercentage  uint64
	Balance           *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "strategyData", arg0)

	outstruct := new(struct {
		StrategyStartDate uint64
		TargetPercentage  uint64
		Balance           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StrategyStartDate = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.TargetPercentage = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Balance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StrategyData is a free data retrieval call binding the contract method 0xdf23b45b.
//
// Solidity: function strategyData(address ) view returns(uint64 strategyStartDate, uint64 targetPercentage, uint128 balance)
func (_Contract *ContractSession) StrategyData(arg0 common.Address) (struct {
	StrategyStartDate uint64
	TargetPercentage  uint64
	Balance           *big.Int
}, error) {
	return _Contract.Contract.StrategyData(&_Contract.CallOpts, arg0)
}

// StrategyData is a free data retrieval call binding the contract method 0xdf23b45b.
//
// Solidity: function strategyData(address ) view returns(uint64 strategyStartDate, uint64 targetPercentage, uint128 balance)
func (_Contract *ContractCallerSession) StrategyData(arg0 common.Address) (struct {
	StrategyStartDate uint64
	TargetPercentage  uint64
	Balance           *big.Int
}, error) {
	return _Contract.Contract.StrategyData(&_Contract.CallOpts, arg0)
}

// ToAmount is a free data retrieval call binding the contract method 0x56623118.
//
// Solidity: function toAmount(address token, uint256 share, bool roundUp) view returns(uint256 amount)
func (_Contract *ContractCaller) ToAmount(opts *bind.CallOpts, token common.Address, share *big.Int, roundUp bool) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "toAmount", token, share, roundUp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToAmount is a free data retrieval call binding the contract method 0x56623118.
//
// Solidity: function toAmount(address token, uint256 share, bool roundUp) view returns(uint256 amount)
func (_Contract *ContractSession) ToAmount(token common.Address, share *big.Int, roundUp bool) (*big.Int, error) {
	return _Contract.Contract.ToAmount(&_Contract.CallOpts, token, share, roundUp)
}

// ToAmount is a free data retrieval call binding the contract method 0x56623118.
//
// Solidity: function toAmount(address token, uint256 share, bool roundUp) view returns(uint256 amount)
func (_Contract *ContractCallerSession) ToAmount(token common.Address, share *big.Int, roundUp bool) (*big.Int, error) {
	return _Contract.Contract.ToAmount(&_Contract.CallOpts, token, share, roundUp)
}

// ToShare is a free data retrieval call binding the contract method 0xda5139ca.
//
// Solidity: function toShare(address token, uint256 amount, bool roundUp) view returns(uint256 share)
func (_Contract *ContractCaller) ToShare(opts *bind.CallOpts, token common.Address, amount *big.Int, roundUp bool) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "toShare", token, amount, roundUp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToShare is a free data retrieval call binding the contract method 0xda5139ca.
//
// Solidity: function toShare(address token, uint256 amount, bool roundUp) view returns(uint256 share)
func (_Contract *ContractSession) ToShare(token common.Address, amount *big.Int, roundUp bool) (*big.Int, error) {
	return _Contract.Contract.ToShare(&_Contract.CallOpts, token, amount, roundUp)
}

// ToShare is a free data retrieval call binding the contract method 0xda5139ca.
//
// Solidity: function toShare(address token, uint256 amount, bool roundUp) view returns(uint256 share)
func (_Contract *ContractCallerSession) ToShare(token common.Address, amount *big.Int, roundUp bool) (*big.Int, error) {
	return _Contract.Contract.ToShare(&_Contract.CallOpts, token, amount, roundUp)
}

// Totals is a free data retrieval call binding the contract method 0x4ffe34db.
//
// Solidity: function totals(address ) view returns(uint128 elastic, uint128 base)
func (_Contract *ContractCaller) Totals(opts *bind.CallOpts, arg0 common.Address) (struct {
	Elastic *big.Int
	Base    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totals", arg0)

	outstruct := new(struct {
		Elastic *big.Int
		Base    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Elastic = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Base = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Totals is a free data retrieval call binding the contract method 0x4ffe34db.
//
// Solidity: function totals(address ) view returns(uint128 elastic, uint128 base)
func (_Contract *ContractSession) Totals(arg0 common.Address) (struct {
	Elastic *big.Int
	Base    *big.Int
}, error) {
	return _Contract.Contract.Totals(&_Contract.CallOpts, arg0)
}

// Totals is a free data retrieval call binding the contract method 0x4ffe34db.
//
// Solidity: function totals(address ) view returns(uint128 elastic, uint128 base)
func (_Contract *ContractCallerSession) Totals(arg0 common.Address) (struct {
	Elastic *big.Int
	Base    *big.Int
}, error) {
	return _Contract.Contract.Totals(&_Contract.CallOpts, arg0)
}

// WhitelistedMasterContracts is a free data retrieval call binding the contract method 0x12a90c8a.
//
// Solidity: function whitelistedMasterContracts(address ) view returns(bool)
func (_Contract *ContractCaller) WhitelistedMasterContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "whitelistedMasterContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedMasterContracts is a free data retrieval call binding the contract method 0x12a90c8a.
//
// Solidity: function whitelistedMasterContracts(address ) view returns(bool)
func (_Contract *ContractSession) WhitelistedMasterContracts(arg0 common.Address) (bool, error) {
	return _Contract.Contract.WhitelistedMasterContracts(&_Contract.CallOpts, arg0)
}

// WhitelistedMasterContracts is a free data retrieval call binding the contract method 0x12a90c8a.
//
// Solidity: function whitelistedMasterContracts(address ) view returns(bool)
func (_Contract *ContractCallerSession) WhitelistedMasterContracts(arg0 common.Address) (bool, error) {
	return _Contract.Contract.WhitelistedMasterContracts(&_Contract.CallOpts, arg0)
}

// Batch is a paid mutator transaction binding the contract method 0xd2423b51.
//
// Solidity: function batch(bytes[] calls, bool revertOnFail) payable returns(bool[] successes, bytes[] results)
func (_Contract *ContractTransactor) Batch(opts *bind.TransactOpts, calls [][]byte, revertOnFail bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "batch", calls, revertOnFail)
}

// Batch is a paid mutator transaction binding the contract method 0xd2423b51.
//
// Solidity: function batch(bytes[] calls, bool revertOnFail) payable returns(bool[] successes, bytes[] results)
func (_Contract *ContractSession) Batch(calls [][]byte, revertOnFail bool) (*types.Transaction, error) {
	return _Contract.Contract.Batch(&_Contract.TransactOpts, calls, revertOnFail)
}

// Batch is a paid mutator transaction binding the contract method 0xd2423b51.
//
// Solidity: function batch(bytes[] calls, bool revertOnFail) payable returns(bool[] successes, bytes[] results)
func (_Contract *ContractTransactorSession) Batch(calls [][]byte, revertOnFail bool) (*types.Transaction, error) {
	return _Contract.Contract.Batch(&_Contract.TransactOpts, calls, revertOnFail)
}

// BatchFlashLoan is a paid mutator transaction binding the contract method 0xf483b3da.
//
// Solidity: function batchFlashLoan(address borrower, address[] receivers, address[] tokens, uint256[] amounts, bytes data) returns()
func (_Contract *ContractTransactor) BatchFlashLoan(opts *bind.TransactOpts, borrower common.Address, receivers []common.Address, tokens []common.Address, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "batchFlashLoan", borrower, receivers, tokens, amounts, data)
}

// BatchFlashLoan is a paid mutator transaction binding the contract method 0xf483b3da.
//
// Solidity: function batchFlashLoan(address borrower, address[] receivers, address[] tokens, uint256[] amounts, bytes data) returns()
func (_Contract *ContractSession) BatchFlashLoan(borrower common.Address, receivers []common.Address, tokens []common.Address, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.BatchFlashLoan(&_Contract.TransactOpts, borrower, receivers, tokens, amounts, data)
}

// BatchFlashLoan is a paid mutator transaction binding the contract method 0xf483b3da.
//
// Solidity: function batchFlashLoan(address borrower, address[] receivers, address[] tokens, uint256[] amounts, bytes data) returns()
func (_Contract *ContractTransactorSession) BatchFlashLoan(borrower common.Address, receivers []common.Address, tokens []common.Address, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.BatchFlashLoan(&_Contract.TransactOpts, borrower, receivers, tokens, amounts, data)
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

// Deploy is a paid mutator transaction binding the contract method 0x1f54245b.
//
// Solidity: function deploy(address masterContract, bytes data, bool useCreate2) payable returns(address cloneAddress)
func (_Contract *ContractTransactor) Deploy(opts *bind.TransactOpts, masterContract common.Address, data []byte, useCreate2 bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deploy", masterContract, data, useCreate2)
}

// Deploy is a paid mutator transaction binding the contract method 0x1f54245b.
//
// Solidity: function deploy(address masterContract, bytes data, bool useCreate2) payable returns(address cloneAddress)
func (_Contract *ContractSession) Deploy(masterContract common.Address, data []byte, useCreate2 bool) (*types.Transaction, error) {
	return _Contract.Contract.Deploy(&_Contract.TransactOpts, masterContract, data, useCreate2)
}

// Deploy is a paid mutator transaction binding the contract method 0x1f54245b.
//
// Solidity: function deploy(address masterContract, bytes data, bool useCreate2) payable returns(address cloneAddress)
func (_Contract *ContractTransactorSession) Deploy(masterContract common.Address, data []byte, useCreate2 bool) (*types.Transaction, error) {
	return _Contract.Contract.Deploy(&_Contract.TransactOpts, masterContract, data, useCreate2)
}

// Deposit is a paid mutator transaction binding the contract method 0x02b9446c.
//
// Solidity: function deposit(address token_, address from, address to, uint256 amount, uint256 share) payable returns(uint256 amountOut, uint256 shareOut)
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, token_ common.Address, from common.Address, to common.Address, amount *big.Int, share *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", token_, from, to, amount, share)
}

// Deposit is a paid mutator transaction binding the contract method 0x02b9446c.
//
// Solidity: function deposit(address token_, address from, address to, uint256 amount, uint256 share) payable returns(uint256 amountOut, uint256 shareOut)
func (_Contract *ContractSession) Deposit(token_ common.Address, from common.Address, to common.Address, amount *big.Int, share *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, token_, from, to, amount, share)
}

// Deposit is a paid mutator transaction binding the contract method 0x02b9446c.
//
// Solidity: function deposit(address token_, address from, address to, uint256 amount, uint256 share) payable returns(uint256 amountOut, uint256 shareOut)
func (_Contract *ContractTransactorSession) Deposit(token_ common.Address, from common.Address, to common.Address, amount *big.Int, share *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, token_, from, to, amount, share)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xf1676d37.
//
// Solidity: function flashLoan(address borrower, address receiver, address token, uint256 amount, bytes data) returns()
func (_Contract *ContractTransactor) FlashLoan(opts *bind.TransactOpts, borrower common.Address, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "flashLoan", borrower, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xf1676d37.
//
// Solidity: function flashLoan(address borrower, address receiver, address token, uint256 amount, bytes data) returns()
func (_Contract *ContractSession) FlashLoan(borrower common.Address, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.FlashLoan(&_Contract.TransactOpts, borrower, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xf1676d37.
//
// Solidity: function flashLoan(address borrower, address receiver, address token, uint256 amount, bytes data) returns()
func (_Contract *ContractTransactorSession) FlashLoan(borrower common.Address, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.FlashLoan(&_Contract.TransactOpts, borrower, receiver, token, amount, data)
}

// Harvest is a paid mutator transaction binding the contract method 0x66c6bb0b.
//
// Solidity: function harvest(address token, bool balance, uint256 maxChangeAmount) returns()
func (_Contract *ContractTransactor) Harvest(opts *bind.TransactOpts, token common.Address, balance bool, maxChangeAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "harvest", token, balance, maxChangeAmount)
}

// Harvest is a paid mutator transaction binding the contract method 0x66c6bb0b.
//
// Solidity: function harvest(address token, bool balance, uint256 maxChangeAmount) returns()
func (_Contract *ContractSession) Harvest(token common.Address, balance bool, maxChangeAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Harvest(&_Contract.TransactOpts, token, balance, maxChangeAmount)
}

// Harvest is a paid mutator transaction binding the contract method 0x66c6bb0b.
//
// Solidity: function harvest(address token, bool balance, uint256 maxChangeAmount) returns()
func (_Contract *ContractTransactorSession) Harvest(token common.Address, balance bool, maxChangeAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Harvest(&_Contract.TransactOpts, token, balance, maxChangeAmount)
}

// PermitToken is a paid mutator transaction binding the contract method 0x7c516e94.
//
// Solidity: function permitToken(address token, address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactor) PermitToken(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "permitToken", token, from, to, amount, deadline, v, r, s)
}

// PermitToken is a paid mutator transaction binding the contract method 0x7c516e94.
//
// Solidity: function permitToken(address token, address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractSession) PermitToken(token common.Address, from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.PermitToken(&_Contract.TransactOpts, token, from, to, amount, deadline, v, r, s)
}

// PermitToken is a paid mutator transaction binding the contract method 0x7c516e94.
//
// Solidity: function permitToken(address token, address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactorSession) PermitToken(token common.Address, from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.PermitToken(&_Contract.TransactOpts, token, from, to, amount, deadline, v, r, s)
}

// RegisterProtocol is a paid mutator transaction binding the contract method 0xaee4d1b2.
//
// Solidity: function registerProtocol() returns()
func (_Contract *ContractTransactor) RegisterProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerProtocol")
}

// RegisterProtocol is a paid mutator transaction binding the contract method 0xaee4d1b2.
//
// Solidity: function registerProtocol() returns()
func (_Contract *ContractSession) RegisterProtocol() (*types.Transaction, error) {
	return _Contract.Contract.RegisterProtocol(&_Contract.TransactOpts)
}

// RegisterProtocol is a paid mutator transaction binding the contract method 0xaee4d1b2.
//
// Solidity: function registerProtocol() returns()
func (_Contract *ContractTransactorSession) RegisterProtocol() (*types.Transaction, error) {
	return _Contract.Contract.RegisterProtocol(&_Contract.TransactOpts)
}

// SetMasterContractApproval is a paid mutator transaction binding the contract method 0xc0a47c93.
//
// Solidity: function setMasterContractApproval(address user, address masterContract, bool approved, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactor) SetMasterContractApproval(opts *bind.TransactOpts, user common.Address, masterContract common.Address, approved bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMasterContractApproval", user, masterContract, approved, v, r, s)
}

// SetMasterContractApproval is a paid mutator transaction binding the contract method 0xc0a47c93.
//
// Solidity: function setMasterContractApproval(address user, address masterContract, bool approved, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractSession) SetMasterContractApproval(user common.Address, masterContract common.Address, approved bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetMasterContractApproval(&_Contract.TransactOpts, user, masterContract, approved, v, r, s)
}

// SetMasterContractApproval is a paid mutator transaction binding the contract method 0xc0a47c93.
//
// Solidity: function setMasterContractApproval(address user, address masterContract, bool approved, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactorSession) SetMasterContractApproval(user common.Address, masterContract common.Address, approved bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetMasterContractApproval(&_Contract.TransactOpts, user, masterContract, approved, v, r, s)
}

// SetStrategy is a paid mutator transaction binding the contract method 0x72cb5d97.
//
// Solidity: function setStrategy(address token, address newStrategy) returns()
func (_Contract *ContractTransactor) SetStrategy(opts *bind.TransactOpts, token common.Address, newStrategy common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setStrategy", token, newStrategy)
}

// SetStrategy is a paid mutator transaction binding the contract method 0x72cb5d97.
//
// Solidity: function setStrategy(address token, address newStrategy) returns()
func (_Contract *ContractSession) SetStrategy(token common.Address, newStrategy common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetStrategy(&_Contract.TransactOpts, token, newStrategy)
}

// SetStrategy is a paid mutator transaction binding the contract method 0x72cb5d97.
//
// Solidity: function setStrategy(address token, address newStrategy) returns()
func (_Contract *ContractTransactorSession) SetStrategy(token common.Address, newStrategy common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetStrategy(&_Contract.TransactOpts, token, newStrategy)
}

// SetStrategyTargetPercentage is a paid mutator transaction binding the contract method 0x3e2a9d4e.
//
// Solidity: function setStrategyTargetPercentage(address token, uint64 targetPercentage_) returns()
func (_Contract *ContractTransactor) SetStrategyTargetPercentage(opts *bind.TransactOpts, token common.Address, targetPercentage_ uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setStrategyTargetPercentage", token, targetPercentage_)
}

// SetStrategyTargetPercentage is a paid mutator transaction binding the contract method 0x3e2a9d4e.
//
// Solidity: function setStrategyTargetPercentage(address token, uint64 targetPercentage_) returns()
func (_Contract *ContractSession) SetStrategyTargetPercentage(token common.Address, targetPercentage_ uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetStrategyTargetPercentage(&_Contract.TransactOpts, token, targetPercentage_)
}

// SetStrategyTargetPercentage is a paid mutator transaction binding the contract method 0x3e2a9d4e.
//
// Solidity: function setStrategyTargetPercentage(address token, uint64 targetPercentage_) returns()
func (_Contract *ContractTransactorSession) SetStrategyTargetPercentage(token common.Address, targetPercentage_ uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetStrategyTargetPercentage(&_Contract.TransactOpts, token, targetPercentage_)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address token, address from, address to, uint256 share) returns()
func (_Contract *ContractTransactor) Transfer(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, share *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transfer", token, from, to, share)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address token, address from, address to, uint256 share) returns()
func (_Contract *ContractSession) Transfer(token common.Address, from common.Address, to common.Address, share *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, token, from, to, share)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address token, address from, address to, uint256 share) returns()
func (_Contract *ContractTransactorSession) Transfer(token common.Address, from common.Address, to common.Address, share *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, token, from, to, share)
}

// TransferMultiple is a paid mutator transaction binding the contract method 0x0fca8843.
//
// Solidity: function transferMultiple(address token, address from, address[] tos, uint256[] shares) returns()
func (_Contract *ContractTransactor) TransferMultiple(opts *bind.TransactOpts, token common.Address, from common.Address, tos []common.Address, shares []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferMultiple", token, from, tos, shares)
}

// TransferMultiple is a paid mutator transaction binding the contract method 0x0fca8843.
//
// Solidity: function transferMultiple(address token, address from, address[] tos, uint256[] shares) returns()
func (_Contract *ContractSession) TransferMultiple(token common.Address, from common.Address, tos []common.Address, shares []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferMultiple(&_Contract.TransactOpts, token, from, tos, shares)
}

// TransferMultiple is a paid mutator transaction binding the contract method 0x0fca8843.
//
// Solidity: function transferMultiple(address token, address from, address[] tos, uint256[] shares) returns()
func (_Contract *ContractTransactorSession) TransferMultiple(token common.Address, from common.Address, tos []common.Address, shares []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferMultiple(&_Contract.TransactOpts, token, from, tos, shares)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner, direct, renounce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner, direct, renounce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner, direct, renounce)
}

// WhitelistMasterContract is a paid mutator transaction binding the contract method 0x733a9d7c.
//
// Solidity: function whitelistMasterContract(address masterContract, bool approved) returns()
func (_Contract *ContractTransactor) WhitelistMasterContract(opts *bind.TransactOpts, masterContract common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "whitelistMasterContract", masterContract, approved)
}

// WhitelistMasterContract is a paid mutator transaction binding the contract method 0x733a9d7c.
//
// Solidity: function whitelistMasterContract(address masterContract, bool approved) returns()
func (_Contract *ContractSession) WhitelistMasterContract(masterContract common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.WhitelistMasterContract(&_Contract.TransactOpts, masterContract, approved)
}

// WhitelistMasterContract is a paid mutator transaction binding the contract method 0x733a9d7c.
//
// Solidity: function whitelistMasterContract(address masterContract, bool approved) returns()
func (_Contract *ContractTransactorSession) WhitelistMasterContract(masterContract common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.WhitelistMasterContract(&_Contract.TransactOpts, masterContract, approved)
}

// Withdraw is a paid mutator transaction binding the contract method 0x97da6d30.
//
// Solidity: function withdraw(address token_, address from, address to, uint256 amount, uint256 share) returns(uint256 amountOut, uint256 shareOut)
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, token_ common.Address, from common.Address, to common.Address, amount *big.Int, share *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", token_, from, to, amount, share)
}

// Withdraw is a paid mutator transaction binding the contract method 0x97da6d30.
//
// Solidity: function withdraw(address token_, address from, address to, uint256 amount, uint256 share) returns(uint256 amountOut, uint256 shareOut)
func (_Contract *ContractSession) Withdraw(token_ common.Address, from common.Address, to common.Address, amount *big.Int, share *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, token_, from, to, amount, share)
}

// Withdraw is a paid mutator transaction binding the contract method 0x97da6d30.
//
// Solidity: function withdraw(address token_, address from, address to, uint256 amount, uint256 share) returns(uint256 amountOut, uint256 shareOut)
func (_Contract *ContractTransactorSession) Withdraw(token_ common.Address, from common.Address, to common.Address, amount *big.Int, share *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, token_, from, to, amount, share)
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

// ContractLogDeployIterator is returned from FilterLogDeploy and is used to iterate over the raw logs and unpacked data for LogDeploy events raised by the Contract contract.
type ContractLogDeployIterator struct {
	Event *ContractLogDeploy // Event containing the contract specifics and raw log

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
func (it *ContractLogDeployIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogDeploy)
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
		it.Event = new(ContractLogDeploy)
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
func (it *ContractLogDeployIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogDeployIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogDeploy represents a LogDeploy event raised by the Contract contract.
type ContractLogDeploy struct {
	MasterContract common.Address
	Data           []byte
	CloneAddress   common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogDeploy is a free log retrieval operation binding the contract event 0xd62166f3c2149208e51788b1401cc356bf5da1fc6c7886a32e18570f57d88b3b.
//
// Solidity: event LogDeploy(address indexed masterContract, bytes data, address indexed cloneAddress)
func (_Contract *ContractFilterer) FilterLogDeploy(opts *bind.FilterOpts, masterContract []common.Address, cloneAddress []common.Address) (*ContractLogDeployIterator, error) {

	var masterContractRule []interface{}
	for _, masterContractItem := range masterContract {
		masterContractRule = append(masterContractRule, masterContractItem)
	}

	var cloneAddressRule []interface{}
	for _, cloneAddressItem := range cloneAddress {
		cloneAddressRule = append(cloneAddressRule, cloneAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogDeploy", masterContractRule, cloneAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogDeployIterator{contract: _Contract.contract, event: "LogDeploy", logs: logs, sub: sub}, nil
}

// WatchLogDeploy is a free log subscription operation binding the contract event 0xd62166f3c2149208e51788b1401cc356bf5da1fc6c7886a32e18570f57d88b3b.
//
// Solidity: event LogDeploy(address indexed masterContract, bytes data, address indexed cloneAddress)
func (_Contract *ContractFilterer) WatchLogDeploy(opts *bind.WatchOpts, sink chan<- *ContractLogDeploy, masterContract []common.Address, cloneAddress []common.Address) (event.Subscription, error) {

	var masterContractRule []interface{}
	for _, masterContractItem := range masterContract {
		masterContractRule = append(masterContractRule, masterContractItem)
	}

	var cloneAddressRule []interface{}
	for _, cloneAddressItem := range cloneAddress {
		cloneAddressRule = append(cloneAddressRule, cloneAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogDeploy", masterContractRule, cloneAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogDeploy)
				if err := _Contract.contract.UnpackLog(event, "LogDeploy", log); err != nil {
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

// ParseLogDeploy is a log parse operation binding the contract event 0xd62166f3c2149208e51788b1401cc356bf5da1fc6c7886a32e18570f57d88b3b.
//
// Solidity: event LogDeploy(address indexed masterContract, bytes data, address indexed cloneAddress)
func (_Contract *ContractFilterer) ParseLogDeploy(log types.Log) (*ContractLogDeploy, error) {
	event := new(ContractLogDeploy)
	if err := _Contract.contract.UnpackLog(event, "LogDeploy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogDepositIterator is returned from FilterLogDeposit and is used to iterate over the raw logs and unpacked data for LogDeposit events raised by the Contract contract.
type ContractLogDepositIterator struct {
	Event *ContractLogDeposit // Event containing the contract specifics and raw log

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
func (it *ContractLogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogDeposit)
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
		it.Event = new(ContractLogDeposit)
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
func (it *ContractLogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogDeposit represents a LogDeposit event raised by the Contract contract.
type ContractLogDeposit struct {
	Token  common.Address
	From   common.Address
	To     common.Address
	Amount *big.Int
	Share  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogDeposit is a free log retrieval operation binding the contract event 0xb2346165e782564f17f5b7e555c21f4fd96fbc93458572bf0113ea35a958fc55.
//
// Solidity: event LogDeposit(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 share)
func (_Contract *ContractFilterer) FilterLogDeposit(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*ContractLogDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogDeposit", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogDepositIterator{contract: _Contract.contract, event: "LogDeposit", logs: logs, sub: sub}, nil
}

// WatchLogDeposit is a free log subscription operation binding the contract event 0xb2346165e782564f17f5b7e555c21f4fd96fbc93458572bf0113ea35a958fc55.
//
// Solidity: event LogDeposit(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 share)
func (_Contract *ContractFilterer) WatchLogDeposit(opts *bind.WatchOpts, sink chan<- *ContractLogDeposit, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogDeposit", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogDeposit)
				if err := _Contract.contract.UnpackLog(event, "LogDeposit", log); err != nil {
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

// ParseLogDeposit is a log parse operation binding the contract event 0xb2346165e782564f17f5b7e555c21f4fd96fbc93458572bf0113ea35a958fc55.
//
// Solidity: event LogDeposit(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 share)
func (_Contract *ContractFilterer) ParseLogDeposit(log types.Log) (*ContractLogDeposit, error) {
	event := new(ContractLogDeposit)
	if err := _Contract.contract.UnpackLog(event, "LogDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogFlashLoanIterator is returned from FilterLogFlashLoan and is used to iterate over the raw logs and unpacked data for LogFlashLoan events raised by the Contract contract.
type ContractLogFlashLoanIterator struct {
	Event *ContractLogFlashLoan // Event containing the contract specifics and raw log

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
func (it *ContractLogFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogFlashLoan)
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
		it.Event = new(ContractLogFlashLoan)
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
func (it *ContractLogFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogFlashLoan represents a LogFlashLoan event raised by the Contract contract.
type ContractLogFlashLoan struct {
	Borrower  common.Address
	Token     common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogFlashLoan is a free log retrieval operation binding the contract event 0x3be9b85936d5d30a1655ea116a17ee3d827b2cd428cc026ce5bf2ac46a223204.
//
// Solidity: event LogFlashLoan(address indexed borrower, address indexed token, uint256 amount, uint256 feeAmount, address indexed receiver)
func (_Contract *ContractFilterer) FilterLogFlashLoan(opts *bind.FilterOpts, borrower []common.Address, token []common.Address, receiver []common.Address) (*ContractLogFlashLoanIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogFlashLoan", borrowerRule, tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogFlashLoanIterator{contract: _Contract.contract, event: "LogFlashLoan", logs: logs, sub: sub}, nil
}

// WatchLogFlashLoan is a free log subscription operation binding the contract event 0x3be9b85936d5d30a1655ea116a17ee3d827b2cd428cc026ce5bf2ac46a223204.
//
// Solidity: event LogFlashLoan(address indexed borrower, address indexed token, uint256 amount, uint256 feeAmount, address indexed receiver)
func (_Contract *ContractFilterer) WatchLogFlashLoan(opts *bind.WatchOpts, sink chan<- *ContractLogFlashLoan, borrower []common.Address, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogFlashLoan", borrowerRule, tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogFlashLoan)
				if err := _Contract.contract.UnpackLog(event, "LogFlashLoan", log); err != nil {
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

// ParseLogFlashLoan is a log parse operation binding the contract event 0x3be9b85936d5d30a1655ea116a17ee3d827b2cd428cc026ce5bf2ac46a223204.
//
// Solidity: event LogFlashLoan(address indexed borrower, address indexed token, uint256 amount, uint256 feeAmount, address indexed receiver)
func (_Contract *ContractFilterer) ParseLogFlashLoan(log types.Log) (*ContractLogFlashLoan, error) {
	event := new(ContractLogFlashLoan)
	if err := _Contract.contract.UnpackLog(event, "LogFlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogRegisterProtocolIterator is returned from FilterLogRegisterProtocol and is used to iterate over the raw logs and unpacked data for LogRegisterProtocol events raised by the Contract contract.
type ContractLogRegisterProtocolIterator struct {
	Event *ContractLogRegisterProtocol // Event containing the contract specifics and raw log

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
func (it *ContractLogRegisterProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogRegisterProtocol)
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
		it.Event = new(ContractLogRegisterProtocol)
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
func (it *ContractLogRegisterProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogRegisterProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogRegisterProtocol represents a LogRegisterProtocol event raised by the Contract contract.
type ContractLogRegisterProtocol struct {
	Protocol common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogRegisterProtocol is a free log retrieval operation binding the contract event 0xdfb44ffabf0d3a8f650d3ce43eff98f6d050e7ea1a396d5794f014e7dadabacb.
//
// Solidity: event LogRegisterProtocol(address indexed protocol)
func (_Contract *ContractFilterer) FilterLogRegisterProtocol(opts *bind.FilterOpts, protocol []common.Address) (*ContractLogRegisterProtocolIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogRegisterProtocol", protocolRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogRegisterProtocolIterator{contract: _Contract.contract, event: "LogRegisterProtocol", logs: logs, sub: sub}, nil
}

// WatchLogRegisterProtocol is a free log subscription operation binding the contract event 0xdfb44ffabf0d3a8f650d3ce43eff98f6d050e7ea1a396d5794f014e7dadabacb.
//
// Solidity: event LogRegisterProtocol(address indexed protocol)
func (_Contract *ContractFilterer) WatchLogRegisterProtocol(opts *bind.WatchOpts, sink chan<- *ContractLogRegisterProtocol, protocol []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogRegisterProtocol", protocolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogRegisterProtocol)
				if err := _Contract.contract.UnpackLog(event, "LogRegisterProtocol", log); err != nil {
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

// ParseLogRegisterProtocol is a log parse operation binding the contract event 0xdfb44ffabf0d3a8f650d3ce43eff98f6d050e7ea1a396d5794f014e7dadabacb.
//
// Solidity: event LogRegisterProtocol(address indexed protocol)
func (_Contract *ContractFilterer) ParseLogRegisterProtocol(log types.Log) (*ContractLogRegisterProtocol, error) {
	event := new(ContractLogRegisterProtocol)
	if err := _Contract.contract.UnpackLog(event, "LogRegisterProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogSetMasterContractApprovalIterator is returned from FilterLogSetMasterContractApproval and is used to iterate over the raw logs and unpacked data for LogSetMasterContractApproval events raised by the Contract contract.
type ContractLogSetMasterContractApprovalIterator struct {
	Event *ContractLogSetMasterContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractLogSetMasterContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogSetMasterContractApproval)
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
		it.Event = new(ContractLogSetMasterContractApproval)
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
func (it *ContractLogSetMasterContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogSetMasterContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogSetMasterContractApproval represents a LogSetMasterContractApproval event raised by the Contract contract.
type ContractLogSetMasterContractApproval struct {
	MasterContract common.Address
	User           common.Address
	Approved       bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogSetMasterContractApproval is a free log retrieval operation binding the contract event 0x5f6ebb64ba012a851c6f014e6cad458ddf213d1512049b31cd06365c2b059257.
//
// Solidity: event LogSetMasterContractApproval(address indexed masterContract, address indexed user, bool approved)
func (_Contract *ContractFilterer) FilterLogSetMasterContractApproval(opts *bind.FilterOpts, masterContract []common.Address, user []common.Address) (*ContractLogSetMasterContractApprovalIterator, error) {

	var masterContractRule []interface{}
	for _, masterContractItem := range masterContract {
		masterContractRule = append(masterContractRule, masterContractItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogSetMasterContractApproval", masterContractRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogSetMasterContractApprovalIterator{contract: _Contract.contract, event: "LogSetMasterContractApproval", logs: logs, sub: sub}, nil
}

// WatchLogSetMasterContractApproval is a free log subscription operation binding the contract event 0x5f6ebb64ba012a851c6f014e6cad458ddf213d1512049b31cd06365c2b059257.
//
// Solidity: event LogSetMasterContractApproval(address indexed masterContract, address indexed user, bool approved)
func (_Contract *ContractFilterer) WatchLogSetMasterContractApproval(opts *bind.WatchOpts, sink chan<- *ContractLogSetMasterContractApproval, masterContract []common.Address, user []common.Address) (event.Subscription, error) {

	var masterContractRule []interface{}
	for _, masterContractItem := range masterContract {
		masterContractRule = append(masterContractRule, masterContractItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogSetMasterContractApproval", masterContractRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogSetMasterContractApproval)
				if err := _Contract.contract.UnpackLog(event, "LogSetMasterContractApproval", log); err != nil {
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

// ParseLogSetMasterContractApproval is a log parse operation binding the contract event 0x5f6ebb64ba012a851c6f014e6cad458ddf213d1512049b31cd06365c2b059257.
//
// Solidity: event LogSetMasterContractApproval(address indexed masterContract, address indexed user, bool approved)
func (_Contract *ContractFilterer) ParseLogSetMasterContractApproval(log types.Log) (*ContractLogSetMasterContractApproval, error) {
	event := new(ContractLogSetMasterContractApproval)
	if err := _Contract.contract.UnpackLog(event, "LogSetMasterContractApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogStrategyDivestIterator is returned from FilterLogStrategyDivest and is used to iterate over the raw logs and unpacked data for LogStrategyDivest events raised by the Contract contract.
type ContractLogStrategyDivestIterator struct {
	Event *ContractLogStrategyDivest // Event containing the contract specifics and raw log

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
func (it *ContractLogStrategyDivestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogStrategyDivest)
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
		it.Event = new(ContractLogStrategyDivest)
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
func (it *ContractLogStrategyDivestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogStrategyDivestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogStrategyDivest represents a LogStrategyDivest event raised by the Contract contract.
type ContractLogStrategyDivest struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogStrategyDivest is a free log retrieval operation binding the contract event 0x39aa22060f8dd4d291720311feedf3b72fef47c06c66ccf5c22b502c62e7550a.
//
// Solidity: event LogStrategyDivest(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) FilterLogStrategyDivest(opts *bind.FilterOpts, token []common.Address) (*ContractLogStrategyDivestIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogStrategyDivest", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogStrategyDivestIterator{contract: _Contract.contract, event: "LogStrategyDivest", logs: logs, sub: sub}, nil
}

// WatchLogStrategyDivest is a free log subscription operation binding the contract event 0x39aa22060f8dd4d291720311feedf3b72fef47c06c66ccf5c22b502c62e7550a.
//
// Solidity: event LogStrategyDivest(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) WatchLogStrategyDivest(opts *bind.WatchOpts, sink chan<- *ContractLogStrategyDivest, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogStrategyDivest", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogStrategyDivest)
				if err := _Contract.contract.UnpackLog(event, "LogStrategyDivest", log); err != nil {
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

// ParseLogStrategyDivest is a log parse operation binding the contract event 0x39aa22060f8dd4d291720311feedf3b72fef47c06c66ccf5c22b502c62e7550a.
//
// Solidity: event LogStrategyDivest(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) ParseLogStrategyDivest(log types.Log) (*ContractLogStrategyDivest, error) {
	event := new(ContractLogStrategyDivest)
	if err := _Contract.contract.UnpackLog(event, "LogStrategyDivest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogStrategyInvestIterator is returned from FilterLogStrategyInvest and is used to iterate over the raw logs and unpacked data for LogStrategyInvest events raised by the Contract contract.
type ContractLogStrategyInvestIterator struct {
	Event *ContractLogStrategyInvest // Event containing the contract specifics and raw log

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
func (it *ContractLogStrategyInvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogStrategyInvest)
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
		it.Event = new(ContractLogStrategyInvest)
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
func (it *ContractLogStrategyInvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogStrategyInvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogStrategyInvest represents a LogStrategyInvest event raised by the Contract contract.
type ContractLogStrategyInvest struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogStrategyInvest is a free log retrieval operation binding the contract event 0xb18e7e4f6eac147a63a3bb6beb2d9039c88698623aff3efc4febbc20b0164ee5.
//
// Solidity: event LogStrategyInvest(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) FilterLogStrategyInvest(opts *bind.FilterOpts, token []common.Address) (*ContractLogStrategyInvestIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogStrategyInvest", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogStrategyInvestIterator{contract: _Contract.contract, event: "LogStrategyInvest", logs: logs, sub: sub}, nil
}

// WatchLogStrategyInvest is a free log subscription operation binding the contract event 0xb18e7e4f6eac147a63a3bb6beb2d9039c88698623aff3efc4febbc20b0164ee5.
//
// Solidity: event LogStrategyInvest(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) WatchLogStrategyInvest(opts *bind.WatchOpts, sink chan<- *ContractLogStrategyInvest, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogStrategyInvest", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogStrategyInvest)
				if err := _Contract.contract.UnpackLog(event, "LogStrategyInvest", log); err != nil {
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

// ParseLogStrategyInvest is a log parse operation binding the contract event 0xb18e7e4f6eac147a63a3bb6beb2d9039c88698623aff3efc4febbc20b0164ee5.
//
// Solidity: event LogStrategyInvest(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) ParseLogStrategyInvest(log types.Log) (*ContractLogStrategyInvest, error) {
	event := new(ContractLogStrategyInvest)
	if err := _Contract.contract.UnpackLog(event, "LogStrategyInvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogStrategyLossIterator is returned from FilterLogStrategyLoss and is used to iterate over the raw logs and unpacked data for LogStrategyLoss events raised by the Contract contract.
type ContractLogStrategyLossIterator struct {
	Event *ContractLogStrategyLoss // Event containing the contract specifics and raw log

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
func (it *ContractLogStrategyLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogStrategyLoss)
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
		it.Event = new(ContractLogStrategyLoss)
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
func (it *ContractLogStrategyLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogStrategyLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogStrategyLoss represents a LogStrategyLoss event raised by the Contract contract.
type ContractLogStrategyLoss struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogStrategyLoss is a free log retrieval operation binding the contract event 0x8f1f26eb9b6aa8689dbdd519ead1999d9c8819d4738e403b2003b18197d9cf97.
//
// Solidity: event LogStrategyLoss(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) FilterLogStrategyLoss(opts *bind.FilterOpts, token []common.Address) (*ContractLogStrategyLossIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogStrategyLoss", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogStrategyLossIterator{contract: _Contract.contract, event: "LogStrategyLoss", logs: logs, sub: sub}, nil
}

// WatchLogStrategyLoss is a free log subscription operation binding the contract event 0x8f1f26eb9b6aa8689dbdd519ead1999d9c8819d4738e403b2003b18197d9cf97.
//
// Solidity: event LogStrategyLoss(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) WatchLogStrategyLoss(opts *bind.WatchOpts, sink chan<- *ContractLogStrategyLoss, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogStrategyLoss", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogStrategyLoss)
				if err := _Contract.contract.UnpackLog(event, "LogStrategyLoss", log); err != nil {
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

// ParseLogStrategyLoss is a log parse operation binding the contract event 0x8f1f26eb9b6aa8689dbdd519ead1999d9c8819d4738e403b2003b18197d9cf97.
//
// Solidity: event LogStrategyLoss(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) ParseLogStrategyLoss(log types.Log) (*ContractLogStrategyLoss, error) {
	event := new(ContractLogStrategyLoss)
	if err := _Contract.contract.UnpackLog(event, "LogStrategyLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogStrategyProfitIterator is returned from FilterLogStrategyProfit and is used to iterate over the raw logs and unpacked data for LogStrategyProfit events raised by the Contract contract.
type ContractLogStrategyProfitIterator struct {
	Event *ContractLogStrategyProfit // Event containing the contract specifics and raw log

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
func (it *ContractLogStrategyProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogStrategyProfit)
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
		it.Event = new(ContractLogStrategyProfit)
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
func (it *ContractLogStrategyProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogStrategyProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogStrategyProfit represents a LogStrategyProfit event raised by the Contract contract.
type ContractLogStrategyProfit struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogStrategyProfit is a free log retrieval operation binding the contract event 0x911c9f20a03edabcbcbd18dca1174cce47a91b234ced7a5a3c60ba0d5b56c5d2.
//
// Solidity: event LogStrategyProfit(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) FilterLogStrategyProfit(opts *bind.FilterOpts, token []common.Address) (*ContractLogStrategyProfitIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogStrategyProfit", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogStrategyProfitIterator{contract: _Contract.contract, event: "LogStrategyProfit", logs: logs, sub: sub}, nil
}

// WatchLogStrategyProfit is a free log subscription operation binding the contract event 0x911c9f20a03edabcbcbd18dca1174cce47a91b234ced7a5a3c60ba0d5b56c5d2.
//
// Solidity: event LogStrategyProfit(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) WatchLogStrategyProfit(opts *bind.WatchOpts, sink chan<- *ContractLogStrategyProfit, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogStrategyProfit", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogStrategyProfit)
				if err := _Contract.contract.UnpackLog(event, "LogStrategyProfit", log); err != nil {
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

// ParseLogStrategyProfit is a log parse operation binding the contract event 0x911c9f20a03edabcbcbd18dca1174cce47a91b234ced7a5a3c60ba0d5b56c5d2.
//
// Solidity: event LogStrategyProfit(address indexed token, uint256 amount)
func (_Contract *ContractFilterer) ParseLogStrategyProfit(log types.Log) (*ContractLogStrategyProfit, error) {
	event := new(ContractLogStrategyProfit)
	if err := _Contract.contract.UnpackLog(event, "LogStrategyProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogStrategyQueuedIterator is returned from FilterLogStrategyQueued and is used to iterate over the raw logs and unpacked data for LogStrategyQueued events raised by the Contract contract.
type ContractLogStrategyQueuedIterator struct {
	Event *ContractLogStrategyQueued // Event containing the contract specifics and raw log

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
func (it *ContractLogStrategyQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogStrategyQueued)
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
		it.Event = new(ContractLogStrategyQueued)
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
func (it *ContractLogStrategyQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogStrategyQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogStrategyQueued represents a LogStrategyQueued event raised by the Contract contract.
type ContractLogStrategyQueued struct {
	Token    common.Address
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogStrategyQueued is a free log retrieval operation binding the contract event 0x6f7ccdf3f86039e5a1dcf6028bf7b4773cbf7a234716ba2e5392b12bb0f8558f.
//
// Solidity: event LogStrategyQueued(address indexed token, address indexed strategy)
func (_Contract *ContractFilterer) FilterLogStrategyQueued(opts *bind.FilterOpts, token []common.Address, strategy []common.Address) (*ContractLogStrategyQueuedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogStrategyQueued", tokenRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogStrategyQueuedIterator{contract: _Contract.contract, event: "LogStrategyQueued", logs: logs, sub: sub}, nil
}

// WatchLogStrategyQueued is a free log subscription operation binding the contract event 0x6f7ccdf3f86039e5a1dcf6028bf7b4773cbf7a234716ba2e5392b12bb0f8558f.
//
// Solidity: event LogStrategyQueued(address indexed token, address indexed strategy)
func (_Contract *ContractFilterer) WatchLogStrategyQueued(opts *bind.WatchOpts, sink chan<- *ContractLogStrategyQueued, token []common.Address, strategy []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogStrategyQueued", tokenRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogStrategyQueued)
				if err := _Contract.contract.UnpackLog(event, "LogStrategyQueued", log); err != nil {
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

// ParseLogStrategyQueued is a log parse operation binding the contract event 0x6f7ccdf3f86039e5a1dcf6028bf7b4773cbf7a234716ba2e5392b12bb0f8558f.
//
// Solidity: event LogStrategyQueued(address indexed token, address indexed strategy)
func (_Contract *ContractFilterer) ParseLogStrategyQueued(log types.Log) (*ContractLogStrategyQueued, error) {
	event := new(ContractLogStrategyQueued)
	if err := _Contract.contract.UnpackLog(event, "LogStrategyQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogStrategySetIterator is returned from FilterLogStrategySet and is used to iterate over the raw logs and unpacked data for LogStrategySet events raised by the Contract contract.
type ContractLogStrategySetIterator struct {
	Event *ContractLogStrategySet // Event containing the contract specifics and raw log

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
func (it *ContractLogStrategySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogStrategySet)
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
		it.Event = new(ContractLogStrategySet)
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
func (it *ContractLogStrategySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogStrategySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogStrategySet represents a LogStrategySet event raised by the Contract contract.
type ContractLogStrategySet struct {
	Token    common.Address
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogStrategySet is a free log retrieval operation binding the contract event 0x03e6352a885adc4cc54767592939c3b1bbd65685658c3beaaba66a888120e217.
//
// Solidity: event LogStrategySet(address indexed token, address indexed strategy)
func (_Contract *ContractFilterer) FilterLogStrategySet(opts *bind.FilterOpts, token []common.Address, strategy []common.Address) (*ContractLogStrategySetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogStrategySet", tokenRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogStrategySetIterator{contract: _Contract.contract, event: "LogStrategySet", logs: logs, sub: sub}, nil
}

// WatchLogStrategySet is a free log subscription operation binding the contract event 0x03e6352a885adc4cc54767592939c3b1bbd65685658c3beaaba66a888120e217.
//
// Solidity: event LogStrategySet(address indexed token, address indexed strategy)
func (_Contract *ContractFilterer) WatchLogStrategySet(opts *bind.WatchOpts, sink chan<- *ContractLogStrategySet, token []common.Address, strategy []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogStrategySet", tokenRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogStrategySet)
				if err := _Contract.contract.UnpackLog(event, "LogStrategySet", log); err != nil {
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

// ParseLogStrategySet is a log parse operation binding the contract event 0x03e6352a885adc4cc54767592939c3b1bbd65685658c3beaaba66a888120e217.
//
// Solidity: event LogStrategySet(address indexed token, address indexed strategy)
func (_Contract *ContractFilterer) ParseLogStrategySet(log types.Log) (*ContractLogStrategySet, error) {
	event := new(ContractLogStrategySet)
	if err := _Contract.contract.UnpackLog(event, "LogStrategySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogStrategyTargetPercentageIterator is returned from FilterLogStrategyTargetPercentage and is used to iterate over the raw logs and unpacked data for LogStrategyTargetPercentage events raised by the Contract contract.
type ContractLogStrategyTargetPercentageIterator struct {
	Event *ContractLogStrategyTargetPercentage // Event containing the contract specifics and raw log

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
func (it *ContractLogStrategyTargetPercentageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogStrategyTargetPercentage)
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
		it.Event = new(ContractLogStrategyTargetPercentage)
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
func (it *ContractLogStrategyTargetPercentageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogStrategyTargetPercentageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogStrategyTargetPercentage represents a LogStrategyTargetPercentage event raised by the Contract contract.
type ContractLogStrategyTargetPercentage struct {
	Token            common.Address
	TargetPercentage *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogStrategyTargetPercentage is a free log retrieval operation binding the contract event 0x7543af99b5602c06e62da0631b5308489a5ff859150105a623b6eb15e8deae0b.
//
// Solidity: event LogStrategyTargetPercentage(address indexed token, uint256 targetPercentage)
func (_Contract *ContractFilterer) FilterLogStrategyTargetPercentage(opts *bind.FilterOpts, token []common.Address) (*ContractLogStrategyTargetPercentageIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogStrategyTargetPercentage", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogStrategyTargetPercentageIterator{contract: _Contract.contract, event: "LogStrategyTargetPercentage", logs: logs, sub: sub}, nil
}

// WatchLogStrategyTargetPercentage is a free log subscription operation binding the contract event 0x7543af99b5602c06e62da0631b5308489a5ff859150105a623b6eb15e8deae0b.
//
// Solidity: event LogStrategyTargetPercentage(address indexed token, uint256 targetPercentage)
func (_Contract *ContractFilterer) WatchLogStrategyTargetPercentage(opts *bind.WatchOpts, sink chan<- *ContractLogStrategyTargetPercentage, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogStrategyTargetPercentage", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogStrategyTargetPercentage)
				if err := _Contract.contract.UnpackLog(event, "LogStrategyTargetPercentage", log); err != nil {
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

// ParseLogStrategyTargetPercentage is a log parse operation binding the contract event 0x7543af99b5602c06e62da0631b5308489a5ff859150105a623b6eb15e8deae0b.
//
// Solidity: event LogStrategyTargetPercentage(address indexed token, uint256 targetPercentage)
func (_Contract *ContractFilterer) ParseLogStrategyTargetPercentage(log types.Log) (*ContractLogStrategyTargetPercentage, error) {
	event := new(ContractLogStrategyTargetPercentage)
	if err := _Contract.contract.UnpackLog(event, "LogStrategyTargetPercentage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogTransferIterator is returned from FilterLogTransfer and is used to iterate over the raw logs and unpacked data for LogTransfer events raised by the Contract contract.
type ContractLogTransferIterator struct {
	Event *ContractLogTransfer // Event containing the contract specifics and raw log

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
func (it *ContractLogTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogTransfer)
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
		it.Event = new(ContractLogTransfer)
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
func (it *ContractLogTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogTransfer represents a LogTransfer event raised by the Contract contract.
type ContractLogTransfer struct {
	Token common.Address
	From  common.Address
	To    common.Address
	Share *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogTransfer is a free log retrieval operation binding the contract event 0x6eabe333476233fd382224f233210cb808a7bc4c4de64f9d76628bf63c677b1a.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 share)
func (_Contract *ContractFilterer) FilterLogTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*ContractLogTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogTransferIterator{contract: _Contract.contract, event: "LogTransfer", logs: logs, sub: sub}, nil
}

// WatchLogTransfer is a free log subscription operation binding the contract event 0x6eabe333476233fd382224f233210cb808a7bc4c4de64f9d76628bf63c677b1a.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 share)
func (_Contract *ContractFilterer) WatchLogTransfer(opts *bind.WatchOpts, sink chan<- *ContractLogTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogTransfer)
				if err := _Contract.contract.UnpackLog(event, "LogTransfer", log); err != nil {
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

// ParseLogTransfer is a log parse operation binding the contract event 0x6eabe333476233fd382224f233210cb808a7bc4c4de64f9d76628bf63c677b1a.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 share)
func (_Contract *ContractFilterer) ParseLogTransfer(log types.Log) (*ContractLogTransfer, error) {
	event := new(ContractLogTransfer)
	if err := _Contract.contract.UnpackLog(event, "LogTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogWhiteListMasterContractIterator is returned from FilterLogWhiteListMasterContract and is used to iterate over the raw logs and unpacked data for LogWhiteListMasterContract events raised by the Contract contract.
type ContractLogWhiteListMasterContractIterator struct {
	Event *ContractLogWhiteListMasterContract // Event containing the contract specifics and raw log

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
func (it *ContractLogWhiteListMasterContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogWhiteListMasterContract)
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
		it.Event = new(ContractLogWhiteListMasterContract)
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
func (it *ContractLogWhiteListMasterContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogWhiteListMasterContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogWhiteListMasterContract represents a LogWhiteListMasterContract event raised by the Contract contract.
type ContractLogWhiteListMasterContract struct {
	MasterContract common.Address
	Approved       bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogWhiteListMasterContract is a free log retrieval operation binding the contract event 0x31a1e0eac44b54ac6c2a2efa87e92c83405ffcf33fceef02a7bca695130e2600.
//
// Solidity: event LogWhiteListMasterContract(address indexed masterContract, bool approved)
func (_Contract *ContractFilterer) FilterLogWhiteListMasterContract(opts *bind.FilterOpts, masterContract []common.Address) (*ContractLogWhiteListMasterContractIterator, error) {

	var masterContractRule []interface{}
	for _, masterContractItem := range masterContract {
		masterContractRule = append(masterContractRule, masterContractItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogWhiteListMasterContract", masterContractRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogWhiteListMasterContractIterator{contract: _Contract.contract, event: "LogWhiteListMasterContract", logs: logs, sub: sub}, nil
}

// WatchLogWhiteListMasterContract is a free log subscription operation binding the contract event 0x31a1e0eac44b54ac6c2a2efa87e92c83405ffcf33fceef02a7bca695130e2600.
//
// Solidity: event LogWhiteListMasterContract(address indexed masterContract, bool approved)
func (_Contract *ContractFilterer) WatchLogWhiteListMasterContract(opts *bind.WatchOpts, sink chan<- *ContractLogWhiteListMasterContract, masterContract []common.Address) (event.Subscription, error) {

	var masterContractRule []interface{}
	for _, masterContractItem := range masterContract {
		masterContractRule = append(masterContractRule, masterContractItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogWhiteListMasterContract", masterContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogWhiteListMasterContract)
				if err := _Contract.contract.UnpackLog(event, "LogWhiteListMasterContract", log); err != nil {
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

// ParseLogWhiteListMasterContract is a log parse operation binding the contract event 0x31a1e0eac44b54ac6c2a2efa87e92c83405ffcf33fceef02a7bca695130e2600.
//
// Solidity: event LogWhiteListMasterContract(address indexed masterContract, bool approved)
func (_Contract *ContractFilterer) ParseLogWhiteListMasterContract(log types.Log) (*ContractLogWhiteListMasterContract, error) {
	event := new(ContractLogWhiteListMasterContract)
	if err := _Contract.contract.UnpackLog(event, "LogWhiteListMasterContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogWithdrawIterator is returned from FilterLogWithdraw and is used to iterate over the raw logs and unpacked data for LogWithdraw events raised by the Contract contract.
type ContractLogWithdrawIterator struct {
	Event *ContractLogWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractLogWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogWithdraw)
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
		it.Event = new(ContractLogWithdraw)
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
func (it *ContractLogWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogWithdraw represents a LogWithdraw event raised by the Contract contract.
type ContractLogWithdraw struct {
	Token  common.Address
	From   common.Address
	To     common.Address
	Amount *big.Int
	Share  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogWithdraw is a free log retrieval operation binding the contract event 0xad9ab9ee6953d4d177f4a03b3a3ac3178ffcb9816319f348060194aa76b14486.
//
// Solidity: event LogWithdraw(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 share)
func (_Contract *ContractFilterer) FilterLogWithdraw(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*ContractLogWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogWithdraw", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogWithdrawIterator{contract: _Contract.contract, event: "LogWithdraw", logs: logs, sub: sub}, nil
}

// WatchLogWithdraw is a free log subscription operation binding the contract event 0xad9ab9ee6953d4d177f4a03b3a3ac3178ffcb9816319f348060194aa76b14486.
//
// Solidity: event LogWithdraw(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 share)
func (_Contract *ContractFilterer) WatchLogWithdraw(opts *bind.WatchOpts, sink chan<- *ContractLogWithdraw, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogWithdraw", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogWithdraw)
				if err := _Contract.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
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

// ParseLogWithdraw is a log parse operation binding the contract event 0xad9ab9ee6953d4d177f4a03b3a3ac3178ffcb9816319f348060194aa76b14486.
//
// Solidity: event LogWithdraw(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 share)
func (_Contract *ContractFilterer) ParseLogWithdraw(log types.Log) (*ContractLogWithdraw, error) {
	event := new(ContractLogWithdraw)
	if err := _Contract.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
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
