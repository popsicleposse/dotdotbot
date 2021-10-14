// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package junglefreaks

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

// JunglefreaksMetaData contains all meta data concerning the Junglefreaks contract.
var JunglefreaksMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Whitelist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"reserveMints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePresale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"who\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"whitelistAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// JunglefreaksABI is the input ABI used to generate the binding from.
// Deprecated: Use JunglefreaksMetaData.ABI instead.
var JunglefreaksABI = JunglefreaksMetaData.ABI

// Junglefreaks is an auto generated Go binding around an Ethereum contract.
type Junglefreaks struct {
	JunglefreaksCaller     // Read-only binding to the contract
	JunglefreaksTransactor // Write-only binding to the contract
	JunglefreaksFilterer   // Log filterer for contract events
}

// JunglefreaksCaller is an auto generated read-only Go binding around an Ethereum contract.
type JunglefreaksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JunglefreaksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JunglefreaksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JunglefreaksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JunglefreaksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JunglefreaksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JunglefreaksSession struct {
	Contract     *Junglefreaks     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JunglefreaksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JunglefreaksCallerSession struct {
	Contract *JunglefreaksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// JunglefreaksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JunglefreaksTransactorSession struct {
	Contract     *JunglefreaksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// JunglefreaksRaw is an auto generated low-level Go binding around an Ethereum contract.
type JunglefreaksRaw struct {
	Contract *Junglefreaks // Generic contract binding to access the raw methods on
}

// JunglefreaksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JunglefreaksCallerRaw struct {
	Contract *JunglefreaksCaller // Generic read-only contract binding to access the raw methods on
}

// JunglefreaksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JunglefreaksTransactorRaw struct {
	Contract *JunglefreaksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJunglefreaks creates a new instance of Junglefreaks, bound to a specific deployed contract.
func NewJunglefreaks(address common.Address, backend bind.ContractBackend) (*Junglefreaks, error) {
	contract, err := bindJunglefreaks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Junglefreaks{JunglefreaksCaller: JunglefreaksCaller{contract: contract}, JunglefreaksTransactor: JunglefreaksTransactor{contract: contract}, JunglefreaksFilterer: JunglefreaksFilterer{contract: contract}}, nil
}

// NewJunglefreaksCaller creates a new read-only instance of Junglefreaks, bound to a specific deployed contract.
func NewJunglefreaksCaller(address common.Address, caller bind.ContractCaller) (*JunglefreaksCaller, error) {
	contract, err := bindJunglefreaks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JunglefreaksCaller{contract: contract}, nil
}

// NewJunglefreaksTransactor creates a new write-only instance of Junglefreaks, bound to a specific deployed contract.
func NewJunglefreaksTransactor(address common.Address, transactor bind.ContractTransactor) (*JunglefreaksTransactor, error) {
	contract, err := bindJunglefreaks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JunglefreaksTransactor{contract: contract}, nil
}

// NewJunglefreaksFilterer creates a new log filterer instance of Junglefreaks, bound to a specific deployed contract.
func NewJunglefreaksFilterer(address common.Address, filterer bind.ContractFilterer) (*JunglefreaksFilterer, error) {
	contract, err := bindJunglefreaks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JunglefreaksFilterer{contract: contract}, nil
}

// bindJunglefreaks binds a generic wrapper to an already deployed contract.
func bindJunglefreaks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JunglefreaksABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Junglefreaks *JunglefreaksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Junglefreaks.Contract.JunglefreaksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Junglefreaks *JunglefreaksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Junglefreaks.Contract.JunglefreaksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Junglefreaks *JunglefreaksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Junglefreaks.Contract.JunglefreaksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Junglefreaks *JunglefreaksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Junglefreaks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Junglefreaks *JunglefreaksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Junglefreaks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Junglefreaks *JunglefreaksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Junglefreaks.Contract.contract.Transact(opts, method, params...)
}

// Whitelist is a free data retrieval call binding the contract method 0xeb73900b.
//
// Solidity: function Whitelist(address ) view returns(uint256)
func (_Junglefreaks *JunglefreaksCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "Whitelist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0xeb73900b.
//
// Solidity: function Whitelist(address ) view returns(uint256)
func (_Junglefreaks *JunglefreaksSession) Whitelist(arg0 common.Address) (*big.Int, error) {
	return _Junglefreaks.Contract.Whitelist(&_Junglefreaks.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0xeb73900b.
//
// Solidity: function Whitelist(address ) view returns(uint256)
func (_Junglefreaks *JunglefreaksCallerSession) Whitelist(arg0 common.Address) (*big.Int, error) {
	return _Junglefreaks.Contract.Whitelist(&_Junglefreaks.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Junglefreaks *JunglefreaksCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Junglefreaks *JunglefreaksSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Junglefreaks.Contract.BalanceOf(&_Junglefreaks.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Junglefreaks *JunglefreaksCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Junglefreaks.Contract.BalanceOf(&_Junglefreaks.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Junglefreaks *JunglefreaksCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Junglefreaks *JunglefreaksSession) BaseURI() (string, error) {
	return _Junglefreaks.Contract.BaseURI(&_Junglefreaks.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Junglefreaks *JunglefreaksCallerSession) BaseURI() (string, error) {
	return _Junglefreaks.Contract.BaseURI(&_Junglefreaks.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Junglefreaks *JunglefreaksCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Junglefreaks *JunglefreaksSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Junglefreaks.Contract.GetApproved(&_Junglefreaks.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Junglefreaks *JunglefreaksCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Junglefreaks.Contract.GetApproved(&_Junglefreaks.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Junglefreaks *JunglefreaksCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Junglefreaks *JunglefreaksSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Junglefreaks.Contract.IsApprovedForAll(&_Junglefreaks.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Junglefreaks *JunglefreaksCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Junglefreaks.Contract.IsApprovedForAll(&_Junglefreaks.CallOpts, owner, operator)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Junglefreaks *JunglefreaksCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Junglefreaks *JunglefreaksSession) MaxSupply() (*big.Int, error) {
	return _Junglefreaks.Contract.MaxSupply(&_Junglefreaks.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Junglefreaks *JunglefreaksCallerSession) MaxSupply() (*big.Int, error) {
	return _Junglefreaks.Contract.MaxSupply(&_Junglefreaks.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Junglefreaks *JunglefreaksCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Junglefreaks *JunglefreaksSession) Name() (string, error) {
	return _Junglefreaks.Contract.Name(&_Junglefreaks.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Junglefreaks *JunglefreaksCallerSession) Name() (string, error) {
	return _Junglefreaks.Contract.Name(&_Junglefreaks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Junglefreaks *JunglefreaksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Junglefreaks *JunglefreaksSession) Owner() (common.Address, error) {
	return _Junglefreaks.Contract.Owner(&_Junglefreaks.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Junglefreaks *JunglefreaksCallerSession) Owner() (common.Address, error) {
	return _Junglefreaks.Contract.Owner(&_Junglefreaks.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Junglefreaks *JunglefreaksCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Junglefreaks *JunglefreaksSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Junglefreaks.Contract.OwnerOf(&_Junglefreaks.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Junglefreaks *JunglefreaksCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Junglefreaks.Contract.OwnerOf(&_Junglefreaks.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Junglefreaks *JunglefreaksCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Junglefreaks *JunglefreaksSession) Paused() (bool, error) {
	return _Junglefreaks.Contract.Paused(&_Junglefreaks.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Junglefreaks *JunglefreaksCallerSession) Paused() (bool, error) {
	return _Junglefreaks.Contract.Paused(&_Junglefreaks.CallOpts)
}

// PresaleOpen is a free data retrieval call binding the contract method 0xbee6348a.
//
// Solidity: function presaleOpen() view returns(bool)
func (_Junglefreaks *JunglefreaksCaller) PresaleOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "presaleOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PresaleOpen is a free data retrieval call binding the contract method 0xbee6348a.
//
// Solidity: function presaleOpen() view returns(bool)
func (_Junglefreaks *JunglefreaksSession) PresaleOpen() (bool, error) {
	return _Junglefreaks.Contract.PresaleOpen(&_Junglefreaks.CallOpts)
}

// PresaleOpen is a free data retrieval call binding the contract method 0xbee6348a.
//
// Solidity: function presaleOpen() view returns(bool)
func (_Junglefreaks *JunglefreaksCallerSession) PresaleOpen() (bool, error) {
	return _Junglefreaks.Contract.PresaleOpen(&_Junglefreaks.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Junglefreaks *JunglefreaksCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Junglefreaks *JunglefreaksSession) Price() (*big.Int, error) {
	return _Junglefreaks.Contract.Price(&_Junglefreaks.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Junglefreaks *JunglefreaksCallerSession) Price() (*big.Int, error) {
	return _Junglefreaks.Contract.Price(&_Junglefreaks.CallOpts)
}

// SaleOpen is a free data retrieval call binding the contract method 0x99288dbb.
//
// Solidity: function saleOpen() view returns(bool)
func (_Junglefreaks *JunglefreaksCaller) SaleOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "saleOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleOpen is a free data retrieval call binding the contract method 0x99288dbb.
//
// Solidity: function saleOpen() view returns(bool)
func (_Junglefreaks *JunglefreaksSession) SaleOpen() (bool, error) {
	return _Junglefreaks.Contract.SaleOpen(&_Junglefreaks.CallOpts)
}

// SaleOpen is a free data retrieval call binding the contract method 0x99288dbb.
//
// Solidity: function saleOpen() view returns(bool)
func (_Junglefreaks *JunglefreaksCallerSession) SaleOpen() (bool, error) {
	return _Junglefreaks.Contract.SaleOpen(&_Junglefreaks.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Junglefreaks *JunglefreaksCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Junglefreaks *JunglefreaksSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Junglefreaks.Contract.SupportsInterface(&_Junglefreaks.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Junglefreaks *JunglefreaksCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Junglefreaks.Contract.SupportsInterface(&_Junglefreaks.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Junglefreaks *JunglefreaksCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Junglefreaks *JunglefreaksSession) Symbol() (string, error) {
	return _Junglefreaks.Contract.Symbol(&_Junglefreaks.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Junglefreaks *JunglefreaksCallerSession) Symbol() (string, error) {
	return _Junglefreaks.Contract.Symbol(&_Junglefreaks.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Junglefreaks *JunglefreaksCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Junglefreaks *JunglefreaksSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Junglefreaks.Contract.TokenByIndex(&_Junglefreaks.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Junglefreaks *JunglefreaksCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Junglefreaks.Contract.TokenByIndex(&_Junglefreaks.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Junglefreaks *JunglefreaksCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Junglefreaks *JunglefreaksSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Junglefreaks.Contract.TokenOfOwnerByIndex(&_Junglefreaks.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Junglefreaks *JunglefreaksCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Junglefreaks.Contract.TokenOfOwnerByIndex(&_Junglefreaks.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Junglefreaks *JunglefreaksCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Junglefreaks *JunglefreaksSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Junglefreaks.Contract.TokenURI(&_Junglefreaks.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Junglefreaks *JunglefreaksCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Junglefreaks.Contract.TokenURI(&_Junglefreaks.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Junglefreaks *JunglefreaksCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Junglefreaks.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Junglefreaks *JunglefreaksSession) TotalSupply() (*big.Int, error) {
	return _Junglefreaks.Contract.TotalSupply(&_Junglefreaks.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Junglefreaks *JunglefreaksCallerSession) TotalSupply() (*big.Int, error) {
	return _Junglefreaks.Contract.TotalSupply(&_Junglefreaks.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.Approve(&_Junglefreaks.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.Approve(&_Junglefreaks.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.Burn(&_Junglefreaks.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.Burn(&_Junglefreaks.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) payable returns()
func (_Junglefreaks *JunglefreaksTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) payable returns()
func (_Junglefreaks *JunglefreaksSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.Mint(&_Junglefreaks.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) payable returns()
func (_Junglefreaks *JunglefreaksTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.Mint(&_Junglefreaks.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Junglefreaks *JunglefreaksTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Junglefreaks *JunglefreaksSession) Pause() (*types.Transaction, error) {
	return _Junglefreaks.Contract.Pause(&_Junglefreaks.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Junglefreaks *JunglefreaksTransactorSession) Pause() (*types.Transaction, error) {
	return _Junglefreaks.Contract.Pause(&_Junglefreaks.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Junglefreaks *JunglefreaksTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Junglefreaks *JunglefreaksSession) RenounceOwnership() (*types.Transaction, error) {
	return _Junglefreaks.Contract.RenounceOwnership(&_Junglefreaks.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Junglefreaks *JunglefreaksTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Junglefreaks.Contract.RenounceOwnership(&_Junglefreaks.TransactOpts)
}

// ReserveMints is a paid mutator transaction binding the contract method 0xda5f6843.
//
// Solidity: function reserveMints(address to) returns()
func (_Junglefreaks *JunglefreaksTransactor) ReserveMints(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "reserveMints", to)
}

// ReserveMints is a paid mutator transaction binding the contract method 0xda5f6843.
//
// Solidity: function reserveMints(address to) returns()
func (_Junglefreaks *JunglefreaksSession) ReserveMints(to common.Address) (*types.Transaction, error) {
	return _Junglefreaks.Contract.ReserveMints(&_Junglefreaks.TransactOpts, to)
}

// ReserveMints is a paid mutator transaction binding the contract method 0xda5f6843.
//
// Solidity: function reserveMints(address to) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) ReserveMints(to common.Address) (*types.Transaction, error) {
	return _Junglefreaks.Contract.ReserveMints(&_Junglefreaks.TransactOpts, to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns()
func (_Junglefreaks *JunglefreaksTransactor) SafeMint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "safeMint", to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns()
func (_Junglefreaks *JunglefreaksSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SafeMint(&_Junglefreaks.TransactOpts, to)
}

// SafeMint is a paid mutator transaction binding the contract method 0x40d097c3.
//
// Solidity: function safeMint(address to) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) SafeMint(to common.Address) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SafeMint(&_Junglefreaks.TransactOpts, to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SafeTransferFrom(&_Junglefreaks.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SafeTransferFrom(&_Junglefreaks.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Junglefreaks *JunglefreaksTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Junglefreaks *JunglefreaksSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SafeTransferFrom0(&_Junglefreaks.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SafeTransferFrom0(&_Junglefreaks.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Junglefreaks *JunglefreaksTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Junglefreaks *JunglefreaksSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SetApprovalForAll(&_Junglefreaks.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SetApprovalForAll(&_Junglefreaks.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Junglefreaks *JunglefreaksTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Junglefreaks *JunglefreaksSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SetBaseURI(&_Junglefreaks.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SetBaseURI(&_Junglefreaks.TransactOpts, newBaseURI)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 newPrice) returns()
func (_Junglefreaks *JunglefreaksTransactor) SetPrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "setPrice", newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 newPrice) returns()
func (_Junglefreaks *JunglefreaksSession) SetPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SetPrice(&_Junglefreaks.TransactOpts, newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 newPrice) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) SetPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.SetPrice(&_Junglefreaks.TransactOpts, newPrice)
}

// TogglePresale is a paid mutator transaction binding the contract method 0x34393743.
//
// Solidity: function togglePresale() returns()
func (_Junglefreaks *JunglefreaksTransactor) TogglePresale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "togglePresale")
}

// TogglePresale is a paid mutator transaction binding the contract method 0x34393743.
//
// Solidity: function togglePresale() returns()
func (_Junglefreaks *JunglefreaksSession) TogglePresale() (*types.Transaction, error) {
	return _Junglefreaks.Contract.TogglePresale(&_Junglefreaks.TransactOpts)
}

// TogglePresale is a paid mutator transaction binding the contract method 0x34393743.
//
// Solidity: function togglePresale() returns()
func (_Junglefreaks *JunglefreaksTransactorSession) TogglePresale() (*types.Transaction, error) {
	return _Junglefreaks.Contract.TogglePresale(&_Junglefreaks.TransactOpts)
}

// ToggleSale is a paid mutator transaction binding the contract method 0x7d8966e4.
//
// Solidity: function toggleSale() returns()
func (_Junglefreaks *JunglefreaksTransactor) ToggleSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "toggleSale")
}

// ToggleSale is a paid mutator transaction binding the contract method 0x7d8966e4.
//
// Solidity: function toggleSale() returns()
func (_Junglefreaks *JunglefreaksSession) ToggleSale() (*types.Transaction, error) {
	return _Junglefreaks.Contract.ToggleSale(&_Junglefreaks.TransactOpts)
}

// ToggleSale is a paid mutator transaction binding the contract method 0x7d8966e4.
//
// Solidity: function toggleSale() returns()
func (_Junglefreaks *JunglefreaksTransactorSession) ToggleSale() (*types.Transaction, error) {
	return _Junglefreaks.Contract.ToggleSale(&_Junglefreaks.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.TransferFrom(&_Junglefreaks.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.TransferFrom(&_Junglefreaks.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Junglefreaks *JunglefreaksTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Junglefreaks *JunglefreaksSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Junglefreaks.Contract.TransferOwnership(&_Junglefreaks.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Junglefreaks.Contract.TransferOwnership(&_Junglefreaks.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Junglefreaks *JunglefreaksTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Junglefreaks *JunglefreaksSession) Unpause() (*types.Transaction, error) {
	return _Junglefreaks.Contract.Unpause(&_Junglefreaks.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Junglefreaks *JunglefreaksTransactorSession) Unpause() (*types.Transaction, error) {
	return _Junglefreaks.Contract.Unpause(&_Junglefreaks.TransactOpts)
}

// WhitelistAddress is a paid mutator transaction binding the contract method 0x6b6384db.
//
// Solidity: function whitelistAddress(address[] who, uint256 amount) returns()
func (_Junglefreaks *JunglefreaksTransactor) WhitelistAddress(opts *bind.TransactOpts, who []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "whitelistAddress", who, amount)
}

// WhitelistAddress is a paid mutator transaction binding the contract method 0x6b6384db.
//
// Solidity: function whitelistAddress(address[] who, uint256 amount) returns()
func (_Junglefreaks *JunglefreaksSession) WhitelistAddress(who []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.WhitelistAddress(&_Junglefreaks.TransactOpts, who, amount)
}

// WhitelistAddress is a paid mutator transaction binding the contract method 0x6b6384db.
//
// Solidity: function whitelistAddress(address[] who, uint256 amount) returns()
func (_Junglefreaks *JunglefreaksTransactorSession) WhitelistAddress(who []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Junglefreaks.Contract.WhitelistAddress(&_Junglefreaks.TransactOpts, who, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Junglefreaks *JunglefreaksTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Junglefreaks.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Junglefreaks *JunglefreaksSession) Withdraw() (*types.Transaction, error) {
	return _Junglefreaks.Contract.Withdraw(&_Junglefreaks.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Junglefreaks *JunglefreaksTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Junglefreaks.Contract.Withdraw(&_Junglefreaks.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Junglefreaks *JunglefreaksTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Junglefreaks.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Junglefreaks *JunglefreaksSession) Receive() (*types.Transaction, error) {
	return _Junglefreaks.Contract.Receive(&_Junglefreaks.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Junglefreaks *JunglefreaksTransactorSession) Receive() (*types.Transaction, error) {
	return _Junglefreaks.Contract.Receive(&_Junglefreaks.TransactOpts)
}

// JunglefreaksApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Junglefreaks contract.
type JunglefreaksApprovalIterator struct {
	Event *JunglefreaksApproval // Event containing the contract specifics and raw log

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
func (it *JunglefreaksApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunglefreaksApproval)
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
		it.Event = new(JunglefreaksApproval)
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
func (it *JunglefreaksApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunglefreaksApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunglefreaksApproval represents a Approval event raised by the Junglefreaks contract.
type JunglefreaksApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Junglefreaks *JunglefreaksFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*JunglefreaksApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Junglefreaks.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &JunglefreaksApprovalIterator{contract: _Junglefreaks.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Junglefreaks *JunglefreaksFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *JunglefreaksApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Junglefreaks.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunglefreaksApproval)
				if err := _Junglefreaks.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Junglefreaks *JunglefreaksFilterer) ParseApproval(log types.Log) (*JunglefreaksApproval, error) {
	event := new(JunglefreaksApproval)
	if err := _Junglefreaks.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunglefreaksApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Junglefreaks contract.
type JunglefreaksApprovalForAllIterator struct {
	Event *JunglefreaksApprovalForAll // Event containing the contract specifics and raw log

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
func (it *JunglefreaksApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunglefreaksApprovalForAll)
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
		it.Event = new(JunglefreaksApprovalForAll)
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
func (it *JunglefreaksApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunglefreaksApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunglefreaksApprovalForAll represents a ApprovalForAll event raised by the Junglefreaks contract.
type JunglefreaksApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Junglefreaks *JunglefreaksFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*JunglefreaksApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Junglefreaks.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &JunglefreaksApprovalForAllIterator{contract: _Junglefreaks.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Junglefreaks *JunglefreaksFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *JunglefreaksApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Junglefreaks.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunglefreaksApprovalForAll)
				if err := _Junglefreaks.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Junglefreaks *JunglefreaksFilterer) ParseApprovalForAll(log types.Log) (*JunglefreaksApprovalForAll, error) {
	event := new(JunglefreaksApprovalForAll)
	if err := _Junglefreaks.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunglefreaksOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Junglefreaks contract.
type JunglefreaksOwnershipTransferredIterator struct {
	Event *JunglefreaksOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *JunglefreaksOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunglefreaksOwnershipTransferred)
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
		it.Event = new(JunglefreaksOwnershipTransferred)
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
func (it *JunglefreaksOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunglefreaksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunglefreaksOwnershipTransferred represents a OwnershipTransferred event raised by the Junglefreaks contract.
type JunglefreaksOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Junglefreaks *JunglefreaksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JunglefreaksOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Junglefreaks.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JunglefreaksOwnershipTransferredIterator{contract: _Junglefreaks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Junglefreaks *JunglefreaksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JunglefreaksOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Junglefreaks.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunglefreaksOwnershipTransferred)
				if err := _Junglefreaks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Junglefreaks *JunglefreaksFilterer) ParseOwnershipTransferred(log types.Log) (*JunglefreaksOwnershipTransferred, error) {
	event := new(JunglefreaksOwnershipTransferred)
	if err := _Junglefreaks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunglefreaksPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Junglefreaks contract.
type JunglefreaksPausedIterator struct {
	Event *JunglefreaksPaused // Event containing the contract specifics and raw log

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
func (it *JunglefreaksPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunglefreaksPaused)
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
		it.Event = new(JunglefreaksPaused)
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
func (it *JunglefreaksPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunglefreaksPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunglefreaksPaused represents a Paused event raised by the Junglefreaks contract.
type JunglefreaksPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Junglefreaks *JunglefreaksFilterer) FilterPaused(opts *bind.FilterOpts) (*JunglefreaksPausedIterator, error) {

	logs, sub, err := _Junglefreaks.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &JunglefreaksPausedIterator{contract: _Junglefreaks.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Junglefreaks *JunglefreaksFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *JunglefreaksPaused) (event.Subscription, error) {

	logs, sub, err := _Junglefreaks.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunglefreaksPaused)
				if err := _Junglefreaks.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Junglefreaks *JunglefreaksFilterer) ParsePaused(log types.Log) (*JunglefreaksPaused, error) {
	event := new(JunglefreaksPaused)
	if err := _Junglefreaks.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunglefreaksTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Junglefreaks contract.
type JunglefreaksTransferIterator struct {
	Event *JunglefreaksTransfer // Event containing the contract specifics and raw log

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
func (it *JunglefreaksTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunglefreaksTransfer)
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
		it.Event = new(JunglefreaksTransfer)
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
func (it *JunglefreaksTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunglefreaksTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunglefreaksTransfer represents a Transfer event raised by the Junglefreaks contract.
type JunglefreaksTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Junglefreaks *JunglefreaksFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*JunglefreaksTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Junglefreaks.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &JunglefreaksTransferIterator{contract: _Junglefreaks.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Junglefreaks *JunglefreaksFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *JunglefreaksTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Junglefreaks.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunglefreaksTransfer)
				if err := _Junglefreaks.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Junglefreaks *JunglefreaksFilterer) ParseTransfer(log types.Log) (*JunglefreaksTransfer, error) {
	event := new(JunglefreaksTransfer)
	if err := _Junglefreaks.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunglefreaksUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Junglefreaks contract.
type JunglefreaksUnpausedIterator struct {
	Event *JunglefreaksUnpaused // Event containing the contract specifics and raw log

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
func (it *JunglefreaksUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunglefreaksUnpaused)
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
		it.Event = new(JunglefreaksUnpaused)
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
func (it *JunglefreaksUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunglefreaksUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunglefreaksUnpaused represents a Unpaused event raised by the Junglefreaks contract.
type JunglefreaksUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Junglefreaks *JunglefreaksFilterer) FilterUnpaused(opts *bind.FilterOpts) (*JunglefreaksUnpausedIterator, error) {

	logs, sub, err := _Junglefreaks.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &JunglefreaksUnpausedIterator{contract: _Junglefreaks.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Junglefreaks *JunglefreaksFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *JunglefreaksUnpaused) (event.Subscription, error) {

	logs, sub, err := _Junglefreaks.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunglefreaksUnpaused)
				if err := _Junglefreaks.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Junglefreaks *JunglefreaksFilterer) ParseUnpaused(log types.Log) (*JunglefreaksUnpaused, error) {
	event := new(JunglefreaksUnpaused)
	if err := _Junglefreaks.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
