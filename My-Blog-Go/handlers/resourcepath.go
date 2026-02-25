// Package handlers 资源路径相关接口处理
package handlers

import (
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// ResourcePathListRequest 资源路径列表请求
type ResourcePathListRequest struct {
	Current      int    `json:"current"`      // 当前页码
	Size         int    `json:"size"`         // 每页条数
	ResourceType string `json:"resourceType"` // 资源类型筛选
	Classify     string `json:"classify"`     // 分类筛选
	Status       *bool  `json:"status"`       // 状态筛选
	IsAdmin      bool   `json:"isAdmin"`      // 是否管理员模式（查看全部，包括未启用）
	NoPagination bool   `json:"noPagination"` // true 表示获取全部数据
}

// SaveResourcePathRequest 保存资源路径请求
type SaveResourcePathRequest struct {
	Title        string  `json:"title" binding:"required"` // 标题
	Classify     *string `json:"classify"`                 // 分类
	Cover        *string `json:"cover"`                    // 封面
	URL          *string `json:"url"`                      // 链接
	Introduction *string `json:"introduction"`             // 简介
	Type         string  `json:"type"`                     // 类型
	FriendAvatar *string `json:"friendAvatar"`             // 友链头像
}

// UpdateResourcePathRequest 更新资源路径请求
type UpdateResourcePathRequest struct {
	ID           uint   `json:"id" binding:"required"` // 记录ID
	Title        string `json:"title"`                 // 标题
	Classify     string `json:"classify"`              // 分类
	Cover        string `json:"cover"`                 // 封面
	URL          string `json:"url"`                   // 链接
	Introduction string `json:"introduction"`          // 简介
	Type         string `json:"type"`                  // 类型
	Remark       string `json:"remark"`                // 备注
	Status       *bool  `json:"status"`                // 状态
}

// LoveListRequest 分类统计请求
type LoveListRequest struct {
	Type string `json:"type"` // 类型筛选
}

// ==================== 接口处理函数 ====================

// ListResourcePath 资源路径列表接口
// @Summary 获取资源路径列表
// @Description 获取资源路径列表，isAdmin=true查看全部，否则只返回已启用的
func ListResourcePath(c *gin.Context) {
	var req ResourcePathListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 验证 isAdmin 参数
	isAdmin := isAdminRequest(c, req.IsAdmin)

	query := config.DB.Model(&models.ResourcePath{})

	// 非管理员模式只查询已启用的资源
	if !isAdmin {
		query = query.Where("status = ?", true)
	}

	// 应用筛选条件
	if req.ResourceType != "" {
		query = query.Where("type = ?", req.ResourceType)
	}
	if req.Classify != "" {
		query = query.Where("classify = ?", req.Classify)
	}
	if req.Status != nil && isAdmin {
		// 管理员模式支持状态筛选（已验证 Token）
		query = query.Where("status = ?", *req.Status)
	}

	var total int64
	query.Count(&total)

	var resourcePaths []models.ResourcePath
	current, size, offset := normalizePagination(req.Current, req.Size)

	if req.NoPagination {
		query.Order("create_time DESC").Find(&resourcePaths)
	} else {
		query.Order("create_time DESC").Offset(offset).Limit(size).Find(&resourcePaths)
	}

	data := make([]map[string]interface{}, 0, len(resourcePaths))
	for _, rp := range resourcePaths {
		data = append(data, buildResourcePathData(rp))
	}

	response.SuccessPage(c, data, total, current, size)
}

// SaveResourcePath 保存资源路径
// @Summary 保存资源路径
// @Description 添加新的资源路径记录，普通用户创建的需要管理员审核
func SaveResourcePath(c *gin.Context) {
	var req SaveResourcePathRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 获取当前登录用户
	client, exists := c.Get("client")
	if !exists {
		response.Failure(c, "未登录")
		return
	}
	currentUser := client.(*models.Client)

	// 管理员创建的资源直接启用，普通用户创建的需要审核
	// Boss(0)、Admin(1) 可以直接启用，其他用户需要审核
	status := currentUser.UserType <= models.UserTypeAdmin

	resourcePath := models.ResourcePath{
		Title:        req.Title,
		Classify:     req.Classify,
		Cover:        req.Cover,
		URL:          req.URL,
		Introduction: req.Introduction,
		Type:         req.Type,
		Remark:       req.FriendAvatar,
		Status:       status,
	}

	if err := config.DB.Create(&resourcePath).Error; err != nil {
		response.Failure(c, "保存失败")
		return
	}

	if status {
	response.SuccessMessage(c, "保存成功")
	} else {
		response.SuccessMessage(c, "保存成功，等待管理员审核")
	}
}

// UpdateResourcePath 更新资源路径
// @Summary 更新资源路径
// @Description 管理员更新资源路径信息
func UpdateResourcePath(c *gin.Context) {
	var req UpdateResourcePathRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 检查记录是否存在
	var count int64
	config.DB.Model(&models.ResourcePath{}).Where("id = ?", req.ID).Count(&count)
	if count == 0 {
		response.Failure(c, "记录不存在")
		return
	}

	// 构建更新字段
	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Classify != "" {
		updates["classify"] = req.Classify
	}
	if req.Cover != "" {
		updates["cover"] = req.Cover
	}
	if req.URL != "" {
		updates["url"] = req.URL
	}
	if req.Introduction != "" {
		updates["introduction"] = req.Introduction
	}
	if req.Type != "" {
		updates["type"] = req.Type
	}
	if req.Remark != "" {
		updates["remark"] = req.Remark
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if len(updates) > 0 {
		if err := config.DB.Model(&models.ResourcePath{}).Where("id = ?", req.ID).Updates(updates).Error; err != nil {
			response.Failure(c, "更新失败")
			return
		}
	}

	response.SuccessMessage(c, "更新成功")
}

// DeleteResourcePath 删除资源路径
// @Summary 删除资源路径
// @Description 管理员删除资源路径记录
func DeleteResourcePath(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if err := config.DB.Delete(&models.ResourcePath{}, req.ID).Error; err != nil {
		response.Failure(c, "删除失败")
		return
	}

	response.SuccessMessage(c, "删除成功")
}

// ListCollect 获取收藏列表
// @Summary 获取收藏列表
// @Description 按分类获取所有收藏的资源
func ListCollect(c *gin.Context) {
	// 获取所有分类
	var classifies []string
	config.DB.Model(&models.ResourcePath{}).
		Where("status = ? AND type = ?", true, "favorites").
		Distinct("classify").
		Pluck("classify", &classifies)

	// 按分类分组返回数据
	data := make(map[string][]map[string]interface{})

	for _, classify := range classifies {
		var resourcePaths []models.ResourcePath
		config.DB.Where("classify = ? AND status = ? AND type = ?", classify, true, "favorites").
			Find(&resourcePaths)

		var items []map[string]interface{}
		for _, rp := range resourcePaths {
			items = append(items, map[string]interface{}{
				"id":           rp.ID,
				"classify":     rp.Classify,
				"createTime":   rp.CreateTime,
				"introduction": rp.Introduction,
				"cover":        rp.Cover,
				"remark":       rp.Remark,
				"status":       rp.Status,
				"title":        rp.Title,
				"type":         rp.Type,
				"url":          rp.URL,
			})
		}
		data[classify] = items
	}

	response.Success(c, data)
}

// LoveList 获取分类统计
// @Summary 获取分类统计
// @Description 按分类统计资源数量
func LoveList(c *gin.Context) {
	var req LoveListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 兼容无参数请求
		req.Type = ""
	}

	type Result struct {
		Classify string
		Count    int64
	}

	query := config.DB.Model(&models.ResourcePath{}).
		Select("classify, COUNT(*) as count").
		Where("status = ?", true)

	if req.Type != "" {
		query = query.Where("type = ?", req.Type)
	}

	var results []Result
	query.Group("classify").Order("MAX(create_time) DESC").Scan(&results)

	data := make([]map[string]interface{}, 0, len(results))
	for _, r := range results {
		data = append(data, map[string]interface{}{
			"classify": r.Classify,
			"count":    r.Count,
		})
	}

	response.Success(c, data)
}

// ==================== 辅助函数 ====================

// buildResourcePathData 构建资源路径数据
func buildResourcePathData(rp models.ResourcePath) map[string]interface{} {
	return map[string]interface{}{
		"id":           rp.ID,
		"classify":     rp.Classify,
		"createTime":   rp.CreateTime,
		"introduction": rp.Introduction,
		"cover":        rp.Cover,
		"friendAvatar": rp.Remark,
		"status":       rp.Status,
		"title":        rp.Title,
		"type":         rp.Type,
		"url":          rp.URL,
	}
}
