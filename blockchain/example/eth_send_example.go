package main

import (
    "context"
    "fmt"
    "github.com/Ankr-network/dccn-common/blockchain/ethereum"
    "math/big"
)

func main() {
    var from = ""
    var to = ""
    var key = ""
    var amount = big.NewInt(0)
    eth, err := ethereum.NewEthService()
    if err != nil {
        fmt.Printf("error has happend %s", err.Error())
    }
    hash, err := eth.EthTransaction(context.Background(), from, to, key, amount)
    if err != nil {
        fmt.Printf("error has happend %s", err)
    } else {
        fmt.Println("transaction hash is ", hash)
    }
}
