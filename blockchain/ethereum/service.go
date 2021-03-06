package ethereum

import (
    "context"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "math/big"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

const (
    EthereumNodeURL = "https://mainnet.infura.io/v3/3141358baff64c0ca4154e356c1dd20d"
)

type EthService struct {
    EthClient *ethclient.Client
}

func NewEthService() (*EthService, error) {
    c, err := ethclient.Dial(EthereumNodeURL)
    if err != nil {
        return nil, err
    }

    return &EthService{c}, nil
}

func (s *EthService) TokenTransfer(assertName, fromKey, fromPassword, toAddress string, amount *big.Float) (hash string, err error) {
    token, err := newToken(assertName, s.EthClient)
    if err != nil {
        return hash, err
    }
    toAddr := common.HexToAddress(toAddress)

    auth, err := bind.NewTransactor(strings.NewReader(fromKey), fromPassword)

    if err != nil {

        return hash, err
    }

    gasPrice, err := s.EthClient.SuggestGasPrice(context.Background())
    if err != nil {
        return hash, err
    }
    auth.GasPrice = gasPrice
    auth.GasLimit = uint64(100000)
    convertAmount, err := token.DecimalsConvert(amount)
    if err != nil {
        return hash, err
    }
    tx, err := token.transfer(auth, toAddr, convertAmount)
    if err != nil {
        return hash, err
    }
    return tx.Hash().Hex(), err
}

func (s *EthService) EthTransaction(ctx context.Context, from, to, key string, amount *big.Int) (hash string, err error) {
    gas, err := s.EthClient.SuggestGasPrice(ctx)
    if err != nil {
        return hash, err
    }
    gasLimit := uint64(21000)
    nonce, err := s.EthClient.PendingNonceAt(ctx, common.HexToAddress(from))
    if err != nil {
        return hash, err
    }
    tx := types.NewTransaction(nonce, common.HexToAddress(to), amount, gasLimit, gas, nil)
    chainID, err := s.EthClient.NetworkID(context.Background())
    if err != nil {
        return hash, err
    }
    privateKey, err := crypto.HexToECDSA(key)
    if err != nil {
        return hash, err
    }
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        return hash, err
    }
    err = s.EthClient.SendTransaction(ctx, signedTx)

    return signedTx.Hash().Hex(), err

}
