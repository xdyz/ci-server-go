package router

import (
	"go-basic-web/service"

	"github.com/gin-gonic/gin"
)

// 接收一个gin.RouterGroup 指针参数
func InitUserRouter(r *gin.RouterGroup) {
	// 这里可以启动中间件 例如jwt 等等
	// middleware.AuthToken()解析token
	// userGroup := r.Group("/user").Use(middleware.AuthToken())
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", service.GetUsers)       // 请求用户列表
		userGroup.GET("/:id", service.GetUserById) // 通过id 查询用户
		userGroup.POST("/", service.CreateUser)    // 创建用户
	}
}
