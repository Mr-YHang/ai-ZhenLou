package router

import (
	"ai-ZhenLou/app/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(router *gin.Engine, h handler.Handler) {
	// 健康检查
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	user := router.Group("/user")
	user.POST("/login", h.Session.Login)

	ai := router.Group("/ai")
	ai.POST("/ack", h.AI.Talk)
}
