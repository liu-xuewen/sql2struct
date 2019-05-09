package main

import (
	"fmt"

	"github.com/liu-xuewen/sql2struct/sql2struct"
)

func main() {
	t2t := sql2struct.NewTable2Struct()

	err := t2t.
		SavePath("/home/go/project/model/model.go").
		Dsn("root:root@tcp(localhost:3306)/test?charset=utf8").
		Run()
	fmt.Println(err)
}
