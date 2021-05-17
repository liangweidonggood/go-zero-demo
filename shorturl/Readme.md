# 短链微服务
```
短链服务就是将长的 URL 网址，通过程序计算等方式，转换为简短的网址字符串。
```
## 准备工作
```
安装 etcd, mysql, redis

go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2

安装 protoc-3.17.0-win64 并配置环境变量

go install github.com/tal-tech/go-zero/tools/goctl@latest


```
## 创建项目
```
mkdir -p shorturl/api && cd shorturl/api

goctl api -o shorturl.api

```
## 编写代码
```
改变shorturl.api文件
type request {
	// TODO: add members here and delete this comment
}

type response {
	// TODO: add members here and delete this comment
}

service shorturl-api {
	@handler GetUser // TODO: set handler name and delete this comment
	get /users/id/:userId(request) returns(response)

	@handler CreateUser // TODO: set handler name and delete this comment
	post /users/create(request)
}
```
## 生成代码
```
goctl api go -api shorturl.api -dir .

目录
.
|-- etc
|   `-- shorturl-api.yaml           // 配置文件
|-- internal
|   |-- config
|   |   `-- config.go               // 定义配置
|   |-- handler
|   |   |-- expandhandler.go        // 实现 expandHandler
|   |   |-- routes.go               // 定义路由处理
|   |   `-- shortenhandler.go       // 实现 shortenHandler
|   |-- logic
|   |   |-- expandlogic.go          // 实现 ExpandLogic
|   |   `-- shortenlogic.go         // 实现 ShortenLogic
|   |-- svc
|   |   `-- servicecontext.go       // 定义 ServiceContext
|   `-- types
|       `-- types.go                // 定义请求、返回结构体
|-- shorturl.api
`-- shorturl.go                     // main 入口定义

```
## 启动 API Gateway 服务，默认侦听在 8888 端口
```
go run shorturl.go -f etc/shorturl-api.yaml
```
## 测试

```
curl -i "http://localhost:8888/shorten?url=http://www.xiaoheiban.cn"

HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 17 May 2021 05:38:01 GMT
Content-Length: 14

{"shorten":""}

```
## 说明
```
可以修改 internal/svc/servicecontext.go 来传递服务依赖（如果需要）

实现逻辑可以修改 internal/logic 下的对应文件

可以通过 goctl 生成各种客户端语言的 api 调用代码

到这里，你已经可以通过 goctl 生成客户端代码给客户端同学并行开发了，支持多种语言，详见文档
```
## 编写 transform rpc 服务
```
mkdir -p rpc/transform && cd rpc/transform

goctl rpc template -o transform.proto
需要修改
syntax = "proto3";

package transform;

message Request {
  string ping = 1;
}

message Response {
  string pong = 1;
}

service Transform {
  rpc Ping(Request) returns(Response);
}

```
## 生成rpc代码
```
确保
GO111MODULE=on
goctl rpc proto -src transform.proto -dir .

目录
.
|-- etc
|   `-- transform.yaml
|-- internal
|   |-- config
|   |   `-- config.go
|   |-- logic
|   |   |-- expandlogic.go
|   |   `-- shortenlogic.go
|   |-- server
|   |   `-- transformerserver.go
|   `-- svc
|       `-- servicecontext.go
|-- transform
|   `-- transform.pb.go
|-- transform.go
|-- transform.proto
`-- transformer
    `-- transformer.go

```
## 修改transform.yaml
```
Name: transform.rpc
ListenOn: 127.0.0.1:8080
Etcd:
  Hosts:
  - 192.168.230.129:2379
  Key: transform.rpc
```
## 启动项目
```
先修改go.mod
require (
	github.com/golang/protobuf v1.4.2
	github.com/tal-tech/go-zero v1.1.7
	google.golang.org/grpc v1.29.1
)
go mod tidy

go run transform.go -f etc/transform.yaml
```
## 查看etcd里面的服务
```
etcdctl get transform.rpc --prefix

transform.rpc/112442071973726213
127.0.0.1:8080
```
## 修改 API Gateway 代码调用 transform rpc 服务
```
修改 internal/config/config.go 如下，增加 transform 服务依赖
修改 internal/svc/servicecontext.go，如下：
修改 internal/logic/expandlogic.go 里的 Expand 方法，如下：
```
## 通过调用 transformer 的 Expand 方法实现短链恢复到 url
```
修改 internal/logic/shortenlogic.go，如下：

```
## 定义数据库表结构，并生成 CRUD+cache 代码
```
mkdir -p rpc/transform/model
创建sql文件 shorturl.sql
CREATE TABLE `shorturl`
(
    `shorten` varchar(255) NOT NULL COMMENT 'shorten key',
    `url` varchar(255) NOT NULL COMMENT 'original url',
    PRIMARY KEY(`shorten`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
建库建表
create database gozero;

-c表示使用 redis cache
goctl model mysql ddl -c -src shorturl.sql -dir .
也可以用 datasource 命令代替 ddl 来指定数据库链接直接从 schema 生成
rpc/transform/model
├── shorturl.sql
├── shorturlmodel.go              // CRUD+cache 代码
└── vars.go                       // 定义常量和变量

```
## 修改 shorten/expand rpc 代码调用 crud+cache 代码
```
修改 rpc/transform/etc/transform.yaml，增加如下内容：
DataSource: root:123@tcp(192.168.230.129:3306)/gozero
Table: shorturl
Cache:
  - Host: 192.168.230.129:6379
    Pass: "ZZkde@#3d99"
可以使用多个 redis 作为 cache，支持 redis 单点或者 redis 集群
修改 rpc/transform/internal/config/config.go，如下：
修改 rpc/transform/internal/svc/servicecontext.go，如下：
修改 rpc/transform/internal/logic/expandlogic.go，如下：
```
```
注意：

undefined cache，你需要 import "github.com/tal-tech/go-zero/core/stores/cache"
undefined model, sqlx, hash 等，你需要在文件中
import "shorturl/rpc/transform/model"

import "github.com/tal-tech/go-zero/core/stores/sqlx"
```
## 完整调用演示
```
go mod download github.com/go-sql-driver/mysql
go mod download github.com/go-xorm/builder

启动rpc服务
cd shorturl/rpc/transform
go run transform.go -f etc/transform.yaml

启动api服务
cd shorturl/api
go run shorturl.go -f etc/shorturl-api.yaml

shorten api 调用
curl -i "http://localhost:8888/shorten?url=http://www.xiaoheiban.cn"

expand api 调用
curl -i "http://localhost:8888/expand?shorten=f35b2a"

```
