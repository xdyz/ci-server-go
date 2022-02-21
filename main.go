package main

import (
	"ci-server-go/initialize"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// 执行路由初始化与启动
	router := initialize.Routers()

	// 注册swagger 添加swagger文档
	// url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", "3001"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":3001")
}
