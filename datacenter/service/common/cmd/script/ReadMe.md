# 公共服务
```

goctl rpc proto -src common.proto -dir .

goctl model mysql datasource -url="root:123@tcp(192.168.160.128:3306)/gozero-datacenter" -table="base_app" -dir ./model -c

goctl model mysql datasource -url="root:123@tcp(192.168.160.128:3306)/gozero-datacenter" -table="app_config" -dir ./model -c


```
