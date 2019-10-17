import {get_key} from "../src/gen_key"

export const get_key_test =()=>{
    get_key("km",(option)=>{
        console.log(option.addr)
    })
}

get_key_test()