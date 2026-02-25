// Package handlers 用户相关接口处理
package handlers

import (
	"fmt"
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"
	"time"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// UserListRequest 用户列表请求
type UserListRequest struct {
	Current      int    `json:"current"`      // 当前页码
	Size         int    `json:"size"`         // 每页条数
	SearchKey    string `json:"searchKey"`    // 搜索关键词（手机号）
	UserStatus   string `json:"userStatus"`   // 用户状态筛选
	UserType     string `json:"userType"`     // 用户类型筛选
	IsAdmin      bool   `json:"isAdmin"`      // 是否管理员视图（返回完整信息）
	NoPagination bool   `json:"noPagination"` // true 表示获取全部数据
}

// UpdateUserAttributeRequest 用户属性修改请求
type UpdateUserAttributeRequest struct {
	UserID   uint    `json:"userId" binding:"required"` // 用户ID
	Field    string  `json:"field" binding:"required"`  // 要修改的字段: status/admire/type
	Value    *string `json:"value"`                     // 字符串值（用于admire）
	IntValue *int    `json:"intValue"`                  // 整数值（用于type）
	BoolFlag *bool   `json:"boolFlag"`                  // 布尔值（用于status）
}

// UpdateUserInfoRequest 更新用户信息请求
type UpdateUserInfoRequest struct {
	UserID          *uint   `json:"userId"`          // 用户ID（可选，不传则修改当前登录用户）
	Username        *string `json:"username"`        // 用户名
	Gender          *int    `json:"gender"`          // 性别
	Introduction    *string `json:"introduction"`    // 简介
	PhoneNumber     *string `json:"phoneNumber"`     // 手机号
	Email           *string `json:"email"`           // 邮箱
	Avatar          *string `json:"avatar"`          // 头像
	Code            *string `json:"code"`            // 验证码（修改邮箱时需要）
	QiniuDomain     *string `json:"qiniuDomain"`     // 七牛云域名
	QiniuAccessKey  *string `json:"qiniuAccessKey"`  // 七牛云AccessKey
	QiniuSecretKey  *string `json:"qiniuSecretKey"`  // 七牛云SecretKey
	QiniuBucketName *string `json:"qiniuBucketName"` // 七牛云Bucket名称
}

// ==================== 接口处理函数 ====================

// ListUsers 用户列表接口
// @Summary 获取用户列表
// @Description 获取用户列表，isAdmin=true返回完整信息，否则只返回公开信息
func ListUsers(c *gin.Context) {
	var req UserListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 验证 isAdmin 参数
	isAdmin := isAdminRequest(c, req.IsAdmin)

	query := config.DB.Model(&models.Client{})

	// 只有验证过的管理员才能使用筛选功能
	if isAdmin {
		if req.SearchKey != "" {
			query = query.Where("phone_number LIKE ?", "%"+req.SearchKey+"%")
		}
		if req.UserStatus != "" {
			query = query.Where("user_status = ?", req.UserStatus == "true")
		}
		if req.UserType != "" {
			query = query.Where("user_type = ?", req.UserType)
		}
	}

	var total int64
	query.Count(&total)

	var clients []models.Client
	current, size, offset := normalizePagination(req.Current, req.Size)

	if req.NoPagination {
		// 获取全部数据（前台只返回公开信息，无安全风险）
		query.Order("user_id").Find(&clients)
	} else {
		// 分页查询
		query.Order("user_id").Offset(offset).Limit(size).Find(&clients)
	}

	// 根据验证后的 isAdmin 构建响应数据
	data := make([]map[string]interface{}, 0, len(clients))
	for _, client := range clients {
		if isAdmin {
			// 管理员视图，包含完整信息
			data = append(data, map[string]interface{}{
				"id":           client.UserID,
				"username":     client.Username,
				"phoneNumber":  client.PhoneNumber,
				"email":        client.Email,
				"admire":       client.Admire,
				"userStatus":   client.UserStatus,
				"avatar":       client.Avatar,
				"gender":       client.Gender,
				"introduction": client.Introduction,
				"userType":     client.UserType,
				"createTime":   client.CreateTime,
				"province":     client.Province,
			})
		} else {
			// 前台视图，只包含公开信息（不返回敏感信息）
			data = append(data, map[string]interface{}{
				"id":       client.UserID, // 需要 id 用于评论匹配
				"username": client.Username,
				"avatar":   client.Avatar,
				"gender":   client.Gender,
			})
		}
	}

	response.SuccessPage(c, data, total, current, size)
}

// UpdateUserAttribute 用户属性修改接口
// @Summary 修改用户属性
// @Description 管理员修改用户的状态、赞赏或类型属性
func UpdateUserAttribute(c *gin.Context) {
	var req UpdateUserAttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	switch req.Field {
	case "status":
		// 修改用户状态
		if req.BoolFlag == nil {
			response.Failure(c, "修改状态需要提供boolFlag参数")
			return
		}
		if err := config.DB.Model(&models.Client{}).Where("user_id = ?", req.UserID).Update("user_status", *req.BoolFlag).Error; err != nil {
			response.Failure(c, "更新失败")
			return
		}

	case "admire":
		// 修改赞赏二维码
		if err := config.DB.Model(&models.Client{}).Where("user_id = ?", req.UserID).Update("admire", req.Value).Error; err != nil {
			response.Failure(c, "更新失败")
			return
		}

	case "type":
		// 修改用户类型
		if req.IntValue == nil {
			response.Failure(c, "修改类型需要提供intValue参数")
			return
		}
		userType := *req.IntValue

		// 仅 Boss 才能将用户提升为 Boss
		if userType == 0 {
			currentClient, exists := c.Get("client")
			if !exists {
				response.Failure(c, "未登录")
				return
			}
			if currentClient.(*models.Client).UserType != 0 {
				response.Failure(c, "仅站长可以设置 Boss 权限")
				return
			}
		}

		if err := config.DB.Model(&models.Client{}).Where("user_id = ?", req.UserID).Update("user_type", userType).Error; err != nil {
			response.Failure(c, "更新失败")
			return
		}
		// 同步更新auth_user的is_staff字段
		// Boss(0)、管理员(1)、访客(3) 设置为管理员
		isStaff := userType < 2 || userType == 3
		if err := config.DB.Model(&models.AuthUser{}).Where("id = ?", req.UserID).Update("is_staff", isStaff).Error; err != nil {
			response.Failure(c, "同步用户权限失败")
			return
		}

	default:
		response.Failure(c, "不支持的字段类型，请使用: status/admire/type")
		return
	}

	response.SuccessMessage(c, "更新成功")
}

// UpdateUserInfo 更新用户信息
// @Summary 更新用户信息
// @Description 用户更新自己的个人信息
func UpdateUserInfo(c *gin.Context) {
	var req UpdateUserInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 获取当前登录用户
	currentClient, exists := c.Get("client")
	if !exists {
		response.Failure(c, "未登录")
		return
	}
	currentUser := currentClient.(*models.Client)

	// 确定要修改的用户ID：不传 userId 则修改当前登录用户
	targetUserID := currentUser.UserID
	isModifyingOther := false
	if req.UserID != nil && *req.UserID > 0 && *req.UserID != currentUser.UserID {
		targetUserID = *req.UserID
		isModifyingOther = true
		// 如果要修改他人信息，必须是管理员
		if currentUser.UserType > models.UserTypeAdmin {
			response.Failure(c, "无权修改他人信息")
			return
		}
		// 记录管理员操作日志
		logAdminAction(currentUser.UserID, currentUser.Username, "修改用户信息", targetUserID, "")
	}

	// 查找用户
	var client models.Client
	if err := config.DB.Where("user_id = ?", targetUserID).First(&client).Error; err != nil {
		response.Failure(c, "用户不存在")
		return
	}

	// 构建更新字段
	updates := make(map[string]interface{})

	// 更新用户名（需要检查唯一性）
	if req.Username != nil && *req.Username != "" && *req.Username != client.Username {
		var existingUser models.AuthUser
		if err := config.DB.Where("username = ?", *req.Username).First(&existingUser).Error; err == nil {
			response.Failure(c, "用户名已存在")
			return
		}
		updates["username"] = *req.Username
		// 同步更新auth_user表
		if err := config.DB.Model(&models.AuthUser{}).Where("id = ?", targetUserID).Update("username", *req.Username).Error; err != nil {
			response.Failure(c, "同步用户名失败: "+err.Error())
			return
		}
	}

	// 更新性别
	if req.Gender != nil && (*req.Gender == 0 || *req.Gender == 1 || *req.Gender == 2) {
		updates["gender"] = *req.Gender
	}

	// 更新简介
	if req.Introduction != nil {
		updates["introduction"] = *req.Introduction
	}

	// 更新七牛云配置
	if req.QiniuDomain != nil && *req.QiniuDomain != "" {
		updates["qiniu_domain"] = *req.QiniuDomain
	}
	if req.QiniuAccessKey != nil && *req.QiniuAccessKey != "" {
		updates["qiniu_access_key"] = *req.QiniuAccessKey
	}
	if req.QiniuSecretKey != nil && *req.QiniuSecretKey != "" {
		updates["qiniu_secret_key"] = *req.QiniuSecretKey
	}
	if req.QiniuBucketName != nil && *req.QiniuBucketName != "" {
		updates["qiniu_bucket_name"] = *req.QiniuBucketName
	}

	// 更新手机号
	if req.PhoneNumber != nil && *req.PhoneNumber != "" {
		updates["phone_number"] = *req.PhoneNumber
	}

	// 更新邮箱（需要验证码验证）
	if req.Email != nil && *req.Email != "" && req.Code != nil && *req.Code != "" {
		var code models.Code
		if err := config.DB.Where("email = ?", *req.Email).Order("id DESC").First(&code).Error; err == nil {
			if time.Since(code.CreateTime).Minutes() <= 5 && code.Code == *req.Code {
				updates["email"] = *req.Email
			}
		}
	}

	// 更新头像
	if req.Avatar != nil && *req.Avatar != "" {
		updates["avatar"] = *req.Avatar
	}

	// 执行更新
	if len(updates) > 0 {
		if err := config.DB.Model(&models.Client{}).Where("user_id = ?", targetUserID).Updates(updates).Error; err != nil {
			response.Failure(c, "更新用户信息失败: "+err.Error())
			return
		}
	}

	// 记录管理员修改他人敏感字段的详细日志
	if isModifyingOther {
		logAdminAction(currentUser.UserID, currentUser.Username, "修改用户信息完成", targetUserID,
			fmt.Sprintf("修改字段数: %d", len(updates)))
	}

	// 返回更新后的用户信息
	var updatedClient models.Client
	if err := config.DB.Where("user_id = ?", targetUserID).First(&updatedClient).Error; err != nil {
		response.Failure(c, "查询更新后的用户信息失败")
		return
	}

	avatar := getValidAvatar(updatedClient.Avatar)

	response.Success(c, map[string]interface{}{
		"id":              updatedClient.UserID,
		"username":        updatedClient.Username,
		"phoneNumber":     updatedClient.PhoneNumber,
		"email":           updatedClient.Email,
		"admire":          updatedClient.Admire,
		"userStatus":      updatedClient.UserStatus,
		"avatar":          avatar,
		"gender":          updatedClient.Gender,
		"introduction":    updatedClient.Introduction,
		"userType":        updatedClient.UserType,
		"createTime":      updatedClient.CreateTime,
		"qiniuDomain":     updatedClient.QiniuDomain,
		"qiniuBucketName": updatedClient.QiniuBucketName,
		"qiniuSecretKey":  updatedClient.QiniuSecretKey,
		"qiniuAccessKey":  updatedClient.QiniuAccessKey,
	})
}
