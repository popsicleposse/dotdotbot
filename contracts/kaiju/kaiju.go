// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kaiju

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

// KaijuMetaData contains all meta data concerning the Kaiju contract.
var KaijuMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genCount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kaijuId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"kaijuBio\",\"type\":\"string\"}],\"name\":\"BioChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kaijuId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"parent1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"parent2\",\"type\":\"uint256\"}],\"name\":\"KaijuFusion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kaijuId\",\"type\":\"uint256\"}],\"name\":\"KaijuRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kaijuId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"kaijuName\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIO_CHANGE_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FUSION_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KKold\",\"outputs\":[{\"internalType\":\"contractIKKold\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME_CHANGE_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RWaste\",\"outputs\":[{\"internalType\":\"contractIRWaste\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kaijuTokens\",\"type\":\"uint256[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"babyCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"babyKaiju\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceGenesis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kaijuId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newBio\",\"type\":\"string\"}],\"name\":\"changeBio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kaijuId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"changeName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"presaleAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"editPresale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"parent1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parent2\",\"type\":\"uint256\"}],\"name\":\"fusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kaijuData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bio\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberOfMints\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberOfMints\",\"type\":\"uint256\"}],\"name\":\"mintPresale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"presaleWhitelist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kaijuId\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"IKKoldAddress\",\"type\":\"address\"}],\"name\":\"setKKold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rWasteAddress\",\"type\":\"address\"}],\"name\":\"setRadioactiveWaste\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePresale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"walletOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KaijuABI is the input ABI used to generate the binding from.
// Deprecated: Use KaijuMetaData.ABI instead.
var KaijuABI = KaijuMetaData.ABI

// Kaiju is an auto generated Go binding around an Ethereum contract.
type Kaiju struct {
	KaijuCaller     // Read-only binding to the contract
	KaijuTransactor // Write-only binding to the contract
	KaijuFilterer   // Log filterer for contract events
}

// KaijuCaller is an auto generated read-only Go binding around an Ethereum contract.
type KaijuCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KaijuTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KaijuTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KaijuFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KaijuFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KaijuSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KaijuSession struct {
	Contract     *Kaiju            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KaijuCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KaijuCallerSession struct {
	Contract *KaijuCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KaijuTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KaijuTransactorSession struct {
	Contract     *KaijuTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KaijuRaw is an auto generated low-level Go binding around an Ethereum contract.
type KaijuRaw struct {
	Contract *Kaiju // Generic contract binding to access the raw methods on
}

// KaijuCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KaijuCallerRaw struct {
	Contract *KaijuCaller // Generic read-only contract binding to access the raw methods on
}

// KaijuTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KaijuTransactorRaw struct {
	Contract *KaijuTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKaiju creates a new instance of Kaiju, bound to a specific deployed contract.
func NewKaiju(address common.Address, backend bind.ContractBackend) (*Kaiju, error) {
	contract, err := bindKaiju(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kaiju{KaijuCaller: KaijuCaller{contract: contract}, KaijuTransactor: KaijuTransactor{contract: contract}, KaijuFilterer: KaijuFilterer{contract: contract}}, nil
}

// NewKaijuCaller creates a new read-only instance of Kaiju, bound to a specific deployed contract.
func NewKaijuCaller(address common.Address, caller bind.ContractCaller) (*KaijuCaller, error) {
	contract, err := bindKaiju(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KaijuCaller{contract: contract}, nil
}

// NewKaijuTransactor creates a new write-only instance of Kaiju, bound to a specific deployed contract.
func NewKaijuTransactor(address common.Address, transactor bind.ContractTransactor) (*KaijuTransactor, error) {
	contract, err := bindKaiju(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KaijuTransactor{contract: contract}, nil
}

// NewKaijuFilterer creates a new log filterer instance of Kaiju, bound to a specific deployed contract.
func NewKaijuFilterer(address common.Address, filterer bind.ContractFilterer) (*KaijuFilterer, error) {
	contract, err := bindKaiju(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KaijuFilterer{contract: contract}, nil
}

// bindKaiju binds a generic wrapper to an already deployed contract.
func bindKaiju(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KaijuABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kaiju *KaijuRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kaiju.Contract.KaijuCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kaiju *KaijuRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kaiju.Contract.KaijuTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kaiju *KaijuRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kaiju.Contract.KaijuTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kaiju *KaijuCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kaiju.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kaiju *KaijuTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kaiju.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kaiju *KaijuTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kaiju.Contract.contract.Transact(opts, method, params...)
}

// BIOCHANGEPRICE is a free data retrieval call binding the contract method 0x74151be0.
//
// Solidity: function BIO_CHANGE_PRICE() view returns(uint256)
func (_Kaiju *KaijuCaller) BIOCHANGEPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "BIO_CHANGE_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BIOCHANGEPRICE is a free data retrieval call binding the contract method 0x74151be0.
//
// Solidity: function BIO_CHANGE_PRICE() view returns(uint256)
func (_Kaiju *KaijuSession) BIOCHANGEPRICE() (*big.Int, error) {
	return _Kaiju.Contract.BIOCHANGEPRICE(&_Kaiju.CallOpts)
}

// BIOCHANGEPRICE is a free data retrieval call binding the contract method 0x74151be0.
//
// Solidity: function BIO_CHANGE_PRICE() view returns(uint256)
func (_Kaiju *KaijuCallerSession) BIOCHANGEPRICE() (*big.Int, error) {
	return _Kaiju.Contract.BIOCHANGEPRICE(&_Kaiju.CallOpts)
}

// FUSIONPRICE is a free data retrieval call binding the contract method 0xe3f5967e.
//
// Solidity: function FUSION_PRICE() view returns(uint256)
func (_Kaiju *KaijuCaller) FUSIONPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "FUSION_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUSIONPRICE is a free data retrieval call binding the contract method 0xe3f5967e.
//
// Solidity: function FUSION_PRICE() view returns(uint256)
func (_Kaiju *KaijuSession) FUSIONPRICE() (*big.Int, error) {
	return _Kaiju.Contract.FUSIONPRICE(&_Kaiju.CallOpts)
}

// FUSIONPRICE is a free data retrieval call binding the contract method 0xe3f5967e.
//
// Solidity: function FUSION_PRICE() view returns(uint256)
func (_Kaiju *KaijuCallerSession) FUSIONPRICE() (*big.Int, error) {
	return _Kaiju.Contract.FUSIONPRICE(&_Kaiju.CallOpts)
}

// KKold is a free data retrieval call binding the contract method 0x46feb710.
//
// Solidity: function KKold() view returns(address)
func (_Kaiju *KaijuCaller) KKold(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "KKold")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KKold is a free data retrieval call binding the contract method 0x46feb710.
//
// Solidity: function KKold() view returns(address)
func (_Kaiju *KaijuSession) KKold() (common.Address, error) {
	return _Kaiju.Contract.KKold(&_Kaiju.CallOpts)
}

// KKold is a free data retrieval call binding the contract method 0x46feb710.
//
// Solidity: function KKold() view returns(address)
func (_Kaiju *KaijuCallerSession) KKold() (common.Address, error) {
	return _Kaiju.Contract.KKold(&_Kaiju.CallOpts)
}

// NAMECHANGEPRICE is a free data retrieval call binding the contract method 0x54b6f161.
//
// Solidity: function NAME_CHANGE_PRICE() view returns(uint256)
func (_Kaiju *KaijuCaller) NAMECHANGEPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "NAME_CHANGE_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NAMECHANGEPRICE is a free data retrieval call binding the contract method 0x54b6f161.
//
// Solidity: function NAME_CHANGE_PRICE() view returns(uint256)
func (_Kaiju *KaijuSession) NAMECHANGEPRICE() (*big.Int, error) {
	return _Kaiju.Contract.NAMECHANGEPRICE(&_Kaiju.CallOpts)
}

// NAMECHANGEPRICE is a free data retrieval call binding the contract method 0x54b6f161.
//
// Solidity: function NAME_CHANGE_PRICE() view returns(uint256)
func (_Kaiju *KaijuCallerSession) NAMECHANGEPRICE() (*big.Int, error) {
	return _Kaiju.Contract.NAMECHANGEPRICE(&_Kaiju.CallOpts)
}

// RWaste is a free data retrieval call binding the contract method 0xcbce9a23.
//
// Solidity: function RWaste() view returns(address)
func (_Kaiju *KaijuCaller) RWaste(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "RWaste")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RWaste is a free data retrieval call binding the contract method 0xcbce9a23.
//
// Solidity: function RWaste() view returns(address)
func (_Kaiju *KaijuSession) RWaste() (common.Address, error) {
	return _Kaiju.Contract.RWaste(&_Kaiju.CallOpts)
}

// RWaste is a free data retrieval call binding the contract method 0xcbce9a23.
//
// Solidity: function RWaste() view returns(address)
func (_Kaiju *KaijuCallerSession) RWaste() (common.Address, error) {
	return _Kaiju.Contract.RWaste(&_Kaiju.CallOpts)
}

// BabyCount is a free data retrieval call binding the contract method 0x539cb1e7.
//
// Solidity: function babyCount() view returns(uint256)
func (_Kaiju *KaijuCaller) BabyCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "babyCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BabyCount is a free data retrieval call binding the contract method 0x539cb1e7.
//
// Solidity: function babyCount() view returns(uint256)
func (_Kaiju *KaijuSession) BabyCount() (*big.Int, error) {
	return _Kaiju.Contract.BabyCount(&_Kaiju.CallOpts)
}

// BabyCount is a free data retrieval call binding the contract method 0x539cb1e7.
//
// Solidity: function babyCount() view returns(uint256)
func (_Kaiju *KaijuCallerSession) BabyCount() (*big.Int, error) {
	return _Kaiju.Contract.BabyCount(&_Kaiju.CallOpts)
}

// BabyKaiju is a free data retrieval call binding the contract method 0x50d1ca19.
//
// Solidity: function babyKaiju(uint256 ) view returns(uint256)
func (_Kaiju *KaijuCaller) BabyKaiju(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "babyKaiju", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BabyKaiju is a free data retrieval call binding the contract method 0x50d1ca19.
//
// Solidity: function babyKaiju(uint256 ) view returns(uint256)
func (_Kaiju *KaijuSession) BabyKaiju(arg0 *big.Int) (*big.Int, error) {
	return _Kaiju.Contract.BabyKaiju(&_Kaiju.CallOpts, arg0)
}

// BabyKaiju is a free data retrieval call binding the contract method 0x50d1ca19.
//
// Solidity: function babyKaiju(uint256 ) view returns(uint256)
func (_Kaiju *KaijuCallerSession) BabyKaiju(arg0 *big.Int) (*big.Int, error) {
	return _Kaiju.Contract.BabyKaiju(&_Kaiju.CallOpts, arg0)
}

// BalanceGenesis is a free data retrieval call binding the contract method 0xe2b26b15.
//
// Solidity: function balanceGenesis(address ) view returns(uint256)
func (_Kaiju *KaijuCaller) BalanceGenesis(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "balanceGenesis", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceGenesis is a free data retrieval call binding the contract method 0xe2b26b15.
//
// Solidity: function balanceGenesis(address ) view returns(uint256)
func (_Kaiju *KaijuSession) BalanceGenesis(arg0 common.Address) (*big.Int, error) {
	return _Kaiju.Contract.BalanceGenesis(&_Kaiju.CallOpts, arg0)
}

// BalanceGenesis is a free data retrieval call binding the contract method 0xe2b26b15.
//
// Solidity: function balanceGenesis(address ) view returns(uint256)
func (_Kaiju *KaijuCallerSession) BalanceGenesis(arg0 common.Address) (*big.Int, error) {
	return _Kaiju.Contract.BalanceGenesis(&_Kaiju.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Kaiju *KaijuCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Kaiju *KaijuSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Kaiju.Contract.BalanceOf(&_Kaiju.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Kaiju *KaijuCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Kaiju.Contract.BalanceOf(&_Kaiju.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Kaiju *KaijuCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Kaiju *KaijuSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Kaiju.Contract.GetApproved(&_Kaiju.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Kaiju *KaijuCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Kaiju.Contract.GetApproved(&_Kaiju.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Kaiju *KaijuCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Kaiju *KaijuSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Kaiju.Contract.IsApprovedForAll(&_Kaiju.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Kaiju *KaijuCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Kaiju.Contract.IsApprovedForAll(&_Kaiju.CallOpts, owner, operator)
}

// KaijuData is a free data retrieval call binding the contract method 0x4b2e0fee.
//
// Solidity: function kaijuData(uint256 ) view returns(string name, string bio)
func (_Kaiju *KaijuCaller) KaijuData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name string
	Bio  string
}, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "kaijuData", arg0)

	outstruct := new(struct {
		Name string
		Bio  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Bio = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// KaijuData is a free data retrieval call binding the contract method 0x4b2e0fee.
//
// Solidity: function kaijuData(uint256 ) view returns(string name, string bio)
func (_Kaiju *KaijuSession) KaijuData(arg0 *big.Int) (struct {
	Name string
	Bio  string
}, error) {
	return _Kaiju.Contract.KaijuData(&_Kaiju.CallOpts, arg0)
}

// KaijuData is a free data retrieval call binding the contract method 0x4b2e0fee.
//
// Solidity: function kaijuData(uint256 ) view returns(string name, string bio)
func (_Kaiju *KaijuCallerSession) KaijuData(arg0 *big.Int) (struct {
	Name string
	Bio  string
}, error) {
	return _Kaiju.Contract.KaijuData(&_Kaiju.CallOpts, arg0)
}

// MaxGenCount is a free data retrieval call binding the contract method 0x3299c120.
//
// Solidity: function maxGenCount() view returns(uint256)
func (_Kaiju *KaijuCaller) MaxGenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "maxGenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGenCount is a free data retrieval call binding the contract method 0x3299c120.
//
// Solidity: function maxGenCount() view returns(uint256)
func (_Kaiju *KaijuSession) MaxGenCount() (*big.Int, error) {
	return _Kaiju.Contract.MaxGenCount(&_Kaiju.CallOpts)
}

// MaxGenCount is a free data retrieval call binding the contract method 0x3299c120.
//
// Solidity: function maxGenCount() view returns(uint256)
func (_Kaiju *KaijuCallerSession) MaxGenCount() (*big.Int, error) {
	return _Kaiju.Contract.MaxGenCount(&_Kaiju.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Kaiju *KaijuCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Kaiju *KaijuSession) MaxSupply() (*big.Int, error) {
	return _Kaiju.Contract.MaxSupply(&_Kaiju.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Kaiju *KaijuCallerSession) MaxSupply() (*big.Int, error) {
	return _Kaiju.Contract.MaxSupply(&_Kaiju.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kaiju *KaijuCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kaiju *KaijuSession) Name() (string, error) {
	return _Kaiju.Contract.Name(&_Kaiju.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kaiju *KaijuCallerSession) Name() (string, error) {
	return _Kaiju.Contract.Name(&_Kaiju.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kaiju *KaijuCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kaiju *KaijuSession) Owner() (common.Address, error) {
	return _Kaiju.Contract.Owner(&_Kaiju.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kaiju *KaijuCallerSession) Owner() (common.Address, error) {
	return _Kaiju.Contract.Owner(&_Kaiju.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Kaiju *KaijuCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Kaiju *KaijuSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Kaiju.Contract.OwnerOf(&_Kaiju.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Kaiju *KaijuCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Kaiju.Contract.OwnerOf(&_Kaiju.CallOpts, tokenId)
}

// PresaleActive is a free data retrieval call binding the contract method 0x53135ca0.
//
// Solidity: function presaleActive() view returns(bool)
func (_Kaiju *KaijuCaller) PresaleActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "presaleActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PresaleActive is a free data retrieval call binding the contract method 0x53135ca0.
//
// Solidity: function presaleActive() view returns(bool)
func (_Kaiju *KaijuSession) PresaleActive() (bool, error) {
	return _Kaiju.Contract.PresaleActive(&_Kaiju.CallOpts)
}

// PresaleActive is a free data retrieval call binding the contract method 0x53135ca0.
//
// Solidity: function presaleActive() view returns(bool)
func (_Kaiju *KaijuCallerSession) PresaleActive() (bool, error) {
	return _Kaiju.Contract.PresaleActive(&_Kaiju.CallOpts)
}

// PresaleWhitelist is a free data retrieval call binding the contract method 0xeb8835ab.
//
// Solidity: function presaleWhitelist(address ) view returns(uint256)
func (_Kaiju *KaijuCaller) PresaleWhitelist(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "presaleWhitelist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleWhitelist is a free data retrieval call binding the contract method 0xeb8835ab.
//
// Solidity: function presaleWhitelist(address ) view returns(uint256)
func (_Kaiju *KaijuSession) PresaleWhitelist(arg0 common.Address) (*big.Int, error) {
	return _Kaiju.Contract.PresaleWhitelist(&_Kaiju.CallOpts, arg0)
}

// PresaleWhitelist is a free data retrieval call binding the contract method 0xeb8835ab.
//
// Solidity: function presaleWhitelist(address ) view returns(uint256)
func (_Kaiju *KaijuCallerSession) PresaleWhitelist(arg0 common.Address) (*big.Int, error) {
	return _Kaiju.Contract.PresaleWhitelist(&_Kaiju.CallOpts, arg0)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Kaiju *KaijuCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Kaiju *KaijuSession) Price() (*big.Int, error) {
	return _Kaiju.Contract.Price(&_Kaiju.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Kaiju *KaijuCallerSession) Price() (*big.Int, error) {
	return _Kaiju.Contract.Price(&_Kaiju.CallOpts)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_Kaiju *KaijuCaller) SaleActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "saleActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_Kaiju *KaijuSession) SaleActive() (bool, error) {
	return _Kaiju.Contract.SaleActive(&_Kaiju.CallOpts)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_Kaiju *KaijuCallerSession) SaleActive() (bool, error) {
	return _Kaiju.Contract.SaleActive(&_Kaiju.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kaiju *KaijuCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kaiju *KaijuSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Kaiju.Contract.SupportsInterface(&_Kaiju.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kaiju *KaijuCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Kaiju.Contract.SupportsInterface(&_Kaiju.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kaiju *KaijuCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kaiju *KaijuSession) Symbol() (string, error) {
	return _Kaiju.Contract.Symbol(&_Kaiju.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kaiju *KaijuCallerSession) Symbol() (string, error) {
	return _Kaiju.Contract.Symbol(&_Kaiju.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Kaiju *KaijuCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Kaiju *KaijuSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Kaiju.Contract.TokenByIndex(&_Kaiju.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Kaiju *KaijuCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Kaiju.Contract.TokenByIndex(&_Kaiju.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Kaiju *KaijuCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Kaiju *KaijuSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Kaiju.Contract.TokenOfOwnerByIndex(&_Kaiju.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Kaiju *KaijuCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Kaiju.Contract.TokenOfOwnerByIndex(&_Kaiju.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Kaiju *KaijuCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Kaiju *KaijuSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Kaiju.Contract.TokenURI(&_Kaiju.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Kaiju *KaijuCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Kaiju.Contract.TokenURI(&_Kaiju.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kaiju *KaijuCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kaiju *KaijuSession) TotalSupply() (*big.Int, error) {
	return _Kaiju.Contract.TotalSupply(&_Kaiju.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kaiju *KaijuCallerSession) TotalSupply() (*big.Int, error) {
	return _Kaiju.Contract.TotalSupply(&_Kaiju.CallOpts)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address owner) view returns(uint256[])
func (_Kaiju *KaijuCaller) WalletOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Kaiju.contract.Call(opts, &out, "walletOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address owner) view returns(uint256[])
func (_Kaiju *KaijuSession) WalletOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Kaiju.Contract.WalletOfOwner(&_Kaiju.CallOpts, owner)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address owner) view returns(uint256[])
func (_Kaiju *KaijuCallerSession) WalletOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Kaiju.Contract.WalletOfOwner(&_Kaiju.CallOpts, owner)
}

// Airdrop is a paid mutator transaction binding the contract method 0x17428b4e.
//
// Solidity: function airdrop(uint256[] kaijuTokens) returns()
func (_Kaiju *KaijuTransactor) Airdrop(opts *bind.TransactOpts, kaijuTokens []*big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "airdrop", kaijuTokens)
}

// Airdrop is a paid mutator transaction binding the contract method 0x17428b4e.
//
// Solidity: function airdrop(uint256[] kaijuTokens) returns()
func (_Kaiju *KaijuSession) Airdrop(kaijuTokens []*big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Airdrop(&_Kaiju.TransactOpts, kaijuTokens)
}

// Airdrop is a paid mutator transaction binding the contract method 0x17428b4e.
//
// Solidity: function airdrop(uint256[] kaijuTokens) returns()
func (_Kaiju *KaijuTransactorSession) Airdrop(kaijuTokens []*big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Airdrop(&_Kaiju.TransactOpts, kaijuTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Kaiju *KaijuTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Kaiju *KaijuSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Approve(&_Kaiju.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Kaiju *KaijuTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Approve(&_Kaiju.TransactOpts, to, tokenId)
}

// ChangeBio is a paid mutator transaction binding the contract method 0x4d426528.
//
// Solidity: function changeBio(uint256 kaijuId, string newBio) returns()
func (_Kaiju *KaijuTransactor) ChangeBio(opts *bind.TransactOpts, kaijuId *big.Int, newBio string) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "changeBio", kaijuId, newBio)
}

// ChangeBio is a paid mutator transaction binding the contract method 0x4d426528.
//
// Solidity: function changeBio(uint256 kaijuId, string newBio) returns()
func (_Kaiju *KaijuSession) ChangeBio(kaijuId *big.Int, newBio string) (*types.Transaction, error) {
	return _Kaiju.Contract.ChangeBio(&_Kaiju.TransactOpts, kaijuId, newBio)
}

// ChangeBio is a paid mutator transaction binding the contract method 0x4d426528.
//
// Solidity: function changeBio(uint256 kaijuId, string newBio) returns()
func (_Kaiju *KaijuTransactorSession) ChangeBio(kaijuId *big.Int, newBio string) (*types.Transaction, error) {
	return _Kaiju.Contract.ChangeBio(&_Kaiju.TransactOpts, kaijuId, newBio)
}

// ChangeName is a paid mutator transaction binding the contract method 0xc39cbef1.
//
// Solidity: function changeName(uint256 kaijuId, string newName) returns()
func (_Kaiju *KaijuTransactor) ChangeName(opts *bind.TransactOpts, kaijuId *big.Int, newName string) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "changeName", kaijuId, newName)
}

// ChangeName is a paid mutator transaction binding the contract method 0xc39cbef1.
//
// Solidity: function changeName(uint256 kaijuId, string newName) returns()
func (_Kaiju *KaijuSession) ChangeName(kaijuId *big.Int, newName string) (*types.Transaction, error) {
	return _Kaiju.Contract.ChangeName(&_Kaiju.TransactOpts, kaijuId, newName)
}

// ChangeName is a paid mutator transaction binding the contract method 0xc39cbef1.
//
// Solidity: function changeName(uint256 kaijuId, string newName) returns()
func (_Kaiju *KaijuTransactorSession) ChangeName(kaijuId *big.Int, newName string) (*types.Transaction, error) {
	return _Kaiju.Contract.ChangeName(&_Kaiju.TransactOpts, kaijuId, newName)
}

// EditPresale is a paid mutator transaction binding the contract method 0x26ed7155.
//
// Solidity: function editPresale(address[] presaleAddresses, uint256[] amount) returns()
func (_Kaiju *KaijuTransactor) EditPresale(opts *bind.TransactOpts, presaleAddresses []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "editPresale", presaleAddresses, amount)
}

// EditPresale is a paid mutator transaction binding the contract method 0x26ed7155.
//
// Solidity: function editPresale(address[] presaleAddresses, uint256[] amount) returns()
func (_Kaiju *KaijuSession) EditPresale(presaleAddresses []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.EditPresale(&_Kaiju.TransactOpts, presaleAddresses, amount)
}

// EditPresale is a paid mutator transaction binding the contract method 0x26ed7155.
//
// Solidity: function editPresale(address[] presaleAddresses, uint256[] amount) returns()
func (_Kaiju *KaijuTransactorSession) EditPresale(presaleAddresses []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.EditPresale(&_Kaiju.TransactOpts, presaleAddresses, amount)
}

// Fusion is a paid mutator transaction binding the contract method 0x65a8a037.
//
// Solidity: function fusion(uint256 parent1, uint256 parent2) returns()
func (_Kaiju *KaijuTransactor) Fusion(opts *bind.TransactOpts, parent1 *big.Int, parent2 *big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "fusion", parent1, parent2)
}

// Fusion is a paid mutator transaction binding the contract method 0x65a8a037.
//
// Solidity: function fusion(uint256 parent1, uint256 parent2) returns()
func (_Kaiju *KaijuSession) Fusion(parent1 *big.Int, parent2 *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Fusion(&_Kaiju.TransactOpts, parent1, parent2)
}

// Fusion is a paid mutator transaction binding the contract method 0x65a8a037.
//
// Solidity: function fusion(uint256 parent1, uint256 parent2) returns()
func (_Kaiju *KaijuTransactorSession) Fusion(parent1 *big.Int, parent2 *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Fusion(&_Kaiju.TransactOpts, parent1, parent2)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 numberOfMints) payable returns()
func (_Kaiju *KaijuTransactor) Mint(opts *bind.TransactOpts, numberOfMints *big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "mint", numberOfMints)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 numberOfMints) payable returns()
func (_Kaiju *KaijuSession) Mint(numberOfMints *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Mint(&_Kaiju.TransactOpts, numberOfMints)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 numberOfMints) payable returns()
func (_Kaiju *KaijuTransactorSession) Mint(numberOfMints *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Mint(&_Kaiju.TransactOpts, numberOfMints)
}

// MintPresale is a paid mutator transaction binding the contract method 0xf759867a.
//
// Solidity: function mintPresale(uint256 numberOfMints) payable returns()
func (_Kaiju *KaijuTransactor) MintPresale(opts *bind.TransactOpts, numberOfMints *big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "mintPresale", numberOfMints)
}

// MintPresale is a paid mutator transaction binding the contract method 0xf759867a.
//
// Solidity: function mintPresale(uint256 numberOfMints) payable returns()
func (_Kaiju *KaijuSession) MintPresale(numberOfMints *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.MintPresale(&_Kaiju.TransactOpts, numberOfMints)
}

// MintPresale is a paid mutator transaction binding the contract method 0xf759867a.
//
// Solidity: function mintPresale(uint256 numberOfMints) payable returns()
func (_Kaiju *KaijuTransactorSession) MintPresale(numberOfMints *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.MintPresale(&_Kaiju.TransactOpts, numberOfMints)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kaiju *KaijuTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kaiju *KaijuSession) RenounceOwnership() (*types.Transaction, error) {
	return _Kaiju.Contract.RenounceOwnership(&_Kaiju.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kaiju *KaijuTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Kaiju.Contract.RenounceOwnership(&_Kaiju.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0xc2ca0ac5.
//
// Solidity: function reveal(uint256 kaijuId) returns()
func (_Kaiju *KaijuTransactor) Reveal(opts *bind.TransactOpts, kaijuId *big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "reveal", kaijuId)
}

// Reveal is a paid mutator transaction binding the contract method 0xc2ca0ac5.
//
// Solidity: function reveal(uint256 kaijuId) returns()
func (_Kaiju *KaijuSession) Reveal(kaijuId *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Reveal(&_Kaiju.TransactOpts, kaijuId)
}

// Reveal is a paid mutator transaction binding the contract method 0xc2ca0ac5.
//
// Solidity: function reveal(uint256 kaijuId) returns()
func (_Kaiju *KaijuTransactorSession) Reveal(kaijuId *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.Reveal(&_Kaiju.TransactOpts, kaijuId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Kaiju *KaijuTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Kaiju *KaijuSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.SafeTransferFrom(&_Kaiju.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Kaiju *KaijuTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.SafeTransferFrom(&_Kaiju.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Kaiju *KaijuTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Kaiju *KaijuSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Kaiju.Contract.SafeTransferFrom0(&_Kaiju.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Kaiju *KaijuTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Kaiju.Contract.SafeTransferFrom0(&_Kaiju.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Kaiju *KaijuTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Kaiju *KaijuSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Kaiju.Contract.SetApprovalForAll(&_Kaiju.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Kaiju *KaijuTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Kaiju.Contract.SetApprovalForAll(&_Kaiju.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_Kaiju *KaijuTransactor) SetBaseURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "setBaseURI", uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_Kaiju *KaijuSession) SetBaseURI(uri string) (*types.Transaction, error) {
	return _Kaiju.Contract.SetBaseURI(&_Kaiju.TransactOpts, uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_Kaiju *KaijuTransactorSession) SetBaseURI(uri string) (*types.Transaction, error) {
	return _Kaiju.Contract.SetBaseURI(&_Kaiju.TransactOpts, uri)
}

// SetKKold is a paid mutator transaction binding the contract method 0xb1a5fee1.
//
// Solidity: function setKKold(address IKKoldAddress) returns()
func (_Kaiju *KaijuTransactor) SetKKold(opts *bind.TransactOpts, IKKoldAddress common.Address) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "setKKold", IKKoldAddress)
}

// SetKKold is a paid mutator transaction binding the contract method 0xb1a5fee1.
//
// Solidity: function setKKold(address IKKoldAddress) returns()
func (_Kaiju *KaijuSession) SetKKold(IKKoldAddress common.Address) (*types.Transaction, error) {
	return _Kaiju.Contract.SetKKold(&_Kaiju.TransactOpts, IKKoldAddress)
}

// SetKKold is a paid mutator transaction binding the contract method 0xb1a5fee1.
//
// Solidity: function setKKold(address IKKoldAddress) returns()
func (_Kaiju *KaijuTransactorSession) SetKKold(IKKoldAddress common.Address) (*types.Transaction, error) {
	return _Kaiju.Contract.SetKKold(&_Kaiju.TransactOpts, IKKoldAddress)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 newPrice) returns()
func (_Kaiju *KaijuTransactor) SetPrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "setPrice", newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 newPrice) returns()
func (_Kaiju *KaijuSession) SetPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.SetPrice(&_Kaiju.TransactOpts, newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 newPrice) returns()
func (_Kaiju *KaijuTransactorSession) SetPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.SetPrice(&_Kaiju.TransactOpts, newPrice)
}

// SetRadioactiveWaste is a paid mutator transaction binding the contract method 0x52d2d249.
//
// Solidity: function setRadioactiveWaste(address rWasteAddress) returns()
func (_Kaiju *KaijuTransactor) SetRadioactiveWaste(opts *bind.TransactOpts, rWasteAddress common.Address) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "setRadioactiveWaste", rWasteAddress)
}

// SetRadioactiveWaste is a paid mutator transaction binding the contract method 0x52d2d249.
//
// Solidity: function setRadioactiveWaste(address rWasteAddress) returns()
func (_Kaiju *KaijuSession) SetRadioactiveWaste(rWasteAddress common.Address) (*types.Transaction, error) {
	return _Kaiju.Contract.SetRadioactiveWaste(&_Kaiju.TransactOpts, rWasteAddress)
}

// SetRadioactiveWaste is a paid mutator transaction binding the contract method 0x52d2d249.
//
// Solidity: function setRadioactiveWaste(address rWasteAddress) returns()
func (_Kaiju *KaijuTransactorSession) SetRadioactiveWaste(rWasteAddress common.Address) (*types.Transaction, error) {
	return _Kaiju.Contract.SetRadioactiveWaste(&_Kaiju.TransactOpts, rWasteAddress)
}

// TogglePresale is a paid mutator transaction binding the contract method 0x34393743.
//
// Solidity: function togglePresale() returns()
func (_Kaiju *KaijuTransactor) TogglePresale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "togglePresale")
}

// TogglePresale is a paid mutator transaction binding the contract method 0x34393743.
//
// Solidity: function togglePresale() returns()
func (_Kaiju *KaijuSession) TogglePresale() (*types.Transaction, error) {
	return _Kaiju.Contract.TogglePresale(&_Kaiju.TransactOpts)
}

// TogglePresale is a paid mutator transaction binding the contract method 0x34393743.
//
// Solidity: function togglePresale() returns()
func (_Kaiju *KaijuTransactorSession) TogglePresale() (*types.Transaction, error) {
	return _Kaiju.Contract.TogglePresale(&_Kaiju.TransactOpts)
}

// ToggleSale is a paid mutator transaction binding the contract method 0x7d8966e4.
//
// Solidity: function toggleSale() returns()
func (_Kaiju *KaijuTransactor) ToggleSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "toggleSale")
}

// ToggleSale is a paid mutator transaction binding the contract method 0x7d8966e4.
//
// Solidity: function toggleSale() returns()
func (_Kaiju *KaijuSession) ToggleSale() (*types.Transaction, error) {
	return _Kaiju.Contract.ToggleSale(&_Kaiju.TransactOpts)
}

// ToggleSale is a paid mutator transaction binding the contract method 0x7d8966e4.
//
// Solidity: function toggleSale() returns()
func (_Kaiju *KaijuTransactorSession) ToggleSale() (*types.Transaction, error) {
	return _Kaiju.Contract.ToggleSale(&_Kaiju.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Kaiju *KaijuTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Kaiju *KaijuSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.TransferFrom(&_Kaiju.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Kaiju *KaijuTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kaiju.Contract.TransferFrom(&_Kaiju.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kaiju *KaijuTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kaiju *KaijuSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kaiju.Contract.TransferOwnership(&_Kaiju.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kaiju *KaijuTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kaiju.Contract.TransferOwnership(&_Kaiju.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Kaiju *KaijuTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kaiju.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Kaiju *KaijuSession) Withdraw() (*types.Transaction, error) {
	return _Kaiju.Contract.Withdraw(&_Kaiju.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Kaiju *KaijuTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Kaiju.Contract.Withdraw(&_Kaiju.TransactOpts)
}

// KaijuApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Kaiju contract.
type KaijuApprovalIterator struct {
	Event *KaijuApproval // Event containing the contract specifics and raw log

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
func (it *KaijuApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaijuApproval)
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
		it.Event = new(KaijuApproval)
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
func (it *KaijuApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaijuApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaijuApproval represents a Approval event raised by the Kaiju contract.
type KaijuApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Kaiju *KaijuFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KaijuApprovalIterator, error) {

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

	logs, sub, err := _Kaiju.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KaijuApprovalIterator{contract: _Kaiju.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Kaiju *KaijuFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KaijuApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Kaiju.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaijuApproval)
				if err := _Kaiju.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Kaiju *KaijuFilterer) ParseApproval(log types.Log) (*KaijuApproval, error) {
	event := new(KaijuApproval)
	if err := _Kaiju.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaijuApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Kaiju contract.
type KaijuApprovalForAllIterator struct {
	Event *KaijuApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KaijuApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaijuApprovalForAll)
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
		it.Event = new(KaijuApprovalForAll)
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
func (it *KaijuApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaijuApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaijuApprovalForAll represents a ApprovalForAll event raised by the Kaiju contract.
type KaijuApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Kaiju *KaijuFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KaijuApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Kaiju.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KaijuApprovalForAllIterator{contract: _Kaiju.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Kaiju *KaijuFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KaijuApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Kaiju.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaijuApprovalForAll)
				if err := _Kaiju.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Kaiju *KaijuFilterer) ParseApprovalForAll(log types.Log) (*KaijuApprovalForAll, error) {
	event := new(KaijuApprovalForAll)
	if err := _Kaiju.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaijuBioChangedIterator is returned from FilterBioChanged and is used to iterate over the raw logs and unpacked data for BioChanged events raised by the Kaiju contract.
type KaijuBioChangedIterator struct {
	Event *KaijuBioChanged // Event containing the contract specifics and raw log

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
func (it *KaijuBioChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaijuBioChanged)
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
		it.Event = new(KaijuBioChanged)
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
func (it *KaijuBioChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaijuBioChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaijuBioChanged represents a BioChanged event raised by the Kaiju contract.
type KaijuBioChanged struct {
	KaijuId  *big.Int
	KaijuBio string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBioChanged is a free log retrieval operation binding the contract event 0xb5d3e30019a90e2b35059d238ddbac8259a3d65d9a4f47395d713901d17cc112.
//
// Solidity: event BioChanged(uint256 kaijuId, string kaijuBio)
func (_Kaiju *KaijuFilterer) FilterBioChanged(opts *bind.FilterOpts) (*KaijuBioChangedIterator, error) {

	logs, sub, err := _Kaiju.contract.FilterLogs(opts, "BioChanged")
	if err != nil {
		return nil, err
	}
	return &KaijuBioChangedIterator{contract: _Kaiju.contract, event: "BioChanged", logs: logs, sub: sub}, nil
}

// WatchBioChanged is a free log subscription operation binding the contract event 0xb5d3e30019a90e2b35059d238ddbac8259a3d65d9a4f47395d713901d17cc112.
//
// Solidity: event BioChanged(uint256 kaijuId, string kaijuBio)
func (_Kaiju *KaijuFilterer) WatchBioChanged(opts *bind.WatchOpts, sink chan<- *KaijuBioChanged) (event.Subscription, error) {

	logs, sub, err := _Kaiju.contract.WatchLogs(opts, "BioChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaijuBioChanged)
				if err := _Kaiju.contract.UnpackLog(event, "BioChanged", log); err != nil {
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

// ParseBioChanged is a log parse operation binding the contract event 0xb5d3e30019a90e2b35059d238ddbac8259a3d65d9a4f47395d713901d17cc112.
//
// Solidity: event BioChanged(uint256 kaijuId, string kaijuBio)
func (_Kaiju *KaijuFilterer) ParseBioChanged(log types.Log) (*KaijuBioChanged, error) {
	event := new(KaijuBioChanged)
	if err := _Kaiju.contract.UnpackLog(event, "BioChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaijuKaijuFusionIterator is returned from FilterKaijuFusion and is used to iterate over the raw logs and unpacked data for KaijuFusion events raised by the Kaiju contract.
type KaijuKaijuFusionIterator struct {
	Event *KaijuKaijuFusion // Event containing the contract specifics and raw log

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
func (it *KaijuKaijuFusionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaijuKaijuFusion)
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
		it.Event = new(KaijuKaijuFusion)
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
func (it *KaijuKaijuFusionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaijuKaijuFusionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaijuKaijuFusion represents a KaijuFusion event raised by the Kaiju contract.
type KaijuKaijuFusion struct {
	KaijuId *big.Int
	Parent1 *big.Int
	Parent2 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKaijuFusion is a free log retrieval operation binding the contract event 0x8afa7b07959f802d11fde473076edd2e6e364aaa294f494baf488b6fd9bf896a.
//
// Solidity: event KaijuFusion(uint256 kaijuId, uint256 parent1, uint256 parent2)
func (_Kaiju *KaijuFilterer) FilterKaijuFusion(opts *bind.FilterOpts) (*KaijuKaijuFusionIterator, error) {

	logs, sub, err := _Kaiju.contract.FilterLogs(opts, "KaijuFusion")
	if err != nil {
		return nil, err
	}
	return &KaijuKaijuFusionIterator{contract: _Kaiju.contract, event: "KaijuFusion", logs: logs, sub: sub}, nil
}

// WatchKaijuFusion is a free log subscription operation binding the contract event 0x8afa7b07959f802d11fde473076edd2e6e364aaa294f494baf488b6fd9bf896a.
//
// Solidity: event KaijuFusion(uint256 kaijuId, uint256 parent1, uint256 parent2)
func (_Kaiju *KaijuFilterer) WatchKaijuFusion(opts *bind.WatchOpts, sink chan<- *KaijuKaijuFusion) (event.Subscription, error) {

	logs, sub, err := _Kaiju.contract.WatchLogs(opts, "KaijuFusion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaijuKaijuFusion)
				if err := _Kaiju.contract.UnpackLog(event, "KaijuFusion", log); err != nil {
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

// ParseKaijuFusion is a log parse operation binding the contract event 0x8afa7b07959f802d11fde473076edd2e6e364aaa294f494baf488b6fd9bf896a.
//
// Solidity: event KaijuFusion(uint256 kaijuId, uint256 parent1, uint256 parent2)
func (_Kaiju *KaijuFilterer) ParseKaijuFusion(log types.Log) (*KaijuKaijuFusion, error) {
	event := new(KaijuKaijuFusion)
	if err := _Kaiju.contract.UnpackLog(event, "KaijuFusion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaijuKaijuRevealedIterator is returned from FilterKaijuRevealed and is used to iterate over the raw logs and unpacked data for KaijuRevealed events raised by the Kaiju contract.
type KaijuKaijuRevealedIterator struct {
	Event *KaijuKaijuRevealed // Event containing the contract specifics and raw log

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
func (it *KaijuKaijuRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaijuKaijuRevealed)
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
		it.Event = new(KaijuKaijuRevealed)
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
func (it *KaijuKaijuRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaijuKaijuRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaijuKaijuRevealed represents a KaijuRevealed event raised by the Kaiju contract.
type KaijuKaijuRevealed struct {
	KaijuId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKaijuRevealed is a free log retrieval operation binding the contract event 0xf549b67fb782b5944f9953a1ef470775d2de1be50196383eacb31ea650310c53.
//
// Solidity: event KaijuRevealed(uint256 kaijuId)
func (_Kaiju *KaijuFilterer) FilterKaijuRevealed(opts *bind.FilterOpts) (*KaijuKaijuRevealedIterator, error) {

	logs, sub, err := _Kaiju.contract.FilterLogs(opts, "KaijuRevealed")
	if err != nil {
		return nil, err
	}
	return &KaijuKaijuRevealedIterator{contract: _Kaiju.contract, event: "KaijuRevealed", logs: logs, sub: sub}, nil
}

// WatchKaijuRevealed is a free log subscription operation binding the contract event 0xf549b67fb782b5944f9953a1ef470775d2de1be50196383eacb31ea650310c53.
//
// Solidity: event KaijuRevealed(uint256 kaijuId)
func (_Kaiju *KaijuFilterer) WatchKaijuRevealed(opts *bind.WatchOpts, sink chan<- *KaijuKaijuRevealed) (event.Subscription, error) {

	logs, sub, err := _Kaiju.contract.WatchLogs(opts, "KaijuRevealed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaijuKaijuRevealed)
				if err := _Kaiju.contract.UnpackLog(event, "KaijuRevealed", log); err != nil {
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

// ParseKaijuRevealed is a log parse operation binding the contract event 0xf549b67fb782b5944f9953a1ef470775d2de1be50196383eacb31ea650310c53.
//
// Solidity: event KaijuRevealed(uint256 kaijuId)
func (_Kaiju *KaijuFilterer) ParseKaijuRevealed(log types.Log) (*KaijuKaijuRevealed, error) {
	event := new(KaijuKaijuRevealed)
	if err := _Kaiju.contract.UnpackLog(event, "KaijuRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaijuNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the Kaiju contract.
type KaijuNameChangedIterator struct {
	Event *KaijuNameChanged // Event containing the contract specifics and raw log

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
func (it *KaijuNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaijuNameChanged)
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
		it.Event = new(KaijuNameChanged)
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
func (it *KaijuNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaijuNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaijuNameChanged represents a NameChanged event raised by the Kaiju contract.
type KaijuNameChanged struct {
	KaijuId   *big.Int
	KaijuName string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0x8edfa912e70e283a8ef6d6f52cd1faef9690ff989eff2f11a134e8478ba7b28b.
//
// Solidity: event NameChanged(uint256 kaijuId, string kaijuName)
func (_Kaiju *KaijuFilterer) FilterNameChanged(opts *bind.FilterOpts) (*KaijuNameChangedIterator, error) {

	logs, sub, err := _Kaiju.contract.FilterLogs(opts, "NameChanged")
	if err != nil {
		return nil, err
	}
	return &KaijuNameChangedIterator{contract: _Kaiju.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0x8edfa912e70e283a8ef6d6f52cd1faef9690ff989eff2f11a134e8478ba7b28b.
//
// Solidity: event NameChanged(uint256 kaijuId, string kaijuName)
func (_Kaiju *KaijuFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *KaijuNameChanged) (event.Subscription, error) {

	logs, sub, err := _Kaiju.contract.WatchLogs(opts, "NameChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaijuNameChanged)
				if err := _Kaiju.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0x8edfa912e70e283a8ef6d6f52cd1faef9690ff989eff2f11a134e8478ba7b28b.
//
// Solidity: event NameChanged(uint256 kaijuId, string kaijuName)
func (_Kaiju *KaijuFilterer) ParseNameChanged(log types.Log) (*KaijuNameChanged, error) {
	event := new(KaijuNameChanged)
	if err := _Kaiju.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaijuOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Kaiju contract.
type KaijuOwnershipTransferredIterator struct {
	Event *KaijuOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KaijuOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaijuOwnershipTransferred)
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
		it.Event = new(KaijuOwnershipTransferred)
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
func (it *KaijuOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaijuOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaijuOwnershipTransferred represents a OwnershipTransferred event raised by the Kaiju contract.
type KaijuOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kaiju *KaijuFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KaijuOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kaiju.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KaijuOwnershipTransferredIterator{contract: _Kaiju.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kaiju *KaijuFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KaijuOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kaiju.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaijuOwnershipTransferred)
				if err := _Kaiju.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Kaiju *KaijuFilterer) ParseOwnershipTransferred(log types.Log) (*KaijuOwnershipTransferred, error) {
	event := new(KaijuOwnershipTransferred)
	if err := _Kaiju.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KaijuTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Kaiju contract.
type KaijuTransferIterator struct {
	Event *KaijuTransfer // Event containing the contract specifics and raw log

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
func (it *KaijuTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KaijuTransfer)
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
		it.Event = new(KaijuTransfer)
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
func (it *KaijuTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KaijuTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KaijuTransfer represents a Transfer event raised by the Kaiju contract.
type KaijuTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Kaiju *KaijuFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KaijuTransferIterator, error) {

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

	logs, sub, err := _Kaiju.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KaijuTransferIterator{contract: _Kaiju.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Kaiju *KaijuFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KaijuTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Kaiju.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KaijuTransfer)
				if err := _Kaiju.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Kaiju *KaijuFilterer) ParseTransfer(log types.Log) (*KaijuTransfer, error) {
	event := new(KaijuTransfer)
	if err := _Kaiju.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
