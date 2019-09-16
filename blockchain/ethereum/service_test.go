package ethereum

import (
    "context"
    "fmt"
    "math/big"
    "testing"
)

func TestStart(t *testing.T) {

    fromKey := ""      //Your from account key
    fromPassword := "" //Your from account password
    ethS, err := NewEthService()
    if ethS == nil || err != nil {
        t.Error("Can't create eth service")
    }

    amount := big.NewFloat(50)
    //assertName ANKR USDT
    hash, err := ethS.TokenTransfer("USDT", fromKey, fromPassword, "0xbbb092e9d4ddaf4e6a793c83eb4fa1a6bcd06389", amount)
    if err != nil {
        t.Errorf("ethS.TokenTransfer err: %s /n", err.Error())
    }
    fmt.Println(hash)
    var from = ""
    var to = ""
    var key = ""
    var amount1 = big.NewInt(10)
    hash, err = ethS.EthTransaction(context.Background(), from, to, key, amount1)
    if err != nil {
        fmt.Printf("error has happend %s", err)
    } else {
        fmt.Println("transaction hash is ", hash)
    }
}
