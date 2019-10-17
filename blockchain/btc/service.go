package btc

import (
    "github.com/huahuayu/go-bitcoind"
)

type Rpcclient struct {
    rpcclient *bitcoind.Bitcoind
}

func NewRpcClient(host, user, password, passphrase string, port int, usessl bool) (bitc *Rpcclient, err error) {
    bc, err := bitcoind.New(host, port, user, password, usessl)
    if err != nil {
        return bitc, err
    }
    return &Rpcclient{rpcclient: bc}, err
}

func (cli *Rpcclient) sendTransaction(fromAddress string, toAddresses map[string]float64) (txHash string, err error) {
    txHash, err = cli.rpcclient.SendMany(fromAddress, toAddresses, 1, "")
    return txHash, err
}
// https://chainquery.com/bitcoin-cli/createrawtransaction
