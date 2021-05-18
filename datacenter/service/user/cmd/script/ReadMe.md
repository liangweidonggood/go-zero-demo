### 用户服务，含登陆


### 创建api 文件
```shell 生成接口
goctl api -o user.api
```
### 生成user api服务
```
goctl api go -api user.api -dir .

goctl rpc proto -src user.proto -dir .
```
### 生成model

```
goctl model mysql datasource -url="root:123@tcp(192.168.160.128:3306)/gozero-datacenter" -table="app_user" -dir ./model -c
goctl model mysql datasource -url="root:123@tcp(192.168.160.128:3306)/gozero-datacenter" -table="base_member" -dir ./model -c
```
