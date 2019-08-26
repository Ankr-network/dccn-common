# wallet-sdk
1. You'll need installed nodejs,see https://nodejs.org/en/
2. Install the build tool babel,see https://babeljs.io
3. Install the dependencies and run npm install
4. Run the babel src -d dist/src and babel sample -d dist/sample command to compile es6
6. Run node dist/sample/privkey_to_keystore _sample.js will generate the keyStore.txt file locally
7. Run node dist/sample/keystore_to_privkey_sample.js, load the local keyStore.txt file to generate the private key.
Test with a browser
First run npm install --global webpack to install webpack
Then run webpack src/*.js tests/index.js -o tests/main.js compile and package javascript
Finally run npm run-script start to open http://localhost:9080 in the browser.
We can click test in the web