// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oneinch_router

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

// GenericRouterSwapDescription is an auto generated low-level Go binding around an user-defined struct.
type GenericRouterSwapDescription struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceiver     common.Address
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
}

// OrderLibOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderLibOrder struct {
	Salt          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	Receiver      common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
	Offsets       *big.Int
	Interactions  []byte
}

// OrderRFQLibOrderRFQ is an auto generated low-level Go binding around an user-defined struct.
type OrderRFQLibOrderRFQ struct {
	Info          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdvanceNonceFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArbitraryStaticCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthDepositRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetAmountCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOneAmountShouldBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitLengthTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PredicateIsNotTrue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQBadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQPrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQSwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQZeroTargetIsForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemainingAmountIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReservesCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReturnAmountIsNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafePermitBadLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"name\":\"SimulationResults\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapAmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromMakerToTakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromTakerToMakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongGetter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMinReturn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroReturnAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTargetIsForbidden\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingRaw\",\"type\":\"uint256\"}],\"name\":\"OrderCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilledRFQ\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"advanceNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"and\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"arbitraryStaticCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderRemaining\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalMask\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"checkPredicate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"clipperSwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"eq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQCompact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderRFQTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderRFQToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"gt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"invalidatorForOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"lt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"name\":\"nonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"or\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remainingRaw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"remainingsRaw\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"executor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structGenericRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"timestampBelow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNonceAccount\",\"type\":\"uint256\"}],\"name\":\"timestampBelowAndNonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3Swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3SwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"unoswapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"unoswapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101e06040523480156200001257600080fd5b506040516200872f3803806200872f83398181016040528101906200003891906200047d565b80808283846040518060400160405280601881526020017f31696e6368204167677265676174696f6e20526f7574657200000000000000008152506040518060400160405280600181526020017f350000000000000000000000000000000000000000000000000000000000000081525060008280519060200120905060008280519060200120905060007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f90508260e081815250508161010081815250504660a0818152505062000112818484620002f760201b60201c565b608081815250503073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505080610120818152505050505050506200017b6200016f6200033360201b60201c565b6200033b60201b60201c565b8073ffffffffffffffffffffffffffffffffffffffff166101408173ffffffffffffffffffffffffffffffffffffffff1681525050508073ffffffffffffffffffffffffffffffffffffffff166101608173ffffffffffffffffffffffffffffffffffffffff1681525050508073ffffffffffffffffffffffffffffffffffffffff166101808173ffffffffffffffffffffffffffffffffffffffff1681525050508073ffffffffffffffffffffffffffffffffffffffff166101a08173ffffffffffffffffffffffffffffffffffffffff1681525050508073ffffffffffffffffffffffffffffffffffffffff166101c08173ffffffffffffffffffffffffffffffffffffffff168152505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620002f0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5062000553565b6000838383463060405160200162000314959493929190620004f6565b6040516020818303038152906040528051906020012090509392505050565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620004318262000404565b9050919050565b6000620004458262000424565b9050919050565b620004578162000438565b81146200046357600080fd5b50565b60008151905062000477816200044c565b92915050565b600060208284031215620004965762000495620003ff565b5b6000620004a68482850162000466565b91505092915050565b6000819050919050565b620004c481620004af565b82525050565b6000819050919050565b620004df81620004ca565b82525050565b620004f08162000424565b82525050565b600060a0820190506200050d6000830188620004b9565b6200051c6020830187620004b9565b6200052b6040830186620004b9565b6200053a6060830185620004d4565b620005496080830184620004e5565b9695505050505050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c0516181046200062b60003960008181614a7a01528181614b0101528181614ba201528181614d0d01528181614daa0152614e2b01526000505060008181612bf001528181612d4b0152612dcc015260008181613e2f0152613feb015260008181610c4701528181610cda01528181610ed301528181610ff8015261104d015260006141590152600061419b0152600061417a015260006140af015260006141050152600061412e01526181046000f3fe6080604052600436106102765760003560e01c80637e54f0921161014f578063bf15fcd8116100c1578063d365c6951161007a578063d365c69514610a84578063e449022e14610ac3578063e5d7bde614610af3578063f2fde38b14610b25578063f78dc25314610b4e578063fa461e3314610b7e57610285565b8063bf15fcd81461093c578063bfa7514314610979578063c53a0292146109b6578063c805a666146109cd578063ca4ece2214610a0a578063cf6fc6e314610a4757610285565b8063942461bb11610113578063942461bb1461080e5780639570eeee1461084b578063bc1ed74c1461087d578063bc80f1a8146108ba578063bd61951d146108ea578063bddccd351461091357610285565b80637e54f09214610736578063825caba11461077357806383197ef01461079c57806384bd6d29146107b35780638da5cb5b146107e357610285565b80635a099843116101e857806370ae92d2116101ac57806370ae92d21461061457806370ccbd3114610651578063715018a61461069057806372c244a8146106a757806374261145146106d057806378e3214f1461070d57610285565b80635a099843146104f957806362e238bb1461052b57806363592c2b1461055d5780636c8382501461059a5780636fe7b0ba146105d757610285565b80632d9a56f61161023a5780632d9a56f61461039557806337e7316f146103d35780633c15fd91146104105780633eca9c0a1461044d5780634f38e2b81461047f57806356f16124146104bc57610285565b80630502b1c51461028a578063093d4fa5146102ba57806312aa3caf146102ea5780632521b9301461031b5780632cc2878d1461035857610285565b3661028557610283610ba7565b005b600080fd5b6102a4600480360381019061029f9190616059565b610bb1565b6040516102b191906160f0565b60405180910390f35b6102d460048036038101906102cf91906161bd565b610bcc565b6040516102e191906160f0565b60405180910390f35b61030460048036038101906102ff919061633f565b6111ab565b6040516103129291906163e9565b60405180910390f35b34801561032757600080fd5b50610342600480360381019061033d9190616412565b6115af565b60405161034f91906160f0565b60405180910390f35b34801561036457600080fd5b5061037f600480360381019061037a91906164e1565b6115f7565b60405161038c9190616529565b60405180910390f35b3480156103a157600080fd5b506103bc60048036038101906103b79190616564565b611644565b6040516103ca9291906165bc565b60405180910390f35b3480156103df57600080fd5b506103fa60048036038101906103f59190616564565b611787565b60405161040791906165e5565b60405180910390f35b34801561041c57600080fd5b5061043760048036038101906104329190616412565b6117aa565b60405161044491906160f0565b60405180910390f35b61046760048036038101906104629190616771565b6117f3565b604051610476939291906167e7565b60405180910390f35b34801561048b57600080fd5b506104a660048036038101906104a1919061681e565b611815565b6040516104b39190616529565b60405180910390f35b3480156104c857600080fd5b506104e360048036038101906104de919061687e565b61183f565b6040516104f091906160f0565b60405180910390f35b610513600480360381019061050e91906168be565b61189a565b604051610522939291906167e7565b60405180910390f35b61054560048036038101906105409190616949565b611a3c565b604051610554939291906167e7565b60405180910390f35b34801561056957600080fd5b50610584600480360381019061057f91906164e1565b611a66565b6040516105919190616529565b60405180910390f35b3480156105a657600080fd5b506105c160048036038101906105bc9190616564565b611a72565b6040516105ce9190616529565b60405180910390f35b3480156105e357600080fd5b506105fe60048036038101906105f9919061681e565b611aa2565b60405161060b9190616529565b60405180910390f35b34801561062057600080fd5b5061063b60048036038101906106369190616a34565b611acc565b60405161064891906160f0565b60405180910390f35b34801561065d57600080fd5b5061067860048036038101906106739190616a61565b611ae4565b604051610687939291906167e7565b60405180910390f35b34801561069c57600080fd5b506106a5611b38565b005b3480156106b357600080fd5b506106ce60048036038101906106c99190616b5a565b611b4c565b005b3480156106dc57600080fd5b506106f760048036038101906106f2919061681e565b611c71565b6040516107049190616529565b60405180910390f35b34801561071957600080fd5b50610734600480360381019061072f9190616b87565b611cf9565b005b34801561074257600080fd5b5061075d60048036038101906107589190616bc7565b611d30565b60405161076a91906160f0565b60405180910390f35b34801561077f57600080fd5b5061079a600480360381019061079591906164e1565b611d4d565b005b3480156107a857600080fd5b506107b1611d5c565b005b6107cd60048036038101906107c89190616bf4565b611d7d565b6040516107da91906160f0565b60405180910390f35b3480156107ef57600080fd5b506107f8611d9e565b6040516108059190616cb9565b60405180910390f35b34801561081a57600080fd5b5061083560048036038101906108309190616d97565b611dc7565b6040516108429190616e9e565b60405180910390f35b61086560048036038101906108609190616ec0565b611e8d565b604051610874939291906167e7565b60405180910390f35b34801561088957600080fd5b506108a4600480360381019061089f9190616bc7565b612036565b6040516108b191906160f0565b60405180910390f35b6108d460048036038101906108cf9190616f2a565b612095565b6040516108e191906160f0565b60405180910390f35b3480156108f657600080fd5b50610911600480360381019061090c9190616fb2565b6120af565b005b34801561091f57600080fd5b5061093a60048036038101906109359190617012565b61215d565b005b34801561094857600080fd5b50610963600480360381019061095e9190616fb2565b61216c565b60405161097091906160f0565b60405180910390f35b34801561098557600080fd5b506109a0600480360381019061099b919061681e565b6121c3565b6040516109ad9190616529565b60405180910390f35b3480156109c257600080fd5b506109cb61224c565b005b3480156109d957600080fd5b506109f460048036038101906109ef9190617052565b612258565b604051610a0191906160f0565b60405180910390f35b348015610a1657600080fd5b50610a316004803603810190610a2c919061681e565b6122a7565b604051610a3e9190616529565b60405180910390f35b348015610a5357600080fd5b50610a6e6004803603810190610a69919061687e565b6122d1565b604051610a7b9190616529565b60405180910390f35b348015610a9057600080fd5b50610aab6004803603810190610aa69190617152565b61231d565b604051610aba939291906167e7565b60405180910390f35b610add6004803603810190610ad89190617286565b6123cb565b604051610aea91906160f0565b60405180910390f35b610b0d6004803603810190610b0891906172fa565b6123e4565b604051610b1c939291906167e7565b60405180910390f35b348015610b3157600080fd5b50610b4c6004803603810190610b479190616a34565b6130b6565b005b610b686004803603810190610b6391906173f8565b613139565b604051610b7591906160f0565b60405180910390f35b348015610b8a57600080fd5b50610ba56004803603810190610ba091906174c8565b613155565b005b610baf61337b565b565b6000610bc13387878787876133e2565b905095945050505050565b600080600073ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff161490508015610c4557863414610c40576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dd6565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1603610d6d576001905060003414610cd6576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f0000000000000000000000000000000000000000000000000000000000000000905060006323b872dd60e01b90506000632e1a7d4d60e01b90506040518281523360048201523060248201528a60448201526000806064836000885af1610d43573d6000823e3d81fd5b8181528a60048201526000806024836000885af1610d64573d6000823e3d81fd5b50505050610dd5565b60003414610da7576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dd4338c898c73ffffffffffffffffffffffffffffffffffffffff16613976909392919063ffffffff16565b5b5b8015610e9b5760008b905060006327a9b42460e01b90506040518181528a60048201528960248201528860448201528760648201528c60848201528560ff1c601b0160a48201528660c48201527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff861660e48201526101206101048201527f0531494e43480000000000000000000000000000000000000000000000000000610143820152600080610149838d875af1610e93573d6000823e3d81fd5b50505061119a565b600073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff161480610f2157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff16145b156110d65760008b90506000634cb6864c60e01b90506040518181528b60048201528960248201528860448201528760648201528a1560018114610f6a57306084830152610f71565b8d60848301525b508560ff1c601b0160a48201528660c48201527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff861660e48201526101206101048201527f0531494e43480000000000000000000000000000000000000000000000000000610143820152600080610149836000875af1610ff5573d6000823e3d81fd5b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff16036110cf5760007f00000000000000000000000000000000000000000000000000000000000000009050600063d0e30db060e01b9050600063a9059cbb60e01b90506040518281526000806004838f885af16110a3573d6000823e3d81fd5b8181528f60048201528b60248201526000806044836000885af16110ca573d6000823e3d81fd5b505050505b5050611199565b60008b90506000632b651a6c60e01b90506040518181528b60048201528a60248201528960448201528860648201528760848201528c60a48201528560ff1c601b0160c48201528660e48201527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff86166101048201526101406101248201527f0531494e43480000000000000000000000000000000000000000000000000000610163820152600080610169836000875af1611195573d6000823e3d81fd5b5050505b5b859150509998505050505050505050565b60008060008760a00135036111ec576040517f0262dde400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000876000016020810190611201919061753c565b90506000886020016020810190611218919061753c565b9050600061123b8373ffffffffffffffffffffffffffffffffffffffff16613a15565b9050600060028b60c00135161461129b578061125857600061125e565b89608001355b3411611296576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112e6565b806112a75760006112ad565b89608001355b34146112e5576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b806113675760008989905011156113235761132289898573ffffffffffffffffffffffffffffffffffffffff16613a979092919063ffffffff16565b5b611366338b604001602081019061133a9190617569565b8c608001358673ffffffffffffffffffffffffffffffffffffffff16613976909392919063ffffffff16565b5b6113788b338c608001358a8a613b2b565b896080013593506113a8308373ffffffffffffffffffffffffffffffffffffffff16613b7490919063ffffffff16565b9450600085036113e4576040517f28ebf24700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848060019003955050600060018b60c0013516146114d3576000611427308573ffffffffffffffffffffffffffffffffffffffff16613b7490919063ffffffff16565b9050600181111561147557808060019003915050808561144791906175c5565b945061147433828673ffffffffffffffffffffffffffffffffffffffff16613c269092919063ffffffff16565b5b848b60a0013561148591906175f9565b8b608001358761149591906175f9565b10156114cd576040517ff32bec2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50611512565b8960a00135851015611511576040517ff32bec2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b60008073ffffffffffffffffffffffffffffffffffffffff168b606001602081019061153e9190617569565b73ffffffffffffffffffffffffffffffffffffffff1614611571578a606001602081019061156c9190617569565b611573565b335b90506115a081878573ffffffffffffffffffffffffffffffffffffffff16613c269092919063ffffffff16565b50505050965096945050505050565b60006115dc83838a73ffffffffffffffffffffffffffffffffffffffff16613a979092919063ffffffff16565b6115e98988888888613d57565b905098975050505050505050565b60008060d083901c65ffffffffffff169050600060a084901c65ffffffffffff169050600084905061162883611a66565b801561163a575061163981836122d1565b5b9350505050919050565b6000803373ffffffffffffffffffffffffffffffffffffffff168360600160208101906116719190616a34565b73ffffffffffffffffffffffffffffffffffffffff16146116be576040517f4ca8886700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116c783611787565b90506002600082815260200190815260200160002054915060018203611719576040517f41a26a6300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167fcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1828460405161176192919061763b565b60405180910390a260016002600083815260200190815260200160002081905550915091565b60006117a36117946140ab565b836141c590919063ffffffff16565b9050919050565b60006117d783838a73ffffffffffffffffffffffffffffffffffffffff16613a979092919063ffffffff16565b6117e58989898989896133e2565b905098975050505050505050565b6000806000611805878787873361189a565b9250925092509450945094915050565b6000806000611824858561423d565b9150915081801561183457508581115b925050509392505050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054905092915050565b60008060006118b96118aa6140ab565b896145b390919063ffffffff16565b905060007f400000000000000000000000000000000000000000000000000000000000000086161461199e5760007f200000000000000000000000000000000000000000000000000000000000000086161415801561191c575060418787905014155b15611953576040517f17c2b1f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119638860600151828989614604565b611999576040517f17c2b1f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119e5565b6119ae8860600151828989614659565b6119e4576040517f17c2b1f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b6119f0888686614712565b80935081945050507fc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca781278184604051611a2992919061763b565b60405180910390a1955095509592505050565b6000806000611a528b8b8b8b8b8b8b8b336123e4565b925092509250985098509895505050505050565b60008142109050919050565b6000806000611a88611a8385614f42565b61423d565b91509150818015611a995750600181145b92505050919050565b6000806000611ab1858561423d565b91509150818015611ac157508581145b925050509392505050565b60016020528060005260406000206000915090505481565b6000806000611b1885858c6040015173ffffffffffffffffffffffffffffffffffffffff16613a979092919063ffffffff16565b611b258a8a8a8a8a61189a565b9250925092509750975097945050505050565b611b40614f59565b611b4a6000614fd7565b565b60008160ff1603611b89576040517fbd71636d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008160ff16600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611bd99190617664565b905080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff167ffc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db82604051611c6591906160f0565b60405180910390a25050565b600080600080600090505b60008188901c63ffffffff1693508314611cea57600080611cad888886908892611ca8939291906176a2565b61423d565b91509150818015611cbe5750600181145b15611cd157600195505050505050611cf2565b8493505050602081611ce39190617664565b9050611c7c565b506000925050505b9392505050565b611d01614f59565b611d2c33828473ffffffffffffffffffffffffffffffffffffffff16613c269092919063ffffffff16565b5050565b600060026000838152602001908152602001600020549050919050565b611d593382600061509b565b50565b611d64614f59565b3373ffffffffffffffffffffffffffffffffffffffff16ff5b6000611d9089338a8a8a8a8a8a8a610bcc565b905098975050505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606000825167ffffffffffffffff811115611de657611de5616616565b5b604051908082528060200260200182016040528015611e145781602001602082028036833780820191505090505b50905060005b8351811015611e835760026000858381518110611e3a57611e396176dd565b5b6020026020010151815260200190815260200160002054828281518110611e6457611e636176dd565b5b6020026020010181815250508080611e7b9061770c565b915050611e1a565b5080915050919050565b6000806000611eac611e9d6140ab565b886145b390919063ffffffff16565b905060007f4000000000000000000000000000000000000000000000000000000000000000851614611f995760007f2000000000000000000000000000000000000000000000000000000000000000851614611f4d57611f12876060015182888861517d565b611f48576040517f17c2b1f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611f94565b611f5d8760600151828888615204565b611f93576040517f17c2b1f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b611fe0565b611fa9876060015182888861525d565b611fdf576040517f17c2b1f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b611feb878533614712565b80935081945050507fc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127818460405161202492919061763b565b60405180910390a19450945094915050565b6000806002600084815260200190815260200160002054905060008103612089576040517fb838de9600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018103915050919050565b60006120a48686868686613d57565b905095945050505050565b6000808473ffffffffffffffffffffffffffffffffffffffff1684846040516120d9929190617793565b600060405180830381855af49150503d8060008114612114576040519150601f19603f3d011682016040523d82523d6000602084013e612119565b606091505b509150915081816040517f1934afc800000000000000000000000000000000000000000000000000000000815260040161215492919061782b565b60405180910390fd5b61216833838361509b565b5050565b600080600061217c8686866152f8565b91509150816121b7576040517f1f1b8f6100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80925050509392505050565b600080600080600090505b60008188901c63ffffffff169350831461223d576000806121ff8888869088926121fa939291906176a2565b61423d565b91509150811580612211575060018114155b1561222457600095505050505050612245565b84935050506020816122369190617664565b90506121ce565b506001925050505b9392505050565b6122566001611b4c565b565b600061228583838c73ffffffffffffffffffffffffffffffffffffffff16613a979092919063ffffffff16565b6122968c8c8c8c8c8c8c8c8c610bcc565b90509b9a5050505050505050505050565b60008060006122b6858561423d565b915091508180156122c657508581105b925050509392505050565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414905092915050565b60008060006014858590501015612360576040517fd9e1c6dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600036600061236f888861532a565b9250925092506123a082828573ffffffffffffffffffffffffffffffffffffffff16613a979092919063ffffffff16565b5050506123b48e8e8e8e8e8e8e8e8e6123e4565b9250925092509b509b509b98505050505050505050565b60006123da3386868686613d57565b9050949350505050565b60008060008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff160361244e576040517fb0c4d05f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124578c611787565b9050368c9050879350869250600060026000848152602001908152602001600020549050600181036124b5576040517fecef366400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168260a00160208101906124e09190616a34565b73ffffffffffffffffffffffffffffffffffffffff161415801561254257503373ffffffffffffffffffffffffffffffffffffffff168260a00160208101906125299190616a34565b73ffffffffffffffffffffffffffffffffffffffff1614155b15612579576040517fd4dfdafe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081036126bf5761259f8260600160208101906125979190616a34565b848f8f614659565b6125d5576040517f5cd5d23300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160c0013590503660006125e884615387565b9150915060007f80000000000000000000000000000000000000000000000000000000000000008a16148015612622575060148282905010155b156126b8576000366000612636858561532a565b92509250925061266782828573ffffffffffffffffffffffffffffffffffffffff16613a979092919063ffffffff16565b6000600260008a815260200190815260200160002054146126b4576040517fc5f2be5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505b50506126c6565b6001810390505b60006126d183614f42565b90501115612719576126e282611a72565b612718576040517fb6629c0200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b6000851415156000851415150361275b576040517ee2a52200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000840361280d578085111561276f578094505b61278d61277b8361539e565b8460c00135888660e0013586896153b5565b935060007f8000000000000000000000000000000000000000000000000000000000000000198816905085816127c391906175f9565b8a866127cf91906175f9565b1115612807576040517ffb8ae12900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061290d565b61282b612819836153f0565b8460e00135878660c001358689615407565b945080851115612893578094506128566128448361539e565b8460c00135888660e0013586896153b5565b935087841115612892576040517f939c420400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b60007f8000000000000000000000000000000000000000000000000000000000000000198816905084816128c791906175f9565b89876128d391906175f9565b101561290b576040517f481ea39200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b600085148061291c5750600084145b15612953576040517ffba5a27600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84810390506001810160026000858152602001908152602001600020819055508d60600160208101906129869190616a34565b73ffffffffffffffffffffffffffffffffffffffff167fb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f0284836040516129cd92919061763b565b60405180910390a260146129e083615442565b905010612a925760003660006129fd6129f886615442565b61532a565b9250925092508273ffffffffffffffffffffffffffffffffffffffff166396a10e3387876060016020810190612a339190616a34565b338c8c8a89896040518963ffffffff1660e01b8152600401612a5c989796959493929190617888565b600060405180830381600087803b158015612a7657600080fd5b505af1158015612a8a573d6000803e3d6000fd5b505050505050505b612acb826020016020810190612aa89190616a34565b836060016020810190612abb9190616a34565b8888612ac687615459565b615470565b612b01576040517f70a03f4800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60148b8b905010612bee576000366000612b1b8e8e61532a565b92509250925060008373ffffffffffffffffffffffffffffffffffffffff1663ccee33d7338b8b87876040518663ffffffff1660e01b8152600401612b64959493929190617900565b6020604051808303816000875af1158015612b83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ba79190617963565b90508781118015612bc65750612bc4612bbf876153f0565b6154c9565b155b8015612be05750612bde612bd98761539e565b6154c9565b155b15612be9578097505b505050505b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16826040016020810190612c389190616a34565b73ffffffffffffffffffffffffffffffffffffffff16148015612c5b5750600034115b15612edf5783341015612c9a576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83341115612d495760003373ffffffffffffffffffffffffffffffffffffffff16853403604051612cca906179b6565b60006040518083038185875af1925050503d8060008114612d07576040519150601f19603f3d011682016040523d82523d6000602084013e612d0c565b606091505b5050905080612d47576040517fb12d13eb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b158015612db157600080fd5b505af1158015612dc5573d6000803e3d6000fd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600073ffffffffffffffffffffffffffffffffffffffff16846080016020810190612e319190616a34565b73ffffffffffffffffffffffffffffffffffffffff1614612e6457836080016020810190612e5f9190616a34565b612e78565b836060016020810190612e779190616a34565b5b866040518363ffffffff1660e01b8152600401612e969291906179cb565b6020604051808303816000875af1158015612eb5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ed99190617a20565b50612fe8565b60003414612f19576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612fb1826040016020810190612f2f9190616a34565b33600073ffffffffffffffffffffffffffffffffffffffff16856080016020810190612f5b9190616a34565b73ffffffffffffffffffffffffffffffffffffffff1614612f8e57846080016020810190612f899190616a34565b612fa2565b846060016020810190612fa19190616a34565b5b87612fac87615544565b615470565b612fe7576040517f478a520500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b6014612ff38361555b565b9050106130a557600036600061301061300b8661555b565b61532a565b9250925092508273ffffffffffffffffffffffffffffffffffffffff16633504ed62878760600160208101906130469190616a34565b338c8c8a89896040518963ffffffff1660e01b815260040161306f989796959493929190617888565b600060405180830381600087803b15801561308957600080fd5b505af115801561309d573d6000803e3d6000fd5b505050505050505b505099509950999650505050505050565b6130be614f59565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361312d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161312490617ad0565b60405180910390fd5b61313681614fd7565b50565b60006131498787878787876133e2565b90509695505050505050565b6131b3565b3d6000803e3d6000fd5b806131725761317161315a565b5b600160005114601f3d11163d1517806131af577ff27f64e40000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b604051601581017f0dfe1681d21220a7ddca3f43a9059cbb23b872dd0000000000000000000000008252602081600484335afa6131f3576131f261315a565b5b60208082016004808501335afa61320d5761320c61315a565b5b602060408201600460088501335afa6132295761322861315a565b5b6000806000881360018114613247576020840151925087915061324f565b835192508891505b507fff1f98431c8ad98523631ae4a59f267346ea31f984000000000000000000000084526060832083527fe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b54602084015273ffffffffffffffffffffffffffffffffffffffff6055852016338118156132eb577fb2c027220000000000000000000000000000000000000000000000000000000060005260046000fd5b6084357f0dfe1681d21220a7ddca3f43a9059cbb23b872dd00000000000000000000000086523081146001811461334a5781601488015233603488015283605488015261334560206000606460108b0160008a5af1613164565b61336e565b33601088015283603088015261336d602060006044600c8b0160008a5af1613164565b5b5050505050505050505050565b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16036133e0576040517f1b10b0f900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6000613549565b3d6000803e3d6000fd5b80613401576134006133e9565b5b600160005114601f3d11163d15178061343e577ff27f64e40000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b60007f0902f1ac000000000000000000000000000000000000000000000000000000008252604082600484875afa61347d5761347c6133e9565b5b60603d146134af577f85cd58dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b8151602083015186156134c55781819250809150505b8785029250633b9aca00820283018184020492507f022c0d9f000000000000000000000000000000000000000000000000000000008452861596508615830260048501528683026024850152886044850152608060648501526000608485015260008060a48660008a5af161353d5761353c6133e9565b5b50509695505050505050565b6dffffffffffffffffffffffffffff851115613589577fcf0b4d3a0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160c081016040528260051b8401843588600081146136385734156135d4577f1841b4e10000000000000000000000000000000000000000000000000000000060005260046000fd5b7f23b872dd00000000000000000000000000000000000000000000000000000000845233600485015273ffffffffffffffffffffffffffffffffffffffff821660248501528860448501526136336020600060648760008f5af16133f3565b613729565b348914613669577f1841b4e10000000000000000000000000000000000000000000000000000000060005260046000fd5b7fd0e30db00000000000000000000000000000000000000000000000000000000084526000806004868c73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25af16136b7576136b66133e9565b5b7fa9059cbb00000000000000000000000000000000000000000000000000000000845273ffffffffffffffffffffffffffffffffffffffff82166004850152886024850152600080604486600073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25af1613728576137276133e9565b5b5b50879350602086015b828110156137c65780356137b573ffffffffffffffffffffffffffffffffffffffff821677ffffffff0000000000000000000000000000000000000000851660a01c7f8000000000000000000000000000000000000000000000000000000000000000861673ffffffffffffffffffffffffffffffffffffffff87168a8a613442565b955080925050602081019050613732565b507f40000000000000000000000000000000000000000000000000000000000000008116600081146138c8576138553077ffffffff0000000000000000000000000000000000000000841660a01c7f8000000000000000000000000000000000000000000000000000000000000000851673ffffffffffffffffffffffffffffffffffffffff86168989613442565b94507f2e1a7d4d000000000000000000000000000000000000000000000000000000008452846004850152600080602486600073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25af16138ac576138ab6133e9565b5b600080600080888f5af16138c3576138c26133e9565b5b61392e565b61392b8b77ffffffff0000000000000000000000000000000000000000841660a01c7f8000000000000000000000000000000000000000000000000000000000000000851673ffffffffffffffffffffffffffffffffffffffff86168989613442565b94505b505050508381101561396c576040517ff32bec2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9695505050505050565b60006323b872dd60e01b905060006040518281528560048201528460248201528360448201526020600060648360008b5af1915081156139d5573d600081146139cb57600160005114601f3d111692506139d3565b6000883b1192505b505b5080613a0d576040517ff405907100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161480613a90575073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b9050919050565b600060e08383905003613abe57613ab78463d505accf60e01b8585615572565b9050613b17565b6101008383905003613ae457613add84638fcbaf0c60e01b8585615572565b9050613b16565b6040517f6827585700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b80613b2557613b246155c7565b5b50505050565b6000634b64e49260e01b90506040518181528560048201528284602483013784836024830101526000808460440183348b5af1613b6b573d6000823e3d81fd5b50505050505050565b6000613b7f83613a15565b15613ba3578173ffffffffffffffffffffffffffffffffffffffff16319050613c20565b8273ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff1660e01b8152600401613bdc9190616cb9565b602060405180830381865afa158015613bf9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c1d9190617963565b90505b92915050565b6000811115613d5257613c3883613a15565b15613d255780471015613c77576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168261138890604051613ca1906179b6565b600060405180830381858888f193505050503d8060008114613cdf576040519150601f19603f3d011682016040523d82523d6000602084013e613ce4565b606091505b5050905080613d1f576040517fb12d13eb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50613d51565b613d5082828573ffffffffffffffffffffffffffffffffffffffff166155d39092919063ffffffff16565b5b5b505050565b60008083839050905060008103613d9a576040517f67e7c0f600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600182039050869250600080341190506000807f2000000000000000000000000000000000000000000000000000000000000000888886818110613de357613de26176dd565b5b90506020020135161190508115613eaf57883414613e2d576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db08a6040518263ffffffff1660e01b81526004016000604051808303818588803b158015613e9557600080fd5b505af1158015613ea9573d6000803e3d6000fd5b50505050505b6001841115613f6857613eea3083613ec75733613ec9565b305b89896000818110613edd57613edc6176dd565b5b9050602002013588615621565b94506000600190505b83811015613f2f57613f2030308a8a85818110613f1357613f126176dd565b5b9050602002013589615621565b95508080600101915050613ef3565b50613f6181613f3e578a613f40565b305b30898987818110613f5457613f536176dd565b5b9050602002013588615621565b9450613fa9565b613fa681613f76578a613f78565b305b83613f835733613f85565b305b89896000818110613f9957613f986176dd565b5b9050602002013588615621565b94505b87851015613fe3576040517ff32bec2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b801561409e577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d866040518263ffffffff1660e01b815260040161404291906160f0565b600060405180830381600087803b15801561405c57600080fd5b505af1158015614070573d6000803e3d6000fd5b5050505061409d858b73ffffffffffffffffffffffffffffffffffffffff1661580390919063ffffffff16565b5b5050505095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614801561412757507f000000000000000000000000000000000000000000000000000000000000000046145b15614154577f000000000000000000000000000000000000000000000000000000000000000090506141c2565b6141bf7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006158f7565b90505b90565b6000366000848061012001906141db9190617aff565b9150915060007f0a244ca8a150ac294c14fcff9277051ced9a5b23e966a0ff0522e989da23116c90506040518284823782812061014082015261012087602083013781815261016081209450506142328585615931565b935050505092915050565b600080600061424c8585615972565b60e01c63ffffffff169050600061426f600487876159bc9290919263ffffffff16565b9050632cc2878d60e01b60e01c63ffffffff1682036142b2576001614293826115f7565b61429e5760006142a1565b60015b8060ff1690509350935050506145ac565b63bf15fcd860e01b60e01c63ffffffff1682101561443357636fe7b0ba60e01b60e01c63ffffffff1682101561437f57634f38e2b860e01b60e01c63ffffffff16820361433957600161431a8261431560648a8a615a0b9290919263ffffffff16565b611815565b614325576000614328565b60015b8060ff1690509350935050506145ac565b6363592c2b60e01b60e01c63ffffffff16820361437a57600161435b82611a66565b614366576000614369565b60015b8060ff1690509350935050506145ac565b61442e565b636fe7b0ba60e01b60e01c63ffffffff1682036143d65760016143b7826143b260648a8a615a0b9290919263ffffffff16565b611aa2565b6143c25760006143c5565b60015b8060ff1690509350935050506145ac565b637426114560e01b60e01c63ffffffff16820361442d57600161440e8261440960648a8a615a0b9290919263ffffffff16565b611c71565b61441957600061441c565b60015b8060ff1690509350935050506145ac565b5b61459a565b63ca4ece2260e01b60e01c63ffffffff168210156144ea5763bf15fcd860e01b60e01c63ffffffff16820361448e5760016144838261447e60648a8a615a0b9290919263ffffffff16565b61216c565b9350935050506145ac565b63bfa7514360e01b60e01c63ffffffff1682036144e55760016144c6826144c160648a8a615a0b9290919263ffffffff16565b6121c3565b6144d15760006144d4565b60015b8060ff1690509350935050506145ac565b614599565b63ca4ece2260e01b60e01c63ffffffff1682036145415760016145228261451d60648a8a615a0b9290919263ffffffff16565b6122a7565b61452d576000614530565b60015b8060ff1690509350935050506145ac565b63cf6fc6e360e01b60e01c63ffffffff1682036145985760016145798261457460248a8a6159bc9290919263ffffffff16565b6122d1565b614584576000614587565b60015b8060ff1690509350935050506145ac565b5b5b6145a53087876152f8565b9350935050505b9250929050565b6000807f74ab4f0cde46aaf927859983f7d04002116dd057d4c4941f6dbfb775c3e31f45905060006020850380518382526101008220925080825250506145fa8482615931565b9250505092915050565b600080631626ba7e60e01b905060405181815285600482015260406024820152836044820152838560648301376020600085606401838a5afa1561464f5760203d1460005183141692505b5050949350505050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603614697576000905061470a565b60408383905014806146ac5750604183839050145b80156146ed57508473ffffffffffffffffffffffffffffffffffffffff166146d5858585615a5d565b73ffffffffffffffffffffffffffffffffffffffff16145b156146fb576001905061470a565b61470785858585614604565b90505b949350505050565b600080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361477b576040517f692e45e000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600085606001519050600073ffffffffffffffffffffffffffffffffffffffff16866080015173ffffffffffffffffffffffffffffffffffffffff16141580156147f557503373ffffffffffffffffffffffffffffffffffffffff16866080015173ffffffffffffffffffffffffffffffffffffffff1614155b1561482c576040517fe8c6632100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008660000151905060006040826fffffffffffffffffffffffffffffffff16901c6fffffffffffffffffffffffffffffffff1690506000811415801561487257508042115b156148a9576040517fc56873ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6148b58383600061509b565b505060008660a00151905060008760c00151905060007f10000000000000000000000000000000000000000000000000000000000000007f20000000000000000000000000000000000000000000000000000000000000007f40000000000000000000000000000000000000000000000000000000000000007f800000000000000000000000000000000000000000000000000000000000000017171719881690506000810361496a57829550819450614a2f565b60007f80000000000000000000000000000000000000000000000000000000000000008916146149e357828111156149ce576040517faa34b69600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8095506149dc838388615b2e565b9450614a2e565b81811115614a1d576040517f7f902a9300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b809450614a2b838387615b67565b95505b5b5050506000831480614a415750600082145b15614a78576040517f07b6e79f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16866020015173ffffffffffffffffffffffffffffffffffffffff16148015614afa575060007f1000000000000000000000000000000000000000000000000000000000000000861614155b15614cd9577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd8230866040518463ffffffff1660e01b8152600401614b5c93929190617b62565b6020604051808303816000875af1158015614b7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b9f9190617a20565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d846040518263ffffffff1660e01b8152600401614bf991906160f0565b600060405180830381600087803b158015614c1357600080fd5b505af1158015614c27573d6000803e3d6000fd5b5050505060008473ffffffffffffffffffffffffffffffffffffffff168461138890604051614c55906179b6565b600060405180830381858888f193505050503d8060008114614c93576040519150601f19603f3d011682016040523d82523d6000602084013e614c98565b606091505b5050905080614cd3576040517fb12d13eb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50614d0b565b614d0a818585896020015173ffffffffffffffffffffffffffffffffffffffff16613976909392919063ffffffff16565b5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16866040015173ffffffffffffffffffffffffffffffffffffffff16148015614d6a5750600034115b15614ecd57813414614da8576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0836040518263ffffffff1660e01b81526004016000604051808303818588803b158015614e1057600080fd5b505af1158015614e24573d6000803e3d6000fd5b50505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb82846040518363ffffffff1660e01b8152600401614e849291906179cb565b6020604051808303816000875af1158015614ea3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ec79190617a20565b50614f39565b60003414614f07576040517f1841b4e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614f38338284896040015173ffffffffffffffffffffffffffffffffffffffff16613976909392919063ffffffff16565b5b50935093915050565b366000614f50836004615b89565b91509150915091565b614f61615bf3565b73ffffffffffffffffffffffffffffffffffffffff16614f7f611d9e565b73ffffffffffffffffffffffffffffffffffffffff1614614fd5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614fcc90617be5565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600060088367ffffffffffffffff16901c67ffffffffffffffff1690506000828460ff166001901b1790506000600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816000858152602001908152602001600020549050828382160361515b576040517ff71fbda200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8281178260008681526020019081526020016000208190555050505050505050565b600080631626ba7e60e01b905060405181815285600482015260406024820152604160448201528460648201527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff841660848201528360ff1c601b0160a48201536020600060a5838a5afa156151fa5760203d1460005183141692505b5050949350505050565b600080631626ba7e60e01b905060405181815285600482015260406024820152604060448201528460648201528360848201526020600060a4838a5afa156152535760203d1460005183141692505b5050949350505050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160361529b57600090506152f0565b8473ffffffffffffffffffffffffffffffffffffffff166152bd858585615bfb565b73ffffffffffffffffffffffffffffffffffffffff16036152e157600190506152f0565b6152ed85858585615204565b90505b949350505050565b60008060405183858237602060008583895afa925060203d148316925082156153215760005191505b50935093915050565b6000366000601485859050101561536d576040517fef356d7a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b843560601c92506014850191506014840390509250925092565b366000615395836005615b89565b91509150915091565b3660006153ac836003615b89565b91509150915091565b60008088889050036153d3576153cc868587615b2e565b90506153e5565b6153e288888888888888615c85565b90505b979650505050505050565b3660006153fe836002615b89565b91509150915091565b60008088889050036154255761541e848787615b67565b9050615437565b61543488888888888888615c85565b90505b979650505050505050565b366000615450836006615b89565b91509150915091565b366000615467836000615b89565b91509150915091565b6000806323b872dd60e01b90506040518181528760048201528660248201528560448201528385606483013760206000856064018360008d5af1600160005114601f3d11163d1517811693505050509695505050505050565b600060018383905014801561553c57507f78000000000000000000000000000000000000000000000000000000000000008383600081811061550e5761550d6176dd565b5b9050013560f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b905092915050565b366000615552836001615b89565b91509150915091565b366000615569836007615b89565b91509150915091565b6000816004016040518581528385600483013760206000838360008b5af1925082156155bd573d600081146155b357600160005114601f3d111693506155bb565b6000883b1193505b505b5050949350505050565b6040513d6000823e3d81fd5b6155e68363a9059cbb60e01b8484615e2a565b61561c576040517ffb7f507900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b60008060007f8000000000000000000000000000000000000000000000000000000000000000851614905080156157215760008473ffffffffffffffffffffffffffffffffffffffff1663128acb08888461567b88615e80565b6401000276a48b6040516020016156929190616cb9565b6040516020818303038152906040526040518663ffffffff1660e01b81526004016156c1959493929190617c23565b60408051808303816000875af11580156156df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906157039190617c92565b9150506157188161571390617cd2565b615eed565b925050506157fb565b60008473ffffffffffffffffffffffffffffffffffffffff1663128acb08888461574a88615e80565b73fffd8963efd1fc6a506488495d951d5263988d258b6040516020016157709190616cb9565b6040516020818303038152906040526040518663ffffffff1660e01b815260040161579f959493929190617c23565b60408051808303816000875af11580156157bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906157e19190617c92565b5090506157f6816157f190617cd2565b615eed565b925050505b949350505050565b80471015615846576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161583d90617d66565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405161586c906179b6565b60006040518083038185875af1925050503d80600081146158a9576040519150601f19603f3d011682016040523d82523d6000602084013e6158ae565b606091505b50509050806158f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016158e990617df8565b60405180910390fd5b505050565b60008383834630604051602001615912959493929190617e18565b6040516020818303038152906040528051906020012090509392505050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b600060048383905010156159b2576040517fef356d7a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8235905092915050565b6000602082018484905010156159fe576040517fef356d7a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8184013590509392505050565b36600082858590501015615a4b576040517fef356d7a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82850191508284039050935093915050565b60006040518260418114615a7c5760408114615a965760009150615ad9565b604085013560001a60208301526040856040840137615ad9565b60208501358060ff1c601b01602084015260208660408501377f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81166060840152505b508015615b26577f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a160608201511015615b2557848152600080526020600060808360015afa5060005191505b5b509392505050565b6000836001858585615b4091906175f9565b615b4a9190617664565b615b5491906175c5565b615b5e9190617e9a565b90509392505050565b6000828483615b7691906175f9565b615b809190617e9a565b90509392505050565b366000806005846007811115615ba257615ba1617ecb565b5b901b905084806101200190615bb79190617aff565b826020886101000135901b901c63ffffffff169083886101000135901c63ffffffff1692615be7939291906176a2565b92509250509250929050565b600033905090565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82167f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a1811015615c7d576040518581528360ff1c601b016020820152846040820152816060820152600080526020600060808360015afa506000519250505b509392505050565b600060018888905003615d1457615c9c88886154c9565b15615ce257858514615cda576040517f49986e7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b839050615e1f565b6040517fbec74c8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000366000615d238b8b61532a565b9250925092506000808473ffffffffffffffffffffffffffffffffffffffff1684848c8b8b604051602001615d5c959493929190617f3c565b604051602081830303815290604052604051615d789190617fb9565b600060405180830381855afa9150503d8060008114615db3576040519150601f19603f3d011682016040523d82523d6000602084013e615db8565b606091505b5091509150811580615dcc57506020815114155b15615e03576040517f110b8e7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80806020019051810190615e179190617963565b955050505050505b979650505050505050565b60006040518481528360048201528260248201526020600060448360008a5af191508115615e77573d60008114615e6d57600160005114601f3d11169250615e75565b6000873b1192505b505b50949350505050565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821115615ee5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615edc90618042565b60405180910390fd5b819050919050565b600080821215615f32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615f29906180ae565b60405180910390fd5b819050919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000615f7982615f4e565b9050919050565b6000615f8b82615f6e565b9050919050565b615f9b81615f80565b8114615fa657600080fd5b50565b600081359050615fb881615f92565b92915050565b6000819050919050565b615fd181615fbe565b8114615fdc57600080fd5b50565b600081359050615fee81615fc8565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261601957616018615ff4565b5b8235905067ffffffffffffffff81111561603657616035615ff9565b5b60208301915083602082028301111561605257616051615ffe565b5b9250929050565b60008060008060006080868803121561607557616074615f44565b5b600061608388828901615fa9565b955050602061609488828901615fdf565b94505060406160a588828901615fdf565b935050606086013567ffffffffffffffff8111156160c6576160c5615f49565b5b6160d288828901616003565b92509250509295509295909350565b6160ea81615fbe565b82525050565b600060208201905061610560008301846160e1565b92915050565b600061611682615f6e565b9050919050565b6161268161610b565b811461613157600080fd5b50565b6000813590506161438161611d565b92915050565b600061615482615f4e565b9050919050565b61616481616149565b811461616f57600080fd5b50565b6000813590506161818161615b565b92915050565b6000819050919050565b61619a81616187565b81146161a557600080fd5b50565b6000813590506161b781616191565b92915050565b60008060008060008060008060006101208a8c0312156161e0576161df615f44565b5b60006161ee8c828d01616134565b99505060206161ff8c828d01616172565b98505060406162108c828d01615fa9565b97505060606162218c828d01615fa9565b96505060806162328c828d01615fdf565b95505060a06162438c828d01615fdf565b94505060c06162548c828d01615fdf565b93505060e06162658c828d016161a8565b9250506101006162778c828d016161a8565b9150509295985092959850929598565b600061629282615f6e565b9050919050565b6162a281616287565b81146162ad57600080fd5b50565b6000813590506162bf81616299565b92915050565b600080fd5b600060e082840312156162e0576162df6162c5565b5b81905092915050565b60008083601f8401126162ff576162fe615ff4565b5b8235905067ffffffffffffffff81111561631c5761631b615ff9565b5b60208301915083600182028301111561633857616337615ffe565b5b9250929050565b600080600080600080610140878903121561635d5761635c615f44565b5b600061636b89828a016162b0565b965050602061637c89828a016162ca565b95505061010087013567ffffffffffffffff81111561639e5761639d615f49565b5b6163aa89828a016162e9565b945094505061012087013567ffffffffffffffff8111156163ce576163cd615f49565b5b6163da89828a016162e9565b92509250509295509295509295565b60006040820190506163fe60008301856160e1565b61640b60208301846160e1565b9392505050565b60008060008060008060008060c0898b03121561643257616431615f44565b5b60006164408b828c01616172565b98505060206164518b828c01615fa9565b97505060406164628b828c01615fdf565b96505060606164738b828c01615fdf565b955050608089013567ffffffffffffffff81111561649457616493615f49565b5b6164a08b828c01616003565b945094505060a089013567ffffffffffffffff8111156164c3576164c2615f49565b5b6164cf8b828c016162e9565b92509250509295985092959890939650565b6000602082840312156164f7576164f6615f44565b5b600061650584828501615fdf565b91505092915050565b60008115159050919050565b6165238161650e565b82525050565b600060208201905061653e600083018461651a565b92915050565b6000610140828403121561655b5761655a6162c5565b5b81905092915050565b60006020828403121561657a57616579615f44565b5b600082013567ffffffffffffffff81111561659857616597615f49565b5b6165a484828501616544565b91505092915050565b6165b681616187565b82525050565b60006040820190506165d160008301856160e1565b6165de60208301846165ad565b9392505050565b60006020820190506165fa60008301846165ad565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61664e82616605565b810181811067ffffffffffffffff8211171561666d5761666c616616565b5b80604052505050565b6000616680615f3a565b905061668c8282616645565b919050565b61669a81615f6e565b81146166a557600080fd5b50565b6000813590506166b781616691565b92915050565b600060e082840312156166d3576166d2616600565b5b6166dd60e0616676565b905060006166ed84828501615fdf565b6000830152506020616701848285016166a8565b6020830152506040616715848285016166a8565b6040830152506060616729848285016166a8565b606083015250608061673d848285016166a8565b60808301525060a061675184828501615fdf565b60a08301525060c061676584828501615fdf565b60c08301525092915050565b600080600080610120858703121561678c5761678b615f44565b5b600061679a878288016166bd565b94505060e085013567ffffffffffffffff8111156167bb576167ba615f49565b5b6167c7878288016162e9565b93509350506101006167db87828801615fdf565b91505092959194509250565b60006060820190506167fc60008301866160e1565b61680960208301856160e1565b61681660408301846165ad565b949350505050565b60008060006040848603121561683757616836615f44565b5b600061684586828701615fdf565b935050602084013567ffffffffffffffff81111561686657616865615f49565b5b616872868287016162e9565b92509250509250925092565b6000806040838503121561689557616894615f44565b5b60006168a3858286016166a8565b92505060206168b485828601615fdf565b9150509250929050565b600080600080600061014086880312156168db576168da615f44565b5b60006168e9888289016166bd565b95505060e086013567ffffffffffffffff81111561690a57616909615f49565b5b616916888289016162e9565b945094505061010061692a88828901615fdf565b92505061012061693c888289016166a8565b9150509295509295909350565b60008060008060008060008060c0898b03121561696957616968615f44565b5b600089013567ffffffffffffffff81111561698757616986615f49565b5b6169938b828c01616544565b985050602089013567ffffffffffffffff8111156169b4576169b3615f49565b5b6169c08b828c016162e9565b9750975050604089013567ffffffffffffffff8111156169e3576169e2615f49565b5b6169ef8b828c016162e9565b95509550506060616a028b828c01615fdf565b9350506080616a138b828c01615fdf565b92505060a0616a248b828c01615fdf565b9150509295985092959890939650565b600060208284031215616a4a57616a49615f44565b5b6000616a58848285016166a8565b91505092915050565b6000806000806000806000610160888a031215616a8157616a80615f44565b5b6000616a8f8a828b016166bd565b97505060e088013567ffffffffffffffff811115616ab057616aaf615f49565b5b616abc8a828b016162e9565b9650965050610100616ad08a828b01615fdf565b945050610120616ae28a828b016166a8565b93505061014088013567ffffffffffffffff811115616b0457616b03615f49565b5b616b108a828b016162e9565b925092505092959891949750929550565b600060ff82169050919050565b616b3781616b21565b8114616b4257600080fd5b50565b600081359050616b5481616b2e565b92915050565b600060208284031215616b7057616b6f615f44565b5b6000616b7e84828501616b45565b91505092915050565b60008060408385031215616b9e57616b9d615f44565b5b6000616bac85828601615fa9565b9250506020616bbd85828601615fdf565b9150509250929050565b600060208284031215616bdd57616bdc615f44565b5b6000616beb848285016161a8565b91505092915050565b600080600080600080600080610100898b031215616c1557616c14615f44565b5b6000616c238b828c01616134565b9850506020616c348b828c01615fa9565b9750506040616c458b828c01615fa9565b9650506060616c568b828c01615fdf565b9550506080616c678b828c01615fdf565b94505060a0616c788b828c01615fdf565b93505060c0616c898b828c016161a8565b92505060e0616c9a8b828c016161a8565b9150509295985092959890939650565b616cb381615f6e565b82525050565b6000602082019050616cce6000830184616caa565b92915050565b600067ffffffffffffffff821115616cef57616cee616616565b5b602082029050602081019050919050565b6000616d13616d0e84616cd4565b616676565b90508083825260208201905060208402830185811115616d3657616d35615ffe565b5b835b81811015616d5f5780616d4b88826161a8565b845260208401935050602081019050616d38565b5050509392505050565b600082601f830112616d7e57616d7d615ff4565b5b8135616d8e848260208601616d00565b91505092915050565b600060208284031215616dad57616dac615f44565b5b600082013567ffffffffffffffff811115616dcb57616dca615f49565b5b616dd784828501616d69565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b616e1581615fbe565b82525050565b6000616e278383616e0c565b60208301905092915050565b6000602082019050919050565b6000616e4b82616de0565b616e558185616deb565b9350616e6083616dfc565b8060005b83811015616e91578151616e788882616e1b565b9750616e8383616e33565b925050600181019050616e64565b5085935050505092915050565b60006020820190508181036000830152616eb88184616e40565b905092915050565b6000806000806101408587031215616edb57616eda615f44565b5b6000616ee9878288016166bd565b94505060e0616efa878288016161a8565b935050610100616f0c878288016161a8565b925050610120616f1e87828801615fdf565b91505092959194509250565b600080600080600060808688031215616f4657616f45615f44565b5b6000616f5488828901616172565b9550506020616f6588828901615fdf565b9450506040616f7688828901615fdf565b935050606086013567ffffffffffffffff811115616f9757616f96615f49565b5b616fa388828901616003565b92509250509295509295909350565b600080600060408486031215616fcb57616fca615f44565b5b6000616fd9868287016166a8565b935050602084013567ffffffffffffffff811115616ffa57616ff9615f49565b5b617006868287016162e9565b92509250509250925092565b6000806040838503121561702957617028615f44565b5b600061703785828601615fdf565b925050602061704885828601615fdf565b9150509250929050565b60008060008060008060008060008060006101408c8e03121561707857617077615f44565b5b60006170868e828f01616134565b9b505060206170978e828f01616172565b9a505060406170a88e828f01615fa9565b99505060606170b98e828f01615fa9565b98505060806170ca8e828f01615fdf565b97505060a06170db8e828f01615fdf565b96505060c06170ec8e828f01615fdf565b95505060e06170fd8e828f016161a8565b94505061010061710f8e828f016161a8565b9350506101208c013567ffffffffffffffff81111561713157617130615f49565b5b61713d8e828f016162e9565b92509250509295989b509295989b9093969950565b60008060008060008060008060008060006101008c8e03121561717857617177615f44565b5b60008c013567ffffffffffffffff81111561719657617195615f49565b5b6171a28e828f01616544565b9b505060208c013567ffffffffffffffff8111156171c3576171c2615f49565b5b6171cf8e828f016162e9565b9a509a505060408c013567ffffffffffffffff8111156171f2576171f1615f49565b5b6171fe8e828f016162e9565b985098505060606172118e828f01615fdf565b96505060806172228e828f01615fdf565b95505060a06172338e828f01615fdf565b94505060c06172448e828f016166a8565b93505060e08c013567ffffffffffffffff81111561726557617264615f49565b5b6172718e828f016162e9565b92509250509295989b509295989b9093969950565b600080600080606085870312156172a05761729f615f44565b5b60006172ae87828801615fdf565b94505060206172bf87828801615fdf565b935050604085013567ffffffffffffffff8111156172e0576172df615f49565b5b6172ec87828801616003565b925092505092959194509250565b600080600080600080600080600060e08a8c03121561731c5761731b615f44565b5b60008a013567ffffffffffffffff81111561733a57617339615f49565b5b6173468c828d01616544565b99505060208a013567ffffffffffffffff81111561736757617366615f49565b5b6173738c828d016162e9565b985098505060408a013567ffffffffffffffff81111561739657617395615f49565b5b6173a28c828d016162e9565b965096505060606173b58c828d01615fdf565b94505060806173c68c828d01615fdf565b93505060a06173d78c828d01615fdf565b92505060c06173e88c828d016166a8565b9150509295985092959850929598565b60008060008060008060a0878903121561741557617414615f44565b5b600061742389828a01616172565b965050602061743489828a01615fa9565b955050604061744589828a01615fdf565b945050606061745689828a01615fdf565b935050608087013567ffffffffffffffff81111561747757617476615f49565b5b61748389828a01616003565b92509250509295509295509295565b6000819050919050565b6174a581617492565b81146174b057600080fd5b50565b6000813590506174c28161749c565b92915050565b600080600080606085870312156174e2576174e1615f44565b5b60006174f0878288016174b3565b9450506020617501878288016174b3565b935050604085013567ffffffffffffffff81111561752257617521615f49565b5b61752e878288016162e9565b925092505092959194509250565b60006020828403121561755257617551615f44565b5b600061756084828501615fa9565b91505092915050565b60006020828403121561757f5761757e615f44565b5b600061758d84828501616172565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006175d082615fbe565b91506175db83615fbe565b92508282039050818111156175f3576175f2617596565b5b92915050565b600061760482615fbe565b915061760f83615fbe565b925082820261761d81615fbe565b9150828204841483151761763457617633617596565b5b5092915050565b600060408201905061765060008301856165ad565b61765d60208301846160e1565b9392505050565b600061766f82615fbe565b915061767a83615fbe565b925082820190508082111561769257617691617596565b5b92915050565b600080fd5b600080fd5b600080858511156176b6576176b5617698565b5b838611156176c7576176c661769d565b5b6001850283019150848603905094509492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061771782615fbe565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361774957617748617596565b5b600182019050919050565b600081905092915050565b82818337600083830152505050565b600061777a8385617754565b935061778783858461775f565b82840190509392505050565b60006177a082848661776e565b91508190509392505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156177e65780820151818401526020810190506177cb565b60008484015250505050565b60006177fd826177ac565b61780781856177b7565b93506178178185602086016177c8565b61782081616605565b840191505092915050565b6000604082019050617840600083018561651a565b818103602083015261785281846177f2565b90509392505050565b600061786783856177b7565b935061787483858461775f565b61787d83616605565b840190509392505050565b600060e08201905061789d600083018b6165ad565b6178aa602083018a616caa565b6178b76040830189616caa565b6178c460608301886160e1565b6178d160808301876160e1565b6178de60a08301866160e1565b81810360c08301526178f181848661785b565b90509998505050505050505050565b60006080820190506179156000830188616caa565b61792260208301876160e1565b61792f60408301866160e1565b818103606083015261794281848661785b565b90509695505050505050565b60008151905061795d81615fc8565b92915050565b60006020828403121561797957617978615f44565b5b60006179878482850161794e565b91505092915050565b50565b60006179a0600083617754565b91506179ab82617990565b600082019050919050565b60006179c182617993565b9150819050919050565b60006040820190506179e06000830185616caa565b6179ed60208301846160e1565b9392505050565b6179fd8161650e565b8114617a0857600080fd5b50565b600081519050617a1a816179f4565b92915050565b600060208284031215617a3657617a35615f44565b5b6000617a4484828501617a0b565b91505092915050565b600082825260208201905092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000617aba602683617a4d565b9150617ac582617a5e565b604082019050919050565b60006020820190508181036000830152617ae981617aad565b9050919050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112617b1c57617b1b617af0565b5b80840192508235915067ffffffffffffffff821115617b3e57617b3d617af5565b5b602083019250600182023603831315617b5a57617b59617afa565b5b509250929050565b6000606082019050617b776000830186616caa565b617b846020830185616caa565b617b9160408301846160e1565b949350505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000617bcf602083617a4d565b9150617bda82617b99565b602082019050919050565b60006020820190508181036000830152617bfe81617bc2565b9050919050565b617c0e81617492565b82525050565b617c1d81615f4e565b82525050565b600060a082019050617c386000830188616caa565b617c45602083018761651a565b617c526040830186617c05565b617c5f6060830185617c14565b8181036080830152617c7181846177f2565b90509695505050505050565b600081519050617c8c8161749c565b92915050565b60008060408385031215617ca957617ca8615f44565b5b6000617cb785828601617c7d565b9250506020617cc885828601617c7d565b9150509250929050565b6000617cdd82617492565b91507f80000000000000000000000000000000000000000000000000000000000000008203617d0f57617d0e617596565b5b816000039050919050565b7f416464726573733a20696e73756666696369656e742062616c616e6365000000600082015250565b6000617d50601d83617a4d565b9150617d5b82617d1a565b602082019050919050565b60006020820190508181036000830152617d7f81617d43565b9050919050565b7f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260008201527f6563697069656e74206d61792068617665207265766572746564000000000000602082015250565b6000617de2603a83617a4d565b9150617ded82617d86565b604082019050919050565b60006020820190508181036000830152617e1181617dd5565b9050919050565b600060a082019050617e2d60008301886165ad565b617e3a60208301876165ad565b617e4760408301866165ad565b617e5460608301856160e1565b617e616080830184616caa565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000617ea582615fbe565b9150617eb083615fbe565b925082617ec057617ebf617e6b565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000819050919050565b617f15617f1082615fbe565b617efa565b82525050565b6000819050919050565b617f36617f3182616187565b617f1b565b82525050565b6000617f4982878961776e565b9150617f558286617f04565b602082019150617f658285617f04565b602082019150617f758284617f25565b6020820191508190509695505050505050565b6000617f93826177ac565b617f9d8185617754565b9350617fad8185602086016177c8565b80840191505092915050565b6000617fc58284617f88565b915081905092915050565b7f53616665436173743a2076616c756520646f65736e27742066697420696e206160008201527f6e20696e74323536000000000000000000000000000000000000000000000000602082015250565b600061802c602883617a4d565b915061803782617fd0565b604082019050919050565b6000602082019050818103600083015261805b8161801f565b9050919050565b7f53616665436173743a2076616c7565206d75737420626520706f736974697665600082015250565b6000618098602083617a4d565b91506180a382618062565b602082019050919050565b600060208201905081810360008301526180c78161808b565b905091905056fea2646970667358221220f5d96d4a463aaf318675f69ec7293a852165489bc485f3cbc64173970f338bec64736f6c63430008110033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, weth common.Address) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, weth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

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

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_Contract *ContractCaller) And(opts *bind.CallOpts, offsets *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "and", offsets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_Contract *ContractSession) And(offsets *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.And(&_Contract.CallOpts, offsets, data)
}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_Contract *ContractCallerSession) And(offsets *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.And(&_Contract.CallOpts, offsets, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_Contract *ContractCaller) ArbitraryStaticCall(opts *bind.CallOpts, target common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "arbitraryStaticCall", target, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_Contract *ContractSession) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _Contract.Contract.ArbitraryStaticCall(&_Contract.CallOpts, target, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_Contract *ContractCallerSession) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _Contract.Contract.ArbitraryStaticCall(&_Contract.CallOpts, target, data)
}

// CheckPredicate is a free data retrieval call binding the contract method 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (_Contract *ContractCaller) CheckPredicate(opts *bind.CallOpts, order OrderLibOrder) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "checkPredicate", order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPredicate is a free data retrieval call binding the contract method 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (_Contract *ContractSession) CheckPredicate(order OrderLibOrder) (bool, error) {
	return _Contract.Contract.CheckPredicate(&_Contract.CallOpts, order)
}

// CheckPredicate is a free data retrieval call binding the contract method 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (_Contract *ContractCallerSession) CheckPredicate(order OrderLibOrder) (bool, error) {
	return _Contract.Contract.CheckPredicate(&_Contract.CallOpts, order)
}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_Contract *ContractCaller) Eq(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "eq", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_Contract *ContractSession) Eq(value *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.Eq(&_Contract.CallOpts, value, data)
}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_Contract *ContractCallerSession) Eq(value *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.Eq(&_Contract.CallOpts, value, data)
}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_Contract *ContractCaller) Gt(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "gt", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_Contract *ContractSession) Gt(value *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.Gt(&_Contract.CallOpts, value, data)
}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_Contract *ContractCallerSession) Gt(value *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.Gt(&_Contract.CallOpts, value, data)
}

// HashOrder is a free data retrieval call binding the contract method 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (_Contract *ContractCaller) HashOrder(opts *bind.CallOpts, order OrderLibOrder) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (_Contract *ContractSession) HashOrder(order OrderLibOrder) ([32]byte, error) {
	return _Contract.Contract.HashOrder(&_Contract.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (_Contract *ContractCallerSession) HashOrder(order OrderLibOrder) ([32]byte, error) {
	return _Contract.Contract.HashOrder(&_Contract.CallOpts, order)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Contract *ContractCaller) InvalidatorForOrderRFQ(opts *bind.CallOpts, maker common.Address, slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "invalidatorForOrderRFQ", maker, slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Contract *ContractSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Contract.Contract.InvalidatorForOrderRFQ(&_Contract.CallOpts, maker, slot)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Contract *ContractCallerSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Contract.Contract.InvalidatorForOrderRFQ(&_Contract.CallOpts, maker, slot)
}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_Contract *ContractCaller) Lt(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lt", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_Contract *ContractSession) Lt(value *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.Lt(&_Contract.CallOpts, value, data)
}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_Contract *ContractCallerSession) Lt(value *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.Lt(&_Contract.CallOpts, value, data)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Contract *ContractCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Contract *ContractSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonce(&_Contract.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Nonce(&_Contract.CallOpts, arg0)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_Contract *ContractCaller) NonceEquals(opts *bind.CallOpts, makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonceEquals", makerAddress, makerNonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_Contract *ContractSession) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _Contract.Contract.NonceEquals(&_Contract.CallOpts, makerAddress, makerNonce)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_Contract *ContractCallerSession) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _Contract.Contract.NonceEquals(&_Contract.CallOpts, makerAddress, makerNonce)
}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_Contract *ContractCaller) Or(opts *bind.CallOpts, offsets *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "or", offsets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_Contract *ContractSession) Or(offsets *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.Or(&_Contract.CallOpts, offsets, data)
}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_Contract *ContractCallerSession) Or(offsets *big.Int, data []byte) (bool, error) {
	return _Contract.Contract.Or(&_Contract.CallOpts, offsets, data)
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

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_Contract *ContractCaller) Remaining(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "remaining", orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_Contract *ContractSession) Remaining(orderHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.Remaining(&_Contract.CallOpts, orderHash)
}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_Contract *ContractCallerSession) Remaining(orderHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.Remaining(&_Contract.CallOpts, orderHash)
}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_Contract *ContractCaller) RemainingRaw(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "remainingRaw", orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_Contract *ContractSession) RemainingRaw(orderHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.RemainingRaw(&_Contract.CallOpts, orderHash)
}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_Contract *ContractCallerSession) RemainingRaw(orderHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.RemainingRaw(&_Contract.CallOpts, orderHash)
}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (_Contract *ContractCaller) RemainingsRaw(opts *bind.CallOpts, orderHashes [][32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "remainingsRaw", orderHashes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (_Contract *ContractSession) RemainingsRaw(orderHashes [][32]byte) ([]*big.Int, error) {
	return _Contract.Contract.RemainingsRaw(&_Contract.CallOpts, orderHashes)
}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (_Contract *ContractCallerSession) RemainingsRaw(orderHashes [][32]byte) ([]*big.Int, error) {
	return _Contract.Contract.RemainingsRaw(&_Contract.CallOpts, orderHashes)
}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_Contract *ContractCaller) TimestampBelow(opts *bind.CallOpts, time *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "timestampBelow", time)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_Contract *ContractSession) TimestampBelow(time *big.Int) (bool, error) {
	return _Contract.Contract.TimestampBelow(&_Contract.CallOpts, time)
}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_Contract *ContractCallerSession) TimestampBelow(time *big.Int) (bool, error) {
	return _Contract.Contract.TimestampBelow(&_Contract.CallOpts, time)
}

// TimestampBelowAndNonceEquals is a free data retrieval call binding the contract method 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (_Contract *ContractCaller) TimestampBelowAndNonceEquals(opts *bind.CallOpts, timeNonceAccount *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "timestampBelowAndNonceEquals", timeNonceAccount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TimestampBelowAndNonceEquals is a free data retrieval call binding the contract method 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (_Contract *ContractSession) TimestampBelowAndNonceEquals(timeNonceAccount *big.Int) (bool, error) {
	return _Contract.Contract.TimestampBelowAndNonceEquals(&_Contract.CallOpts, timeNonceAccount)
}

// TimestampBelowAndNonceEquals is a free data retrieval call binding the contract method 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (_Contract *ContractCallerSession) TimestampBelowAndNonceEquals(timeNonceAccount *big.Int) (bool, error) {
	return _Contract.Contract.TimestampBelowAndNonceEquals(&_Contract.CallOpts, timeNonceAccount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_Contract *ContractTransactor) AdvanceNonce(opts *bind.TransactOpts, amount uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "advanceNonce", amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_Contract *ContractSession) AdvanceNonce(amount uint8) (*types.Transaction, error) {
	return _Contract.Contract.AdvanceNonce(&_Contract.TransactOpts, amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_Contract *ContractTransactorSession) AdvanceNonce(amount uint8) (*types.Transaction, error) {
	return _Contract.Contract.AdvanceNonce(&_Contract.TransactOpts, amount)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (_Contract *ContractTransactor) CancelOrder(opts *bind.TransactOpts, order OrderLibOrder) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (_Contract *ContractSession) CancelOrder(order OrderLibOrder) (*types.Transaction, error) {
	return _Contract.Contract.CancelOrder(&_Contract.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (_Contract *ContractTransactorSession) CancelOrder(order OrderLibOrder) (*types.Transaction, error) {
	return _Contract.Contract.CancelOrder(&_Contract.TransactOpts, order)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Contract *ContractTransactor) CancelOrderRFQ(opts *bind.TransactOpts, orderInfo *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelOrderRFQ", orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Contract *ContractSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelOrderRFQ(&_Contract.TransactOpts, orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Contract *ContractTransactorSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelOrderRFQ(&_Contract.TransactOpts, orderInfo)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_Contract *ContractTransactor) CancelOrderRFQ0(opts *bind.TransactOpts, orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancelOrderRFQ0", orderInfo, additionalMask)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_Contract *ContractSession) CancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelOrderRFQ0(&_Contract.TransactOpts, orderInfo, additionalMask)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_Contract *ContractTransactorSession) CancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CancelOrderRFQ0(&_Contract.TransactOpts, orderInfo, additionalMask)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactor) ClipperSwap(opts *bind.TransactOpts, clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "clipperSwap", clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Contract *ContractSession) ClipperSwap(clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ClipperSwap(&_Contract.TransactOpts, clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactorSession) ClipperSwap(clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ClipperSwap(&_Contract.TransactOpts, clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactor) ClipperSwapTo(opts *bind.TransactOpts, clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "clipperSwapTo", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Contract *ContractSession) ClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ClipperSwapTo(&_Contract.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactorSession) ClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.ClipperSwapTo(&_Contract.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (_Contract *ContractTransactor) ClipperSwapToWithPermit(opts *bind.TransactOpts, clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "clipperSwapToWithPermit", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (_Contract *ContractSession) ClipperSwapToWithPermit(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.ClipperSwapToWithPermit(&_Contract.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (_Contract *ContractTransactorSession) ClipperSwapToWithPermit(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.ClipperSwapToWithPermit(&_Contract.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Contract *ContractTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Contract *ContractSession) Destroy() (*types.Transaction, error) {
	return _Contract.Contract.Destroy(&_Contract.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Contract *ContractTransactorSession) Destroy() (*types.Transaction, error) {
	return _Contract.Contract.Destroy(&_Contract.TransactOpts)
}

// FillOrder is a paid mutator transaction binding the contract method 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (_Contract *ContractTransactor) FillOrder(opts *bind.TransactOpts, order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fillOrder", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (_Contract *ContractSession) FillOrder(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.FillOrder(&_Contract.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (_Contract *ContractTransactorSession) FillOrder(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.FillOrder(&_Contract.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_Contract *ContractTransactor) FillOrderRFQ(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fillOrderRFQ", order, signature, flagsAndAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_Contract *ContractSession) FillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderRFQ(&_Contract.TransactOpts, order, signature, flagsAndAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_Contract *ContractTransactorSession) FillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderRFQ(&_Contract.TransactOpts, order, signature, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Contract *ContractTransactor) FillOrderRFQCompact(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fillOrderRFQCompact", order, r, vs, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Contract *ContractSession) FillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderRFQCompact(&_Contract.TransactOpts, order, r, vs, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Contract *ContractTransactorSession) FillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderRFQCompact(&_Contract.TransactOpts, order, r, vs, flagsAndAmount)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Contract *ContractTransactor) FillOrderRFQTo(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fillOrderRFQTo", order, signature, flagsAndAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Contract *ContractSession) FillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderRFQTo(&_Contract.TransactOpts, order, signature, flagsAndAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Contract *ContractTransactorSession) FillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderRFQTo(&_Contract.TransactOpts, order, signature, flagsAndAmount, target)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Contract *ContractTransactor) FillOrderRFQToWithPermit(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fillOrderRFQToWithPermit", order, signature, flagsAndAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Contract *ContractSession) FillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderRFQToWithPermit(&_Contract.TransactOpts, order, signature, flagsAndAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Contract *ContractTransactorSession) FillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderRFQToWithPermit(&_Contract.TransactOpts, order, signature, flagsAndAmount, target, permit)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (_Contract *ContractTransactor) FillOrderTo(opts *bind.TransactOpts, order_ OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fillOrderTo", order_, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (_Contract *ContractSession) FillOrderTo(order_ OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderTo(&_Contract.TransactOpts, order_, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (_Contract *ContractTransactorSession) FillOrderTo(order_ OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderTo(&_Contract.TransactOpts, order_, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Contract *ContractTransactor) FillOrderToWithPermit(opts *bind.TransactOpts, order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "fillOrderToWithPermit", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Contract *ContractSession) FillOrderToWithPermit(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderToWithPermit(&_Contract.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Contract *ContractTransactorSession) FillOrderToWithPermit(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.FillOrderToWithPermit(&_Contract.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_Contract *ContractTransactor) IncreaseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "increaseNonce")
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_Contract *ContractSession) IncreaseNonce() (*types.Transaction, error) {
	return _Contract.Contract.IncreaseNonce(&_Contract.TransactOpts)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_Contract *ContractTransactorSession) IncreaseNonce() (*types.Transaction, error) {
	return _Contract.Contract.IncreaseNonce(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Contract *ContractTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Contract *ContractSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RescueFunds(&_Contract.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Contract *ContractTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RescueFunds(&_Contract.TransactOpts, token, amount)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_Contract *ContractTransactor) Simulate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "simulate", target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_Contract *ContractSession) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.Simulate(&_Contract.TransactOpts, target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_Contract *ContractTransactorSession) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.Simulate(&_Contract.TransactOpts, target, data)
}

// Swap is a paid mutator transaction binding the contract method 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_Contract *ContractTransactor) Swap(opts *bind.TransactOpts, executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "swap", executor, desc, permit, data)
}

// Swap is a paid mutator transaction binding the contract method 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_Contract *ContractSession) Swap(executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.Swap(&_Contract.TransactOpts, executor, desc, permit, data)
}

// Swap is a paid mutator transaction binding the contract method 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_Contract *ContractTransactorSession) Swap(executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.Swap(&_Contract.TransactOpts, executor, desc, permit, data)
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

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactor) UniswapV3Swap(opts *bind.TransactOpts, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "uniswapV3Swap", amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractSession) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3Swap(&_Contract.TransactOpts, amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactorSession) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3Swap(&_Contract.TransactOpts, amount, minReturn, pools)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Contract *ContractTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Contract *ContractSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapCallback(&_Contract.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_Contract *ContractTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapCallback(&_Contract.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactor) UniswapV3SwapTo(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "uniswapV3SwapTo", recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractSession) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapTo(&_Contract.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactorSession) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapTo(&_Contract.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Contract *ContractTransactor) UniswapV3SwapToWithPermit(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "uniswapV3SwapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Contract *ContractSession) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapToWithPermit(&_Contract.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Contract *ContractTransactorSession) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapToWithPermit(&_Contract.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// Unoswap is a paid mutator transaction binding the contract method 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactor) Unoswap(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unoswap", srcToken, amount, minReturn, pools)
}

// Unoswap is a paid mutator transaction binding the contract method 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractSession) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Unoswap(&_Contract.TransactOpts, srcToken, amount, minReturn, pools)
}

// Unoswap is a paid mutator transaction binding the contract method 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactorSession) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Unoswap(&_Contract.TransactOpts, srcToken, amount, minReturn, pools)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactor) UnoswapTo(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unoswapTo", recipient, srcToken, amount, minReturn, pools)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractSession) UnoswapTo(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnoswapTo(&_Contract.TransactOpts, recipient, srcToken, amount, minReturn, pools)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_Contract *ContractTransactorSession) UnoswapTo(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnoswapTo(&_Contract.TransactOpts, recipient, srcToken, amount, minReturn, pools)
}

// UnoswapToWithPermit is a paid mutator transaction binding the contract method 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Contract *ContractTransactor) UnoswapToWithPermit(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unoswapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UnoswapToWithPermit is a paid mutator transaction binding the contract method 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Contract *ContractSession) UnoswapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.UnoswapToWithPermit(&_Contract.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// UnoswapToWithPermit is a paid mutator transaction binding the contract method 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_Contract *ContractTransactorSession) UnoswapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _Contract.Contract.UnoswapToWithPermit(&_Contract.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
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

// ContractNonceIncreasedIterator is returned from FilterNonceIncreased and is used to iterate over the raw logs and unpacked data for NonceIncreased events raised by the Contract contract.
type ContractNonceIncreasedIterator struct {
	Event *ContractNonceIncreased // Event containing the contract specifics and raw log

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
func (it *ContractNonceIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNonceIncreased)
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
		it.Event = new(ContractNonceIncreased)
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
func (it *ContractNonceIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNonceIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNonceIncreased represents a NonceIncreased event raised by the Contract contract.
type ContractNonceIncreased struct {
	Maker    common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncreased is a free log retrieval operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_Contract *ContractFilterer) FilterNonceIncreased(opts *bind.FilterOpts, maker []common.Address) (*ContractNonceIncreasedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return &ContractNonceIncreasedIterator{contract: _Contract.contract, event: "NonceIncreased", logs: logs, sub: sub}, nil
}

// WatchNonceIncreased is a free log subscription operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_Contract *ContractFilterer) WatchNonceIncreased(opts *bind.WatchOpts, sink chan<- *ContractNonceIncreased, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNonceIncreased)
				if err := _Contract.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
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

// ParseNonceIncreased is a log parse operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_Contract *ContractFilterer) ParseNonceIncreased(log types.Log) (*ContractNonceIncreased, error) {
	event := new(ContractNonceIncreased)
	if err := _Contract.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderCanceledIterator is returned from FilterOrderCanceled and is used to iterate over the raw logs and unpacked data for OrderCanceled events raised by the Contract contract.
type ContractOrderCanceledIterator struct {
	Event *ContractOrderCanceled // Event containing the contract specifics and raw log

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
func (it *ContractOrderCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderCanceled)
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
		it.Event = new(ContractOrderCanceled)
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
func (it *ContractOrderCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderCanceled represents a OrderCanceled event raised by the Contract contract.
type ContractOrderCanceled struct {
	Maker        common.Address
	OrderHash    [32]byte
	RemainingRaw *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderCanceled is a free log retrieval operation binding the contract event 0xcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (_Contract *ContractFilterer) FilterOrderCanceled(opts *bind.FilterOpts, maker []common.Address) (*ContractOrderCanceledIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OrderCanceled", makerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOrderCanceledIterator{contract: _Contract.contract, event: "OrderCanceled", logs: logs, sub: sub}, nil
}

// WatchOrderCanceled is a free log subscription operation binding the contract event 0xcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (_Contract *ContractFilterer) WatchOrderCanceled(opts *bind.WatchOpts, sink chan<- *ContractOrderCanceled, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OrderCanceled", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderCanceled)
				if err := _Contract.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
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

// ParseOrderCanceled is a log parse operation binding the contract event 0xcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (_Contract *ContractFilterer) ParseOrderCanceled(log types.Log) (*ContractOrderCanceled, error) {
	event := new(ContractOrderCanceled)
	if err := _Contract.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the Contract contract.
type ContractOrderFilledIterator struct {
	Event *ContractOrderFilled // Event containing the contract specifics and raw log

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
func (it *ContractOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderFilled)
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
		it.Event = new(ContractOrderFilled)
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
func (it *ContractOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderFilled represents a OrderFilled event raised by the Contract contract.
type ContractOrderFilled struct {
	Maker     common.Address
	OrderHash [32]byte
	Remaining *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_Contract *ContractFilterer) FilterOrderFilled(opts *bind.FilterOpts, maker []common.Address) (*ContractOrderFilledIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OrderFilled", makerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOrderFilledIterator{contract: _Contract.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_Contract *ContractFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *ContractOrderFilled, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OrderFilled", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderFilled)
				if err := _Contract.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_Contract *ContractFilterer) ParseOrderFilled(log types.Log) (*ContractOrderFilled, error) {
	event := new(ContractOrderFilled)
	if err := _Contract.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOrderFilledRFQIterator is returned from FilterOrderFilledRFQ and is used to iterate over the raw logs and unpacked data for OrderFilledRFQ events raised by the Contract contract.
type ContractOrderFilledRFQIterator struct {
	Event *ContractOrderFilledRFQ // Event containing the contract specifics and raw log

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
func (it *ContractOrderFilledRFQIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOrderFilledRFQ)
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
		it.Event = new(ContractOrderFilledRFQ)
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
func (it *ContractOrderFilledRFQIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOrderFilledRFQIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOrderFilledRFQ represents a OrderFilledRFQ event raised by the Contract contract.
type ContractOrderFilledRFQ struct {
	OrderHash    [32]byte
	MakingAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderFilledRFQ is a free log retrieval operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_Contract *ContractFilterer) FilterOrderFilledRFQ(opts *bind.FilterOpts) (*ContractOrderFilledRFQIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return &ContractOrderFilledRFQIterator{contract: _Contract.contract, event: "OrderFilledRFQ", logs: logs, sub: sub}, nil
}

// WatchOrderFilledRFQ is a free log subscription operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_Contract *ContractFilterer) WatchOrderFilledRFQ(opts *bind.WatchOpts, sink chan<- *ContractOrderFilledRFQ) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOrderFilledRFQ)
				if err := _Contract.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
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

// ParseOrderFilledRFQ is a log parse operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_Contract *ContractFilterer) ParseOrderFilledRFQ(log types.Log) (*ContractOrderFilledRFQ, error) {
	event := new(ContractOrderFilledRFQ)
	if err := _Contract.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
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
