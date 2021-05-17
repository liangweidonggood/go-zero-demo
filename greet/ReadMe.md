# 生成项目

## 工具安装
```
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero

GOPROXY=https://goproxy.cn/,direct go install github.com/tal-tech/go-zero/tools/goctl@latest

```
## 开始
```
mkdir go-zero-demo && cd go-zero-demo
go mod init go-zero-demo
```

## 生成一个单体项目
```
goctl api new greet
cd greet
go mod tidy
go run greet.go -f etc/greet-api.yaml

```
## 测试
```
curl -i http://localhost:8888/from/you

HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 22 Oct 2020 14:03:18 GMT
Content-Length: 14
{"message":""}
```
## 项目结构
```
|-- etc
|   `-- greet-api.yaml
|-- greet.api
|-- greet.go
`-- internal
    |-- config
    |   `-- config.go
    |-- handler
    |   |-- greethandler.go
    |   `-- routes.go
    |-- logic
    |   `-- greetlogic.go
    |-- svc
    |   `-- servicecontext.go
    `-- types
        `-- types.go

```
## 编写逻辑
```
/go-zero-demo/greet/internal/logic/greetlogic.go
func (l *GreetLogic) Greet(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line
    
	return &types.Response{}, nil
}

```




