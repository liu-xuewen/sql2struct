package sql2struct

import "fmt"

func init() {
	// git update-index --assume-unchanged localConfig.go  //git忽略已经被提交的文件

	fmt.Println("config Dsn")
	Dsn = "/stt_match?charset=utf8&parseTime=True&loc=Local"
}
