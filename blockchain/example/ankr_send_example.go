package main

import (
    "fmt"
    "math/big"

    "github.com/Ankr-network/dccn-common/blockchain/ethereum"
)

func main() {
    var key = "{\"address\":\"fe688261dc29678693b98163a6532dcbb0ce265b\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"73ece1e686c4c304d796862da553d5b03008d0ae5472c10a76eae8d5327ca878\",\"cipherparams\":{\"iv\":\"02dc37db89a423f229f25ef170f43e44\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"c5016aa0735e0f32b6695a408f938b1df2b1411bbbaadd60087bcd3e29a33360\"},\"mac\":\"b0243c076490b91a20ad989c513b3d1b524093558c4c28495726c094d4d11764\"},\"id\":\"cedf152d-516f-4d24-a0b0-6bd599ddacb3\",\"version\":3}"
    var password = "lvyanlong"
    var to=""
    ethS, err := ethereum.NewEthService()
    if ethS == nil || err != nil {
        fmt.Println("Can't create eth service")
    }

    amount := big.NewFloat(0)
    //assertName ANKR USDT
    hash, err := ethS.TokenTransfer("ANKR", key, password, to, amount)
    if err != nil {
        fmt.Println("ethS.TokenTransfer err:", err.Error())
    }
    fmt.Println(hash)
}