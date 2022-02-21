package router

import (
	"github.com/gin-gonic/gin"
)

// router 可以当做是contruller

// 接收一个gin.RouterGroup 指针参数
func InitLogin(r *gin.RouterGroup) {
	//  这里可以启动中间件
	loginGroup := r.Group("/login")
	{
		loginGroup.POST("/")
	}

}
