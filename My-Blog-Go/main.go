// Package main 博客后端服务入口
package main

import (
	"context"
	"fmt"
	"log"
	"my-blog-go/config"
	"my-blog-go/middleware"
	"my-blog-go/routes"
	"my-blog-go/utils"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化数据库
	config.InitDB()

	// 初始化限流器（须在 LoadConfig 之后）
	middleware.InitRateLimiters()

	// 设置Gin模式（减少路由注册日志）
	gin.SetMode(gin.ReleaseMode)

	// 创建Gin引擎（仅保留必要的中间件）
	r := gin.New()
	r.Use(gin.Recovery()) // 恢复中间件，防止 panic 导致服务崩溃

	// 开发环境开启请求日志
	if config.AppConfig.Environment == "development" {
		r.Use(gin.Logger())
	}

	// 设置API路由
	routes.SetupRoutes(r)

	// 启动后台清理任务
	utils.StartCleanupTasks()

	// 确保端口可用（开发环境自动杀掉占用进程）
	port := config.AppConfig.ServerPort
	ensurePortAvailable(port)

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 在 goroutine 中启动服务器
	go func() {
		log.Printf("🚀 服务启动在端口 %s...", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("服务启动失败:", err)
		}
	}()

	// 优雅关闭
	gracefulShutdown(srv)
}

// ensurePortAvailable 确保端口可用
// 开发环境下自动杀掉占用进程，生产环境仅报错退出
func ensurePortAvailable(port string) {
	addr := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", addr)
	if err == nil {
		// 端口可用，直接返回
		listener.Close()
		return
	}

	// 生产环境不自动杀进程，直接报错退出
	if config.AppConfig.Environment == "production" {
		log.Fatalf("❌ 端口 %s 被占用，请手动检查并释放", port)
	}

	// 开发环境：尝试杀掉占用进程
	log.Printf("⚠️  端口 %s 被占用，正在尝试释放（仅开发环境）...", port)

	if killProcessOnPort(port) {
		// 等待端口释放
		time.Sleep(500 * time.Millisecond)

		// 再次检查端口是否可用
		listener, err = net.Listen("tcp", addr)
		if err == nil {
			listener.Close()
			log.Printf("✅ 端口 %s 已成功释放", port)
			return
		}
	}

	log.Fatalf("❌ 无法释放端口 %s，请手动杀掉占用进程", port)
}

// killProcessOnPort 杀掉占用指定端口的进程
func killProcessOnPort(port string) bool {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin", "linux":
		// macOS/Linux: 使用 lsof 查找并杀掉进程
		// lsof -ti :8000 | xargs kill -9
		script := fmt.Sprintf("lsof -ti :%s | xargs kill -9 2>/dev/null", port)
		cmd = exec.Command("sh", "-c", script)
	case "windows":
		// Windows: 使用 netstat 和 taskkill
		// 先找到 PID，再杀掉
		script := fmt.Sprintf(`for /f "tokens=5" %%a in ('netstat -aon ^| findstr :%s') do taskkill /F /PID %%a`, port)
		cmd = exec.Command("cmd", "/C", script)
	default:
		log.Printf("⚠️  不支持的操作系统: %s", runtime.GOOS)
		return false
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		// 如果没有进程被杀掉，lsof 可能返回空，这不算错误
		outputStr := strings.TrimSpace(string(output))
		if outputStr != "" {
			log.Printf("⚠️  杀进程输出: %s", outputStr)
		}
	}

	return true
}

// gracefulShutdown 优雅关闭服务
func gracefulShutdown(srv *http.Server) {
	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("⏳ 正在关闭服务...")

	// 设置关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 关闭 HTTP 服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("服务关闭出错: %v", err)
	}

	// 关闭数据库连接
	if sqlDB, err := config.DB.DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Printf("数据库关闭出错: %v", err)
		}
	}

	log.Println("✅ 服务已安全关闭")
}
