package router

import (
	"github.com/gin-gonic/gin"

	"custom-schedule/internal/controller"
	"custom-schedule/internal/middleware"
	"custom-schedule/internal/service"
)

// SetupRouter 初始化 Gin 路由
func SetupRouter(userService service.UserService) *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.CORSMiddleware())

	// 健康检查
	api := r.Group("/api")
	controller.RegisterHealthRoutes(api)

	// v1 API
	v1 := api.Group("/v1")

	// 用户模块
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(v1.Group("/users"))

	// TODO: 注册任务、计划等模块的路由

	return r
}
