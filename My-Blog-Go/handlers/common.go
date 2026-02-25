// Package handlers 提供所有API处理函数
// 本文件定义公共请求结构体和辅助函数
package handlers

import (
	"log"
	"my-blog-go/config"
	"my-blog-go/models"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ==================== 共享 HTTP Client ====================

// SharedHTTPClient 全局共享的 HTTP 客户端（带连接池和超时）
var SharedHTTPClient = &http.Client{
	Timeout: 15 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	},
}

// ==================== 共享工具函数 ====================

// GetClientIP 获取客户端真实IP（统一实现，供 handlers 和 middleware 使用）
func GetClientIP(c *gin.Context) string {
	// 优先级: X-Forwarded-For > X-Real-IP > CF-Connecting-IP > ClientIP
	if xForwardedFor := c.GetHeader("X-Forwarded-For"); xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(ips[0])
	}
	if xRealIP := c.GetHeader("X-Real-IP"); xRealIP != "" {
		return strings.TrimSpace(xRealIP)
	}
	if cfIP := c.GetHeader("CF-Connecting-IP"); cfIP != "" {
		return strings.TrimSpace(cfIP)
	}
	return c.ClientIP()
}

// GetStaticFilePath 获取静态文件的绝对路径（基于可执行文件所在目录）
func GetStaticFilePath(relativePath string) string {
	return filepath.Join("static", relativePath)
}

// ==================== 公共请求结构体 ====================

// PageRequest 分页请求基础结构
type PageRequest struct {
	Current int `json:"current" binding:"min=1"` // 当前页码，从1开始
	Size    int `json:"size" binding:"min=1"`    // 每页条数
}

// PaginationParams 分页参数（通用）
type PaginationParams struct {
	Current      int  // 当前页码
	Size         int  // 每页条数
	NoPagination bool // 是否禁用分页
}

// NormalizePagination 规范化分页参数
// 返回规范化后的 current, size, noPagination
func NormalizePagination(current, size int, noPagination bool) (int, int, bool) {
	if noPagination {
		return 1, 0, true
	}
	if current <= 0 {
		current = 1
	}
	if size <= 0 {
		size = 10 // 默认每页10条
	}
	return current, size, false
}

// IDRequest 通用ID请求
type IDRequest struct {
	ID uint `json:"id" binding:"required"` // 记录ID
}

// IDsRequest 批量ID请求
type IDsRequest struct {
	IDs []uint `json:"ids" binding:"required,min=1"` // ID列表
}

// StatusRequest 状态修改请求
type StatusRequest struct {
	ID     uint `json:"id" binding:"required"` // 记录ID
	Status bool `json:"status"`                // 状态值
}

// ==================== 权限辅助函数 ====================

// isAdminRequest 验证请求是否具有管理员权限
// 如果 reqIsAdmin 为 false 直接返回 false
// 如果为 true，从 Token 中验证用户身份（Boss/Admin/Visitor 通过，Normal 不通过）
func isAdminRequest(c *gin.Context, reqIsAdmin bool) bool {
	if !reqIsAdmin {
		return false
	}
	client, exists := c.Get("client")
	if !exists {
		return false
	}
	clientInfo, ok := client.(*models.Client)
	if !ok {
		return false
	}
	return clientInfo.UserType != models.UserTypeNormal
}

// ==================== 批量查询辅助函数 ====================

// batchGetUserClients 批量获取用户Client信息
// 返回 map[userID]*models.Client
func batchGetUserClients(userIDs []uint) map[uint]*models.Client {
	result := make(map[uint]*models.Client, len(userIDs))
	if len(userIDs) == 0 {
		return result
	}
	// 去重
	uniqueIDs := make([]uint, 0, len(userIDs))
	seen := make(map[uint]bool, len(userIDs))
	for _, id := range userIDs {
		if !seen[id] {
			seen[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}
	var clients []models.Client
	if err := config.DB.Where("user_id IN ?", uniqueIDs).Find(&clients).Error; err != nil {
		log.Printf("[batchGetUserClients] 批量查询用户失败: %v", err)
		return result
	}
	for i := range clients {
		result[clients[i].UserID] = &clients[i]
	}
	return result
}

// batchGetUsernames 批量获取用户名
// 返回 map[userID]username
func batchGetUsernames(userIDs []uint) map[uint]string {
	result := make(map[uint]string, len(userIDs))
	if len(userIDs) == 0 {
		return result
	}
	uniqueIDs := make([]uint, 0, len(userIDs))
	seen := make(map[uint]bool, len(userIDs))
	for _, id := range userIDs {
		if !seen[id] {
			seen[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}
	var users []models.AuthUser
	if err := config.DB.Where("id IN ?", uniqueIDs).Find(&users).Error; err != nil {
		log.Printf("[batchGetUsernames] 批量查询用户名失败: %v", err)
		return result
	}
	for _, u := range users {
		result[u.ID] = u.Username
	}
	return result
}

// ==================== 日志辅助函数 ====================

// logAdminAction 记录管理员操作日志
func logAdminAction(adminID uint, adminName string, action string, targetID uint, detail string) {
	log.Printf("[ADMIN] 管理员 %s(ID:%d) %s 目标ID:%d %s", adminName, adminID, action, targetID, detail)
}

// ==================== 资源辅助函数 ====================

// getValidAvatar 检查头像资源是否有效
// 如果资源存在且状态为启用则返回原路径，否则返回空字符串
func getValidAvatar(avatar *string) string {
	if avatar == nil || *avatar == "" {
		return ""
	}
	var resource models.Resource
	if err := config.DB.Where("path = ?", *avatar).First(&resource).Error; err != nil {
		return *avatar // 资源表中不存在，返回原值（可能是外链）
	}
	if resource.Status {
		return *avatar
	}
	return ""
}

// getValidCover 检查封面资源是否有效
// 仅当资源存在且状态为启用时返回路径
func getValidCover(cover *string) string {
	if cover == nil || *cover == "" {
		return ""
	}
	var resource models.Resource
	if err := config.DB.Where("path = ?", *cover).First(&resource).Error; err != nil {
		return ""
	}
	if resource.Status {
		return *cover
	}
	return ""
}

// checkResourceStatus 检查资源状态
// fallback: 如果资源不存在于Resource表，是否返回原值
func checkResourceStatus(path string, fallback bool) string {
	if path == "" {
		return ""
	}
	var resource models.Resource
	if err := config.DB.Where("path = ?", path).First(&resource).Error; err != nil {
		if fallback {
			return path
		}
		return ""
	}
	if resource.Status {
		return path
	}
	return ""
}

// ptrToString 指针转字符串，空指针返回空字符串
func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// getUsername 根据用户ID获取用户名
func getUsername(userID uint) string {
	var user models.AuthUser
	if err := config.DB.First(&user, userID).Error; err != nil {
		return ""
	}
	return user.Username
}

// getCommentCount 获取文章评论数
func getCommentCount(articleID uint) int64 {
	var count int64
	config.DB.Model(&models.Comment{}).Where("source = ?", articleID).Count(&count)
	return count
}

// getLabel 获取标签信息
func getLabel(labelID uint, articleID ...uint) []map[string]interface{} {
	var label models.Label
	if err := config.DB.First(&label, labelID).Error; err != nil {
		return nil
	}
	var count int64
	config.DB.Model(&models.Article{}).Where("label_id = ?", labelID).Count(&count)
	result := map[string]interface{}{
		"id":               label.ID,
		"labelName":        label.LabelName,
		"labelDescription": label.LabelDescription,
		"sortId":           label.SortID,
		"countOfLabel":     count,
	}
	if len(articleID) > 0 {
		result["articleId"] = articleID[0]
	}
	return []map[string]interface{}{result}
}

// getSort 获取分类信息
func getSort(sortID uint) []map[string]interface{} {
	var sort models.Sort
	if err := config.DB.First(&sort, sortID).Error; err != nil {
		return nil
	}
	var articleCount, labelCount int64
	config.DB.Model(&models.Article{}).Where("sort_id = ?", sortID).Count(&articleCount)
	config.DB.Model(&models.Label{}).Where("sort_id = ?", sortID).Count(&labelCount)
	return []map[string]interface{}{
		{
			"id":              sort.ID,
			"sortName":        sort.SortName,
			"sortDescription": sort.SortDescription,
			"sortType":        sort.SortType,
			"priority":        sort.Priority,
			"countOfSort":     articleCount,
			"labels":          labelCount,
		},
	}
}

// getUserClient 根据用户ID获取Client信息
func getUserClient(userID uint) (*models.Client, error) {
	var client models.Client
	err := config.DB.Where("user_id = ?", userID).First(&client).Error
	return &client, err
}

// calculateOffset 计算分页偏移量（旧版本，仅返回 offset）
// 已废弃，推荐使用 normalizePagination
func calculateOffset(current, size int) int {
	_, _, offset := normalizePagination(current, size)
	return offset
}

// normalizePagination 规范化分页参数并计算偏移量
// 返回规范化后的 current, size 和 offset
// 处理 size=0 或 current=0 的边界情况
func normalizePagination(current, size int) (int, int, int) {
	if current < 1 {
		current = 1
	}
	if size < 1 {
		size = 10 // 默认每页10条
	}
	offset := (current - 1) * size
	return current, size, offset
}
