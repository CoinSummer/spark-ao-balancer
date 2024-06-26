// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ao

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

// IDistributionPool is an auto generated low-level Go binding around an user-defined struct.
type IDistributionPool struct {
	PayoutStart                  *big.Int
	DecreaseInterval             *big.Int
	WithdrawLockPeriod           *big.Int
	ClaimLockPeriod              *big.Int
	WithdrawLockPeriodAfterStake *big.Int
	InitialReward                *big.Int
	RewardDecrease               *big.Int
	MinimalStake                 *big.Int
	IsPublic                     bool
}

// AoMetaData contains all meta data concerning the Ao contract.
var AoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OverplusBridged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"payoutStart\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"decreaseInterval\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriodAfterStake\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"initialReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDecrease\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIDistribution.Pool\",\"name\":\"pool\",\"type\":\"tuple\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"payoutStart\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"decreaseInterval\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriodAfterStake\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"initialReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDecrease\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIDistribution.Pool\",\"name\":\"pool\",\"type\":\"tuple\"}],\"name\":\"PoolEdited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UserClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"arweaveAddress\",\"type\":\"bytes32\"}],\"name\":\"UserStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"arweaveAddress\",\"type\":\"bytes32\"}],\"name\":\"UserWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aoDistributionWallet_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"payoutStart\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"decreaseInterval\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriodAfterStake\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"initialReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDecrease\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"internalType\":\"structIDistribution.Pool[]\",\"name\":\"poolsInfo_\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"refunderAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackAddress_\",\"type\":\"address\"}],\"name\":\"Distribution_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aoDistributionWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeOverplus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"payoutStart\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"decreaseInterval\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriodAfterStake\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"initialReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDecrease\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"internalType\":\"structIDistribution.Pool\",\"name\":\"pool_\",\"type\":\"tuple\"}],\"name\":\"createPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ejectStakedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"getCurrentUserReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId_\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"startTime_\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"endTime_\",\"type\":\"uint128\"}],\"name\":\"getPeriodReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isNotUpgradeable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overplus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"payoutStart\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"decreaseInterval\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimLockPeriod\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawLockPeriodAfterStake\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"initialReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDecrease\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolsData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"lastUpdate\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDeposited\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeUpgradeability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"arweaveAddress_\",\"type\":\"bytes32\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDepositedInPublicPools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usersData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"lastStake\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"arweaveAddress_\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AoABI is the input ABI used to generate the binding from.
// Deprecated: Use AoMetaData.ABI instead.
var AoABI = AoMetaData.ABI

// Ao is an auto generated Go binding around an Ethereum contract.
type Ao struct {
	AoCaller     // Read-only binding to the contract
	AoTransactor // Write-only binding to the contract
	AoFilterer   // Log filterer for contract events
}

// AoCaller is an auto generated read-only Go binding around an Ethereum contract.
type AoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AoSession struct {
	Contract     *Ao               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AoCallerSession struct {
	Contract *AoCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AoTransactorSession struct {
	Contract     *AoTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AoRaw is an auto generated low-level Go binding around an Ethereum contract.
type AoRaw struct {
	Contract *Ao // Generic contract binding to access the raw methods on
}

// AoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AoCallerRaw struct {
	Contract *AoCaller // Generic read-only contract binding to access the raw methods on
}

// AoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AoTransactorRaw struct {
	Contract *AoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAo creates a new instance of Ao, bound to a specific deployed contract.
func NewAo(address common.Address, backend bind.ContractBackend) (*Ao, error) {
	contract, err := bindAo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ao{AoCaller: AoCaller{contract: contract}, AoTransactor: AoTransactor{contract: contract}, AoFilterer: AoFilterer{contract: contract}}, nil
}

// NewAoCaller creates a new read-only instance of Ao, bound to a specific deployed contract.
func NewAoCaller(address common.Address, caller bind.ContractCaller) (*AoCaller, error) {
	contract, err := bindAo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AoCaller{contract: contract}, nil
}

// NewAoTransactor creates a new write-only instance of Ao, bound to a specific deployed contract.
func NewAoTransactor(address common.Address, transactor bind.ContractTransactor) (*AoTransactor, error) {
	contract, err := bindAo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AoTransactor{contract: contract}, nil
}

// NewAoFilterer creates a new log filterer instance of Ao, bound to a specific deployed contract.
func NewAoFilterer(address common.Address, filterer bind.ContractFilterer) (*AoFilterer, error) {
	contract, err := bindAo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AoFilterer{contract: contract}, nil
}

// bindAo binds a generic wrapper to an already deployed contract.
func bindAo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ao *AoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ao.Contract.AoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ao *AoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ao.Contract.AoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ao *AoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ao.Contract.AoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ao *AoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ao.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ao *AoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ao.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ao *AoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ao.Contract.contract.Transact(opts, method, params...)
}

// AoDistributionWallet is a free data retrieval call binding the contract method 0xa06b2a39.
//
// Solidity: function aoDistributionWallet() view returns(address)
func (_Ao *AoCaller) AoDistributionWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "aoDistributionWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AoDistributionWallet is a free data retrieval call binding the contract method 0xa06b2a39.
//
// Solidity: function aoDistributionWallet() view returns(address)
func (_Ao *AoSession) AoDistributionWallet() (common.Address, error) {
	return _Ao.Contract.AoDistributionWallet(&_Ao.CallOpts)
}

// AoDistributionWallet is a free data retrieval call binding the contract method 0xa06b2a39.
//
// Solidity: function aoDistributionWallet() view returns(address)
func (_Ao *AoCallerSession) AoDistributionWallet() (common.Address, error) {
	return _Ao.Contract.AoDistributionWallet(&_Ao.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_Ao *AoCaller) DepositToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "depositToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_Ao *AoSession) DepositToken() (common.Address, error) {
	return _Ao.Contract.DepositToken(&_Ao.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_Ao *AoCallerSession) DepositToken() (common.Address, error) {
	return _Ao.Contract.DepositToken(&_Ao.CallOpts)
}

// GetCurrentUserReward is a free data retrieval call binding the contract method 0x8468a4c4.
//
// Solidity: function getCurrentUserReward(uint256 poolId_, address user_) view returns(uint256)
func (_Ao *AoCaller) GetCurrentUserReward(opts *bind.CallOpts, poolId_ *big.Int, user_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "getCurrentUserReward", poolId_, user_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentUserReward is a free data retrieval call binding the contract method 0x8468a4c4.
//
// Solidity: function getCurrentUserReward(uint256 poolId_, address user_) view returns(uint256)
func (_Ao *AoSession) GetCurrentUserReward(poolId_ *big.Int, user_ common.Address) (*big.Int, error) {
	return _Ao.Contract.GetCurrentUserReward(&_Ao.CallOpts, poolId_, user_)
}

// GetCurrentUserReward is a free data retrieval call binding the contract method 0x8468a4c4.
//
// Solidity: function getCurrentUserReward(uint256 poolId_, address user_) view returns(uint256)
func (_Ao *AoCallerSession) GetCurrentUserReward(poolId_ *big.Int, user_ common.Address) (*big.Int, error) {
	return _Ao.Contract.GetCurrentUserReward(&_Ao.CallOpts, poolId_, user_)
}

// GetPeriodReward is a free data retrieval call binding the contract method 0x535f583f.
//
// Solidity: function getPeriodReward(uint256 poolId_, uint128 startTime_, uint128 endTime_) view returns(uint256)
func (_Ao *AoCaller) GetPeriodReward(opts *bind.CallOpts, poolId_ *big.Int, startTime_ *big.Int, endTime_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "getPeriodReward", poolId_, startTime_, endTime_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPeriodReward is a free data retrieval call binding the contract method 0x535f583f.
//
// Solidity: function getPeriodReward(uint256 poolId_, uint128 startTime_, uint128 endTime_) view returns(uint256)
func (_Ao *AoSession) GetPeriodReward(poolId_ *big.Int, startTime_ *big.Int, endTime_ *big.Int) (*big.Int, error) {
	return _Ao.Contract.GetPeriodReward(&_Ao.CallOpts, poolId_, startTime_, endTime_)
}

// GetPeriodReward is a free data retrieval call binding the contract method 0x535f583f.
//
// Solidity: function getPeriodReward(uint256 poolId_, uint128 startTime_, uint128 endTime_) view returns(uint256)
func (_Ao *AoCallerSession) GetPeriodReward(poolId_ *big.Int, startTime_ *big.Int, endTime_ *big.Int) (*big.Int, error) {
	return _Ao.Contract.GetPeriodReward(&_Ao.CallOpts, poolId_, startTime_, endTime_)
}

// IsNotUpgradeable is a free data retrieval call binding the contract method 0xe0b4b386.
//
// Solidity: function isNotUpgradeable() view returns(bool)
func (_Ao *AoCaller) IsNotUpgradeable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "isNotUpgradeable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNotUpgradeable is a free data retrieval call binding the contract method 0xe0b4b386.
//
// Solidity: function isNotUpgradeable() view returns(bool)
func (_Ao *AoSession) IsNotUpgradeable() (bool, error) {
	return _Ao.Contract.IsNotUpgradeable(&_Ao.CallOpts)
}

// IsNotUpgradeable is a free data retrieval call binding the contract method 0xe0b4b386.
//
// Solidity: function isNotUpgradeable() view returns(bool)
func (_Ao *AoCallerSession) IsNotUpgradeable() (bool, error) {
	return _Ao.Contract.IsNotUpgradeable(&_Ao.CallOpts)
}

// Overplus is a free data retrieval call binding the contract method 0x78df87b2.
//
// Solidity: function overplus() view returns(uint256)
func (_Ao *AoCaller) Overplus(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "overplus")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Overplus is a free data retrieval call binding the contract method 0x78df87b2.
//
// Solidity: function overplus() view returns(uint256)
func (_Ao *AoSession) Overplus() (*big.Int, error) {
	return _Ao.Contract.Overplus(&_Ao.CallOpts)
}

// Overplus is a free data retrieval call binding the contract method 0x78df87b2.
//
// Solidity: function overplus() view returns(uint256)
func (_Ao *AoCallerSession) Overplus() (*big.Int, error) {
	return _Ao.Contract.Overplus(&_Ao.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ao *AoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ao *AoSession) Owner() (common.Address, error) {
	return _Ao.Contract.Owner(&_Ao.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ao *AoCallerSession) Owner() (common.Address, error) {
	return _Ao.Contract.Owner(&_Ao.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint128 payoutStart, uint128 decreaseInterval, uint128 withdrawLockPeriod, uint128 claimLockPeriod, uint128 withdrawLockPeriodAfterStake, uint256 initialReward, uint256 rewardDecrease, uint256 minimalStake, bool isPublic)
func (_Ao *AoCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PayoutStart                  *big.Int
	DecreaseInterval             *big.Int
	WithdrawLockPeriod           *big.Int
	ClaimLockPeriod              *big.Int
	WithdrawLockPeriodAfterStake *big.Int
	InitialReward                *big.Int
	RewardDecrease               *big.Int
	MinimalStake                 *big.Int
	IsPublic                     bool
}, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "pools", arg0)

	outstruct := new(struct {
		PayoutStart                  *big.Int
		DecreaseInterval             *big.Int
		WithdrawLockPeriod           *big.Int
		ClaimLockPeriod              *big.Int
		WithdrawLockPeriodAfterStake *big.Int
		InitialReward                *big.Int
		RewardDecrease               *big.Int
		MinimalStake                 *big.Int
		IsPublic                     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PayoutStart = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DecreaseInterval = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WithdrawLockPeriod = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ClaimLockPeriod = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.WithdrawLockPeriodAfterStake = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.InitialReward = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RewardDecrease = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MinimalStake = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.IsPublic = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint128 payoutStart, uint128 decreaseInterval, uint128 withdrawLockPeriod, uint128 claimLockPeriod, uint128 withdrawLockPeriodAfterStake, uint256 initialReward, uint256 rewardDecrease, uint256 minimalStake, bool isPublic)
func (_Ao *AoSession) Pools(arg0 *big.Int) (struct {
	PayoutStart                  *big.Int
	DecreaseInterval             *big.Int
	WithdrawLockPeriod           *big.Int
	ClaimLockPeriod              *big.Int
	WithdrawLockPeriodAfterStake *big.Int
	InitialReward                *big.Int
	RewardDecrease               *big.Int
	MinimalStake                 *big.Int
	IsPublic                     bool
}, error) {
	return _Ao.Contract.Pools(&_Ao.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint128 payoutStart, uint128 decreaseInterval, uint128 withdrawLockPeriod, uint128 claimLockPeriod, uint128 withdrawLockPeriodAfterStake, uint256 initialReward, uint256 rewardDecrease, uint256 minimalStake, bool isPublic)
func (_Ao *AoCallerSession) Pools(arg0 *big.Int) (struct {
	PayoutStart                  *big.Int
	DecreaseInterval             *big.Int
	WithdrawLockPeriod           *big.Int
	ClaimLockPeriod              *big.Int
	WithdrawLockPeriodAfterStake *big.Int
	InitialReward                *big.Int
	RewardDecrease               *big.Int
	MinimalStake                 *big.Int
	IsPublic                     bool
}, error) {
	return _Ao.Contract.Pools(&_Ao.CallOpts, arg0)
}

// PoolsData is a free data retrieval call binding the contract method 0xf343e858.
//
// Solidity: function poolsData(uint256 ) view returns(uint128 lastUpdate, uint256 rate, uint256 totalDeposited)
func (_Ao *AoCaller) PoolsData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LastUpdate     *big.Int
	Rate           *big.Int
	TotalDeposited *big.Int
}, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "poolsData", arg0)

	outstruct := new(struct {
		LastUpdate     *big.Int
		Rate           *big.Int
		TotalDeposited *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastUpdate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalDeposited = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolsData is a free data retrieval call binding the contract method 0xf343e858.
//
// Solidity: function poolsData(uint256 ) view returns(uint128 lastUpdate, uint256 rate, uint256 totalDeposited)
func (_Ao *AoSession) PoolsData(arg0 *big.Int) (struct {
	LastUpdate     *big.Int
	Rate           *big.Int
	TotalDeposited *big.Int
}, error) {
	return _Ao.Contract.PoolsData(&_Ao.CallOpts, arg0)
}

// PoolsData is a free data retrieval call binding the contract method 0xf343e858.
//
// Solidity: function poolsData(uint256 ) view returns(uint128 lastUpdate, uint256 rate, uint256 totalDeposited)
func (_Ao *AoCallerSession) PoolsData(arg0 *big.Int) (struct {
	LastUpdate     *big.Int
	Rate           *big.Int
	TotalDeposited *big.Int
}, error) {
	return _Ao.Contract.PoolsData(&_Ao.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ao *AoCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ao *AoSession) ProxiableUUID() ([32]byte, error) {
	return _Ao.Contract.ProxiableUUID(&_Ao.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ao *AoCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Ao.Contract.ProxiableUUID(&_Ao.CallOpts)
}

// TotalDepositedInPublicPools is a free data retrieval call binding the contract method 0xd2ba5e3a.
//
// Solidity: function totalDepositedInPublicPools() view returns(uint256)
func (_Ao *AoCaller) TotalDepositedInPublicPools(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "totalDepositedInPublicPools")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositedInPublicPools is a free data retrieval call binding the contract method 0xd2ba5e3a.
//
// Solidity: function totalDepositedInPublicPools() view returns(uint256)
func (_Ao *AoSession) TotalDepositedInPublicPools() (*big.Int, error) {
	return _Ao.Contract.TotalDepositedInPublicPools(&_Ao.CallOpts)
}

// TotalDepositedInPublicPools is a free data retrieval call binding the contract method 0xd2ba5e3a.
//
// Solidity: function totalDepositedInPublicPools() view returns(uint256)
func (_Ao *AoCallerSession) TotalDepositedInPublicPools() (*big.Int, error) {
	return _Ao.Contract.TotalDepositedInPublicPools(&_Ao.CallOpts)
}

// UsersData is a free data retrieval call binding the contract method 0x30dc6308.
//
// Solidity: function usersData(address , uint256 ) view returns(uint128 lastStake, uint256 deposited, uint256 rate, uint256 pendingRewards)
func (_Ao *AoCaller) UsersData(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LastStake      *big.Int
	Deposited      *big.Int
	Rate           *big.Int
	PendingRewards *big.Int
}, error) {
	var out []interface{}
	err := _Ao.contract.Call(opts, &out, "usersData", arg0, arg1)

	outstruct := new(struct {
		LastStake      *big.Int
		Deposited      *big.Int
		Rate           *big.Int
		PendingRewards *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Deposited = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PendingRewards = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UsersData is a free data retrieval call binding the contract method 0x30dc6308.
//
// Solidity: function usersData(address , uint256 ) view returns(uint128 lastStake, uint256 deposited, uint256 rate, uint256 pendingRewards)
func (_Ao *AoSession) UsersData(arg0 common.Address, arg1 *big.Int) (struct {
	LastStake      *big.Int
	Deposited      *big.Int
	Rate           *big.Int
	PendingRewards *big.Int
}, error) {
	return _Ao.Contract.UsersData(&_Ao.CallOpts, arg0, arg1)
}

// UsersData is a free data retrieval call binding the contract method 0x30dc6308.
//
// Solidity: function usersData(address , uint256 ) view returns(uint128 lastStake, uint256 deposited, uint256 rate, uint256 pendingRewards)
func (_Ao *AoCallerSession) UsersData(arg0 common.Address, arg1 *big.Int) (struct {
	LastStake      *big.Int
	Deposited      *big.Int
	Rate           *big.Int
	PendingRewards *big.Int
}, error) {
	return _Ao.Contract.UsersData(&_Ao.CallOpts, arg0, arg1)
}

// DistributionInit is a paid mutator transaction binding the contract method 0xf3b122e0.
//
// Solidity: function Distribution_init(address depositToken_, address aoDistributionWallet_, (uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool)[] poolsInfo_, address refunderAddress_, address fallbackAddress_) returns()
func (_Ao *AoTransactor) DistributionInit(opts *bind.TransactOpts, depositToken_ common.Address, aoDistributionWallet_ common.Address, poolsInfo_ []IDistributionPool, refunderAddress_ common.Address, fallbackAddress_ common.Address) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "Distribution_init", depositToken_, aoDistributionWallet_, poolsInfo_, refunderAddress_, fallbackAddress_)
}

// DistributionInit is a paid mutator transaction binding the contract method 0xf3b122e0.
//
// Solidity: function Distribution_init(address depositToken_, address aoDistributionWallet_, (uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool)[] poolsInfo_, address refunderAddress_, address fallbackAddress_) returns()
func (_Ao *AoSession) DistributionInit(depositToken_ common.Address, aoDistributionWallet_ common.Address, poolsInfo_ []IDistributionPool, refunderAddress_ common.Address, fallbackAddress_ common.Address) (*types.Transaction, error) {
	return _Ao.Contract.DistributionInit(&_Ao.TransactOpts, depositToken_, aoDistributionWallet_, poolsInfo_, refunderAddress_, fallbackAddress_)
}

// DistributionInit is a paid mutator transaction binding the contract method 0xf3b122e0.
//
// Solidity: function Distribution_init(address depositToken_, address aoDistributionWallet_, (uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool)[] poolsInfo_, address refunderAddress_, address fallbackAddress_) returns()
func (_Ao *AoTransactorSession) DistributionInit(depositToken_ common.Address, aoDistributionWallet_ common.Address, poolsInfo_ []IDistributionPool, refunderAddress_ common.Address, fallbackAddress_ common.Address) (*types.Transaction, error) {
	return _Ao.Contract.DistributionInit(&_Ao.TransactOpts, depositToken_, aoDistributionWallet_, poolsInfo_, refunderAddress_, fallbackAddress_)
}

// BridgeOverplus is a paid mutator transaction binding the contract method 0x8901579e.
//
// Solidity: function bridgeOverplus() returns()
func (_Ao *AoTransactor) BridgeOverplus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "bridgeOverplus")
}

// BridgeOverplus is a paid mutator transaction binding the contract method 0x8901579e.
//
// Solidity: function bridgeOverplus() returns()
func (_Ao *AoSession) BridgeOverplus() (*types.Transaction, error) {
	return _Ao.Contract.BridgeOverplus(&_Ao.TransactOpts)
}

// BridgeOverplus is a paid mutator transaction binding the contract method 0x8901579e.
//
// Solidity: function bridgeOverplus() returns()
func (_Ao *AoTransactorSession) BridgeOverplus() (*types.Transaction, error) {
	return _Ao.Contract.BridgeOverplus(&_Ao.TransactOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x5bd2e387.
//
// Solidity: function createPool((uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool) pool_) returns()
func (_Ao *AoTransactor) CreatePool(opts *bind.TransactOpts, pool_ IDistributionPool) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "createPool", pool_)
}

// CreatePool is a paid mutator transaction binding the contract method 0x5bd2e387.
//
// Solidity: function createPool((uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool) pool_) returns()
func (_Ao *AoSession) CreatePool(pool_ IDistributionPool) (*types.Transaction, error) {
	return _Ao.Contract.CreatePool(&_Ao.TransactOpts, pool_)
}

// CreatePool is a paid mutator transaction binding the contract method 0x5bd2e387.
//
// Solidity: function createPool((uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool) pool_) returns()
func (_Ao *AoTransactorSession) CreatePool(pool_ IDistributionPool) (*types.Transaction, error) {
	return _Ao.Contract.CreatePool(&_Ao.TransactOpts, pool_)
}

// EjectStakedFunds is a paid mutator transaction binding the contract method 0xdd2518f3.
//
// Solidity: function ejectStakedFunds(uint256 poolId, address user) returns()
func (_Ao *AoTransactor) EjectStakedFunds(opts *bind.TransactOpts, poolId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "ejectStakedFunds", poolId, user)
}

// EjectStakedFunds is a paid mutator transaction binding the contract method 0xdd2518f3.
//
// Solidity: function ejectStakedFunds(uint256 poolId, address user) returns()
func (_Ao *AoSession) EjectStakedFunds(poolId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Ao.Contract.EjectStakedFunds(&_Ao.TransactOpts, poolId, user)
}

// EjectStakedFunds is a paid mutator transaction binding the contract method 0xdd2518f3.
//
// Solidity: function ejectStakedFunds(uint256 poolId, address user) returns()
func (_Ao *AoTransactorSession) EjectStakedFunds(poolId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Ao.Contract.EjectStakedFunds(&_Ao.TransactOpts, poolId, user)
}

// RemoveUpgradeability is a paid mutator transaction binding the contract method 0xbeee9fa9.
//
// Solidity: function removeUpgradeability() returns()
func (_Ao *AoTransactor) RemoveUpgradeability(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "removeUpgradeability")
}

// RemoveUpgradeability is a paid mutator transaction binding the contract method 0xbeee9fa9.
//
// Solidity: function removeUpgradeability() returns()
func (_Ao *AoSession) RemoveUpgradeability() (*types.Transaction, error) {
	return _Ao.Contract.RemoveUpgradeability(&_Ao.TransactOpts)
}

// RemoveUpgradeability is a paid mutator transaction binding the contract method 0xbeee9fa9.
//
// Solidity: function removeUpgradeability() returns()
func (_Ao *AoTransactorSession) RemoveUpgradeability() (*types.Transaction, error) {
	return _Ao.Contract.RemoveUpgradeability(&_Ao.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ao *AoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ao *AoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ao.Contract.RenounceOwnership(&_Ao.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ao *AoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ao.Contract.RenounceOwnership(&_Ao.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3d0a42e7.
//
// Solidity: function stake(uint256 poolId_, uint256 amount_, bytes32 arweaveAddress_) returns()
func (_Ao *AoTransactor) Stake(opts *bind.TransactOpts, poolId_ *big.Int, amount_ *big.Int, arweaveAddress_ [32]byte) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "stake", poolId_, amount_, arweaveAddress_)
}

// Stake is a paid mutator transaction binding the contract method 0x3d0a42e7.
//
// Solidity: function stake(uint256 poolId_, uint256 amount_, bytes32 arweaveAddress_) returns()
func (_Ao *AoSession) Stake(poolId_ *big.Int, amount_ *big.Int, arweaveAddress_ [32]byte) (*types.Transaction, error) {
	return _Ao.Contract.Stake(&_Ao.TransactOpts, poolId_, amount_, arweaveAddress_)
}

// Stake is a paid mutator transaction binding the contract method 0x3d0a42e7.
//
// Solidity: function stake(uint256 poolId_, uint256 amount_, bytes32 arweaveAddress_) returns()
func (_Ao *AoTransactorSession) Stake(poolId_ *big.Int, amount_ *big.Int, arweaveAddress_ [32]byte) (*types.Transaction, error) {
	return _Ao.Contract.Stake(&_Ao.TransactOpts, poolId_, amount_, arweaveAddress_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ao *AoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ao *AoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ao.Contract.TransferOwnership(&_Ao.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ao *AoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ao.Contract.TransferOwnership(&_Ao.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Ao *AoTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Ao *AoSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Ao.Contract.UpgradeTo(&_Ao.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Ao *AoTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Ao.Contract.UpgradeTo(&_Ao.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ao *AoTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ao *AoSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ao.Contract.UpgradeToAndCall(&_Ao.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ao *AoTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ao.Contract.UpgradeToAndCall(&_Ao.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3756c011.
//
// Solidity: function withdraw(uint256 poolId_, uint256 amount_, bytes32 arweaveAddress_) returns()
func (_Ao *AoTransactor) Withdraw(opts *bind.TransactOpts, poolId_ *big.Int, amount_ *big.Int, arweaveAddress_ [32]byte) (*types.Transaction, error) {
	return _Ao.contract.Transact(opts, "withdraw", poolId_, amount_, arweaveAddress_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3756c011.
//
// Solidity: function withdraw(uint256 poolId_, uint256 amount_, bytes32 arweaveAddress_) returns()
func (_Ao *AoSession) Withdraw(poolId_ *big.Int, amount_ *big.Int, arweaveAddress_ [32]byte) (*types.Transaction, error) {
	return _Ao.Contract.Withdraw(&_Ao.TransactOpts, poolId_, amount_, arweaveAddress_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3756c011.
//
// Solidity: function withdraw(uint256 poolId_, uint256 amount_, bytes32 arweaveAddress_) returns()
func (_Ao *AoTransactorSession) Withdraw(poolId_ *big.Int, amount_ *big.Int, arweaveAddress_ [32]byte) (*types.Transaction, error) {
	return _Ao.Contract.Withdraw(&_Ao.TransactOpts, poolId_, amount_, arweaveAddress_)
}

// AoAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Ao contract.
type AoAdminChangedIterator struct {
	Event *AoAdminChanged // Event containing the contract specifics and raw log

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
func (it *AoAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoAdminChanged)
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
		it.Event = new(AoAdminChanged)
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
func (it *AoAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoAdminChanged represents a AdminChanged event raised by the Ao contract.
type AoAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Ao *AoFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AoAdminChangedIterator, error) {

	logs, sub, err := _Ao.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AoAdminChangedIterator{contract: _Ao.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Ao *AoFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AoAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Ao.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoAdminChanged)
				if err := _Ao.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Ao *AoFilterer) ParseAdminChanged(log types.Log) (*AoAdminChanged, error) {
	event := new(AoAdminChanged)
	if err := _Ao.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Ao contract.
type AoBeaconUpgradedIterator struct {
	Event *AoBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AoBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoBeaconUpgraded)
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
		it.Event = new(AoBeaconUpgraded)
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
func (it *AoBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoBeaconUpgraded represents a BeaconUpgraded event raised by the Ao contract.
type AoBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Ao *AoFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AoBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Ao.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AoBeaconUpgradedIterator{contract: _Ao.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Ao *AoFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AoBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Ao.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoBeaconUpgraded)
				if err := _Ao.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Ao *AoFilterer) ParseBeaconUpgraded(log types.Log) (*AoBeaconUpgraded, error) {
	event := new(AoBeaconUpgraded)
	if err := _Ao.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ao contract.
type AoInitializedIterator struct {
	Event *AoInitialized // Event containing the contract specifics and raw log

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
func (it *AoInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoInitialized)
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
		it.Event = new(AoInitialized)
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
func (it *AoInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoInitialized represents a Initialized event raised by the Ao contract.
type AoInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ao *AoFilterer) FilterInitialized(opts *bind.FilterOpts) (*AoInitializedIterator, error) {

	logs, sub, err := _Ao.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AoInitializedIterator{contract: _Ao.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ao *AoFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AoInitialized) (event.Subscription, error) {

	logs, sub, err := _Ao.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoInitialized)
				if err := _Ao.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ao *AoFilterer) ParseInitialized(log types.Log) (*AoInitialized, error) {
	event := new(AoInitialized)
	if err := _Ao.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoOverplusBridgedIterator is returned from FilterOverplusBridged and is used to iterate over the raw logs and unpacked data for OverplusBridged events raised by the Ao contract.
type AoOverplusBridgedIterator struct {
	Event *AoOverplusBridged // Event containing the contract specifics and raw log

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
func (it *AoOverplusBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoOverplusBridged)
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
		it.Event = new(AoOverplusBridged)
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
func (it *AoOverplusBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoOverplusBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoOverplusBridged represents a OverplusBridged event raised by the Ao contract.
type AoOverplusBridged struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOverplusBridged is a free log retrieval operation binding the contract event 0x1c1beeacf263316031be95f61cdcab6d052033248371579d5a954667ee9df8ad.
//
// Solidity: event OverplusBridged(uint256 amount)
func (_Ao *AoFilterer) FilterOverplusBridged(opts *bind.FilterOpts) (*AoOverplusBridgedIterator, error) {

	logs, sub, err := _Ao.contract.FilterLogs(opts, "OverplusBridged")
	if err != nil {
		return nil, err
	}
	return &AoOverplusBridgedIterator{contract: _Ao.contract, event: "OverplusBridged", logs: logs, sub: sub}, nil
}

// WatchOverplusBridged is a free log subscription operation binding the contract event 0x1c1beeacf263316031be95f61cdcab6d052033248371579d5a954667ee9df8ad.
//
// Solidity: event OverplusBridged(uint256 amount)
func (_Ao *AoFilterer) WatchOverplusBridged(opts *bind.WatchOpts, sink chan<- *AoOverplusBridged) (event.Subscription, error) {

	logs, sub, err := _Ao.contract.WatchLogs(opts, "OverplusBridged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoOverplusBridged)
				if err := _Ao.contract.UnpackLog(event, "OverplusBridged", log); err != nil {
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

// ParseOverplusBridged is a log parse operation binding the contract event 0x1c1beeacf263316031be95f61cdcab6d052033248371579d5a954667ee9df8ad.
//
// Solidity: event OverplusBridged(uint256 amount)
func (_Ao *AoFilterer) ParseOverplusBridged(log types.Log) (*AoOverplusBridged, error) {
	event := new(AoOverplusBridged)
	if err := _Ao.contract.UnpackLog(event, "OverplusBridged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ao contract.
type AoOwnershipTransferredIterator struct {
	Event *AoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoOwnershipTransferred)
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
		it.Event = new(AoOwnershipTransferred)
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
func (it *AoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoOwnershipTransferred represents a OwnershipTransferred event raised by the Ao contract.
type AoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ao *AoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ao.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AoOwnershipTransferredIterator{contract: _Ao.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ao *AoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ao.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoOwnershipTransferred)
				if err := _Ao.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ao *AoFilterer) ParseOwnershipTransferred(log types.Log) (*AoOwnershipTransferred, error) {
	event := new(AoOwnershipTransferred)
	if err := _Ao.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Ao contract.
type AoPoolCreatedIterator struct {
	Event *AoPoolCreated // Event containing the contract specifics and raw log

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
func (it *AoPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoPoolCreated)
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
		it.Event = new(AoPoolCreated)
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
func (it *AoPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoPoolCreated represents a PoolCreated event raised by the Ao contract.
type AoPoolCreated struct {
	PoolId *big.Int
	Pool   IDistributionPool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x9b9b09f76d4db7c7b0d783b9ba64d003e95462ef13cbb7cb782d74b085548cac.
//
// Solidity: event PoolCreated(uint256 indexed poolId, (uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool) pool)
func (_Ao *AoFilterer) FilterPoolCreated(opts *bind.FilterOpts, poolId []*big.Int) (*AoPoolCreatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Ao.contract.FilterLogs(opts, "PoolCreated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &AoPoolCreatedIterator{contract: _Ao.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x9b9b09f76d4db7c7b0d783b9ba64d003e95462ef13cbb7cb782d74b085548cac.
//
// Solidity: event PoolCreated(uint256 indexed poolId, (uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool) pool)
func (_Ao *AoFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *AoPoolCreated, poolId []*big.Int) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Ao.contract.WatchLogs(opts, "PoolCreated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoPoolCreated)
				if err := _Ao.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x9b9b09f76d4db7c7b0d783b9ba64d003e95462ef13cbb7cb782d74b085548cac.
//
// Solidity: event PoolCreated(uint256 indexed poolId, (uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool) pool)
func (_Ao *AoFilterer) ParsePoolCreated(log types.Log) (*AoPoolCreated, error) {
	event := new(AoPoolCreated)
	if err := _Ao.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoPoolEditedIterator is returned from FilterPoolEdited and is used to iterate over the raw logs and unpacked data for PoolEdited events raised by the Ao contract.
type AoPoolEditedIterator struct {
	Event *AoPoolEdited // Event containing the contract specifics and raw log

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
func (it *AoPoolEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoPoolEdited)
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
		it.Event = new(AoPoolEdited)
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
func (it *AoPoolEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoPoolEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoPoolEdited represents a PoolEdited event raised by the Ao contract.
type AoPoolEdited struct {
	PoolId *big.Int
	Pool   IDistributionPool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolEdited is a free log retrieval operation binding the contract event 0x3a267bddb165bbfddc7b47b4a65707d7b637e98f445a07ba27ba471b9e0bea54.
//
// Solidity: event PoolEdited(uint256 indexed poolId, (uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool) pool)
func (_Ao *AoFilterer) FilterPoolEdited(opts *bind.FilterOpts, poolId []*big.Int) (*AoPoolEditedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Ao.contract.FilterLogs(opts, "PoolEdited", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &AoPoolEditedIterator{contract: _Ao.contract, event: "PoolEdited", logs: logs, sub: sub}, nil
}

// WatchPoolEdited is a free log subscription operation binding the contract event 0x3a267bddb165bbfddc7b47b4a65707d7b637e98f445a07ba27ba471b9e0bea54.
//
// Solidity: event PoolEdited(uint256 indexed poolId, (uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool) pool)
func (_Ao *AoFilterer) WatchPoolEdited(opts *bind.WatchOpts, sink chan<- *AoPoolEdited, poolId []*big.Int) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Ao.contract.WatchLogs(opts, "PoolEdited", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoPoolEdited)
				if err := _Ao.contract.UnpackLog(event, "PoolEdited", log); err != nil {
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

// ParsePoolEdited is a log parse operation binding the contract event 0x3a267bddb165bbfddc7b47b4a65707d7b637e98f445a07ba27ba471b9e0bea54.
//
// Solidity: event PoolEdited(uint256 indexed poolId, (uint128,uint128,uint128,uint128,uint128,uint256,uint256,uint256,bool) pool)
func (_Ao *AoFilterer) ParsePoolEdited(log types.Log) (*AoPoolEdited, error) {
	event := new(AoPoolEdited)
	if err := _Ao.contract.UnpackLog(event, "PoolEdited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Ao contract.
type AoUpgradedIterator struct {
	Event *AoUpgraded // Event containing the contract specifics and raw log

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
func (it *AoUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoUpgraded)
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
		it.Event = new(AoUpgraded)
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
func (it *AoUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoUpgraded represents a Upgraded event raised by the Ao contract.
type AoUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ao *AoFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AoUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Ao.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AoUpgradedIterator{contract: _Ao.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ao *AoFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AoUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Ao.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoUpgraded)
				if err := _Ao.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ao *AoFilterer) ParseUpgraded(log types.Log) (*AoUpgraded, error) {
	event := new(AoUpgraded)
	if err := _Ao.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoUserClaimedIterator is returned from FilterUserClaimed and is used to iterate over the raw logs and unpacked data for UserClaimed events raised by the Ao contract.
type AoUserClaimedIterator struct {
	Event *AoUserClaimed // Event containing the contract specifics and raw log

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
func (it *AoUserClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoUserClaimed)
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
		it.Event = new(AoUserClaimed)
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
func (it *AoUserClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoUserClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoUserClaimed represents a UserClaimed event raised by the Ao contract.
type AoUserClaimed struct {
	PoolId   *big.Int
	User     common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUserClaimed is a free log retrieval operation binding the contract event 0xace6f3f8956413e2875b9070e2616d13687dfb251cf63b343028c32822dfa263.
//
// Solidity: event UserClaimed(uint256 indexed poolId, address indexed user, address receiver, uint256 amount)
func (_Ao *AoFilterer) FilterUserClaimed(opts *bind.FilterOpts, poolId []*big.Int, user []common.Address) (*AoUserClaimedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ao.contract.FilterLogs(opts, "UserClaimed", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AoUserClaimedIterator{contract: _Ao.contract, event: "UserClaimed", logs: logs, sub: sub}, nil
}

// WatchUserClaimed is a free log subscription operation binding the contract event 0xace6f3f8956413e2875b9070e2616d13687dfb251cf63b343028c32822dfa263.
//
// Solidity: event UserClaimed(uint256 indexed poolId, address indexed user, address receiver, uint256 amount)
func (_Ao *AoFilterer) WatchUserClaimed(opts *bind.WatchOpts, sink chan<- *AoUserClaimed, poolId []*big.Int, user []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ao.contract.WatchLogs(opts, "UserClaimed", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoUserClaimed)
				if err := _Ao.contract.UnpackLog(event, "UserClaimed", log); err != nil {
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

// ParseUserClaimed is a log parse operation binding the contract event 0xace6f3f8956413e2875b9070e2616d13687dfb251cf63b343028c32822dfa263.
//
// Solidity: event UserClaimed(uint256 indexed poolId, address indexed user, address receiver, uint256 amount)
func (_Ao *AoFilterer) ParseUserClaimed(log types.Log) (*AoUserClaimed, error) {
	event := new(AoUserClaimed)
	if err := _Ao.contract.UnpackLog(event, "UserClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoUserStakedIterator is returned from FilterUserStaked and is used to iterate over the raw logs and unpacked data for UserStaked events raised by the Ao contract.
type AoUserStakedIterator struct {
	Event *AoUserStaked // Event containing the contract specifics and raw log

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
func (it *AoUserStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoUserStaked)
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
		it.Event = new(AoUserStaked)
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
func (it *AoUserStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoUserStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoUserStaked represents a UserStaked event raised by the Ao contract.
type AoUserStaked struct {
	PoolId         *big.Int
	User           common.Address
	Amount         *big.Int
	ArweaveAddress [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUserStaked is a free log retrieval operation binding the contract event 0x9612520d2ead6ce7570316d46e8b55381a1a66b9d591b1a2e770a36ba276a388.
//
// Solidity: event UserStaked(uint256 indexed poolId, address indexed user, uint256 amount, bytes32 arweaveAddress)
func (_Ao *AoFilterer) FilterUserStaked(opts *bind.FilterOpts, poolId []*big.Int, user []common.Address) (*AoUserStakedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ao.contract.FilterLogs(opts, "UserStaked", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AoUserStakedIterator{contract: _Ao.contract, event: "UserStaked", logs: logs, sub: sub}, nil
}

// WatchUserStaked is a free log subscription operation binding the contract event 0x9612520d2ead6ce7570316d46e8b55381a1a66b9d591b1a2e770a36ba276a388.
//
// Solidity: event UserStaked(uint256 indexed poolId, address indexed user, uint256 amount, bytes32 arweaveAddress)
func (_Ao *AoFilterer) WatchUserStaked(opts *bind.WatchOpts, sink chan<- *AoUserStaked, poolId []*big.Int, user []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ao.contract.WatchLogs(opts, "UserStaked", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoUserStaked)
				if err := _Ao.contract.UnpackLog(event, "UserStaked", log); err != nil {
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

// ParseUserStaked is a log parse operation binding the contract event 0x9612520d2ead6ce7570316d46e8b55381a1a66b9d591b1a2e770a36ba276a388.
//
// Solidity: event UserStaked(uint256 indexed poolId, address indexed user, uint256 amount, bytes32 arweaveAddress)
func (_Ao *AoFilterer) ParseUserStaked(log types.Log) (*AoUserStaked, error) {
	event := new(AoUserStaked)
	if err := _Ao.contract.UnpackLog(event, "UserStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AoUserWithdrawnIterator is returned from FilterUserWithdrawn and is used to iterate over the raw logs and unpacked data for UserWithdrawn events raised by the Ao contract.
type AoUserWithdrawnIterator struct {
	Event *AoUserWithdrawn // Event containing the contract specifics and raw log

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
func (it *AoUserWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AoUserWithdrawn)
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
		it.Event = new(AoUserWithdrawn)
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
func (it *AoUserWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AoUserWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AoUserWithdrawn represents a UserWithdrawn event raised by the Ao contract.
type AoUserWithdrawn struct {
	PoolId         *big.Int
	User           common.Address
	Amount         *big.Int
	ArweaveAddress [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUserWithdrawn is a free log retrieval operation binding the contract event 0xa746a3d46437245ec92a03d65f8e29d475586c621a9c23e71fffbdbecdcf07ca.
//
// Solidity: event UserWithdrawn(uint256 indexed poolId, address indexed user, uint256 amount, bytes32 arweaveAddress)
func (_Ao *AoFilterer) FilterUserWithdrawn(opts *bind.FilterOpts, poolId []*big.Int, user []common.Address) (*AoUserWithdrawnIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ao.contract.FilterLogs(opts, "UserWithdrawn", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AoUserWithdrawnIterator{contract: _Ao.contract, event: "UserWithdrawn", logs: logs, sub: sub}, nil
}

// WatchUserWithdrawn is a free log subscription operation binding the contract event 0xa746a3d46437245ec92a03d65f8e29d475586c621a9c23e71fffbdbecdcf07ca.
//
// Solidity: event UserWithdrawn(uint256 indexed poolId, address indexed user, uint256 amount, bytes32 arweaveAddress)
func (_Ao *AoFilterer) WatchUserWithdrawn(opts *bind.WatchOpts, sink chan<- *AoUserWithdrawn, poolId []*big.Int, user []common.Address) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ao.contract.WatchLogs(opts, "UserWithdrawn", poolIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AoUserWithdrawn)
				if err := _Ao.contract.UnpackLog(event, "UserWithdrawn", log); err != nil {
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

// ParseUserWithdrawn is a log parse operation binding the contract event 0xa746a3d46437245ec92a03d65f8e29d475586c621a9c23e71fffbdbecdcf07ca.
//
// Solidity: event UserWithdrawn(uint256 indexed poolId, address indexed user, uint256 amount, bytes32 arweaveAddress)
func (_Ao *AoFilterer) ParseUserWithdrawn(log types.Log) (*AoUserWithdrawn, error) {
	event := new(AoUserWithdrawn)
	if err := _Ao.contract.UnpackLog(event, "UserWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
