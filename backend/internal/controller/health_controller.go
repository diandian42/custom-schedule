package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHealthRoutes 注册健康检查相关路由
func RegisterHealthRoutes(router *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "custom-schedule backend is running",
		})
	})
}
