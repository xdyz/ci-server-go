package main

import "ci-server-go/initialize"

func main() {

	// 执行路由初始化与启动
	router := initialize.Routers()

	router.Run(":3000")
}
