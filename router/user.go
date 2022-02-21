package router

import "github.com/gin-gonic/gin"

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func GetUsers(g *gin.Context) {

}

// 接收一个gin.RouterGroup 指针参数
func InitUserRouter(r *gin.RouterGroup) {
	// 这里可以启动中间件 例如jwt 等等
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", GetUsers)
	}
}
