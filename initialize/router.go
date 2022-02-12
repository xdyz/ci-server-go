package initialize

import (
	"ci-server-go/router"

	"github.com/gin-gonic/gin"
)

// 必须指名有返回值 这样用到的地方就可以接收返回值了
func Routers() *gin.Engine {
	Router := gin.Default()

	// 配置全局的路径 api/v1/

	ApiGroup := Router.Group("/api/v1")

	router.InitUserRouter(ApiGroup)

	return Router

}
