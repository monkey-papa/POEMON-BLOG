// Package utils 定时清理任务
package utils

import (
	"log"
	"my-blog-go/config"
	"my-blog-go/models"
	"time"
)

// StartCleanupTasks 启动后台清理任务
func StartCleanupTasks() {
	// 验证码清理（每小时执行）
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		// 启动时先执行一次
		cleanupExpiredCodes()

		for range ticker.C {
			cleanupExpiredCodes()
		}
	}()

	log.Println("✅ 后台清理任务已启动")
}

// cleanupExpiredCodes 清理过期验证码
// 验证码有效期为 5 分钟
func cleanupExpiredCodes() {
	// 删除创建时间超过 10 分钟的验证码
	expireTime := time.Now().Add(-10 * time.Minute)
	result := config.DB.Where("create_time < ?", expireTime).Delete(&models.Code{})

	if result.Error != nil {
		log.Printf("🧹 清理过期验证码失败: %v", result.Error)
		return
	}
	if result.RowsAffected > 0 {
		log.Printf("🧹 清理了 %d 条过期验证码", result.RowsAffected)
	}
}
