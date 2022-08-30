- # 基于以太坊Ethereum-IPFS的去中心化Ebay区块链项目实战

[solidity 0.8.0](https://learnblockchain.cn/docs/solidity/0.8.0/index.html) 
[solidity-by-example](https://solidity-by-example.org/)   

- [Ebay开发基础](#ebay开发基础)
  - [truffle 使用](#truffle-使用)
    - [安装](#安装)
    - [测试及开发](#测试及开发)
  - [](#)
- [ipfs](#ipfs)
- [bay 维克里拍卖](#bay-维克里拍卖)
- [Web Product](#web-product)
- [Web Auction](#web-auction)
- [增加托管仲裁合约](#增加托管仲裁合约)
- [离线处理](#离线处理)

# Ebay开发基础
## truffle 使用 

[官网](https://trufflesuite.com/docs/truffle/)  

### 安装 
- `Node.js v12 - v16`    

```shell
npm install -g truffle
```

- `contracts/`: Directory for Solidity contracts
- `migrations/`: Directory for scriptable deployment files
- `test/`: Directory for test files for testing your application and contracts
- `truffle.js`: Truffle configuration file  

### 测试及开发

创建项目`truffle init`

输出日志
```shell
▶ truffle init
Debugger listening on ws://127.0.0.1:58528/4249a889-8b03-4caf-89b2-6795c9a77a3f
For help, see: https://nodejs.org/en/docs/inspector
Debugger attached.

Starting init...
================

> Copying project files to /Users/ymm/work/mygithub/ethereum/code/truffle/1-init

Init successful, sweet!

Try our scaffold commands to get started:
  $ truffle create contract YourContractName # scaffold a contract
  $ truffle create test YourTestName         # scaffold a test

http://trufflesuite.com/docs

Waiting for the debugger to disconnect...
```

可以先测试`truffle test`
```shell
truffle test .

Compiling your contracts...
===========================
> Everything is up to date, there is nothing to compile.
✓ Downloading compiler. Attempt #1.

  0 passing (2ms)
```

编译`truffle compile`  
```shell
Compiling your contracts...
===========================
> Everything is up to date, there is nothing to compile.
```

在含有`truffle-config.js` 的目录执行 `truffle develop`
```shell
truffle develop
Truffle Develop started at http://127.0.0.1:9545/

Accounts:
(0) 0x6153ecc1c902b938192c89ffe582e2ff37587d1d
(1) 0xb74f6ae85fc5b6d0c7be64ba4836eee7fee96c38
(2) 0x806638079c615bfc0fc2c82d1d89d9ddabf798c4
(3) 0x97d3aae2660cdf92a950349ebf45d33a1eb06884
(4) 0x9727c3c147c01c87e84207f5bf9d3011db2f9f0c
(5) 0xf36a8a1529d2d32cbd8f5a88efbddd0dd675e11d
(6) 0xe9df8416e3a1fdcdaea8bb6906dd8cc83e37dd56
(7) 0xe017140a56c24ffaf9d60f8b5809425065a6539f
(8) 0x7f93c03315554355701b2683607abf48b1ace4e6
(9) 0x7978c7990851b74e0f385acce8c9a54d44d5e447

Private Keys:
(0) e5b118f0facc9eda9b7201c48f9e85c59c50cd3ac8a2020dabc70d640e8d6cda
(1) fb0298ca681253747479e1ee0fc6cc0e3a1d5c19843da46f8847ca43c41ad25e
(2) f132fe70d632b27a3fd7a7e0b3db1555616b7ded12677d9a09b3609112a28d0a
(3) 0f44065095ddc88503113634c4b2e1f3eb5fef6f297451d46e349d49af65d478
(4) e0a9438cf9ca64f3e55542eb8f3b48b8676b4ea37b7766406357e9226b6a07e4
(5) 507c62060acb56911fcbbf8e3a5286e3e0371da6a14c9e094c43dc4e0a42e332
(6) d7289570f9c27cac2a168ac2970492d8b30a18698257fd522195b5966b85883e
(7) fc9382932821b3ebb92dfec034891912356da5ea98f3c3019069667fe777e08d
(8) 3e495d976ddb0886a9b6165c2edc18d5d96625b4693aaa35285db170e5f8183e
(9) 9f9a9e51186a51c7e33fcea366248cef07ecb25cd84c60239a6a5140ea08c87e

Mnemonic: liquid fruit frozen giant border weasel clock empower soldier sweet nation extra

⚠️  Important ⚠️  : This mnemonic was created for you by Truffle. It is not secure.
Ensure you do not use it on production blockchains, or else you risk losing funds.
```

在`develop`环境可以直接执行编译及部署`truffle(develop)> migrate`   

## 前端与合约交互  
### 


# ipfs
# bay 维克里拍卖
# Web Product
# Web Auction
# 增加托管仲裁合约 
# 离线处理