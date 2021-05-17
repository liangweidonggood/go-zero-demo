# 情景
假设我们在开发一个商城项目，而开发者小明负责用户模块(user)和订单模块(order)的开发，我们姑且将这两个模块拆分成两个微服务①
- 订单服务(order)提供一个查询接口
- 用户服务(user)提供一个方法供订单服务获取用户信息
# 服务设计分析
根据情景提要我们可以得知，订单是直接面向用户，通过http协议访问数据，
而订单内部需要获取用户的一些基础数据，既然我们的服务是采用微服务的架构设计， 
那么两个服务（user,order）就必须要进行数据交换，服务间的数据交换即服务间的通讯，
到了这里，采用合理的通讯协议也是一个开发人员需要 考虑的事情，可以通过http，rpc等方式来进行通讯，
这里我们选择rpc来实现服务间的通讯，相信这里我已经对"rpc服务存在有什么作用？"已经作了一个比较好的场景描述。
当然，一个服务开发前远不止这点设计分析，我们这里就不详细描述了。从上文得知，我们需要一个
- user rpc
- order api
# 创建mall工程
```
mkdir mall && cd mall

```
# 创建user rpc服务
- 创建user rpc服务
```
mkdir -p user/rpc && cd user/rpc

```
- 添加user.proto文件，增加getUser方法
```
vim ~/go-zero-demo/mall/user/rpc/user.proto
```
```
syntax = "proto3";

  package user;

  message IdRequest {
      string id = 1;
  }

  message UserResponse {
      // 用户id
      string id = 1;
      // 用户名称
      string name = 2;
      // 用户性别
      string gender = 3;
  }

  service User {
      rpc getUser(IdRequest) returns(UserResponse);
  }
```
- 生成代码
```
$ cd ~/go-zero-demo/mall/user/rpc
$ goctl rpc proto -src user.proto -dir .
protoc  -I=/Users/xx/mall/user user.proto --go_out=plugins=grpc:/Users/xx/mall/user/user
Done.
```
- 填充业务逻辑
```
$ vim internal/logic/getuserlogic.go
```
```go
package logic

  import (
      "context"

      "go-zero-demo/mall/user/internal/svc"
      "go-zero-demo/mall/user/user"

      "github.com/tal-tech/go-zero/core/logx"
  )

  type GetUserLogic struct {
      ctx    context.Context
      svcCtx *svc.ServiceContext
      logx.Logger
  }

  func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
      return &GetUserLogic{
          ctx:    ctx,
          svcCtx: svcCtx,
          Logger: logx.WithContext(ctx),
      }
  }

  func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
      return &user.UserResponse{
          Id:   "1",
          Name: "test",
      }, nil
  }
```
# 创建order api服务
- 创建 order api服务
```
$ cd ~/go-zero-demo/mall
$ mkdir -p order/api && cd order/api
```
- 添加api文件
```
$ vim order.api
```
```
type(
      OrderReq {
          Id string `path:"id"`
      }

      OrderReply {
          Id string `json:"id"`
          Name string `json:"name"`
      }
  )
  service order {
      @handler getOrder
      get /api/order/get/:id (OrderReq) returns (OrderReply)
  }
```
- 生成order服务
```
$ goctl api go -api order.api -dir .
  Done.
```
- 添加user rpc配置
```
$ vim internal/config/config.go
```
```go
 package config

  import "github.com/tal-tech/go-zero/rest"
  import "github.com/tal-tech/go-zero/zrpc"

  type Config struct {
      rest.RestConf
      UserRpc zrpc.RpcClientConf
  }
```
- 添加yaml配置
```
$ vim etc/order.yaml
```
```
Name: order
  Host: 0.0.0.0
  Port: 8888
  UserRpc:
    Etcd:
      Hosts:
      - 127.0.0.1:2379
      Key: user.rpc
```
- 完善服务依赖
```
$ vim internal/svc/servicecontext.go
```

```go
package svc

  import (
      "go-zero-demo/mall/order/api/internal/config"
      "go-zero-demo/mall/user/rpc/userclient"

      "github.com/tal-tech/go-zero/zrpc"
  )

  type ServiceContext struct {
      Config  config.Config
      UserRpc userclient.User
  }

  func NewServiceContext(c config.Config) *ServiceContext {
      return &ServiceContext{
          Config:  c,
          UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
      }
  }
```
- 添加order演示逻辑   

给`getorderlogic`添加业务逻辑
  
```
 $ vim ~/go-zero-demo/mall/order/api/internal/logic/getorderlogic.go
```
```go

user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdRequest{
    Id: "1",
})
if err != nil {
    return nil, err
}

if user.Name != "test" {
    return nil, errors.New("用户不存在")
}

return &types.OrderReply{
    Id:   req.Id,
    Name: "test order",
}, nil
```

# 启动服务并验证

- 启动etcd

```
$ etcd
```
- 启动user rp
```
$ go run user.go -f etc/user.yaml
Starting rpc server at 127.0.0.1:8080...
```
- 启动order api
```
$ go run order.go -f etc/order.yaml
Starting server at 0.0.0.0:8888...
```
- 访问order api
```
curl -i -X GET \
http://localhost:8888/api/order/get/1

HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 07 Feb 2021 03:45:05 GMT
Content-Length: 30

{"id":"1","name":"test order"}
```