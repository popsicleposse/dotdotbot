// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dotdotbot

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

// DotdotbotMetaData contains all meta data concerning the Dotdotbot contract.
var DotdotbotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferToOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"tryMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DotdotbotABI is the input ABI used to generate the binding from.
// Deprecated: Use DotdotbotMetaData.ABI instead.
var DotdotbotABI = DotdotbotMetaData.ABI

// Dotdotbot is an auto generated Go binding around an Ethereum contract.
type Dotdotbot struct {
	DotdotbotCaller     // Read-only binding to the contract
	DotdotbotTransactor // Write-only binding to the contract
	DotdotbotFilterer   // Log filterer for contract events
}

// DotdotbotCaller is an auto generated read-only Go binding around an Ethereum contract.
type DotdotbotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DotdotbotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DotdotbotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DotdotbotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DotdotbotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DotdotbotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DotdotbotSession struct {
	Contract     *Dotdotbot        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DotdotbotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DotdotbotCallerSession struct {
	Contract *DotdotbotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DotdotbotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DotdotbotTransactorSession struct {
	Contract     *DotdotbotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DotdotbotRaw is an auto generated low-level Go binding around an Ethereum contract.
type DotdotbotRaw struct {
	Contract *Dotdotbot // Generic contract binding to access the raw methods on
}

// DotdotbotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DotdotbotCallerRaw struct {
	Contract *DotdotbotCaller // Generic read-only contract binding to access the raw methods on
}

// DotdotbotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DotdotbotTransactorRaw struct {
	Contract *DotdotbotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDotdotbot creates a new instance of Dotdotbot, bound to a specific deployed contract.
func NewDotdotbot(address common.Address, backend bind.ContractBackend) (*Dotdotbot, error) {
	contract, err := bindDotdotbot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dotdotbot{DotdotbotCaller: DotdotbotCaller{contract: contract}, DotdotbotTransactor: DotdotbotTransactor{contract: contract}, DotdotbotFilterer: DotdotbotFilterer{contract: contract}}, nil
}

// NewDotdotbotCaller creates a new read-only instance of Dotdotbot, bound to a specific deployed contract.
func NewDotdotbotCaller(address common.Address, caller bind.ContractCaller) (*DotdotbotCaller, error) {
	contract, err := bindDotdotbot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DotdotbotCaller{contract: contract}, nil
}

// NewDotdotbotTransactor creates a new write-only instance of Dotdotbot, bound to a specific deployed contract.
func NewDotdotbotTransactor(address common.Address, transactor bind.ContractTransactor) (*DotdotbotTransactor, error) {
	contract, err := bindDotdotbot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DotdotbotTransactor{contract: contract}, nil
}

// NewDotdotbotFilterer creates a new log filterer instance of Dotdotbot, bound to a specific deployed contract.
func NewDotdotbotFilterer(address common.Address, filterer bind.ContractFilterer) (*DotdotbotFilterer, error) {
	contract, err := bindDotdotbot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DotdotbotFilterer{contract: contract}, nil
}

// bindDotdotbot binds a generic wrapper to an already deployed contract.
func bindDotdotbot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DotdotbotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dotdotbot *DotdotbotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dotdotbot.Contract.DotdotbotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dotdotbot *DotdotbotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotbot.Contract.DotdotbotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dotdotbot *DotdotbotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dotdotbot.Contract.DotdotbotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dotdotbot *DotdotbotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dotdotbot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dotdotbot *DotdotbotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotbot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dotdotbot *DotdotbotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dotdotbot.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dotdotbot *DotdotbotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dotdotbot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dotdotbot *DotdotbotSession) Owner() (common.Address, error) {
	return _Dotdotbot.Contract.Owner(&_Dotdotbot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dotdotbot *DotdotbotCallerSession) Owner() (common.Address, error) {
	return _Dotdotbot.Contract.Owner(&_Dotdotbot.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Dotdotbot *DotdotbotTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotbot.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Dotdotbot *DotdotbotSession) Deposit() (*types.Transaction, error) {
	return _Dotdotbot.Contract.Deposit(&_Dotdotbot.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Dotdotbot *DotdotbotTransactorSession) Deposit() (*types.Transaction, error) {
	return _Dotdotbot.Contract.Deposit(&_Dotdotbot.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Dotdotbot *DotdotbotTransactor) OnERC721Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Dotdotbot.contract.Transact(opts, "onERC721Received", _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Dotdotbot *DotdotbotSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Dotdotbot.Contract.OnERC721Received(&_Dotdotbot.TransactOpts, _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) returns(bytes4)
func (_Dotdotbot *DotdotbotTransactorSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Dotdotbot.Contract.OnERC721Received(&_Dotdotbot.TransactOpts, _operator, _from, _tokenId, _data)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address addr) returns()
func (_Dotdotbot *DotdotbotTransactor) SetImplementation(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Dotdotbot.contract.Transact(opts, "setImplementation", addr)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address addr) returns()
func (_Dotdotbot *DotdotbotSession) SetImplementation(addr common.Address) (*types.Transaction, error) {
	return _Dotdotbot.Contract.SetImplementation(&_Dotdotbot.TransactOpts, addr)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address addr) returns()
func (_Dotdotbot *DotdotbotTransactorSession) SetImplementation(addr common.Address) (*types.Transaction, error) {
	return _Dotdotbot.Contract.SetImplementation(&_Dotdotbot.TransactOpts, addr)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 price) returns()
func (_Dotdotbot *DotdotbotTransactor) SetPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Dotdotbot.contract.Transact(opts, "setPrice", price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 price) returns()
func (_Dotdotbot *DotdotbotSession) SetPrice(price *big.Int) (*types.Transaction, error) {
	return _Dotdotbot.Contract.SetPrice(&_Dotdotbot.TransactOpts, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 price) returns()
func (_Dotdotbot *DotdotbotTransactorSession) SetPrice(price *big.Int) (*types.Transaction, error) {
	return _Dotdotbot.Contract.SetPrice(&_Dotdotbot.TransactOpts, price)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address addr, bool status) returns()
func (_Dotdotbot *DotdotbotTransactor) SetWhitelisted(opts *bind.TransactOpts, addr common.Address, status bool) (*types.Transaction, error) {
	return _Dotdotbot.contract.Transact(opts, "setWhitelisted", addr, status)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address addr, bool status) returns()
func (_Dotdotbot *DotdotbotSession) SetWhitelisted(addr common.Address, status bool) (*types.Transaction, error) {
	return _Dotdotbot.Contract.SetWhitelisted(&_Dotdotbot.TransactOpts, addr, status)
}

// SetWhitelisted is a paid mutator transaction binding the contract method 0x9281aa0b.
//
// Solidity: function setWhitelisted(address addr, bool status) returns()
func (_Dotdotbot *DotdotbotTransactorSession) SetWhitelisted(addr common.Address, status bool) (*types.Transaction, error) {
	return _Dotdotbot.Contract.SetWhitelisted(&_Dotdotbot.TransactOpts, addr, status)
}

// TransferToOwner is a paid mutator transaction binding the contract method 0x2d90ae94.
//
// Solidity: function transferToOwner() returns()
func (_Dotdotbot *DotdotbotTransactor) TransferToOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotbot.contract.Transact(opts, "transferToOwner")
}

// TransferToOwner is a paid mutator transaction binding the contract method 0x2d90ae94.
//
// Solidity: function transferToOwner() returns()
func (_Dotdotbot *DotdotbotSession) TransferToOwner() (*types.Transaction, error) {
	return _Dotdotbot.Contract.TransferToOwner(&_Dotdotbot.TransactOpts)
}

// TransferToOwner is a paid mutator transaction binding the contract method 0x2d90ae94.
//
// Solidity: function transferToOwner() returns()
func (_Dotdotbot *DotdotbotTransactorSession) TransferToOwner() (*types.Transaction, error) {
	return _Dotdotbot.Contract.TransferToOwner(&_Dotdotbot.TransactOpts)
}

// TryMint is a paid mutator transaction binding the contract method 0x55486283.
//
// Solidity: function tryMint(uint256 count) returns()
func (_Dotdotbot *DotdotbotTransactor) TryMint(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _Dotdotbot.contract.Transact(opts, "tryMint", count)
}

// TryMint is a paid mutator transaction binding the contract method 0x55486283.
//
// Solidity: function tryMint(uint256 count) returns()
func (_Dotdotbot *DotdotbotSession) TryMint(count *big.Int) (*types.Transaction, error) {
	return _Dotdotbot.Contract.TryMint(&_Dotdotbot.TransactOpts, count)
}

// TryMint is a paid mutator transaction binding the contract method 0x55486283.
//
// Solidity: function tryMint(uint256 count) returns()
func (_Dotdotbot *DotdotbotTransactorSession) TryMint(count *big.Int) (*types.Transaction, error) {
	return _Dotdotbot.Contract.TryMint(&_Dotdotbot.TransactOpts, count)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dotdotbot *DotdotbotTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotbot.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dotdotbot *DotdotbotSession) Withdraw() (*types.Transaction, error) {
	return _Dotdotbot.Contract.Withdraw(&_Dotdotbot.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dotdotbot *DotdotbotTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Dotdotbot.Contract.Withdraw(&_Dotdotbot.TransactOpts)
}
