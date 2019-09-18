import nacl from "tweetnacl";
import { get_nonce } from "./get_nonce";
import { sha256 } from 'js-sha256';

let bc_addr = '//localhost:26657/broadcast_tx_commit';
let bc_addr_query = '//localhost:26657/abci_query';
export const set_blockchain_addr = (blockchain_addr) => {
    bc_addr = '//' + blockchain_addr + '/broadcast_tx_commit';
    bc_addr_query = '//' + blockchain_addr + '/store/account';
}

const hex2bytes = (hex) => {
    let bytes = [];
    for (var i = 0; i < hex.length - 1; i += 2) {
        bytes.push(parseInt(hex.substr(i, 2), 16));
    }

    return bytes
}

export const send_coin = async (from, to, amount, private_key) => {

    const nonce = await get_nonce(bc_addr_query, from); //TODO check balance is FAIL.
    if (nonce === "") {
        throw "get balance fail"
    }

    // handle nonce
    let nonceInt = parseInt(nonce); //TODO check parseInt fail.
    nonceInt++

    // for test
    //nonceInt = 10

    // handle hash
    const sha245_sum = sha256(from + to + amount + nonceInt);
    let arr = new Uint8Array(32);
    arr.set(hex2bytes(sha245_sum));

    // handle priv key
    let raw = atob(private_key); //TODO check DOMException
    const rawLength = raw.length;
    let pk_array = new Uint8Array(new ArrayBuffer(rawLength));
    for (let i = 0; i < rawLength; i++) {
        pk_array[i] = raw.charCodeAt(i);
    }

    // handle sign
    let sign_array = nacl.sign(arr, pk_array)
    const sign = btoa(String.fromCharCode.apply(null, sign_array.slice(0, 64)));

    await fetch(bc_addr + '?tx="trx_send=' + from + ":" + to + ":" + amount + ":" + nonceInt + ":" + encodeURIComponent(public_key) + ":" + encodeURIComponent(sign) + '"', {
        method: 'GET',
        mode: 'cors',
    }).then(res => {
        return res.json();
    }).then(json => {
        if (json['error']) {
            throw json['error'];
        }
        return json.result;
    }).catch(err => {
        console.log('请求错误', err);
    });
}
