// Package handlers 健康检查接口
package handlers

import (
	"my-blog-go/config"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

// HealthCheck 健康检查接口
// @Summary 健康检查
// @Description 返回服务健康状态，用于负载均衡器和监控系统
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"uptime":    time.Since(startTime).String(),
	})
}

// ReadinessCheck 就绪检查接口
// @Summary 就绪检查
// @Description 检查服务是否准备好接收流量（包括数据库连接）
func ReadinessCheck(c *gin.Context) {
	// 检查数据库连接
	sqlDB, err := config.DB.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "not ready",
			"error":  "database connection failed",
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "not ready",
			"error":  "database ping failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ready",
		"database": "connected",
	})
}

// SystemInfo 系统信息接口（仅开发环境）
// @Summary 系统信息
// @Description 返回系统运行状态信息
func SystemInfo(c *gin.Context) {
	// 仅开发环境可用
	if config.AppConfig.Environment != "development" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "only available in development mode",
		})
		return
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.JSON(http.StatusOK, gin.H{
		"runtime": gin.H{
			"go_version":    runtime.Version(),
			"go_os":         runtime.GOOS,
			"go_arch":       runtime.GOARCH,
			"cpu_num":       runtime.NumCPU(),
			"goroutine_num": runtime.NumGoroutine(),
		},
		"memory": gin.H{
			"alloc_mb":       m.Alloc / 1024 / 1024,
			"total_alloc_mb": m.TotalAlloc / 1024 / 1024,
			"sys_mb":         m.Sys / 1024 / 1024,
			"gc_num":         m.NumGC,
		},
		"uptime": time.Since(startTime).String(),
	})
}
