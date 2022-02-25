package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 	// 不用中间件的话，可以直接写在这里
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, UPDATE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) // 204
			return
		}

		// 使用中间件
		// cors.New(cors.Options{
		// 	AllowedOrigins:   []string{"*"},
		// 	AllowCredentials: true,
		// 	AllowedHeaders:   []string{"*"},
		// 	AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		// })
		c.Next()
	}
}
