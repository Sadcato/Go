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