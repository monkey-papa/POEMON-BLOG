// Package middleware 请求 ID 追踪中间件
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDKey 上下文中的请求 ID 键名
const RequestIDKey = "X-Request-ID"

// RequestIDMiddleware 为每个请求生成唯一 ID
// 便于日志追踪和问题排查
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 优先使用客户端传递的请求 ID
		requestID := c.GetHeader(RequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 存入上下文
		c.Set(RequestIDKey, requestID)

		// 设置响应头
		c.Header(RequestIDKey, requestID)

		c.Next()
	}
}

// GetRequestID 从上下文获取请求 ID
func GetRequestID(c *gin.Context) string {
	if id, exists := c.Get(RequestIDKey); exists {
		return id.(string)
	}
	return ""
}
