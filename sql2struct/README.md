
```go

	// git update-index --assume-unchanged ./sql2struct/localConfig.go  //git忽略已经被提交的文件
	// 恢复跟踪  git update-index --no-assume-unchanged   sp_edaijia/protected/controllers/ApiController.php  //恢复跟踪
	// 如果忽略的文件多了，可以使用以下命令查看忽略列表
	// git uls-files -v | grep '^h\ '
	//提取文件路径，方法如下
	// git ls-files -v | grep '^h\ ' | awk '{print $2}'
	//所有被忽略的文件，取消忽略的方法，如下
	//git ls-files -v | grep '^h' | awk '{print $2}' |xargs git update-index --no-assume-unchanged

```
