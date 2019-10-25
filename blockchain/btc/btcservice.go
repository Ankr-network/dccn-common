package btc

import (
    "fmt"
    "github.com/btcsuite/btcd/chaincfg"
    "github.com/btcsuite/btcutil"
    "github.com/stevenroose/go-bitcoin-core-rpc"
)

func TransferBtc(from, to, passWord string, amount float64) (hash string, err error) {
    config := &rpcclient.ConnConfig{Host: "btc-01.dccn.ankr.com:8332", User: "user", Pass: "1q2w3e4r",}
    client, err := rpcclient.New(config)
    if err != nil {
        fmt.Errorf("%s", err.Error())
        return hash, err
    }
    err = client.WalletPassphrase(passWord, 5)
    if err != nil {
        fmt.Errorf("WalletPassphrase %s", err.Error())
    }
    address, err := btcutil.DecodeAddress(to, &chaincfg.MainNetParams)
    if err != nil {
        fmt.Errorf("DecodeAddress  %s", err.Error())
        return hash, err
    }
    btcAmount, _ := btcutil.NewAmount(amount)
    hashTx, err := client.SendFrom(from, address, btcAmount)

    if err != nil {
        fmt.Errorf("sendFrom error %s", err.Error())
        return hash, err
    }
    return hashTx.String(), nil
}
