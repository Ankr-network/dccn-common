package ethereum

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	EthereumNodeURL = "https://kindly-civil-escargot.quiknode.io/baa643a7-78b4-438e-b53a-76569410a448/bEpPhdsXJUrNvWzednALKg==/"
)

type EthService struct {
	EthClient *ethclient.Client
}

func NewEthService() (*EthService, error){
	c, err := ethclient.Dial(EthereumNodeURL)
	if err != nil {
		return  nil, err
	}

	return &EthService{c }, nil
}

func (s *EthService) TokenTransfer(assertName, fromKey, fromPassword, toAddress string, amount *big.Float) error {
	token, err := newToken(assertName, s.EthClient)
	if err != nil {
		return err
	}

	toAddr := common.HexToAddress(toAddress)

	auth, err := bind.NewTransactor(strings.NewReader(fromKey), fromPassword)
	if err != nil {

		return err
	}

	decimal, err := token.Decimals(nil)
	if err != nil {
		return err
	}

	tenDecimal := big.NewFloat(math.Pow(10, float64(decimal)))
	tens := tenDecimal.String()

	fmt.Printf("%s, %s\n", tens, amount.String())
	convertAmountF := new(big.Float).Mul(tenDecimal, amount)
	convertAmount  := new(big.Int)
	convertAmountF.Int(convertAmount)

	fmt.Printf("%s", convertAmount.String())

	_, err = token.transfer(auth, toAddr, convertAmount)

	return err
}
