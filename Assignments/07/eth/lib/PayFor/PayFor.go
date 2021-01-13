// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PayFor

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

// PayForABI is the input ABI used to generate the binding from.
const PayForABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"forProduct\",\"type\":\"uint256\"}],\"name\":\"ReceiveFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getNPayments\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalanceContract\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"getPaymentInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"application\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loc\",\"type\":\"uint256\"}],\"name\":\"ReceivedFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PayFor is an auto generated Go binding around an Ethereum contract.
type PayFor struct {
	PayForCaller     // Read-only binding to the contract
	PayForTransactor // Write-only binding to the contract
	PayForFilterer   // Log filterer for contract events
}

// PayForCaller is an auto generated read-only Go binding around an Ethereum contract.
type PayForCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayForTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PayForTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayForFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PayForFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayForSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PayForSession struct {
	Contract     *PayFor           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayForCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PayForCallerSession struct {
	Contract *PayForCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PayForTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PayForTransactorSession struct {
	Contract     *PayForTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayForRaw is an auto generated low-level Go binding around an Ethereum contract.
type PayForRaw struct {
	Contract *PayFor // Generic contract binding to access the raw methods on
}

// PayForCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PayForCallerRaw struct {
	Contract *PayForCaller // Generic read-only contract binding to access the raw methods on
}

// PayForTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PayForTransactorRaw struct {
	Contract *PayForTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayFor creates a new instance of PayFor, bound to a specific deployed contract.
func NewPayFor(address common.Address, backend bind.ContractBackend) (*PayFor, error) {
	contract, err := bindPayFor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PayFor{PayForCaller: PayForCaller{contract: contract}, PayForTransactor: PayForTransactor{contract: contract}, PayForFilterer: PayForFilterer{contract: contract}}, nil
}

// NewPayForCaller creates a new read-only instance of PayFor, bound to a specific deployed contract.
func NewPayForCaller(address common.Address, caller bind.ContractCaller) (*PayForCaller, error) {
	contract, err := bindPayFor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PayForCaller{contract: contract}, nil
}

// NewPayForTransactor creates a new write-only instance of PayFor, bound to a specific deployed contract.
func NewPayForTransactor(address common.Address, transactor bind.ContractTransactor) (*PayForTransactor, error) {
	contract, err := bindPayFor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PayForTransactor{contract: contract}, nil
}

// NewPayForFilterer creates a new log filterer instance of PayFor, bound to a specific deployed contract.
func NewPayForFilterer(address common.Address, filterer bind.ContractFilterer) (*PayForFilterer, error) {
	contract, err := bindPayFor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PayForFilterer{contract: contract}, nil
}

// bindPayFor binds a generic wrapper to an already deployed contract.
func bindPayFor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PayForABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayFor *PayForRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PayFor.Contract.PayForCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayFor *PayForRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayFor.Contract.PayForTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayFor *PayForRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayFor.Contract.PayForTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayFor *PayForCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PayFor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayFor *PayForTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayFor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayFor *PayForTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayFor.Contract.contract.Transact(opts, method, params...)
}

// GetBalanceContract is a free data retrieval call binding the contract method 0x66bd78fd.
//
// Solidity: function getBalanceContract() constant returns(uint256)
func (_PayFor *PayForCaller) GetBalanceContract(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PayFor.contract.Call(opts, out, "getBalanceContract")
	return *ret0, err
}

// GetBalanceContract is a free data retrieval call binding the contract method 0x66bd78fd.
//
// Solidity: function getBalanceContract() constant returns(uint256)
func (_PayFor *PayForSession) GetBalanceContract() (*big.Int, error) {
	return _PayFor.Contract.GetBalanceContract(&_PayFor.CallOpts)
}

// GetBalanceContract is a free data retrieval call binding the contract method 0x66bd78fd.
//
// Solidity: function getBalanceContract() constant returns(uint256)
func (_PayFor *PayForCallerSession) GetBalanceContract() (*big.Int, error) {
	return _PayFor.Contract.GetBalanceContract(&_PayFor.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_PayFor *PayForCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PayFor.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_PayFor *PayForSession) IsOwner() (bool, error) {
	return _PayFor.Contract.IsOwner(&_PayFor.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_PayFor *PayForCallerSession) IsOwner() (bool, error) {
	return _PayFor.Contract.IsOwner(&_PayFor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PayFor *PayForCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PayFor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PayFor *PayForSession) Owner() (common.Address, error) {
	return _PayFor.Contract.Owner(&_PayFor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PayFor *PayForCallerSession) Owner() (common.Address, error) {
	return _PayFor.Contract.Owner(&_PayFor.CallOpts)
}

// ReceiveFunds is a paid mutator transaction binding the contract method 0x46d37eea.
//
// Solidity: function ReceiveFunds(forProduct uint256) returns(bool)
func (_PayFor *PayForTransactor) ReceiveFunds(opts *bind.TransactOpts, forProduct *big.Int) (*types.Transaction, error) {
	return _PayFor.contract.Transact(opts, "ReceiveFunds", forProduct)
}

// ReceiveFunds is a paid mutator transaction binding the contract method 0x46d37eea.
//
// Solidity: function ReceiveFunds(forProduct uint256) returns(bool)
func (_PayFor *PayForSession) ReceiveFunds(forProduct *big.Int) (*types.Transaction, error) {
	return _PayFor.Contract.ReceiveFunds(&_PayFor.TransactOpts, forProduct)
}

// ReceiveFunds is a paid mutator transaction binding the contract method 0x46d37eea.
//
// Solidity: function ReceiveFunds(forProduct uint256) returns(bool)
func (_PayFor *PayForTransactorSession) ReceiveFunds(forProduct *big.Int) (*types.Transaction, error) {
	return _PayFor.Contract.ReceiveFunds(&_PayFor.TransactOpts, forProduct)
}

// GetNPayments is a paid mutator transaction binding the contract method 0x5e5ef79c.
//
// Solidity: function getNPayments() returns(uint256)
func (_PayFor *PayForTransactor) GetNPayments(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayFor.contract.Transact(opts, "getNPayments")
}

// GetNPayments is a paid mutator transaction binding the contract method 0x5e5ef79c.
//
// Solidity: function getNPayments() returns(uint256)
func (_PayFor *PayForSession) GetNPayments() (*types.Transaction, error) {
	return _PayFor.Contract.GetNPayments(&_PayFor.TransactOpts)
}

// GetNPayments is a paid mutator transaction binding the contract method 0x5e5ef79c.
//
// Solidity: function getNPayments() returns(uint256)
func (_PayFor *PayForTransactorSession) GetNPayments() (*types.Transaction, error) {
	return _PayFor.Contract.GetNPayments(&_PayFor.TransactOpts)
}

// GetPaymentInfo is a paid mutator transaction binding the contract method 0xd810ff0f.
//
// Solidity: function getPaymentInfo(n uint256) returns(address, uint256, uint256)
func (_PayFor *PayForTransactor) GetPaymentInfo(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _PayFor.contract.Transact(opts, "getPaymentInfo", n)
}

// GetPaymentInfo is a paid mutator transaction binding the contract method 0xd810ff0f.
//
// Solidity: function getPaymentInfo(n uint256) returns(address, uint256, uint256)
func (_PayFor *PayForSession) GetPaymentInfo(n *big.Int) (*types.Transaction, error) {
	return _PayFor.Contract.GetPaymentInfo(&_PayFor.TransactOpts, n)
}

// GetPaymentInfo is a paid mutator transaction binding the contract method 0xd810ff0f.
//
// Solidity: function getPaymentInfo(n uint256) returns(address, uint256, uint256)
func (_PayFor *PayForTransactorSession) GetPaymentInfo(n *big.Int) (*types.Transaction, error) {
	return _PayFor.Contract.GetPaymentInfo(&_PayFor.TransactOpts, n)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PayFor *PayForTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayFor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PayFor *PayForSession) RenounceOwnership() (*types.Transaction, error) {
	return _PayFor.Contract.RenounceOwnership(&_PayFor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PayFor *PayForTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PayFor.Contract.RenounceOwnership(&_PayFor.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PayFor *PayForTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PayFor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PayFor *PayForSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PayFor.Contract.TransferOwnership(&_PayFor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PayFor *PayForTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PayFor.Contract.TransferOwnership(&_PayFor.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns(bool)
func (_PayFor *PayForTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PayFor.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns(bool)
func (_PayFor *PayForSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _PayFor.Contract.Withdraw(&_PayFor.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns(bool)
func (_PayFor *PayForTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _PayFor.Contract.Withdraw(&_PayFor.TransactOpts, amount)
}

// PayForOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PayFor contract.
type PayForOwnershipTransferredIterator struct {
	Event *PayForOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PayForOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayForOwnershipTransferred)
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
		it.Event = new(PayForOwnershipTransferred)
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
func (it *PayForOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayForOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayForOwnershipTransferred represents a OwnershipTransferred event raised by the PayFor contract.
type PayForOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PayFor *PayForFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PayForOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PayFor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PayForOwnershipTransferredIterator{contract: _PayFor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PayFor *PayForFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PayForOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PayFor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayForOwnershipTransferred)
				if err := _PayFor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PayForReceivedFundsIterator is returned from FilterReceivedFunds and is used to iterate over the raw logs and unpacked data for ReceivedFunds events raised by the PayFor contract.
type PayForReceivedFundsIterator struct {
	Event *PayForReceivedFunds // Event containing the contract specifics and raw log

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
func (it *PayForReceivedFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayForReceivedFunds)
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
		it.Event = new(PayForReceivedFunds)
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
func (it *PayForReceivedFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayForReceivedFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayForReceivedFunds represents a ReceivedFunds event raised by the PayFor contract.
type PayForReceivedFunds struct {
	Sender      common.Address
	Value       *big.Int
	Application *big.Int
	Loc         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedFunds is a free log retrieval operation binding the contract event 0x5f74095c3cf9cd2b4407992767113d77fc65cb2423a83ecc71977f3ace6b34f4.
//
// Solidity: e ReceivedFunds(sender address, value uint256, application uint256, loc uint256)
func (_PayFor *PayForFilterer) FilterReceivedFunds(opts *bind.FilterOpts) (*PayForReceivedFundsIterator, error) {

	logs, sub, err := _PayFor.contract.FilterLogs(opts, "ReceivedFunds")
	if err != nil {
		return nil, err
	}
	return &PayForReceivedFundsIterator{contract: _PayFor.contract, event: "ReceivedFunds", logs: logs, sub: sub}, nil
}

// WatchReceivedFunds is a free log subscription operation binding the contract event 0x5f74095c3cf9cd2b4407992767113d77fc65cb2423a83ecc71977f3ace6b34f4.
//
// Solidity: e ReceivedFunds(sender address, value uint256, application uint256, loc uint256)
func (_PayFor *PayForFilterer) WatchReceivedFunds(opts *bind.WatchOpts, sink chan<- *PayForReceivedFunds) (event.Subscription, error) {

	logs, sub, err := _PayFor.contract.WatchLogs(opts, "ReceivedFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayForReceivedFunds)
				if err := _PayFor.contract.UnpackLog(event, "ReceivedFunds", log); err != nil {
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

// PayForWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the PayFor contract.
type PayForWithdrawnIterator struct {
	Event *PayForWithdrawn // Event containing the contract specifics and raw log

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
func (it *PayForWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayForWithdrawn)
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
		it.Event = new(PayForWithdrawn)
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
func (it *PayForWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayForWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayForWithdrawn represents a Withdrawn event raised by the PayFor contract.
type PayForWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(to address, amount uint256)
func (_PayFor *PayForFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*PayForWithdrawnIterator, error) {

	logs, sub, err := _PayFor.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &PayForWithdrawnIterator{contract: _PayFor.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(to address, amount uint256)
func (_PayFor *PayForFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *PayForWithdrawn) (event.Subscription, error) {

	logs, sub, err := _PayFor.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayForWithdrawn)
				if err := _PayFor.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
