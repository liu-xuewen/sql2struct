package sql2struct

import "fmt"

func init() {
	// git update-index --assume-unchanged ./sql2struct/localConfig.go  //git忽略已经被提交的文件


	//Dsn = "root:root@tcp(localhost:3306)/tt?charset=utf8"

	 Dsn = "root:root@tcp(localhost:3306)/tt?charset=utf8"
	fmt.Println("config Dsn",Dsn)
}
