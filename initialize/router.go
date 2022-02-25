package initialize

import (
	"go-basic-web/middleware"
	"go-basic-web/router"

	"github.com/gin-gonic/gin"
)

// 必须指名有返回值 这样用到的地方就可以接收返回值了
func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(
		// cors.New(cors.Options{
		// 	AllowedOrigins:   []string{"*"},
		// 	AllowCredentials: true,
		// 	AllowedHeaders:   []string{"*"},
		// 	AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "UPDATE"},

		// }),
		middleware.CorsMiddle(),   // cors 中间件
		middleware.LoggerMiddle(), // 接口日志中间件
	)

	// 配置全局的路径 api/v1/

	ApiGroup := Router.Group("/api/v1")

	// 将路由在这里进行注册
	router.InitUserRouter(ApiGroup)
	router.InitLogin(ApiGroup)

	return Router

}
