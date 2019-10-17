package main

import (
    "fmt"
    "github.com/Ankr-network/dccn-common/blockchain/ethereum/ustd"
    "math/big"
)

func main() {
    var to = ""
    var key = ""
    var password = ""
    var amount = big.NewInt(10)
    usdt, _ := ustd.NewUsdtService()
    hash, err := usdt.TokenTransfer(to,key, password, amount)
    if err != nil {
        fmt.Println("error", err)
    } else {
        fmt.Printf("%s", hash)
    }

}
