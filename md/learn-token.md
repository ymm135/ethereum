- # WETH/ERC-20/ERC-721

## WETH

```js
    string public name     = "Wrapped Ether";
    string public symbol   = "WETH";
    uint8  public decimals = 18;
```

三个方法描述WETH的基本信息


如何存款呢: deposit 
```js
    function deposit() public payable {
        balanceOf[msg.sender] += msg.value;
        emit Deposit(msg.sender, msg.value);
    }
```

那个用户调用该方法，就是往该合约里存钱。  
比如这个交易:  https://sepolia.etherscan.io/tx/0x03121e3249eb1ad7caed621842b4c11f4946982b9b895a7c7fecb5ee6c9d770a  

<br>
<div align=center>
  <img src="../res/images/1/leran-token-1.png" width="85%"></img>
</div>


提取
```js
    function withdraw(uint wad) public {
        require(balanceOf[msg.sender] >= wad, "");
        balanceOf[msg.sender] -= wad;
        msg.sender.transfer(wad);
        emit Withdrawal(msg.sender, wad);
    }
```

<br>
<div align=center>
  <img src="../res/images/1/leran-token-2.png" width="85%"></img>
</div>

<br>
<div align=center>
  <img src="../res/images/1/leran-token-3.png" width="85%"></img>
</div>

总量
```sh
  function totalSupply() public view returns (uint) {
        return address(this).balance;
    }
```

> 是这个合约的总量，多个用户存款的总和。   

每次转账之前要授权:  
```js
    function approve(address guy, uint wad) public returns (bool) {
        allowance[msg.sender][guy] = wad;
        emit Approval(msg.sender, guy, wad);
        return true;
    }
```

```js
function transfer(address dst, uint wad) public returns (bool) {
        return transferFrom(msg.sender, dst, wad);
    }

    function transferFrom(address src, address dst, uint wad)
        public
        returns (bool)
    {
        require(balanceOf[src] >= wad, "");

        if (src != msg.sender && allowance[src][msg.sender] != uint(-1)) {
            require(allowance[src][msg.sender] >= wad, "");
            allowance[src][msg.sender] -= wad;
        }

        balanceOf[src] -= wad;
        balanceOf[dst] += wad;

        emit Transfer(src, dst, wad);

        return true;
    }
```

每次授权的金额会减少，超过授权额度之后，会再次授权。  

<br>
<div align=center>
  <img src="../res/images/1/leran-token-4.png" width="30%"></img>
</div>


<br>
<div align=center>
  <img src="../res/images/1/leran-token-5.png" width="85%"></img>
</div>

## ERC-20

ERC20标准方法
```js
function name() public view returns (string)
function symbol() public view returns (string)
function decimals() public view returns (uint8)
function totalSupply() public view returns (uint256)
function balanceOf(address _owner) public view returns (uint256 balance)
function transfer(address _to, uint256 _value) public returns (bool success)
function transferFrom(address _from, address _to, uint256 _value) public returns (bool success)
function approve(address _spender, uint256 _value) public returns (bool success)
function allowance(address _owner, address _spender) public view returns (uint256 remaining)
```

这个示例中在构造时，需要传入总量,并且往该账户中铸造币  
```js
    constructor(uint _totalSupply) public {
        uint chainId;
        assembly {
            chainId := chainid()
        }
        DOMAIN_SEPARATOR = keccak256(
            abi.encode(
                keccak256('EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)'),
                keccak256(bytes(name)),
                keccak256(bytes('1')),
                chainId,
                address(this)
            )
        );
        _mint(msg.sender, _totalSupply);
    }

    function _mint(address to, uint value) internal {
        totalSupply = totalSupply.add(value);
        balanceOf[to] = balanceOf[to].add(value);
        emit Transfer(address(0), to, value);
    }
```

接下来是转账授权  
```js

    function _approve(address owner, address spender, uint value) private {
        allowance[owner][spender] = value;
        emit Approval(owner, spender, value);
    }

    function approve(address spender, uint value) external returns (bool) {
        _approve(msg.sender, spender, value);
        return true;
    }

    function transferFrom(address from, address to, uint value) external returns (bool) {
        if (allowance[from][msg.sender] != uint(-1)) {
            allowance[from][msg.sender] = allowance[from][msg.sender].sub(value);
        }
        _transfer(from, to, value);
        return true;
    }
```

另外转账还需要销毁
```js
    function _burn(address from, uint value) internal {
        balanceOf[from] = balanceOf[from].sub(value);
        totalSupply = totalSupply.sub(value);
        emit Transfer(from, address(0), value);
    }

    function _transfer(address from, address to, uint value) private {
        uint burnAmount = value / 100;
        _burn(from, burnAmount);
        uint transferAmount = value.sub(burnAmount);
        balanceOf[from] = balanceOf[from].sub(transferAmount);
        balanceOf[to] = balanceOf[to].add(transferAmount);
        emit Transfer(from, to, transferAmount);
    }

    function transfer(address to, uint value) external returns (bool) {
        _transfer(msg.sender, to, value);
        return true;
    }
```

另外还有一个`permit`  
```js
    function permit(address owner, address spender, uint value, uint deadline, uint8 v, bytes32 r, bytes32 s) external {
        require(deadline >= block.timestamp, 'EXPIRED');
        bytes32 digest = keccak256(
            abi.encodePacked(
                '\x19\x01',
                DOMAIN_SEPARATOR,
                keccak256(abi.encode(PERMIT_TYPEHASH, owner, spender, value, nonces[owner]++, deadline))
            )
        );
        address recoveredAddress = ecrecover(digest, v, r, s);
        require(recoveredAddress != address(0) && recoveredAddress == owner, 'INVALID_SIGNATURE');
        _approve(owner, spender, value);
    }
```  

如果使用approve，则允许spender最多使用wad个代币。

如果你给某人提供有效的签名，则该人可以调用permit以允许spender  使用你的代币。

因此，基本上，“无 gas”交易背后的模式是制作可以提供给某人的签名，以便他们可以安全地执行特殊交易。这就像授予某人执行函数的权限。

这是一种授权模式。  



## ERC-721
