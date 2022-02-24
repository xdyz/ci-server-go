package middleware

import "github.com/gin-gonic/gin"

func LoggerMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
