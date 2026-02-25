// Package handlers 网站信息相关接口处理
package handlers

import (
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// GetWebInfoRequest 获取网站信息请求
type GetWebInfoRequest struct {
	IsAdmin bool `json:"isAdmin"` // 是否管理员模式（返回全部，包括未启用）
}

// UpdateWebInfoRequest 更新网站信息请求
type UpdateWebInfoRequest struct {
	ID              uint   `json:"id" binding:"required"` // 记录ID
	UpdateMode      string `json:"updateMode"`            // 更新模式: "full"-完整更新, "partial"-部分更新（默认根据字段判断）
	WebName         string `json:"webName"`               // 网站名称
	WebTitle        string `json:"webTitle"`              // 网站标题
	Notices         string `json:"notices"`               // 公告
	Footer          string `json:"footer"`                // 页脚
	BackgroundImage string `json:"backgroundImage"`       // 背景图
	Avatar          string `json:"avatar"`                // 头像
	RandomAvatar    string `json:"randomAvatar"`          // 随机头像
	RandomName      string `json:"randomName"`            // 随机名称
	RandomCover     string `json:"randomCover"`           // 随机封面
	WaifuJson       string `json:"waifuJson"`             // 看板娘配置
	Status          bool   `json:"status"`                // 状态
}

// ==================== 接口处理函数 ====================

// GetWebInfo 统一获取网站信息接口
// @Summary 获取网站信息
// @Description 获取网站配置信息，isAdmin=true返回全部，否则只返回已启用的
func GetWebInfo(c *gin.Context) {
	var req GetWebInfoRequest
	// 允许空参数请求（兼容旧前端）
	c.ShouldBindJSON(&req)

	// 验证 isAdmin 参数
	isAdmin := isAdminRequest(c, req.IsAdmin)

	query := config.DB.Model(&models.WebInfo{})

	// 非管理员模式只查询已启用的配置
	if !isAdmin {
		query = query.Where("status = ?", true)
	}

	var webinfos []models.WebInfo
	query.Order("id DESC").Find(&webinfos)

	// 无数据时返回 null
	if len(webinfos) == 0 {
		response.Success(c, nil)
		return
	}

	// 统一返回第一条记录（对象格式）
	webinfo := webinfos[0]

	var data map[string]interface{}
	if isAdmin {
		// 管理员视图：原样返回所有字段（已验证 Token）
		data = map[string]interface{}{
			"id":              webinfo.ID,
			"webName":         webinfo.WebName,
			"webTitle":        webinfo.WebTitle,
			"notices":         webinfo.Notices,
			"footer":          webinfo.Footer,
			"backgroundImage": webinfo.BackgroundImage,
			"avatar":          webinfo.Avatar,
			"randomAvatar":    webinfo.RandomAvatar,
			"randomName":      webinfo.RandomName,
			"randomCover":     webinfo.RandomCover,
			"waifuJson":       webinfo.WaifuJson,
			"status":          webinfo.Status,
		}
	} else {
		// 前台视图：检查资源状态
		avatar := checkResourceStatus(webinfo.Avatar, true)
		backgroundImage := checkResourceStatus(ptrToString(webinfo.BackgroundImage), false)
		randomAvatar := checkResourceStatus(ptrToString(webinfo.RandomAvatar), false)
		randomCover := checkResourceStatus(ptrToString(webinfo.RandomCover), true)

		data = map[string]interface{}{
			"id":              webinfo.ID,
			"webName":         webinfo.WebName,
			"webTitle":        webinfo.WebTitle,
			"notices":         webinfo.Notices,
			"footer":          webinfo.Footer,
			"backgroundImage": backgroundImage,
			"avatar":          avatar,
			"randomAvatar":    randomAvatar,
			"randomName":      webinfo.RandomName,
			"randomCover":     randomCover,
			"waifuJson":       webinfo.WaifuJson,
			"status":          webinfo.Status,
		}
	}

	response.Success(c, data)
}

// UpdateWebInfo 更新网站信息
// @Summary 更新网站信息
// @Description 管理员更新网站配置
func UpdateWebInfo(c *gin.Context) {
	var req UpdateWebInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	updates := make(map[string]interface{})
	updates["status"] = req.Status

	// 使用显式 updateMode 字段判断更新模式
	// 如果没传 updateMode，则根据 webName 是否为空来判断
	isFullUpdate := req.UpdateMode == "full" || (req.UpdateMode == "" && req.WebName != "")

	if isFullUpdate {
		// 完整更新模式：更新主要配置字段
		updates["web_name"] = req.WebName
		updates["web_title"] = req.WebTitle
		updates["footer"] = req.Footer
		updates["background_image"] = req.BackgroundImage
		updates["avatar"] = req.Avatar
		updates["waifu_json"] = req.WaifuJson
	} else {
		// 部分更新模式：只更新非空字段
		if req.RandomAvatar != "" {
			updates["random_avatar"] = req.RandomAvatar
		}
		if req.RandomName != "" {
			updates["random_name"] = req.RandomName
		}
		if req.RandomCover != "" {
			updates["random_cover"] = req.RandomCover
		}
		if req.Notices != "" {
			updates["notices"] = req.Notices
		}
	}

	if err := config.DB.Model(&models.WebInfo{}).Where("id = ?", req.ID).Updates(updates).Error; err != nil {
		response.Failure(c, "更新失败")
		return
	}

	response.SuccessMessage(c, "更新成功")
}
