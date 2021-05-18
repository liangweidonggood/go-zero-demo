# greet 单体项目

# shorturl 短链微服务

# book 图书查阅系统

# 尼玛

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
