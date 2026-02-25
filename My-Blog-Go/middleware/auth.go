// Package middleware 提供HTTP中间件
package middleware

import (
	"errors"
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"
	"my-blog-go/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// 上下文键名常量
const (
	ContextKeyUser   = "user"   // 用户信息
	ContextKeyClient = "client" // 客户端信息
	ContextKeyClaims = "claims" // JWT声明
)

// AuthMiddleware 认证中间件
// 验证JWT Token，确保用户已登录
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			response.Unauthorized(c, "请先登录")
			c.Abort()
			return
		}

		// 解析JWT
		claims, err := utils.ParseJWT(token)
		if err != nil {
			handleJWTError(c, err)
			return
		}

		// 查询用户信息
		client, user, err := validateUser(claims.UserID)
		if err != nil {
			response.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set(ContextKeyUser, user)
		c.Set(ContextKeyClient, client)
		c.Set(ContextKeyClaims, claims)

		// 检查是否需要刷新Token
		refreshTokenIfNeeded(c, claims)

		c.Next()
	}
}

// AdminMiddleware 管理员权限中间件
// 验证用户是否具有管理员权限
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			response.Unauthorized(c, "请先登录")
			c.Abort()
			return
		}

		// 解析JWT
		claims, err := utils.ParseJWT(token)
		if err != nil {
			handleJWTError(c, err)
			return
		}

		// 快速检查管理员权限（从JWT中判断）
		if !claims.IsStaff || claims.UserType > models.UserTypeAdmin {
			response.Forbidden(c, "需要管理员权限")
			c.Abort()
			return
		}

		// 查询用户信息
		client, user, err := validateUser(claims.UserID)
		if err != nil {
			response.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		// 再次验证权限（防止Token签发后权限被修改）
		if client.UserType > models.UserTypeAdmin {
			response.Forbidden(c, "需要管理员权限")
			c.Abort()
			return
		}

		c.Set(ContextKeyUser, user)
		c.Set(ContextKeyClient, client)
		c.Set(ContextKeyClaims, claims)

		refreshTokenIfNeeded(c, claims)

		c.Next()
	}
}

// OptionalAuthMiddleware 可选认证中间件
// Token可选，有Token则解析用户信息，无Token则继续
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			c.Next()
			return
		}

		claims, err := utils.ParseJWT(token)
		if err != nil {
			c.Next()
			return
		}

		client, user, err := validateUser(claims.UserID)
		if err != nil {
			c.Next()
			return
		}

		c.Set(ContextKeyUser, user)
		c.Set(ContextKeyClient, client)
		c.Set(ContextKeyClaims, claims)

		c.Next()
	}
}

// ==================== 辅助函数 ====================

// extractToken 从请求头提取Token
// 支持 "Token xxx" 和 "Bearer xxx" 格式
func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 {
		return authHeader
	}

	tokenType := strings.ToLower(parts[0])
	if tokenType == "token" || tokenType == "bearer" {
		return parts[1]
	}

	return authHeader
}

// handleJWTError 处理JWT错误
func handleJWTError(c *gin.Context, err error) {
	if errors.Is(err, utils.ErrTokenExpired) {
		response.Unauthorized(c, "登录已过期，请重新登录")
	} else {
		response.Unauthorized(c, "无效的Token")
	}
	c.Abort()
}

// validateUser 验证用户信息
func validateUser(userID uint) (*models.Client, *models.AuthUser, error) {
	var client models.Client
	if err := config.DB.Where("user_id = ?", userID).First(&client).Error; err != nil {
		return nil, nil, errors.New("用户不存在")
	}

	if !client.UserStatus || client.Deleted {
		return nil, nil, errors.New("账户已被禁用或删除")
	}

	var user models.AuthUser
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, nil, errors.New("用户不存在")
	}

	return &client, &user, nil
}

// refreshTokenIfNeeded 如果Token快过期则刷新
func refreshTokenIfNeeded(c *gin.Context, claims *utils.JWTClaims) {
	if newToken, needRefresh := utils.RefreshJWT(claims); needRefresh {
		c.Header("X-New-Token", newToken)
	}
}

// GetCurrentUser 获取当前登录用户
func GetCurrentUser(c *gin.Context) (*models.AuthUser, bool) {
	user, exists := c.Get(ContextKeyUser)
	if !exists {
		return nil, false
	}
	return user.(*models.AuthUser), true
}

// GetCurrentClient 获取当前登录用户的Client信息
func GetCurrentClient(c *gin.Context) (*models.Client, bool) {
	client, exists := c.Get(ContextKeyClient)
	if !exists {
		return nil, false
	}
	return client.(*models.Client), true
}
