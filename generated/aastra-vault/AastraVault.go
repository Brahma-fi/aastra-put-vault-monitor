// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AastraVault

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
)

// AastraVaultMetaData contains all meta data concerning the AastraVault contract.
var AastraVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_strategyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToVault0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToVault1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToStrategy0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToStrategy1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accruedProtocolFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedStrategyFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedStrategyFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"name\":\"burnAndCollect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"burned0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burned1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesToVault0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesToVault1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"collectProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"collectStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compoundFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emergencyBurnAndCollect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"freezeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"freezeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"getPositionAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_lowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_upperTick\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"swapEnabled\",\"type\":\"bool\"}],\"name\":\"mintOptimalLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStrategy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"poke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"}],\"name\":\"position\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_baseLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_baseUpper\",\"type\":\"int24\"}],\"name\":\"setBaseTicks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_limitLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitUpper\",\"type\":\"int24\"}],\"name\":\"setLimitTicks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"direction\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountInToSwap\",\"type\":\"uint256\"}],\"name\":\"swapTokensFromPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AastraVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use AastraVaultMetaData.ABI instead.
var AastraVaultABI = AastraVaultMetaData.ABI

// AastraVault is an auto generated Go binding around an Ethereum contract.
type AastraVault struct {
	AastraVaultCaller     // Read-only binding to the contract
	AastraVaultTransactor // Write-only binding to the contract
	AastraVaultFilterer   // Log filterer for contract events
}

// AastraVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type AastraVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AastraVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AastraVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AastraVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AastraVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AastraVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AastraVaultSession struct {
	Contract     *AastraVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AastraVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AastraVaultCallerSession struct {
	Contract *AastraVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AastraVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AastraVaultTransactorSession struct {
	Contract     *AastraVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AastraVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type AastraVaultRaw struct {
	Contract *AastraVault // Generic contract binding to access the raw methods on
}

// AastraVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AastraVaultCallerRaw struct {
	Contract *AastraVaultCaller // Generic read-only contract binding to access the raw methods on
}

// AastraVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AastraVaultTransactorRaw struct {
	Contract *AastraVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAastraVault creates a new instance of AastraVault, bound to a specific deployed contract.
func NewAastraVault(address common.Address, backend bind.ContractBackend) (*AastraVault, error) {
	contract, err := bindAastraVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AastraVault{AastraVaultCaller: AastraVaultCaller{contract: contract}, AastraVaultTransactor: AastraVaultTransactor{contract: contract}, AastraVaultFilterer: AastraVaultFilterer{contract: contract}}, nil
}

// NewAastraVaultCaller creates a new read-only instance of AastraVault, bound to a specific deployed contract.
func NewAastraVaultCaller(address common.Address, caller bind.ContractCaller) (*AastraVaultCaller, error) {
	contract, err := bindAastraVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AastraVaultCaller{contract: contract}, nil
}

// NewAastraVaultTransactor creates a new write-only instance of AastraVault, bound to a specific deployed contract.
func NewAastraVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*AastraVaultTransactor, error) {
	contract, err := bindAastraVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AastraVaultTransactor{contract: contract}, nil
}

// NewAastraVaultFilterer creates a new log filterer instance of AastraVault, bound to a specific deployed contract.
func NewAastraVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*AastraVaultFilterer, error) {
	contract, err := bindAastraVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AastraVaultFilterer{contract: contract}, nil
}

// bindAastraVault binds a generic wrapper to an already deployed contract.
func bindAastraVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AastraVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AastraVault *AastraVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AastraVault.Contract.AastraVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AastraVault *AastraVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AastraVault.Contract.AastraVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AastraVault *AastraVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AastraVault.Contract.AastraVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AastraVault *AastraVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AastraVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AastraVault *AastraVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AastraVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AastraVault *AastraVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AastraVault.Contract.contract.Transact(opts, method, params...)
}

// AccruedProtocolFees0 is a free data retrieval call binding the contract method 0xeae989a2.
//
// Solidity: function accruedProtocolFees0() view returns(uint256)
func (_AastraVault *AastraVaultCaller) AccruedProtocolFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "accruedProtocolFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedProtocolFees0 is a free data retrieval call binding the contract method 0xeae989a2.
//
// Solidity: function accruedProtocolFees0() view returns(uint256)
func (_AastraVault *AastraVaultSession) AccruedProtocolFees0() (*big.Int, error) {
	return _AastraVault.Contract.AccruedProtocolFees0(&_AastraVault.CallOpts)
}

// AccruedProtocolFees0 is a free data retrieval call binding the contract method 0xeae989a2.
//
// Solidity: function accruedProtocolFees0() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) AccruedProtocolFees0() (*big.Int, error) {
	return _AastraVault.Contract.AccruedProtocolFees0(&_AastraVault.CallOpts)
}

// AccruedProtocolFees1 is a free data retrieval call binding the contract method 0xa00fa77f.
//
// Solidity: function accruedProtocolFees1() view returns(uint256)
func (_AastraVault *AastraVaultCaller) AccruedProtocolFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "accruedProtocolFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedProtocolFees1 is a free data retrieval call binding the contract method 0xa00fa77f.
//
// Solidity: function accruedProtocolFees1() view returns(uint256)
func (_AastraVault *AastraVaultSession) AccruedProtocolFees1() (*big.Int, error) {
	return _AastraVault.Contract.AccruedProtocolFees1(&_AastraVault.CallOpts)
}

// AccruedProtocolFees1 is a free data retrieval call binding the contract method 0xa00fa77f.
//
// Solidity: function accruedProtocolFees1() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) AccruedProtocolFees1() (*big.Int, error) {
	return _AastraVault.Contract.AccruedProtocolFees1(&_AastraVault.CallOpts)
}

// AccruedStrategyFees0 is a free data retrieval call binding the contract method 0xae7fb76b.
//
// Solidity: function accruedStrategyFees0() view returns(uint256)
func (_AastraVault *AastraVaultCaller) AccruedStrategyFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "accruedStrategyFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedStrategyFees0 is a free data retrieval call binding the contract method 0xae7fb76b.
//
// Solidity: function accruedStrategyFees0() view returns(uint256)
func (_AastraVault *AastraVaultSession) AccruedStrategyFees0() (*big.Int, error) {
	return _AastraVault.Contract.AccruedStrategyFees0(&_AastraVault.CallOpts)
}

// AccruedStrategyFees0 is a free data retrieval call binding the contract method 0xae7fb76b.
//
// Solidity: function accruedStrategyFees0() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) AccruedStrategyFees0() (*big.Int, error) {
	return _AastraVault.Contract.AccruedStrategyFees0(&_AastraVault.CallOpts)
}

// AccruedStrategyFees1 is a free data retrieval call binding the contract method 0x14964478.
//
// Solidity: function accruedStrategyFees1() view returns(uint256)
func (_AastraVault *AastraVaultCaller) AccruedStrategyFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "accruedStrategyFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedStrategyFees1 is a free data retrieval call binding the contract method 0x14964478.
//
// Solidity: function accruedStrategyFees1() view returns(uint256)
func (_AastraVault *AastraVaultSession) AccruedStrategyFees1() (*big.Int, error) {
	return _AastraVault.Contract.AccruedStrategyFees1(&_AastraVault.CallOpts)
}

// AccruedStrategyFees1 is a free data retrieval call binding the contract method 0x14964478.
//
// Solidity: function accruedStrategyFees1() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) AccruedStrategyFees1() (*big.Int, error) {
	return _AastraVault.Contract.AccruedStrategyFees1(&_AastraVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AastraVault *AastraVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AastraVault *AastraVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AastraVault.Contract.Allowance(&_AastraVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AastraVault.Contract.Allowance(&_AastraVault.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AastraVault *AastraVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AastraVault *AastraVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AastraVault.Contract.BalanceOf(&_AastraVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AastraVault.Contract.BalanceOf(&_AastraVault.CallOpts, account)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_AastraVault *AastraVaultCaller) BaseLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "baseLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_AastraVault *AastraVaultSession) BaseLower() (*big.Int, error) {
	return _AastraVault.Contract.BaseLower(&_AastraVault.CallOpts)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_AastraVault *AastraVaultCallerSession) BaseLower() (*big.Int, error) {
	return _AastraVault.Contract.BaseLower(&_AastraVault.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_AastraVault *AastraVaultCaller) BaseUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "baseUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_AastraVault *AastraVaultSession) BaseUpper() (*big.Int, error) {
	return _AastraVault.Contract.BaseUpper(&_AastraVault.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_AastraVault *AastraVaultCallerSession) BaseUpper() (*big.Int, error) {
	return _AastraVault.Contract.BaseUpper(&_AastraVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AastraVault *AastraVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AastraVault *AastraVaultSession) Decimals() (uint8, error) {
	return _AastraVault.Contract.Decimals(&_AastraVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AastraVault *AastraVaultCallerSession) Decimals() (uint8, error) {
	return _AastraVault.Contract.Decimals(&_AastraVault.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AastraVault *AastraVaultCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AastraVault *AastraVaultSession) Factory() (common.Address, error) {
	return _AastraVault.Contract.Factory(&_AastraVault.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AastraVault *AastraVaultCallerSession) Factory() (common.Address, error) {
	return _AastraVault.Contract.Factory(&_AastraVault.CallOpts)
}

// GetBalance0 is a free data retrieval call binding the contract method 0x629d9405.
//
// Solidity: function getBalance0() view returns(uint256)
func (_AastraVault *AastraVaultCaller) GetBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "getBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance0 is a free data retrieval call binding the contract method 0x629d9405.
//
// Solidity: function getBalance0() view returns(uint256)
func (_AastraVault *AastraVaultSession) GetBalance0() (*big.Int, error) {
	return _AastraVault.Contract.GetBalance0(&_AastraVault.CallOpts)
}

// GetBalance0 is a free data retrieval call binding the contract method 0x629d9405.
//
// Solidity: function getBalance0() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) GetBalance0() (*big.Int, error) {
	return _AastraVault.Contract.GetBalance0(&_AastraVault.CallOpts)
}

// GetBalance1 is a free data retrieval call binding the contract method 0x41aec538.
//
// Solidity: function getBalance1() view returns(uint256)
func (_AastraVault *AastraVaultCaller) GetBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "getBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance1 is a free data retrieval call binding the contract method 0x41aec538.
//
// Solidity: function getBalance1() view returns(uint256)
func (_AastraVault *AastraVaultSession) GetBalance1() (*big.Int, error) {
	return _AastraVault.Contract.GetBalance1(&_AastraVault.CallOpts)
}

// GetBalance1 is a free data retrieval call binding the contract method 0x41aec538.
//
// Solidity: function getBalance1() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) GetBalance1() (*big.Int, error) {
	return _AastraVault.Contract.GetBalance1(&_AastraVault.CallOpts)
}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xa91ef6eb.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper) view returns(uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultCaller) GetPositionAmounts(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "getPositionAmounts", tickLower, tickUpper)

	outstruct := new(struct {
		Amount0 *big.Int
		Amount1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xa91ef6eb.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper) view returns(uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultSession) GetPositionAmounts(tickLower *big.Int, tickUpper *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _AastraVault.Contract.GetPositionAmounts(&_AastraVault.CallOpts, tickLower, tickUpper)
}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xa91ef6eb.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper) view returns(uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultCallerSession) GetPositionAmounts(tickLower *big.Int, tickUpper *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _AastraVault.Contract.GetPositionAmounts(&_AastraVault.CallOpts, tickLower, tickUpper)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_AastraVault *AastraVaultCaller) GetTotalAmounts(opts *bind.CallOpts) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "getTotalAmounts")

	outstruct := new(struct {
		Total0 *big.Int
		Total1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Total0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Total1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_AastraVault *AastraVaultSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _AastraVault.Contract.GetTotalAmounts(&_AastraVault.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_AastraVault *AastraVaultCallerSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _AastraVault.Contract.GetTotalAmounts(&_AastraVault.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_AastraVault *AastraVaultCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_AastraVault *AastraVaultSession) Governance() (common.Address, error) {
	return _AastraVault.Contract.Governance(&_AastraVault.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_AastraVault *AastraVaultCallerSession) Governance() (common.Address, error) {
	return _AastraVault.Contract.Governance(&_AastraVault.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_AastraVault *AastraVaultCaller) LimitLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "limitLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_AastraVault *AastraVaultSession) LimitLower() (*big.Int, error) {
	return _AastraVault.Contract.LimitLower(&_AastraVault.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_AastraVault *AastraVaultCallerSession) LimitLower() (*big.Int, error) {
	return _AastraVault.Contract.LimitLower(&_AastraVault.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_AastraVault *AastraVaultCaller) LimitUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "limitUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_AastraVault *AastraVaultSession) LimitUpper() (*big.Int, error) {
	return _AastraVault.Contract.LimitUpper(&_AastraVault.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_AastraVault *AastraVaultCallerSession) LimitUpper() (*big.Int, error) {
	return _AastraVault.Contract.LimitUpper(&_AastraVault.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_AastraVault *AastraVaultCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_AastraVault *AastraVaultSession) MaxTotalSupply() (*big.Int, error) {
	return _AastraVault.Contract.MaxTotalSupply(&_AastraVault.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _AastraVault.Contract.MaxTotalSupply(&_AastraVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AastraVault *AastraVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AastraVault *AastraVaultSession) Name() (string, error) {
	return _AastraVault.Contract.Name(&_AastraVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AastraVault *AastraVaultCallerSession) Name() (string, error) {
	return _AastraVault.Contract.Name(&_AastraVault.CallOpts)
}

// PauseStrategy is a free data retrieval call binding the contract method 0xd8c143f7.
//
// Solidity: function pauseStrategy() view returns(bool)
func (_AastraVault *AastraVaultCaller) PauseStrategy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "pauseStrategy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PauseStrategy is a free data retrieval call binding the contract method 0xd8c143f7.
//
// Solidity: function pauseStrategy() view returns(bool)
func (_AastraVault *AastraVaultSession) PauseStrategy() (bool, error) {
	return _AastraVault.Contract.PauseStrategy(&_AastraVault.CallOpts)
}

// PauseStrategy is a free data retrieval call binding the contract method 0xd8c143f7.
//
// Solidity: function pauseStrategy() view returns(bool)
func (_AastraVault *AastraVaultCallerSession) PauseStrategy() (bool, error) {
	return _AastraVault.Contract.PauseStrategy(&_AastraVault.CallOpts)
}

// PauseUser is a free data retrieval call binding the contract method 0xc0c3132c.
//
// Solidity: function pauseUser() view returns(bool)
func (_AastraVault *AastraVaultCaller) PauseUser(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "pauseUser")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PauseUser is a free data retrieval call binding the contract method 0xc0c3132c.
//
// Solidity: function pauseUser() view returns(bool)
func (_AastraVault *AastraVaultSession) PauseUser() (bool, error) {
	return _AastraVault.Contract.PauseUser(&_AastraVault.CallOpts)
}

// PauseUser is a free data retrieval call binding the contract method 0xc0c3132c.
//
// Solidity: function pauseUser() view returns(bool)
func (_AastraVault *AastraVaultCallerSession) PauseUser() (bool, error) {
	return _AastraVault.Contract.PauseUser(&_AastraVault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_AastraVault *AastraVaultCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_AastraVault *AastraVaultSession) Pool() (common.Address, error) {
	return _AastraVault.Contract.Pool(&_AastraVault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_AastraVault *AastraVaultCallerSession) Pool() (common.Address, error) {
	return _AastraVault.Contract.Pool(&_AastraVault.CallOpts)
}

// Position is a free data retrieval call binding the contract method 0x20ca5d32.
//
// Solidity: function position(int24 tickLower, int24 tickUpper) view returns(uint128, uint256, uint256, uint128, uint128)
func (_AastraVault *AastraVaultCaller) Position(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "position", tickLower, tickUpper)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// Position is a free data retrieval call binding the contract method 0x20ca5d32.
//
// Solidity: function position(int24 tickLower, int24 tickUpper) view returns(uint128, uint256, uint256, uint128, uint128)
func (_AastraVault *AastraVaultSession) Position(tickLower *big.Int, tickUpper *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AastraVault.Contract.Position(&_AastraVault.CallOpts, tickLower, tickUpper)
}

// Position is a free data retrieval call binding the contract method 0x20ca5d32.
//
// Solidity: function position(int24 tickLower, int24 tickUpper) view returns(uint128, uint256, uint256, uint128, uint128)
func (_AastraVault *AastraVaultCallerSession) Position(tickLower *big.Int, tickUpper *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AastraVault.Contract.Position(&_AastraVault.CallOpts, tickLower, tickUpper)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_AastraVault *AastraVaultCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_AastraVault *AastraVaultSession) ProtocolFee() (*big.Int, error) {
	return _AastraVault.Contract.ProtocolFee(&_AastraVault.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) ProtocolFee() (*big.Int, error) {
	return _AastraVault.Contract.ProtocolFee(&_AastraVault.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_AastraVault *AastraVaultCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_AastraVault *AastraVaultSession) Router() (common.Address, error) {
	return _AastraVault.Contract.Router(&_AastraVault.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_AastraVault *AastraVaultCallerSession) Router() (common.Address, error) {
	return _AastraVault.Contract.Router(&_AastraVault.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_AastraVault *AastraVaultCaller) Strategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "strategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_AastraVault *AastraVaultSession) Strategy() (common.Address, error) {
	return _AastraVault.Contract.Strategy(&_AastraVault.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_AastraVault *AastraVaultCallerSession) Strategy() (common.Address, error) {
	return _AastraVault.Contract.Strategy(&_AastraVault.CallOpts)
}

// StrategyFee is a free data retrieval call binding the contract method 0xacdb94ec.
//
// Solidity: function strategyFee() view returns(uint256)
func (_AastraVault *AastraVaultCaller) StrategyFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "strategyFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrategyFee is a free data retrieval call binding the contract method 0xacdb94ec.
//
// Solidity: function strategyFee() view returns(uint256)
func (_AastraVault *AastraVaultSession) StrategyFee() (*big.Int, error) {
	return _AastraVault.Contract.StrategyFee(&_AastraVault.CallOpts)
}

// StrategyFee is a free data retrieval call binding the contract method 0xacdb94ec.
//
// Solidity: function strategyFee() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) StrategyFee() (*big.Int, error) {
	return _AastraVault.Contract.StrategyFee(&_AastraVault.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_AastraVault *AastraVaultCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_AastraVault *AastraVaultSession) SwapRouter() (common.Address, error) {
	return _AastraVault.Contract.SwapRouter(&_AastraVault.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_AastraVault *AastraVaultCallerSession) SwapRouter() (common.Address, error) {
	return _AastraVault.Contract.SwapRouter(&_AastraVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AastraVault *AastraVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AastraVault *AastraVaultSession) Symbol() (string, error) {
	return _AastraVault.Contract.Symbol(&_AastraVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AastraVault *AastraVaultCallerSession) Symbol() (string, error) {
	return _AastraVault.Contract.Symbol(&_AastraVault.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AastraVault *AastraVaultCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AastraVault *AastraVaultSession) TickSpacing() (*big.Int, error) {
	return _AastraVault.Contract.TickSpacing(&_AastraVault.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AastraVault *AastraVaultCallerSession) TickSpacing() (*big.Int, error) {
	return _AastraVault.Contract.TickSpacing(&_AastraVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AastraVault *AastraVaultCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AastraVault *AastraVaultSession) Token0() (common.Address, error) {
	return _AastraVault.Contract.Token0(&_AastraVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AastraVault *AastraVaultCallerSession) Token0() (common.Address, error) {
	return _AastraVault.Contract.Token0(&_AastraVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AastraVault *AastraVaultCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AastraVault *AastraVaultSession) Token1() (common.Address, error) {
	return _AastraVault.Contract.Token1(&_AastraVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AastraVault *AastraVaultCallerSession) Token1() (common.Address, error) {
	return _AastraVault.Contract.Token1(&_AastraVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AastraVault *AastraVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AastraVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AastraVault *AastraVaultSession) TotalSupply() (*big.Int, error) {
	return _AastraVault.Contract.TotalSupply(&_AastraVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AastraVault *AastraVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _AastraVault.Contract.TotalSupply(&_AastraVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AastraVault *AastraVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AastraVault *AastraVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.Approve(&_AastraVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AastraVault *AastraVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.Approve(&_AastraVault.TransactOpts, spender, amount)
}

// BurnAndCollect is a paid mutator transaction binding the contract method 0xd5195d79.
//
// Solidity: function burnAndCollect(int24 tickLower, int24 tickUpper, uint128 liquidity) returns(uint256 burned0, uint256 burned1, uint256 feesToVault0, uint256 feesToVault1)
func (_AastraVault *AastraVaultTransactor) BurnAndCollect(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "burnAndCollect", tickLower, tickUpper, liquidity)
}

// BurnAndCollect is a paid mutator transaction binding the contract method 0xd5195d79.
//
// Solidity: function burnAndCollect(int24 tickLower, int24 tickUpper, uint128 liquidity) returns(uint256 burned0, uint256 burned1, uint256 feesToVault0, uint256 feesToVault1)
func (_AastraVault *AastraVaultSession) BurnAndCollect(tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.BurnAndCollect(&_AastraVault.TransactOpts, tickLower, tickUpper, liquidity)
}

// BurnAndCollect is a paid mutator transaction binding the contract method 0xd5195d79.
//
// Solidity: function burnAndCollect(int24 tickLower, int24 tickUpper, uint128 liquidity) returns(uint256 burned0, uint256 burned1, uint256 feesToVault0, uint256 feesToVault1)
func (_AastraVault *AastraVaultTransactorSession) BurnAndCollect(tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.BurnAndCollect(&_AastraVault.TransactOpts, tickLower, tickUpper, liquidity)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x0430c130.
//
// Solidity: function collectProtocol(uint256 amount0, uint256 amount1, address to) returns()
func (_AastraVault *AastraVaultTransactor) CollectProtocol(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "collectProtocol", amount0, amount1, to)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x0430c130.
//
// Solidity: function collectProtocol(uint256 amount0, uint256 amount1, address to) returns()
func (_AastraVault *AastraVaultSession) CollectProtocol(amount0 *big.Int, amount1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.CollectProtocol(&_AastraVault.TransactOpts, amount0, amount1, to)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x0430c130.
//
// Solidity: function collectProtocol(uint256 amount0, uint256 amount1, address to) returns()
func (_AastraVault *AastraVaultTransactorSession) CollectProtocol(amount0 *big.Int, amount1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.CollectProtocol(&_AastraVault.TransactOpts, amount0, amount1, to)
}

// CollectStrategy is a paid mutator transaction binding the contract method 0xd2f5ee4b.
//
// Solidity: function collectStrategy(uint256 amount0, uint256 amount1, address to) returns()
func (_AastraVault *AastraVaultTransactor) CollectStrategy(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "collectStrategy", amount0, amount1, to)
}

// CollectStrategy is a paid mutator transaction binding the contract method 0xd2f5ee4b.
//
// Solidity: function collectStrategy(uint256 amount0, uint256 amount1, address to) returns()
func (_AastraVault *AastraVaultSession) CollectStrategy(amount0 *big.Int, amount1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.CollectStrategy(&_AastraVault.TransactOpts, amount0, amount1, to)
}

// CollectStrategy is a paid mutator transaction binding the contract method 0xd2f5ee4b.
//
// Solidity: function collectStrategy(uint256 amount0, uint256 amount1, address to) returns()
func (_AastraVault *AastraVaultTransactorSession) CollectStrategy(amount0 *big.Int, amount1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.CollectStrategy(&_AastraVault.TransactOpts, amount0, amount1, to)
}

// CompoundFee is a paid mutator transaction binding the contract method 0x0a6c3bb2.
//
// Solidity: function compoundFee() returns()
func (_AastraVault *AastraVaultTransactor) CompoundFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "compoundFee")
}

// CompoundFee is a paid mutator transaction binding the contract method 0x0a6c3bb2.
//
// Solidity: function compoundFee() returns()
func (_AastraVault *AastraVaultSession) CompoundFee() (*types.Transaction, error) {
	return _AastraVault.Contract.CompoundFee(&_AastraVault.TransactOpts)
}

// CompoundFee is a paid mutator transaction binding the contract method 0x0a6c3bb2.
//
// Solidity: function compoundFee() returns()
func (_AastraVault *AastraVaultTransactorSession) CompoundFee() (*types.Transaction, error) {
	return _AastraVault.Contract.CompoundFee(&_AastraVault.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AastraVault *AastraVaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AastraVault *AastraVaultSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.DecreaseAllowance(&_AastraVault.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AastraVault *AastraVaultTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.DecreaseAllowance(&_AastraVault.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultTransactor) Deposit(opts *bind.TransactOpts, amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "deposit", amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultSession) Deposit(amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.Deposit(&_AastraVault.TransactOpts, amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultTransactorSession) Deposit(amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.Deposit(&_AastraVault.TransactOpts, amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// EmergencyBurnAndCollect is a paid mutator transaction binding the contract method 0x3869c916.
//
// Solidity: function emergencyBurnAndCollect(address to) returns()
func (_AastraVault *AastraVaultTransactor) EmergencyBurnAndCollect(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "emergencyBurnAndCollect", to)
}

// EmergencyBurnAndCollect is a paid mutator transaction binding the contract method 0x3869c916.
//
// Solidity: function emergencyBurnAndCollect(address to) returns()
func (_AastraVault *AastraVaultSession) EmergencyBurnAndCollect(to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.EmergencyBurnAndCollect(&_AastraVault.TransactOpts, to)
}

// EmergencyBurnAndCollect is a paid mutator transaction binding the contract method 0x3869c916.
//
// Solidity: function emergencyBurnAndCollect(address to) returns()
func (_AastraVault *AastraVaultTransactorSession) EmergencyBurnAndCollect(to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.EmergencyBurnAndCollect(&_AastraVault.TransactOpts, to)
}

// FreezeStrategy is a paid mutator transaction binding the contract method 0x6d5130d4.
//
// Solidity: function freezeStrategy(bool value) returns()
func (_AastraVault *AastraVaultTransactor) FreezeStrategy(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "freezeStrategy", value)
}

// FreezeStrategy is a paid mutator transaction binding the contract method 0x6d5130d4.
//
// Solidity: function freezeStrategy(bool value) returns()
func (_AastraVault *AastraVaultSession) FreezeStrategy(value bool) (*types.Transaction, error) {
	return _AastraVault.Contract.FreezeStrategy(&_AastraVault.TransactOpts, value)
}

// FreezeStrategy is a paid mutator transaction binding the contract method 0x6d5130d4.
//
// Solidity: function freezeStrategy(bool value) returns()
func (_AastraVault *AastraVaultTransactorSession) FreezeStrategy(value bool) (*types.Transaction, error) {
	return _AastraVault.Contract.FreezeStrategy(&_AastraVault.TransactOpts, value)
}

// FreezeUser is a paid mutator transaction binding the contract method 0x94cf2d7c.
//
// Solidity: function freezeUser(bool value) returns()
func (_AastraVault *AastraVaultTransactor) FreezeUser(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "freezeUser", value)
}

// FreezeUser is a paid mutator transaction binding the contract method 0x94cf2d7c.
//
// Solidity: function freezeUser(bool value) returns()
func (_AastraVault *AastraVaultSession) FreezeUser(value bool) (*types.Transaction, error) {
	return _AastraVault.Contract.FreezeUser(&_AastraVault.TransactOpts, value)
}

// FreezeUser is a paid mutator transaction binding the contract method 0x94cf2d7c.
//
// Solidity: function freezeUser(bool value) returns()
func (_AastraVault *AastraVaultTransactorSession) FreezeUser(value bool) (*types.Transaction, error) {
	return _AastraVault.Contract.FreezeUser(&_AastraVault.TransactOpts, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AastraVault *AastraVaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AastraVault *AastraVaultSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.IncreaseAllowance(&_AastraVault.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AastraVault *AastraVaultTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.IncreaseAllowance(&_AastraVault.TransactOpts, spender, addedValue)
}

// MintOptimalLiquidity is a paid mutator transaction binding the contract method 0x6b8d1d49.
//
// Solidity: function mintOptimalLiquidity(int24 _lowerTick, int24 _upperTick, uint256 amount0, uint256 amount1, bool swapEnabled) returns()
func (_AastraVault *AastraVaultTransactor) MintOptimalLiquidity(opts *bind.TransactOpts, _lowerTick *big.Int, _upperTick *big.Int, amount0 *big.Int, amount1 *big.Int, swapEnabled bool) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "mintOptimalLiquidity", _lowerTick, _upperTick, amount0, amount1, swapEnabled)
}

// MintOptimalLiquidity is a paid mutator transaction binding the contract method 0x6b8d1d49.
//
// Solidity: function mintOptimalLiquidity(int24 _lowerTick, int24 _upperTick, uint256 amount0, uint256 amount1, bool swapEnabled) returns()
func (_AastraVault *AastraVaultSession) MintOptimalLiquidity(_lowerTick *big.Int, _upperTick *big.Int, amount0 *big.Int, amount1 *big.Int, swapEnabled bool) (*types.Transaction, error) {
	return _AastraVault.Contract.MintOptimalLiquidity(&_AastraVault.TransactOpts, _lowerTick, _upperTick, amount0, amount1, swapEnabled)
}

// MintOptimalLiquidity is a paid mutator transaction binding the contract method 0x6b8d1d49.
//
// Solidity: function mintOptimalLiquidity(int24 _lowerTick, int24 _upperTick, uint256 amount0, uint256 amount1, bool swapEnabled) returns()
func (_AastraVault *AastraVaultTransactorSession) MintOptimalLiquidity(_lowerTick *big.Int, _upperTick *big.Int, amount0 *big.Int, amount1 *big.Int, swapEnabled bool) (*types.Transaction, error) {
	return _AastraVault.Contract.MintOptimalLiquidity(&_AastraVault.TransactOpts, _lowerTick, _upperTick, amount0, amount1, swapEnabled)
}

// Poke is a paid mutator transaction binding the contract method 0xfd7fd491.
//
// Solidity: function poke(int24 tickLower, int24 tickUpper) returns()
func (_AastraVault *AastraVaultTransactor) Poke(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "poke", tickLower, tickUpper)
}

// Poke is a paid mutator transaction binding the contract method 0xfd7fd491.
//
// Solidity: function poke(int24 tickLower, int24 tickUpper) returns()
func (_AastraVault *AastraVaultSession) Poke(tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.Poke(&_AastraVault.TransactOpts, tickLower, tickUpper)
}

// Poke is a paid mutator transaction binding the contract method 0xfd7fd491.
//
// Solidity: function poke(int24 tickLower, int24 tickUpper) returns()
func (_AastraVault *AastraVaultTransactorSession) Poke(tickLower *big.Int, tickUpper *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.Poke(&_AastraVault.TransactOpts, tickLower, tickUpper)
}

// SetBaseTicks is a paid mutator transaction binding the contract method 0x5542c87f.
//
// Solidity: function setBaseTicks(int24 _baseLower, int24 _baseUpper) returns()
func (_AastraVault *AastraVaultTransactor) SetBaseTicks(opts *bind.TransactOpts, _baseLower *big.Int, _baseUpper *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "setBaseTicks", _baseLower, _baseUpper)
}

// SetBaseTicks is a paid mutator transaction binding the contract method 0x5542c87f.
//
// Solidity: function setBaseTicks(int24 _baseLower, int24 _baseUpper) returns()
func (_AastraVault *AastraVaultSession) SetBaseTicks(_baseLower *big.Int, _baseUpper *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.SetBaseTicks(&_AastraVault.TransactOpts, _baseLower, _baseUpper)
}

// SetBaseTicks is a paid mutator transaction binding the contract method 0x5542c87f.
//
// Solidity: function setBaseTicks(int24 _baseLower, int24 _baseUpper) returns()
func (_AastraVault *AastraVaultTransactorSession) SetBaseTicks(_baseLower *big.Int, _baseUpper *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.SetBaseTicks(&_AastraVault.TransactOpts, _baseLower, _baseUpper)
}

// SetLimitTicks is a paid mutator transaction binding the contract method 0x64d54930.
//
// Solidity: function setLimitTicks(int24 _limitLower, int24 _limitUpper) returns()
func (_AastraVault *AastraVaultTransactor) SetLimitTicks(opts *bind.TransactOpts, _limitLower *big.Int, _limitUpper *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "setLimitTicks", _limitLower, _limitUpper)
}

// SetLimitTicks is a paid mutator transaction binding the contract method 0x64d54930.
//
// Solidity: function setLimitTicks(int24 _limitLower, int24 _limitUpper) returns()
func (_AastraVault *AastraVaultSession) SetLimitTicks(_limitLower *big.Int, _limitUpper *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.SetLimitTicks(&_AastraVault.TransactOpts, _limitLower, _limitUpper)
}

// SetLimitTicks is a paid mutator transaction binding the contract method 0x64d54930.
//
// Solidity: function setLimitTicks(int24 _limitLower, int24 _limitUpper) returns()
func (_AastraVault *AastraVaultTransactorSession) SetLimitTicks(_limitLower *big.Int, _limitUpper *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.SetLimitTicks(&_AastraVault.TransactOpts, _limitLower, _limitUpper)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_AastraVault *AastraVaultTransactor) SetMaxTotalSupply(opts *bind.TransactOpts, _maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "setMaxTotalSupply", _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_AastraVault *AastraVaultSession) SetMaxTotalSupply(_maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.SetMaxTotalSupply(&_AastraVault.TransactOpts, _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_AastraVault *AastraVaultTransactorSession) SetMaxTotalSupply(_maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.SetMaxTotalSupply(&_AastraVault.TransactOpts, _maxTotalSupply)
}

// SwapTokensFromPool is a paid mutator transaction binding the contract method 0xc3e895dc.
//
// Solidity: function swapTokensFromPool(bool direction, uint256 amountInToSwap) returns(uint256 amountOut)
func (_AastraVault *AastraVaultTransactor) SwapTokensFromPool(opts *bind.TransactOpts, direction bool, amountInToSwap *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "swapTokensFromPool", direction, amountInToSwap)
}

// SwapTokensFromPool is a paid mutator transaction binding the contract method 0xc3e895dc.
//
// Solidity: function swapTokensFromPool(bool direction, uint256 amountInToSwap) returns(uint256 amountOut)
func (_AastraVault *AastraVaultSession) SwapTokensFromPool(direction bool, amountInToSwap *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.SwapTokensFromPool(&_AastraVault.TransactOpts, direction, amountInToSwap)
}

// SwapTokensFromPool is a paid mutator transaction binding the contract method 0xc3e895dc.
//
// Solidity: function swapTokensFromPool(bool direction, uint256 amountInToSwap) returns(uint256 amountOut)
func (_AastraVault *AastraVaultTransactorSession) SwapTokensFromPool(direction bool, amountInToSwap *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.SwapTokensFromPool(&_AastraVault.TransactOpts, direction, amountInToSwap)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AastraVault *AastraVaultTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AastraVault *AastraVaultSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.Transfer(&_AastraVault.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AastraVault *AastraVaultTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.Transfer(&_AastraVault.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AastraVault *AastraVaultTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AastraVault *AastraVaultSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.TransferFrom(&_AastraVault.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AastraVault *AastraVaultTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AastraVault.Contract.TransferFrom(&_AastraVault.TransactOpts, sender, recipient, amount)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_AastraVault *AastraVaultTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "uniswapV3MintCallback", amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_AastraVault *AastraVaultSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AastraVault.Contract.UniswapV3MintCallback(&_AastraVault.TransactOpts, amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_AastraVault *AastraVaultTransactorSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AastraVault.Contract.UniswapV3MintCallback(&_AastraVault.TransactOpts, amount0, amount1, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_AastraVault *AastraVaultTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_AastraVault *AastraVaultSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _AastraVault.Contract.UniswapV3SwapCallback(&_AastraVault.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_AastraVault *AastraVaultTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _AastraVault.Contract.UniswapV3SwapCallback(&_AastraVault.TransactOpts, amount0Delta, amount1Delta, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.contract.Transact(opts, "withdraw", shares, amount0Min, amount1Min, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultSession) Withdraw(shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.Withdraw(&_AastraVault.TransactOpts, shares, amount0Min, amount1Min, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultTransactorSession) Withdraw(shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AastraVault.Contract.Withdraw(&_AastraVault.TransactOpts, shares, amount0Min, amount1Min, to)
}

// AastraVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AastraVault contract.
type AastraVaultApprovalIterator struct {
	Event *AastraVaultApproval // Event containing the contract specifics and raw log

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
func (it *AastraVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AastraVaultApproval)
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
		it.Event = new(AastraVaultApproval)
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
func (it *AastraVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AastraVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AastraVaultApproval represents a Approval event raised by the AastraVault contract.
type AastraVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AastraVault *AastraVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AastraVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AastraVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AastraVaultApprovalIterator{contract: _AastraVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AastraVault *AastraVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AastraVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AastraVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AastraVaultApproval)
				if err := _AastraVault.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AastraVault *AastraVaultFilterer) ParseApproval(log types.Log) (*AastraVaultApproval, error) {
	event := new(AastraVaultApproval)
	if err := _AastraVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AastraVaultCollectFeesIterator is returned from FilterCollectFees and is used to iterate over the raw logs and unpacked data for CollectFees events raised by the AastraVault contract.
type AastraVaultCollectFeesIterator struct {
	Event *AastraVaultCollectFees // Event containing the contract specifics and raw log

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
func (it *AastraVaultCollectFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AastraVaultCollectFees)
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
		it.Event = new(AastraVaultCollectFees)
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
func (it *AastraVaultCollectFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AastraVaultCollectFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AastraVaultCollectFees represents a CollectFees event raised by the AastraVault contract.
type AastraVaultCollectFees struct {
	FeesToVault0    *big.Int
	FeesToVault1    *big.Int
	FeesToStrategy0 *big.Int
	FeesToStrategy1 *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0x1ac56d7e866e3f5ea9aa92aa11758ead39a0a5f013f3fefb0f47cb9d008edd27.
//
// Solidity: event CollectFees(uint256 feesToVault0, uint256 feesToVault1, uint256 feesToStrategy0, uint256 feesToStrategy1)
func (_AastraVault *AastraVaultFilterer) FilterCollectFees(opts *bind.FilterOpts) (*AastraVaultCollectFeesIterator, error) {

	logs, sub, err := _AastraVault.contract.FilterLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return &AastraVaultCollectFeesIterator{contract: _AastraVault.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0x1ac56d7e866e3f5ea9aa92aa11758ead39a0a5f013f3fefb0f47cb9d008edd27.
//
// Solidity: event CollectFees(uint256 feesToVault0, uint256 feesToVault1, uint256 feesToStrategy0, uint256 feesToStrategy1)
func (_AastraVault *AastraVaultFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *AastraVaultCollectFees) (event.Subscription, error) {

	logs, sub, err := _AastraVault.contract.WatchLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AastraVaultCollectFees)
				if err := _AastraVault.contract.UnpackLog(event, "CollectFees", log); err != nil {
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

// ParseCollectFees is a log parse operation binding the contract event 0x1ac56d7e866e3f5ea9aa92aa11758ead39a0a5f013f3fefb0f47cb9d008edd27.
//
// Solidity: event CollectFees(uint256 feesToVault0, uint256 feesToVault1, uint256 feesToStrategy0, uint256 feesToStrategy1)
func (_AastraVault *AastraVaultFilterer) ParseCollectFees(log types.Log) (*AastraVaultCollectFees, error) {
	event := new(AastraVaultCollectFees)
	if err := _AastraVault.contract.UnpackLog(event, "CollectFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AastraVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the AastraVault contract.
type AastraVaultDepositIterator struct {
	Event *AastraVaultDeposit // Event containing the contract specifics and raw log

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
func (it *AastraVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AastraVaultDeposit)
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
		it.Event = new(AastraVaultDeposit)
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
func (it *AastraVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AastraVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AastraVaultDeposit represents a Deposit event raised by the AastraVault contract.
type AastraVaultDeposit struct {
	Sender  common.Address
	To      common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*AastraVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AastraVault.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AastraVaultDepositIterator{contract: _AastraVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AastraVaultDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AastraVault.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AastraVaultDeposit)
				if err := _AastraVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultFilterer) ParseDeposit(log types.Log) (*AastraVaultDeposit, error) {
	event := new(AastraVaultDeposit)
	if err := _AastraVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AastraVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AastraVault contract.
type AastraVaultTransferIterator struct {
	Event *AastraVaultTransfer // Event containing the contract specifics and raw log

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
func (it *AastraVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AastraVaultTransfer)
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
		it.Event = new(AastraVaultTransfer)
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
func (it *AastraVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AastraVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AastraVaultTransfer represents a Transfer event raised by the AastraVault contract.
type AastraVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AastraVault *AastraVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AastraVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AastraVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AastraVaultTransferIterator{contract: _AastraVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AastraVault *AastraVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AastraVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AastraVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AastraVaultTransfer)
				if err := _AastraVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AastraVault *AastraVaultFilterer) ParseTransfer(log types.Log) (*AastraVaultTransfer, error) {
	event := new(AastraVaultTransfer)
	if err := _AastraVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AastraVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AastraVault contract.
type AastraVaultWithdrawIterator struct {
	Event *AastraVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *AastraVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AastraVaultWithdraw)
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
		it.Event = new(AastraVaultWithdraw)
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
func (it *AastraVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AastraVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AastraVaultWithdraw represents a Withdraw event raised by the AastraVault contract.
type AastraVaultWithdraw struct {
	Sender  common.Address
	To      common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*AastraVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AastraVault.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AastraVaultWithdrawIterator{contract: _AastraVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AastraVaultWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AastraVault.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AastraVaultWithdraw)
				if err := _AastraVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AastraVault *AastraVaultFilterer) ParseWithdraw(log types.Log) (*AastraVaultWithdraw, error) {
	event := new(AastraVaultWithdraw)
	if err := _AastraVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
