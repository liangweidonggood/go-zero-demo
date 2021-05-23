# 部署说明

## 1 windows下生成linux二进制文件
```
set CGO_ENABLED=0
set GOOS=linux
go build -ldflags="-s -w" -installsuffix cgo -o main .

GOOS：目标系统为linux
CGO_ENABLED：默认为1，启用C语言版本的GO编译器，通过设置成0禁用它
GOARCH：32位系统为386，64位系统为amd64
-ldflags：用于传递每个go工具链接调用的参数。
    -s: 省略符号表和调试信息
    -w: 省略DWARF符号表
-installsuffix：在软件包安装的目录中增加后缀标识，用于区分默认版本
-o：指定编译后的可执行文件名称
```
## 2 编写Dockerfile
```
FROM scratch
WORKDIR /app

构建基本镜像
docker build -t go-scratch:1.0.0 .
```
## 3 编写docker-compose.yml
```
version: "3"
services:
  greet:
    image: go-scratch:1.0.0
    container_name: greet
    restart: always
    ports:
      - 8888:8888
    volumes:
    - /home/project/greet/etc/greet-api.yaml:/app/etc/greet-api.yaml
    - /home/project/greet/logs:/app/logs
    - /home/project/greet/main:/app/main
    - /etc/timezone:/etc/timezone:ro
    - /etc/localtime:/etc/localtime:ro
    command: ./main -f etc/greet-api.yaml
```
## 4 上传至服务器构建镜像并运行docker-compose
```
重要
chmod +x main

docker-compose up -d

```

## 5 测试
```
curl -i http://192.168.230.129:8888/from/you

HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 21 May 2021 05:02:01 GMT
Content-Length: 28

{"message":"hello,go-zero!"}

```
