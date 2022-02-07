package main

import (
	"github.com/gin-gonic/gin"
)

func Helo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	// 返回 gin 引擎
	e := gin.Default()

	e.GET("/ping", Helo)

	// port 端口为3000  如果不写默认为8080
	e.Run(":3000")
}
