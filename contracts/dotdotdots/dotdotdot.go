// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dotdotdots

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

// DotdotdotsMetaData contains all meta data concerning the Dotdotdots contract.
var DotdotdotsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_NFT_PURCHASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NFT_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcStartingIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencySetStartingIndexBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipSaleState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberOfTokensMax5\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"reserveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxTokenSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingIndexBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DotdotdotsABI is the input ABI used to generate the binding from.
// Deprecated: Use DotdotdotsMetaData.ABI instead.
var DotdotdotsABI = DotdotdotsMetaData.ABI

// Dotdotdots is an auto generated Go binding around an Ethereum contract.
type Dotdotdots struct {
	DotdotdotsCaller     // Read-only binding to the contract
	DotdotdotsTransactor // Write-only binding to the contract
	DotdotdotsFilterer   // Log filterer for contract events
}

// DotdotdotsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DotdotdotsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DotdotdotsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DotdotdotsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DotdotdotsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DotdotdotsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DotdotdotsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DotdotdotsSession struct {
	Contract     *Dotdotdots       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DotdotdotsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DotdotdotsCallerSession struct {
	Contract *DotdotdotsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DotdotdotsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DotdotdotsTransactorSession struct {
	Contract     *DotdotdotsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DotdotdotsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DotdotdotsRaw struct {
	Contract *Dotdotdots // Generic contract binding to access the raw methods on
}

// DotdotdotsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DotdotdotsCallerRaw struct {
	Contract *DotdotdotsCaller // Generic read-only contract binding to access the raw methods on
}

// DotdotdotsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DotdotdotsTransactorRaw struct {
	Contract *DotdotdotsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDotdotdots creates a new instance of Dotdotdots, bound to a specific deployed contract.
func NewDotdotdots(address common.Address, backend bind.ContractBackend) (*Dotdotdots, error) {
	contract, err := bindDotdotdots(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dotdotdots{DotdotdotsCaller: DotdotdotsCaller{contract: contract}, DotdotdotsTransactor: DotdotdotsTransactor{contract: contract}, DotdotdotsFilterer: DotdotdotsFilterer{contract: contract}}, nil
}

// NewDotdotdotsCaller creates a new read-only instance of Dotdotdots, bound to a specific deployed contract.
func NewDotdotdotsCaller(address common.Address, caller bind.ContractCaller) (*DotdotdotsCaller, error) {
	contract, err := bindDotdotdots(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DotdotdotsCaller{contract: contract}, nil
}

// NewDotdotdotsTransactor creates a new write-only instance of Dotdotdots, bound to a specific deployed contract.
func NewDotdotdotsTransactor(address common.Address, transactor bind.ContractTransactor) (*DotdotdotsTransactor, error) {
	contract, err := bindDotdotdots(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DotdotdotsTransactor{contract: contract}, nil
}

// NewDotdotdotsFilterer creates a new log filterer instance of Dotdotdots, bound to a specific deployed contract.
func NewDotdotdotsFilterer(address common.Address, filterer bind.ContractFilterer) (*DotdotdotsFilterer, error) {
	contract, err := bindDotdotdots(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DotdotdotsFilterer{contract: contract}, nil
}

// bindDotdotdots binds a generic wrapper to an already deployed contract.
func bindDotdotdots(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DotdotdotsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dotdotdots *DotdotdotsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dotdotdots.Contract.DotdotdotsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dotdotdots *DotdotdotsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotdots.Contract.DotdotdotsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dotdotdots *DotdotdotsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dotdotdots.Contract.DotdotdotsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dotdotdots *DotdotdotsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dotdotdots.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dotdotdots *DotdotdotsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotdots.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dotdotdots *DotdotdotsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dotdotdots.Contract.contract.Transact(opts, method, params...)
}

// MAXNFTPURCHASE is a free data retrieval call binding the contract method 0x020b39cc.
//
// Solidity: function MAX_NFT_PURCHASE() view returns(uint256)
func (_Dotdotdots *DotdotdotsCaller) MAXNFTPURCHASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "MAX_NFT_PURCHASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNFTPURCHASE is a free data retrieval call binding the contract method 0x020b39cc.
//
// Solidity: function MAX_NFT_PURCHASE() view returns(uint256)
func (_Dotdotdots *DotdotdotsSession) MAXNFTPURCHASE() (*big.Int, error) {
	return _Dotdotdots.Contract.MAXNFTPURCHASE(&_Dotdotdots.CallOpts)
}

// MAXNFTPURCHASE is a free data retrieval call binding the contract method 0x020b39cc.
//
// Solidity: function MAX_NFT_PURCHASE() view returns(uint256)
func (_Dotdotdots *DotdotdotsCallerSession) MAXNFTPURCHASE() (*big.Int, error) {
	return _Dotdotdots.Contract.MAXNFTPURCHASE(&_Dotdotdots.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Dotdotdots *DotdotdotsCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Dotdotdots *DotdotdotsSession) MAXSUPPLY() (*big.Int, error) {
	return _Dotdotdots.Contract.MAXSUPPLY(&_Dotdotdots.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Dotdotdots *DotdotdotsCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Dotdotdots.Contract.MAXSUPPLY(&_Dotdotdots.CallOpts)
}

// NFTPRICE is a free data retrieval call binding the contract method 0x676dd563.
//
// Solidity: function NFT_PRICE() view returns(uint256)
func (_Dotdotdots *DotdotdotsCaller) NFTPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "NFT_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NFTPRICE is a free data retrieval call binding the contract method 0x676dd563.
//
// Solidity: function NFT_PRICE() view returns(uint256)
func (_Dotdotdots *DotdotdotsSession) NFTPRICE() (*big.Int, error) {
	return _Dotdotdots.Contract.NFTPRICE(&_Dotdotdots.CallOpts)
}

// NFTPRICE is a free data retrieval call binding the contract method 0x676dd563.
//
// Solidity: function NFT_PRICE() view returns(uint256)
func (_Dotdotdots *DotdotdotsCallerSession) NFTPRICE() (*big.Int, error) {
	return _Dotdotdots.Contract.NFTPRICE(&_Dotdotdots.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dotdotdots *DotdotdotsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dotdotdots *DotdotdotsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Dotdotdots.Contract.BalanceOf(&_Dotdotdots.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dotdotdots *DotdotdotsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Dotdotdots.Contract.BalanceOf(&_Dotdotdots.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dotdotdots *DotdotdotsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dotdotdots *DotdotdotsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Dotdotdots.Contract.GetApproved(&_Dotdotdots.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dotdotdots *DotdotdotsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Dotdotdots.Contract.GetApproved(&_Dotdotdots.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dotdotdots *DotdotdotsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dotdotdots *DotdotdotsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Dotdotdots.Contract.IsApprovedForAll(&_Dotdotdots.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dotdotdots *DotdotdotsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Dotdotdots.Contract.IsApprovedForAll(&_Dotdotdots.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dotdotdots *DotdotdotsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dotdotdots *DotdotdotsSession) Name() (string, error) {
	return _Dotdotdots.Contract.Name(&_Dotdotdots.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dotdotdots *DotdotdotsCallerSession) Name() (string, error) {
	return _Dotdotdots.Contract.Name(&_Dotdotdots.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dotdotdots *DotdotdotsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dotdotdots *DotdotdotsSession) Owner() (common.Address, error) {
	return _Dotdotdots.Contract.Owner(&_Dotdotdots.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dotdotdots *DotdotdotsCallerSession) Owner() (common.Address, error) {
	return _Dotdotdots.Contract.Owner(&_Dotdotdots.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dotdotdots *DotdotdotsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dotdotdots *DotdotdotsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Dotdotdots.Contract.OwnerOf(&_Dotdotdots.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dotdotdots *DotdotdotsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Dotdotdots.Contract.OwnerOf(&_Dotdotdots.CallOpts, tokenId)
}

// SaleIsActive is a free data retrieval call binding the contract method 0xeb8d2444.
//
// Solidity: function saleIsActive() view returns(bool)
func (_Dotdotdots *DotdotdotsCaller) SaleIsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "saleIsActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleIsActive is a free data retrieval call binding the contract method 0xeb8d2444.
//
// Solidity: function saleIsActive() view returns(bool)
func (_Dotdotdots *DotdotdotsSession) SaleIsActive() (bool, error) {
	return _Dotdotdots.Contract.SaleIsActive(&_Dotdotdots.CallOpts)
}

// SaleIsActive is a free data retrieval call binding the contract method 0xeb8d2444.
//
// Solidity: function saleIsActive() view returns(bool)
func (_Dotdotdots *DotdotdotsCallerSession) SaleIsActive() (bool, error) {
	return _Dotdotdots.Contract.SaleIsActive(&_Dotdotdots.CallOpts)
}

// StartingIndex is a free data retrieval call binding the contract method 0xcb774d47.
//
// Solidity: function startingIndex() view returns(uint256)
func (_Dotdotdots *DotdotdotsCaller) StartingIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "startingIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingIndex is a free data retrieval call binding the contract method 0xcb774d47.
//
// Solidity: function startingIndex() view returns(uint256)
func (_Dotdotdots *DotdotdotsSession) StartingIndex() (*big.Int, error) {
	return _Dotdotdots.Contract.StartingIndex(&_Dotdotdots.CallOpts)
}

// StartingIndex is a free data retrieval call binding the contract method 0xcb774d47.
//
// Solidity: function startingIndex() view returns(uint256)
func (_Dotdotdots *DotdotdotsCallerSession) StartingIndex() (*big.Int, error) {
	return _Dotdotdots.Contract.StartingIndex(&_Dotdotdots.CallOpts)
}

// StartingIndexBlock is a free data retrieval call binding the contract method 0xe36d6498.
//
// Solidity: function startingIndexBlock() view returns(uint256)
func (_Dotdotdots *DotdotdotsCaller) StartingIndexBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "startingIndexBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingIndexBlock is a free data retrieval call binding the contract method 0xe36d6498.
//
// Solidity: function startingIndexBlock() view returns(uint256)
func (_Dotdotdots *DotdotdotsSession) StartingIndexBlock() (*big.Int, error) {
	return _Dotdotdots.Contract.StartingIndexBlock(&_Dotdotdots.CallOpts)
}

// StartingIndexBlock is a free data retrieval call binding the contract method 0xe36d6498.
//
// Solidity: function startingIndexBlock() view returns(uint256)
func (_Dotdotdots *DotdotdotsCallerSession) StartingIndexBlock() (*big.Int, error) {
	return _Dotdotdots.Contract.StartingIndexBlock(&_Dotdotdots.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dotdotdots *DotdotdotsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dotdotdots *DotdotdotsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dotdotdots.Contract.SupportsInterface(&_Dotdotdots.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dotdotdots *DotdotdotsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dotdotdots.Contract.SupportsInterface(&_Dotdotdots.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dotdotdots *DotdotdotsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dotdotdots *DotdotdotsSession) Symbol() (string, error) {
	return _Dotdotdots.Contract.Symbol(&_Dotdotdots.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dotdotdots *DotdotdotsCallerSession) Symbol() (string, error) {
	return _Dotdotdots.Contract.Symbol(&_Dotdotdots.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Dotdotdots *DotdotdotsCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Dotdotdots *DotdotdotsSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Dotdotdots.Contract.TokenByIndex(&_Dotdotdots.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Dotdotdots *DotdotdotsCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Dotdotdots.Contract.TokenByIndex(&_Dotdotdots.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Dotdotdots *DotdotdotsCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Dotdotdots *DotdotdotsSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Dotdotdots.Contract.TokenOfOwnerByIndex(&_Dotdotdots.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Dotdotdots *DotdotdotsCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Dotdotdots.Contract.TokenOfOwnerByIndex(&_Dotdotdots.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Dotdotdots *DotdotdotsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Dotdotdots *DotdotdotsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Dotdotdots.Contract.TokenURI(&_Dotdotdots.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Dotdotdots *DotdotdotsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Dotdotdots.Contract.TokenURI(&_Dotdotdots.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dotdotdots *DotdotdotsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dotdotdots.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dotdotdots *DotdotdotsSession) TotalSupply() (*big.Int, error) {
	return _Dotdotdots.Contract.TotalSupply(&_Dotdotdots.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dotdotdots *DotdotdotsCallerSession) TotalSupply() (*big.Int, error) {
	return _Dotdotdots.Contract.TotalSupply(&_Dotdotdots.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Dotdotdots *DotdotdotsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Dotdotdots *DotdotdotsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.Approve(&_Dotdotdots.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Dotdotdots *DotdotdotsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.Approve(&_Dotdotdots.TransactOpts, to, tokenId)
}

// CalcStartingIndex is a paid mutator transaction binding the contract method 0xe13f351a.
//
// Solidity: function calcStartingIndex() returns()
func (_Dotdotdots *DotdotdotsTransactor) CalcStartingIndex(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "calcStartingIndex")
}

// CalcStartingIndex is a paid mutator transaction binding the contract method 0xe13f351a.
//
// Solidity: function calcStartingIndex() returns()
func (_Dotdotdots *DotdotdotsSession) CalcStartingIndex() (*types.Transaction, error) {
	return _Dotdotdots.Contract.CalcStartingIndex(&_Dotdotdots.TransactOpts)
}

// CalcStartingIndex is a paid mutator transaction binding the contract method 0xe13f351a.
//
// Solidity: function calcStartingIndex() returns()
func (_Dotdotdots *DotdotdotsTransactorSession) CalcStartingIndex() (*types.Transaction, error) {
	return _Dotdotdots.Contract.CalcStartingIndex(&_Dotdotdots.TransactOpts)
}

// EmergencySetStartingIndexBlock is a paid mutator transaction binding the contract method 0x7d17fcbe.
//
// Solidity: function emergencySetStartingIndexBlock() returns()
func (_Dotdotdots *DotdotdotsTransactor) EmergencySetStartingIndexBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "emergencySetStartingIndexBlock")
}

// EmergencySetStartingIndexBlock is a paid mutator transaction binding the contract method 0x7d17fcbe.
//
// Solidity: function emergencySetStartingIndexBlock() returns()
func (_Dotdotdots *DotdotdotsSession) EmergencySetStartingIndexBlock() (*types.Transaction, error) {
	return _Dotdotdots.Contract.EmergencySetStartingIndexBlock(&_Dotdotdots.TransactOpts)
}

// EmergencySetStartingIndexBlock is a paid mutator transaction binding the contract method 0x7d17fcbe.
//
// Solidity: function emergencySetStartingIndexBlock() returns()
func (_Dotdotdots *DotdotdotsTransactorSession) EmergencySetStartingIndexBlock() (*types.Transaction, error) {
	return _Dotdotdots.Contract.EmergencySetStartingIndexBlock(&_Dotdotdots.TransactOpts)
}

// FlipSaleState is a paid mutator transaction binding the contract method 0x34918dfd.
//
// Solidity: function flipSaleState() returns()
func (_Dotdotdots *DotdotdotsTransactor) FlipSaleState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "flipSaleState")
}

// FlipSaleState is a paid mutator transaction binding the contract method 0x34918dfd.
//
// Solidity: function flipSaleState() returns()
func (_Dotdotdots *DotdotdotsSession) FlipSaleState() (*types.Transaction, error) {
	return _Dotdotdots.Contract.FlipSaleState(&_Dotdotdots.TransactOpts)
}

// FlipSaleState is a paid mutator transaction binding the contract method 0x34918dfd.
//
// Solidity: function flipSaleState() returns()
func (_Dotdotdots *DotdotdotsTransactorSession) FlipSaleState() (*types.Transaction, error) {
	return _Dotdotdots.Contract.FlipSaleState(&_Dotdotdots.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 numberOfTokensMax5) payable returns()
func (_Dotdotdots *DotdotdotsTransactor) Mint(opts *bind.TransactOpts, numberOfTokensMax5 *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "mint", numberOfTokensMax5)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 numberOfTokensMax5) payable returns()
func (_Dotdotdots *DotdotdotsSession) Mint(numberOfTokensMax5 *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.Mint(&_Dotdotdots.TransactOpts, numberOfTokensMax5)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 numberOfTokensMax5) payable returns()
func (_Dotdotdots *DotdotdotsTransactorSession) Mint(numberOfTokensMax5 *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.Mint(&_Dotdotdots.TransactOpts, numberOfTokensMax5)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dotdotdots *DotdotdotsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dotdotdots *DotdotdotsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dotdotdots.Contract.RenounceOwnership(&_Dotdotdots.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dotdotdots *DotdotdotsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dotdotdots.Contract.RenounceOwnership(&_Dotdotdots.TransactOpts)
}

// ReserveTokens is a paid mutator transaction binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 num) returns()
func (_Dotdotdots *DotdotdotsTransactor) ReserveTokens(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "reserveTokens", num)
}

// ReserveTokens is a paid mutator transaction binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 num) returns()
func (_Dotdotdots *DotdotdotsSession) ReserveTokens(num *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.ReserveTokens(&_Dotdotdots.TransactOpts, num)
}

// ReserveTokens is a paid mutator transaction binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 num) returns()
func (_Dotdotdots *DotdotdotsTransactorSession) ReserveTokens(num *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.ReserveTokens(&_Dotdotdots.TransactOpts, num)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Dotdotdots *DotdotdotsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Dotdotdots *DotdotdotsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SafeTransferFrom(&_Dotdotdots.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Dotdotdots *DotdotdotsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SafeTransferFrom(&_Dotdotdots.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Dotdotdots *DotdotdotsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Dotdotdots *DotdotdotsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SafeTransferFrom0(&_Dotdotdots.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Dotdotdots *DotdotdotsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SafeTransferFrom0(&_Dotdotdots.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dotdotdots *DotdotdotsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dotdotdots *DotdotdotsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SetApprovalForAll(&_Dotdotdots.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dotdotdots *DotdotdotsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SetApprovalForAll(&_Dotdotdots.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Dotdotdots *DotdotdotsTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "setBaseURI", baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Dotdotdots *DotdotdotsSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SetBaseURI(&_Dotdotdots.TransactOpts, baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Dotdotdots *DotdotdotsTransactorSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SetBaseURI(&_Dotdotdots.TransactOpts, baseURI_)
}

// SetMaxTokenSupply is a paid mutator transaction binding the contract method 0xb07ed982.
//
// Solidity: function setMaxTokenSupply(uint256 maxSupply) returns()
func (_Dotdotdots *DotdotdotsTransactor) SetMaxTokenSupply(opts *bind.TransactOpts, maxSupply *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "setMaxTokenSupply", maxSupply)
}

// SetMaxTokenSupply is a paid mutator transaction binding the contract method 0xb07ed982.
//
// Solidity: function setMaxTokenSupply(uint256 maxSupply) returns()
func (_Dotdotdots *DotdotdotsSession) SetMaxTokenSupply(maxSupply *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SetMaxTokenSupply(&_Dotdotdots.TransactOpts, maxSupply)
}

// SetMaxTokenSupply is a paid mutator transaction binding the contract method 0xb07ed982.
//
// Solidity: function setMaxTokenSupply(uint256 maxSupply) returns()
func (_Dotdotdots *DotdotdotsTransactorSession) SetMaxTokenSupply(maxSupply *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.SetMaxTokenSupply(&_Dotdotdots.TransactOpts, maxSupply)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Dotdotdots *DotdotdotsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Dotdotdots *DotdotdotsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.TransferFrom(&_Dotdotdots.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Dotdotdots *DotdotdotsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dotdotdots.Contract.TransferFrom(&_Dotdotdots.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dotdotdots *DotdotdotsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dotdotdots *DotdotdotsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dotdotdots.Contract.TransferOwnership(&_Dotdotdots.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dotdotdots *DotdotdotsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dotdotdots.Contract.TransferOwnership(&_Dotdotdots.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dotdotdots *DotdotdotsTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dotdotdots.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dotdotdots *DotdotdotsSession) Withdraw() (*types.Transaction, error) {
	return _Dotdotdots.Contract.Withdraw(&_Dotdotdots.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dotdotdots *DotdotdotsTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Dotdotdots.Contract.Withdraw(&_Dotdotdots.TransactOpts)
}

// DotdotdotsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Dotdotdots contract.
type DotdotdotsApprovalIterator struct {
	Event *DotdotdotsApproval // Event containing the contract specifics and raw log

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
func (it *DotdotdotsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DotdotdotsApproval)
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
		it.Event = new(DotdotdotsApproval)
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
func (it *DotdotdotsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DotdotdotsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DotdotdotsApproval represents a Approval event raised by the Dotdotdots contract.
type DotdotdotsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dotdotdots *DotdotdotsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DotdotdotsApprovalIterator, error) {

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

	logs, sub, err := _Dotdotdots.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DotdotdotsApprovalIterator{contract: _Dotdotdots.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dotdotdots *DotdotdotsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DotdotdotsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Dotdotdots.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DotdotdotsApproval)
				if err := _Dotdotdots.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Dotdotdots *DotdotdotsFilterer) ParseApproval(log types.Log) (*DotdotdotsApproval, error) {
	event := new(DotdotdotsApproval)
	if err := _Dotdotdots.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DotdotdotsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Dotdotdots contract.
type DotdotdotsApprovalForAllIterator struct {
	Event *DotdotdotsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DotdotdotsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DotdotdotsApprovalForAll)
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
		it.Event = new(DotdotdotsApprovalForAll)
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
func (it *DotdotdotsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DotdotdotsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DotdotdotsApprovalForAll represents a ApprovalForAll event raised by the Dotdotdots contract.
type DotdotdotsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dotdotdots *DotdotdotsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DotdotdotsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Dotdotdots.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DotdotdotsApprovalForAllIterator{contract: _Dotdotdots.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dotdotdots *DotdotdotsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DotdotdotsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Dotdotdots.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DotdotdotsApprovalForAll)
				if err := _Dotdotdots.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Dotdotdots *DotdotdotsFilterer) ParseApprovalForAll(log types.Log) (*DotdotdotsApprovalForAll, error) {
	event := new(DotdotdotsApprovalForAll)
	if err := _Dotdotdots.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DotdotdotsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dotdotdots contract.
type DotdotdotsOwnershipTransferredIterator struct {
	Event *DotdotdotsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DotdotdotsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DotdotdotsOwnershipTransferred)
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
		it.Event = new(DotdotdotsOwnershipTransferred)
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
func (it *DotdotdotsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DotdotdotsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DotdotdotsOwnershipTransferred represents a OwnershipTransferred event raised by the Dotdotdots contract.
type DotdotdotsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dotdotdots *DotdotdotsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DotdotdotsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dotdotdots.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DotdotdotsOwnershipTransferredIterator{contract: _Dotdotdots.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dotdotdots *DotdotdotsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DotdotdotsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dotdotdots.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DotdotdotsOwnershipTransferred)
				if err := _Dotdotdots.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dotdotdots *DotdotdotsFilterer) ParseOwnershipTransferred(log types.Log) (*DotdotdotsOwnershipTransferred, error) {
	event := new(DotdotdotsOwnershipTransferred)
	if err := _Dotdotdots.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DotdotdotsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dotdotdots contract.
type DotdotdotsTransferIterator struct {
	Event *DotdotdotsTransfer // Event containing the contract specifics and raw log

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
func (it *DotdotdotsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DotdotdotsTransfer)
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
		it.Event = new(DotdotdotsTransfer)
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
func (it *DotdotdotsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DotdotdotsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DotdotdotsTransfer represents a Transfer event raised by the Dotdotdots contract.
type DotdotdotsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dotdotdots *DotdotdotsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DotdotdotsTransferIterator, error) {

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

	logs, sub, err := _Dotdotdots.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DotdotdotsTransferIterator{contract: _Dotdotdots.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dotdotdots *DotdotdotsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DotdotdotsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Dotdotdots.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DotdotdotsTransfer)
				if err := _Dotdotdots.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Dotdotdots *DotdotdotsFilterer) ParseTransfer(log types.Log) (*DotdotdotsTransfer, error) {
	event := new(DotdotdotsTransfer)
	if err := _Dotdotdots.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
