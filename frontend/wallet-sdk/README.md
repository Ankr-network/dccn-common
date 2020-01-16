# ANKR Wallet JavaScript SDK

## Introduction
This API document introduces the web APIs for third party to integrate as an Ankr Wallet. The basic functionalities include,

* generate an address
* send ankr native tokens from one address to another address
* query balance
* query transaction history
* convert private key to keystore format
* convert keystore format to private key

## Installation

1. You'll need to install nodejs, see https://nodejs.org/en/
2. Install the build tool babel, see https://babeljs.io
3. Install the dependencies and run "npm install"
4. Run the "npx babel src -d dist/src" and "npx babel sample -d dist/sample" command to compile es6

### Test in command line
1. Run "node dist/sample/privkey_to_keystore_sample.js" to generate the keyStore.txt file locally
2. Run "node dist/sample/keystore_to_privkey_sample.js" to load the local keyStore.txt file to generate the private key.

### Test with a browser
1. First run "npm install --global webpack" to install webpack
2. Then run "webpack src/*.js tests/index.js -o tests/main.js" to compile and package javascript
3. Finally run "npm run-script start" and open http://localhost:9080 in the browser.
4. Then you can click buttons in the webpage to test.

## APIs

### gen_key

generate a key pair

* `input`: salt
* `output`: private key, public key, address

### get_balance

get the balance of specific address

* `input`: an wallet address
* `output`: the current balance

### send_coin

send ankr native tokens from one address to another address

* `input`: from address, to address, amount, private key, public key
* `output`: result status

### get_history 

query the transaction history

* `input`: a wallet address
* `output`: the array of transactions.

### privkey_to_keystore

convert private key to keystore format

* `input`: private key, password, wallet address, wallet public key, wallet name
* `output`: keystore string or error

### keystore_to_privkey

convert keystore format to private key

* `input`: keystore string, password
* `output`: private key or error

----------------------------------
Please follow the sample code in the sample and tests folder.
