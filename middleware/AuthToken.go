package middleware

import (
	"go-basic-web/utils"

	"github.com/gin-gonic/gin"
)

// 中间件内部需要返回一个gin.HandlerFunc函数
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token  Authorization 为key
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			utils.Faild(c, "请求头中没有token")
			// 取消执行后面的handler 放弃此次请求
			c.Abort()
			return
		}

		// 如果是有的 那么就 解析token
		claims, err := utils.ParseToken(tokenString)

		if err != nil {
			utils.Faild(c, "token解析失败")
			// 取消执行后面的handler 放弃此次请求
			c.Abort()
			return
		}

		// 将解析出来的信息放到上下文中
		c.Set("claims", claims)

		c.Next()
	}
}
