// Package handlers 禁用词相关接口处理
package handlers

import (
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// WordsListRequest 禁用词列表请求
type WordsListRequest struct {
	Current   int    `json:"current"`   // 当前页码
	Size      int    `json:"size"`      // 每页条数
	SearchKey string `json:"searchKey"` // 搜索关键词
}

// SaveWordsRequest 保存禁用词请求
type SaveWordsRequest struct {
	Avatar   *string `json:"avatar"`                     // 头像
	Username string  `json:"username"`                   // 用户名
	Message  string  `json:"message" binding:"required"` // 禁用词内容
}

// UpdateWordsRequest 更新禁用词请求
type UpdateWordsRequest struct {
	ID      uint   `json:"id" binding:"required"`      // 记录ID
	Message string `json:"message" binding:"required"` // 禁用词内容
}

// ==================== 接口处理函数 ====================

// ListWords 获取禁用词列表
// @Summary 获取禁用词列表
// @Description 获取所有禁用词记录
func ListWords(c *gin.Context) {
	var req WordsListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	query := config.DB.Model(&models.Words{})

	if req.SearchKey != "" {
		query = query.Where("message LIKE ?", "%"+req.SearchKey+"%")
	}

	var total int64
	query.Count(&total)

	var words []models.Words
	current, size, offset := normalizePagination(req.Current, req.Size)
	query.Order("id").Offset(offset).Limit(size).Find(&words)

	data := make([]map[string]interface{}, 0, len(words))
	for _, w := range words {
		data = append(data, map[string]interface{}{
			"id":       w.ID,
			"username": w.Username,
			"message":  w.Message,
			"avatar":   w.Avatar,
		})
	}

	response.SuccessPage(c, data, total, current, size)
}

// SaveWords 保存禁用词
// @Summary 添加禁用词
// @Description 管理员添加禁用词
func SaveWords(c *gin.Context) {
	var req SaveWordsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 检查禁用词是否已存在
	var existing models.Words
	if err := config.DB.Where("message = ?", req.Message).First(&existing).Error; err == nil {
		response.Failure(c, "该禁用词已存在")
		return
	}

	words := models.Words{
		Avatar:   req.Avatar,
		Username: req.Username,
		Message:  req.Message,
	}

	if err := config.DB.Create(&words).Error; err != nil {
		response.Failure(c, "保存失败")
		return
	}

	response.SuccessMessage(c, "保存成功")
}

// UpdateWords 更新禁用词
// @Summary 更新禁用词
// @Description 管理员更新禁用词
func UpdateWords(c *gin.Context) {
	var req UpdateWordsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 检查记录是否存在
	var word models.Words
	if err := config.DB.First(&word, req.ID).Error; err != nil {
		response.Failure(c, "记录不存在")
		return
	}

	// 检查新禁用词是否与其他记录重复
	var existing models.Words
	if err := config.DB.Where("message = ? AND id != ?", req.Message, req.ID).First(&existing).Error; err == nil {
		response.Failure(c, "该禁用词已存在")
		return
	}

	if err := config.DB.Model(&models.Words{}).Where("id = ?", req.ID).Update("message", req.Message).Error; err != nil {
		response.Failure(c, "更新失败")
		return
	}

	response.SuccessMessage(c, "更新成功")
}

// DeleteWords 删除禁用词
// @Summary 删除禁用词
// @Description 管理员删除禁用词
func DeleteWords(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if err := config.DB.Delete(&models.Words{}, req.ID).Error; err != nil {
		response.Failure(c, "删除失败")
		return
	}

	response.SuccessMessage(c, "删除成功")
}
