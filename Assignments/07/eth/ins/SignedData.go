// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SignedData

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

// SignedDataABI is the input ABI used to generate the binding from.
const SignedDataABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newVersion\",\"type\":\"uint256\"},{\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"_implementation\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_version\",\"type\":\"uint256\"},{\"name\":\"_implementation\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proxyToContractVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"proxyToAddress\",\"type\":\"address\"}],\"name\":\"ContractUpgradeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SignedData is an auto generated Go binding around an Ethereum contract.
type SignedData struct {
	SignedDataCaller     // Read-only binding to the contract
	SignedDataTransactor // Write-only binding to the contract
	SignedDataFilterer   // Log filterer for contract events
}

// SignedDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignedDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignedDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignedDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignedDataSession struct {
	Contract     *SignedData       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignedDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignedDataCallerSession struct {
	Contract *SignedDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SignedDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignedDataTransactorSession struct {
	Contract     *SignedDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SignedDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignedDataRaw struct {
	Contract *SignedData // Generic contract binding to access the raw methods on
}

// SignedDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignedDataCallerRaw struct {
	Contract *SignedDataCaller // Generic read-only contract binding to access the raw methods on
}

// SignedDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignedDataTransactorRaw struct {
	Contract *SignedDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignedData creates a new instance of SignedData, bound to a specific deployed contract.
func NewSignedData(address common.Address, backend bind.ContractBackend) (*SignedData, error) {
	contract, err := bindSignedData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignedData{SignedDataCaller: SignedDataCaller{contract: contract}, SignedDataTransactor: SignedDataTransactor{contract: contract}, SignedDataFilterer: SignedDataFilterer{contract: contract}}, nil
}

// NewSignedDataCaller creates a new read-only instance of SignedData, bound to a specific deployed contract.
func NewSignedDataCaller(address common.Address, caller bind.ContractCaller) (*SignedDataCaller, error) {
	contract, err := bindSignedData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignedDataCaller{contract: contract}, nil
}

// NewSignedDataTransactor creates a new write-only instance of SignedData, bound to a specific deployed contract.
func NewSignedDataTransactor(address common.Address, transactor bind.ContractTransactor) (*SignedDataTransactor, error) {
	contract, err := bindSignedData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignedDataTransactor{contract: contract}, nil
}

// NewSignedDataFilterer creates a new log filterer instance of SignedData, bound to a specific deployed contract.
func NewSignedDataFilterer(address common.Address, filterer bind.ContractFilterer) (*SignedDataFilterer, error) {
	contract, err := bindSignedData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignedDataFilterer{contract: contract}, nil
}

// bindSignedData binds a generic wrapper to an already deployed contract.
func bindSignedData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignedDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedData *SignedDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignedData.Contract.SignedDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedData *SignedDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedData.Contract.SignedDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedData *SignedDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedData.Contract.SignedDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedData *SignedDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignedData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedData *SignedDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedData *SignedDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedData.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentContractVersion is a free data retrieval call binding the contract method 0x43ed6d77.
//
// Solidity: function getCurrentContractVersion() constant returns(uint256)
func (_SignedData *SignedDataCaller) GetCurrentContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SignedData.contract.Call(opts, out, "getCurrentContractVersion")
	return *ret0, err
}

// GetCurrentContractVersion is a free data retrieval call binding the contract method 0x43ed6d77.
//
// Solidity: function getCurrentContractVersion() constant returns(uint256)
func (_SignedData *SignedDataSession) GetCurrentContractVersion() (*big.Int, error) {
	return _SignedData.Contract.GetCurrentContractVersion(&_SignedData.CallOpts)
}

// GetCurrentContractVersion is a free data retrieval call binding the contract method 0x43ed6d77.
//
// Solidity: function getCurrentContractVersion() constant returns(uint256)
func (_SignedData *SignedDataCallerSession) GetCurrentContractVersion() (*big.Int, error) {
	return _SignedData.Contract.GetCurrentContractVersion(&_SignedData.CallOpts)
}

// GetToAddress is a free data retrieval call binding the contract method 0xe411842f.
//
// Solidity: function getToAddress() constant returns(address)
func (_SignedData *SignedDataCaller) GetToAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SignedData.contract.Call(opts, out, "getToAddress")
	return *ret0, err
}

// GetToAddress is a free data retrieval call binding the contract method 0xe411842f.
//
// Solidity: function getToAddress() constant returns(address)
func (_SignedData *SignedDataSession) GetToAddress() (common.Address, error) {
	return _SignedData.Contract.GetToAddress(&_SignedData.CallOpts)
}

// GetToAddress is a free data retrieval call binding the contract method 0xe411842f.
//
// Solidity: function getToAddress() constant returns(address)
func (_SignedData *SignedDataCallerSession) GetToAddress() (common.Address, error) {
	return _SignedData.Contract.GetToAddress(&_SignedData.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() constant returns(_implementation address)
func (_SignedData *SignedDataCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SignedData.contract.Call(opts, out, "implementation")
	return *ret0, err
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() constant returns(_implementation address)
func (_SignedData *SignedDataSession) Implementation() (common.Address, error) {
	return _SignedData.Contract.Implementation(&_SignedData.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() constant returns(_implementation address)
func (_SignedData *SignedDataCallerSession) Implementation() (common.Address, error) {
	return _SignedData.Contract.Implementation(&_SignedData.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SignedData *SignedDataCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SignedData.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SignedData *SignedDataSession) IsOwner() (bool, error) {
	return _SignedData.Contract.IsOwner(&_SignedData.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SignedData *SignedDataCallerSession) IsOwner() (bool, error) {
	return _SignedData.Contract.IsOwner(&_SignedData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SignedData *SignedDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SignedData.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SignedData *SignedDataSession) Owner() (common.Address, error) {
	return _SignedData.Contract.Owner(&_SignedData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SignedData *SignedDataCallerSession) Owner() (common.Address, error) {
	return _SignedData.Contract.Owner(&_SignedData.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SignedData *SignedDataTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedData.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SignedData *SignedDataSession) RenounceOwnership() (*types.Transaction, error) {
	return _SignedData.Contract.RenounceOwnership(&_SignedData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SignedData *SignedDataTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SignedData.Contract.RenounceOwnership(&_SignedData.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SignedData *SignedDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SignedData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SignedData *SignedDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SignedData.Contract.TransferOwnership(&_SignedData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SignedData *SignedDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SignedData.Contract.TransferOwnership(&_SignedData.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3ad06d16.
//
// Solidity: function upgradeTo(_newVersion uint256, _implementation address) returns()
func (_SignedData *SignedDataTransactor) UpgradeTo(opts *bind.TransactOpts, _newVersion *big.Int, _implementation common.Address) (*types.Transaction, error) {
	return _SignedData.contract.Transact(opts, "upgradeTo", _newVersion, _implementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3ad06d16.
//
// Solidity: function upgradeTo(_newVersion uint256, _implementation address) returns()
func (_SignedData *SignedDataSession) UpgradeTo(_newVersion *big.Int, _implementation common.Address) (*types.Transaction, error) {
	return _SignedData.Contract.UpgradeTo(&_SignedData.TransactOpts, _newVersion, _implementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3ad06d16.
//
// Solidity: function upgradeTo(_newVersion uint256, _implementation address) returns()
func (_SignedData *SignedDataTransactorSession) UpgradeTo(_newVersion *big.Int, _implementation common.Address) (*types.Transaction, error) {
	return _SignedData.Contract.UpgradeTo(&_SignedData.TransactOpts, _newVersion, _implementation)
}

// SignedDataContractUpgradeEventIterator is returned from FilterContractUpgradeEvent and is used to iterate over the raw logs and unpacked data for ContractUpgradeEvent events raised by the SignedData contract.
type SignedDataContractUpgradeEventIterator struct {
	Event *SignedDataContractUpgradeEvent // Event containing the contract specifics and raw log

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
func (it *SignedDataContractUpgradeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignedDataContractUpgradeEvent)
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
		it.Event = new(SignedDataContractUpgradeEvent)
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
func (it *SignedDataContractUpgradeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignedDataContractUpgradeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignedDataContractUpgradeEvent represents a ContractUpgradeEvent event raised by the SignedData contract.
type SignedDataContractUpgradeEvent struct {
	ProxyToContractVersion *big.Int
	ProxyToAddress         common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterContractUpgradeEvent is a free log retrieval operation binding the contract event 0x88ee542495575e7963fd4f7e164feb7d209e16d8f5f9900cc094cacc84123e2e.
//
// Solidity: e ContractUpgradeEvent(proxyToContractVersion uint256, proxyToAddress address)
func (_SignedData *SignedDataFilterer) FilterContractUpgradeEvent(opts *bind.FilterOpts) (*SignedDataContractUpgradeEventIterator, error) {

	logs, sub, err := _SignedData.contract.FilterLogs(opts, "ContractUpgradeEvent")
	if err != nil {
		return nil, err
	}
	return &SignedDataContractUpgradeEventIterator{contract: _SignedData.contract, event: "ContractUpgradeEvent", logs: logs, sub: sub}, nil
}

// WatchContractUpgradeEvent is a free log subscription operation binding the contract event 0x88ee542495575e7963fd4f7e164feb7d209e16d8f5f9900cc094cacc84123e2e.
//
// Solidity: e ContractUpgradeEvent(proxyToContractVersion uint256, proxyToAddress address)
func (_SignedData *SignedDataFilterer) WatchContractUpgradeEvent(opts *bind.WatchOpts, sink chan<- *SignedDataContractUpgradeEvent) (event.Subscription, error) {

	logs, sub, err := _SignedData.contract.WatchLogs(opts, "ContractUpgradeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignedDataContractUpgradeEvent)
				if err := _SignedData.contract.UnpackLog(event, "ContractUpgradeEvent", log); err != nil {
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

// SignedDataOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SignedData contract.
type SignedDataOwnershipTransferredIterator struct {
	Event *SignedDataOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SignedDataOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignedDataOwnershipTransferred)
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
		it.Event = new(SignedDataOwnershipTransferred)
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
func (it *SignedDataOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignedDataOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignedDataOwnershipTransferred represents a OwnershipTransferred event raised by the SignedData contract.
type SignedDataOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SignedData *SignedDataFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SignedDataOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SignedData.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SignedDataOwnershipTransferredIterator{contract: _SignedData.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SignedData *SignedDataFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SignedDataOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SignedData.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignedDataOwnershipTransferred)
				if err := _SignedData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
