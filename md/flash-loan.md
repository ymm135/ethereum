- ## flash-loan  

## 参考文章
- [aave performing a flash loan ](https://docs.aave.com/developers/v/1.0/tutorials/performing-a-flash-loan)  
- [Create Flashloan Combo on Furucombo](https://medium.com/furucombo/create-flashloan-combo-on-furucombo-c7c3b23267f0)  

## Furucombo
### [Learn](https://docs.furucombo.app/) 

`Furucombo` 从为终端用户构建的工具开始，只需通过`拖放`即可优化他们的 `DeFi` 策略，可帮助您优化您的加密财富。  
- `投资模式` 轻松探索和投资富康理财的财富管理农场，让我们帮助您充分利用您的资金。  
- `创建模式` 它将复杂的 `DeFi` 协议可视化为立方体。 用户设置输入/输出和    

### 展示
```sh
Rate difference: 20+%
1 DAI = 0.9927 sUSD on Kyberswap
1 DAI = 0.8057 sUSD on Uniswap
👉🏻 Buy low sell high: Buy sUSD on Uniswap and sell it on Kyberswap
```

现在您找到了速率差异，让我们开始创建组合。完整的组合应该是这样的：

```sh
1️⃣ Borrow 100DAI from Flashloan
2️⃣ Swap 100DAI to 122.83649sUSD on Uniswap
3️⃣ Swap 122.83649sUSD to 122.83429DAI on Kyberswap
4️⃣ Repay 100.09DAI to Flashloan
5️⃣ You keep 22.74429DAI profit.
```

<br>
<div align=center>
  <img src="../res/images/defi-5.png" width="100%"></img>
</div>

## 疑问及拓展
### 矿池(Pools)  

矿池（英文：Mining Pool），最早指`比特币矿池`，后来泛指`POW`矿池，目前更发展出`POC`矿池。

在加密货币挖矿的背景下，挖矿池是由矿工在网络共享其处理能力的资源池，以根据他们为找到区块的可能性贡献的工作量平均分配奖励。向出示有效的部分工作量证明的矿池成员授予“份额”。当采矿的难度增加到缓慢的采矿者可能需要几个世纪才能形成块的程度时，才开始在池中采矿。

截止2021年12月，全球算力前三的比特币矿池为AntPool、F2Pool、ViaBTC。全球约七成算力在中国矿工手中。  

> POW（Proof of work工作量证明）——区块链世界最早的共识机制  
> POS（Proof of stake权益证明）——POW强有力的挑战者  
> POC（Proof of capacity容量证明）——即将成为区块链世界受众最多的共识？





