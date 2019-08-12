import {decryptDataV3} from "../src/decryptDataV3"
import btoa from "btoa"
import fs from "fs"

const expectedPrivateKey = Uint8Array.from(Buffer.from(btoa(String.fromCharCode.apply(null, [236, 215, 73, 70, 33, 204, 45, 186, 221, 182, 81, 203, 14, 120, 214, 81, 148, 80, 175, 91, 223, 39, 36, 238, 236, 7, 143, 255, 7, 0, 91, 115, 191, 149, 132, 37, 81, 154, 197, 30, 227, 238, 142, 29, 24, 153, 187, 38, 146, 108, 192, 137, 182, 81, 183, 56, 232, 111, 91, 133, 199, 159, 145, 19])), 'utf8'));
const password = '12345678qw';
const wrongPassword = '12345678qw1';
function test(){
    var bb =fs.readFileSync("bb.txt",'utf-8')
    bb=JSON.parse(bb)
    return decryptDataV3(bb.crypto,password).then((result)=>{
        console.log(result)
    })
}

test()