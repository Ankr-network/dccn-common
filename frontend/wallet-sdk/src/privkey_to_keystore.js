import scrypt from 'scrypt-js';
import unorm from 'unorm';
import keccak256 from 'keccak256';
import {Counter, ModeOfOperation} from 'aes-js';
export const privkeyToKeystore = (data, auth, address, publickey, name) => {
    const salt = getRandomArray(32);
    const iv = getRandomArray(16);
    return privkeyToKeystoreDeterministic(data, auth, address, publickey, salt, iv, name)
}

const getRandomArray = (length) => {
    return new Uint8Array(
        Array.from({length}, () => Math.floor(Math.random() * 256)),
    );
};


export const privkeyToKeystoreDeterministic = (data, auth, address, publickey, salt, iv, name)=> {
    const n = 262144;
    const r = 8;
    const p = 1;
    const dklen = 32;
    const authBytes = convertStringToBytes(unorm.nfkc(auth));
    const dataBytes = convertStringToBytes(data);
    return new Promise((resolve, reject) => {
        scrypt(authBytes, salt, n, r, p, dklen, (error, progress, key) => {
            if (error) {
                reject(error);
            } else if (key) {
                const encryptKey = key.slice(0, 16);
                const counter = new Counter(5);
                counter.setBytes(iv);
                const cbc = new ModeOfOperation.ctr(encryptKey, counter);
                const cipherText = cbc.encrypt(dataBytes);

                const mac = new Uint8Array(
                    keccak256(
                        Buffer.from([...key.slice(16, 32), ...Array.from(cipherText)]),
                    ),
                );

                resolve({
                    name,
                    address,
                    publickey,
                    crypto: {
                        cipher: 'aes-128-ctr',
                        ciphertext: toHexString(cipherText),
                        cipherparams: {
                            iv: toHexString(iv),
                        },
                        kdf: 'scrypt',
                        kdfparams: {
                            dklen: dklen,
                            n,
                            p,
                            r,
                            salt: toHexString(salt),
                        },
                        mac: toHexString(mac),
                    },
                    version: 3,
                });
            } else {
                // progress
            }
        });
    });
}


const convertStringToBytes = (str)=> {
    var buf = new ArrayBuffer(str.length);
    var bufView = new Uint8Array(buf);
    for (var i = 0, strLen = str.length; i < strLen; i++) {
        bufView[i] = str.charCodeAt(i);
    }
    return bufView;
}

const toHexString = (byteArray)=> {
    return Array.prototype.map
        .call(byteArray, (byte) => {
            return ('0' + (byte & 0xff).toString(16)).slice(-2);
        })
        .join('');
}
