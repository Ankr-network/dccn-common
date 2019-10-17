import {send_coin,set_blockchain_addr} from "../src/send_coin"

function send_blockchain() {
    set_blockchain_addr("chain-dev.dccn.ankr.com:26657")
    send_coin("17FE9087DF26BF874B40820201911E37D9BE1FEEFA1A19","03F02BC04AF233907C0578D31D8DCD6F0872414A2A450F",10,"0mqsOtVueE7uq/I5J/dAhesumWXTu619xXuRgtj4l0d0ELMH6X9ZjGqT6Lnhrhp13LVeGIgrm3QgBnk4q16BZg==")

}

send_blockchain()