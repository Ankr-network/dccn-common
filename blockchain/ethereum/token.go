package ethereum

import (
    "github.com/ethereum/go-ethereum/core/types"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

type Token interface {
    DecimalsConvert(amount *big.Float) (convertAmount *big.Int, err error)
    transfer(opts *bind.TransactOpts, toAddress common.Address, value *big.Int) (*types.Transaction, error)
}

func newToken(name string, ethClient *ethclient.Client) (Token, error) {
    ContractAddreses = make(map[string]string)
    ContractAddreses["ANKR"]=AnkrContractAddress
    ContractAddreses["ANKR-T"]=AnkrContractAddressTest
    switch name {
    case "ANKR","ANKR-T":
        token, err := NewAnkrToken(name, ethClient)
        return token, err
    case "USDT":
        token, err := NewUsdtToken(name, ethClient)
        return token, err
    }
    return nil, nil
}
