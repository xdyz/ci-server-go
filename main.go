package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 返回 gin 引擎
	e := gin.Default()

	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// port 端口为3000
	e.Run(":3000")
}
