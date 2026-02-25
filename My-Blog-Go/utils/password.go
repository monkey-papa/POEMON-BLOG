package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// CheckDjangoPassword 验证Django风格的密码哈希
func CheckDjangoPassword(password, encoded string) bool {
	parts := strings.Split(encoded, "$")
	if len(parts) != 4 {
		return false
	}

	algorithm := parts[0]
	iterations := parts[1]
	salt := parts[2]
	hash := parts[3]

	if algorithm != "pbkdf2_sha256" {
		return false
	}

	var iterCount int
	fmt.Sscanf(iterations, "%d", &iterCount)

	// 使用PBKDF2-SHA256生成哈希
	dk := pbkdf2.Key([]byte(password), []byte(salt), iterCount, sha256.Size, sha256.New)
	computed := base64.StdEncoding.EncodeToString(dk)

	return computed == hash
}

// MakeDjangoPassword 生成Django风格的密码哈希
func MakeDjangoPassword(password string) string {
	iterations := 720000
	salt := generateSalt(22)

	dk := pbkdf2.Key([]byte(password), []byte(salt), iterations, sha256.Size, sha256.New)
	hash := base64.StdEncoding.EncodeToString(dk)

	return fmt.Sprintf("pbkdf2_sha256$%d$%s$%s", iterations, salt, hash)
}

func generateSalt(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return base64.RawURLEncoding.EncodeToString(bytes)[:length]
}

// GenerateToken 生成Token
func GenerateToken() string {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	return fmt.Sprintf("%x", bytes)
}
