# solidity
[github源码](https://github.com/ethereum/solidity)  

Solidity, the Smart Contract Programming Language  

- ### [FISCO BCOS 技术文档](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/index.html)  

## 源码编译
[官方文档](https://docs.soliditylang.org/en/latest/installing-solidity.html#building-from-source) 

环境准备:  
- CMake (version 3.13+)  
- Boost (version 1.77+ on Windows, 1.65+ otherwise)  
- Git
- [z3](https://github.com/Z3Prover/z3)(version 4.8+, Optional)
- cvc4 (Optional) 


环境安装:
boost
```
wget https://boostorg.jfrog.io/artifactory/main/release/1.69.0/source/boost_1_69_0.tar.gz
tar -zxvf boost_1_69_0.tar.gz  

#编译
# 进入源码目录
./bootstrap.sh
sudo ./b2 install  

# 查看库
ls -l /usr/local/lib/libboost_*  
ls -l /usr/local/include/boost/
```

z3
```shell
git clone https://github.com/Z3Prover/z3

# 编译
python scripts/mk_make.py
cd build
make
sudo make install
```

> 需要安装`4.8.17`指定版本，版本过高也无法编译。  

solidity编译
```
git clone git@github.com:ethereum/solidity.git  # 需要有.git目录

mkdir build && cd build
cmake .. && make
make install 
```

编译后的文件
```shell
CMakeCache.txt
CMakeFiles
cmake_install.cmake
_deps
deps
include
libevmasm
liblangutil
libsmtutil
libsolc
libsolidity
libsolutil
libyul
Makefile
solc
test
tools
```

还有一些编译选项
```shell
# disables only Z3 SMT Solver.
cmake .. -DUSE_Z3=OFF

# disables only CVC4 SMT Solver.
cmake .. -DUSE_CVC4=OFF

# disables both Z3 and CVC4
cmake .. -DUSE_CVC4=OFF -DUSE_Z3=OFF
```

## vscode 调试  

- ### [编译文档](https://docs.soliditylang.org/en/latest/using-the-compiler.html)  

solc的用法示例
```shell
solc --bin -o /tmp/solcoutput dapp-bin=/usr/local/lib/dapp-bin contract.sol
```

测试合约:  
```js
// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;
contract ReceiverPays {
    address owner = msg.sender;

    mapping(uint256 => bool) usedNonces;

    constructor() payable {}

    function claimPayment(uint256 amount, uint256 nonce, bytes memory signature) external {
        require(!usedNonces[nonce]);
        usedNonces[nonce] = true;

        // this recreates the message that was signed on the client
        bytes32 message = prefixed(keccak256(abi.encodePacked(msg.sender, amount, nonce, this)));

        require(recoverSigner(message, signature) == owner);

        payable(msg.sender).transfer(amount);
    }

    /// destroy the contract and reclaim the leftover funds.
    function shutdown() external {
        require(msg.sender == owner);
        selfdestruct(payable(msg.sender));
    }

    /// signature methods.
    function splitSignature(bytes memory sig)
        internal
        pure
        returns (uint8 v, bytes32 r, bytes32 s)
    {
        require(sig.length == 65);

        assembly {
            // first 32 bytes, after the length prefix.
            r := mload(add(sig, 32))
            // second 32 bytes.
            s := mload(add(sig, 64))
            // final byte (first byte of the next 32 bytes).
            v := byte(0, mload(add(sig, 96)))
        }

        return (v, r, s);
    }

    function recoverSigner(bytes32 message, bytes memory sig)
        internal
        pure
        returns (address)
    {
        (uint8 v, bytes32 r, bytes32 s) = splitSignature(sig);

        return ecrecover(message, v, r, s);
    }

    /// builds a prefixed hash to mimic the behavior of eth_sign.
    function prefixed(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
}
```

```shell
./solc/solc --bin -o test.bin  test.sol

# 输出 Bin文件  
.
├── test.bin
│   └── ReceiverPays.bin       # 编译后的bin文件, 二进制文件  
└── test.sol
```  

使用指令输出指令集`./../solc/solc --bin -o test.bin --ast-compact-json --asm  test.sol`  
```shell
├── test.bin
│   ├── ReceiverPays.bin
│   ├── ReceiverPays.evm       # 
│   └── test.sol_json.ast      # 抽象语法树
└── test.sol
```

>   --asm                EVM assembly of the contracts.
>   --asm-json           EVM assembly of the contracts in JSON format.  



`ReceiverPays.evm`文件详情:  
```asm
 /* "test.sol":68:1820  contract ReceiverPays {... */
  mstore(0x40, 0x80)
    /* "test.sol":112:122  msg.sender */
  caller
    /* "test.sol":96:122  address owner = msg.sender */
  0x00
  dup1
  0x0100
  exp
  dup2
  sload
  dup2
  0xffffffffffffffffffffffffffffffffffffffff
  mul
  not
  and
  swap1
  dup4
  0xffffffffffffffffffffffffffffffffffffffff
  and
  mul
  or
  swap1
  sstore
  pop
    /* "test.sol":68:1820  contract ReceiverPays {... */
  dataSize(sub_0)
  dup1
  dataOffset(sub_0)
  0x00
  codecopy
  0x00
  return
stop

sub_0: assembly {
        /* "test.sol":68:1820  contract ReceiverPays {... */
      mstore(0x40, 0x80)
      callvalue
      dup1
      iszero
      tag_1
      jumpi
      0x00
      dup1
      revert
...
```

remix 复制的字节码
```json
{
    "functionDebugData": {
        "@_13": {
            "entryPoint": null, 
            "id": 13, 
            "parameterSlots": 0, 
            "returnSlots": 0
        }
    }, 
    "generatedSources": [ ], 
    "linkReferences": { }, 
    "object": "6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061077d806100536000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063a90ae8871461003b578063fc0e74d114610057575b600080fd5b61005560048036038101906100509190610470565b610061565b005b61005f61019e565b005b6001600083815260200190815260200160002060009054906101000a900460ff161561008c57600080fd5b600180600084815260200190815260200160002060006101000a81548160ff02191690831515021790555060006100ee338585306040516020016100d394939291906105e1565b6040516020818303038152906040528051906020012061020f565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610131828461023f565b73ffffffffffffffffffffffffffffffffffffffff161461015157600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc859081150290604051600060405180830381858888f19350505050158015610197573d6000803e3d6000fd5b5050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101f657600080fd5b3373ffffffffffffffffffffffffffffffffffffffff16ff5b60008160405160200161022291906106b1565b604051602081830303815290604052805190602001209050919050565b60008060008061024e856102ae565b925092509250600186848484604051600081526020016040526040516102779493929190610702565b6020604051602081039080840390855afa158015610299573d6000803e3d6000fd5b50505060206040510351935050505092915050565b600080600060418451146102c157600080fd5b6020840151915060408401519050606084015160001a92509193909250565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610307816102f4565b811461031257600080fd5b50565b600081359050610324816102fe565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61037d82610334565b810181811067ffffffffffffffff8211171561039c5761039b610345565b5b80604052505050565b60006103af6102e0565b90506103bb8282610374565b919050565b600067ffffffffffffffff8211156103db576103da610345565b5b6103e482610334565b9050602081019050919050565b82818337600083830152505050565b600061041361040e846103c0565b6103a5565b90508281526020810184848401111561042f5761042e61032f565b5b61043a8482856103f1565b509392505050565b600082601f8301126104575761045661032a565b5b8135610467848260208601610400565b91505092915050565b600080600060608486031215610489576104886102ea565b5b600061049786828701610315565b93505060206104a886828701610315565b925050604084013567ffffffffffffffff8111156104c9576104c86102ef565b5b6104d586828701610442565b9150509250925092565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061050a826104df565b9050919050565b60008160601b9050919050565b600061052982610511565b9050919050565b600061053b8261051e565b9050919050565b61055361054e826104ff565b610530565b82525050565b6000819050919050565b61057461056f826102f4565b610559565b82525050565b6000819050919050565b600061059f61059a610595846104df565b61057a565b6104df565b9050919050565b60006105b182610584565b9050919050565b60006105c3826105a6565b9050919050565b6105db6105d6826105b8565b610530565b82525050565b60006105ed8287610542565b6014820191506105fd8286610563565b60208201915061060d8285610563565b60208201915061061d82846105ca565b60148201915081905095945050505050565b600081905092915050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b6000610670601c8361062f565b915061067b8261063a565b601c82019050919050565b6000819050919050565b6000819050919050565b6106ab6106a682610686565b610690565b82525050565b60006106bc82610663565b91506106c8828461069a565b60208201915081905092915050565b6106e081610686565b82525050565b600060ff82169050919050565b6106fc816106e6565b82525050565b600060808201905061071760008301876106d7565b61072460208301866106f3565b61073160408301856106d7565b61073e60608301846106d7565b9594505050505056fea2646970667358221220aaf19d2292322ba117a13c1d5a3ab83fee819c4545b74c206cdc9dab1c53b55a64736f6c634300080f0033", 
    "opcodes": "PUSH1 0x80 PUSH1 0x40 MSTORE CALLER PUSH1 0x0 DUP1 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP PUSH2 0x77D DUP1 PUSH2 0x53 PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN INVALID PUSH1 0x80 PUSH1 0x40 MSTORE CALLVALUE DUP1 ISZERO PUSH2 0x10 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x4 CALLDATASIZE LT PUSH2 0x36 JUMPI PUSH1 0x0 CALLDATALOAD PUSH1 0xE0 SHR DUP1 PUSH4 0xA90AE887 EQ PUSH2 0x3B JUMPI DUP1 PUSH4 0xFC0E74D1 EQ PUSH2 0x57 JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x55 PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 PUSH2 0x50 SWAP2 SWAP1 PUSH2 0x470 JUMP JUMPDEST PUSH2 0x61 JUMP JUMPDEST STOP JUMPDEST PUSH2 0x5F PUSH2 0x19E JUMP JUMPDEST STOP JUMPDEST PUSH1 0x1 PUSH1 0x0 DUP4 DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND ISZERO PUSH2 0x8C JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x1 DUP1 PUSH1 0x0 DUP5 DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH1 0xFF MUL NOT AND SWAP1 DUP4 ISZERO ISZERO MUL OR SWAP1 SSTORE POP PUSH1 0x0 PUSH2 0xEE CALLER DUP6 DUP6 ADDRESS PUSH1 0x40 MLOAD PUSH1 0x20 ADD PUSH2 0xD3 SWAP5 SWAP4 SWAP3 SWAP2 SWAP1 PUSH2 0x5E1 JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH1 0x20 DUP2 DUP4 SUB SUB DUP2 MSTORE SWAP1 PUSH1 0x40 MSTORE DUP1 MLOAD SWAP1 PUSH1 0x20 ADD KECCAK256 PUSH2 0x20F JUMP JUMPDEST SWAP1 POP PUSH1 0x0 DUP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH2 0x131 DUP3 DUP5 PUSH2 0x23F JUMP JUMPDEST PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ PUSH2 0x151 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH2 0x8FC DUP6 SWAP1 DUP2 ISZERO MUL SWAP1 PUSH1 0x40 MLOAD PUSH1 0x0 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 DUP6 DUP9 DUP9 CALL SWAP4 POP POP POP POP ISZERO DUP1 ISZERO PUSH2 0x197 JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP POP JUMP JUMPDEST PUSH1 0x0 DUP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ PUSH2 0x1F6 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SELFDESTRUCT JUMPDEST PUSH1 0x0 DUP2 PUSH1 0x40 MLOAD PUSH1 0x20 ADD PUSH2 0x222 SWAP2 SWAP1 PUSH2 0x6B1 JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH1 0x20 DUP2 DUP4 SUB SUB DUP2 MSTORE SWAP1 PUSH1 0x40 MSTORE DUP1 MLOAD SWAP1 PUSH1 0x20 ADD KECCAK256 SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x0 DUP1 PUSH2 0x24E DUP6 PUSH2 0x2AE JUMP JUMPDEST SWAP3 POP SWAP3 POP SWAP3 POP PUSH1 0x1 DUP7 DUP5 DUP5 DUP5 PUSH1 0x40 MLOAD PUSH1 0x0 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x40 MSTORE PUSH1 0x40 MLOAD PUSH2 0x277 SWAP5 SWAP4 SWAP3 SWAP2 SWAP1 PUSH2 0x702 JUMP JUMPDEST PUSH1 0x20 PUSH1 0x40 MLOAD PUSH1 0x20 DUP2 SUB SWAP1 DUP1 DUP5 SUB SWAP1 DUP6 GAS STATICCALL ISZERO DUP1 ISZERO PUSH2 0x299 JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP PUSH1 0x20 PUSH1 0x40 MLOAD SUB MLOAD SWAP4 POP POP POP POP SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x0 PUSH1 0x41 DUP5 MLOAD EQ PUSH2 0x2C1 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x20 DUP5 ADD MLOAD SWAP2 POP PUSH1 0x40 DUP5 ADD MLOAD SWAP1 POP PUSH1 0x60 DUP5 ADD MLOAD PUSH1 0x0 BYTE SWAP3 POP SWAP2 SWAP4 SWAP1 SWAP3 POP JUMP JUMPDEST PUSH1 0x0 PUSH1 0x40 MLOAD SWAP1 POP SWAP1 JUMP JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 DUP2 SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH2 0x307 DUP2 PUSH2 0x2F4 JUMP JUMPDEST DUP2 EQ PUSH2 0x312 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP JUMP JUMPDEST PUSH1 0x0 DUP2 CALLDATALOAD SWAP1 POP PUSH2 0x324 DUP2 PUSH2 0x2FE JUMP JUMPDEST SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 PUSH1 0x1F NOT PUSH1 0x1F DUP4 ADD AND SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH32 0x4E487B7100000000000000000000000000000000000000000000000000000000 PUSH1 0x0 MSTORE PUSH1 0x41 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST PUSH2 0x37D DUP3 PUSH2 0x334 JUMP JUMPDEST DUP2 ADD DUP2 DUP2 LT PUSH8 0xFFFFFFFFFFFFFFFF DUP3 GT OR ISZERO PUSH2 0x39C JUMPI PUSH2 0x39B PUSH2 0x345 JUMP JUMPDEST JUMPDEST DUP1 PUSH1 0x40 MSTORE POP POP POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x3AF PUSH2 0x2E0 JUMP JUMPDEST SWAP1 POP PUSH2 0x3BB DUP3 DUP3 PUSH2 0x374 JUMP JUMPDEST SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 PUSH8 0xFFFFFFFFFFFFFFFF DUP3 GT ISZERO PUSH2 0x3DB JUMPI PUSH2 0x3DA PUSH2 0x345 JUMP JUMPDEST JUMPDEST PUSH2 0x3E4 DUP3 PUSH2 0x334 JUMP JUMPDEST SWAP1 POP PUSH1 0x20 DUP2 ADD SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST DUP3 DUP2 DUP4 CALLDATACOPY PUSH1 0x0 DUP4 DUP4 ADD MSTORE POP POP POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x413 PUSH2 0x40E DUP5 PUSH2 0x3C0 JUMP JUMPDEST PUSH2 0x3A5 JUMP JUMPDEST SWAP1 POP DUP3 DUP2 MSTORE PUSH1 0x20 DUP2 ADD DUP5 DUP5 DUP5 ADD GT ISZERO PUSH2 0x42F JUMPI PUSH2 0x42E PUSH2 0x32F JUMP JUMPDEST JUMPDEST PUSH2 0x43A DUP5 DUP3 DUP6 PUSH2 0x3F1 JUMP JUMPDEST POP SWAP4 SWAP3 POP POP POP JUMP JUMPDEST PUSH1 0x0 DUP3 PUSH1 0x1F DUP4 ADD SLT PUSH2 0x457 JUMPI PUSH2 0x456 PUSH2 0x32A JUMP JUMPDEST JUMPDEST DUP2 CALLDATALOAD PUSH2 0x467 DUP5 DUP3 PUSH1 0x20 DUP7 ADD PUSH2 0x400 JUMP JUMPDEST SWAP2 POP POP SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x0 PUSH1 0x60 DUP5 DUP7 SUB SLT ISZERO PUSH2 0x489 JUMPI PUSH2 0x488 PUSH2 0x2EA JUMP JUMPDEST JUMPDEST PUSH1 0x0 PUSH2 0x497 DUP7 DUP3 DUP8 ADD PUSH2 0x315 JUMP JUMPDEST SWAP4 POP POP PUSH1 0x20 PUSH2 0x4A8 DUP7 DUP3 DUP8 ADD PUSH2 0x315 JUMP JUMPDEST SWAP3 POP POP PUSH1 0x40 DUP5 ADD CALLDATALOAD PUSH8 0xFFFFFFFFFFFFFFFF DUP2 GT ISZERO PUSH2 0x4C9 JUMPI PUSH2 0x4C8 PUSH2 0x2EF JUMP JUMPDEST JUMPDEST PUSH2 0x4D5 DUP7 DUP3 DUP8 ADD PUSH2 0x442 JUMP JUMPDEST SWAP2 POP POP SWAP3 POP SWAP3 POP SWAP3 JUMP JUMPDEST PUSH1 0x0 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF DUP3 AND SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x50A DUP3 PUSH2 0x4DF JUMP JUMPDEST SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP2 PUSH1 0x60 SHL SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x529 DUP3 PUSH2 0x511 JUMP JUMPDEST SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x53B DUP3 PUSH2 0x51E JUMP JUMPDEST SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH2 0x553 PUSH2 0x54E DUP3 PUSH2 0x4FF JUMP JUMPDEST PUSH2 0x530 JUMP JUMPDEST DUP3 MSTORE POP POP JUMP JUMPDEST PUSH1 0x0 DUP2 SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH2 0x574 PUSH2 0x56F DUP3 PUSH2 0x2F4 JUMP JUMPDEST PUSH2 0x559 JUMP JUMPDEST DUP3 MSTORE POP POP JUMP JUMPDEST PUSH1 0x0 DUP2 SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x59F PUSH2 0x59A PUSH2 0x595 DUP5 PUSH2 0x4DF JUMP JUMPDEST PUSH2 0x57A JUMP JUMPDEST PUSH2 0x4DF JUMP JUMPDEST SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x5B1 DUP3 PUSH2 0x584 JUMP JUMPDEST SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x5C3 DUP3 PUSH2 0x5A6 JUMP JUMPDEST SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH2 0x5DB PUSH2 0x5D6 DUP3 PUSH2 0x5B8 JUMP JUMPDEST PUSH2 0x530 JUMP JUMPDEST DUP3 MSTORE POP POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x5ED DUP3 DUP8 PUSH2 0x542 JUMP JUMPDEST PUSH1 0x14 DUP3 ADD SWAP2 POP PUSH2 0x5FD DUP3 DUP7 PUSH2 0x563 JUMP JUMPDEST PUSH1 0x20 DUP3 ADD SWAP2 POP PUSH2 0x60D DUP3 DUP6 PUSH2 0x563 JUMP JUMPDEST PUSH1 0x20 DUP3 ADD SWAP2 POP PUSH2 0x61D DUP3 DUP5 PUSH2 0x5CA JUMP JUMPDEST PUSH1 0x14 DUP3 ADD SWAP2 POP DUP2 SWAP1 POP SWAP6 SWAP5 POP POP POP POP POP JUMP JUMPDEST PUSH1 0x0 DUP2 SWAP1 POP SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH32 0x19457468657265756D205369676E6564204D6573736167653A0A333200000000 PUSH1 0x0 DUP3 ADD MSTORE POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x670 PUSH1 0x1C DUP4 PUSH2 0x62F JUMP JUMPDEST SWAP2 POP PUSH2 0x67B DUP3 PUSH2 0x63A JUMP JUMPDEST PUSH1 0x1C DUP3 ADD SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP2 SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP2 SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH2 0x6AB PUSH2 0x6A6 DUP3 PUSH2 0x686 JUMP JUMPDEST PUSH2 0x690 JUMP JUMPDEST DUP3 MSTORE POP POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x6BC DUP3 PUSH2 0x663 JUMP JUMPDEST SWAP2 POP PUSH2 0x6C8 DUP3 DUP5 PUSH2 0x69A JUMP JUMPDEST PUSH1 0x20 DUP3 ADD SWAP2 POP DUP2 SWAP1 POP SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH2 0x6E0 DUP2 PUSH2 0x686 JUMP JUMPDEST DUP3 MSTORE POP POP JUMP JUMPDEST PUSH1 0x0 PUSH1 0xFF DUP3 AND SWAP1 POP SWAP2 SWAP1 POP JUMP JUMPDEST PUSH2 0x6FC DUP2 PUSH2 0x6E6 JUMP JUMPDEST DUP3 MSTORE POP POP JUMP JUMPDEST PUSH1 0x0 PUSH1 0x80 DUP3 ADD SWAP1 POP PUSH2 0x717 PUSH1 0x0 DUP4 ADD DUP8 PUSH2 0x6D7 JUMP JUMPDEST PUSH2 0x724 PUSH1 0x20 DUP4 ADD DUP7 PUSH2 0x6F3 JUMP JUMPDEST PUSH2 0x731 PUSH1 0x40 DUP4 ADD DUP6 PUSH2 0x6D7 JUMP JUMPDEST PUSH2 0x73E PUSH1 0x60 DUP4 ADD DUP5 PUSH2 0x6D7 JUMP JUMPDEST SWAP6 SWAP5 POP POP POP POP POP JUMP INVALID LOG2 PUSH5 0x6970667358 0x22 SLT KECCAK256 0xAA CALL SWAP14 0x22 SWAP3 ORIGIN 0x2B LOG1 OR LOG1 EXTCODECOPY SAR GAS GASPRICE 0xB8 EXTCODEHASH 0xEE DUP2 SWAP13 GASLIMIT GASLIMIT 0xB7 0x4C KECCAK256 PUSH13 0xDC9DAB1C53B55A64736F6C6343 STOP ADDMOD 0xF STOP CALLER ", 
    "sourceMap": "68:1752:0:-:0;;;112:10;96:26;;;;;;;;;;;;;;;;;;;;68:1752;;;;;;"
}
```

remix abi
```json
[
	{
		"inputs": [],
		"stateMutability": "payable",
		"type": "constructor"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "nonce",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "claimPayment",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "shutdown",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
```

> 编译后的bin文件，可以使用evm运行


- ### [EVM操作码手册](https://www.ethervm.io/)  

<br>
<div align=center>
  <img src="../res/images/evm-opcodes.png" width="100%"></img>  
</div>


### [evmone](https://github.com/ethereum/evmone)   

```shell
git clone --recursive https://github.com/ethereum/evmone
cd evmone

# configure
cmake -S . -B build -DEVMONE_TESTING=ON

#build
cmake --build build --parallel

```

编译输出:
```shell
~/work/evmone# cmake --build build --parallel
[  3%] Built target tooling
[  8%] Built target instructions
[ 12%] Built target loader
[ 15%] Built target evmone-state
[ 20%] Built target evmone-bench-internal
[ 24%] Built target testutils
[ 41%] Built target evmone
[ 44%] Built target evmc-tool
[ 50%] Built target evmc-vmtester
[ 55%] Built target evmone-bench
[ 58%] Built target testutils-dump
[100%] Built target evmone-unittests
```



## 智能合约运行原理  
- ### [github FISCO BCOS 文章](https://github.com/FISCO-BCOS/FISCO-BCOS-DOC/blob/release-2/2.x/docs/articles/3_features/35_contract/solidity_operation_principle.md)  

### 引 言

作为一门面向智能合约的语言，Solidity与其他经典语言既有差异也有相似之处。一方面，服务于区块链的属性使其与其他语言存在差异。例如，合约的部署与调用均要经过区块链网络确认；执行成本需要被严格控制，以防止恶意代码消耗节点资源。另一方面，身为编程语言，Solidity的实现并未脱离经典语言，比如Solidity中包含类似栈、堆的设计，采用栈式虚拟机来进行字节码处理。本系列前几篇文章介绍了如何开发Solidity程序，为了让读者知其然更知其所以然，本文将进一步介绍Solidity的内部运行原理，聚焦于Solidity程序的生命周期和EVM工作机制。

### Solidity的生命周期

与其他语言一样，Solidity的代码生命周期离不开编译、部署、执行、销毁这四个阶段。下图整理展现了Solidity程序的完整生命周期：

![](../res/images/solidity_operation_principle/IMG_5474.PNG)

经编译后，Solidity文件会生成字节码。这是一种类似jvm字节码的代码。部署时，字节码与构造参数会被构建成交易，这笔交易会被打包到区块中，经由网络共识过程，最后在各区块链节点上构建合约，并将合约地址返还用户。当用户准备调用该合约上的函数时，调用请求同样也会经历交易、区块、共识的过程，最终在各节点上由EVM虚拟机来执行。

下面是一个示例程序，我们通过remix探索它的生命周期。

```
pragma solidity ^0.4.25;

contract Demo{
    uint private _state;
    constructor(uint state){
        _state = state;
    }
    function set(uint state) public {
        _state = state;
    }
}
```

### 编译

源代码编译完后，可以通过ByteCode按钮得到它的二进制：

```
608060405234801561001057600080fd5b506040516020806100ed83398101806040528101908080519060200190929190505050806000819055505060a4806100496000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806360fe47b1146044575b600080fd5b348015604f57600080fd5b50606c60048036038101908080359060200190929190505050606e565b005b80600081905550505600a165627a7a723058204ed906444cc4c9aabd183c52b2d486dfc5dea9801260c337185dad20e11f811b0029
```

还可以得到对应的字节码（OpCode）：

```
PUSH1 0x80 PUSH1 0x40 MSTORE CALLVALUE DUP1 ISZERO PUSH2 0x10 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x40 MLOAD PUSH1 0x20 DUP1 PUSH2 0xED DUP4 CODECOPY DUP2 ADD DUP1 PUSH1 0x40 MSTORE DUP2 ADD SWAP1 DUP1 DUP1 MLOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP DUP1 PUSH1 0x0 DUP2 SWAP1 SSTORE POP POP PUSH1 0xA4 DUP1 PUSH2 0x49 PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN STOP PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH1 0x3F JUMPI PUSH1 0x0 CALLDATALOAD PUSH29 0x100000000000000000000000000000000000000000000000000000000 SWAP1 DIV PUSH4 0xFFFFFFFF AND DUP1 PUSH4 0x60FE47B1 EQ PUSH1 0x44 JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST CALLVALUE DUP1 ISZERO PUSH1 0x4F JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x6C PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP PUSH1 0x6E JUMP JUMPDEST STOP JUMPDEST DUP1 PUSH1 0x0 DUP2 SWAP1 SSTORE POP POP JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 0x4e 0xd9 MOD DIFFICULTY 0x4c 0xc4 0xc9 0xaa 0xbd XOR EXTCODECOPY MSTORE 0xb2 0xd4 DUP7 0xdf 0xc5 0xde 0xa9 DUP1 SLT PUSH1 0xC3 CALLDATACOPY XOR 0x5d 0xad KECCAK256 0xe1 0x1f DUP2 SHL STOP 0x29 
```

其中下述指令集为set函数对应的代码，后面会解释set函数如何运行。

```
JUMPDEST DUP1 PUSH1 0x0 DUP2 SWAP1 SSTORE POP POP JUMP STOP
```

### 部署

编译完后，即可在remix上对代码进行部署，构造参数传入0x123:

![](../res/images/solidity_operation_principle/IMG_5475.PNG)

部署成功后，可得到一条交易回执：

![](../res/images/solidity_operation_principle/IMG_5476.PNG)

点开input，可以看到具体的交易输入数据：

![](../res/images/solidity_operation_principle/IMG_5477.PNG)

上面这段数据中，标黄的部分正好是前文中的合约二进制；而标紫的部分，则对应了传入的构造参数0x123。这些都表明，合约部署以交易作为介质。结合区块链交易知识，我们可以还原出整个部署过程：

- 客户端将部署请求(合约二进制，构造参数)作为交易的输入数据，以此构造出一笔交易
- 交易经过rlp编码，然后由发送者进行私钥签名
- 已签名的交易被推送到区块链上的节点
- 区块链节点验证交易后，存入交易池
- 轮到该节点出块时，打包交易构建区块，广播给其他节点
- 其他节点验证区块并取得共识。不同区块链可能采用不同共识算法，FISCO BCOS中采用PBFT取得共识，这要求经历三阶段提交（pre-prepare，prepare, commit）
- 节点执行交易，结果就是智能合约Demo被创建，状态字段_state的存储空间被分配，并被初始化为0x123

### 执行

根据是否带有修饰符view，我们可将函数分为两类：调用与交易。由于在编译期就确定了调用不会引起合约状态的变更，故对于这类函数调用，节点直接提供查询即可，无需与其他区块链节点确认。而由于交易可能引起状态变更，故会在网络间确认。下面将以用户调用了set(0x10)为假设，看看具体的运行过程。首先，函数set没有配置view/pure修饰符，这意味着其可能更改合约状态。所以这个调用信息会被放入一笔交易，经由交易编码、交易签名、交易推送、交易池缓存、打包出块、网络共识等过程，最终被交由各节点的EVM执行。在EVM中，由SSTORE字节码将参数0xa存储到合约字段_state中。该字节码先从栈上拿到状态字段_state的地址与新值0xa，随后完成实际存储。下图展示了运行过程：

![](../res/images/solidity_operation_principle/IMG_5478.PNG)

这里仅粗略介绍了set(0xa)是如何运行，下节将进一步展开介绍EVM的工作机制以及数据存储机制。

### 销毁

由于合约上链后就无法篡改，所以合约生命可持续到底层区块链被彻底关停。若要手动销毁合约，可通过字节码selfdestruct。销毁合约也需要进行交易确认，在此不多作赘述。

## EVM原理

在前文中，我们介绍了Solidity程序的运行原理。经过交易确认后，最终由EVM执行字节码。对EVM，上文只是一笔带过，这一节将具体介绍其工作机制。

### 运行原理

EVM是栈式虚拟机，其核心特征就是所有操作数都会被存储在栈上。下面我们将通过一段简单的Solidity语句代码看看其运行原理：

```
uint a = 1;
uint b = 2;
uint c = a + b;
```

这段代码经过编译后，得到的字节码如下：

```
PUSH1 0x1
PUSH1 0x2
ADD
```

为了读者更好了解其概念，这里精简为上述3条语句，但实际的字节码可能更复杂，且会掺杂SWAP和DUP之类的语句。我们可以看到，在上述代码中，包含两个指令：PUSH1和ADD，它们的含义如下：

- PUSH1：将数据压入栈顶。
- ADD：POP两个栈顶元素，将它们相加，并压回栈顶。

这里用半动画的方式解释其执行过程。下图中，sp表示栈顶指针，pc表示程序计数器。当执行完push1 0x1后，pc和sp均往下移：

![](../res/images/solidity_operation_principle/IMG_5479.PNG)

类似地，执行push1 0x2后，pc和sp状态如下：

![](../res/images/solidity_operation_principle/IMG_5480.PNG)

最后，当add执行完后，栈顶的两个操作数都被弹出作为add指令的输入，两者的和则会被压入栈：

![](../res/images/solidity_operation_principle/IMG_5481.PNG)



### 存储探究

在开发过程中，我们常会遇到令人迷惑的memory修饰符；阅读开源代码时，也会看到各种直接针对内存进行的assembly操作。不了解存储机制的开发者遇到这些情况就会一头雾水，所以，这节将探究EVM的存储原理。在前文《[智能合约编写之Solidity的基础特性](http://mp.weixin.qq.com/s?__biz=MzA3MTI5Njg4Mw==&mid=2247485625&idx=1&sn=9af6032cbf0ad0a3f7f8b7e85faebc77&chksm=9f2efaa5a85973b3fb118b3f1a6e2cd6aef8c1852ee97e93d98afeae71975c3cffc24a0b28fd&scene=21#wechat_redirect)》中我们介绍过，一段Solidity代码，通常会涉及到局部变量、合约状态变量。而这些变量的存储方式存在差别，下面代码表明了变量与存储方式之间的关系。

```
contract Demo{
    //状态存储
    uint private _state;

    function set(uint state) public {
        //栈存储
        uint i = 0;
        //内存存储
        string memory str = "aaa";
    }
}
```

#### 栈

栈用于存储字节码指令的操作数。在Solidity中，局部变量若是整型、定长字节数组等类型，就会随着指令的运行入栈、出栈。例如，在下面这条简单的语句中，变量值1会被读出，通过PUSH操作压入栈顶：

```
uint i = 1;
```

对于这类变量，无法强行改变它们的存储方式，如果在它们之前放置memory修饰符，编译会报错。

#### 内存

内存类似java中的堆，它用于储存"对象"。在Solidity编程中，如果一个局部变量属于变长字节数组、字符串、结构体等类型，其通常会被memory修饰符修饰，以表明存储在内存中。

本节中，我们将以字符串为例，分析内存如何存储这些对象。

##### 1. 对象存储结构

下面将用assembly语句对复杂对象的存储方式进行分析。assembly语句用于调用字节码操作。mload指令将被用于对这些字节码进行调用。mload(p)表示从地址p读取32字节的数据。开发者可将对象变量看作指针直接传入mload。在下面代码中，经过mload调用，data变量保存了字符串str在内存中的前32字节。

```
string memory str = "aaa";
bytes32 data;
assembly{
    data := mload(str)
}  
```

掌握mload，即可用此分析string变量是如何存储的。下面的代码将揭示字符串数据的存储方式：

```
function strStorage() public view returns(bytes32, bytes32){
    string memory str = "你好";
    bytes32 data;
    bytes32 data2;
    assembly{
        data := mload(str)
        data2 := mload(add(str, 0x20))
    }   
    return (data, data2);
}
```

data变量表示str的`0~31`字节，data2表示str的`32~63`字节。运行strStorage函数的结果如下：

```
0: bytes32: 0x0000000000000000000000000000000000000000000000000000000000000006
1: bytes32: 0xe4bda0e5a5bd0000000000000000000000000000000000000000000000000000
```

可以看到，第一个数据字得到的值为6，正好是字符串"你好"经UTF-8编码后的字节数。第二个数据字则保存的是"你好"本身的UTF-8编码。熟练掌握了字符串的存储格式之后，我们就可以运用assembly修改、拷贝、拼接字符串。读者可搜索Solidity的字符串库，了解如何实现string的concat。

##### 2. 内存分配方式

既然内存用于存储对象，就必然涉及到内存分配方式。memory的分配方式非常简单，就是顺序分配。下面我们将分配两个对象，并查看它们的地址：

```
function memAlloc() public view returns(bytes32, bytes32){
    string memory str = "aaa";
    string memory str2 = "bbb";
    bytes32 p1;
    bytes32 p2;
    assembly{
        p1 := str
        p2 := str2
    }   
    return (p1, p2);
}
```

运行此函数后，返回结果将包含两个数据字：

```
0: bytes32: 0x0000000000000000000000000000000000000000000000000000000000000080
1: bytes32: 0x00000000000000000000000000000000000000000000000000000000000000c0
```

这说明，第一个字符串str1的起始地址是0x80，第二个字符串str2的起始地址是0xc0，之间64字节，正好是str1本身占据的空间。此时的内存布局如下，其中一格表示32字节（一个数据字，EVM采用32字节作为一个数据字，而非4字节）：

![](../res/images/solidity_operation_principle/IMG_5482.PNG)

- 0x40~0x60：空闲指针，保存可用地址，本例中是0x100，说明新的对象将从0x100处分配。可以用mload(0x40)获取到新对象的分配地址。
- 0x80~0xc0：对象分配的起始地址。这里分配了字符串aaa
- 0xc0~0x100：分配了字符串bbb
- 0x100~...：因为是顺序分配，新的对象将会分配到这里。

#### 状态存储

顾名思义，状态存储用于存储合约的状态字段。从模型而言，存储由多个32字节的存储槽构成。在前文中，我们介绍了Demo合约的set函数，里面0x0表示的是状态变量_state的存储槽。所有固定长度变量会依序放到这组存储槽中。对于mapping和数组，存储会更复杂，其自身会占据1槽，所包含数据则会按相应规则占据其他槽，比如mapping中，数据项的存储槽位由键值k、mapping自身槽位p经keccak计算得来。从实现而言，不同的链可能采用不同实现，比较经典的是以太坊所采用的MPT树。由于MPT树性能、扩展性等问题，FISCO BCOS放弃了这一结构，而采用了分布式存储，通过rocksdb或mysql来存储状态数据，使存储的性能、可扩展性得到提高。

## 结语

本文介绍了Solidity的运行原理，运行原理总结如下。首先，Solidity源码会被编译为字节码，部署时，字节码会以交易为载体在网络间确认，并在节点上形成合约；合约函数调用，如果是交易类型，会经过网络确认，最终由EVM执行。EVM是栈式虚拟机，它会读取合约的字节码并执行。在执行过程中，会与栈、内存、合约存储进行交互。其中，栈用于存储普通的局部变量，这些局部变量就是字节码的操作数；内存用于存储对象，采用length+body进行存储，顺序分配方式进行内存分配；状态存储用于存储状态变量。理解Solidity的运行方式及其背后原理，是成为Solidity编程高手必经之路。












