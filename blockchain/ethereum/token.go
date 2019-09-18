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
    switch name {
    case "ANKR":
        contractAddress:="0x8290333cef9e6d528dd5618fb97a76f268f3edd4"
        token, err := NewAnkrToken(name, ethClient,contractAddress)
        return token, err
    case "ANKR-T":
        contractAddress:="0x38A03b35a6662D35B226879E57d499eE26A6D4B5"
        token, err := NewAnkrToken(name, ethClient,contractAddress)
        return token, err
    case "USDT":
        token, err := NewUsdtToken(name, ethClient)
        return token, err
    }
    return nil, nil
}
