// Package middleware CORS跨域中间件
package middleware

import (
	"my-blog-go/config"
	"strings"

	"github.com/gin-gonic/gin"
)

// isAllowedOrigin 检查来源是否在白名单中（白名单从配置读取）
func isAllowedOrigin(origin string) bool {
	// 开发环境允许所有来源
	if config.AppConfig.Environment == "development" {
		return true
	}

	// 生产环境检查白名单
	for _, allowed := range config.AppConfig.CORSAllowedOrigins {
		if strings.HasPrefix(origin, allowed) {
			return true
		}
	}
	return false
}

// CORSMiddleware 跨域中间件
// 生产环境限制允许的来源，开发环境允许所有
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		// 检查来源是否允许
		if origin != "" && isAllowedOrigin(origin) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
		} else if config.AppConfig.Environment == "development" {
			// 开发环境回退
			c.Header("Access-Control-Allow-Origin", "*")
		}

		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-New-Token")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "X-New-Token")
		c.Header("Access-Control-Max-Age", "86400") // 预检请求缓存24小时

		// 安全相关头部
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
