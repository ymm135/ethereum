# go-ethereum
- ### [https://github.com/ethereum/go-ethereum](https://github.com/ethereum/go-ethereum)  Go语言实现  
- ### [https://github.com/hyperledger/fabric](https://github.com/hyperledger/fabric)  Go语言实现  


## 安装
### 源码编译
```shell
git clone git@github.com:ethereum/go-ethereum.git

# geth
make geth 

# to build the full suite of utilities 
make all
```

### Ubuntu via PPAs
```shell
sudo add-apt-repository -y ppa:ethereum/ethereum

sudo apt-get update
sudo apt-get install ethereum
```

如果已经下载需要更新版本
```bash
sudo apt-get update
sudo apt-get install ethereum
sudo apt-get upgrade geth
```

安装的bin文件
Setting up `rlpdump` (1.10.21+build27994+focal) ...  
Setting up `puppeth` (1.10.21+build27994+focal) ...  
Setting up `clef` (1.10.21+build27994+focal) ...  
Setting up `bootnode` (1.10.21+build27994+focal) ...  
Setting up `geth` (1.10.21+build27994+focal) ...  
Setting up `evm` (1.10.21+build27994+focal) ...  
Setting up `abigen` (1.10.21+build27994+focal) ...  
Setting up `ethereum` (1.10.21+build27994+focal) ...  


### 安装包下载
[https://geth.ethereum.org/downloads/](https://geth.ethereum.org/downloads/)  

## 搭建私有链(Ubuntu)
### 安装geth
:point_right: [安装方法](#安装)  

也可以下载离线包，添加到环境PATH即可,离线包内容为:  
```shell
COPYING  abigen   bootnode clef     evm      geth     puppeth  rlpdump
```

### 配置`创世文件`
`genesis.json`  
```json
{
  "config": {
    "chainId": <arbitrary positive integer>,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "berlinBlock": 0,
    "londonBlock": 0
  },
  "alloc": {},
  "coinbase": "0x0000000000000000000000000000000000000000",
  "difficulty": "0x20000",
  "extraData": "",
  "gasLimit": "0x2fefd8",
  "nonce": "0x0000000000000042",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
```

如果您想预先为某些帐户提供资金以便于测试，请创建帐户并`alloc`使用其地址填充该字段。

```json
"alloc": {
  "0x0000000000000000000000000000000000000001": {
    "balance": "111111111"
  },
  "0x0000000000000000000000000000000000000002": {
    "balance": "222222222"
  }
}
```

最终配置文件如下:
```json
{
	"config": {
		"chainId": 123,
		"homesteadBlock": 0,
		"eip150Block": 0,
		"eip155Block": 0,
		"eip158Block": 0,
		"byzantiumBlock": 0,
		"constantinopleBlock": 0,
		"petersburgBlock": 0,
		"istanbulBlock": 0,
		"berlinBlock": 0,
		"londonBlock": 0
	},
	"alloc": {
		"0x0000000000000000000000000000000000000001": {
			"balance": "111111111"
		},
		"0x0000000000000000000000000000000000000002": {
			"balance": "222222222"
		}
	},
	"coinbase": "0x0000000000000000000000000000000000000000",
	"difficulty": "0x20000",
	"extraData": "",
	"gasLimit": "0x2fefd8",
	"nonce": "0x0000000000000042",
	"mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"timestamp": "0x00"
}
```

### 初始化
```shell
geth init genesis.json
```

初始化输出
```shell
INFO [08-12|17:30:43.367] Maximum peer count                       ETH=50 LES=0 total=50
INFO [08-12|17:30:43.369] Smartcard socket not found, disabling    err="stat /run/pcscd/pcscd.comm: no such file or directory"
INFO [08-12|17:30:43.376] Set global gas cap                       cap=50,000,000
INFO [08-12|17:30:43.378] Allocated cache and file handles         database=/root/.ethereum/geth/chaindata cache=16.00MiB handles=16
INFO [08-12|17:30:43.415] Opened ancient database                  database=/root/.ethereum/geth/chaindata/ancient readonly=false
INFO [08-12|17:30:43.416] Writing custom genesis block 
INFO [08-12|17:30:43.416] Persisted trie from memory database      nodes=3 size=399.00B time="70.652µs" gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
INFO [08-12|17:30:43.417] Freezer shutting down 
INFO [08-12|17:30:43.417] Successfully wrote genesis state         database=chaindata                              hash=6ab19c..97db44
INFO [08-12|17:30:43.417] Allocated cache and file handles         database=/root/.ethereum/geth/lightchaindata    cache=16.00MiB handles=16
INFO [08-12|17:30:43.434] Opened ancient database                  database=/root/.ethereum/geth/lightchaindata/ancient readonly=false
INFO [08-12|17:30:43.435] Writing custom genesis block 
INFO [08-12|17:30:43.435] Persisted trie from memory database      nodes=3 size=399.00B time="54.665µs" gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
INFO [08-12|17:30:43.435] Successfully wrote genesis state         database=lightchaindata                         hash=6ab19c..97db44
```

初始化并创建数据库

> 如果本地已经创建，可以删除数据`geth removedb`  

```shell
$ geth removedb                                  
INFO [08-14|11:42:39.420] Maximum peer count                       ETH=50 LES=0 total=50
INFO [08-14|11:42:39.422] Set global gas cap                       cap=50,000,000
Remove full node state database (/Users/ymm/Library/Ethereum/geth/chaindata)? [y/n] y
Remove full node state database (/Users/ymm/Library/Ethereum/geth/chaindata)? [y/n] y
INFO [08-14|11:42:42.620] Database successfully deleted            path=/Users/ymm/Library/Ethereum/geth/chaindata elapsed=4.043ms
Remove full node ancient database (/Users/ymm/Library/Ethereum/geth/chaindata/ancient)? [y/n] y
Remove full node ancient database (/Users/ymm/Library/Ethereum/geth/chaindata/ancient)? [y/n] y
INFO [08-14|11:42:45.504] Database successfully deleted            path=/Users/ymm/Library/Ethereum/geth/chaindata/ancient elapsed=5.639ms
INFO [08-14|11:42:45.504] Light node database missing              path=/Users/ymm/Library/Ethereum/geth/lightchaindata
```

### 创建一个集合节点

```shell
bootnode --genkey=boot.key
bootnode --nodekey=boot.key
```

查看`boot.ket`
```shell
4ebf6df1dd94b40f6e7486a558531c0b2ddca05b34e9bc46640fd95553c3afa6
```

启动日志
```shell
# bootnode --nodekey=boot.key
enode://644204e69a456937bafce4fe7a7609c485d8902e571bb8ea86f73108ecb7e0ea430865669bf8e62c3cb138bbd91194bca8b924898ac1382b4443015b1866f4d9@127.0.0.1:0?discport=30301
Note: you're using cmd/bootnode, a developer tool.
We recommend using a regular node as bootstrap node for production deployments.
INFO [08-12|17:33:43.570] New local node record                    seq=1,660,296,823,568 id=cc7190d45564ac73 ip=<nil> udp=0 tcp=0
```

当 bootnode 在线时，它将显示一个`enode` `URL` ，其他节点可以使用它来连接它并交换对等点信息。确保将显示的 IP 地址信息（很可能[::]）替换为您的外部可访问 IP 以获取实际`enodeURL`。  

> 您也可以使用成熟的geth节点作为引导节点，但这是不太推荐的方式。  

这里的url就是`enode://644204e69a456937bafce4fe7a7609c485d8902e571bb8ea86f73108ecb7e0ea430865669bf8e62c3cb138bbd91194bca8b924898ac1382b4443015b1866f4d9@127.0.0.1:0?discport=30301`  


### 启动您的成员节点
在引导节点可操作且外部可访问的情况下（您可以尝试 `telnet <ip> <port>`确保它确实可访问），启动 `geth` 指向引导节点的每个后续节点以通过`--bootnodes`标志进行对等发现。可能还需要将您的专用网络的数据目录分开，因此还要指定一个自定义`--datadir`标志。

```shell
geth --datadir=path/to/custom/data/folder --bootnodes=<bootnode-enode-url-from-above>
```

也可以使用
```shell
geth --datadir=path/to/custom/data/folder --networkid=  
```

执行指令
```shell
geth --datadir node1/ --networkid 123 --port 3100 --nodiscover
```


> `console`  命令行启动,可以直接交互  
> 如果不是交互模式，找到`data/geth.ipc`, 通过`geth attach ipc:data/geth.ipc`,可以使用的模块有:  

- admin
- debug
- engine
- eth
- ethash
- miner
- personal
- rpc
- txpool
- web3

> 交互控制就是JS控制台  

attach日志
```shell
# geth attach ipc:node1/geth.ipc
Welcome to the Geth JavaScript console!

instance: Geth/v1.10.21-stable-67109427/linux-amd64/go1.18.4
at block: 0 (Thu Jan 01 1970 08:00:00 GMT+0800 (CST))
 datadir: /root/work/ethereum/node1
 modules: admin:1.0 debug:1.0 engine:1.0 eth:1.0 ethash:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

To exit, press ctrl-d or type exit
```

启动日志
```shell
INFO [08-13|17:02:17.746] Starting Geth on Ethereum mainnet... 
INFO [08-13|17:02:17.747] Bumping default cache on mainnet         provided=1024 updated=4096
INFO [08-13|17:02:17.748] Maximum peer count                       ETH=50 LES=0 total=50
INFO [08-13|17:02:17.751] Smartcard socket not found, disabling    err="stat /run/pcscd/pcscd.comm: no such file or directory"
WARN [08-13|17:02:17.757] Sanitizing cache to Go's GC limits       provided=4096 updated=1308
INFO [08-13|17:02:17.758] Set global gas cap                       cap=50,000,000
INFO [08-13|17:02:17.759] Allocated trie memory caches             clean=196.00MiB dirty=327.00MiB
INFO [08-13|17:02:17.761] Allocated cache and file handles         database=/root/work/ethereum/data/geth/chaindata cache=653.00MiB handles=524,288
INFO [08-13|17:02:17.788] Opened ancient database                  database=/root/work/ethereum/data/geth/chaindata/ancient readonly=false
INFO [08-13|17:02:17.788] Writing default main-net genesis block 
INFO [08-13|17:02:18.176] Persisted trie from memory database      nodes=12356 size=1.78MiB time=71.751864ms gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
```

然后循环提示一下信息
```shell
INFO [08-13|17:07:59.453] State sync in progress                   synced=0.00% state=3.99MiB   accounts=1376@331.22KiB slots=13681@2.72MiB   codes=629@966.38KiB eta=8456h2m13.746s
INFO [08-13|17:08:01.569] Imported new block headers               count=192 elapsed=42.729ms  number=4416 hash=25a7e2..a7d61a age=7y1mo2w
INFO [08-13|17:08:01.571] Imported new block receipts              count=47  elapsed=38.743ms  number=410  hash=d83833..7aa4ff age=7y1mo2w  size=20.97KiB
INFO [08-13|17:08:03.166] Looking for peers                        peercount=1 tried=31 static=0
```

data目录结构
```shell
├── data
│   ├── geth
│   │   ├── chaindata
│   │   │   ├── 000001.log
│   │   │   ├── ancient
│   │   │   │   ├── bodies.0000.cdat
│   │   │   │   ├── bodies.cidx
│   │   │   │   ├── bodies.meta
│   │   │   │   ├── diffs.0000.rdat
│   │   │   │   ├── diffs.meta
│   │   │   │   ├── diffs.ridx
│   │   │   │   ├── FLOCK
│   │   │   │   ├── hashes.0000.rdat
│   │   │   │   ├── hashes.meta
│   │   │   │   ├── hashes.ridx
│   │   │   │   ├── headers.0000.cdat
│   │   │   │   ├── headers.cidx
│   │   │   │   ├── headers.meta
│   │   │   │   ├── receipts.0000.cdat
│   │   │   │   ├── receipts.cidx
│   │   │   │   └── receipts.meta
│   │   │   ├── CURRENT
│   │   │   ├── LOCK
│   │   │   ├── LOG
│   │   │   └── MANIFEST-000000
│   │   ├── ethash
│   │   │   ├── cache-R23-0000000000000000
│   │   │   └── cache-R23-290decd9548b62a8
│   │   ├── jwtsecret
│   │   ├── LOCK
│   │   ├── nodekey
│   │   ├── nodes
│   │   │   ├── 000001.log
│   │   │   ├── CURRENT
│   │   │   ├── LOCK
│   │   │   ├── LOG
│   │   │   └── MANIFEST-000000
│   │   ├── transactions.rlp
│   │   └── triecache
│   │       ├── data.0.bin
│   │       ├── data.1.bin
│   │       └── metadata.bin
│   └── keystore
```

#### 查看账户`eth.accounts`
```shell
> eth.accounts
[]
```

#### 创建账户`personal.newAccount`

```shell
# 创建账户, 123456是密码
> personal.newAccount("123456")
INFO [08-13|17:34:41.328] Looking for peers                        peercount=0 tried=6  static=0
INFO [08-13|17:34:41.377] Your new key was generated               address=0x160781b56FF6E366e8B41F1ef910Eb3c90B46D13
WARN [08-13|17:34:41.377] Please backup your key file!             path=/root/work/ethereum/data/keystore/UTC--2022-08-13T09-34-39.377406787Z--160781b56ff6e366e8b41f1ef910eb3c90b46d13
WARN [08-13|17:34:41.377] Please remember your password! 
"0x160781b56ff6e366e8b41f1ef910eb3c90b46d13"
```

keyfile
```json
{
	"address": "160781b56ff6e366e8b41f1ef910eb3c90b46d13",
	"crypto": {
		"cipher": "aes-128-ctr",
		"ciphertext": "c420c847549b683a60f12558dd8a16466479fe7a6aee51973896bf12d00d7796",
		"cipherparams": {
			"iv": "3ef50abb93c294729bd3a36e28d08474"
		},
		"kdf": "scrypt",
		"kdfparams": {
			"dklen": 32,
			"n": 262144,
			"p": 1,
			"r": 8,
			"salt": "126d5d3a2d7fa1cd6678d90b141121cdd2c099c5a8cdc0899945d601e583003f"
		},
		"mac": "c3a1e4eae1910bb024ab722b735cee3058f6e97e975734433e631fd5ad696636"
	},
	"id": "1a2af59f-0857-4159-a573-72cfdb6c49fc",
	"version": 3
}
```

最后是公钥`0x160781b56ff6e366e8b41f1ef910eb3c90b46d13`  

#### 查看账户信息`eth.getBalance`
```shell
> eth.getBalance("0x160781b56ff6e366e8b41f1ef910eb3c90b46d13")
0
```


### 命令行操作 
#### 查看区块个数`eth.blockNumber`
```shell
> eth.blockNumber
0
```

#### 挖矿`miner.start`
```shell
> miner.start()

# 过一会手动停止
> miner.stop()
```

打印信息
```shell
INFO [08-14|17:52:12.096] Commit new sealing work                  number=1 sealhash=2ebac1..19cd2e uncles=0 txs=0 gas=0 fees=0 elapsed="542.164µs"
INFO [08-14|17:52:12.096] Commit new sealing work                  number=1 sealhash=2ebac1..19cd2e uncles=0 txs=0 gas=0 fees=0 elapsed="829.805µs"
INFO [08-14|17:52:12.749] Successfully sealed new block            number=1 sealhash=2ebac1..19cd2e hash=08637f..dea52c elapsed=652.847ms
INFO [08-14|17:52:12.749] 🔨 mined potential block                  number=1 hash=08637f..dea52c
INFO [08-14|17:52:12.749] Commit new sealing work                  number=2 sealhash=fddd08..6e292f uncles=0 txs=0 gas=0 fees=0 elapsed="353.867µs"
INFO [08-14|17:52:12.750] Commit new sealing work                  number=2 sealhash=fddd08..6e292f uncles=0 txs=0 gas=0 fees=0 elapsed="945.385µs"
INFO [08-14|17:52:13.127] Successfully sealed new block            number=2 sealhash=fddd08..6e292f hash=ed4b07..d61838 elapsed=378.047ms
INFO [08-14|17:52:13.128] 🔨 mined potential block                  number=2 hash=ed4b07..d61838
INFO [08-14|17:52:13.129] Commit new sealing work                  number=3 sealhash=bc06ea..75efbe uncles=0 txs=0 gas=0 fees=0 elapsed="446.811µs"
INFO [08-14|17:52:13.130] Commit new sealing work                  number=3 sealhash=bc06ea..75efbe uncles=0 txs=0 gas=0 fees=0 elapsed=1.085ms
INFO [08-14|17:52:15.336] Successfully sealed new block            number=3 sealhash=bc06ea..75efbe hash=67c33d..a45c84 elapsed=2.207s
INFO [08-14|17:52:15.336] 🔨 mined potential block                  number=3 hash=67c33d..a45c84
INFO [08-14|17:52:15.337] Commit new sealing work                  number=4 sealhash=1cd772..ac1ff2 uncles=0 txs=0 gas=0 fees=0 elapsed="297.746µs"
INFO [08-14|17:52:15.337] Commit new sealing work                  number=4 sealhash=1cd772..ac1ff2 uncles=0 txs=0 gas=0 fees=0 elapsed="468.251µs"
INFO [08-14|17:52:17.583] Successfully sealed new block            number=4 sealhash=1cd772..ac1ff2 hash=652926..0d449e elapsed=2.245s
INFO [08-14|17:52:17.583] 🔨 mined potential block                  number=4 hash=652926..0d449e
INFO [08-14|17:52:17.583] Commit new sealing work                  number=5 sealhash=d54b1e..5a9eeb uncles=0 txs=0 gas=0 fees=0 elapsed="250.564µs"
INFO [08-14|17:52:17.584] Commit new sealing work                  number=5 sealhash=d54b1e..5a9eeb uncles=0 txs=0 gas=0 fees=0 elapsed="709.186µs"
INFO [08-14|17:52:18.580] Looking for peers                        peercount=2 tried=130 static=0
INFO [08-14|17:52:20.664] Successfully sealed new block            number=5 sealhash=d54b1e..5a9eeb hash=b2534e..41cca9 elapsed=3.080s
INFO [08-14|17:52:20.664] 🔨 mined potential block                  number=5 hash=b2534e..41cca9
INFO [08-14|17:52:20.664] Commit new sealing work                  number=6 sealhash=041235..662389 uncles=0 txs=0 gas=0 fees=0 elapsed="206.314µs"
INFO [08-14|17:52:20.665] Commit new sealing work                  number=6 sealhash=041235..662389 uncles=0 txs=0 gas=0 fees=0 elapsed="501.492µs"
INFO [08-14|17:52:20.712] Successfully sealed new block            number=6 sealhash=041235..662389 hash=081f4c..980c6c elapsed=48.142ms
INFO [08-14|17:52:20.712] 🔨 mined potential block                  number=6 hash=081f4c..980c6c
INFO [08-14|17:52:20.713] Commit new sealing work                  number=7 sealhash=64135e..91e586 uncles=0 txs=0 gas=0 fees=0 elapsed="205.802µs"
INFO [08-14|17:52:20.713] Commit new sealing work                  number=7 sealhash=64135e..91e586 uncles=0 txs=0 gas=0 fees=0 elapsed="482.088µs"
INFO [08-14|17:52:25.895] Successfully sealed new block            number=7 sealhash=64135e..91e586 hash=087452..42e15f elapsed=5.181s
INFO [08-14|17:52:25.895] 🔨 mined potential block                  number=7 hash=087452..42e15f
INFO [08-14|17:52:25.895] Commit new sealing work                  number=8 sealhash=b1c024..a9da80 uncles=0 txs=0 gas=0 fees=0 elapsed="188.795µs"
INFO [08-14|17:52:25.896] Commit new sealing work                  number=8 sealhash=b1c024..a9da80 uncles=0 txs=0 gas=0 fees=0 elapsed="834.275µs"
INFO [08-14|17:52:26.988] Successfully sealed new block            number=8 sealhash=b1c024..a9da80 hash=ffb84d..794d30 elapsed=1.092s
INFO [08-14|17:52:26.988] 🔗 block reached canonical chain          number=1 hash=08637f..dea52c
INFO [08-14|17:52:26.988] 🔨 mined potential block                  number=8 hash=ffb84d..794d30
```

查看区块信息:
```shell
> eth.blockNumber
8
> eth.accounts
["0xa32662696e6d7c3a3348141daaf642605c7e8fca"]
> eth.getBalance(eth.accounts[0])
16000000000000000000
# 转换成ether
> web3.fromWei(16000000000000000000,'ether')
"16"
```

####  `admin.peers`
```shell
> admin.peers
[{
    caps: ["diff/1", "eth/65", "eth/66", "eth/67", "snap/1"],
    enode: "enode://eb57103f1ddd58cb13d40d5b45b6c05a91e6748b195c9886c390e123fd1643ace03a4b0ce04bfc1cdb9a997ff670c0980a8005aa0449a2b784085d9425857dd0@54.155.229.0:30311",
    id: "3cd20ccf82586affc7ec52652a7bd2e659c2001567a18a0103792992af075497",
    name: "Geth/v1.1.11-70d08a57/linux-amd64/go1.16.10",
    network: {
      inbound: false,
      localAddress: "10.211.55.10:44354",
      remoteAddress: "54.155.229.0:30311",
      static: false,
      trusted: false
    },
    protocols: {
      eth: "handshake",
      snap: "handshake"
    }
}]
> admin.peers
[]
```

#### 多个节点之间交易  
创建另个一节点:  
```shell
geth --datadir data2 init genesis.json
```

输出日志
```shell
root@matrix:~/work/ethereum# geth --datadir data2 init genesis.json
INFO [08-14|18:52:06.151] Maximum peer count                       ETH=50 LES=0 total=50
INFO [08-14|18:52:06.153] Smartcard socket not found, disabling    err="stat /run/pcscd/pcscd.comm: no such file or directory"
INFO [08-14|18:52:06.156] Set global gas cap                       cap=50,000,000
INFO [08-14|18:52:06.158] Allocated cache and file handles         database=/root/work/ethereum/data2/geth/chaindata cache=16.00MiB handles=16
INFO [08-14|18:52:06.198] Opened ancient database                  database=/root/work/ethereum/data2/geth/chaindata/ancient readonly=false
INFO [08-14|18:52:06.198] Writing custom genesis block 
INFO [08-14|18:52:06.199] Persisted trie from memory database      nodes=3 size=399.00B time="59.575µs" gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
INFO [08-14|18:52:06.200] Successfully wrote genesis state         database=chaindata                                hash=6ab19c..97db44
INFO [08-14|18:52:06.200] Allocated cache and file handles         database=/root/work/ethereum/data2/geth/lightchaindata cache=16.00MiB handles=16
INFO [08-14|18:52:06.226] Opened ancient database                  database=/root/work/ethereum/data2/geth/lightchaindata/ancient readonly=false
INFO [08-14|18:52:06.226] Writing custom genesis block 
INFO [08-14|18:52:06.227] Persisted trie from memory database      nodes=3 size=399.00B time="72.35µs"  gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
INFO [08-14|18:52:06.227] Successfully wrote genesis state         database=lightchaindata                           hash=6ab19c..97db44
```

启动第二个节点
```shell
geth --datadir data2/ --networkid 123 --port 3101 --nodiscover --authrpc.port 8552
```

启动日志
```shell
INFO [08-14|19:01:23.168] Disk storage enabled for ethash caches   dir=/root/work/ethereum/data2/geth/ethash count=3
INFO [08-14|19:01:23.168] Disk storage enabled for ethash DAGs     dir=/root/.ethash                         count=2
INFO [08-14|19:01:23.168] Initialising Ethereum protocol           network=123 dbversion=8
INFO [08-14|19:01:23.169] Loaded most recent local header          number=0 hash=6ab19c..97db44 td=131,072 age=53y4mo2w
INFO [08-14|19:01:23.169] Loaded most recent local full block      number=0 hash=6ab19c..97db44 td=131,072 age=53y4mo2w
INFO [08-14|19:01:23.169] Loaded most recent local fast block      number=0 hash=6ab19c..97db44 td=131,072 age=53y4mo2w
INFO [08-14|19:01:23.170] Loaded local transaction journal         transactions=0 dropped=0
INFO [08-14|19:01:23.170] Regenerated local transaction journal    transactions=0 accounts=0
INFO [08-14|19:01:23.170] Gasprice oracle is ignoring threshold set threshold=2
WARN [08-14|19:01:23.170] Unclean shutdown detected                booted=2022-08-14T18:53:52+0800 age=7m31s
WARN [08-14|19:01:23.170] Unclean shutdown detected                booted=2022-08-14T18:55:35+0800 age=5m48s
WARN [08-14|19:01:23.170] Unclean shutdown detected                booted=2022-08-14T18:55:52+0800 age=5m31s
WARN [08-14|19:01:23.170] Unclean shutdown detected                booted=2022-08-14T18:57:22+0800 age=4m1s
WARN [08-14|19:01:23.170] Engine API enabled                       protocol=eth
WARN [08-14|19:01:23.170] Engine API started but chain not configured for merge yet 
INFO [08-14|19:01:23.170] Starting peer-to-peer node               instance=Geth/v1.10.21-stable-67109427/linux-amd64/go1.18.4
INFO [08-14|19:01:23.188] New local node record                    seq=1,660,474,432,570 id=088c3bc19542e630 ip=127.0.0.1 udp=0 tcp=3101
INFO [08-14|19:01:23.188] Started P2P networking                   self="enode://febe674e1415e8c2ae949ba06f4dad43da28905161e1e48c636de97c473605b6cb939f9dadaf0f61fe587ae9bce16115df6ae41601fc23f09f9deb340748027e@127.0.0.1:3101?discport=0"
INFO [08-14|19:01:23.189] IPC endpoint opened                      url=/root/work/ethereum/data2/geth.ipc
INFO [08-14|19:01:23.189] Loaded JWT secret file                   path=/root/work/ethereum/data2/geth/jwtsecret crc32=0x5ed23a5d
INFO [08-14|19:01:23.191] WebSocket enabled                        url=ws://127.0.0.1:8552
INFO [08-14|19:01:23.191] HTTP server started                      endpoint=127.0.0.1:8552 auth=true prefix= cors=localhost vhosts=localhost
INFO [08-14|19:01:25.340] Mapped network port                      proto=tcp extport=3101 intport=3101 interface="UPNP IGDv1-IP1"
```

连接节点2
```shell
geth attach ipc:data2/geth.ipc
```

创建两个账户:
```shell
> personal.newAccount("1234567")
"0x6616e2e7e372c648830a472bd9816bb371f1be6e"
> personal.newAccount("12345678")
"0x9e0f0703def69563f994a8120c8fc8909f6337ee"
> admin.nodeInfo
{
  enode: "enode://febe674e1415e8c2ae949ba06f4dad43da28905161e1e48c636de97c473605b6cb939f9dadaf0f61fe587ae9bce16115df6ae41601fc23f09f9deb340748027e@192.168.1.4:3101?discport=0",
  enr: "enr:-Jy4QLSBwlEdthgZPAU4xAEw_X5cTHOvERkfNxh0EIcuvIx2JUAso79ZZso5EikoG4iuBkUrdXkloRPyUbPoHJKPHaKGAYKb_Rw7g2V0aMfGhG3BCFmAgmlkgnY0gmlwhMCoAQSJc2VjcDI1NmsxoQL-vmdOFBXowq6Um6BvTa1D2iiQUWHh5Ixjbel8RzYFtoRzbmFwwIN0Y3CCDB0",
  id: "088c3bc19542e630eeb4997c5f8478453e9f4a778d1fba0aa851470a6a73beef",
  ip: "192.168.1.4",
  listenAddr: "[::]:3101",
  name: "Geth/v1.10.21-stable-67109427/linux-amd64/go1.18.4",
  ports: {
    discovery: 0,
    listener: 3101
  },
  protocols: {
    eth: {
      config: {
        berlinBlock: 0,
        byzantiumBlock: 0,
        chainId: 123,
        constantinopleBlock: 0,
        eip150Block: 0,
        eip150Hash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        eip155Block: 0,
        eip158Block: 0,
        homesteadBlock: 0,
        istanbulBlock: 0,
        londonBlock: 0,
        petersburgBlock: 0
      },
      difficulty: 131072,
      genesis: "0x6ab19c9fed11060a0bd11534a277e4e60e8c831c6953b196fc95c867bc97db44",
      head: "0x6ab19c9fed11060a0bd11534a277e4e60e8c831c6953b196fc95c867bc97db44",
      network: 123
    },
    snap: {}
  }
}
```

可以看到data2的节点信息为:`enode://febe674e1415e8c2ae949ba06f4dad43da28905161e1e48c636de97c473605b6cb939f9dadaf0f61fe587ae9bce16115df6ae41601fc23f09f9deb340748027e@192.168.1.4:3101?discport=0`  

可以到`data`中添加`data2`节点  

```shell
> admin.addPeer("enode://febe674e1415e8c2ae949ba06f4dad43da28905161e1e48c636de97c473605b6cb939f9dadaf0f61fe587ae9bce16115df6ae41601fc23f09f9deb340748027e@192.168.1.4:3101?discport=0")
true
```

> `data2`虽然不能被发现，但是可以被手动添加  

添加成功后，但是`admin.peers`还是为空，增加日志等级`--verbosity 5`,查看错误信息 

```shell
DEBUG[08-14|19:15:52.491] Served admin_addPeer                     reqid=33 duration="214.025µs"
TRACE[08-14|19:15:52.491] Adding static node                       id=088c3bc19542e630 ip=192.168.1.4 added=false
TRACE[08-14|19:15:59.562] Dial error                               id=088c3bc19542e630 addr=192.168.1.4:3101 conn=staticdial err="i/o timeout"
```

> 网上说没有进行`geth init`，直接运行导致。重试`init`也不行  


#### 交易`eth.sendTransaction`  

```shell
eth.sendTransaction({
    from: '0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe',
    to: '0x11f4d0A3c12e86B4b5F39B213F7E19D048276DAe',
    value: '1000000000000000'
})
```







