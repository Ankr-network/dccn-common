# wallet-sdk

1.安装 nodejs ,详见:https://nodejs.org/en/。
2.安装 babel https://babeljs.io
3.运行 npm install。
4.运行babel src -d dist/src。
5.运行babel sample dist/sample。
6.node encryptDataV3_sample.js，会在本地生成keyStore.txt文件。
7.node decryptDataV3_sample.js, 加载本地的keyStore.txt文件生成私钥。
8.在浏览器上面测试
首先安装npm install --global webpack 然后运行 webpack src/*.js tests/index.js -o tests/main.js
最后运行命令 npm install --global webpack
可以在web中点击测试