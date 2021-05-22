package main

import (
	"log-api/common/log"
	"log-api/irisapp"
	_ "log-api/model"//初始化数据库模型
)

func main() {
	log.INFO("main", "开始启动app")
	//初始化mvc框架
	irisapp.InitApp()
	//启动运行
	irisapp.Start()
	log.INFO("main", "app退出")

}
