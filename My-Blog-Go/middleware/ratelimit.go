// Package middleware 请求频率限制中间件
package middleware

import (
	"my-blog-go/config"
	"my-blog-go/response"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter 请求频率限制器
type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
	limit    int           // 时间窗口内允许的最大请求数
	window   time.Duration // 时间窗口
}

// 全局限制器实例
var (
	globalLimiter        *RateLimiter
	loginLimiter         *RateLimiter
	sensitiveRateLimiter *RateLimiter
)

// InitRateLimiters 初始化限流器（从配置读取参数，须在 LoadConfig 之后调用）
func InitRateLimiters() {
	cfg := config.AppConfig

	globalLimiter = NewRateLimiter(cfg.RateLimitGlobal, cfg.RateLimitGlobalWindow)
	loginLimiter = NewRateLimiter(cfg.RateLimitLogin, cfg.RateLimitLoginWindow)
	sensitiveRateLimiter = NewRateLimiter(cfg.RateLimitSensitive, cfg.RateLimitSensitiveWindow)

	// 定期清理过期记录
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			globalLimiter.cleanup()
			loginLimiter.cleanup()
			sensitiveRateLimiter.cleanup()
		}
	}()
}

// NewRateLimiter 创建新的限制器
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

// Allow 检查是否允许请求
func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	// 过滤掉过期的请求
	var validRequests []time.Time
	for _, t := range rl.requests[key] {
		if t.After(windowStart) {
			validRequests = append(validRequests, t)
		}
	}

	// 检查是否超限
	if len(validRequests) >= rl.limit {
		rl.requests[key] = validRequests
		return false
	}

	// 添加新请求
	validRequests = append(validRequests, now)
	rl.requests[key] = validRequests
	return true
}

// cleanup 清理过期记录
func (rl *RateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	for key, times := range rl.requests {
		var valid []time.Time
		for _, t := range times {
			if t.After(windowStart) {
				valid = append(valid, t)
			}
		}
		if len(valid) == 0 {
			delete(rl.requests, key)
		} else {
			rl.requests[key] = valid
		}
	}
}

// getClientIPForRateLimit 获取客户端IP用于限流
func getClientIPForRateLimit(c *gin.Context) string {
	if xff := c.GetHeader("X-Forwarded-For"); xff != "" {
		for i, ch := range xff {
			if ch == ',' {
				return xff[:i]
			}
		}
		return xff
	}
	if xri := c.GetHeader("X-Real-IP"); xri != "" {
		return xri
	}
	return c.ClientIP()
}

// RateLimitMiddleware 全局请求频率限制中间件
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIPForRateLimit(c)
		if !globalLimiter.Allow(ip) {
			response.Error(c, 429, "请求过于频繁，请稍后再试")
			c.Abort()
			return
		}
		c.Next()
	}
}

// LoginRateLimitMiddleware 登录请求频率限制
func LoginRateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIPForRateLimit(c)
		if !loginLimiter.Allow(ip) {
			response.Error(c, 429, "登录尝试过于频繁，请稍后再试")
			c.Abort()
			return
		}
		c.Next()
	}
}

// SensitiveRateLimitMiddleware 敏感操作频率限制
func SensitiveRateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIPForRateLimit(c)
		if !sensitiveRateLimiter.Allow(ip) {
			response.Error(c, 429, "操作过于频繁，请稍后再试")
			c.Abort()
			return
		}
		c.Next()
	}
}
