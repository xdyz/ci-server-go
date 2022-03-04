package main

import (
	"fmt"
	_ "go-basic-web/common"
	_ "go-basic-web/docs"
	"go-basic-web/global"
	"go-basic-web/initialize"

	"github.com/fvbock/endless"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title        后端基础架构api文档
// @version      1.0
// @description  后端基础架构api文档
// @host         localhost:3001
// @BasePath     /api/v1
func main() {

	// 执行路由初始化与启动
	router := initialize.Routers()

	// 注册swagger 添加swagger文档
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", "3001"))
	// 设置访问路径
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	global.Logger.Sugar().Infof("后端基础架构api文档启动成功，监听端口：%s", "3001")
	// router.Run(":3001")

	endless.ListenAndServe(":3001", router)
}
