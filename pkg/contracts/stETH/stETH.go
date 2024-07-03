// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stETH

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

// StETHMetaData contains all meta data concerning the StETH contract.
var StETHMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CONTROL_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStakingPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"_stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"setStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lidoLocator\",\"type\":\"address\"},{\"name\":\"_eip712StETH\",\"type\":\"address\"}],\"name\":\"finalizeUpgrade_v2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDepositedValidators\",\"type\":\"uint256\"}],\"name\":\"unsafeChangeDepositedValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBufferedEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lidoLocator\",\"type\":\"address\"},{\"name\":\"_eip712StETH\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveELRewards\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentStakeLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeLimitFullInfo\",\"outputs\":[{\"name\":\"isStakingPaused\",\"type\":\"bool\"},{\"name\":\"isStakingLimitSet\",\"type\":\"bool\"},{\"name\":\"currentStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimitGrowthBlocks\",\"type\":\"uint256\"},{\"name\":\"prevStakeLimit\",\"type\":\"uint256\"},{\"name\":\"prevStakeBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferSharesFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resumeStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeDistribution\",\"outputs\":[{\"name\":\"treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveWithdrawals\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEIP712StETH\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxDepositsCount\",\"type\":\"uint256\"},{\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"name\":\"_depositCalldata\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBeaconStat\",\"outputs\":[{\"name\":\"depositedValidators\",\"type\":\"uint256\"},{\"name\":\"beaconValidators\",\"type\":\"uint256\"},{\"name\":\"beaconBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reportTimestamp\",\"type\":\"uint256\"},{\"name\":\"_timeElapsed\",\"type\":\"uint256\"},{\"name\":\"_clValidators\",\"type\":\"uint256\"},{\"name\":\"_clBalance\",\"type\":\"uint256\"},{\"name\":\"_withdrawalVaultBalance\",\"type\":\"uint256\"},{\"name\":\"_elRewardsVaultBalance\",\"type\":\"uint256\"},{\"name\":\"_sharesRequestedToBurn\",\"type\":\"uint256\"},{\"name\":\"_withdrawalFinalizationBatches\",\"type\":\"uint256[]\"},{\"name\":\"_simulatedShareRate\",\"type\":\"uint256\"}],\"name\":\"handleOracleReport\",\"outputs\":[{\"name\":\"postRebaseAmounts\",\"type\":\"uint256[4]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"name\":\"totalFee\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLidoLocator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositableEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalELRewardsCollected\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"StakingLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingLimitRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preCLValidators\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postCLValidators\",\"type\":\"uint256\"}],\"name\":\"CLValidatorsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositedValidators\",\"type\":\"uint256\"}],\"name\":\"DepositedValidatorsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preCLBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postCLBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"executionLayerRewardsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postBufferedEther\",\"type\":\"uint256\"}],\"name\":\"ETHDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeElapsed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preTotalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preTotalEther\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postTotalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postTotalEther\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesMintedAsFees\",\"type\":\"uint256\"}],\"name\":\"TokenRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lidoLocator\",\"type\":\"address\"}],\"name\":\"LidoLocatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unbuffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"eip712StETH\",\"type\":\"address\"}],\"name\":\"EIP712StETHInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"TransferShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"preRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesAmount\",\"type\":\"uint256\"}],\"name\":\"SharesBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"}]",
}

// StETHABI is the input ABI used to generate the binding from.
// Deprecated: Use StETHMetaData.ABI instead.
var StETHABI = StETHMetaData.ABI

// StETH is an auto generated Go binding around an Ethereum contract.
type StETH struct {
	StETHCaller     // Read-only binding to the contract
	StETHTransactor // Write-only binding to the contract
	StETHFilterer   // Log filterer for contract events
}

// StETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type StETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StETHSession struct {
	Contract     *StETH            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StETHCallerSession struct {
	Contract *StETHCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StETHTransactorSession struct {
	Contract     *StETHTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type StETHRaw struct {
	Contract *StETH // Generic contract binding to access the raw methods on
}

// StETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StETHCallerRaw struct {
	Contract *StETHCaller // Generic read-only contract binding to access the raw methods on
}

// StETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StETHTransactorRaw struct {
	Contract *StETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStETH creates a new instance of StETH, bound to a specific deployed contract.
func NewStETH(address common.Address, backend bind.ContractBackend) (*StETH, error) {
	contract, err := bindStETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StETH{StETHCaller: StETHCaller{contract: contract}, StETHTransactor: StETHTransactor{contract: contract}, StETHFilterer: StETHFilterer{contract: contract}}, nil
}

// NewStETHCaller creates a new read-only instance of StETH, bound to a specific deployed contract.
func NewStETHCaller(address common.Address, caller bind.ContractCaller) (*StETHCaller, error) {
	contract, err := bindStETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StETHCaller{contract: contract}, nil
}

// NewStETHTransactor creates a new write-only instance of StETH, bound to a specific deployed contract.
func NewStETHTransactor(address common.Address, transactor bind.ContractTransactor) (*StETHTransactor, error) {
	contract, err := bindStETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StETHTransactor{contract: contract}, nil
}

// NewStETHFilterer creates a new log filterer instance of StETH, bound to a specific deployed contract.
func NewStETHFilterer(address common.Address, filterer bind.ContractFilterer) (*StETHFilterer, error) {
	contract, err := bindStETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StETHFilterer{contract: contract}, nil
}

// bindStETH binds a generic wrapper to an already deployed contract.
func bindStETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StETH *StETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StETH.Contract.StETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StETH *StETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StETH.Contract.StETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StETH *StETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StETH.Contract.StETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StETH *StETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StETH *StETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StETH *StETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StETH.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StETH *StETHCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StETH *StETHSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StETH.Contract.DOMAINSEPARATOR(&_StETH.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StETH *StETHCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StETH.Contract.DOMAINSEPARATOR(&_StETH.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StETH *StETHCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StETH *StETHSession) PAUSEROLE() ([32]byte, error) {
	return _StETH.Contract.PAUSEROLE(&_StETH.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StETH *StETHCallerSession) PAUSEROLE() ([32]byte, error) {
	return _StETH.Contract.PAUSEROLE(&_StETH.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_StETH *StETHCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_StETH *StETHSession) RESUMEROLE() ([32]byte, error) {
	return _StETH.Contract.RESUMEROLE(&_StETH.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_StETH *StETHCallerSession) RESUMEROLE() ([32]byte, error) {
	return _StETH.Contract.RESUMEROLE(&_StETH.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_StETH *StETHCaller) STAKINGCONTROLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "STAKING_CONTROL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_StETH *StETHSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _StETH.Contract.STAKINGCONTROLROLE(&_StETH.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_StETH *StETHCallerSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _StETH.Contract.STAKINGCONTROLROLE(&_StETH.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_StETH *StETHCaller) STAKINGPAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "STAKING_PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_StETH *StETHSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _StETH.Contract.STAKINGPAUSEROLE(&_StETH.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_StETH *StETHCallerSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _StETH.Contract.STAKINGPAUSEROLE(&_StETH.CallOpts)
}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_StETH *StETHCaller) UNSAFECHANGEDEPOSITEDVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_StETH *StETHSession) UNSAFECHANGEDEPOSITEDVALIDATORSROLE() ([32]byte, error) {
	return _StETH.Contract.UNSAFECHANGEDEPOSITEDVALIDATORSROLE(&_StETH.CallOpts)
}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_StETH *StETHCallerSession) UNSAFECHANGEDEPOSITEDVALIDATORSROLE() ([32]byte, error) {
	return _StETH.Contract.UNSAFECHANGEDEPOSITEDVALIDATORSROLE(&_StETH.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_StETH *StETHCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_StETH *StETHSession) AllowRecoverability(token common.Address) (bool, error) {
	return _StETH.Contract.AllowRecoverability(&_StETH.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_StETH *StETHCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _StETH.Contract.AllowRecoverability(&_StETH.CallOpts, token)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StETH *StETHCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StETH *StETHSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StETH.Contract.Allowance(&_StETH.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StETH *StETHCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StETH.Contract.Allowance(&_StETH.CallOpts, _owner, _spender)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_StETH *StETHCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_StETH *StETHSession) AppId() ([32]byte, error) {
	return _StETH.Contract.AppId(&_StETH.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_StETH *StETHCallerSession) AppId() ([32]byte, error) {
	return _StETH.Contract.AppId(&_StETH.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_StETH *StETHCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_StETH *StETHSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _StETH.Contract.BalanceOf(&_StETH.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_StETH *StETHCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _StETH.Contract.BalanceOf(&_StETH.CallOpts, _account)
}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_StETH *StETHCaller) CanDeposit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "canDeposit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_StETH *StETHSession) CanDeposit() (bool, error) {
	return _StETH.Contract.CanDeposit(&_StETH.CallOpts)
}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_StETH *StETHCallerSession) CanDeposit() (bool, error) {
	return _StETH.Contract.CanDeposit(&_StETH.CallOpts)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_StETH *StETHCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_StETH *StETHSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _StETH.Contract.CanPerform(&_StETH.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_StETH *StETHCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _StETH.Contract.CanPerform(&_StETH.CallOpts, _sender, _role, _params)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StETH *StETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StETH *StETHSession) Decimals() (uint8, error) {
	return _StETH.Contract.Decimals(&_StETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StETH *StETHCallerSession) Decimals() (uint8, error) {
	return _StETH.Contract.Decimals(&_StETH.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_StETH *StETHCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_StETH *StETHSession) Eip712Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _StETH.Contract.Eip712Domain(&_StETH.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_StETH *StETHCallerSession) Eip712Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _StETH.Contract.Eip712Domain(&_StETH.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_StETH *StETHCaller) GetBeaconStat(opts *bind.CallOpts) (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getBeaconStat")

	outstruct := new(struct {
		DepositedValidators *big.Int
		BeaconValidators    *big.Int
		BeaconBalance       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositedValidators = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BeaconValidators = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BeaconBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_StETH *StETHSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _StETH.Contract.GetBeaconStat(&_StETH.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_StETH *StETHCallerSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _StETH.Contract.GetBeaconStat(&_StETH.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_StETH *StETHCaller) GetBufferedEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getBufferedEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_StETH *StETHSession) GetBufferedEther() (*big.Int, error) {
	return _StETH.Contract.GetBufferedEther(&_StETH.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_StETH *StETHCallerSession) GetBufferedEther() (*big.Int, error) {
	return _StETH.Contract.GetBufferedEther(&_StETH.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_StETH *StETHCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_StETH *StETHSession) GetContractVersion() (*big.Int, error) {
	return _StETH.Contract.GetContractVersion(&_StETH.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_StETH *StETHCallerSession) GetContractVersion() (*big.Int, error) {
	return _StETH.Contract.GetContractVersion(&_StETH.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_StETH *StETHCaller) GetCurrentStakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getCurrentStakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_StETH *StETHSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _StETH.Contract.GetCurrentStakeLimit(&_StETH.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_StETH *StETHCallerSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _StETH.Contract.GetCurrentStakeLimit(&_StETH.CallOpts)
}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_StETH *StETHCaller) GetDepositableEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getDepositableEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_StETH *StETHSession) GetDepositableEther() (*big.Int, error) {
	return _StETH.Contract.GetDepositableEther(&_StETH.CallOpts)
}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_StETH *StETHCallerSession) GetDepositableEther() (*big.Int, error) {
	return _StETH.Contract.GetDepositableEther(&_StETH.CallOpts)
}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_StETH *StETHCaller) GetEIP712StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getEIP712StETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_StETH *StETHSession) GetEIP712StETH() (common.Address, error) {
	return _StETH.Contract.GetEIP712StETH(&_StETH.CallOpts)
}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_StETH *StETHCallerSession) GetEIP712StETH() (common.Address, error) {
	return _StETH.Contract.GetEIP712StETH(&_StETH.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_StETH *StETHCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_StETH *StETHSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _StETH.Contract.GetEVMScriptExecutor(&_StETH.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_StETH *StETHCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _StETH.Contract.GetEVMScriptExecutor(&_StETH.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_StETH *StETHCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_StETH *StETHSession) GetEVMScriptRegistry() (common.Address, error) {
	return _StETH.Contract.GetEVMScriptRegistry(&_StETH.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_StETH *StETHCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _StETH.Contract.GetEVMScriptRegistry(&_StETH.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_StETH *StETHCaller) GetFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_StETH *StETHSession) GetFee() (uint16, error) {
	return _StETH.Contract.GetFee(&_StETH.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_StETH *StETHCallerSession) GetFee() (uint16, error) {
	return _StETH.Contract.GetFee(&_StETH.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_StETH *StETHCaller) GetFeeDistribution(opts *bind.CallOpts) (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getFeeDistribution")

	outstruct := new(struct {
		TreasuryFeeBasisPoints  uint16
		InsuranceFeeBasisPoints uint16
		OperatorsFeeBasisPoints uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TreasuryFeeBasisPoints = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.InsuranceFeeBasisPoints = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.OperatorsFeeBasisPoints = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_StETH *StETHSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _StETH.Contract.GetFeeDistribution(&_StETH.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_StETH *StETHCallerSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _StETH.Contract.GetFeeDistribution(&_StETH.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_StETH *StETHCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_StETH *StETHSession) GetInitializationBlock() (*big.Int, error) {
	return _StETH.Contract.GetInitializationBlock(&_StETH.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_StETH *StETHCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _StETH.Contract.GetInitializationBlock(&_StETH.CallOpts)
}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_StETH *StETHCaller) GetLidoLocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getLidoLocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_StETH *StETHSession) GetLidoLocator() (common.Address, error) {
	return _StETH.Contract.GetLidoLocator(&_StETH.CallOpts)
}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_StETH *StETHCallerSession) GetLidoLocator() (common.Address, error) {
	return _StETH.Contract.GetLidoLocator(&_StETH.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_StETH *StETHCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_StETH *StETHSession) GetOracle() (common.Address, error) {
	return _StETH.Contract.GetOracle(&_StETH.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_StETH *StETHCallerSession) GetOracle() (common.Address, error) {
	return _StETH.Contract.GetOracle(&_StETH.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_StETH *StETHCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_StETH *StETHSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _StETH.Contract.GetPooledEthByShares(&_StETH.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_StETH *StETHCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _StETH.Contract.GetPooledEthByShares(&_StETH.CallOpts, _sharesAmount)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_StETH *StETHCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_StETH *StETHSession) GetRecoveryVault() (common.Address, error) {
	return _StETH.Contract.GetRecoveryVault(&_StETH.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_StETH *StETHCallerSession) GetRecoveryVault() (common.Address, error) {
	return _StETH.Contract.GetRecoveryVault(&_StETH.CallOpts)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_StETH *StETHCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_StETH *StETHSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _StETH.Contract.GetSharesByPooledEth(&_StETH.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_StETH *StETHCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _StETH.Contract.GetSharesByPooledEth(&_StETH.CallOpts, _ethAmount)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_StETH *StETHCaller) GetStakeLimitFullInfo(opts *bind.CallOpts) (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getStakeLimitFullInfo")

	outstruct := new(struct {
		IsStakingPaused           bool
		IsStakingLimitSet         bool
		CurrentStakeLimit         *big.Int
		MaxStakeLimit             *big.Int
		MaxStakeLimitGrowthBlocks *big.Int
		PrevStakeLimit            *big.Int
		PrevStakeBlockNumber      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStakingPaused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsStakingLimitSet = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.CurrentStakeLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimitGrowthBlocks = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeLimit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeBlockNumber = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_StETH *StETHSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _StETH.Contract.GetStakeLimitFullInfo(&_StETH.CallOpts)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_StETH *StETHCallerSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _StETH.Contract.GetStakeLimitFullInfo(&_StETH.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_StETH *StETHCaller) GetTotalELRewardsCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getTotalELRewardsCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_StETH *StETHSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _StETH.Contract.GetTotalELRewardsCollected(&_StETH.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_StETH *StETHCallerSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _StETH.Contract.GetTotalELRewardsCollected(&_StETH.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_StETH *StETHCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_StETH *StETHSession) GetTotalPooledEther() (*big.Int, error) {
	return _StETH.Contract.GetTotalPooledEther(&_StETH.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_StETH *StETHCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _StETH.Contract.GetTotalPooledEther(&_StETH.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_StETH *StETHCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_StETH *StETHSession) GetTotalShares() (*big.Int, error) {
	return _StETH.Contract.GetTotalShares(&_StETH.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_StETH *StETHCallerSession) GetTotalShares() (*big.Int, error) {
	return _StETH.Contract.GetTotalShares(&_StETH.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_StETH *StETHCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_StETH *StETHSession) GetTreasury() (common.Address, error) {
	return _StETH.Contract.GetTreasury(&_StETH.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_StETH *StETHCallerSession) GetTreasury() (common.Address, error) {
	return _StETH.Contract.GetTreasury(&_StETH.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_StETH *StETHCaller) GetWithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "getWithdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_StETH *StETHSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _StETH.Contract.GetWithdrawalCredentials(&_StETH.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_StETH *StETHCallerSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _StETH.Contract.GetWithdrawalCredentials(&_StETH.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_StETH *StETHCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_StETH *StETHSession) HasInitialized() (bool, error) {
	return _StETH.Contract.HasInitialized(&_StETH.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_StETH *StETHCallerSession) HasInitialized() (bool, error) {
	return _StETH.Contract.HasInitialized(&_StETH.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_StETH *StETHCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_StETH *StETHSession) IsPetrified() (bool, error) {
	return _StETH.Contract.IsPetrified(&_StETH.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_StETH *StETHCallerSession) IsPetrified() (bool, error) {
	return _StETH.Contract.IsPetrified(&_StETH.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_StETH *StETHCaller) IsStakingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "isStakingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_StETH *StETHSession) IsStakingPaused() (bool, error) {
	return _StETH.Contract.IsStakingPaused(&_StETH.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_StETH *StETHCallerSession) IsStakingPaused() (bool, error) {
	return _StETH.Contract.IsStakingPaused(&_StETH.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_StETH *StETHCaller) IsStopped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "isStopped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_StETH *StETHSession) IsStopped() (bool, error) {
	return _StETH.Contract.IsStopped(&_StETH.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_StETH *StETHCallerSession) IsStopped() (bool, error) {
	return _StETH.Contract.IsStopped(&_StETH.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_StETH *StETHCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_StETH *StETHSession) Kernel() (common.Address, error) {
	return _StETH.Contract.Kernel(&_StETH.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_StETH *StETHCallerSession) Kernel() (common.Address, error) {
	return _StETH.Contract.Kernel(&_StETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StETH *StETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StETH *StETHSession) Name() (string, error) {
	return _StETH.Contract.Name(&_StETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StETH *StETHCallerSession) Name() (string, error) {
	return _StETH.Contract.Name(&_StETH.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_StETH *StETHCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_StETH *StETHSession) Nonces(owner common.Address) (*big.Int, error) {
	return _StETH.Contract.Nonces(&_StETH.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_StETH *StETHCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _StETH.Contract.Nonces(&_StETH.CallOpts, owner)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_StETH *StETHCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_StETH *StETHSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _StETH.Contract.SharesOf(&_StETH.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_StETH *StETHCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _StETH.Contract.SharesOf(&_StETH.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StETH *StETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StETH *StETHSession) Symbol() (string, error) {
	return _StETH.Contract.Symbol(&_StETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StETH *StETHCallerSession) Symbol() (string, error) {
	return _StETH.Contract.Symbol(&_StETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StETH *StETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StETH *StETHSession) TotalSupply() (*big.Int, error) {
	return _StETH.Contract.TotalSupply(&_StETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StETH *StETHCallerSession) TotalSupply() (*big.Int, error) {
	return _StETH.Contract.TotalSupply(&_StETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StETH *StETHTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StETH *StETHSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.Approve(&_StETH.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StETH *StETHTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.Approve(&_StETH.TransactOpts, _spender, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_StETH *StETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_StETH *StETHSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.DecreaseAllowance(&_StETH.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_StETH *StETHTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.DecreaseAllowance(&_StETH.TransactOpts, _spender, _subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_StETH *StETHTransactor) Deposit(opts *bind.TransactOpts, _maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "deposit", _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_StETH *StETHSession) Deposit(_maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _StETH.Contract.Deposit(&_StETH.TransactOpts, _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_StETH *StETHTransactorSession) Deposit(_maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _StETH.Contract.Deposit(&_StETH.TransactOpts, _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x2f85e57c.
//
// Solidity: function finalizeUpgrade_v2(address _lidoLocator, address _eip712StETH) returns()
func (_StETH *StETHTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts, _lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "finalizeUpgrade_v2", _lidoLocator, _eip712StETH)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x2f85e57c.
//
// Solidity: function finalizeUpgrade_v2(address _lidoLocator, address _eip712StETH) returns()
func (_StETH *StETHSession) FinalizeUpgradeV2(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StETH.Contract.FinalizeUpgradeV2(&_StETH.TransactOpts, _lidoLocator, _eip712StETH)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x2f85e57c.
//
// Solidity: function finalizeUpgrade_v2(address _lidoLocator, address _eip712StETH) returns()
func (_StETH *StETHTransactorSession) FinalizeUpgradeV2(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StETH.Contract.FinalizeUpgradeV2(&_StETH.TransactOpts, _lidoLocator, _eip712StETH)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0xbac3f3c5.
//
// Solidity: function handleOracleReport(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _clValidators, uint256 _clBalance, uint256 _withdrawalVaultBalance, uint256 _elRewardsVaultBalance, uint256 _sharesRequestedToBurn, uint256[] _withdrawalFinalizationBatches, uint256 _simulatedShareRate) returns(uint256[4] postRebaseAmounts)
func (_StETH *StETHTransactor) HandleOracleReport(opts *bind.TransactOpts, _reportTimestamp *big.Int, _timeElapsed *big.Int, _clValidators *big.Int, _clBalance *big.Int, _withdrawalVaultBalance *big.Int, _elRewardsVaultBalance *big.Int, _sharesRequestedToBurn *big.Int, _withdrawalFinalizationBatches []*big.Int, _simulatedShareRate *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "handleOracleReport", _reportTimestamp, _timeElapsed, _clValidators, _clBalance, _withdrawalVaultBalance, _elRewardsVaultBalance, _sharesRequestedToBurn, _withdrawalFinalizationBatches, _simulatedShareRate)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0xbac3f3c5.
//
// Solidity: function handleOracleReport(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _clValidators, uint256 _clBalance, uint256 _withdrawalVaultBalance, uint256 _elRewardsVaultBalance, uint256 _sharesRequestedToBurn, uint256[] _withdrawalFinalizationBatches, uint256 _simulatedShareRate) returns(uint256[4] postRebaseAmounts)
func (_StETH *StETHSession) HandleOracleReport(_reportTimestamp *big.Int, _timeElapsed *big.Int, _clValidators *big.Int, _clBalance *big.Int, _withdrawalVaultBalance *big.Int, _elRewardsVaultBalance *big.Int, _sharesRequestedToBurn *big.Int, _withdrawalFinalizationBatches []*big.Int, _simulatedShareRate *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.HandleOracleReport(&_StETH.TransactOpts, _reportTimestamp, _timeElapsed, _clValidators, _clBalance, _withdrawalVaultBalance, _elRewardsVaultBalance, _sharesRequestedToBurn, _withdrawalFinalizationBatches, _simulatedShareRate)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0xbac3f3c5.
//
// Solidity: function handleOracleReport(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _clValidators, uint256 _clBalance, uint256 _withdrawalVaultBalance, uint256 _elRewardsVaultBalance, uint256 _sharesRequestedToBurn, uint256[] _withdrawalFinalizationBatches, uint256 _simulatedShareRate) returns(uint256[4] postRebaseAmounts)
func (_StETH *StETHTransactorSession) HandleOracleReport(_reportTimestamp *big.Int, _timeElapsed *big.Int, _clValidators *big.Int, _clBalance *big.Int, _withdrawalVaultBalance *big.Int, _elRewardsVaultBalance *big.Int, _sharesRequestedToBurn *big.Int, _withdrawalFinalizationBatches []*big.Int, _simulatedShareRate *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.HandleOracleReport(&_StETH.TransactOpts, _reportTimestamp, _timeElapsed, _clValidators, _clBalance, _withdrawalVaultBalance, _elRewardsVaultBalance, _sharesRequestedToBurn, _withdrawalFinalizationBatches, _simulatedShareRate)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_StETH *StETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_StETH *StETHSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.IncreaseAllowance(&_StETH.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_StETH *StETHTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.IncreaseAllowance(&_StETH.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_StETH *StETHTransactor) Initialize(opts *bind.TransactOpts, _lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "initialize", _lidoLocator, _eip712StETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_StETH *StETHSession) Initialize(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StETH.Contract.Initialize(&_StETH.TransactOpts, _lidoLocator, _eip712StETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_StETH *StETHTransactorSession) Initialize(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StETH.Contract.Initialize(&_StETH.TransactOpts, _lidoLocator, _eip712StETH)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_StETH *StETHTransactor) PauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "pauseStaking")
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_StETH *StETHSession) PauseStaking() (*types.Transaction, error) {
	return _StETH.Contract.PauseStaking(&_StETH.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_StETH *StETHTransactorSession) PauseStaking() (*types.Transaction, error) {
	return _StETH.Contract.PauseStaking(&_StETH.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_StETH *StETHTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_StETH *StETHSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StETH.Contract.Permit(&_StETH.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_StETH *StETHTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StETH.Contract.Permit(&_StETH.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_StETH *StETHTransactor) ReceiveELRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "receiveELRewards")
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_StETH *StETHSession) ReceiveELRewards() (*types.Transaction, error) {
	return _StETH.Contract.ReceiveELRewards(&_StETH.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_StETH *StETHTransactorSession) ReceiveELRewards() (*types.Transaction, error) {
	return _StETH.Contract.ReceiveELRewards(&_StETH.TransactOpts)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_StETH *StETHTransactor) ReceiveWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "receiveWithdrawals")
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_StETH *StETHSession) ReceiveWithdrawals() (*types.Transaction, error) {
	return _StETH.Contract.ReceiveWithdrawals(&_StETH.TransactOpts)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_StETH *StETHTransactorSession) ReceiveWithdrawals() (*types.Transaction, error) {
	return _StETH.Contract.ReceiveWithdrawals(&_StETH.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_StETH *StETHTransactor) RemoveStakingLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "removeStakingLimit")
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_StETH *StETHSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _StETH.Contract.RemoveStakingLimit(&_StETH.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_StETH *StETHTransactorSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _StETH.Contract.RemoveStakingLimit(&_StETH.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_StETH *StETHTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_StETH *StETHSession) Resume() (*types.Transaction, error) {
	return _StETH.Contract.Resume(&_StETH.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_StETH *StETHTransactorSession) Resume() (*types.Transaction, error) {
	return _StETH.Contract.Resume(&_StETH.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_StETH *StETHTransactor) ResumeStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "resumeStaking")
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_StETH *StETHSession) ResumeStaking() (*types.Transaction, error) {
	return _StETH.Contract.ResumeStaking(&_StETH.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_StETH *StETHTransactorSession) ResumeStaking() (*types.Transaction, error) {
	return _StETH.Contract.ResumeStaking(&_StETH.TransactOpts)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_StETH *StETHTransactor) SetStakingLimit(opts *bind.TransactOpts, _maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "setStakingLimit", _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_StETH *StETHSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.SetStakingLimit(&_StETH.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_StETH *StETHTransactorSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.SetStakingLimit(&_StETH.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_StETH *StETHTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_StETH *StETHSession) Stop() (*types.Transaction, error) {
	return _StETH.Contract.Stop(&_StETH.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_StETH *StETHTransactorSession) Stop() (*types.Transaction, error) {
	return _StETH.Contract.Stop(&_StETH.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_StETH *StETHTransactor) Submit(opts *bind.TransactOpts, _referral common.Address) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "submit", _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_StETH *StETHSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _StETH.Contract.Submit(&_StETH.TransactOpts, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_StETH *StETHTransactorSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _StETH.Contract.Submit(&_StETH.TransactOpts, _referral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_StETH *StETHTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_StETH *StETHSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.Transfer(&_StETH.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_StETH *StETHTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.Transfer(&_StETH.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_StETH *StETHTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_StETH *StETHSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.TransferFrom(&_StETH.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_StETH *StETHTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.TransferFrom(&_StETH.TransactOpts, _sender, _recipient, _amount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StETH *StETHTransactor) TransferShares(opts *bind.TransactOpts, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "transferShares", _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StETH *StETHSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.TransferShares(&_StETH.TransactOpts, _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StETH *StETHTransactorSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.TransferShares(&_StETH.TransactOpts, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StETH *StETHTransactor) TransferSharesFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "transferSharesFrom", _sender, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StETH *StETHSession) TransferSharesFrom(_sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.TransferSharesFrom(&_StETH.TransactOpts, _sender, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StETH *StETHTransactorSession) TransferSharesFrom(_sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.TransferSharesFrom(&_StETH.TransactOpts, _sender, _recipient, _sharesAmount)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_StETH *StETHTransactor) TransferToVault(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "transferToVault", arg0)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_StETH *StETHSession) TransferToVault(arg0 common.Address) (*types.Transaction, error) {
	return _StETH.Contract.TransferToVault(&_StETH.TransactOpts, arg0)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_StETH *StETHTransactorSession) TransferToVault(arg0 common.Address) (*types.Transaction, error) {
	return _StETH.Contract.TransferToVault(&_StETH.TransactOpts, arg0)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_StETH *StETHTransactor) UnsafeChangeDepositedValidators(opts *bind.TransactOpts, _newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _StETH.contract.Transact(opts, "unsafeChangeDepositedValidators", _newDepositedValidators)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_StETH *StETHSession) UnsafeChangeDepositedValidators(_newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.UnsafeChangeDepositedValidators(&_StETH.TransactOpts, _newDepositedValidators)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_StETH *StETHTransactorSession) UnsafeChangeDepositedValidators(_newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _StETH.Contract.UnsafeChangeDepositedValidators(&_StETH.TransactOpts, _newDepositedValidators)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StETH *StETHTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StETH.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StETH *StETHSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StETH.Contract.Fallback(&_StETH.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StETH *StETHTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StETH.Contract.Fallback(&_StETH.TransactOpts, calldata)
}

// StETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StETH contract.
type StETHApprovalIterator struct {
	Event *StETHApproval // Event containing the contract specifics and raw log

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
func (it *StETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHApproval)
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
		it.Event = new(StETHApproval)
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
func (it *StETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHApproval represents a Approval event raised by the StETH contract.
type StETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StETH *StETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StETHApprovalIterator{contract: _StETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StETH *StETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHApproval)
				if err := _StETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StETH *StETHFilterer) ParseApproval(log types.Log) (*StETHApproval, error) {
	event := new(StETHApproval)
	if err := _StETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHCLValidatorsUpdatedIterator is returned from FilterCLValidatorsUpdated and is used to iterate over the raw logs and unpacked data for CLValidatorsUpdated events raised by the StETH contract.
type StETHCLValidatorsUpdatedIterator struct {
	Event *StETHCLValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *StETHCLValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHCLValidatorsUpdated)
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
		it.Event = new(StETHCLValidatorsUpdated)
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
func (it *StETHCLValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHCLValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHCLValidatorsUpdated represents a CLValidatorsUpdated event raised by the StETH contract.
type StETHCLValidatorsUpdated struct {
	ReportTimestamp  *big.Int
	PreCLValidators  *big.Int
	PostCLValidators *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCLValidatorsUpdated is a free log retrieval operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_StETH *StETHFilterer) FilterCLValidatorsUpdated(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*StETHCLValidatorsUpdatedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "CLValidatorsUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &StETHCLValidatorsUpdatedIterator{contract: _StETH.contract, event: "CLValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchCLValidatorsUpdated is a free log subscription operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_StETH *StETHFilterer) WatchCLValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *StETHCLValidatorsUpdated, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "CLValidatorsUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHCLValidatorsUpdated)
				if err := _StETH.contract.UnpackLog(event, "CLValidatorsUpdated", log); err != nil {
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

// ParseCLValidatorsUpdated is a log parse operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_StETH *StETHFilterer) ParseCLValidatorsUpdated(log types.Log) (*StETHCLValidatorsUpdated, error) {
	event := new(StETHCLValidatorsUpdated)
	if err := _StETH.contract.UnpackLog(event, "CLValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the StETH contract.
type StETHContractVersionSetIterator struct {
	Event *StETHContractVersionSet // Event containing the contract specifics and raw log

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
func (it *StETHContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHContractVersionSet)
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
		it.Event = new(StETHContractVersionSet)
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
func (it *StETHContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHContractVersionSet represents a ContractVersionSet event raised by the StETH contract.
type StETHContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_StETH *StETHFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*StETHContractVersionSetIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &StETHContractVersionSetIterator{contract: _StETH.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_StETH *StETHFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *StETHContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHContractVersionSet)
				if err := _StETH.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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

// ParseContractVersionSet is a log parse operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_StETH *StETHFilterer) ParseContractVersionSet(log types.Log) (*StETHContractVersionSet, error) {
	event := new(StETHContractVersionSet)
	if err := _StETH.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHDepositedValidatorsChangedIterator is returned from FilterDepositedValidatorsChanged and is used to iterate over the raw logs and unpacked data for DepositedValidatorsChanged events raised by the StETH contract.
type StETHDepositedValidatorsChangedIterator struct {
	Event *StETHDepositedValidatorsChanged // Event containing the contract specifics and raw log

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
func (it *StETHDepositedValidatorsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHDepositedValidatorsChanged)
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
		it.Event = new(StETHDepositedValidatorsChanged)
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
func (it *StETHDepositedValidatorsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHDepositedValidatorsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHDepositedValidatorsChanged represents a DepositedValidatorsChanged event raised by the StETH contract.
type StETHDepositedValidatorsChanged struct {
	DepositedValidators *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDepositedValidatorsChanged is a free log retrieval operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_StETH *StETHFilterer) FilterDepositedValidatorsChanged(opts *bind.FilterOpts) (*StETHDepositedValidatorsChangedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "DepositedValidatorsChanged")
	if err != nil {
		return nil, err
	}
	return &StETHDepositedValidatorsChangedIterator{contract: _StETH.contract, event: "DepositedValidatorsChanged", logs: logs, sub: sub}, nil
}

// WatchDepositedValidatorsChanged is a free log subscription operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_StETH *StETHFilterer) WatchDepositedValidatorsChanged(opts *bind.WatchOpts, sink chan<- *StETHDepositedValidatorsChanged) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "DepositedValidatorsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHDepositedValidatorsChanged)
				if err := _StETH.contract.UnpackLog(event, "DepositedValidatorsChanged", log); err != nil {
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

// ParseDepositedValidatorsChanged is a log parse operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_StETH *StETHFilterer) ParseDepositedValidatorsChanged(log types.Log) (*StETHDepositedValidatorsChanged, error) {
	event := new(StETHDepositedValidatorsChanged)
	if err := _StETH.contract.UnpackLog(event, "DepositedValidatorsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHEIP712StETHInitializedIterator is returned from FilterEIP712StETHInitialized and is used to iterate over the raw logs and unpacked data for EIP712StETHInitialized events raised by the StETH contract.
type StETHEIP712StETHInitializedIterator struct {
	Event *StETHEIP712StETHInitialized // Event containing the contract specifics and raw log

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
func (it *StETHEIP712StETHInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHEIP712StETHInitialized)
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
		it.Event = new(StETHEIP712StETHInitialized)
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
func (it *StETHEIP712StETHInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHEIP712StETHInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHEIP712StETHInitialized represents a EIP712StETHInitialized event raised by the StETH contract.
type StETHEIP712StETHInitialized struct {
	Eip712StETH common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEIP712StETHInitialized is a free log retrieval operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_StETH *StETHFilterer) FilterEIP712StETHInitialized(opts *bind.FilterOpts) (*StETHEIP712StETHInitializedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "EIP712StETHInitialized")
	if err != nil {
		return nil, err
	}
	return &StETHEIP712StETHInitializedIterator{contract: _StETH.contract, event: "EIP712StETHInitialized", logs: logs, sub: sub}, nil
}

// WatchEIP712StETHInitialized is a free log subscription operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_StETH *StETHFilterer) WatchEIP712StETHInitialized(opts *bind.WatchOpts, sink chan<- *StETHEIP712StETHInitialized) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "EIP712StETHInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHEIP712StETHInitialized)
				if err := _StETH.contract.UnpackLog(event, "EIP712StETHInitialized", log); err != nil {
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

// ParseEIP712StETHInitialized is a log parse operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_StETH *StETHFilterer) ParseEIP712StETHInitialized(log types.Log) (*StETHEIP712StETHInitialized, error) {
	event := new(StETHEIP712StETHInitialized)
	if err := _StETH.contract.UnpackLog(event, "EIP712StETHInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHELRewardsReceivedIterator is returned from FilterELRewardsReceived and is used to iterate over the raw logs and unpacked data for ELRewardsReceived events raised by the StETH contract.
type StETHELRewardsReceivedIterator struct {
	Event *StETHELRewardsReceived // Event containing the contract specifics and raw log

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
func (it *StETHELRewardsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHELRewardsReceived)
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
		it.Event = new(StETHELRewardsReceived)
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
func (it *StETHELRewardsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHELRewardsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHELRewardsReceived represents a ELRewardsReceived event raised by the StETH contract.
type StETHELRewardsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterELRewardsReceived is a free log retrieval operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_StETH *StETHFilterer) FilterELRewardsReceived(opts *bind.FilterOpts) (*StETHELRewardsReceivedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return &StETHELRewardsReceivedIterator{contract: _StETH.contract, event: "ELRewardsReceived", logs: logs, sub: sub}, nil
}

// WatchELRewardsReceived is a free log subscription operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_StETH *StETHFilterer) WatchELRewardsReceived(opts *bind.WatchOpts, sink chan<- *StETHELRewardsReceived) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHELRewardsReceived)
				if err := _StETH.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
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

// ParseELRewardsReceived is a log parse operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_StETH *StETHFilterer) ParseELRewardsReceived(log types.Log) (*StETHELRewardsReceived, error) {
	event := new(StETHELRewardsReceived)
	if err := _StETH.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHETHDistributedIterator is returned from FilterETHDistributed and is used to iterate over the raw logs and unpacked data for ETHDistributed events raised by the StETH contract.
type StETHETHDistributedIterator struct {
	Event *StETHETHDistributed // Event containing the contract specifics and raw log

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
func (it *StETHETHDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHETHDistributed)
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
		it.Event = new(StETHETHDistributed)
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
func (it *StETHETHDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHETHDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHETHDistributed represents a ETHDistributed event raised by the StETH contract.
type StETHETHDistributed struct {
	ReportTimestamp                *big.Int
	PreCLBalance                   *big.Int
	PostCLBalance                  *big.Int
	WithdrawalsWithdrawn           *big.Int
	ExecutionLayerRewardsWithdrawn *big.Int
	PostBufferedEther              *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterETHDistributed is a free log retrieval operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_StETH *StETHFilterer) FilterETHDistributed(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*StETHETHDistributedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "ETHDistributed", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &StETHETHDistributedIterator{contract: _StETH.contract, event: "ETHDistributed", logs: logs, sub: sub}, nil
}

// WatchETHDistributed is a free log subscription operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_StETH *StETHFilterer) WatchETHDistributed(opts *bind.WatchOpts, sink chan<- *StETHETHDistributed, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "ETHDistributed", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHETHDistributed)
				if err := _StETH.contract.UnpackLog(event, "ETHDistributed", log); err != nil {
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

// ParseETHDistributed is a log parse operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_StETH *StETHFilterer) ParseETHDistributed(log types.Log) (*StETHETHDistributed, error) {
	event := new(StETHETHDistributed)
	if err := _StETH.contract.UnpackLog(event, "ETHDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHLidoLocatorSetIterator is returned from FilterLidoLocatorSet and is used to iterate over the raw logs and unpacked data for LidoLocatorSet events raised by the StETH contract.
type StETHLidoLocatorSetIterator struct {
	Event *StETHLidoLocatorSet // Event containing the contract specifics and raw log

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
func (it *StETHLidoLocatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHLidoLocatorSet)
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
		it.Event = new(StETHLidoLocatorSet)
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
func (it *StETHLidoLocatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHLidoLocatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHLidoLocatorSet represents a LidoLocatorSet event raised by the StETH contract.
type StETHLidoLocatorSet struct {
	LidoLocator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLidoLocatorSet is a free log retrieval operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_StETH *StETHFilterer) FilterLidoLocatorSet(opts *bind.FilterOpts) (*StETHLidoLocatorSetIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "LidoLocatorSet")
	if err != nil {
		return nil, err
	}
	return &StETHLidoLocatorSetIterator{contract: _StETH.contract, event: "LidoLocatorSet", logs: logs, sub: sub}, nil
}

// WatchLidoLocatorSet is a free log subscription operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_StETH *StETHFilterer) WatchLidoLocatorSet(opts *bind.WatchOpts, sink chan<- *StETHLidoLocatorSet) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "LidoLocatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHLidoLocatorSet)
				if err := _StETH.contract.UnpackLog(event, "LidoLocatorSet", log); err != nil {
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

// ParseLidoLocatorSet is a log parse operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_StETH *StETHFilterer) ParseLidoLocatorSet(log types.Log) (*StETHLidoLocatorSet, error) {
	event := new(StETHLidoLocatorSet)
	if err := _StETH.contract.UnpackLog(event, "LidoLocatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the StETH contract.
type StETHRecoverToVaultIterator struct {
	Event *StETHRecoverToVault // Event containing the contract specifics and raw log

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
func (it *StETHRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHRecoverToVault)
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
		it.Event = new(StETHRecoverToVault)
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
func (it *StETHRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHRecoverToVault represents a RecoverToVault event raised by the StETH contract.
type StETHRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_StETH *StETHFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*StETHRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StETHRecoverToVaultIterator{contract: _StETH.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_StETH *StETHFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *StETHRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHRecoverToVault)
				if err := _StETH.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_StETH *StETHFilterer) ParseRecoverToVault(log types.Log) (*StETHRecoverToVault, error) {
	event := new(StETHRecoverToVault)
	if err := _StETH.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the StETH contract.
type StETHResumedIterator struct {
	Event *StETHResumed // Event containing the contract specifics and raw log

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
func (it *StETHResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHResumed)
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
		it.Event = new(StETHResumed)
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
func (it *StETHResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHResumed represents a Resumed event raised by the StETH contract.
type StETHResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_StETH *StETHFilterer) FilterResumed(opts *bind.FilterOpts) (*StETHResumedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &StETHResumedIterator{contract: _StETH.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_StETH *StETHFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *StETHResumed) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHResumed)
				if err := _StETH.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_StETH *StETHFilterer) ParseResumed(log types.Log) (*StETHResumed, error) {
	event := new(StETHResumed)
	if err := _StETH.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the StETH contract.
type StETHScriptResultIterator struct {
	Event *StETHScriptResult // Event containing the contract specifics and raw log

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
func (it *StETHScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHScriptResult)
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
		it.Event = new(StETHScriptResult)
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
func (it *StETHScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHScriptResult represents a ScriptResult event raised by the StETH contract.
type StETHScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_StETH *StETHFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*StETHScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &StETHScriptResultIterator{contract: _StETH.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_StETH *StETHFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *StETHScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHScriptResult)
				if err := _StETH.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_StETH *StETHFilterer) ParseScriptResult(log types.Log) (*StETHScriptResult, error) {
	event := new(StETHScriptResult)
	if err := _StETH.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHSharesBurntIterator is returned from FilterSharesBurnt and is used to iterate over the raw logs and unpacked data for SharesBurnt events raised by the StETH contract.
type StETHSharesBurntIterator struct {
	Event *StETHSharesBurnt // Event containing the contract specifics and raw log

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
func (it *StETHSharesBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHSharesBurnt)
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
		it.Event = new(StETHSharesBurnt)
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
func (it *StETHSharesBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHSharesBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHSharesBurnt represents a SharesBurnt event raised by the StETH contract.
type StETHSharesBurnt struct {
	Account               common.Address
	PreRebaseTokenAmount  *big.Int
	PostRebaseTokenAmount *big.Int
	SharesAmount          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSharesBurnt is a free log retrieval operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_StETH *StETHFilterer) FilterSharesBurnt(opts *bind.FilterOpts, account []common.Address) (*StETHSharesBurntIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return &StETHSharesBurntIterator{contract: _StETH.contract, event: "SharesBurnt", logs: logs, sub: sub}, nil
}

// WatchSharesBurnt is a free log subscription operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_StETH *StETHFilterer) WatchSharesBurnt(opts *bind.WatchOpts, sink chan<- *StETHSharesBurnt, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHSharesBurnt)
				if err := _StETH.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
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

// ParseSharesBurnt is a log parse operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_StETH *StETHFilterer) ParseSharesBurnt(log types.Log) (*StETHSharesBurnt, error) {
	event := new(StETHSharesBurnt)
	if err := _StETH.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHStakingLimitRemovedIterator is returned from FilterStakingLimitRemoved and is used to iterate over the raw logs and unpacked data for StakingLimitRemoved events raised by the StETH contract.
type StETHStakingLimitRemovedIterator struct {
	Event *StETHStakingLimitRemoved // Event containing the contract specifics and raw log

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
func (it *StETHStakingLimitRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHStakingLimitRemoved)
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
		it.Event = new(StETHStakingLimitRemoved)
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
func (it *StETHStakingLimitRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHStakingLimitRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHStakingLimitRemoved represents a StakingLimitRemoved event raised by the StETH contract.
type StETHStakingLimitRemoved struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitRemoved is a free log retrieval operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_StETH *StETHFilterer) FilterStakingLimitRemoved(opts *bind.FilterOpts) (*StETHStakingLimitRemovedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return &StETHStakingLimitRemovedIterator{contract: _StETH.contract, event: "StakingLimitRemoved", logs: logs, sub: sub}, nil
}

// WatchStakingLimitRemoved is a free log subscription operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_StETH *StETHFilterer) WatchStakingLimitRemoved(opts *bind.WatchOpts, sink chan<- *StETHStakingLimitRemoved) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHStakingLimitRemoved)
				if err := _StETH.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
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

// ParseStakingLimitRemoved is a log parse operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_StETH *StETHFilterer) ParseStakingLimitRemoved(log types.Log) (*StETHStakingLimitRemoved, error) {
	event := new(StETHStakingLimitRemoved)
	if err := _StETH.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHStakingLimitSetIterator is returned from FilterStakingLimitSet and is used to iterate over the raw logs and unpacked data for StakingLimitSet events raised by the StETH contract.
type StETHStakingLimitSetIterator struct {
	Event *StETHStakingLimitSet // Event containing the contract specifics and raw log

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
func (it *StETHStakingLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHStakingLimitSet)
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
		it.Event = new(StETHStakingLimitSet)
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
func (it *StETHStakingLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHStakingLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHStakingLimitSet represents a StakingLimitSet event raised by the StETH contract.
type StETHStakingLimitSet struct {
	MaxStakeLimit              *big.Int
	StakeLimitIncreasePerBlock *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitSet is a free log retrieval operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_StETH *StETHFilterer) FilterStakingLimitSet(opts *bind.FilterOpts) (*StETHStakingLimitSetIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return &StETHStakingLimitSetIterator{contract: _StETH.contract, event: "StakingLimitSet", logs: logs, sub: sub}, nil
}

// WatchStakingLimitSet is a free log subscription operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_StETH *StETHFilterer) WatchStakingLimitSet(opts *bind.WatchOpts, sink chan<- *StETHStakingLimitSet) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHStakingLimitSet)
				if err := _StETH.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
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

// ParseStakingLimitSet is a log parse operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_StETH *StETHFilterer) ParseStakingLimitSet(log types.Log) (*StETHStakingLimitSet, error) {
	event := new(StETHStakingLimitSet)
	if err := _StETH.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHStakingPausedIterator is returned from FilterStakingPaused and is used to iterate over the raw logs and unpacked data for StakingPaused events raised by the StETH contract.
type StETHStakingPausedIterator struct {
	Event *StETHStakingPaused // Event containing the contract specifics and raw log

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
func (it *StETHStakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHStakingPaused)
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
		it.Event = new(StETHStakingPaused)
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
func (it *StETHStakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHStakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHStakingPaused represents a StakingPaused event raised by the StETH contract.
type StETHStakingPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingPaused is a free log retrieval operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_StETH *StETHFilterer) FilterStakingPaused(opts *bind.FilterOpts) (*StETHStakingPausedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return &StETHStakingPausedIterator{contract: _StETH.contract, event: "StakingPaused", logs: logs, sub: sub}, nil
}

// WatchStakingPaused is a free log subscription operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_StETH *StETHFilterer) WatchStakingPaused(opts *bind.WatchOpts, sink chan<- *StETHStakingPaused) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHStakingPaused)
				if err := _StETH.contract.UnpackLog(event, "StakingPaused", log); err != nil {
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

// ParseStakingPaused is a log parse operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_StETH *StETHFilterer) ParseStakingPaused(log types.Log) (*StETHStakingPaused, error) {
	event := new(StETHStakingPaused)
	if err := _StETH.contract.UnpackLog(event, "StakingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHStakingResumedIterator is returned from FilterStakingResumed and is used to iterate over the raw logs and unpacked data for StakingResumed events raised by the StETH contract.
type StETHStakingResumedIterator struct {
	Event *StETHStakingResumed // Event containing the contract specifics and raw log

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
func (it *StETHStakingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHStakingResumed)
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
		it.Event = new(StETHStakingResumed)
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
func (it *StETHStakingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHStakingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHStakingResumed represents a StakingResumed event raised by the StETH contract.
type StETHStakingResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingResumed is a free log retrieval operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_StETH *StETHFilterer) FilterStakingResumed(opts *bind.FilterOpts) (*StETHStakingResumedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return &StETHStakingResumedIterator{contract: _StETH.contract, event: "StakingResumed", logs: logs, sub: sub}, nil
}

// WatchStakingResumed is a free log subscription operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_StETH *StETHFilterer) WatchStakingResumed(opts *bind.WatchOpts, sink chan<- *StETHStakingResumed) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHStakingResumed)
				if err := _StETH.contract.UnpackLog(event, "StakingResumed", log); err != nil {
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

// ParseStakingResumed is a log parse operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_StETH *StETHFilterer) ParseStakingResumed(log types.Log) (*StETHStakingResumed, error) {
	event := new(StETHStakingResumed)
	if err := _StETH.contract.UnpackLog(event, "StakingResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHStoppedIterator is returned from FilterStopped and is used to iterate over the raw logs and unpacked data for Stopped events raised by the StETH contract.
type StETHStoppedIterator struct {
	Event *StETHStopped // Event containing the contract specifics and raw log

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
func (it *StETHStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHStopped)
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
		it.Event = new(StETHStopped)
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
func (it *StETHStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHStopped represents a Stopped event raised by the StETH contract.
type StETHStopped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopped is a free log retrieval operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_StETH *StETHFilterer) FilterStopped(opts *bind.FilterOpts) (*StETHStoppedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return &StETHStoppedIterator{contract: _StETH.contract, event: "Stopped", logs: logs, sub: sub}, nil
}

// WatchStopped is a free log subscription operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_StETH *StETHFilterer) WatchStopped(opts *bind.WatchOpts, sink chan<- *StETHStopped) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHStopped)
				if err := _StETH.contract.UnpackLog(event, "Stopped", log); err != nil {
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

// ParseStopped is a log parse operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_StETH *StETHFilterer) ParseStopped(log types.Log) (*StETHStopped, error) {
	event := new(StETHStopped)
	if err := _StETH.contract.UnpackLog(event, "Stopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the StETH contract.
type StETHSubmittedIterator struct {
	Event *StETHSubmitted // Event containing the contract specifics and raw log

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
func (it *StETHSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHSubmitted)
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
		it.Event = new(StETHSubmitted)
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
func (it *StETHSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHSubmitted represents a Submitted event raised by the StETH contract.
type StETHSubmitted struct {
	Sender   common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_StETH *StETHFilterer) FilterSubmitted(opts *bind.FilterOpts, sender []common.Address) (*StETHSubmittedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return &StETHSubmittedIterator{contract: _StETH.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_StETH *StETHFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *StETHSubmitted, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHSubmitted)
				if err := _StETH.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_StETH *StETHFilterer) ParseSubmitted(log types.Log) (*StETHSubmitted, error) {
	event := new(StETHSubmitted)
	if err := _StETH.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHTokenRebasedIterator is returned from FilterTokenRebased and is used to iterate over the raw logs and unpacked data for TokenRebased events raised by the StETH contract.
type StETHTokenRebasedIterator struct {
	Event *StETHTokenRebased // Event containing the contract specifics and raw log

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
func (it *StETHTokenRebasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHTokenRebased)
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
		it.Event = new(StETHTokenRebased)
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
func (it *StETHTokenRebasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHTokenRebasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHTokenRebased represents a TokenRebased event raised by the StETH contract.
type StETHTokenRebased struct {
	ReportTimestamp    *big.Int
	TimeElapsed        *big.Int
	PreTotalShares     *big.Int
	PreTotalEther      *big.Int
	PostTotalShares    *big.Int
	PostTotalEther     *big.Int
	SharesMintedAsFees *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenRebased is a free log retrieval operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_StETH *StETHFilterer) FilterTokenRebased(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*StETHTokenRebasedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "TokenRebased", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &StETHTokenRebasedIterator{contract: _StETH.contract, event: "TokenRebased", logs: logs, sub: sub}, nil
}

// WatchTokenRebased is a free log subscription operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_StETH *StETHFilterer) WatchTokenRebased(opts *bind.WatchOpts, sink chan<- *StETHTokenRebased, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "TokenRebased", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHTokenRebased)
				if err := _StETH.contract.UnpackLog(event, "TokenRebased", log); err != nil {
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

// ParseTokenRebased is a log parse operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_StETH *StETHFilterer) ParseTokenRebased(log types.Log) (*StETHTokenRebased, error) {
	event := new(StETHTokenRebased)
	if err := _StETH.contract.UnpackLog(event, "TokenRebased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StETH contract.
type StETHTransferIterator struct {
	Event *StETHTransfer // Event containing the contract specifics and raw log

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
func (it *StETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHTransfer)
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
		it.Event = new(StETHTransfer)
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
func (it *StETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHTransfer represents a Transfer event raised by the StETH contract.
type StETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StETH *StETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StETHTransferIterator{contract: _StETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StETH *StETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHTransfer)
				if err := _StETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StETH *StETHFilterer) ParseTransfer(log types.Log) (*StETHTransfer, error) {
	event := new(StETHTransfer)
	if err := _StETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHTransferSharesIterator is returned from FilterTransferShares and is used to iterate over the raw logs and unpacked data for TransferShares events raised by the StETH contract.
type StETHTransferSharesIterator struct {
	Event *StETHTransferShares // Event containing the contract specifics and raw log

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
func (it *StETHTransferSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHTransferShares)
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
		it.Event = new(StETHTransferShares)
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
func (it *StETHTransferSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHTransferSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHTransferShares represents a TransferShares event raised by the StETH contract.
type StETHTransferShares struct {
	From        common.Address
	To          common.Address
	SharesValue *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferShares is a free log retrieval operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_StETH *StETHFilterer) FilterTransferShares(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StETHTransferSharesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StETH.contract.FilterLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StETHTransferSharesIterator{contract: _StETH.contract, event: "TransferShares", logs: logs, sub: sub}, nil
}

// WatchTransferShares is a free log subscription operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_StETH *StETHFilterer) WatchTransferShares(opts *bind.WatchOpts, sink chan<- *StETHTransferShares, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StETH.contract.WatchLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHTransferShares)
				if err := _StETH.contract.UnpackLog(event, "TransferShares", log); err != nil {
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

// ParseTransferShares is a log parse operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_StETH *StETHFilterer) ParseTransferShares(log types.Log) (*StETHTransferShares, error) {
	event := new(StETHTransferShares)
	if err := _StETH.contract.UnpackLog(event, "TransferShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHUnbufferedIterator is returned from FilterUnbuffered and is used to iterate over the raw logs and unpacked data for Unbuffered events raised by the StETH contract.
type StETHUnbufferedIterator struct {
	Event *StETHUnbuffered // Event containing the contract specifics and raw log

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
func (it *StETHUnbufferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHUnbuffered)
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
		it.Event = new(StETHUnbuffered)
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
func (it *StETHUnbufferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHUnbufferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHUnbuffered represents a Unbuffered event raised by the StETH contract.
type StETHUnbuffered struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbuffered is a free log retrieval operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_StETH *StETHFilterer) FilterUnbuffered(opts *bind.FilterOpts) (*StETHUnbufferedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return &StETHUnbufferedIterator{contract: _StETH.contract, event: "Unbuffered", logs: logs, sub: sub}, nil
}

// WatchUnbuffered is a free log subscription operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_StETH *StETHFilterer) WatchUnbuffered(opts *bind.WatchOpts, sink chan<- *StETHUnbuffered) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHUnbuffered)
				if err := _StETH.contract.UnpackLog(event, "Unbuffered", log); err != nil {
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

// ParseUnbuffered is a log parse operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_StETH *StETHFilterer) ParseUnbuffered(log types.Log) (*StETHUnbuffered, error) {
	event := new(StETHUnbuffered)
	if err := _StETH.contract.UnpackLog(event, "Unbuffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StETHWithdrawalsReceivedIterator is returned from FilterWithdrawalsReceived and is used to iterate over the raw logs and unpacked data for WithdrawalsReceived events raised by the StETH contract.
type StETHWithdrawalsReceivedIterator struct {
	Event *StETHWithdrawalsReceived // Event containing the contract specifics and raw log

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
func (it *StETHWithdrawalsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StETHWithdrawalsReceived)
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
		it.Event = new(StETHWithdrawalsReceived)
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
func (it *StETHWithdrawalsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StETHWithdrawalsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StETHWithdrawalsReceived represents a WithdrawalsReceived event raised by the StETH contract.
type StETHWithdrawalsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsReceived is a free log retrieval operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_StETH *StETHFilterer) FilterWithdrawalsReceived(opts *bind.FilterOpts) (*StETHWithdrawalsReceivedIterator, error) {

	logs, sub, err := _StETH.contract.FilterLogs(opts, "WithdrawalsReceived")
	if err != nil {
		return nil, err
	}
	return &StETHWithdrawalsReceivedIterator{contract: _StETH.contract, event: "WithdrawalsReceived", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsReceived is a free log subscription operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_StETH *StETHFilterer) WatchWithdrawalsReceived(opts *bind.WatchOpts, sink chan<- *StETHWithdrawalsReceived) (event.Subscription, error) {

	logs, sub, err := _StETH.contract.WatchLogs(opts, "WithdrawalsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StETHWithdrawalsReceived)
				if err := _StETH.contract.UnpackLog(event, "WithdrawalsReceived", log); err != nil {
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

// ParseWithdrawalsReceived is a log parse operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_StETH *StETHFilterer) ParseWithdrawalsReceived(log types.Log) (*StETHWithdrawalsReceived, error) {
	event := new(StETHWithdrawalsReceived)
	if err := _StETH.contract.UnpackLog(event, "WithdrawalsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
