package keystore_test

import (
	"testing"
	"fmt"

	ks "github.com/Ankr-network/dccn-common/wallet/keystore"
)
// test data
//I8coAiVgzqR14lAQtbM/F44u8PKxKTb9gIymvDpTxficDVWtMMFKKIk2DhT89ONafR4dvKHqXGVV/9B7n6CNZA==
//nA1VrTDBSiiJNg4U/PTjWn0eHbyh6lxlVf/Qe5+gjWQ=
//3897CE9274875578E17402C080DA3D66DE19DA85038DD0

const (
        PRIV_KEY = "I8coAiVgzqR14lAQtbM/F44u8PKxKTb9gIymvDpTxficDVWtMMFKKIk2DhT89ONafR4dvKHqXGVV/9B7n6CNZA=="
	ORIG_PASSWORD = "12345678"
	NEW_PASSWORD = "11111111"
)

func TestEncrytDecrypt(t *testing.T) {
	t.Log("Testing encryption decryption")

	cryptoStruct, err := ks.EncryptDataV3([]byte(PRIV_KEY),
		[]byte(ORIG_PASSWORD), ks.StandardScryptN, ks.StandardScryptP)

	if err == nil {
		//fmt.Printf("\ncryptoStruct: %v\n ", cryptoStruct)
	} else {
		t.Error(err)
	}

	re, err := ks.DecryptDataV3(cryptoStruct, ORIG_PASSWORD)
	if err == nil {
		//fmt.Printf("\nresult: %s\n", string(re))
	} else {
		t.Error(err)
	}

	if string(re) != PRIV_KEY {
                t.Error("result doesn't match!")
        }
}

func TestPrivKeyAndKeyStore(t *testing.T) {
	t.Log("Testing privkey and keystore")

	re, _, err := ks.PrivKeyToKeyStoreJsonString("myname",
		PRIV_KEY, ORIG_PASSWORD)
	if err == nil {
		//fmt.Printf(re)
	} else {
		t.Error(err)
	}

	priv, err := ks.KeyStoreJsonStringToPrivKey(re, ORIG_PASSWORD)
	if err == nil {
		//fmt.Printf(priv)
	} else {
		t.Error(err)
	}

	if priv != PRIV_KEY {
                t.Error("result doesn't match!")
        }
}

func TestResetPassword(t *testing.T) {
	t.Log("Testing reset password")
        re, _, err := ks.PrivKeyToKeyStoreJsonString("myname",
                PRIV_KEY, ORIG_PASSWORD)
        if err == nil {
                //fmt.Printf(re)
        } else {
                t.Error(err)
        }

        newjson, _, err := ks.ResetPasswordKeyStoreJsonString(re, ORIG_PASSWORD, NEW_PASSWORD)
        if err == nil {
                //fmt.Printf(newjson)
        } else {
                t.Error(err)
        }

	priv, err := ks.KeyStoreJsonStringToPrivKey(newjson, NEW_PASSWORD)
        if err == nil {
                //fmt.Printf(priv)
        } else {
                t.Error(err)
        }

	if priv != PRIV_KEY {
                fmt.Printf("result doesn't match!")
                t.Error("result doesn't match!")
        }
}
