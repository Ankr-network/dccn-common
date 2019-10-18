package main

import (
    "fmt"
    "encoding/json"
    "github.com/Ankr-network/dccn-common/blockchain/btc"
)

func main() {
   transaction, err := btc.CreateTransaction("cQpb8dJ3B36c25HGBSSDkzu4d8CRzP1tbWJg2vsWyAiEfetQPtdX", "2MytUZedXFVTCKXrgRc2cJJf8NTaurMegdA", 1, "959ee3e0971d12c9449d51d29b26cdd198821b899a6cb37ce3b51a2294bd6c8e")
   if err != nil {
       fmt.Println(err)
       return
   }
   data, _ := json.Marshal(transaction)
   fmt.Println(string(data))
}
//wisdom simple consider flip faculty live shove liar dune remove dial toe

//const (
//    SERVER_HOST        = "47.75.70.201"
//    SERVER_PORT        = 18332
//    USER               = ""
//    PASSWD             = ""
//    USESSL             = false
//    WALLET_PASSPHRASE  = ""
//    WALLET_PASSPHRASE2 = ""
//)
//
//func main(){
//    bc, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
//    if err != nil {
//        log.Fatalln(err)
//    }
//    tx,err:=bc.SendFrom("","",0.00001,1,"","")
//    if err!=nil{
//         log.Printf("%s",err)
//    }
//    fmt.Print(tx)
//
//}