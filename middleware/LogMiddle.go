package middleware

import (
	"go-basic-web/global"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware 日志中间件  洋葱模型
func LoggerMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()             // 开始时间
		path := c.Request.URL.Path      // 请求路径
		query := c.Request.URL.RawQuery // 请求参数 query 类型参数
		params := c.Request.PostForm    // 请求参数 form 类型参数

		c.Next() // 执行后面的handler

		cost := time.Since(start) // 耗时
		global.Logger.Info(
			path,
			zap.Int("status", c.Writer.Status()),            // 状态码
			zap.String("method", c.Request.Method),          // 请求方式
			zap.String("path", path),                        // 请求路径
			zap.String("query", query),                      // 请求参数 query 类型参数
			zap.Any("params", params),                       // 请求参数 params 类型参数
			zap.String("ip", c.ClientIP()),                  // 请求客户端ip
			zap.String("user-agent", c.Request.UserAgent()), // 请求客户端浏览器信息
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()), // 错误信息
			zap.Duration("cost", cost), // 耗时
		)
	}
}
