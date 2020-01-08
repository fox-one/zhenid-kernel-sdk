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