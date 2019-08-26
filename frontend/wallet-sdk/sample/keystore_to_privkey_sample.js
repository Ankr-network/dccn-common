import {keystoreToPrivkey} from "../src/keystore_to_privkey"
import btoa from "btoa"
import fs from "fs"

const password = '12345678qw';
const wrongPassword = '12345678qw1';
function keystoreToPrivkeyTest(){
    var keystore =fs.readFileSync("keystore.txt",'utf-8')
    keystore=JSON.parse(keystore)
    return keystoreToPrivkey(keystore.crypto,password).then((result)=>{
        console.log(JSON.stringify(result))
    })
}

keystoreToPrivkeyTest()