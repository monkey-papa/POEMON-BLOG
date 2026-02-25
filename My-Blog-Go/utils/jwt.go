package utils

import (
	"errors"
	"my-blog-go/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 自定义错误
var (
	ErrTokenExpired = errors.New("token expired")
	ErrTokenInvalid = errors.New("token invalid")
)

// JWTClaims JWT 声明
type JWTClaims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	UserType int    `json:"userType"`
	IsStaff  bool   `json:"isStaff"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成 JWT Token
func GenerateJWT(userID uint, username string, userType int, isStaff bool) (string, error) {
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		UserType: userType,
		IsStaff:  isStaff,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7天过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-blog-go",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

// ParseJWT 解析 JWT Token
func ParseJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil {
		// 检查是否是过期错误
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshJWT 刷新 JWT Token（如果快过期则生成新的）
func RefreshJWT(claims *JWTClaims) (string, bool) {
	// 如果距离过期时间小于1天，则刷新
	if time.Until(claims.ExpiresAt.Time) < 24*time.Hour {
		newToken, err := GenerateJWT(claims.UserID, claims.Username, claims.UserType, claims.IsStaff)
		if err != nil {
			return "", false
		}
		return newToken, true
	}
	return "", false
}
