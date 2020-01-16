let bc_addr = '//localhost:26657/abci_query';
export const set_blockchain_addr = (blockchain_addr) => {
    bc_addr = '//' + blockchain_addr + '/store/account';
}

export const get_balance = async (wallet_addr) => {
    let balance = 'FAIL';
    await fetch(bc_addr + '?data="bal:' + wallet_addr.toString() + '"', {
        method: 'GET',
        mode: 'cors',
    }).then(res => {
        return res.json();
    }).then(json => {
        if (json['error']) {
            throw json['error'];
        }
        return json.result.response;
    }).then(data => {
        return data.value;
    }).then(val => {
        balance = atob(val).split(':', 1);
    }).catch(err => {
        console.log('request error:', err);
    });

    return balance;
}
