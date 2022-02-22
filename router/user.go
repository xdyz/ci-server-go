package router

import (
	"go-basic-web/middleware"

	"github.com/gin-gonic/gin"
)

// @Router /user [get]
// @Summary 请求用户列表
// @Schemes
func GetUsers(g *gin.Context) {

}

// 接收一个gin.RouterGroup 指针参数
func InitUserRouter(r *gin.RouterGroup) {
	// 这里可以启动中间件 例如jwt 等等
	// middleware.AuthToken()解析token
	userGroup := r.Group("/user").Use(middleware.AuthToken())
	{
		userGroup.GET("/", GetUsers)
	}
}
