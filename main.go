package main

import (
	"context"
	"encoding/json"
	"errors"
	"eth_spark_ao/pkg/base"
	"eth_spark_ao/pkg/contracts/ao"
	spark_address_privider "eth_spark_ao/pkg/contracts/spark_address_provider"
	"eth_spark_ao/pkg/contracts/spark_price_oracle"
	"eth_spark_ao/pkg/contracts/spark_staking"
	"eth_spark_ao/pkg/contracts/wstETH"
	"eth_spark_ao/pkg/slack"
	"eth_spark_ao/util"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/imskyd/argus_generator/argus"
	"github.com/imskyd/ethers_helper"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"sync"
	"time"
)

const SparkStakingAddress = "0xC13e21B648A5Ee794902342038FF3aDAB66BE987"
const AOStakingAddress = "0xfe08d40eee53d64936d3128838867c867602665c"
const wstETHAddress = "0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0"
const DAIAddress = "0x6B175474E89094C44Da98b954EedeAC495271d0F"

var SparkTooLowHealthFactor = big.NewInt(15 * 1e17)
var SparkTooHighHealthFactor = big.NewInt(155 * 1e16)

var logger *logrus.Logger
var client *ethclient.Client
var sparkStaking *spark_staking.SparkStaking
var aoStaking *ao.Ao
var wstETHContract *wstETH.WstETH
var argusService *argus.ArgusService
var arweaveAddress [32]byte
var movingFromSparkToAO, movingFromAOToSpark bool
var locker *sync.RWMutex
var BlockTime uint64
var config *base.ArgusConfig
var notify *slack.Slack

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.StampMilli,
	})

	//========config
	config = base.GetArgusConfig()
	//config check
	res, err, pkAddress := ethers_helper.AddressMustBelongToPrivateKey(config.BotAddress, config.BotPk)
	if err != nil {
		handlePanic("config ethers_helper.AddressMustBelongToPrivateKey check err: " + err.Error())
	}
	if !res {
		handlePanic("bot pk address not match bot address: " + pkAddress)
	}
	data := common.FromHex(config.ArweaveAddressHex)
	if len(data) != 32 {
		handlePanic("arweave address lenth not 32")
	}
	copy(arweaveAddress[:], data)

	slackConfig := base.GetSlackConfig()
	notify = slack.NewSlack(slackConfig.Token, slackConfig.Channel)

	client, _ = ethclient.Dial(config.Ws)
	sparkStaking, _ = spark_staking.NewSparkStaking(common.HexToAddress(SparkStakingAddress), client)
	aoStaking, _ = ao.NewAo(common.HexToAddress(AOStakingAddress), client)
	wstETHContract, _ = wstETH.NewWstETH(common.HexToAddress(wstETHAddress), client)
	_argusService, err := argus.NewArgusService(config.Rpc, config.Argus, config.BotPk)
	if err != nil {
		handlePanic("argus.NewArgusService err: %s" + err.Error())
	}
	argusService = _argusService

	locker = &sync.RWMutex{}
}

func LiveClock() {
	time.Sleep(time.Second * 15)
	for true {
		if time.Now().Unix()-int64(BlockTime) > 36 {
			logger.Errorf("live check failed")
			notify.SendMsg("Spark AO live check failed and restart", "Spark AO live check failed and restart")
			handlePanic("live check failed")
		}
		time.Sleep(time.Second)
	}
}

func handleErr(err error) {
	logger.Errorf(err.Error())
	go notify.SendMsg("❌Spark&AO get error", "❌ "+err.Error())
}

func handlePanic(msg string) {
	notify.SendMsg("❌Spark&AO get panic", "❌ go check auto restart and error msg: "+msg)
	log.Panicf(msg)
}

func main() {

	aoStaking, _ = ao.NewAo(common.HexToAddress(AOStakingAddress), client)
	//notify.SendMsg("Spark AO Start", "Spark AO Start")
	configJ, _ := json.Marshal(config)
	logger.Infof("config auto check: %s", configJ)
	go LiveClock()

	headers := make(chan *types.Header)
	sub, _ := client.SubscribeNewHead(context.Background(), headers)

	for {
		select {
		case err := <-sub.Err():
			handlePanic("wss err: " + err.Error())
		case header := <-headers:
			logger.Infof("==============================")
			logger.Infof("[Block] block height: %d, blockTime: %d", header.Number, header.Time)
			BlockTime = header.Time
			suggestGas, _ := client.SuggestGasPrice(context.Background())
			priorityGas := new(big.Int).Mul(suggestGas, big.NewInt(17))
			priorityGas = new(big.Int).Div(priorityGas, big.NewInt(10))
			logger.Infof("[Gas] suggest gas: %f, priority gas: %f", util.ToDecimal(suggestGas, 9).InexactFloat64(), util.ToDecimal(priorityGas, 9).InexactFloat64())

			//get now block's oracle address and price info
			sparkAddressProviderAddr, err := sparkStaking.ADDRESSESPROVIDER(nil)
			if err != nil {
				handleErr(fmt.Errorf("get ADDRESSESPROVIDER err: %s", err.Error()))
				continue
			}
			sparkAddressProviderContract, err := spark_address_privider.NewSparkAddressPrivider(sparkAddressProviderAddr, client)
			if err != nil {
				handleErr(fmt.Errorf("spark_address_privider.NewSparkAddressPrivider err: %s", err.Error()))
				continue
			}
			priceOracleAddr, err := sparkAddressProviderContract.GetPriceOracle(nil)
			if err != nil {
				handleErr(fmt.Errorf("sparkAddressProviderContract.GetPriceOracle err: %s", err.Error()))
				continue
			}
			priceOracle, err := spark_price_oracle.NewSparkPriceOracle(priceOracleAddr, client)
			if err != nil {
				handleErr(fmt.Errorf("spark_price_oracle.NewSparkPriceOracle err: %s", err.Error()))
				continue
			}
			wstETHPrice, err := priceOracle.GetAssetPrice(nil, common.HexToAddress(wstETHAddress))
			if err != nil {
				handleErr(fmt.Errorf("priceOracle.GetAssetPrice err: %s", err.Error()))
				continue
			}
			decimals, err := priceOracle.BASECURRENCYUNIT(nil)
			if err != nil {
				handleErr(fmt.Errorf("priceOracle.BASECURRENCYUNIT err: %s", err.Error()))
				continue
			}

			num, _ := decimal.NewFromString(wstETHPrice.String())
			mul := decimal.NewFromInt(decimals.Int64())
			wstETHPriceFormat := num.Div(mul)

			logger.Infof("[Oracle] wstETH price: %f", wstETHPriceFormat.InexactFloat64())
			if wstETHPriceFormat.InexactFloat64() == 0 {
				handleErr(fmt.Errorf("[Oracle] wstETH price is 0"))
				continue
			}

			userAccountData, err := sparkStaking.GetUserAccountData(nil, common.HexToAddress(config.Safe))
			if err != nil {
				handleErr(fmt.Errorf("get healthFactor err: %s", err.Error()))
				continue
			}
			//抵押物 || 债务 为0时跳过
			if big.NewInt(0).Cmp(userAccountData.TotalCollateralBase) == 0 {
				logger.Warnf("TotalCollateralBase is 0, continue ...")
				continue
			}
			if big.NewInt(0).Cmp(userAccountData.TotalDebtBase) == 0 {
				logger.Warnf("TotalDebtBase is 0, continue ...")
				continue
			}
			healthFactor := userAccountData.HealthFactor
			logger.Infof("health factor: %f", util.ToDecimal(healthFactor, 18).InexactFloat64())
			liquidationThreshold := util.ToDecimal(userAccountData.CurrentLiquidationThreshold, 4).InexactFloat64()
			logger.Infof("liquidition threshold: %f", liquidationThreshold)

			totalCollateralBase := util.ToDecimal(userAccountData.TotalCollateralBase, 8).InexactFloat64()
			totalDebtBase := util.ToDecimal(userAccountData.TotalDebtBase, 8).InexactFloat64()

			if healthFactor.Cmp(SparkTooLowHealthFactor) == -1 {
				// health factor < 1.5 move ao's stETH to spark
				if !startMovingFromAOToSpark() {
					continue
				}
				logger.Warnf("health factor < 1.5(%f), move AO stETH Spark", util.ToDecimal(healthFactor, 18).InexactFloat64())
				targetSupplyAmount := 1.52 * totalDebtBase / liquidationThreshold
				lackAmount := targetSupplyAmount - totalCollateralBase
				lackWstETH := lackAmount / wstETHPriceFormat.InexactFloat64()
				lackStETH, err := wstETHContract.GetStETHByWstETH(nil, util.ToWei(lackWstETH, 18))
				if err != nil {
					handleErr(fmt.Errorf("wstETHContract.GetStETHByWstETH err: %s", err.Error()))
					finishMovingFromSparkToAO()
					continue
				}
				wstETHToStETHFormat := util.ToDecimal(lackStETH, 18)
				logger.Warnf("AO need remove %f stETH and wrap to wstETH(%f) to supply Spark", wstETHToStETHFormat.InexactFloat64(), lackWstETH)
				finalWstETH, err := wstETHContract.GetWstETHByStETH(nil, lackStETH)
				if err != nil {
					handleErr(fmt.Errorf("wstETHContract.GetWstETHByStETH err: %s", err.Error()))
					finishMovingFromSparkToAO()
					continue
				}
				//start combine argus tx
				var cdl []argus.CallData
				c1 := MakeAOWithdrawCallData(big.NewInt(0), lackStETH, arweaveAddress)
				cdl = append(cdl, c1)
				c2 := MakeWrapStETHCallData(lackStETH)
				cdl = append(cdl, c2)
				c3 := MakeSupplyWstETHToSparkCallData(common.HexToAddress(wstETHAddress), finalWstETH, common.HexToAddress(config.Safe), 0)
				cdl = append(cdl, c3)
				opt, _ := argusService.GetTransactOpts()
				//opt.NoSend = true
				opt.GasPrice = priorityGas
				tx, err := argusService.ExecTransactions(opt, cdl)
				if err != nil {
					logger.Errorf("exec ao -> spark err: %s", err.Error())
					finishMovingFromAOToSpark()
					continue
				}
				j, _ := json.Marshal(tx)
				logger.Infof("exec tx: %s", string(j))
				if errC := checkTx(tx); errC != nil {
					handleErr(errC)
				} else {
					logger.Warnf("exec ao -> spark success: %s", tx.Hash().String())
					msg := fmt.Sprintf("✅ Spark health %f, move %f stETH from AO to Spark. tx: https://etherscan.io/tx/%s", util.ToDecimal(healthFactor, 18).InexactFloat64(), util.ToDecimal(lackStETH, 18).InexactFloat64(), tx.Hash().String())
					go notify.SendMsg("Spark&AO Health Factor Changing", msg)
				}
				finishMovingFromAOToSpark()
			}
			if healthFactor.Cmp(SparkTooHighHealthFactor) == 1 {
				// health factor > 1.55, move wstETH and supply to AO
				if !startMovingFromSparkToAO() {
					continue
				}
				logger.Warnf("health factor > 1.55(%f), move Spark wstETH and supply to AO", util.ToDecimal(healthFactor, 18).InexactFloat64())
				targetSupplyAmount := 1.52 * totalDebtBase / liquidationThreshold
				logger.Warnf("target amount value is: %f", targetSupplyAmount)
				logger.Warnf("remove amount value is: %f", totalCollateralBase-targetSupplyAmount)
				removeWstETHAmount := (totalCollateralBase - targetSupplyAmount) / wstETHPriceFormat.InexactFloat64()
				logger.Warnf("remove wstETH amount is: %f", removeWstETHAmount)
				//unwrap to stETH -> deposit to ao
				stETHAmount, err := wstETHContract.GetStETHByWstETH(nil, util.ToWei(removeWstETHAmount, 18))
				if err != nil {
					handleErr(fmt.Errorf("wstETHContract.GetStETHByWstETH err: %s", err.Error()))
					finishMovingFromSparkToAO()
					continue
				}
				stETHAmountFormat := util.ToDecimal(stETHAmount, 18)
				logger.Warnf("unwrap to stETH amount is: %f", stETHAmountFormat.InexactFloat64())

				//start combine argus tx
				var cdl []argus.CallData
				c1 := MakeWithdrawWstETHToSparkCallData(common.HexToAddress(wstETHAddress), util.ToWei(removeWstETHAmount, 18), common.HexToAddress(config.Safe))
				cdl = append(cdl, c1)
				c2 := MakeUnWrapWstETHCallData(util.ToWei(removeWstETHAmount, 18))
				cdl = append(cdl, c2)
				c3 := MakeAOStakeCallData(big.NewInt(0), stETHAmount, arweaveAddress)
				cdl = append(cdl, c3)
				opt, _ := argusService.GetTransactOpts()
				//opt.NoSend = true
				opt.GasPrice = priorityGas
				tx, err := argusService.ExecTransactions(opt, cdl)
				if err != nil {
					logger.Errorf("exec spark -> ao err: %s", err.Error())
					finishMovingFromSparkToAO()
					continue
				}
				j, _ := json.Marshal(tx)
				logger.Infof("exec tx: %s", string(j))
				if errC := checkTx(tx); errC != nil {
					handleErr(errC)
				} else {
					logger.Warnf("exec spark -> ao success: %s", tx.Hash().String())
					msg := fmt.Sprintf("✅ Spark health %f, move %f stETH from Spark to AO. tx: https://etherscan.io/tx/%s", util.ToDecimal(healthFactor, 18).InexactFloat64(), util.ToDecimal(stETHAmount, 18).InexactFloat64(), tx.Hash().String())
					go notify.SendMsg("Spark&AO Health Factor Changing", msg)
				}
				finishMovingFromSparkToAO()
			}
		}
	}
}

// =============================================== ao -> spark
// 1. withdraw stETH from ao
// 2. wrap stETH to wstETH
// 3. supply wstETH to spark

// MakeAOWithdrawCallData
// e.g: https://etherscan.io/tx/0x6748a91d3d34e4d26eaee974d39b51c1c2fbb2eac616ea626fefc3931bf628e0
func MakeAOWithdrawCallData(poolId_ *big.Int, amount_ *big.Int, arweaveAddress_ [32]byte) argus.CallData {
	txData, err := argus.MakeTxDataByAbi(ao.AoMetaData.ABI, "withdraw", poolId_, amount_, arweaveAddress_)
	if err != nil {
		handlePanic("argus.MakeTxDataByAbi err: " + err.Error())
	}
	callData := argus.MakeCallData(common.HexToAddress(AOStakingAddress), big.NewInt(0), txData)
	return callData
}

// MakeWrapStETHCallData
// e.g: https://etherscan.io/tx/0x7b2195ade05c91919357a9ce14c5c4e1095b084ee2ff5226a0f15995957be558
func MakeWrapStETHCallData(_stETHAmount *big.Int) argus.CallData {
	txData, err := argus.MakeTxDataByAbi(wstETH.WstETHMetaData.ABI, "wrap", _stETHAmount)
	if err != nil {
		handlePanic("argus.MakeTxDataByAbi err: " + err.Error())
	}
	callData := argus.MakeCallData(common.HexToAddress(wstETHAddress), big.NewInt(0), txData)
	return callData
}

// MakeSupplyWstETHToSparkCallData
// e.g: https://etherscan.io/tx/0x447a88cd9245cf5d9e86e2ca90029ce49a097de974fca17152bd4897b69f0613
func MakeSupplyWstETHToSparkCallData(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) argus.CallData {
	txData, err := argus.MakeTxDataByAbi(spark_staking.SparkStakingMetaData.ABI, "supply", asset, amount, onBehalfOf, referralCode)
	if err != nil {
		handlePanic("argus.MakeTxDataByAbi err: " + err.Error())
	}
	callData := argus.MakeCallData(common.HexToAddress(SparkStakingAddress), big.NewInt(0), txData)
	return callData
}

// =============================================== spark -> ao
// 1. withdraw from spark
// 2. unwrap wstETH to stETH
// 3. deposit stETH to ao

// MakeWithdrawWstETHToSparkCallData
// e.g: https://etherscan.io/tx/0x658969817aea91ffd534cceed5776d4027d9527d4071d60f35e5b8537f684199
func MakeWithdrawWstETHToSparkCallData(asset common.Address, amount *big.Int, to common.Address) argus.CallData {
	txData, err := argus.MakeTxDataByAbi(spark_staking.SparkStakingMetaData.ABI, "withdraw", asset, amount, to)
	if err != nil {
		handlePanic("argus.MakeTxDataByAbi err: " + err.Error())
	}
	callData := argus.MakeCallData(common.HexToAddress(SparkStakingAddress), big.NewInt(0), txData)
	return callData
}

// MakeUnWrapWstETHCallData
// e.g: https://etherscan.io/tx/0xa6a91abbe04db70745143ec72821d1f922e4048ae5bedde9ffb23dc2de69e592
func MakeUnWrapWstETHCallData(_wstETHAmount *big.Int) argus.CallData {
	txData, err := argus.MakeTxDataByAbi(wstETH.WstETHMetaData.ABI, "unwrap", _wstETHAmount)
	if err != nil {
		handlePanic("argus.MakeTxDataByAbi err: %s" + err.Error())
	}
	callData := argus.MakeCallData(common.HexToAddress(wstETHAddress), big.NewInt(0), txData)
	return callData
}

// MakeAOStakeCallData
// e.g: https://etherscan.io/tx/0x40043b879b506f25d4fce7d721d7574583e3e603a379680ec725e6f6632234a0
func MakeAOStakeCallData(poolId_ *big.Int, amount_ *big.Int, arweaveAddress_ [32]byte) argus.CallData {
	txData, err := argus.MakeTxDataByAbi(ao.AoMetaData.ABI, "stake", poolId_, amount_, arweaveAddress_)
	if err != nil {
		handlePanic("argus.MakeTxDataByAbi err: " + err.Error())
	}
	callData := argus.MakeCallData(common.HexToAddress(AOStakingAddress), big.NewInt(0), txData)
	return callData
}

// ======================================== Lockers
// movingFromSparkToAO, movingFromAOToSpark
func startMovingFromSparkToAO() bool {
	locker.Lock()
	if movingFromSparkToAO {
		return false
	}
	movingFromSparkToAO = true
	locker.Unlock()
	return true
}

func finishMovingFromSparkToAO() {
	locker.Lock()
	movingFromSparkToAO = false
	locker.Unlock()
}

func startMovingFromAOToSpark() bool {
	locker.Lock()
	if movingFromAOToSpark {
		return false
	}
	movingFromAOToSpark = true
	locker.Unlock()
	return true
}

func finishMovingFromAOToSpark() {
	locker.Lock()
	movingFromAOToSpark = false
	locker.Unlock()
}

func checkTx(tx *types.Transaction) error {
	failedTime := 0
	for true {
		if failedTime > 30 {
			return errors.New("tx trigger max failed time: " + tx.Hash().String())
		}
		res, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			logger.Errorf("client.TransactionReceipt err: %s", err.Error())
			failedTime++
			time.Sleep(time.Second)
			continue
		}
		if res.Status == types.ReceiptStatusSuccessful {
			logger.Warnf("tx success: %s", tx.Hash())
			return nil
		}
		if res.Status == types.ReceiptStatusFailed {
			logger.Errorf("tx failed: %s", tx.Hash())
			return errors.New("tx failed: " + tx.Hash().String())
		}
	}
	return nil
}

// ======================= init

func MakeSparkBorrowDAICallData(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) argus.CallData {
	txData, err := argus.MakeTxDataByAbi(spark_staking.SparkStakingMetaData.ABI, "borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
	if err != nil {
		handlePanic("argus.MakeTxDataByAbi err: " + err.Error())
	}
	callData := argus.MakeCallData(common.HexToAddress(SparkStakingAddress), big.NewInt(0), txData)
	return callData
}

func MakeSparkRepayDAICallData(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) argus.CallData {
	txData, err := argus.MakeTxDataByAbi(spark_staking.SparkStakingMetaData.ABI, "repay", asset, amount, interestRateMode, onBehalfOf)
	if err != nil {
		handlePanic("argus.MakeTxDataByAbi err: " + err.Error())
	}
	callData := argus.MakeCallData(common.HexToAddress(SparkStakingAddress), big.NewInt(0), txData)
	return callData
}
