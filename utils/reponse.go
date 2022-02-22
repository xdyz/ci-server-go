package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reponse(c *gin.Context, code int, message, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// 接口成功时 统一进行返回
func Success(c *gin.Context, message string, data interface{}) {
	Reponse(c, 0, "请求成功", "")
}

// 接口错误的时候 统一的请求返回
func Faild(c *gin.Context, message string) {
	Reponse(c, 1, message, nil)
}
