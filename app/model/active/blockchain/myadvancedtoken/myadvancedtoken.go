// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package myadvancedtoken

import (
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"strings"
)

// MyadvancedtokenABI is the input ABI used to generate the binding from.
const MyadvancedtokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newSellPrice\",\"type\":\"uint256\"},{\"name\":\"newBuyPrice\",\"type\":\"uint256\"}],\"name\":\"setPrices\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sellPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buyPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"freeze\",\"type\":\"bool\"}],\"name\":\"freezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"FrozenFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"}]"

// Myadvancedtoken is an auto generated Go binding around an Ethereum contract.
type Myadvancedtoken struct {
	MyadvancedtokenCaller     // Read-only binding to the contract
	MyadvancedtokenTransactor // Write-only binding to the contract
	MyadvancedtokenFilterer   // Log filterer for contract events
}

// MyadvancedtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyadvancedtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyadvancedtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyadvancedtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyadvancedtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyadvancedtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyadvancedtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyadvancedtokenSession struct {
	Contract     *Myadvancedtoken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyadvancedtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyadvancedtokenCallerSession struct {
	Contract *MyadvancedtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MyadvancedtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyadvancedtokenTransactorSession struct {
	Contract     *MyadvancedtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MyadvancedtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyadvancedtokenRaw struct {
	Contract *Myadvancedtoken // Generic contract binding to access the raw methods on
}

// MyadvancedtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyadvancedtokenCallerRaw struct {
	Contract *MyadvancedtokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyadvancedtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyadvancedtokenTransactorRaw struct {
	Contract *MyadvancedtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyadvancedtoken creates a new instance of Myadvancedtoken, bound to a specific deployed contract.
func NewMyadvancedtoken(address common.Address, backend bind.ContractBackend) (*Myadvancedtoken, error) {
	contract, err := bindMyadvancedtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Myadvancedtoken{MyadvancedtokenCaller: MyadvancedtokenCaller{contract: contract}, MyadvancedtokenTransactor: MyadvancedtokenTransactor{contract: contract}, MyadvancedtokenFilterer: MyadvancedtokenFilterer{contract: contract}}, nil
}

// NewMyadvancedtokenCaller creates a new read-only instance of Myadvancedtoken, bound to a specific deployed contract.
func NewMyadvancedtokenCaller(address common.Address, caller bind.ContractCaller) (*MyadvancedtokenCaller, error) {
	contract, err := bindMyadvancedtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyadvancedtokenCaller{contract: contract}, nil
}

// NewMyadvancedtokenTransactor creates a new write-only instance of Myadvancedtoken, bound to a specific deployed contract.
func NewMyadvancedtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MyadvancedtokenTransactor, error) {
	contract, err := bindMyadvancedtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyadvancedtokenTransactor{contract: contract}, nil
}

// NewMyadvancedtokenFilterer creates a new log filterer instance of Myadvancedtoken, bound to a specific deployed contract.
func NewMyadvancedtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MyadvancedtokenFilterer, error) {
	contract, err := bindMyadvancedtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyadvancedtokenFilterer{contract: contract}, nil
}

// bindMyadvancedtoken binds a generic wrapper to an already deployed contract.
func bindMyadvancedtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyadvancedtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Myadvancedtoken *MyadvancedtokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Myadvancedtoken.Contract.MyadvancedtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Myadvancedtoken *MyadvancedtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.MyadvancedtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Myadvancedtoken *MyadvancedtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.MyadvancedtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Myadvancedtoken *MyadvancedtokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Myadvancedtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Myadvancedtoken *MyadvancedtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Myadvancedtoken *MyadvancedtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Myadvancedtoken.Contract.Allowance(&_Myadvancedtoken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Myadvancedtoken.Contract.Allowance(&_Myadvancedtoken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Myadvancedtoken.Contract.BalanceOf(&_Myadvancedtoken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Myadvancedtoken.Contract.BalanceOf(&_Myadvancedtoken.CallOpts, arg0)
}

// BuyPrice is a free data retrieval call binding the contract method 0x8620410b.
//
// Solidity: function buyPrice() constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCaller) BuyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "buyPrice")
	return *ret0, err
}

// BuyPrice is a free data retrieval call binding the contract method 0x8620410b.
//
// Solidity: function buyPrice() constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenSession) BuyPrice() (*big.Int, error) {
	return _Myadvancedtoken.Contract.BuyPrice(&_Myadvancedtoken.CallOpts)
}

// BuyPrice is a free data retrieval call binding the contract method 0x8620410b.
//
// Solidity: function buyPrice() constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) BuyPrice() (*big.Int, error) {
	return _Myadvancedtoken.Contract.BuyPrice(&_Myadvancedtoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Myadvancedtoken *MyadvancedtokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Myadvancedtoken *MyadvancedtokenSession) Decimals() (uint8, error) {
	return _Myadvancedtoken.Contract.Decimals(&_Myadvancedtoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) Decimals() (uint8, error) {
	return _Myadvancedtoken.Contract.Decimals(&_Myadvancedtoken.CallOpts)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xb414d4b6.
//
// Solidity: function frozenAccount( address) constant returns(bool)
func (_Myadvancedtoken *MyadvancedtokenCaller) FrozenAccount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "frozenAccount", arg0)
	return *ret0, err
}

// FrozenAccount is a free data retrieval call binding the contract method 0xb414d4b6.
//
// Solidity: function frozenAccount( address) constant returns(bool)
func (_Myadvancedtoken *MyadvancedtokenSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _Myadvancedtoken.Contract.FrozenAccount(&_Myadvancedtoken.CallOpts, arg0)
}

// FrozenAccount is a free data retrieval call binding the contract method 0xb414d4b6.
//
// Solidity: function frozenAccount( address) constant returns(bool)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) FrozenAccount(arg0 common.Address) (bool, error) {
	return _Myadvancedtoken.Contract.FrozenAccount(&_Myadvancedtoken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Myadvancedtoken *MyadvancedtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Myadvancedtoken *MyadvancedtokenSession) Name() (string, error) {
	return _Myadvancedtoken.Contract.Name(&_Myadvancedtoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) Name() (string, error) {
	return _Myadvancedtoken.Contract.Name(&_Myadvancedtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Myadvancedtoken *MyadvancedtokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Myadvancedtoken *MyadvancedtokenSession) Owner() (common.Address, error) {
	return _Myadvancedtoken.Contract.Owner(&_Myadvancedtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) Owner() (common.Address, error) {
	return _Myadvancedtoken.Contract.Owner(&_Myadvancedtoken.CallOpts)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCaller) SellPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "sellPrice")
	return *ret0, err
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenSession) SellPrice() (*big.Int, error) {
	return _Myadvancedtoken.Contract.SellPrice(&_Myadvancedtoken.CallOpts)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) SellPrice() (*big.Int, error) {
	return _Myadvancedtoken.Contract.SellPrice(&_Myadvancedtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Myadvancedtoken *MyadvancedtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Myadvancedtoken *MyadvancedtokenSession) Symbol() (string, error) {
	return _Myadvancedtoken.Contract.Symbol(&_Myadvancedtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) Symbol() (string, error) {
	return _Myadvancedtoken.Contract.Symbol(&_Myadvancedtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Myadvancedtoken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenSession) TotalSupply() (*big.Int, error) {
	return _Myadvancedtoken.Contract.TotalSupply(&_Myadvancedtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Myadvancedtoken *MyadvancedtokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Myadvancedtoken.Contract.TotalSupply(&_Myadvancedtoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Approve(&_Myadvancedtoken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Approve(&_Myadvancedtoken.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.ApproveAndCall(&_Myadvancedtoken.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.ApproveAndCall(&_Myadvancedtoken.TransactOpts, _spender, _value, _extraData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Burn(&_Myadvancedtoken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Burn(&_Myadvancedtoken.TransactOpts, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactor) BurnFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "burnFrom", _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.BurnFrom(&_Myadvancedtoken.TransactOpts, _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.BurnFrom(&_Myadvancedtoken.TransactOpts, _from, _value)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_Myadvancedtoken *MyadvancedtokenTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_Myadvancedtoken *MyadvancedtokenSession) Buy() (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Buy(&_Myadvancedtoken.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) Buy() (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Buy(&_Myadvancedtoken.TransactOpts)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xe724529c.
//
// Solidity: function freezeAccount(target address, freeze bool) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactor) FreezeAccount(opts *bind.TransactOpts, target common.Address, freeze bool) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "freezeAccount", target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xe724529c.
//
// Solidity: function freezeAccount(target address, freeze bool) returns()
func (_Myadvancedtoken *MyadvancedtokenSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.FreezeAccount(&_Myadvancedtoken.TransactOpts, target, freeze)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xe724529c.
//
// Solidity: function freezeAccount(target address, freeze bool) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) FreezeAccount(target common.Address, freeze bool) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.FreezeAccount(&_Myadvancedtoken.TransactOpts, target, freeze)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactor) MintToken(opts *bind.TransactOpts, target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "mintToken", target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns()
func (_Myadvancedtoken *MyadvancedtokenSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.MintToken(&_Myadvancedtoken.TransactOpts, target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.MintToken(&_Myadvancedtoken.TransactOpts, target, mintedAmount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(amount uint256) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactor) Sell(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "sell", amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(amount uint256) returns()
func (_Myadvancedtoken *MyadvancedtokenSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Sell(&_Myadvancedtoken.TransactOpts, amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(amount uint256) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Sell(&_Myadvancedtoken.TransactOpts, amount)
}

// SetPrices is a paid mutator transaction binding the contract method 0x05fefda7.
//
// Solidity: function setPrices(newSellPrice uint256, newBuyPrice uint256) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactor) SetPrices(opts *bind.TransactOpts, newSellPrice *big.Int, newBuyPrice *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "setPrices", newSellPrice, newBuyPrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0x05fefda7.
//
// Solidity: function setPrices(newSellPrice uint256, newBuyPrice uint256) returns()
func (_Myadvancedtoken *MyadvancedtokenSession) SetPrices(newSellPrice *big.Int, newBuyPrice *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.SetPrices(&_Myadvancedtoken.TransactOpts, newSellPrice, newBuyPrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0x05fefda7.
//
// Solidity: function setPrices(newSellPrice uint256, newBuyPrice uint256) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) SetPrices(newSellPrice *big.Int, newBuyPrice *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.SetPrices(&_Myadvancedtoken.TransactOpts, newSellPrice, newBuyPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Transfer(&_Myadvancedtoken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.Transfer(&_Myadvancedtoken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.TransferFrom(&_Myadvancedtoken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.TransferFrom(&_Myadvancedtoken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Myadvancedtoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Myadvancedtoken *MyadvancedtokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.TransferOwnership(&_Myadvancedtoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Myadvancedtoken *MyadvancedtokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Myadvancedtoken.Contract.TransferOwnership(&_Myadvancedtoken.TransactOpts, newOwner)
}

// MyadvancedtokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Myadvancedtoken contract.
type MyadvancedtokenBurnIterator struct {
	Event *MyadvancedtokenBurn // Event containing the contract specifics and raw log

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
func (it *MyadvancedtokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyadvancedtokenBurn)
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
		it.Event = new(MyadvancedtokenBurn)
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
func (it *MyadvancedtokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyadvancedtokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyadvancedtokenBurn represents a Burn event raised by the Myadvancedtoken contract.
type MyadvancedtokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(from indexed address, value uint256)
func (_Myadvancedtoken *MyadvancedtokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*MyadvancedtokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Myadvancedtoken.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &MyadvancedtokenBurnIterator{contract: _Myadvancedtoken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(from indexed address, value uint256)
func (_Myadvancedtoken *MyadvancedtokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *MyadvancedtokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Myadvancedtoken.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyadvancedtokenBurn)
				if err := _Myadvancedtoken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// MyadvancedtokenFrozenFundsIterator is returned from FilterFrozenFunds and is used to iterate over the raw logs and unpacked data for FrozenFunds events raised by the Myadvancedtoken contract.
type MyadvancedtokenFrozenFundsIterator struct {
	Event *MyadvancedtokenFrozenFunds // Event containing the contract specifics and raw log

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
func (it *MyadvancedtokenFrozenFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyadvancedtokenFrozenFunds)
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
		it.Event = new(MyadvancedtokenFrozenFunds)
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
func (it *MyadvancedtokenFrozenFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyadvancedtokenFrozenFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyadvancedtokenFrozenFunds represents a FrozenFunds event raised by the Myadvancedtoken contract.
type MyadvancedtokenFrozenFunds struct {
	Target common.Address
	Frozen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFrozenFunds is a free log retrieval operation binding the contract event 0x48335238b4855f35377ed80f164e8c6f3c366e54ac00b96a6402d4a9814a03a5.
//
// Solidity: e FrozenFunds(target address, frozen bool)
func (_Myadvancedtoken *MyadvancedtokenFilterer) FilterFrozenFunds(opts *bind.FilterOpts) (*MyadvancedtokenFrozenFundsIterator, error) {

	logs, sub, err := _Myadvancedtoken.contract.FilterLogs(opts, "FrozenFunds")
	if err != nil {
		return nil, err
	}
	return &MyadvancedtokenFrozenFundsIterator{contract: _Myadvancedtoken.contract, event: "FrozenFunds", logs: logs, sub: sub}, nil
}

// WatchFrozenFunds is a free log subscription operation binding the contract event 0x48335238b4855f35377ed80f164e8c6f3c366e54ac00b96a6402d4a9814a03a5.
//
// Solidity: e FrozenFunds(target address, frozen bool)
func (_Myadvancedtoken *MyadvancedtokenFilterer) WatchFrozenFunds(opts *bind.WatchOpts, sink chan<- *MyadvancedtokenFrozenFunds) (event.Subscription, error) {

	logs, sub, err := _Myadvancedtoken.contract.WatchLogs(opts, "FrozenFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyadvancedtokenFrozenFunds)
				if err := _Myadvancedtoken.contract.UnpackLog(event, "FrozenFunds", log); err != nil {
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

// MyadvancedtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Myadvancedtoken contract.
type MyadvancedtokenTransferIterator struct {
	Event *MyadvancedtokenTransfer // Event containing the contract specifics and raw log

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
func (it *MyadvancedtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyadvancedtokenTransfer)
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
		it.Event = new(MyadvancedtokenTransfer)
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
func (it *MyadvancedtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyadvancedtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyadvancedtokenTransfer represents a Transfer event raised by the Myadvancedtoken contract.
type MyadvancedtokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Myadvancedtoken *MyadvancedtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyadvancedtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Myadvancedtoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyadvancedtokenTransferIterator{contract: _Myadvancedtoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Myadvancedtoken *MyadvancedtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyadvancedtokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Myadvancedtoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyadvancedtokenTransfer)
				if err := _Myadvancedtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
