// Package handlers 微言相关接口处理
package handlers

import (
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// WeiYanListRequest 微言列表请求
type WeiYanListRequest struct {
	Current int  `json:"current"` // 当前页码
	Size    int  `json:"size"`    // 每页条数
	All     bool `json:"all"`     // 是否显示所有（包括非公开）
}

// WeiYanNewsListRequest 文章进展列表请求
type WeiYanNewsListRequest struct {
	Current      int  `json:"current"`      // 当前页码
	Size         int  `json:"size"`         // 每页条数
	Source       *int `json:"source"`       // 文章ID
	All          bool `json:"all"`          // 是否显示所有
	NoPagination bool `json:"noPagination"` // true 表示获取全部数据
}

// SaveWeiYanRequest 保存微言请求
// 有 source 则为文章进展，无 source 则为普通微言
type SaveWeiYanRequest struct {
	Content  string `json:"content" binding:"required"` // 内容
	Type     string `json:"type"`                       // 类型
	Source   *int   `json:"source"`                     // 关联文章ID（有值则为文章进展）
	IsPublic *bool  `json:"isPublic"`                   // 是否公开（文章进展默认公开）
}

// ChangeWeiYanPublicRequest 修改公开状态请求
type ChangeWeiYanPublicRequest struct {
	ID       uint `json:"id" binding:"required"` // 微言ID
	IsPublic bool `json:"isPublic"`              // 是否公开
}

// ==================== 接口处理函数 ====================

// WeiYanList 获取微言列表
// @Summary 获取微言列表
// @Description 获取用户的微言列表。未登录看Boss公开微言，登录看自己的公开微言
func WeiYanList(c *gin.Context) {
	var req WeiYanListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 验证 All 参数：只有后台用户才能查看所有人的微言（后台管理用）
	// Boss(0)、Admin(1)、Visitor(3) 可以，Normal(2) 不能
	showAll := req.All
	if showAll {
		if client, exists := c.Get("client"); exists {
			if clientInfo, ok := client.(*models.Client); ok {
				showAll = clientInfo.UserType != models.UserTypeNormal
			} else {
				showAll = false
			}
		} else {
			showAll = false
		}
	}

	// 获取当前登录用户，确定要查询的用户ID
	var userID uint = 0
	if client, exists := c.Get("client"); exists {
		if clientInfo, ok := client.(*models.Client); ok {
			// 登录用户：看自己的微言
			userID = clientInfo.UserID
		}
	}
	if userID == 0 {
		// 未登录：看 Boss 的微言
		var user models.AuthUser
		if err := config.DB.Where("is_staff = ?", true).First(&user).Error; err != nil {
			response.Success(c, []interface{}{})
			return
		}
		userID = user.ID
	}

	// 查询没有关联文章的微言
	query := config.DB.Model(&models.WeiYan{}).Where("source IS NULL")

	if showAll {
		// 管理员模式：查看所有人的所有微言（后台管理用）
		// 不加任何过滤
	} else {
		// 前台模式：只返回指定用户的公开微言
		// 无论登录与否，都只显示公开的微言
		query = query.Where("user_id = ? AND is_public = ?", userID, true)
	}

	var total int64
	query.Count(&total)

	var weiyans []models.WeiYan
	current, size, offset := normalizePagination(req.Current, req.Size)
	query.Order("id DESC").Offset(offset).Limit(size).Find(&weiyans)

	data := make([]map[string]interface{}, 0, len(weiyans))
	for _, w := range weiyans {
		data = append(data, buildWeiYanData(w))
	}

	response.SuccessPage(c, data, total, current, size)
}

// WeiYanNewsList 获取文章进展列表
// @Summary 获取文章进展列表
// @Description 获取与文章关联的微言（进展记录），任何人都可以查看公开的进展
func WeiYanNewsList(c *gin.Context) {
	var req WeiYanNewsListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 验证 All 参数：只有后台用户才能查看所有进展（后台管理用）
	// Boss(0)、Admin(1)、Visitor(3) 可以，Normal(2) 不能
	showAll := req.All
	if showAll {
		if client, exists := c.Get("client"); exists {
			if clientInfo, ok := client.(*models.Client); ok {
				showAll = clientInfo.UserType != models.UserTypeNormal
			} else {
				showAll = false
			}
		} else {
			showAll = false
		}
	}

	query := config.DB.Model(&models.WeiYan{})

	if showAll {
		// 管理员模式：返回所有有关联文章的微言（后台管理用）
		query = query.Where("source IS NOT NULL")
	} else {
		// 前台模式：只返回公开的，且必须指定文章ID
		query = query.Where("source IS NOT NULL AND is_public = ?", true)
		if req.Source != nil && *req.Source > 0 {
			query = query.Where("source = ?", *req.Source)
		}
	}

	var total int64
	query.Count(&total)

	var weiyans []models.WeiYan
	current, size, offset := normalizePagination(req.Current, req.Size)

	if req.NoPagination {
		// 获取全部数据
		query.Order("id DESC").Find(&weiyans)
	} else {
		// 分页查询
		query.Order("id DESC").Offset(offset).Limit(size).Find(&weiyans)
	}

	data := make([]map[string]interface{}, 0, len(weiyans))
	for _, w := range weiyans {
		data = append(data, buildWeiYanData(w))
	}

	response.SuccessPage(c, data, total, current, size)
}

// SaveWeiYan 保存微言
// @Summary 保存微言
// @Description 发布微言，有source参数则为文章进展，否则为普通微言
func SaveWeiYan(c *gin.Context) {
	var req SaveWeiYanRequest
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

	// 设置默认类型
	if req.Type == "" {
		req.Type = "a"
	}

	// 如果有关联文章ID，验证文章是否存在（文章进展模式）
	if req.Source != nil && *req.Source > 0 {
		var article models.Article
		if err := config.DB.Where("id = ?", *req.Source).First(&article).Error; err != nil {
			response.Failure(c, "文章不存在")
			return
		}
	}

	// 确定公开状态：文章进展默认公开，普通微言使用传入值或默认false
	isPublic := false
	if req.Source != nil && *req.Source > 0 {
		isPublic = true // 文章进展默认公开
	}
	if req.IsPublic != nil {
		isPublic = *req.IsPublic // 如果显式传入则使用传入值
	}

	weiyan := models.WeiYan{
		UserID:   currentUser.UserID, // 直接使用当前登录用户ID
		Content:  req.Content,
		Type:     req.Type,
		Source:   req.Source,
		IsPublic: isPublic,
	}

	if err := config.DB.Create(&weiyan).Error; err != nil {
		response.Failure(c, "保存失败")
		return
	}

	response.SuccessMessage(c, "保存成功")
}

// DeleteWeiYan 删除微言
// @Summary 删除微言
// @Description 删除指定微言，只有作者或管理员可以删除
func DeleteWeiYan(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 获取微言信息
	var weiyan models.WeiYan
	if err := config.DB.First(&weiyan, req.ID).Error; err != nil {
		response.Failure(c, "微言不存在")
		return
	}

	// 获取当前用户
	client, exists := c.Get("client")
	if !exists {
		response.Failure(c, "未登录")
		return
	}
	currentUser := client.(*models.Client)

	// 验证权限：只有作者或管理员可以删除
	isAuthor := weiyan.UserID == currentUser.UserID
	isAdmin := currentUser.UserType <= models.UserTypeAdmin
	if !isAuthor && !isAdmin {
		response.Failure(c, "无权删除他人的微言")
		return
	}

	if err := config.DB.Delete(&models.WeiYan{}, req.ID).Error; err != nil {
		response.Failure(c, "删除失败")
		return
	}

	response.SuccessMessage(c, "删除成功")
}

// ChangeWeiYanPublic 修改微言公开状态
// @Summary 修改微言公开状态
// @Description 管理员修改微言的公开/私密状态
func ChangeWeiYanPublic(c *gin.Context) {
	var req ChangeWeiYanPublicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 检查记录是否存在
	var count int64
	config.DB.Model(&models.WeiYan{}).Where("id = ?", req.ID).Count(&count)
	if count == 0 {
		response.Failure(c, "记录不存在")
		return
	}

	if err := config.DB.Model(&models.WeiYan{}).Where("id = ?", req.ID).Update("is_public", req.IsPublic).Error; err != nil {
		response.Failure(c, "更新失败")
		return
	}

	response.SuccessMessage(c, "更新成功")
}

// ==================== 辅助函数 ====================

// buildWeiYanData 构建微言数据
func buildWeiYanData(w models.WeiYan) map[string]interface{} {
	client, _ := getUserClient(w.UserID)
	username := "未知用户"
	var avatar interface{} = nil

	if client != nil {
		username = client.Username
		if client.Avatar != nil {
			avatar = *client.Avatar
		}
	}

	return map[string]interface{}{
		"id":         w.ID,
		"userId":     w.UserID,
		"username":   username,
		"avatar":     avatar,
		"likeCount":  w.LikeCount,
		"content":    w.Content,
		"type":       w.Type,
		"source":     w.Source,
		"isPublic":   w.IsPublic,
		"createTime": w.CreateTime,
	}
}
