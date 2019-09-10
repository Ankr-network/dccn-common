package ethereum

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"math/big"
	"testing"
)

func TestStart(t *testing.T) {

	fromKey := ""//Your from account key
	fromPassword := ""//Your from account password
	ethS, err := NewEthService()
	if ethS == nil || err != nil {
		t.Error("Can't create eth service")
	}

	amount := big.NewFloat(50)

	hash,err:= ethS.TokenTransfer("ANKR", fromKey, fromPassword, "0xbbb092e9d4ddaf4e6a793c83eb4fa1a6bcd06389", amount)
	if err != nil {
		t.Errorf("ethS.TokenTransfer err: %s", err.Error())
	}
	fmt.Println(hash)

}
