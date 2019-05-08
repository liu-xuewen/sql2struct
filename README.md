## 使用方法

1. 直接运行
```sh
go run cli.go  --file model.go --dsn root:root@tcp\(localhost:3306\)/test?charset=utf8 --table user
```

2. 编译运行
```sh
 go build -o mysql2struct cli.go
 ./mysql2struct  -table tt
./mysql2struct -dsn "aa:bb@tcp(192.168.0.0:3323)/tt?charset=utf8&parseTime=True&loc=Local" -table tt
```

3. 参数说明  
```sh
-dsn            string 数据库dsn配置
-enableJsonTag  bool 是否添加json的tag
-file           string 保存路径
-packageName    string 包名
-prefix         string 表前缀
-realNameMethod string 结构体对应的表名
-table          string 要迁移的表
-tagKey         string tag的key
```
