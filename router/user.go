package router

import (
	"go-basic-web/middleware"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Tags     用户
// @Summary  请求用户列表
// @Param    page  query  int     true   "页码"    default(1)
// @Param    size  query  int     true   "每页数量"  default(10)
// @Param    name  query  string  false  "用户名"
// @Router   /user [get]
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
