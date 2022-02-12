package router

import "github.com/gin-gonic/gin"

// 接收一个gin.RouterGroup 指针参数
func InitUserRouter(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/")
	}
}
