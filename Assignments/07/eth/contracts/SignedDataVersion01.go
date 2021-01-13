// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SignedDataVersion01

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// SignedDataVersion01ABI is the input ABI used to generate the binding from.
const SignedDataVersion01ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_app\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes2\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalanceContract\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minPayment\",\"type\":\"uint256\"}],\"name\":\"setMinPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinPayment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_app\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"uint256\"},{\"name\":\"_data_r\",\"type\":\"bytes32\"},{\"name\":\"_data_s\",\"type\":\"bytes32\"},{\"name\":\"_data_v\",\"type\":\"bytes2\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"application\",\"type\":\"uint256\"},{\"name\":\"payFor\",\"type\":\"uint256\"}],\"name\":\"genReceivedFunds\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"App\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"Name\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ValueR\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ValueS\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ValueV\",\"type\":\"bytes2\"},{\"indexed\":false,\"name\":\"By\",\"type\":\"address\"}],\"name\":\"DataChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"application\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"payFor\",\"type\":\"uint256\"}],\"name\":\"ReceivedFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SignedDataVersion01 is an auto generated Go binding around an Ethereum contract.
type SignedDataVersion01 struct {
	SignedDataVersion01Caller     // Read-only binding to the contract
	SignedDataVersion01Transactor // Write-only binding to the contract
	SignedDataVersion01Filterer   // Log filterer for contract events
}

// SignedDataVersion01Caller is an auto generated read-only Go binding around an Ethereum contract.
type SignedDataVersion01Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedDataVersion01Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SignedDataVersion01Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedDataVersion01Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignedDataVersion01Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedDataVersion01Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignedDataVersion01Session struct {
	Contract     *SignedDataVersion01 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SignedDataVersion01CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignedDataVersion01CallerSession struct {
	Contract *SignedDataVersion01Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SignedDataVersion01TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignedDataVersion01TransactorSession struct {
	Contract     *SignedDataVersion01Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SignedDataVersion01Raw is an auto generated low-level Go binding around an Ethereum contract.
type SignedDataVersion01Raw struct {
	Contract *SignedDataVersion01 // Generic contract binding to access the raw methods on
}

// SignedDataVersion01CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignedDataVersion01CallerRaw struct {
	Contract *SignedDataVersion01Caller // Generic read-only contract binding to access the raw methods on
}

// SignedDataVersion01TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignedDataVersion01TransactorRaw struct {
	Contract *SignedDataVersion01Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSignedDataVersion01 creates a new instance of SignedDataVersion01, bound to a specific deployed contract.
func NewSignedDataVersion01(address common.Address, backend bind.ContractBackend) (*SignedDataVersion01, error) {
	contract, err := bindSignedDataVersion01(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignedDataVersion01{SignedDataVersion01Caller: SignedDataVersion01Caller{contract: contract}, SignedDataVersion01Transactor: SignedDataVersion01Transactor{contract: contract}, SignedDataVersion01Filterer: SignedDataVersion01Filterer{contract: contract}}, nil
}

// NewSignedDataVersion01Caller creates a new read-only instance of SignedDataVersion01, bound to a specific deployed contract.
func NewSignedDataVersion01Caller(address common.Address, caller bind.ContractCaller) (*SignedDataVersion01Caller, error) {
	contract, err := bindSignedDataVersion01(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignedDataVersion01Caller{contract: contract}, nil
}

// NewSignedDataVersion01Transactor creates a new write-only instance of SignedDataVersion01, bound to a specific deployed contract.
func NewSignedDataVersion01Transactor(address common.Address, transactor bind.ContractTransactor) (*SignedDataVersion01Transactor, error) {
	contract, err := bindSignedDataVersion01(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignedDataVersion01Transactor{contract: contract}, nil
}

// NewSignedDataVersion01Filterer creates a new log filterer instance of SignedDataVersion01, bound to a specific deployed contract.
func NewSignedDataVersion01Filterer(address common.Address, filterer bind.ContractFilterer) (*SignedDataVersion01Filterer, error) {
	contract, err := bindSignedDataVersion01(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignedDataVersion01Filterer{contract: contract}, nil
}

// bindSignedDataVersion01 binds a generic wrapper to an already deployed contract.
func bindSignedDataVersion01(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignedDataVersion01ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedDataVersion01 *SignedDataVersion01Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignedDataVersion01.Contract.SignedDataVersion01Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedDataVersion01 *SignedDataVersion01Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.SignedDataVersion01Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedDataVersion01 *SignedDataVersion01Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.SignedDataVersion01Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedDataVersion01 *SignedDataVersion01CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignedDataVersion01.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedDataVersion01 *SignedDataVersion01TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedDataVersion01 *SignedDataVersion01TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.contract.Transact(opts, method, params...)
}

// GetBalanceContract is a free data retrieval call binding the contract method 0x66bd78fd.
//
// Solidity: function getBalanceContract() constant returns(uint256)
func (_SignedDataVersion01 *SignedDataVersion01Caller) GetBalanceContract(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SignedDataVersion01.contract.Call(opts, out, "getBalanceContract")
	return *ret0, err
}

// GetBalanceContract is a free data retrieval call binding the contract method 0x66bd78fd.
//
// Solidity: function getBalanceContract() constant returns(uint256)
func (_SignedDataVersion01 *SignedDataVersion01Session) GetBalanceContract() (*big.Int, error) {
	return _SignedDataVersion01.Contract.GetBalanceContract(&_SignedDataVersion01.CallOpts)
}

// GetBalanceContract is a free data retrieval call binding the contract method 0x66bd78fd.
//
// Solidity: function getBalanceContract() constant returns(uint256)
func (_SignedDataVersion01 *SignedDataVersion01CallerSession) GetBalanceContract() (*big.Int, error) {
	return _SignedDataVersion01.Contract.GetBalanceContract(&_SignedDataVersion01.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x4e4fe306.
//
// Solidity: function getData(_app uint256, _name uint256) constant returns(bytes32, bytes32, bytes2)
func (_SignedDataVersion01 *SignedDataVersion01Caller) GetData(opts *bind.CallOpts, _app *big.Int, _name *big.Int) ([32]byte, [32]byte, [2]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
		ret2 = new([2]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SignedDataVersion01.contract.Call(opts, out, "getData", _app, _name)
	return *ret0, *ret1, *ret2, err
}

// GetData is a free data retrieval call binding the contract method 0x4e4fe306.
//
// Solidity: function getData(_app uint256, _name uint256) constant returns(bytes32, bytes32, bytes2)
func (_SignedDataVersion01 *SignedDataVersion01Session) GetData(_app *big.Int, _name *big.Int) ([32]byte, [32]byte, [2]byte, error) {
	return _SignedDataVersion01.Contract.GetData(&_SignedDataVersion01.CallOpts, _app, _name)
}

// GetData is a free data retrieval call binding the contract method 0x4e4fe306.
//
// Solidity: function getData(_app uint256, _name uint256) constant returns(bytes32, bytes32, bytes2)
func (_SignedDataVersion01 *SignedDataVersion01CallerSession) GetData(_app *big.Int, _name *big.Int) ([32]byte, [32]byte, [2]byte, error) {
	return _SignedDataVersion01.Contract.GetData(&_SignedDataVersion01.CallOpts, _app, _name)
}

// GetMinPayment is a free data retrieval call binding the contract method 0x77a890cd.
//
// Solidity: function getMinPayment() constant returns(uint256)
func (_SignedDataVersion01 *SignedDataVersion01Caller) GetMinPayment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SignedDataVersion01.contract.Call(opts, out, "getMinPayment")
	return *ret0, err
}

// GetMinPayment is a free data retrieval call binding the contract method 0x77a890cd.
//
// Solidity: function getMinPayment() constant returns(uint256)
func (_SignedDataVersion01 *SignedDataVersion01Session) GetMinPayment() (*big.Int, error) {
	return _SignedDataVersion01.Contract.GetMinPayment(&_SignedDataVersion01.CallOpts)
}

// GetMinPayment is a free data retrieval call binding the contract method 0x77a890cd.
//
// Solidity: function getMinPayment() constant returns(uint256)
func (_SignedDataVersion01 *SignedDataVersion01CallerSession) GetMinPayment() (*big.Int, error) {
	return _SignedDataVersion01.Contract.GetMinPayment(&_SignedDataVersion01.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SignedDataVersion01 *SignedDataVersion01Caller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SignedDataVersion01.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SignedDataVersion01 *SignedDataVersion01Session) IsOwner() (bool, error) {
	return _SignedDataVersion01.Contract.IsOwner(&_SignedDataVersion01.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SignedDataVersion01 *SignedDataVersion01CallerSession) IsOwner() (bool, error) {
	return _SignedDataVersion01.Contract.IsOwner(&_SignedDataVersion01.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SignedDataVersion01 *SignedDataVersion01Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SignedDataVersion01.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SignedDataVersion01 *SignedDataVersion01Session) Owner() (common.Address, error) {
	return _SignedDataVersion01.Contract.Owner(&_SignedDataVersion01.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SignedDataVersion01 *SignedDataVersion01CallerSession) Owner() (common.Address, error) {
	return _SignedDataVersion01.Contract.Owner(&_SignedDataVersion01.CallOpts)
}

// GenReceivedFunds is a paid mutator transaction binding the contract method 0xbe2b1636.
//
// Solidity: function genReceivedFunds(application uint256, payFor uint256) returns()
func (_SignedDataVersion01 *SignedDataVersion01Transactor) GenReceivedFunds(opts *bind.TransactOpts, application *big.Int, payFor *big.Int) (*types.Transaction, error) {
	return _SignedDataVersion01.contract.Transact(opts, "genReceivedFunds", application, payFor)
}

// GenReceivedFunds is a paid mutator transaction binding the contract method 0xbe2b1636.
//
// Solidity: function genReceivedFunds(application uint256, payFor uint256) returns()
func (_SignedDataVersion01 *SignedDataVersion01Session) GenReceivedFunds(application *big.Int, payFor *big.Int) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.GenReceivedFunds(&_SignedDataVersion01.TransactOpts, application, payFor)
}

// GenReceivedFunds is a paid mutator transaction binding the contract method 0xbe2b1636.
//
// Solidity: function genReceivedFunds(application uint256, payFor uint256) returns()
func (_SignedDataVersion01 *SignedDataVersion01TransactorSession) GenReceivedFunds(application *big.Int, payFor *big.Int) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.GenReceivedFunds(&_SignedDataVersion01.TransactOpts, application, payFor)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_SignedDataVersion01 *SignedDataVersion01Transactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedDataVersion01.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_SignedDataVersion01 *SignedDataVersion01Session) Init() (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.Init(&_SignedDataVersion01.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_SignedDataVersion01 *SignedDataVersion01TransactorSession) Init() (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.Init(&_SignedDataVersion01.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_SignedDataVersion01 *SignedDataVersion01Transactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedDataVersion01.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_SignedDataVersion01 *SignedDataVersion01Session) Kill() (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.Kill(&_SignedDataVersion01.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_SignedDataVersion01 *SignedDataVersion01TransactorSession) Kill() (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.Kill(&_SignedDataVersion01.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SignedDataVersion01 *SignedDataVersion01Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedDataVersion01.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SignedDataVersion01 *SignedDataVersion01Session) RenounceOwnership() (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.RenounceOwnership(&_SignedDataVersion01.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SignedDataVersion01 *SignedDataVersion01TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.RenounceOwnership(&_SignedDataVersion01.TransactOpts)
}

// SetData is a paid mutator transaction binding the contract method 0xa5b7026a.
//
// Solidity: function setData(_app uint256, _name uint256, _data_r bytes32, _data_s bytes32, _data_v bytes2) returns()
func (_SignedDataVersion01 *SignedDataVersion01Transactor) SetData(opts *bind.TransactOpts, _app *big.Int, _name *big.Int, _data_r [32]byte, _data_s [32]byte, _data_v [2]byte) (*types.Transaction, error) {
	return _SignedDataVersion01.contract.Transact(opts, "setData", _app, _name, _data_r, _data_s, _data_v)
}

// SetData is a paid mutator transaction binding the contract method 0xa5b7026a.
//
// Solidity: function setData(_app uint256, _name uint256, _data_r bytes32, _data_s bytes32, _data_v bytes2) returns()
func (_SignedDataVersion01 *SignedDataVersion01Session) SetData(_app *big.Int, _name *big.Int, _data_r [32]byte, _data_s [32]byte, _data_v [2]byte) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.SetData(&_SignedDataVersion01.TransactOpts, _app, _name, _data_r, _data_s, _data_v)
}

// SetData is a paid mutator transaction binding the contract method 0xa5b7026a.
//
// Solidity: function setData(_app uint256, _name uint256, _data_r bytes32, _data_s bytes32, _data_v bytes2) returns()
func (_SignedDataVersion01 *SignedDataVersion01TransactorSession) SetData(_app *big.Int, _name *big.Int, _data_r [32]byte, _data_s [32]byte, _data_v [2]byte) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.SetData(&_SignedDataVersion01.TransactOpts, _app, _name, _data_r, _data_s, _data_v)
}

// SetMinPayment is a paid mutator transaction binding the contract method 0x6d427fa3.
//
// Solidity: function setMinPayment(_minPayment uint256) returns()
func (_SignedDataVersion01 *SignedDataVersion01Transactor) SetMinPayment(opts *bind.TransactOpts, _minPayment *big.Int) (*types.Transaction, error) {
	return _SignedDataVersion01.contract.Transact(opts, "setMinPayment", _minPayment)
}

// SetMinPayment is a paid mutator transaction binding the contract method 0x6d427fa3.
//
// Solidity: function setMinPayment(_minPayment uint256) returns()
func (_SignedDataVersion01 *SignedDataVersion01Session) SetMinPayment(_minPayment *big.Int) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.SetMinPayment(&_SignedDataVersion01.TransactOpts, _minPayment)
}

// SetMinPayment is a paid mutator transaction binding the contract method 0x6d427fa3.
//
// Solidity: function setMinPayment(_minPayment uint256) returns()
func (_SignedDataVersion01 *SignedDataVersion01TransactorSession) SetMinPayment(_minPayment *big.Int) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.SetMinPayment(&_SignedDataVersion01.TransactOpts, _minPayment)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SignedDataVersion01 *SignedDataVersion01Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SignedDataVersion01.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SignedDataVersion01 *SignedDataVersion01Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.TransferOwnership(&_SignedDataVersion01.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SignedDataVersion01 *SignedDataVersion01TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.TransferOwnership(&_SignedDataVersion01.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns(bool)
func (_SignedDataVersion01 *SignedDataVersion01Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SignedDataVersion01.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns(bool)
func (_SignedDataVersion01 *SignedDataVersion01Session) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.Withdraw(&_SignedDataVersion01.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns(bool)
func (_SignedDataVersion01 *SignedDataVersion01TransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _SignedDataVersion01.Contract.Withdraw(&_SignedDataVersion01.TransactOpts, amount)
}

// SignedDataVersion01DataChangeIterator is returned from FilterDataChange and is used to iterate over the raw logs and unpacked data for DataChange events raised by the SignedDataVersion01 contract.
type SignedDataVersion01DataChangeIterator struct {
	Event *SignedDataVersion01DataChange // Event containing the contract specifics and raw log

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
func (it *SignedDataVersion01DataChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignedDataVersion01DataChange)
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
		it.Event = new(SignedDataVersion01DataChange)
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
func (it *SignedDataVersion01DataChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignedDataVersion01DataChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignedDataVersion01DataChange represents a DataChange event raised by the SignedDataVersion01 contract.
type SignedDataVersion01DataChange struct {
	App    *big.Int
	Name   *big.Int
	ValueR [32]byte
	ValueS [32]byte
	ValueV [2]byte
	By     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDataChange is a free log retrieval operation binding the contract event 0x6169a3fa8eef34d8cbe067f1d83da287bada999a0779d71d6dc71cf7b93096c9.
//
// Solidity: e DataChange(App uint256, Name uint256, ValueR bytes32, ValueS bytes32, ValueV bytes2, By address)
func (_SignedDataVersion01 *SignedDataVersion01Filterer) FilterDataChange(opts *bind.FilterOpts) (*SignedDataVersion01DataChangeIterator, error) {

	logs, sub, err := _SignedDataVersion01.contract.FilterLogs(opts, "DataChange")
	if err != nil {
		return nil, err
	}
	return &SignedDataVersion01DataChangeIterator{contract: _SignedDataVersion01.contract, event: "DataChange", logs: logs, sub: sub}, nil
}

// WatchDataChange is a free log subscription operation binding the contract event 0x6169a3fa8eef34d8cbe067f1d83da287bada999a0779d71d6dc71cf7b93096c9.
//
// Solidity: e DataChange(App uint256, Name uint256, ValueR bytes32, ValueS bytes32, ValueV bytes2, By address)
func (_SignedDataVersion01 *SignedDataVersion01Filterer) WatchDataChange(opts *bind.WatchOpts, sink chan<- *SignedDataVersion01DataChange) (event.Subscription, error) {

	logs, sub, err := _SignedDataVersion01.contract.WatchLogs(opts, "DataChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignedDataVersion01DataChange)
				if err := _SignedDataVersion01.contract.UnpackLog(event, "DataChange", log); err != nil {
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

// SignedDataVersion01OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SignedDataVersion01 contract.
type SignedDataVersion01OwnershipTransferredIterator struct {
	Event *SignedDataVersion01OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SignedDataVersion01OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignedDataVersion01OwnershipTransferred)
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
		it.Event = new(SignedDataVersion01OwnershipTransferred)
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
func (it *SignedDataVersion01OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignedDataVersion01OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignedDataVersion01OwnershipTransferred represents a OwnershipTransferred event raised by the SignedDataVersion01 contract.
type SignedDataVersion01OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SignedDataVersion01 *SignedDataVersion01Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SignedDataVersion01OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SignedDataVersion01.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SignedDataVersion01OwnershipTransferredIterator{contract: _SignedDataVersion01.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SignedDataVersion01 *SignedDataVersion01Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SignedDataVersion01OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SignedDataVersion01.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignedDataVersion01OwnershipTransferred)
				if err := _SignedDataVersion01.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SignedDataVersion01ReceivedFundsIterator is returned from FilterReceivedFunds and is used to iterate over the raw logs and unpacked data for ReceivedFunds events raised by the SignedDataVersion01 contract.
type SignedDataVersion01ReceivedFundsIterator struct {
	Event *SignedDataVersion01ReceivedFunds // Event containing the contract specifics and raw log

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
func (it *SignedDataVersion01ReceivedFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignedDataVersion01ReceivedFunds)
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
		it.Event = new(SignedDataVersion01ReceivedFunds)
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
func (it *SignedDataVersion01ReceivedFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignedDataVersion01ReceivedFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignedDataVersion01ReceivedFunds represents a ReceivedFunds event raised by the SignedDataVersion01 contract.
type SignedDataVersion01ReceivedFunds struct {
	Sender      common.Address
	Value       *big.Int
	Application *big.Int
	PayFor      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedFunds is a free log retrieval operation binding the contract event 0x5f74095c3cf9cd2b4407992767113d77fc65cb2423a83ecc71977f3ace6b34f4.
//
// Solidity: e ReceivedFunds(sender address, value uint256, application uint256, payFor uint256)
func (_SignedDataVersion01 *SignedDataVersion01Filterer) FilterReceivedFunds(opts *bind.FilterOpts) (*SignedDataVersion01ReceivedFundsIterator, error) {

	logs, sub, err := _SignedDataVersion01.contract.FilterLogs(opts, "ReceivedFunds")
	if err != nil {
		return nil, err
	}
	return &SignedDataVersion01ReceivedFundsIterator{contract: _SignedDataVersion01.contract, event: "ReceivedFunds", logs: logs, sub: sub}, nil
}

// WatchReceivedFunds is a free log subscription operation binding the contract event 0x5f74095c3cf9cd2b4407992767113d77fc65cb2423a83ecc71977f3ace6b34f4.
//
// Solidity: e ReceivedFunds(sender address, value uint256, application uint256, payFor uint256)
func (_SignedDataVersion01 *SignedDataVersion01Filterer) WatchReceivedFunds(opts *bind.WatchOpts, sink chan<- *SignedDataVersion01ReceivedFunds) (event.Subscription, error) {

	logs, sub, err := _SignedDataVersion01.contract.WatchLogs(opts, "ReceivedFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignedDataVersion01ReceivedFunds)
				if err := _SignedDataVersion01.contract.UnpackLog(event, "ReceivedFunds", log); err != nil {
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

// SignedDataVersion01WithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the SignedDataVersion01 contract.
type SignedDataVersion01WithdrawnIterator struct {
	Event *SignedDataVersion01Withdrawn // Event containing the contract specifics and raw log

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
func (it *SignedDataVersion01WithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignedDataVersion01Withdrawn)
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
		it.Event = new(SignedDataVersion01Withdrawn)
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
func (it *SignedDataVersion01WithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignedDataVersion01WithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignedDataVersion01Withdrawn represents a Withdrawn event raised by the SignedDataVersion01 contract.
type SignedDataVersion01Withdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(to address, amount uint256)
func (_SignedDataVersion01 *SignedDataVersion01Filterer) FilterWithdrawn(opts *bind.FilterOpts) (*SignedDataVersion01WithdrawnIterator, error) {

	logs, sub, err := _SignedDataVersion01.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &SignedDataVersion01WithdrawnIterator{contract: _SignedDataVersion01.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(to address, amount uint256)
func (_SignedDataVersion01 *SignedDataVersion01Filterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *SignedDataVersion01Withdrawn) (event.Subscription, error) {

	logs, sub, err := _SignedDataVersion01.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignedDataVersion01Withdrawn)
				if err := _SignedDataVersion01.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
