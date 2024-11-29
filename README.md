# Go 
## 交叉编译
在mac上进行交叉编译
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o ./out/app .   //linux
CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build -o ./out/app.exe . //windows    

在windows上进行交叉编译
$Env:CGO_ENABLE=0;$Env:GOARCH="amd64";$Env:GOOS="linux" //linux
go build -o ./out/app .

$Env:CGO_ENABLE=0;$Env:GOARCH="amd64";$Env:GOOS="darwin"    //mac
go build -o ./out/app .

what i learn
如何创建 go 项目；
如何使用 VSCode 的内置命令行；
如何使用 go run 命令来运行 Go 程序；
如何使用 git 做版本控制；
Go 应用中 package main 的规则；
标准库 fmt 包的基本使用；
标准库 http 包的基本使用；
http 包中如何通过 url 路径来处理业务逻辑；
如何使用 air 来自动重载代码；
如何使用 go proxy 来加速 go get 命令；
如何添加 http 标头；
如何在不需要联网的情况下查看 Go 文档；
如何在 Go 文档中快速定位所需内容。
路由与 Go Modules 的使用。
添加sql驱动
连接数据库