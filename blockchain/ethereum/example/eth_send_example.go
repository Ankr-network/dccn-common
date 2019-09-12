package example

import (
    "context"
    "fmt"
    "github.com/Ankr-network/dccn-common/blockchain/ethereum"
    "math/big"
)

func main() {
    var from = "0x83891E6db12868E6118327B6a917265B854480a6"
    var to = "0x93e144bEa008850d90c9e86912Ae73b3d7A0578D"
    var key = "31098708AA769DBB4732A32EF08FBB3DE46925F1FF47A775FC9A50AE900021CE"
    var amount = big.NewInt(10)
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
