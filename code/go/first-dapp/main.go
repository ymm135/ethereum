package main

import (
	"context"
	counterSol "counter/sol"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("Hello Sol")

	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("Dial Error", err.Error())
		return
	}
	defer conn.Close()

	accountAddress := common.HexToAddress("0xd8b431ecef36518b7fa24224cc7678eca963f4e3")
	contractAddress := common.HexToAddress("0x8b3A1C22E6ADFCa111fF07Ab0a06e29B273717b0")

	counterObj, err := counterSol.NewCounter(contractAddress, conn)
	if err != nil {
		fmt.Println("NewCounter Error", err.Error())
		return
	}

	getParam := &bind.CallOpts{
		Pending: false,
		From:    accountAddress,
		Context: context.Background(),
	}
	counterVal, err := counterObj.CounterCaller.Get(getParam)
	if err != nil {
		fmt.Println("Get Error ", err.Error())
		return
	}

	fmt.Println("counterVal:", counterVal)
}
