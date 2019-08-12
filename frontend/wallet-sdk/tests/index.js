import {gen_key} from '../src/gen_key';
import {get_balance, set_blockchain_addr as set_balance_bc_addr} from '../src/get_balance';
import {
    get_history,
    get_recv_history,
    get_send_history,
    set_blockchain_addr as set_history_bc_addr
} from '../src/get_history';
import {send_coin, set_blockchain_addr as set_coin_bc_addr} from '../src/send_coin';
import {encryptDataV3} from '../src/encryptDataV3'
import btoa from "btoa"
import {decryptDataV3} from '../src/decryptDataV3'

let element = null;

const echo =
    (text) => {
        element.appendChild(document.createElement('br'))
        if (text != null) {
            element.appendChild(document.createTextNode(text))
        }
    }

document.getElementById('gen_keys_btn')
    .onclick =
    async() => {
        gen_key(
            'me@example.com', 'brig alert rope welsh foss rang orb', (key) => {
                element = document.getElementById('gen_keys');
                echo();
                echo('public key:');
                echo(key.public_key);
                echo();
                echo('secret key:');
                echo(key.private_key);
                echo();
                echo('wallet address:');
                echo(key.addr);
                echo();
            });
    }

document.getElementById('get_balance_btn')
    .onclick =
    async() => {
        set_balance_bc_addr('chain-stage.dccn.ankr.com:26657')

        const balance =
            await get_balance('B508ED0D54597D516A680E7951F18CAD24C7EC9FCFCD67');
        element = document.getElementById('get_balance');

        echo();
        echo('balance:');
        echo(balance);
    }

document.getElementById('send_coin_btn')
    .onclick =
    async() => {
        set_coin_bc_addr('chain-stage.dccn.ankr.com:26657')

        await send_coin(
            'B508ED0D54597D516A680E7951F18CAD24C7EC9FCFCD67',
            '0D9FE6A785C830D2BE66FE40E0E7FE3D9838456CE15D2C', '88000000000000000000',
            'wmyZZoMedWlsPUDVCOy+TiVcrIBPcn3WJN8k5cPQgIvC8cbcR10FtdAdzIlqXQJL9hBw1i0RsVjF6Oep/06Ezg==',
            'wvHG3EddBbXQHcyJal0CS/YQcNYtEbFYxejnqf9OhM4=');

        set_balance_bc_addr('chain-stage.dccn.ankr.com:26657')
        const balance =
            await get_balance('B508ED0D54597D516A680E7951F18CAD24C7EC9FCFCD67');
        element = document.getElementById('send_coin');

        echo();
        echo('new balance:');
        echo(balance);
    }

document.getElementById('get_send_history_btn')
    .onclick =
    async() => {
        set_history_bc_addr('chain-stage.dccn.ankr.com:26657')

        const historys =
            await get_send_history('B508ED0D54597D516A680E7951F18CAD24C7EC9FCFCD67');
        element = document.getElementById('get_send_history');

        echo();
        echo('history:');
        historys.forEach((history) => {
            echo('From: ' + history.from_addr)
            echo('To: ' + history.to_addr)
            echo('Amount: ' + history.amount)
            echo('At: ' + history.timestamp)
            echo()
        });
    }

document.getElementById('get_recv_history_btn')
    .onclick =
    async() => {
        set_history_bc_addr('chain-stage.dccn.ankr.com:26657')

        const historys =
            await get_recv_history('0D9FE6A785C830D2BE66FE40E0E7FE3D9838456CE15D2C');
        element = document.getElementById('get_recv_history');

        echo();
        echo('history:');
        historys.forEach((history) => {
            echo('From: ' + history.from_addr)
            echo('To: ' + history.to_addr)
            echo('Amount: ' + history.amount)
            echo('At: ' + history.timestamp)
            echo()
        });
    }

document.getElementById('get_history_btn')
    .onclick = async() => {
    set_history_bc_addr('chain-stage.dccn.ankr.com:26657')

    const historys =
        await get_history('B508ED0D54597D516A680E7951F18CAD24C7EC9FCFCD67');
    element = document.getElementById('get_history');

    echo();
    echo('history:');
    historys.forEach((history) => {
        echo('From: ' + history.from_addr)
        echo('To: ' + history.to_addr)
        echo('Amount: ' + history.amount)
        echo('At: ' + history.timestamp)
        echo()
    });
}
document.getElementById('encryptDataV3_btn')
    .onclick = async() => {
    // prettier-ignore
    var privateKey = btoa(String.fromCharCode.apply(null, [236, 215, 73, 70, 33, 204, 45, 186, 221, 182, 81, 203, 14, 120, 214, 81, 148, 80, 175, 91, 223, 39, 36, 238, 236, 7, 143, 255, 7, 0, 91, 115, 191, 149, 132, 37, 81, 154, 197, 30, 227, 238, 142, 29, 24, 153, 187, 38, 146, 108, 192, 137, 182, 81, 183, 56, 232, 111, 91, 133, 199, 159, 145, 19]));

    var password = '12345678qw';
    var address = '7541C223B7DE30B611A26E1F794F905EC10E3022AE33E9';
    var publickey = 'hJQKACadsobArr0SXSqEhyC8dwTlkkRLu3dt9UQtDwc=';
// prettier-ignore
    var salt = new Uint8Array([83, 174, 68, 194, 13, 87, 23, 31, 61, 121, 2, 24, 170, 3, 101, 35, 221, 180, 65, 45, 85, 218, 220, 9, 28, 222, 121, 67, 90, 195, 124, 224]);
// prettier-ignore
    var iv = new Uint8Array([37, 8, 192, 48, 234, 198, 43, 102, 211, 83, 231, 250, 78, 104, 219, 236]);
    var kestore = await encryptDataV3(privateKey, password, address, publickey, salt, iv)
    document.getElementById('decryptDataV3_text').innerHTML="decryptDataV3 :"+JSON.stringify(kestore)
}

document.getElementById('decryptDataV3_btn')
    .onclick = async() => {
    var keystore = {
        "address": "7541C223B7DE30B611A26E1F794F905EC10E3022AE33E9",
        "publickey": "hJQKACadsobArr0SXSqEhyC8dwTlkkRLu3dt9UQtDwc=",
        "crypto": {
            "cipher": "aes-128-ctr",
            "ciphertext": "edaa547056785b16d0224a880a15902f70c0628f7add425061bcf071e79dd206273682390fafa28ef4edf5d546b2dfc70a24c6ba728e8718438bcb1340ad5b461aa1fc038287ea4074a889a1114b688066d2ace440bc9c8e",
            "cipherparams": {
                "iv": "2508c030eac62b66d353e7fa4e68dbec"
            },
            "kdf": "scrypt",
            "kdfparams": {
                "dklen": 32,
                "n": 262144,
                "p": 1,
                "r": 8,
                "salt": "53ae44c20d57171f3d790218aa036523ddb4412d55dadc091cde79435ac37ce0"
            },
            "mac": "938706762ce85d2681d4eb1ad4003d6298cf4f0a4f3f52b9e34322871aafd0cf"
        },
        "version": 3
    }
    var password = '12345678qw';
    var decry = await decryptDataV3(keystore.crypto,password)
    document.getElementById('encryptDataV3_text').innerHTML="encryptDataV3 :"+JSON.stringify(decry)
}
