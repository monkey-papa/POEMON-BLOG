// Package handlers 资源相关接口处理
package handlers

import (
	"my-blog-go/config"
	"my-blog-go/middleware"
	"my-blog-go/models"
	"my-blog-go/response"
	"my-blog-go/utils"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// ResourceListRequest 资源列表请求
type ResourceListRequest struct {
	Current      int    `json:"current"`      // 当前页码
	Size         int    `json:"size"`         // 每页条数
	ResourceType string `json:"resourceType"` // 资源类型筛选
	NoPagination bool   `json:"noPagination"` // true 表示获取全部数据
	IsAdmin      bool   `json:"isAdmin"`      // 是否管理员模式（返回全部，包括未启用）
}

// SaveResourceRequest 保存资源请求
type SaveResourceRequest struct {
	Type     string `json:"type" binding:"required"`     // 资源类型
	Path     string `json:"path" binding:"required"`     // 资源路径
	Size     int    `json:"size" binding:"required"`     // 文件大小
	MimeType string `json:"mimeType" binding:"required"` // MIME类型
}

// ChangeResourceStatusRequest 修改资源状态请求
type ChangeResourceStatusRequest struct {
	ID   uint `json:"id" binding:"required"` // 资源ID
	Flag bool `json:"flag"`                  // 状态值
}

// DeleteImageRequest 删除图片请求
type DeleteImageRequest struct {
	URL string `json:"url"` // 图片URL
	ID  *uint  `json:"id"`  // 资源ID
}

// ==================== 接口处理函数 ====================

// ListResource 获取资源列表
// @Summary 获取资源列表
// @Description 获取资源记录，isAdmin=true返回全部，否则只返回已启用的
func ListResource(c *gin.Context) {
	var req ResourceListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 验证 isAdmin 参数
	isAdmin := isAdminRequest(c, req.IsAdmin)

	query := config.DB.Model(&models.Resource{})

	// 非管理员模式只查询已启用的资源
	if !isAdmin {
		query = query.Where("status = ?", true)
	}

	if req.ResourceType != "" {
		query = query.Where("mime_type = ?", req.ResourceType)
	}

	var total int64
	query.Count(&total)

	var resources []models.Resource
	current, size, offset := normalizePagination(req.Current, req.Size)

	if req.NoPagination {
		query.Order("create_time DESC").Find(&resources)
	} else {
		query.Order("create_time DESC").Offset(offset).Limit(size).Find(&resources)
	}

	// 批量获取用户信息
	userIDs := make([]uint, 0, len(resources))
	for _, r := range resources {
		userIDs = append(userIDs, r.UserID)
	}
	clientMap := batchGetUserClients(userIDs)

	data := make([]map[string]interface{}, 0, len(resources))
	for _, r := range resources {
		username := "未知用户"
		if client, ok := clientMap[r.UserID]; ok {
			username = client.Username
		}

		data = append(data, map[string]interface{}{
			"id":         r.ID,
			"createTime": r.CreateTime,
			"mimeType":   r.MimeType,
			"path":       r.Path,
			"size":       r.Size,
			"status":     r.Status,
			"type":       r.Type,
			"userId":     r.UserID,
			"username":   username,
		})
	}

	response.SuccessPage(c, data, total, current, size)
}

// SaveResource 保存资源记录
// @Summary 保存资源记录
// @Description 保存上传的资源信息到数据库
func SaveResource(c *gin.Context) {
	var req SaveResourceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 获取当前登录用户
	currentUser, exists := middleware.GetCurrentClient(c)
	if !exists {
		response.Failure(c, "未登录")
		return
	}

	resource := models.Resource{
		UserID:   currentUser.UserID, // 直接使用当前登录用户ID
		Type:     req.Type,
		Path:     req.Path,
		Size:     &req.Size,
		MimeType: &req.MimeType,
		Status:   true,
	}

	if err := config.DB.Create(&resource).Error; err != nil {
		response.Failure(c, "保存失败")
		return
	}

	response.SuccessMessage(c, "保存成功")
}

// ChangeResourceStatus 修改资源状态
// @Summary 修改资源状态
// @Description 管理员启用或禁用资源
func ChangeResourceStatus(c *gin.Context) {
	var req ChangeResourceStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if err := config.DB.Model(&models.Resource{}).Where("id = ?", req.ID).Update("status", req.Flag).Error; err != nil {
		response.Failure(c, "更新失败")
		return
	}

	response.SuccessMessage(c, "更新成功")
}

// DeleteImage 删除图片
// @Summary 删除图片
// @Description 从七牛云和数据库中删除图片
func DeleteImage(c *gin.Context) {
	var req DeleteImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 获取当前登录用户
	currentUser, exists := middleware.GetCurrentClient(c)
	if !exists {
		response.Failure(c, "未登录")
		return
	}

	// 如果传了资源ID，校验资源归属后再删除
	if req.ID != nil && *req.ID > 0 {
		var resource models.Resource
		if err := config.DB.First(&resource, *req.ID).Error; err != nil {
			response.Failure(c, "资源不存在")
			return
		}
		// 只有资源拥有者或管理员可以删除
		if resource.UserID != currentUser.UserID && currentUser.UserType > models.UserTypeAdmin {
			response.Failure(c, "无权删除该资源")
			return
		}
		if err := config.DB.Delete(&models.Resource{}, *req.ID).Error; err != nil {
			response.Failure(c, "删除资源记录失败")
			return
		}
	}

	// 获取当前用户的七牛云配置
	client, err := getUserClient(currentUser.UserID)
	if err != nil {
		response.Failure(c, "用户不存在")
		return
	}

	// 检查七牛云配置
	if client.QiniuAccessKey == nil || client.QiniuSecretKey == nil ||
		client.QiniuBucketName == nil || client.QiniuDomain == nil {
		response.Failure(c, "用户未配置七牛云")
		return
	}

	// 从七牛云删除
	qiniuCfg := utils.QiniuConfig{
		AccessKey:  *client.QiniuAccessKey,
		SecretKey:  *client.QiniuSecretKey,
		BucketName: *client.QiniuBucketName,
		Domain:     *client.QiniuDomain,
	}

	if err := utils.DeleteFromQiniu(req.URL, qiniuCfg); err != nil {
		response.Failure(c, "七牛云删除失败")
		return
	}

	response.SuccessMessage(c, "删除成功")
}
