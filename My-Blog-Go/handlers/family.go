// Package handlers 表白墙/爱链相关接口处理
package handlers

import (
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// FamilyListRequest 表白墙列表请求
type FamilyListRequest struct {
	Current int   `json:"current"` // 当前页码
	Size    int   `json:"size"`    // 每页条数
	Status  *bool `json:"status"`  // 状态筛选
	// IsAdmin 已废弃，改为从 Token 验证管理员身份
	IsAdmin bool `json:"isAdmin"` // 是否管理员模式（分页+查看全部）
}

// AddFamilyRequest 添加表白墙请求
type AddFamilyRequest struct {
	BgCover        string  `json:"bgCover" binding:"required"` // 背景图
	ManCover       string  `json:"manCover"`                   // 男方头像
	WomanCover     string  `json:"womanCover"`                 // 女方头像
	ManName        string  `json:"manName"`                    // 男方名称
	WomanName      string  `json:"womanName"`                  // 女方名称
	Timing         string  `json:"timing"`                     // 计时开始时间
	CountdownTitle *string `json:"countdownTitle"`             // 倒计时标题
	CountdownTime  *string `json:"countdownTime"`              // 倒计时时间
	FamilyInfo     *string `json:"familyInfo"`                 // 表白内容
	Status         *bool   `json:"status"`                     // 状态
	LikeCount      *int    `json:"likeCount"`                  // 点赞数
}

// ChangeFamilyStatusRequest 修改表白墙状态请求
type ChangeFamilyStatusRequest struct {
	ID   uint `json:"id" binding:"required"` // 记录ID
	Flag bool `json:"flag"`                  // 状态值
}

// ==================== 接口处理函数 ====================

// ListFamily 表白墙列表接口
// @Description 获取表白墙列表，isAdmin=true返回全部（分页），否则只返回已通过审核的
// @Description 获取表白墙列表，管理员返回全部（分页），普通用户只返回已通过审核的
func ListFamily(c *gin.Context) {
	var req FamilyListRequest
	// 允许空参数请求（兼容旧前端）
	c.ShouldBindJSON(&req)

	// 验证 isAdmin 参数
	isAdmin := isAdminRequest(c, req.IsAdmin)

	query := config.DB.Model(&models.Family{})
	if isAdmin {
		// 管理员模式：支持状态筛选和分页（已验证 Token）
		if req.Status != nil {
			query = query.Where("status = ?", *req.Status)
		}

		var total int64
		query.Count(&total)

		var families []models.Family
		current, size, offset := normalizePagination(req.Current, req.Size)
		query.Order("id").Offset(offset).Limit(size).Find(&families)

		data := make([]map[string]interface{}, 0, len(families))
		for _, family := range families {
			data = append(data, buildFamilyData(family, false))
		}

		response.SuccessPage(c, data, total, current, size)
	} else {
		// 前台模式：只返回已通过审核的
		var families []models.Family
		query.Where("status = ?", true).Find(&families)

		data := make([]map[string]interface{}, 0, len(families))
		for _, family := range families {
			data = append(data, buildFamilyData(family, true))
		}

		response.Success(c, data)
	}
}

// GetAdminFamily 获取管理员的表白墙
// @Summary 获取站长表白墙
// @Description 获取站长（管理员）的表白墙信息
func GetAdminFamily(c *gin.Context) {
	// 查找管理员用户
	var adminUser models.AuthUser
	if err := config.DB.Where("is_staff = ?", true).First(&adminUser).Error; err != nil {
		response.Success(c, []interface{}{})
		return
	}

	// 查找表白墙
	var family models.Family
	if err := config.DB.Where("user_id_id = ?", adminUser.ID).First(&family).Error; err != nil {
		response.Success(c, []interface{}{})
		return
	}

	// 只返回已启用的
	if !family.Status {
		response.Success(c, []interface{}{})
		return
	}

	response.Success(c, buildFamilyData(family, true))
}

// AddFamily 添加表白墙
// @Summary 添加表白墙
// @Description 用户添加表白墙信息
func AddFamily(c *gin.Context) {
	var req AddFamilyRequest
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

	// 设置默认值
	status := false
	if req.Status != nil {
		status = *req.Status
	}
	likeCount := 0
	if req.LikeCount != nil {
		likeCount = *req.LikeCount
	}

	family := models.Family{
		UserID:         currentUser.UserID, // 直接使用当前登录用户ID
		BgCover:        req.BgCover,
		ManCover:       req.ManCover,
		WomanCover:     req.WomanCover,
		ManName:        req.ManName,
		WomanName:      req.WomanName,
		Timing:         req.Timing,
		CountdownTitle: req.CountdownTitle,
		CountdownTime:  req.CountdownTime,
		FamilyInfo:     req.FamilyInfo,
		Status:         status,
		LikeCount:      likeCount,
	}

	if err := config.DB.Create(&family).Error; err != nil {
		response.Failure(c, "添加失败")
		return
	}

	response.SuccessMessage(c, "添加成功")
}

// DeleteFamily 删除表白墙
// @Summary 删除表白墙
// @Description 管理员删除指定表白墙记录
func DeleteFamily(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if err := config.DB.Delete(&models.Family{}, req.ID).Error; err != nil {
		response.Failure(c, "删除失败")
		return
	}

	response.SuccessMessage(c, "删除成功")
}

// ChangeFamilyStatus 修改表白墙状态
// @Summary 修改表白墙状态
// @Description 管理员审核表白墙
func ChangeFamilyStatus(c *gin.Context) {
	var req ChangeFamilyStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if err := config.DB.Model(&models.Family{}).Where("id = ?", req.ID).Update("status", req.Flag).Error; err != nil {
		response.Failure(c, "更新失败")
		return
	}

	response.SuccessMessage(c, "更新成功")
}

// ==================== 辅助函数 ====================

// buildFamilyData 构建表白墙数据
func buildFamilyData(family models.Family, checkResource bool) map[string]interface{} {
	bgCover := family.BgCover
	manCover := family.ManCover
	womanCover := family.WomanCover

	if checkResource {
		bgCover = checkResourceStatus(family.BgCover, false)
		manCover = checkResourceStatus(family.ManCover, false)
		womanCover = checkResourceStatus(family.WomanCover, false)
	}

	return map[string]interface{}{
		"id":             family.ID,
		"userId":         family.UserID,
		"bgCover":        bgCover,
		"manCover":       manCover,
		"womanCover":     womanCover,
		"manName":        family.ManName,
		"womanName":      family.WomanName,
		"timing":         family.Timing,
		"countdownTitle": family.CountdownTitle,
		"countdownTime":  family.CountdownTime,
		"status":         family.Status,
		"familyInfo":     family.FamilyInfo,
		"likeCount":      family.LikeCount,
		"createTime":     family.CreateTime,
	}
}
