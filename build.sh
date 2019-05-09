#!/usr/bin/env bash

# 版本
version="v0.0.1"

# 名字
linux=mysql2struct
# win=mysql2struct-win."$version".exe

# 打包
go build -o "$linux" main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$linux" main.go
# CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "$win" main.go

# 压缩
#upx $linux
#upx $win
