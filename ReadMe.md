# greet 单体项目

# shorturl 短链微服务

# book 图书查阅系统

# datacenter 数据中台

# mall 商城微服务

# 代码生成

## 1.生成api
```
1、api目录生成一个api模板
goctl api -o xxx.api
2、修改模板内容
3、根据api文件生成代码
goctl api go -api xxx.api -dir .

```
## 2.生成rpc
```
1.rpc目录下新建proto文件
goctl rpc template -o xxx.proto

2.根据proto文件生成代码
goctl rpc proto -src xxx.proto -dir .

```

## 3.生成model
```
1.model目录下新建sql文件

2.数据库执行sql建表

3.生成代码
-c表示使用 redis cache
goctl model mysql ddl -c -src xxx.sql -dir .
```
# 项目目录示例
```
mall // 工程名称
├── common // 通用库
│   ├── randx
│   └── stringx
├── go.mod
├── go.sum
└── service // 服务存放目录
    ├── afterSale
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── cart
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── order
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── pay
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── product
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    └── user
        ├── cmd
        │   ├── api
        │   ├── cronjob
        │   ├── rmq
        │   ├── rpc
        │   └── script
        └── model
```

# api目录
```
.
├── etc
│   └── greet-api.yaml              // 配置文件
├── go.mod                          // mod文件
├── greet.api                       // api描述文件
├── greet.go                        // main函数入口
└── internal                        
    ├── config  
    │   └── config.go               // 配置声明type
    ├── handler                     // 路由及handler转发
    │   ├── greethandler.go
    │   └── routes.go
    ├── logic                       // 业务逻辑
    │   └── greetlogic.go
    ├── middleware                  // 中间件文件
    │   └── greetmiddleware.go
    ├── svc                         // logic所依赖的资源池
    │   └── servicecontext.go
    └── types                       // request、response的struct，根据api自动生成，不建议编辑
        └── types.go
```
# rpc目录
```
.
├── etc             // yaml配置文件
│   └── greet.yaml
├── go.mod
├── greet           // pb.go文件夹①
│   └── greet.pb.go
├── greet.go        // main函数
├── greet.proto     // proto 文件
├── greetclient     // call logic ②
│   └── greet.go
└── internal        
    ├── config      // yaml配置对应的实体
    │   └── config.go
    ├── logic       // 业务代码
    │   └── pinglogic.go
    ├── server      // rpc server
    │   └── greetserver.go
    └── svc         // 依赖资源
        └── servicecontext.go
```
