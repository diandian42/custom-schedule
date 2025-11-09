package response

import "github.com/gin-gonic/gin"

// JSON 通用返回格式
func JSON(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(code, gin.H{
		"data":    data,
		"message": message,
	})
}
