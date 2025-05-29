package main

import (
	"ai-ZhenLou/global"
	"ai-ZhenLou/initialize"
	"fmt"
)

func main() {
	fmt.Println("开始搞起！")
	// 初始化配置
	initialize.InitializeConfig()
	// 初始化日志
	initialize.InitializeLog(*global.App.Config)
	// 初始化redis
	initialize.InitializeRedis()
	// 初始化mysql
	initialize.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		initialize.CloseDB()
	}()

	// 启动服务
	initialize.RunServer()
}
