import scrypt from 'scrypt-js';
import { Counter, ModeOfOperation } from 'aes-js';
import keccak256 from 'keccak256';

const keystoreToPrivkeyErrors = {
    CYPHER_IS_NOT_SUPPORTED: 'CYPHER_IS_NOT_SUPPORTED',
    WRONG_PASSWORD: 'WRONG_PASSWORD',
}
const SUPPORTED_CYPHER_NAME = 'aes-128-ctr';

export const keystoreToPrivkey =(encryptedPrivateKey,auth) =>{
    if (encryptedPrivateKey.cipher !== SUPPORTED_CYPHER_NAME) {
        throw new Error(decryptDataV3Errors.CYPHER_IS_NOT_SUPPORTED);
    }
    const cypherText = Uint8Array.from(
        Buffer.from(encryptedPrivateKey.ciphertext, 'hex'),
    );
    const mac = Uint8Array.from(Buffer.from(encryptedPrivateKey.mac, 'hex'));
    const iv = Uint8Array.from(
        Buffer.from(encryptedPrivateKey.cipherparams.iv, 'hex'),
    );
    return getKDFKey(encryptedPrivateKey, auth).then(key => {
        const encryptKey = key.slice(0, 16);

        const counter = new Counter(5);
        counter.setBytes(iv);
        const cbc = new ModeOfOperation.ctr(encryptKey, counter);
        const privateKey = cbc.decrypt(cypherText);

        const calculatedMac = new Uint8Array(
            keccak256(
                Buffer.from([
                    ...Array.from(key.slice(16, 32)),
                    ...Array.from(cypherText),
                ]),
            ),
        );

        if (Buffer.compare(Buffer.from(mac), Buffer.from(calculatedMac)) !== 0) {
            throw new Error(decryptDataV3Errors.WRONG_PASSWORD);
        }

        return {
            privateKey,
        };
    });
}

const getKDFKey =(encryptedPrivateKey, auth)=> {
    const authArray = Buffer.from(auth, 'utf8');
    const salt = Buffer.from(encryptedPrivateKey.kdfparams.salt, 'hex');
    const dkLen = encryptedPrivateKey.kdfparams.dklen;

    const n = encryptedPrivateKey.kdfparams.n;
    const r = encryptedPrivateKey.kdfparams.r;
    const p = encryptedPrivateKey.kdfparams.p;

    return new Promise((resolve, reject) => {
        scrypt(authArray, salt, n, r, p, dkLen, (error, progress, key)=> {
            if (error) {
                reject(error);
            } else if (key) {
                resolve(key);
            } else {
                // progress
            }
        });
    });
}