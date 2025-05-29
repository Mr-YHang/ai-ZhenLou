package main

import (
	"ai-ZhenLou/config"
	"ai-ZhenLou/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("开始搞起！")
	// 初始化配置
	initialize.InitializeConfig()

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	r.Run(":" + config.Conf.App.Port)

}
