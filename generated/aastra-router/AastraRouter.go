// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AastraRouter

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

// AastraRouterMetaData contains all meta data concerning the AastraRouter contract.
var AastraRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"compoundFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"getBaseAmounts\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"getBaseTicks\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"getRangeAmounts\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"getRangeTicks\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_baseLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_baseUpper\",\"type\":\"int24\"},{\"internalType\":\"uint8\",\"name\":\"_percentage\",\"type\":\"uint8\"}],\"name\":\"newBaseLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_limitLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitUpper\",\"type\":\"int24\"},{\"internalType\":\"uint8\",\"name\":\"_percentage\",\"type\":\"uint8\"}],\"name\":\"newLimitLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AastraRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use AastraRouterMetaData.ABI instead.
var AastraRouterABI = AastraRouterMetaData.ABI

// AastraRouter is an auto generated Go binding around an Ethereum contract.
type AastraRouter struct {
	AastraRouterCaller     // Read-only binding to the contract
	AastraRouterTransactor // Write-only binding to the contract
	AastraRouterFilterer   // Log filterer for contract events
}

// AastraRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AastraRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AastraRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AastraRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AastraRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AastraRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AastraRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AastraRouterSession struct {
	Contract     *AastraRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AastraRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AastraRouterCallerSession struct {
	Contract *AastraRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AastraRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AastraRouterTransactorSession struct {
	Contract     *AastraRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AastraRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AastraRouterRaw struct {
	Contract *AastraRouter // Generic contract binding to access the raw methods on
}

// AastraRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AastraRouterCallerRaw struct {
	Contract *AastraRouterCaller // Generic read-only contract binding to access the raw methods on
}

// AastraRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AastraRouterTransactorRaw struct {
	Contract *AastraRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAastraRouter creates a new instance of AastraRouter, bound to a specific deployed contract.
func NewAastraRouter(address common.Address, backend bind.ContractBackend) (*AastraRouter, error) {
	contract, err := bindAastraRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AastraRouter{AastraRouterCaller: AastraRouterCaller{contract: contract}, AastraRouterTransactor: AastraRouterTransactor{contract: contract}, AastraRouterFilterer: AastraRouterFilterer{contract: contract}}, nil
}

// NewAastraRouterCaller creates a new read-only instance of AastraRouter, bound to a specific deployed contract.
func NewAastraRouterCaller(address common.Address, caller bind.ContractCaller) (*AastraRouterCaller, error) {
	contract, err := bindAastraRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AastraRouterCaller{contract: contract}, nil
}

// NewAastraRouterTransactor creates a new write-only instance of AastraRouter, bound to a specific deployed contract.
func NewAastraRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*AastraRouterTransactor, error) {
	contract, err := bindAastraRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AastraRouterTransactor{contract: contract}, nil
}

// NewAastraRouterFilterer creates a new log filterer instance of AastraRouter, bound to a specific deployed contract.
func NewAastraRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*AastraRouterFilterer, error) {
	contract, err := bindAastraRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AastraRouterFilterer{contract: contract}, nil
}

// bindAastraRouter binds a generic wrapper to an already deployed contract.
func bindAastraRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AastraRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AastraRouter *AastraRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AastraRouter.Contract.AastraRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AastraRouter *AastraRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AastraRouter.Contract.AastraRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AastraRouter *AastraRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AastraRouter.Contract.AastraRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AastraRouter *AastraRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AastraRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AastraRouter *AastraRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AastraRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AastraRouter *AastraRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AastraRouter.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AastraRouter *AastraRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AastraRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AastraRouter *AastraRouterSession) Factory() (common.Address, error) {
	return _AastraRouter.Contract.Factory(&_AastraRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AastraRouter *AastraRouterCallerSession) Factory() (common.Address, error) {
	return _AastraRouter.Contract.Factory(&_AastraRouter.CallOpts)
}

// GetBaseAmounts is a free data retrieval call binding the contract method 0x7a163a50.
//
// Solidity: function getBaseAmounts(address _vault) view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AastraRouter *AastraRouterCaller) GetBaseAmounts(opts *bind.CallOpts, _vault common.Address) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _AastraRouter.contract.Call(opts, &out, "getBaseAmounts", _vault)

	outstruct := new(struct {
		Liquidity *big.Int
		Amount0   *big.Int
		Amount1   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount0 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBaseAmounts is a free data retrieval call binding the contract method 0x7a163a50.
//
// Solidity: function getBaseAmounts(address _vault) view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AastraRouter *AastraRouterSession) GetBaseAmounts(_vault common.Address) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _AastraRouter.Contract.GetBaseAmounts(&_AastraRouter.CallOpts, _vault)
}

// GetBaseAmounts is a free data retrieval call binding the contract method 0x7a163a50.
//
// Solidity: function getBaseAmounts(address _vault) view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AastraRouter *AastraRouterCallerSession) GetBaseAmounts(_vault common.Address) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _AastraRouter.Contract.GetBaseAmounts(&_AastraRouter.CallOpts, _vault)
}

// GetBaseTicks is a free data retrieval call binding the contract method 0x1adfaf15.
//
// Solidity: function getBaseTicks(address _vault) view returns(int24, int24)
func (_AastraRouter *AastraRouterCaller) GetBaseTicks(opts *bind.CallOpts, _vault common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AastraRouter.contract.Call(opts, &out, "getBaseTicks", _vault)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBaseTicks is a free data retrieval call binding the contract method 0x1adfaf15.
//
// Solidity: function getBaseTicks(address _vault) view returns(int24, int24)
func (_AastraRouter *AastraRouterSession) GetBaseTicks(_vault common.Address) (*big.Int, *big.Int, error) {
	return _AastraRouter.Contract.GetBaseTicks(&_AastraRouter.CallOpts, _vault)
}

// GetBaseTicks is a free data retrieval call binding the contract method 0x1adfaf15.
//
// Solidity: function getBaseTicks(address _vault) view returns(int24, int24)
func (_AastraRouter *AastraRouterCallerSession) GetBaseTicks(_vault common.Address) (*big.Int, *big.Int, error) {
	return _AastraRouter.Contract.GetBaseTicks(&_AastraRouter.CallOpts, _vault)
}

// GetRangeAmounts is a free data retrieval call binding the contract method 0xd0961ba5.
//
// Solidity: function getRangeAmounts(address _vault) view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AastraRouter *AastraRouterCaller) GetRangeAmounts(opts *bind.CallOpts, _vault common.Address) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _AastraRouter.contract.Call(opts, &out, "getRangeAmounts", _vault)

	outstruct := new(struct {
		Liquidity *big.Int
		Amount0   *big.Int
		Amount1   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount0 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRangeAmounts is a free data retrieval call binding the contract method 0xd0961ba5.
//
// Solidity: function getRangeAmounts(address _vault) view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AastraRouter *AastraRouterSession) GetRangeAmounts(_vault common.Address) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _AastraRouter.Contract.GetRangeAmounts(&_AastraRouter.CallOpts, _vault)
}

// GetRangeAmounts is a free data retrieval call binding the contract method 0xd0961ba5.
//
// Solidity: function getRangeAmounts(address _vault) view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AastraRouter *AastraRouterCallerSession) GetRangeAmounts(_vault common.Address) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _AastraRouter.Contract.GetRangeAmounts(&_AastraRouter.CallOpts, _vault)
}

// GetRangeTicks is a free data retrieval call binding the contract method 0x0b23d6ac.
//
// Solidity: function getRangeTicks(address _vault) view returns(int24, int24)
func (_AastraRouter *AastraRouterCaller) GetRangeTicks(opts *bind.CallOpts, _vault common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AastraRouter.contract.Call(opts, &out, "getRangeTicks", _vault)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRangeTicks is a free data retrieval call binding the contract method 0x0b23d6ac.
//
// Solidity: function getRangeTicks(address _vault) view returns(int24, int24)
func (_AastraRouter *AastraRouterSession) GetRangeTicks(_vault common.Address) (*big.Int, *big.Int, error) {
	return _AastraRouter.Contract.GetRangeTicks(&_AastraRouter.CallOpts, _vault)
}

// GetRangeTicks is a free data retrieval call binding the contract method 0x0b23d6ac.
//
// Solidity: function getRangeTicks(address _vault) view returns(int24, int24)
func (_AastraRouter *AastraRouterCallerSession) GetRangeTicks(_vault common.Address) (*big.Int, *big.Int, error) {
	return _AastraRouter.Contract.GetRangeTicks(&_AastraRouter.CallOpts, _vault)
}

// CompoundFee is a paid mutator transaction binding the contract method 0x4d27349d.
//
// Solidity: function compoundFee(address _vault) returns()
func (_AastraRouter *AastraRouterTransactor) CompoundFee(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _AastraRouter.contract.Transact(opts, "compoundFee", _vault)
}

// CompoundFee is a paid mutator transaction binding the contract method 0x4d27349d.
//
// Solidity: function compoundFee(address _vault) returns()
func (_AastraRouter *AastraRouterSession) CompoundFee(_vault common.Address) (*types.Transaction, error) {
	return _AastraRouter.Contract.CompoundFee(&_AastraRouter.TransactOpts, _vault)
}

// CompoundFee is a paid mutator transaction binding the contract method 0x4d27349d.
//
// Solidity: function compoundFee(address _vault) returns()
func (_AastraRouter *AastraRouterTransactorSession) CompoundFee(_vault common.Address) (*types.Transaction, error) {
	return _AastraRouter.Contract.CompoundFee(&_AastraRouter.TransactOpts, _vault)
}

// NewBaseLiquidity is a paid mutator transaction binding the contract method 0x729e4235.
//
// Solidity: function newBaseLiquidity(int24 _baseLower, int24 _baseUpper, uint8 _percentage) returns()
func (_AastraRouter *AastraRouterTransactor) NewBaseLiquidity(opts *bind.TransactOpts, _baseLower *big.Int, _baseUpper *big.Int, _percentage uint8) (*types.Transaction, error) {
	return _AastraRouter.contract.Transact(opts, "newBaseLiquidity", _baseLower, _baseUpper, _percentage)
}

// NewBaseLiquidity is a paid mutator transaction binding the contract method 0x729e4235.
//
// Solidity: function newBaseLiquidity(int24 _baseLower, int24 _baseUpper, uint8 _percentage) returns()
func (_AastraRouter *AastraRouterSession) NewBaseLiquidity(_baseLower *big.Int, _baseUpper *big.Int, _percentage uint8) (*types.Transaction, error) {
	return _AastraRouter.Contract.NewBaseLiquidity(&_AastraRouter.TransactOpts, _baseLower, _baseUpper, _percentage)
}

// NewBaseLiquidity is a paid mutator transaction binding the contract method 0x729e4235.
//
// Solidity: function newBaseLiquidity(int24 _baseLower, int24 _baseUpper, uint8 _percentage) returns()
func (_AastraRouter *AastraRouterTransactorSession) NewBaseLiquidity(_baseLower *big.Int, _baseUpper *big.Int, _percentage uint8) (*types.Transaction, error) {
	return _AastraRouter.Contract.NewBaseLiquidity(&_AastraRouter.TransactOpts, _baseLower, _baseUpper, _percentage)
}

// NewLimitLiquidity is a paid mutator transaction binding the contract method 0x37bcc0a8.
//
// Solidity: function newLimitLiquidity(int24 _limitLower, int24 _limitUpper, uint8 _percentage) returns()
func (_AastraRouter *AastraRouterTransactor) NewLimitLiquidity(opts *bind.TransactOpts, _limitLower *big.Int, _limitUpper *big.Int, _percentage uint8) (*types.Transaction, error) {
	return _AastraRouter.contract.Transact(opts, "newLimitLiquidity", _limitLower, _limitUpper, _percentage)
}

// NewLimitLiquidity is a paid mutator transaction binding the contract method 0x37bcc0a8.
//
// Solidity: function newLimitLiquidity(int24 _limitLower, int24 _limitUpper, uint8 _percentage) returns()
func (_AastraRouter *AastraRouterSession) NewLimitLiquidity(_limitLower *big.Int, _limitUpper *big.Int, _percentage uint8) (*types.Transaction, error) {
	return _AastraRouter.Contract.NewLimitLiquidity(&_AastraRouter.TransactOpts, _limitLower, _limitUpper, _percentage)
}

// NewLimitLiquidity is a paid mutator transaction binding the contract method 0x37bcc0a8.
//
// Solidity: function newLimitLiquidity(int24 _limitLower, int24 _limitUpper, uint8 _percentage) returns()
func (_AastraRouter *AastraRouterTransactorSession) NewLimitLiquidity(_limitLower *big.Int, _limitUpper *big.Int, _percentage uint8) (*types.Transaction, error) {
	return _AastraRouter.Contract.NewLimitLiquidity(&_AastraRouter.TransactOpts, _limitLower, _limitUpper, _percentage)
}
