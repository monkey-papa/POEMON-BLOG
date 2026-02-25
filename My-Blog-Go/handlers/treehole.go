// Package handlers 树洞相关接口处理
package handlers

import (
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"
	"my-blog-go/utils"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// TreeHoleListRequest 树洞列表请求
type TreeHoleListRequest struct {
	Current      int  `json:"current"`      // 当前页码
	Size         int  `json:"size"`         // 每页条数
	NoPagination bool `json:"noPagination"` // true 表示获取全部数据
}

// SaveTreeHoleRequest 保存树洞请求
type SaveTreeHoleRequest struct {
	Avatar   string `json:"avatar" binding:"required"`  // 头像
	Username string `json:"username"`                   // 用户名
	Message  string `json:"message" binding:"required"` // 留言内容
}

// ==================== 接口处理函数 ====================

// ListTreeHole 获取树洞列表
// @Summary 获取树洞列表
// @Description 获取所有树洞留言
func ListTreeHole(c *gin.Context) {
	var req TreeHoleListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	var total int64
	config.DB.Model(&models.TreeHole{}).Count(&total)

	var treeHoles []models.TreeHole
	query := config.DB.Model(&models.TreeHole{})

	// 规范化分页参数
	current, size, offset := normalizePagination(req.Current, req.Size)

	if req.NoPagination {
		// 获取全部数据
		query.Find(&treeHoles)
	} else {
		// 分页查询
		query.Offset(offset).Limit(size).Find(&treeHoles)
	}

	data := make([]map[string]interface{}, 0, len(treeHoles))
	for _, th := range treeHoles {
		data = append(data, map[string]interface{}{
			"id":         th.ID,
			"userId":     th.UserID,
			"username":   th.Username,
			"createTime": th.CreateTime,
			"message":    th.Message,
			"avatar":     th.Avatar,
		})
	}

	response.SuccessPage(c, data, total, current, size)
}

// SaveTreeHole 保存树洞
// @Summary 保存树洞留言
// @Description 用户添加树洞留言，记录实际发布者ID
func SaveTreeHole(c *gin.Context) {
	var req SaveTreeHoleRequest
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

	// 验证留言内容
	if err := utils.ValidateTreeHoleMessage(req.Message); err != nil {
		response.Failure(c, err.Error())
		return
	}

	// 清理内容（防XSS）
	req.Message = utils.SanitizeString(req.Message)
	req.Username = utils.SanitizeString(req.Username)

	treeHole := models.TreeHole{
		UserID:   currentUser.UserID, // 记录实际发布者ID
		Avatar:   &req.Avatar,
		Username: req.Username,
		Message:  req.Message,
	}

	if err := config.DB.Create(&treeHole).Error; err != nil {
		response.Failure(c, "保存失败")
		return
	}

	// 通知博主有新的树洞留言（异步，不阻塞响应）
	if currentUser.UserType != models.UserTypeBoss {
		go notifyBossUser(req.Username, req.Message, "treehole")
	}

	// 返回创建的记录
	response.Success(c, map[string]interface{}{
		"id":         treeHole.ID,
		"userId":     treeHole.UserID,
		"username":   treeHole.Username,
		"createTime": treeHole.CreateTime,
		"message":    treeHole.Message,
		"avatar":     treeHole.Avatar,
	})
}

// DeleteTreeHole 删除树洞
// @Summary 删除树洞留言
// @Description 管理员删除树洞留言
func DeleteTreeHole(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if err := config.DB.Delete(&models.TreeHole{}, req.ID).Error; err != nil {
		response.Failure(c, "删除失败")
		return
	}

	response.SuccessMessage(c, "删除成功")
}
