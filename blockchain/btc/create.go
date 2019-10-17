package btc

import (
    "bytes"
    "encoding/hex"
    "github.com/btcsuite/btcd/chaincfg"
    "github.com/btcsuite/btcd/chaincfg/chainhash"
    "github.com/btcsuite/btcd/txscript"
    "github.com/btcsuite/btcd/wire"
    "github.com/btcsuite/btcutil"
)

//import (
//    "bytes"
//    "encoding/hex"
//    "github.com/prettymuchbryce/hellobitcoin/base58check"
//    "go-bitcoin-multisig/btcutils"
//    "log"
//    "github.com/btcsuite/btcd/btcjson"
//    "github.com/prettymuchbryce/hellobitcoin/base58check/base58"
//    "math/big"
//    "strconv"
//)
//
//
//var decodeMap [256]byte
//
//func  generateTransaction(private, txHash string){
//    privateKey :=Decode(private)
//    publicKey, err := btcutils.NewPublicKey(privateKey)
//    if err != nil {
//        log.Fatal(err)
//    }
//    publicKeyHash, err := btcutils.Hash160(publicKey)
//    if err != nil {
//        log.Fatal(err)
//    }
//    tempScriptSig, err := btcutils.NewP2PKHScriptPubKey(publicKeyHash)
//    if err != nil {
//        log.Fatal(err)
//    }
//
//
//}
//
//type CorruptInputError int64
//
//func (e CorruptInputError) Error() string {
//    return "illegal base58 data at input byte " + strconv.FormatInt(int64(e), 10)
//}
//
//
//func DecodeToBig(src []byte) (*big.Int, error) {
//    n := new(big.Int)
//    radix := big.NewInt(58)
//    for i := 0; i < len(src); i++ {
//        b := decodeMap[src[i]]
//        if b == 0xFF {
//            return nil, CorruptInputError(i)
//        }
//        n.Mul(n, radix)
//        n.Add(n, big.NewInt(int64(b)))
//    }
//    return n, nil
//}
//
//
//func Decode(value string) []byte {
//    zeroBytes := 0
//    for i := 0; i < len(value); i++ {
//        if value[i] == 49 {
//            zeroBytes += 1
//        } else {
//            break
//        }
//    }
//
//    publicKeyInt, err := base58.DecodeToBig([]byte(value))
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    encodedChecksum := publicKeyInt.Bytes()
//
//    encoded := encodedChecksum[0 : len(encodedChecksum)-4]
//
//    var buffer bytes.Buffer
//    for i := 0; i < zeroBytes; i++ {
//        zeroByte, err := hex.DecodeString("00")
//        if err != nil {
//            log.Fatal(err)
//        }
//        buffer.WriteByte(zeroByte[0])
//    }
//
//    buffer.Write(encoded)
//
//    return buffer.Bytes()[1:len(buffer.Bytes())]
//}

type Transaction struct {
    TxId               string `json:"txid"`
    SourceAddress      string `json:"source_address"`
    DestinationAddress string `json:"destination_address"`
    Amount             int64  `json:"amount"`
    UnsignedTx         string `json:"unsignedtx"`
    SignedTx           string `json:"signedtx"`
}


func CreateTransaction(secret string, destination string, amount int64, txHash string) (Transaction, error) {
    var transaction Transaction
    wif, err := btcutil.DecodeWIF(secret)
    if err != nil {
        return Transaction{}, err
    }
    // get the public key script information for both the sender and the recipient
    addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), &chaincfg.MainNetParams)
    sourceTx := wire.NewMsgTx(wire.TxVersion)
    sourceUtxoHash, _ := chainhash.NewHashFromStr(txHash)
    sourceUtxo := wire.NewOutPoint(sourceUtxoHash, 0)
    sourceTxIn := wire.NewTxIn(sourceUtxo, nil, nil)
    destinationAddress, err := btcutil.DecodeAddress(destination, &chaincfg.MainNetParams)
    sourceAddress, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), &chaincfg.MainNetParams)
    if err != nil {
        return Transaction{}, err
    }
    destinationPkScript, _ := txscript.PayToAddrScript(destinationAddress)
    sourcePkScript, _ := txscript.PayToAddrScript(sourceAddress)
    sourceTxOut := wire.NewTxOut(amount, sourcePkScript)
    sourceTx.AddTxIn(sourceTxIn)
    sourceTx.AddTxOut(sourceTxOut)
    sourceTxHash := sourceTx.TxHash()
    redeemTx := wire.NewMsgTx(wire.TxVersion)
    prevOut := wire.NewOutPoint(&sourceTxHash, 0)
    redeemTxIn := wire.NewTxIn(prevOut, nil, nil)
    redeemTx.AddTxIn(redeemTxIn)
    redeemTxOut := wire.NewTxOut(amount, destinationPkScript)
    redeemTx.AddTxOut(redeemTxOut)
    sigScript, err := txscript.SignatureScript(redeemTx, 0, sourceTx.TxOut[0].PkScript, txscript.SigHashAll, wif.PrivKey, false)
    if err != nil {
        return Transaction{}, err
    }
    redeemTx.TxIn[0].SignatureScript = sigScript
    flags := txscript.StandardVerifyFlags
    vm, err := txscript.NewEngine(sourceTx.TxOut[0].PkScript, redeemTx, 0, flags, nil, nil, amount)
    if err != nil {
        return Transaction{}, err
    }
    if err := vm.Execute(); err != nil {
        return Transaction{}, err
    }
    var unsignedTx bytes.Buffer
    var signedTx bytes.Buffer
    sourceTx.Serialize(&unsignedTx)
    redeemTx.Serialize(&signedTx)
    transaction.TxId = sourceTxHash.String()
    transaction.UnsignedTx = hex.EncodeToString(unsignedTx.Bytes())
    transaction.Amount = amount
    transaction.SignedTx = hex.EncodeToString(signedTx.Bytes())
    transaction.SourceAddress = sourceAddress.EncodeAddress()
    transaction.DestinationAddress = destinationAddress.EncodeAddress()
    return transaction, nil
}
//https://github.com/M-AMAIRI/BlockChain-Go-Wallet/blob/master/BitCTWG.go