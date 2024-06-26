# Spark & AO 调仓平衡

使用Cobo Argus来保证资金安全，监控到调仓条件时通过argus bot发起操作。  
操作的bot需要加一个`余额监控`来保证有足够的资金来发起调仓tx


## 代码逻辑：  
 1. ws监听到新块时，获取spark合约中address provider-> oracle address -> wstETH price
 2. 根据每个块当时的`health factor`来判断是否需要调仓，调仓额度可以通过 spark中 缺少/溢出的health factor来计算，再换算成 stETH/wstETH 

Health Factor计算方法 ([from spark contract: GenericLogic.sol: 172](https://etherscan.io/address/0x5ae329203e00f76891094dcfedd5aca082a50e1b#code)):  
```solidity
    vars.healthFactor = (vars.totalDebtInBaseCurrency == 0)
      ? type(uint256).max
      : (vars.totalCollateralInBaseCurrency.percentMul(vars.avgLiquidationThreshold)).wadDiv(
        vars.totalDebtInBaseCurrency
      );
```

可以通过公式来计算出要调仓的额度来进一步计算对应的stETH的数量

## 运行代码

代码打包成docker运行，[镜像地址](https://hub.docker.com/r/horizont9/spark_ao/tags)

```dockerfile
docker run -d --restart=always --name spark_ao -v /path/to/config.yaml:/app/conf/config.yaml horizont9/spark_ao:v0.0.2
```

根据[config模板](./conf/config.yaml.example)填写参数并挂载到容器下  

代码中如果检测到block时间与本地时间戳存在超时的情况（32s）下会触发panic
```go
func handlePanic(msg string) {
	notify.SendMsg("❌Spark&AO get panic and try restart", "❌ go check auto restart status, error msg: "+msg)
	log.Panicf(msg)
}
```
会导致程序停止，配合docker中 ```--restart=aways``` 来先进行自动重启并发送通知，需要查看日志进行进一步排查原因。

## Cobo Argus相关规则配置

1. [Cobo Argus ACL规则](./acl/spark_ao_acl.sol)
2. 提前进行Approve操作，涉及的Approve行为有：

```text
1. 合约 0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84 (stETH)
	approve(address _spender,uint256 _amount) 
	#	Name		Type	Data
	0	_spender	address	0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0 （wstETH）
	1	_amount		uint256	额度

	用途：wrap stETH用到, spark只接受wstETH

2. 0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0 (wstETH)
	approve approve(address _spender,uint256 _amount)
	#	Name		Type	Data
	0	_spender	address	0xC13e21B648A5Ee794902342038FF3aDAB66BE987 （spark protocol:Staking）
	1	_amount		uint256	额度

	用途：supply wstETH 到 spark

3. 0x6B175474E89094C44Da98b954EedeAC495271d0F (DAI)
	approve(address _spender,uint256 _amount)  
	#	Name		Type	Data
	0	_spender	address	0xC13e21B648A5Ee794902342038FF3aDAB66BE987 （spark protocol:Staking）
	1	_amount		uint256	额度

	用途：repay DAI给spark。还款DAI时用到

4. 0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84 (stETH) 
	approve(address _spender,uint256 _amount)  
	#	Name		Type	Data
	0	_spender	address	0xfE08D40Eee53d64936D3128838867c867602665c （ao staking合约）
	1	_amount		uint256	额度

	用途：stake stETH 到 ao stake 合约
```