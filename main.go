package main

import (
	"ai-ZhenLou/global"
	"ai-ZhenLou/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("开始搞起！")
	// 初始化配置
	initialize.InitializeConfig()
	// 初始化日志
	initialize.InitializeLog(*global.App.Config)
	// 初始化mysql
	initialize.InitializeDB()

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	r.Run(":" + global.App.Config.App.Port)

}
