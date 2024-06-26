// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package spark_price_oracle

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

// SparkPriceOracleMetaData contains all meta data concerning the SparkPriceOracle contract.
var SparkPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseCurrencyUnit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AssetSourceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseCurrencyUnit\",\"type\":\"uint256\"}],\"name\":\"BaseCurrencySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"}],\"name\":\"FallbackOracleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_CURRENCY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_CURRENCY_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"getAssetsPrices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFallbackOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getSourceOfAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"}],\"name\":\"setAssetSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fallbackOracle\",\"type\":\"address\"}],\"name\":\"setFallbackOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SparkPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use SparkPriceOracleMetaData.ABI instead.
var SparkPriceOracleABI = SparkPriceOracleMetaData.ABI

// SparkPriceOracle is an auto generated Go binding around an Ethereum contract.
type SparkPriceOracle struct {
	SparkPriceOracleCaller     // Read-only binding to the contract
	SparkPriceOracleTransactor // Write-only binding to the contract
	SparkPriceOracleFilterer   // Log filterer for contract events
}

// SparkPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SparkPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SparkPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SparkPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SparkPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SparkPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SparkPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SparkPriceOracleSession struct {
	Contract     *SparkPriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SparkPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SparkPriceOracleCallerSession struct {
	Contract *SparkPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SparkPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SparkPriceOracleTransactorSession struct {
	Contract     *SparkPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SparkPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SparkPriceOracleRaw struct {
	Contract *SparkPriceOracle // Generic contract binding to access the raw methods on
}

// SparkPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SparkPriceOracleCallerRaw struct {
	Contract *SparkPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// SparkPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SparkPriceOracleTransactorRaw struct {
	Contract *SparkPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSparkPriceOracle creates a new instance of SparkPriceOracle, bound to a specific deployed contract.
func NewSparkPriceOracle(address common.Address, backend bind.ContractBackend) (*SparkPriceOracle, error) {
	contract, err := bindSparkPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SparkPriceOracle{SparkPriceOracleCaller: SparkPriceOracleCaller{contract: contract}, SparkPriceOracleTransactor: SparkPriceOracleTransactor{contract: contract}, SparkPriceOracleFilterer: SparkPriceOracleFilterer{contract: contract}}, nil
}

// NewSparkPriceOracleCaller creates a new read-only instance of SparkPriceOracle, bound to a specific deployed contract.
func NewSparkPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*SparkPriceOracleCaller, error) {
	contract, err := bindSparkPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SparkPriceOracleCaller{contract: contract}, nil
}

// NewSparkPriceOracleTransactor creates a new write-only instance of SparkPriceOracle, bound to a specific deployed contract.
func NewSparkPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*SparkPriceOracleTransactor, error) {
	contract, err := bindSparkPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SparkPriceOracleTransactor{contract: contract}, nil
}

// NewSparkPriceOracleFilterer creates a new log filterer instance of SparkPriceOracle, bound to a specific deployed contract.
func NewSparkPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*SparkPriceOracleFilterer, error) {
	contract, err := bindSparkPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SparkPriceOracleFilterer{contract: contract}, nil
}

// bindSparkPriceOracle binds a generic wrapper to an already deployed contract.
func bindSparkPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SparkPriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SparkPriceOracle *SparkPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SparkPriceOracle.Contract.SparkPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SparkPriceOracle *SparkPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SparkPriceOracle.Contract.SparkPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SparkPriceOracle *SparkPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SparkPriceOracle.Contract.SparkPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SparkPriceOracle *SparkPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SparkPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SparkPriceOracle *SparkPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SparkPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SparkPriceOracle *SparkPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SparkPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_SparkPriceOracle *SparkPriceOracleCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SparkPriceOracle.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_SparkPriceOracle *SparkPriceOracleSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _SparkPriceOracle.Contract.ADDRESSESPROVIDER(&_SparkPriceOracle.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_SparkPriceOracle *SparkPriceOracleCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _SparkPriceOracle.Contract.ADDRESSESPROVIDER(&_SparkPriceOracle.CallOpts)
}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_SparkPriceOracle *SparkPriceOracleCaller) BASECURRENCY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SparkPriceOracle.contract.Call(opts, &out, "BASE_CURRENCY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_SparkPriceOracle *SparkPriceOracleSession) BASECURRENCY() (common.Address, error) {
	return _SparkPriceOracle.Contract.BASECURRENCY(&_SparkPriceOracle.CallOpts)
}

// BASECURRENCY is a free data retrieval call binding the contract method 0xe19f4700.
//
// Solidity: function BASE_CURRENCY() view returns(address)
func (_SparkPriceOracle *SparkPriceOracleCallerSession) BASECURRENCY() (common.Address, error) {
	return _SparkPriceOracle.Contract.BASECURRENCY(&_SparkPriceOracle.CallOpts)
}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_SparkPriceOracle *SparkPriceOracleCaller) BASECURRENCYUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SparkPriceOracle.contract.Call(opts, &out, "BASE_CURRENCY_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_SparkPriceOracle *SparkPriceOracleSession) BASECURRENCYUNIT() (*big.Int, error) {
	return _SparkPriceOracle.Contract.BASECURRENCYUNIT(&_SparkPriceOracle.CallOpts)
}

// BASECURRENCYUNIT is a free data retrieval call binding the contract method 0x8c89b64f.
//
// Solidity: function BASE_CURRENCY_UNIT() view returns(uint256)
func (_SparkPriceOracle *SparkPriceOracleCallerSession) BASECURRENCYUNIT() (*big.Int, error) {
	return _SparkPriceOracle.Contract.BASECURRENCYUNIT(&_SparkPriceOracle.CallOpts)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_SparkPriceOracle *SparkPriceOracleCaller) GetAssetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SparkPriceOracle.contract.Call(opts, &out, "getAssetPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_SparkPriceOracle *SparkPriceOracleSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _SparkPriceOracle.Contract.GetAssetPrice(&_SparkPriceOracle.CallOpts, asset)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_SparkPriceOracle *SparkPriceOracleCallerSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _SparkPriceOracle.Contract.GetAssetPrice(&_SparkPriceOracle.CallOpts, asset)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_SparkPriceOracle *SparkPriceOracleCaller) GetAssetsPrices(opts *bind.CallOpts, assets []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _SparkPriceOracle.contract.Call(opts, &out, "getAssetsPrices", assets)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_SparkPriceOracle *SparkPriceOracleSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _SparkPriceOracle.Contract.GetAssetsPrices(&_SparkPriceOracle.CallOpts, assets)
}

// GetAssetsPrices is a free data retrieval call binding the contract method 0x9d23d9f2.
//
// Solidity: function getAssetsPrices(address[] assets) view returns(uint256[])
func (_SparkPriceOracle *SparkPriceOracleCallerSession) GetAssetsPrices(assets []common.Address) ([]*big.Int, error) {
	return _SparkPriceOracle.Contract.GetAssetsPrices(&_SparkPriceOracle.CallOpts, assets)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_SparkPriceOracle *SparkPriceOracleCaller) GetFallbackOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SparkPriceOracle.contract.Call(opts, &out, "getFallbackOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_SparkPriceOracle *SparkPriceOracleSession) GetFallbackOracle() (common.Address, error) {
	return _SparkPriceOracle.Contract.GetFallbackOracle(&_SparkPriceOracle.CallOpts)
}

// GetFallbackOracle is a free data retrieval call binding the contract method 0x6210308c.
//
// Solidity: function getFallbackOracle() view returns(address)
func (_SparkPriceOracle *SparkPriceOracleCallerSession) GetFallbackOracle() (common.Address, error) {
	return _SparkPriceOracle.Contract.GetFallbackOracle(&_SparkPriceOracle.CallOpts)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_SparkPriceOracle *SparkPriceOracleCaller) GetSourceOfAsset(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _SparkPriceOracle.contract.Call(opts, &out, "getSourceOfAsset", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_SparkPriceOracle *SparkPriceOracleSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _SparkPriceOracle.Contract.GetSourceOfAsset(&_SparkPriceOracle.CallOpts, asset)
}

// GetSourceOfAsset is a free data retrieval call binding the contract method 0x92bf2be0.
//
// Solidity: function getSourceOfAsset(address asset) view returns(address)
func (_SparkPriceOracle *SparkPriceOracleCallerSession) GetSourceOfAsset(asset common.Address) (common.Address, error) {
	return _SparkPriceOracle.Contract.GetSourceOfAsset(&_SparkPriceOracle.CallOpts, asset)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_SparkPriceOracle *SparkPriceOracleTransactor) SetAssetSources(opts *bind.TransactOpts, assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _SparkPriceOracle.contract.Transact(opts, "setAssetSources", assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_SparkPriceOracle *SparkPriceOracleSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _SparkPriceOracle.Contract.SetAssetSources(&_SparkPriceOracle.TransactOpts, assets, sources)
}

// SetAssetSources is a paid mutator transaction binding the contract method 0xabfd5310.
//
// Solidity: function setAssetSources(address[] assets, address[] sources) returns()
func (_SparkPriceOracle *SparkPriceOracleTransactorSession) SetAssetSources(assets []common.Address, sources []common.Address) (*types.Transaction, error) {
	return _SparkPriceOracle.Contract.SetAssetSources(&_SparkPriceOracle.TransactOpts, assets, sources)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_SparkPriceOracle *SparkPriceOracleTransactor) SetFallbackOracle(opts *bind.TransactOpts, fallbackOracle common.Address) (*types.Transaction, error) {
	return _SparkPriceOracle.contract.Transact(opts, "setFallbackOracle", fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_SparkPriceOracle *SparkPriceOracleSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _SparkPriceOracle.Contract.SetFallbackOracle(&_SparkPriceOracle.TransactOpts, fallbackOracle)
}

// SetFallbackOracle is a paid mutator transaction binding the contract method 0x170aee73.
//
// Solidity: function setFallbackOracle(address fallbackOracle) returns()
func (_SparkPriceOracle *SparkPriceOracleTransactorSession) SetFallbackOracle(fallbackOracle common.Address) (*types.Transaction, error) {
	return _SparkPriceOracle.Contract.SetFallbackOracle(&_SparkPriceOracle.TransactOpts, fallbackOracle)
}

// SparkPriceOracleAssetSourceUpdatedIterator is returned from FilterAssetSourceUpdated and is used to iterate over the raw logs and unpacked data for AssetSourceUpdated events raised by the SparkPriceOracle contract.
type SparkPriceOracleAssetSourceUpdatedIterator struct {
	Event *SparkPriceOracleAssetSourceUpdated // Event containing the contract specifics and raw log

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
func (it *SparkPriceOracleAssetSourceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SparkPriceOracleAssetSourceUpdated)
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
		it.Event = new(SparkPriceOracleAssetSourceUpdated)
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
func (it *SparkPriceOracleAssetSourceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SparkPriceOracleAssetSourceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SparkPriceOracleAssetSourceUpdated represents a AssetSourceUpdated event raised by the SparkPriceOracle contract.
type SparkPriceOracleAssetSourceUpdated struct {
	Asset  common.Address
	Source common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetSourceUpdated is a free log retrieval operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_SparkPriceOracle *SparkPriceOracleFilterer) FilterAssetSourceUpdated(opts *bind.FilterOpts, asset []common.Address, source []common.Address) (*SparkPriceOracleAssetSourceUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _SparkPriceOracle.contract.FilterLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &SparkPriceOracleAssetSourceUpdatedIterator{contract: _SparkPriceOracle.contract, event: "AssetSourceUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetSourceUpdated is a free log subscription operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_SparkPriceOracle *SparkPriceOracleFilterer) WatchAssetSourceUpdated(opts *bind.WatchOpts, sink chan<- *SparkPriceOracleAssetSourceUpdated, asset []common.Address, source []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _SparkPriceOracle.contract.WatchLogs(opts, "AssetSourceUpdated", assetRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SparkPriceOracleAssetSourceUpdated)
				if err := _SparkPriceOracle.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
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

// ParseAssetSourceUpdated is a log parse operation binding the contract event 0x22c5b7b2d8561d39f7f210b6b326a1aa69f15311163082308ac4877db6339dc1.
//
// Solidity: event AssetSourceUpdated(address indexed asset, address indexed source)
func (_SparkPriceOracle *SparkPriceOracleFilterer) ParseAssetSourceUpdated(log types.Log) (*SparkPriceOracleAssetSourceUpdated, error) {
	event := new(SparkPriceOracleAssetSourceUpdated)
	if err := _SparkPriceOracle.contract.UnpackLog(event, "AssetSourceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SparkPriceOracleBaseCurrencySetIterator is returned from FilterBaseCurrencySet and is used to iterate over the raw logs and unpacked data for BaseCurrencySet events raised by the SparkPriceOracle contract.
type SparkPriceOracleBaseCurrencySetIterator struct {
	Event *SparkPriceOracleBaseCurrencySet // Event containing the contract specifics and raw log

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
func (it *SparkPriceOracleBaseCurrencySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SparkPriceOracleBaseCurrencySet)
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
		it.Event = new(SparkPriceOracleBaseCurrencySet)
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
func (it *SparkPriceOracleBaseCurrencySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SparkPriceOracleBaseCurrencySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SparkPriceOracleBaseCurrencySet represents a BaseCurrencySet event raised by the SparkPriceOracle contract.
type SparkPriceOracleBaseCurrencySet struct {
	BaseCurrency     common.Address
	BaseCurrencyUnit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBaseCurrencySet is a free log retrieval operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_SparkPriceOracle *SparkPriceOracleFilterer) FilterBaseCurrencySet(opts *bind.FilterOpts, baseCurrency []common.Address) (*SparkPriceOracleBaseCurrencySetIterator, error) {

	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := _SparkPriceOracle.contract.FilterLogs(opts, "BaseCurrencySet", baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return &SparkPriceOracleBaseCurrencySetIterator{contract: _SparkPriceOracle.contract, event: "BaseCurrencySet", logs: logs, sub: sub}, nil
}

// WatchBaseCurrencySet is a free log subscription operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_SparkPriceOracle *SparkPriceOracleFilterer) WatchBaseCurrencySet(opts *bind.WatchOpts, sink chan<- *SparkPriceOracleBaseCurrencySet, baseCurrency []common.Address) (event.Subscription, error) {

	var baseCurrencyRule []interface{}
	for _, baseCurrencyItem := range baseCurrency {
		baseCurrencyRule = append(baseCurrencyRule, baseCurrencyItem)
	}

	logs, sub, err := _SparkPriceOracle.contract.WatchLogs(opts, "BaseCurrencySet", baseCurrencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SparkPriceOracleBaseCurrencySet)
				if err := _SparkPriceOracle.contract.UnpackLog(event, "BaseCurrencySet", log); err != nil {
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

// ParseBaseCurrencySet is a log parse operation binding the contract event 0xe27c4c1372396a3d15a9922f74f9dfc7c72b1ad6d63868470787249c356454c1.
//
// Solidity: event BaseCurrencySet(address indexed baseCurrency, uint256 baseCurrencyUnit)
func (_SparkPriceOracle *SparkPriceOracleFilterer) ParseBaseCurrencySet(log types.Log) (*SparkPriceOracleBaseCurrencySet, error) {
	event := new(SparkPriceOracleBaseCurrencySet)
	if err := _SparkPriceOracle.contract.UnpackLog(event, "BaseCurrencySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SparkPriceOracleFallbackOracleUpdatedIterator is returned from FilterFallbackOracleUpdated and is used to iterate over the raw logs and unpacked data for FallbackOracleUpdated events raised by the SparkPriceOracle contract.
type SparkPriceOracleFallbackOracleUpdatedIterator struct {
	Event *SparkPriceOracleFallbackOracleUpdated // Event containing the contract specifics and raw log

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
func (it *SparkPriceOracleFallbackOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SparkPriceOracleFallbackOracleUpdated)
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
		it.Event = new(SparkPriceOracleFallbackOracleUpdated)
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
func (it *SparkPriceOracleFallbackOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SparkPriceOracleFallbackOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SparkPriceOracleFallbackOracleUpdated represents a FallbackOracleUpdated event raised by the SparkPriceOracle contract.
type SparkPriceOracleFallbackOracleUpdated struct {
	FallbackOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFallbackOracleUpdated is a free log retrieval operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_SparkPriceOracle *SparkPriceOracleFilterer) FilterFallbackOracleUpdated(opts *bind.FilterOpts, fallbackOracle []common.Address) (*SparkPriceOracleFallbackOracleUpdatedIterator, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _SparkPriceOracle.contract.FilterLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return &SparkPriceOracleFallbackOracleUpdatedIterator{contract: _SparkPriceOracle.contract, event: "FallbackOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchFallbackOracleUpdated is a free log subscription operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_SparkPriceOracle *SparkPriceOracleFilterer) WatchFallbackOracleUpdated(opts *bind.WatchOpts, sink chan<- *SparkPriceOracleFallbackOracleUpdated, fallbackOracle []common.Address) (event.Subscription, error) {

	var fallbackOracleRule []interface{}
	for _, fallbackOracleItem := range fallbackOracle {
		fallbackOracleRule = append(fallbackOracleRule, fallbackOracleItem)
	}

	logs, sub, err := _SparkPriceOracle.contract.WatchLogs(opts, "FallbackOracleUpdated", fallbackOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SparkPriceOracleFallbackOracleUpdated)
				if err := _SparkPriceOracle.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
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

// ParseFallbackOracleUpdated is a log parse operation binding the contract event 0xce7a780d33665b1ea097af5f155e3821b809ecbaa839d3b33aa83ba28168cefb.
//
// Solidity: event FallbackOracleUpdated(address indexed fallbackOracle)
func (_SparkPriceOracle *SparkPriceOracleFilterer) ParseFallbackOracleUpdated(log types.Log) (*SparkPriceOracleFallbackOracleUpdated, error) {
	event := new(SparkPriceOracleFallbackOracleUpdated)
	if err := _SparkPriceOracle.contract.UnpackLog(event, "FallbackOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
