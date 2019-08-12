import {decryptDataV3} from "../src/decryptDataV3"
import btoa from "btoa"
import fs from "fs"

const password = '12345678qw';
const wrongPassword = '12345678qw1';
function test_decryptDataV3(){
    var keystore =fs.readFileSync("keystore.txt",'utf-8')
    keystore=JSON.parse(keystore)
    return decryptDataV3(keystore.crypto,password).then((result)=>{
        console.log(result)
    })
}

test_decryptDataV3()