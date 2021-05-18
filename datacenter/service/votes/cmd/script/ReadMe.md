```
goctl rpc template -o=votes.proto
goctl rpc proto -src votes.proto -dir .
```
# 生成model
```
goctl model mysql datasource -url="root:123@tcp(192.168.160.128:3306)/gozero-datacenter" -table="app_enroll" -dir ./model -c
goctl model mysql datasource -url="root:123@tcp(192.168.160.128:3306)/gozero-datacenter" -table="app_votes" -dir ./model -c
goctl model mysql datasource -url="root:123@tcp(192.168.160.128:3306)/gozero-datacenter" -table="app_votes_activity" -dir ./model -c
```