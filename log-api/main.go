package main

import (
	"log-api/irisapp"
	"log-api/model"
)

func main() {
	//初始化数据库模型
	if err := model.Init(); nil != err {
		panic(err)
	}
	//初始化mvc框架
	irisapp.InitApp()
	//启动运行
	irisapp.Start()
}
