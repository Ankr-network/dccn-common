import {encryptDataV3} from "../src/encryptDataV3"
import btoa from "btoa"
import fs from "fs"

const encryptedKeys = {
    address: '7541C223B7DE30B611A26E1F794F905EC10E3022AE33E9',
    publickey: 'hJQKACadsobArr0SXSqEhyC8dwTlkkRLu3dt9UQtDwc=',
    crypto: {
        cipher: 'aes-128-ctr',
        ciphertext: 'edaa547056785b16d0224a880a15902f70c0628f7add425061bcf071e79dd206273682390fafa28ef4edf5d546b2dfc70a24c6ba728e8718438bcb1340ad5b461aa1fc038287ea4074a889a1114b688066d2ace440bc9c8e',
        cipherparams: {
            iv: '2508c030eac62b66d353e7fa4e68dbec'
        },
        kdf: 'scrypt',
        kdfparams: {
            dklen: 32,
            n: 262144,
            p: 1,
            r: 8,
            salt: '53ae44c20d57171f3d790218aa036523ddb4412d55dadc091cde79435ac37ce0'
        },
        mac: '938706762ce85d2681d4eb1ad4003d6298cf4f0a4f3f52b9e34322871aafd0cf'
    },
    version: 3
};

// prettier-ignore
var privateKey = btoa(String.fromCharCode.apply(null, [236, 215, 73, 70, 33, 204, 45, 186, 221, 182, 81, 203, 14, 120, 214, 81, 148, 80, 175, 91, 223, 39, 36, 238, 236, 7, 143, 255, 7, 0, 91, 115, 191, 149, 132, 37, 81, 154, 197, 30, 227, 238, 142, 29, 24, 153, 187, 38, 146, 108, 192, 137, 182, 81, 183, 56, 232, 111, 91, 133, 199, 159, 145, 19]));

var password = '12345678qw';
var address = '7541C223B7DE30B611A26E1F794F905EC10E3022AE33E9';
var publickey = 'hJQKACadsobArr0SXSqEhyC8dwTlkkRLu3dt9UQtDwc=';
// prettier-ignore
var salt = new Uint8Array([83, 174, 68, 194, 13, 87, 23, 31, 61, 121, 2, 24, 170, 3, 101, 35, 221, 180, 65, 45, 85, 218, 220, 9, 28, 222, 121, 67, 90, 195, 124, 224]);
// prettier-ignore
var iv = new Uint8Array([37, 8, 192, 48, 234, 198, 43, 102, 211, 83, 231, 250, 78, 104, 219, 236]);

const encryptDataV3_test =()=>{
    return encryptDataV3(privateKey, password, address, publickey, salt, iv)
        .then((data)=>{
            console.log(data)
            fs.writeFileSync("keystore.txt",JSON.stringify(data),"utf-8");
        })
}

encryptDataV3_test()