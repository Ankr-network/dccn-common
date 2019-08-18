# wallet-sdk

1. 安装 nodejs ,详见:https://nodejs.org/en/
1. 运行 npm install
1. 运行 npx babel src -d dist/src
1. 运行 npx babel sample -d dist/sample
1. node dist/sample/encryptDataV3_sample.js，会在本地生成 keyStore.txt 文件。
1. node dist/sample/decryptDataV3_sample.js, 加载本地的 keyStore.txt 文件生成私钥。
1. 在浏览器上面测试
   运行 npx webpack src/\*.js tests/index.js -o tests/main.js
   最后运行命令 npm install --global webpack
   可以在 web 中点击测试
