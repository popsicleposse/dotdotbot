// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shroomscouts

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

// ShroomscoutsMetaData contains all meta data concerning the Shroomscouts contract.
var ShroomscoutsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ShroomscoutsABI is the input ABI used to generate the binding from.
// Deprecated: Use ShroomscoutsMetaData.ABI instead.
var ShroomscoutsABI = ShroomscoutsMetaData.ABI

// Shroomscouts is an auto generated Go binding around an Ethereum contract.
type Shroomscouts struct {
	ShroomscoutsCaller     // Read-only binding to the contract
	ShroomscoutsTransactor // Write-only binding to the contract
	ShroomscoutsFilterer   // Log filterer for contract events
}

// ShroomscoutsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShroomscoutsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShroomscoutsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShroomscoutsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShroomscoutsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShroomscoutsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShroomscoutsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShroomscoutsSession struct {
	Contract     *Shroomscouts     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShroomscoutsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShroomscoutsCallerSession struct {
	Contract *ShroomscoutsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ShroomscoutsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShroomscoutsTransactorSession struct {
	Contract     *ShroomscoutsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ShroomscoutsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShroomscoutsRaw struct {
	Contract *Shroomscouts // Generic contract binding to access the raw methods on
}

// ShroomscoutsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShroomscoutsCallerRaw struct {
	Contract *ShroomscoutsCaller // Generic read-only contract binding to access the raw methods on
}

// ShroomscoutsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShroomscoutsTransactorRaw struct {
	Contract *ShroomscoutsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShroomscouts creates a new instance of Shroomscouts, bound to a specific deployed contract.
func NewShroomscouts(address common.Address, backend bind.ContractBackend) (*Shroomscouts, error) {
	contract, err := bindShroomscouts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Shroomscouts{ShroomscoutsCaller: ShroomscoutsCaller{contract: contract}, ShroomscoutsTransactor: ShroomscoutsTransactor{contract: contract}, ShroomscoutsFilterer: ShroomscoutsFilterer{contract: contract}}, nil
}

// NewShroomscoutsCaller creates a new read-only instance of Shroomscouts, bound to a specific deployed contract.
func NewShroomscoutsCaller(address common.Address, caller bind.ContractCaller) (*ShroomscoutsCaller, error) {
	contract, err := bindShroomscouts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShroomscoutsCaller{contract: contract}, nil
}

// NewShroomscoutsTransactor creates a new write-only instance of Shroomscouts, bound to a specific deployed contract.
func NewShroomscoutsTransactor(address common.Address, transactor bind.ContractTransactor) (*ShroomscoutsTransactor, error) {
	contract, err := bindShroomscouts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShroomscoutsTransactor{contract: contract}, nil
}

// NewShroomscoutsFilterer creates a new log filterer instance of Shroomscouts, bound to a specific deployed contract.
func NewShroomscoutsFilterer(address common.Address, filterer bind.ContractFilterer) (*ShroomscoutsFilterer, error) {
	contract, err := bindShroomscouts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShroomscoutsFilterer{contract: contract}, nil
}

// bindShroomscouts binds a generic wrapper to an already deployed contract.
func bindShroomscouts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShroomscoutsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shroomscouts *ShroomscoutsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shroomscouts.Contract.ShroomscoutsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shroomscouts *ShroomscoutsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shroomscouts.Contract.ShroomscoutsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shroomscouts *ShroomscoutsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shroomscouts.Contract.ShroomscoutsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shroomscouts *ShroomscoutsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shroomscouts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shroomscouts *ShroomscoutsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shroomscouts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shroomscouts *ShroomscoutsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shroomscouts.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_Shroomscouts *ShroomscoutsCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Shroomscouts.contract.Call(opts, &out, "balanceOf", _owner, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_Shroomscouts *ShroomscoutsSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _Shroomscouts.Contract.BalanceOf(&_Shroomscouts.CallOpts, _owner, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_Shroomscouts *ShroomscoutsCallerSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _Shroomscouts.Contract.BalanceOf(&_Shroomscouts.CallOpts, _owner, _id)
}

// MintToken is a paid mutator transaction binding the contract method 0xc634d032.
//
// Solidity: function mintToken(uint256 amount) returns()
func (_Shroomscouts *ShroomscoutsTransactor) MintToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Shroomscouts.contract.Transact(opts, "mintToken", amount)
}

// MintToken is a paid mutator transaction binding the contract method 0xc634d032.
//
// Solidity: function mintToken(uint256 amount) returns()
func (_Shroomscouts *ShroomscoutsSession) MintToken(amount *big.Int) (*types.Transaction, error) {
	return _Shroomscouts.Contract.MintToken(&_Shroomscouts.TransactOpts, amount)
}

// MintToken is a paid mutator transaction binding the contract method 0xc634d032.
//
// Solidity: function mintToken(uint256 amount) returns()
func (_Shroomscouts *ShroomscoutsTransactorSession) MintToken(amount *big.Int) (*types.Transaction, error) {
	return _Shroomscouts.Contract.MintToken(&_Shroomscouts.TransactOpts, amount)
}
