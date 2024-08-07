// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.19;

import "./BaseReadOnlyACL.sol";

contract Spark_AO_ACL is BaseReadOnlyACL {
    bytes32 public constant NAME = "Spark_AO_ACL";
    uint256 public constant VERSION = 3;

    address public constant AO_Stake_Contract =
    0xfE08D40Eee53d64936D3128838867c867602665c;

    address public constant Spark_Stake_Contract =
    0xC13e21B648A5Ee794902342038FF3aDAB66BE987;

    bytes32 public constant Arweave_Address_Hex =
    0xac202616a9aa22fe8ae0476a7dadba4ae8a2feef200466b290f05f4a54ca1e64;

    address public constant Erc20DAI =
    0x6B175474E89094C44Da98b954EedeAC495271d0F;
    address public constant Erc20wstETH =
    0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0;
    address public constant Erc20stETH =
    0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84;

    address public constant SAFE =
    0xaCce2d0b77BcB844a1B95c64696b071aB8c7aDA3;

    constructor() BaseOwnable(msg.sender) {}

    //ao stake & withdraw
    function stake(
        uint256 poolId_,
        uint256 amount_,
        bytes32 arweaveAddress_
    ) external pure onlyContract(AO_Stake_Contract) {
        require(poolId_ == 0, "invalid poolId");
        require(arweaveAddress_ == Arweave_Address_Hex, "invalid arweave address");
    }

    function withdraw(
        uint256 poolId_,
        uint256 amount_,
        bytes32 arweaveAddress_
    ) external pure onlyContract(AO_Stake_Contract) {
        require(poolId_ == 0, "invalid poolId");
        require(arweaveAddress_ == Arweave_Address_Hex, "invalid arweave address");
    }

    //spark supply & withdraw, borrow & repay
    function supply(
        address asset,
        uint256 amount,
        address onBehalfOf,
        uint16 referralCode
    ) external view onlyContract(Spark_Stake_Contract)  {
        require(onBehalfOf == SAFE, "invalid to address");
    }

    function withdraw(
        address asset,
        uint256 amount,
        address to
    ) external view onlyContract(Spark_Stake_Contract) {
        require(to == SAFE, "invalid to address");
    }


    function borrow(
        address asset,
        uint256 amount,
        uint256 interestRateMode,
        uint16 referralCode,
        address onBehalfOf
    ) external view onlyContract(Spark_Stake_Contract) {
        require(asset == Erc20DAI, "asset not DAI");
        require(interestRateMode == 2, "invalid interestRateMode address");
        require(onBehalfOf == SAFE, "invalid onBehalfOf address");
    }

    function repay(
        address asset,
        uint256 amount,
        uint256 interestRateMode,
        address onBehalfOf
    ) external view onlyContract(Spark_Stake_Contract) {
        require(asset == Erc20DAI, "asset not DAI");
        require(interestRateMode == 2, "invalid interestRateMode address");
        require(onBehalfOf == SAFE, "invalid onBehalfOf address");
    }

    //wstETH wrap & unwrap
    function wrap(uint256 _stETHAmount) external view onlyContract(Erc20wstETH) {}
    function unwrap(uint256 _wstETHAmount) external view onlyContract(Erc20wstETH) {}

    //stETH
    function submit(address _referral) external view onlyContract(Erc20stETH) {}
}