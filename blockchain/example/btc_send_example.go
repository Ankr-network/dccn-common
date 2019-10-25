package main

import (
    "fmt"
    "github.com/Ankr-network/dccn-common/blockchain/btc"
)

func main() {
    hash, err := btc.TransferBtc("test3", "e68748dcfcd846f89831134424205007", "",  1)
    if err != nil {
       fmt.Printf("error %s", err.Error())
    }
    fmt.Printf("tx hash : ", hash)

}
