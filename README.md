# ZhenID Kernel SDK


## 编译iOS SDK

安装编译库

```
go get golang.org/x/mobile/cmd/gomobile
```

在工程目录下目录执行

```
gomobile init
gomobile bind -target=ios
```

生成的 framework 是可以在iOS环境下运行的framework

## Client API

* [ ] 生成账户（公私钥账号）
* [ ] 根据账户生成（JWT1，用于访问接入点）
* [ ] 签名校验 （传入账户，校验内容）-> 校验结果
* [ ] 签发ECDSA的JWT2 （传入账户，payload）-> JWT2
* [ ] 对signdata进行签名 （传入账户，需签名数据）-> 签名结果
* [ ] 对signdata进行校验 （传入账户，需校验数据）-> 校验结果
* [ ] file key加密文件 （传入账户，传入文件）-> 加密后文件
* [ ] file key解密文件 （传入账户，传入文件）-> 解密后文件
* [ ] 获取文件加密算法列表 （）-> [AES]
* [ ] 生成文件加密密钥  （加密方式，随机密钥）-> [AES，密钥]
* [ ] 加密文件 （文件，加密密钥，加密方式）-> 加密后文件
* [ ] 使用E（encript key 公钥）加密 （文件密钥，加密方式） -> 加密后内容
* [ ] 生成Transaction (转入用户的公钥，Transaction 对象） -> 生成好的 Tranaction 对象
* [ ] 签名Transaction （Transaction，账户） -> 签名后的Transaction
