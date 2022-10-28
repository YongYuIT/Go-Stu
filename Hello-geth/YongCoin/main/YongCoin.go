// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// YongCoinMetaData contains all meta data concerning the YongCoin contract.
var YongCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reciver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reciver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YongCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use YongCoinMetaData.ABI instead.
var YongCoinABI = YongCoinMetaData.ABI

// YongCoin is an auto generated Go binding around an Ethereum contract.
type YongCoin struct {
	YongCoinCaller     // Read-only binding to the contract
	YongCoinTransactor // Write-only binding to the contract
	YongCoinFilterer   // Log filterer for contract events
}

// YongCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type YongCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YongCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YongCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YongCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YongCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YongCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YongCoinSession struct {
	Contract     *YongCoin         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YongCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YongCoinCallerSession struct {
	Contract *YongCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// YongCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YongCoinTransactorSession struct {
	Contract     *YongCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// YongCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type YongCoinRaw struct {
	Contract *YongCoin // Generic contract binding to access the raw methods on
}

// YongCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YongCoinCallerRaw struct {
	Contract *YongCoinCaller // Generic read-only contract binding to access the raw methods on
}

// YongCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YongCoinTransactorRaw struct {
	Contract *YongCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYongCoin creates a new instance of YongCoin, bound to a specific deployed contract.
func NewYongCoin(address common.Address, backend bind.ContractBackend) (*YongCoin, error) {
	contract, err := bindYongCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YongCoin{YongCoinCaller: YongCoinCaller{contract: contract}, YongCoinTransactor: YongCoinTransactor{contract: contract}, YongCoinFilterer: YongCoinFilterer{contract: contract}}, nil
}

// NewYongCoinCaller creates a new read-only instance of YongCoin, bound to a specific deployed contract.
func NewYongCoinCaller(address common.Address, caller bind.ContractCaller) (*YongCoinCaller, error) {
	contract, err := bindYongCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YongCoinCaller{contract: contract}, nil
}

// NewYongCoinTransactor creates a new write-only instance of YongCoin, bound to a specific deployed contract.
func NewYongCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*YongCoinTransactor, error) {
	contract, err := bindYongCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YongCoinTransactor{contract: contract}, nil
}

// NewYongCoinFilterer creates a new log filterer instance of YongCoin, bound to a specific deployed contract.
func NewYongCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*YongCoinFilterer, error) {
	contract, err := bindYongCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YongCoinFilterer{contract: contract}, nil
}

// bindYongCoin binds a generic wrapper to an already deployed contract.
func bindYongCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YongCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YongCoin *YongCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YongCoin.Contract.YongCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YongCoin *YongCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YongCoin.Contract.YongCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YongCoin *YongCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YongCoin.Contract.YongCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YongCoin *YongCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YongCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YongCoin *YongCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YongCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YongCoin *YongCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YongCoin.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_YongCoin *YongCoinCaller) Balance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YongCoin.contract.Call(opts, &out, "balance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_YongCoin *YongCoinSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _YongCoin.Contract.Balance(&_YongCoin.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address ) view returns(uint256)
func (_YongCoin *YongCoinCallerSession) Balance(arg0 common.Address) (*big.Int, error) {
	return _YongCoin.Contract.Balance(&_YongCoin.CallOpts, arg0)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_YongCoin *YongCoinCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YongCoin.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_YongCoin *YongCoinSession) Minter() (common.Address, error) {
	return _YongCoin.Contract.Minter(&_YongCoin.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_YongCoin *YongCoinCallerSession) Minter() (common.Address, error) {
	return _YongCoin.Contract.Minter(&_YongCoin.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address reciver, uint256 amount) returns()
func (_YongCoin *YongCoinTransactor) Mint(opts *bind.TransactOpts, reciver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YongCoin.contract.Transact(opts, "mint", reciver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address reciver, uint256 amount) returns()
func (_YongCoin *YongCoinSession) Mint(reciver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YongCoin.Contract.Mint(&_YongCoin.TransactOpts, reciver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address reciver, uint256 amount) returns()
func (_YongCoin *YongCoinTransactorSession) Mint(reciver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YongCoin.Contract.Mint(&_YongCoin.TransactOpts, reciver, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address reciver, uint256 amount) returns()
func (_YongCoin *YongCoinTransactor) Send(opts *bind.TransactOpts, reciver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YongCoin.contract.Transact(opts, "send", reciver, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address reciver, uint256 amount) returns()
func (_YongCoin *YongCoinSession) Send(reciver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YongCoin.Contract.Send(&_YongCoin.TransactOpts, reciver, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(address reciver, uint256 amount) returns()
func (_YongCoin *YongCoinTransactorSession) Send(reciver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YongCoin.Contract.Send(&_YongCoin.TransactOpts, reciver, amount)
}

// YongCoinSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the YongCoin contract.
type YongCoinSentIterator struct {
	Event *YongCoinSent // Event containing the contract specifics and raw log

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
func (it *YongCoinSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YongCoinSent)
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
		it.Event = new(YongCoinSent)
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
func (it *YongCoinSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YongCoinSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YongCoinSent represents a Sent event raised by the YongCoin contract.
type YongCoinSent struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd3345.
//
// Solidity: event Sent(address from, address to, uint256 amount)
func (_YongCoin *YongCoinFilterer) FilterSent(opts *bind.FilterOpts) (*YongCoinSentIterator, error) {

	logs, sub, err := _YongCoin.contract.FilterLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return &YongCoinSentIterator{contract: _YongCoin.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd3345.
//
// Solidity: event Sent(address from, address to, uint256 amount)
func (_YongCoin *YongCoinFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *YongCoinSent) (event.Subscription, error) {

	logs, sub, err := _YongCoin.contract.WatchLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YongCoinSent)
				if err := _YongCoin.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x3990db2d31862302a685e8086b5755072a6e2b5b780af1ee81ece35ee3cd3345.
//
// Solidity: event Sent(address from, address to, uint256 amount)
func (_YongCoin *YongCoinFilterer) ParseSent(log types.Log) (*YongCoinSent, error) {
	event := new(YongCoinSent)
	if err := _YongCoin.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
