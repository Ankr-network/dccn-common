package wallet

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	_ "encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const PubKeyEd25519Size = 32
const PrivKeyEd25519Size = 64
const PubkeyStart = "PubKeyEd25519{"

/*
  Generate new key-pair. return priv_key, pub_key, address in string format.
  Caller does not need to care about the encoding, and API will handle internally.

  example output for priv_key, pub_key, address:
  mB3bPsj9kUsAPmFw11Gqb6AYi7nQ8PFrNI+G62IRyYnstkObfl1KIeQi8pOfpNovC2iikxivhW9baCLStM2hyA==
  7LZDm35dSiHkIvKTn6TaLwtoopMYr4VvW2gi0rTNocg=
  0D9FE6A785C830D2BE66FE40E0E7FE3D9838456C
*/
func GenerateKeys() (priv_key_base64, pub_key_base64, address string) {
	mykey := ed25519.GenPrivKey()

	// data for test
	//mykey := [194 108 153 102 131 30 117 105 108 61 64 213 8 236 190 78 37 92 172 128 
	//79 114 125 214 36 223 36 229 195 208 128 139 194 241 198 220 71 93 5 181 208 29 204 
	//137 106 93 2 75 246 16 112 214 45 17 177 88 197 232 231 169 255 78 132 206]
	//mykey := ed25519.PrivKeyEd25519{194,108, 153, 102, 131, 30, 117, 105, 108, 61, 64, 213, 
	//8, 236, 190, 78, 37, 92, 172, 128, 79, 114, 125, 214, 36, 223, 36, 229, 195, 208, 128, 
	//139, 194, 241, 198, 220, 71, 93, 5, 181, 208, 29, 204, 137, 106, 93, 2, 75, 246, 16, 
	//112, 214, 45, 17, 177, 88, 197, 232, 231, 169, 255, 78, 132, 206}

	privArray := [PrivKeyEd25519Size]byte(mykey)
	privBytes := privArray[:]
	privB64 := base64.StdEncoding.EncodeToString(privBytes)
	priv_key_base64 = privB64

	mystr := fmt.Sprintf("%s", mykey.PubKey())
	mystr = mystr[len(PubkeyStart) : len(PubkeyStart)+PrivKeyEd25519Size]

	src := []byte(mystr)
	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)

	pubB64 := base64.StdEncoding.EncodeToString([]byte(dst))
	pub_key_base64 = pubB64

	address = fmt.Sprintf("%s", mykey.PubKey().Address())
	return priv_key_base64, pub_key_base64, address
}

/*
  Get address by public key.
  With this function, if a user has public_key, address is not very necessary to save in database.
*/
func GetAddressByPublicKey(pub_key string) (string, error) {
	pubKeyObject, err := deserilizePubKey(pub_key)
	if err != nil {
		return "", err
	}

	address := fmt.Sprintf("%s", pubKeyObject.Address())
	return address, nil
}

/*
 do sha256 and sign.
 metering feature needs this.
*/
func Sign(input, priv_key string) (result string, err_ret error) {
	privKeyObject, err := deserilizePrivKey(priv_key)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256([]byte(input))
	encrypted_result, err := privKeyObject.Sign([]byte(string(sum[:32])))
	if err != nil {
		return "", err
	}

	encrypted_result_b64 := base64.StdEncoding.EncodeToString([]byte(encrypted_result))

	result = encrypted_result_b64
	err_ret = nil

	return
}

func SetValidators(ip, port, pubkey, power string) (result string, err_ret error) {

        _, err := strconv.ParseInt(string(power), 10, 64)
        if err != nil {
                return "", err
        }

	cl := getHTTPClient(ip, port)
	_, err = cl.Status()
	if err != nil {
		return "", err
	}

	pubKeyObject, err := deserilizePubKey(pubkey)
        if err != nil {
                return "", err
        }

	mystr := fmt.Sprintf("%s", pubKeyObject)
        mystr = mystr[len(PubkeyStart) : len(PubkeyStart)+PrivKeyEd25519Size]
	fmt.Println(mystr)

	btr, err := cl.BroadcastTxCommit(types.Tx(
		fmt.Sprintf("%s:%s/%s", string("val"), mystr, power)))

        if err != nil {
                return "", err
        } else if btr.CheckTx.Code != 0 {
                return "", errors.New(btr.CheckTx.Log)
        } else if btr.DeliverTx.Code != 0{
                return "", errors.New(btr.DeliverTx.Log)
        }

        client.WaitForHeight(cl, btr.Height+1, nil)

	fmt.Println(btr)
	return "", nil
}

/*
 get balance based on (ip, port, address).
 ip: blockchain node ip or domain
 port: blockchain node port, usually 26657
 address: wallet address
 return: balance or error
*/
func GetBalance(ip, port, address string) (balance string, err_ret error) {

	cl := getHTTPClient(ip, port)

	_, err := cl.Status()
	if err != nil {
		return "", err
	}

	//fmt.Println("LatestBlockHeight:", status.SyncInfo.LatestBlockHeight)
	// curl  'localhost:26657/abci_query?data="bal:1234567890123456789012345678901234567890"'

	res, err := cl.ABCIQuery("/websocket", cmn.HexBytes(fmt.Sprintf("%s:%s", "bal", address)))
	qres := res.Response
	if !qres.IsOK() {
		return "", errors.New("Query balance failure, connect error.")
	} else {
		balanceNonceSlices := strings.Split(string(qres.Value), ":")
		if len(balanceNonceSlices) == 2 {
			balance = balanceNonceSlices[0]
		} else {
			return "", errors.New("Query balance failure, balance format incorrect.")
		}
	}

	return balance, nil
}

/*
Example: send 10 tokens from address1 to address2.

from_address: address1
to_address: address2
amount: 10
priv_key: address1's priv_key
public_key: address1's public key

note: priv_key is used for signature, and it will not be sent or saved.
*/
func SendCoins(ip, port, priv_key, from_address, to_address, amount, public_key string) error {
	var nonce string
	cl := getHTTPClient(ip, port)

	_, err := cl.Status()
	if err != nil {
		return err
	}

	res, err := cl.ABCIQuery("/websocket", cmn.HexBytes(fmt.Sprintf("%s:%s", "bal", from_address)))
	qres := res.Response
	if !qres.IsOK() {
		return errors.New("Query nonce failure, connect error.")
	} else {
		balanceNonceSlices := strings.Split(string(qres.Value), ":")
		if len(balanceNonceSlices) == 2 {
			nonce = balanceNonceSlices[1]
		} else {
			return errors.New("Query nonce failure, balance format incorrect.")
		}
	}

	nonceInt, err := strconv.ParseInt(string(nonce), 10, 64)
	if err != nil {
		return err
	}

	nonceInt++
	nonce = fmt.Sprintf("%d", nonceInt)

	sig, err := Sign(fmt.Sprintf("%s%s%s%s", from_address, to_address, amount, nonce), priv_key)
	if err != nil {
		return err
	}

	//fmt.Printf("%s=%s:%s:%s:%s:%s:%s\n", string("trx_send"), from_address, to_address, amount, nonce, public_key, sig)
	btr, err := cl.BroadcastTxCommit(types.Tx(
		fmt.Sprintf("%s=%s:%s:%s:%s:%s:%s", string("trx_send"), from_address, to_address, amount, nonce, public_key, sig)))

	if err != nil {
		return err
	} else if btr.CheckTx.Code != 0 {
		return errors.New(btr.CheckTx.Log)
	} else if btr.DeliverTx.Code != 0{
		return errors.New(btr.DeliverTx.Log)
	}

	client.WaitForHeight(cl, btr.Height+1, nil)

	return nil
}

/*
query stake based on (ip, port)
everyone can call this function, because it's read only.
*/
func GetStake(ip, port string) (stake string, err_ret error) {

	cl := getHTTPClient(ip, port)

	_, err := cl.Status()
	if err != nil {
		return "", err
	}

	// curl  'localhost:26657/abci_query?data="stk"'
	res, err := cl.ABCIQuery("/websocket", cmn.HexBytes(fmt.Sprintf("%s", "stk")))
	qres := res.Response
	if !qres.IsOK() {
		return "", errors.New("Query balance failure, connect error.")
	} else {
		stakeNonceSlices := strings.Split(string(qres.Value), ":")
		if len(stakeNonceSlices) == 2 {
			stake = stakeNonceSlices[0]
		} else {
			return "", errors.New("Query balance failure, balance format incorrect.")
		}
	}

	return stake, nil
}

/*
ADMIN API:
set stake based on (ip, port, priv_key, amount, pub_key).
 ip: blockchain node ip or domain
 port: blockchain node port, usually 26657
 priv_key: admin private key
 amount: token amount
 pub_key: admin pubic key
*/
func SetStake(ip, port, priv_key, amount, public_key string) error {
	var nonce string
	cl := getHTTPClient(ip, port)

	_, err := cl.Status()
	if err != nil {
		return err
	}

	res, err := cl.ABCIQuery("/websocket", cmn.HexBytes(fmt.Sprintf("%s", "stk")))
	qres := res.Response
	if !qres.IsOK() {
		return errors.New("Query nonce failure, connect error.")
	} else {
		balanceNonceSlices := strings.Split(string(qres.Value), ":")
		if len(balanceNonceSlices) == 2 {
			nonce = balanceNonceSlices[1]
		} else {
			return errors.New("Query nonce failure, balance format incorrect.")
		}
	}

	nonceInt, err := strconv.ParseInt(string(nonce), 10, 64)
	if err != nil {
		return err
	}

	nonceInt++
	nonce = fmt.Sprintf("%d", nonceInt)

	sig, err := Sign(fmt.Sprintf("%s%s", amount, nonce), priv_key)
	if err != nil {
		return err
	}

	//fmt.Printf("%s=%s:%s:%s:%s\n", string("set_stk"), amount, nonce, public_key, sig)
	btr, err := cl.BroadcastTxCommit(types.Tx(
		fmt.Sprintf("%s=%s:%s:%s:%s", string("set_stk"), amount, nonce, public_key, sig)))

	if err != nil {
		return err
	}
	client.WaitForHeight(cl, btr.Height+1, nil)

	return nil
}

/*
ADMIN API:
set balance based on (ip, port, priv_key, address, amount, public_key)
priv_key: admin priv_key
address: set amount to this address
amount: token amount
public_key: admin pubic_key
*/
func SetBalance(ip, port, priv_key, address, amount, public_key string) error {
	var nonce string = "0"
	cl := getHTTPClient(ip, port)

	_, err := cl.Status()
	if err != nil {
		return err
	}

	res, err := cl.ABCIQuery("/websocket", cmn.HexBytes(fmt.Sprintf("%s:%s", "bal", address)))
	qres := res.Response
	if !qres.IsOK() {
		return errors.New("Query nonce failure, connect error.")
	} else {
		if string(qres.Value) != "" {
		balanceNonceSlices := strings.Split(string(qres.Value), ":")
		if len(balanceNonceSlices) == 2 {
			nonce = balanceNonceSlices[1]
		} else {
			return errors.New("Query nonce failure, balance format incorrect.")
		}
	}
	}

	nonceInt, err := strconv.ParseInt(string(nonce), 10, 64)
	if err != nil {
		return err
	}

	nonceInt++
	nonce = fmt.Sprintf("%d", nonceInt)

	sig, err := Sign(fmt.Sprintf("%s%s%s", address, amount, nonce), priv_key)
	if err != nil {
		return err
	}

	//fmt.Printf("%s=%s:%s:%s:%s:%s\n", string("set_bal"), address, amount, nonce, public_key, sig)
	btr, err := cl.BroadcastTxCommit(types.Tx(
		fmt.Sprintf("%s=%s:%s:%s:%s:%s", string("set_bal"), address, amount, nonce, public_key, sig)))

	if err != nil {
		return err
	}
	client.WaitForHeight(cl, btr.Height+1, nil)

	return nil
}

/**
based on the address, return the send history of such address. 
For example, such address has sent out tokens 5 times, and the trx details will be returned.

Parameter	Type	Default	Required	Description
query		string	""	true		Query
prove		bool	false	false		Include proofs of the transactions inclusion in the block
page		int	1	false		Page number (1-based)
per_page	int	30	false		Number of entries per page (max: 100)

TxSearch API detail:
https://tendermint.com/rpc/#txsearch
*/
func GetHistorySend(ip, port, address string, prove bool, page, perPage int) (*ctypes.ResultTxSearch, error) {
	if page == 0 {
		page = 1
	}

	if (perPage == 0) {
		perPage = 30
	}

	cl := getHTTPClient(ip, port)

	//curl "localhost:26657/tx_search?query=\"app.fromaddress='B508ED0D54597D516A680E7951F18CAD24C7EC9F'\"&prove=true"
	query := "app.fromaddress="+"'"+address+"'"
	btr, err := cl.TxSearch(query, prove, page, perPage)
	if err != nil {
		return nil, err
	} 

	return btr, err
}

/**
based on the address, return the receive history of such address. 
For example, such address has received tokens 5 times, and the trx details will be returned.

Parameter	Type	Default	Required	Description
query		string	""	true		Query
prove		bool	false	false		Include proofs of the transactions inclusion in the block
page		int	1	false		Page number (1-based)
per_page	int	30	false		Number of entries per page (max: 100)

TxSearch API detail:
https://tendermint.com/rpc/#txsearch
*/
func GetHistoryReceive(ip, port, address string, prove bool, page, perPage int) (*ctypes.ResultTxSearch, error) {
	if page == 0 {
		page = 1
	}

	if (perPage == 0) {
		perPage = 30
	}

	cl := getHTTPClient(ip, port)

	//curl "localhost:26657/tx_search?query=\"app.toaddress='B508ED0D54597D516A680E7951F18CAD24C7EC9F'\"&prove=true"
	query := "app.toaddress="+"'"+address+"'"
	btr, err := cl.TxSearch(query, prove, page, perPage)
	if err != nil {
		return nil, err
	} 

	return btr, err
}

/* helper functions */
/*-------------------------------------------------------------------------------------------------*/
func getHTTPClient(ip, port string) *client.HTTP {
	return client.NewHTTP(ip+":"+port, "/websocket")
}

func deserilizePrivKey(priv_key_b64 string) (ed25519.PrivKeyEd25519, error) {
	kDec, err := base64.StdEncoding.DecodeString(priv_key_b64)
	if err != nil {
		return ed25519.PrivKeyEd25519{}, err
	}

	pp := []byte(kDec)
	keyObject := ed25519.PrivKeyEd25519{pp[0], pp[1], pp[2], pp[3], pp[4], pp[5], pp[6], pp[7], pp[8], pp[9],
		pp[10], pp[11], pp[12], pp[13], pp[14], pp[15], pp[16], pp[17], pp[18], pp[19], pp[20], pp[21],
		pp[22], pp[23], pp[24], pp[25], pp[26], pp[27], pp[28], pp[29], pp[30], pp[31], pp[32], pp[33],
		pp[34], pp[35], pp[36], pp[37], pp[38], pp[39], pp[40], pp[41], pp[42], pp[43], pp[44], pp[45],
		pp[46], pp[47], pp[48], pp[49], pp[50], pp[51], pp[52], pp[53], pp[54], pp[55], pp[56], pp[57],
		pp[58], pp[59], pp[60], pp[61], pp[62], pp[PrivKeyEd25519Size-1]}

	return keyObject, nil
}

func deserilizePubKey(pub_key_b64 string) (ed25519.PubKeyEd25519, error) {
	pDec, err := base64.StdEncoding.DecodeString(pub_key_b64)
	if err != nil {
		return ed25519.PubKeyEd25519{}, err
	}

	pk := []byte(pDec)
	var pubObject ed25519.PubKeyEd25519 = ed25519.PubKeyEd25519{pk[0], pk[1], pk[2], pk[3],pk[4], pk[5],pk[6],
		pk[7],pk[8], pk[9], pk[10], pk[11], pk[12], pk[13], pk[14], pk[15], pk[16], pk[17], pk[18], pk[19],
		pk[20], pk[21],pk[22], pk[23],pk[24], pk[25],pk[26], pk[27],pk[28], pk[29],pk[30], pk[PubKeyEd25519Size - 1]}

	return pubObject, nil
}
