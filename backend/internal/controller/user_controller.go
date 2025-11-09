package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"custom-schedule/internal/service"
)

// UserController 用户相关接口控制器
type UserController struct {
	UserService service.UserService
}

// NewUserController 创建 UserController
func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// RegisterRoutes 注册用户相关路由
func (uc *UserController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", uc.listUsers)
	router.POST("", uc.createUser)
}

func (uc *UserController) listUsers(c *gin.Context) {
	// TODO: 调用 service 获取用户列表
	c.JSON(http.StatusOK, gin.H{
		"message": "list users - to be implemented",
	})
}

func (uc *UserController) createUser(c *gin.Context) {
	// TODO: 调用 service 创建用户
	c.JSON(http.StatusCreated, gin.H{
		"message": "create user - to be implemented",
	})
}
