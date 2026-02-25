// Package handlers 认证相关接口处理
package handlers

import (
	"log"
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"
	"my-blog-go/utils"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ==================== 验证辅助函数 ====================

// validateLogin 验证登录参数
func validateLogin(account, password string) error {
	return utils.ValidateLoginRequest(account, password)
}

// validateRegister 验证注册参数
func validateRegister(username, password, email, code string) error {
	return utils.ValidateRegisterRequest(username, password, email, code)
}

// ==================== 请求结构体 ====================

// LoginRequest 登录请求
type LoginRequest struct {
	Account  string `json:"account" binding:"required"`  // 账号（用户名或邮箱）
	Password string `json:"password" binding:"required"` // 密码
	Province string `json:"province"`                    // 省份信息
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	Email    string `json:"email" binding:"required"`    // 邮箱
	Code     string `json:"code" binding:"required"`     // 验证码
	Province string `json:"province"`                    // 省份信息
}

// ForgetPasswordRequest 忘记密码请求
type ForgetPasswordRequest struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 新密码
	Email    string `json:"email" binding:"required"`    // 邮箱
	Code     string `json:"code" binding:"required"`     // 验证码
}

// SendCodeRequest 发送验证码请求
type SendCodeRequest struct {
	Email string `json:"email" binding:"required,email"` // 邮箱地址
}

// SendCommentNotifyRequest 发送评论通知请求
type SendCommentNotifyRequest struct {
	Comment string `json:"comment" binding:"required"` // 评论内容
	Name    string `json:"name" binding:"required"`    // 评论者名称
	Email   string `json:"email"`                      // 接收者邮箱（可选，优先使用）
	UserId  uint   `json:"userId"`                     // 接收者用户ID（可选，后端查邮箱）
	Type    string `json:"type"`                       // 通知类型: comment/approve
}

// ==================== 接口处理函数 ====================

// FrontLogin 前台登录
// @Summary 前台用户登录
// @Description 支持用户名或邮箱登录，返回用户信息和Token
func FrontLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "请输入账号和密码")
		return
	}

	// 参数验证
	if err := validateLogin(req.Account, req.Password); err != nil {
		response.Failure(c, err.Error())
		return
	}

	var user models.AuthUser
	var client models.Client

	// 判断是邮箱还是用户名登录
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)*@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	if emailRegex.MatchString(req.Account) {
		// 邮箱登录
		if err := config.DB.Where("email = ?", req.Account).First(&client).Error; err != nil {
			response.Failure(c, "用户名或密码错误")
			return
		}
		if err := config.DB.First(&user, client.UserID).Error; err != nil {
			response.Failure(c, "用户名或密码错误")
			return
		}
	} else {
		// 用户名登录
		if err := config.DB.Where("username = ?", req.Account).First(&user).Error; err != nil {
			response.Failure(c, "用户名或密码错误")
			return
		}
		if err := config.DB.Where("user_id = ?", user.ID).First(&client).Error; err != nil {
			response.Failure(c, "用户名或密码错误")
			return
		}
	}

	// 检查账户状态
	if !client.UserStatus || client.Deleted {
		response.Failure(c, "账户已被禁用或删除")
		return
	}

	// 验证密码
	if !utils.CheckDjangoPassword(req.Password, user.Password) {
		response.Failure(c, "用户名或密码错误")
		return
	}

	// 生成JWT Token
	token, err := utils.GenerateJWT(user.ID, user.Username, client.UserType, user.IsStaff)
	if err != nil {
		response.Failure(c, "登录失败，请稍后重试")
		return
	}

	// 更新省份信息（首次登录时记录）
	if req.Province != "" && (client.Province == nil || *client.Province == "") {
		config.DB.Model(&client).Update("province", req.Province)
	}

	// 构建响应（统一返回对象格式）
	response.Success(c, buildUserResponse(&client, token))
}

// AdminLogin 后台登录
// @Summary 管理员登录
// @Description 管理员专用登录接口，需要管理员权限
func AdminLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "请输入账号和密码")
		return
	}

	// 参数验证
	if err := validateLogin(req.Account, req.Password); err != nil {
		response.Failure(c, err.Error())
		return
	}

	var user models.AuthUser
	if err := config.DB.Where("username = ?", req.Account).First(&user).Error; err != nil {
		response.Failure(c, "用户名或密码错误，或非管理员账户")
		return
	}

	var client models.Client
	if err := config.DB.Where("user_id = ?", user.ID).First(&client).Error; err != nil {
		response.Failure(c, "用户名或密码错误，或非管理员账户")
		return
	}

	// 检查后台登录权限：Boss(0)、Admin(1)、Visitor(3) 可以登录，普通用户(2)不能
	if client.UserType == models.UserTypeNormal || !client.UserStatus {
		response.Failure(c, "用户名或密码错误，或非管理员账户")
		return
	}

	// 验证密码
	if !utils.CheckDjangoPassword(req.Password, user.Password) {
		response.Failure(c, "用户名或密码错误，或非管理员账户")
		return
	}

	// 生成JWT Token
	token, err := utils.GenerateJWT(user.ID, user.Username, client.UserType, user.IsStaff)
	if err != nil {
		response.Failure(c, "登录失败，请稍后重试")
		return
	}

	// 构建响应（统一返回对象格式）
	response.Success(c, buildUserResponse(&client, token))
}

// Register 用户注册
// @Summary 用户注册
// @Description 新用户注册，需要邮箱验证码
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "请填写完整注册信息")
		return
	}

	// 参数验证
	if err := validateRegister(req.Username, req.Password, req.Email, req.Code); err != nil {
		response.Failure(c, err.Error())
		return
	}

	// 检查用户名是否已存在
	var count int64
	config.DB.Model(&models.AuthUser{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		response.Failure(c, "用户名已存在")
		return
	}

	// 检查邮箱是否已存在
	config.DB.Model(&models.Client{}).Where("email = ?", req.Email).Count(&count)
	if count > 0 {
		response.Failure(c, "邮箱已被注册")
		return
	}

	// 验证验证码
	if !verifyCode(req.Email, req.Code) {
		response.Failure(c, "验证码错误或已过期")
		return
	}

	// 使用事务创建用户，确保原子性
	var user models.AuthUser
	var client models.Client
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// 创建 AuthUser
		user = models.AuthUser{
			Username:   req.Username,
			Password:   utils.MakeDjangoPassword(req.Password),
			IsActive:   true,
			DateJoined: time.Now(),
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		// 创建 Client 记录
		client = models.Client{
			UserID:     user.ID,
			Username:   req.Username,
			Email:      &req.Email,
			Province:   &req.Province,
			UserStatus: true,
		}
		if err := tx.Create(&client).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		response.Failure(c, "注册失败，请稍后重试")
		return
	}

	// 生成JWT Token
	token, err := utils.GenerateJWT(user.ID, user.Username, client.UserType, user.IsStaff)
	if err != nil {
		response.Failure(c, "注册成功，但自动登录失败")
		return
	}

	// 构建响应（统一返回对象格式）
	response.Success(c, buildUserResponse(&client, token))
}

// ForgetPassword 忘记密码
// @Summary 重置密码
// @Description 通过邮箱验证码重置密码
func ForgetPassword(c *gin.Context) {
	var req ForgetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "请填写完整信息")
		return
	}

	// 检查用户是否存在
	var user models.AuthUser
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		response.Failure(c, "用户不存在")
		return
	}

	// 校验邮箱归属：确保请求中的邮箱与该用户注册邮箱一致
	var client models.Client
	if err := config.DB.Where("user_id = ?", user.ID).First(&client).Error; err != nil {
		response.Failure(c, "用户信息异常")
		return
	}
	if client.Email == nil || *client.Email != req.Email {
		response.Failure(c, "邮箱与该用户注册邮箱不匹配")
		return
	}

	// 验证验证码
	if !verifyCode(req.Email, req.Code) {
		response.Failure(c, "验证码错误或已过期")
		return
	}

	// 密码已由前端 AES 加密，直接哈希加密后的字符串
	user.Password = utils.MakeDjangoPassword(req.Password)
	if err := config.DB.Save(&user).Error; err != nil {
		response.Failure(c, "密码重置失败")
		return
	}

	response.SuccessMessage(c, "密码重置成功")
}

// SendCode 发送验证码
// @Summary 发送邮箱验证码
// @Description 向指定邮箱发送验证码
func SendCode(c *gin.Context) {
	var req SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "请输入有效的邮箱地址")
		return
	}

	email := strings.TrimSpace(req.Email)

	// 参数验证
	if err := utils.ValidateEmail(email); err != nil {
		response.Failure(c, err.Error())
		return
	}

	// 生成验证码
	code := strings.ToUpper(utils.GenerateCode())

	// 发送邮件
	if err := utils.SendVerificationCode(email, code); err != nil {
		log.Printf("[SendCode] 邮件发送失败: email=%s, err=%v", email, err)
		response.Failure(c, "验证码发送失败，请稍后重试")
		return
	}

	// 保存验证码到数据库
	codeRecord := models.Code{
		Email: email,
		Code:  code,
	}
	if err := config.DB.Create(&codeRecord).Error; err != nil {
		log.Printf("[SendCode] 保存验证码失败: email=%s, err=%v", email, err)
		response.Failure(c, "验证码发送成功但保存失败，请重试")
		return
	}

	response.SuccessMessage(c, "验证码已发送")
}

// SendCodeComment 发送评论通知
// @Summary 发送评论通知邮件
// @Description 向用户发送评论或审核通知邮件，支持直接传邮箱或传用户ID
func SendCodeComment(c *gin.Context) {
	var req SendCommentNotifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 确定接收者邮箱
	email := req.Email
	if email == "" && req.UserId > 0 {
		// 根据 userId 查询邮箱（无论用户是否启用都能获取）
		var client models.Client
		if err := config.DB.First(&client, req.UserId).Error; err != nil {
			response.Failure(c, "用户不存在")
			return
		}
		if client.Email == nil || *client.Email == "" {
			response.Failure(c, "用户未设置邮箱")
			return
		}
		email = *client.Email
	}

	if email == "" {
		response.Failure(c, "请提供邮箱或用户ID")
		return
	}

	var err error
	if req.Type == "approve" {
		err = utils.SendApproveNotification(email, req.Comment, req.Name)
	} else {
		err = utils.SendCommentNotification(email, req.Comment, req.Name)
	}

	if err != nil {
		response.Failure(c, "通知发送失败: "+err.Error())
		return
	}

	response.SuccessMessage(c, "通知已发送")
}

// ==================== 辅助函数 ====================

// verifyCode 验证验证码
// 验证码有效期为5分钟
func verifyCode(email, code string) bool {
	var codeRecord models.Code
	if err := config.DB.Where("email = ?", email).Order("create_time DESC").First(&codeRecord).Error; err != nil {
		return false
	}

	// 检查是否过期（5分钟）
	if time.Since(codeRecord.CreateTime) > 5*time.Minute {
		return false
	}

	return codeRecord.Code == code
}

// buildUserResponse 构建用户响应数据
func buildUserResponse(client *models.Client, token string) map[string]interface{} {
	avatar := getValidAvatar(client.Avatar)
	return map[string]interface{}{
		"accessToken":     token,
		"id":              client.UserID,
		"username":        client.Username,
		"phoneNumber":     client.PhoneNumber,
		"email":           client.Email,
		"admire":          client.Admire,
		"userStatus":      client.UserStatus,
		"avatar":          avatar,
		"gender":          client.Gender,
		"introduction":    client.Introduction,
		"userType":        client.UserType,
		"createTime":      client.CreateTime,
		"qiniuDomain":     client.QiniuDomain,
		"qiniuBucketName": client.QiniuBucketName,
		"qiniuSecretKey":  client.QiniuSecretKey,
		"qiniuAccessKey":  client.QiniuAccessKey,
	}
}
