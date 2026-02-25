// Package config 应用配置
package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Config 应用配置结构
type Config struct {
	// 数据库配置
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// JWT配置
	JWTSecret string

	// 邮件配置
	EmailFrom     string
	EmailPassword string
	EmailSMTPHost string
	EmailSMTPPort string

	// AI摘要服务配置（OpenAI兼容，支持 DeepSeek / OpenAI / 硅基流动等）
	AIAPIKey string
	AIAPIURL string
	AIModel  string

	// 站点信息配置
	SiteName         string // 站点名称，用于邮件模板等
	SiteURL          string // 博客前端URL
	BossWechat       string // 站长微信号
	EmailHeaderImage string // 邮件头部图片URL

	// 天气API配置（apihz.cn）
	WeatherAPIID  string
	WeatherAPIKey string

	// CORS配置
	CORSAllowedOrigins []string // 生产环境允许的域名白名单

	// 服务器配置
	ServerPort  string
	Environment string // development/production
	DBLogSQL    bool   // 是否打印SQL日志

	// 默认值配置
	DefaultSortID  uint // 删除分类时文章移动到的默认分类
	DefaultLabelID uint // 删除标签时文章移动到的默认标签

	// 验证配置
	MaxUsernameLength     int // 用户名最大长度
	MaxPasswordLength     int // 密码最大长度
	MinPasswordLength     int // 密码最小长度
	MaxArticleTitleLength int // 文章标题最大长度
	MaxCommentLength      int // 评论最大长度
	MaxTreeHoleLength     int // 树洞留言最大长度

	// 限流配置
	RateLimitGlobal          int           // 全局限流：时间窗口内最大请求数
	RateLimitGlobalWindow    time.Duration // 全局限流：时间窗口
	RateLimitLogin           int           // 登录限流：时间窗口内最大尝试次数
	RateLimitLoginWindow     time.Duration // 登录限流：时间窗口
	RateLimitSensitive       int           // 敏感操作限流：时间窗口内最大次数
	RateLimitSensitiveWindow time.Duration // 敏感操作限流：时间窗口
}

// AppConfig 全局应用配置实例
var AppConfig *Config

// LoadConfig 加载配置
// 优先从环境变量读取，环境变量不存在则使用默认值
func LoadConfig() {
	// 尝试加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("提示: 未找到 .env 文件，使用环境变量或默认值")
	}

	AppConfig = &Config{
		// 数据库配置
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "myblog"),

		// JWT配置
		JWTSecret: getEnv("JWT_SECRET", "change-this-secret-in-production"),

		// 邮件配置
		EmailFrom:     getEnv("EMAIL_FROM", ""),
		EmailPassword: getEnv("EMAIL_PASSWORD", ""),
		EmailSMTPHost: getEnv("EMAIL_SMTP_HOST", "smtp.qq.com"),
		EmailSMTPPort: getEnv("EMAIL_SMTP_PORT", "465"),

		// AI摘要服务配置（OpenAI兼容）
		AIAPIKey: getEnv("AI_API_KEY", ""),
		AIAPIURL: getEnv("AI_API_URL", "https://api.deepseek.com/v1"),
		AIModel:  getEnv("AI_MODEL", "deepseek-chat"),

		// 站点信息配置
		SiteName:         getEnv("SITE_NAME", "My-Blog"),
		SiteURL:          getEnv("SITE_URL", "http://localhost:81"),
		BossWechat:       getEnv("BOSS_WECHAT", ""),
		EmailHeaderImage: getEnv("EMAIL_HEADER_IMAGE", ""),

		// 天气API配置（apihz.cn）- 必须在 .env 中配置
		WeatherAPIID:  getEnv("WEATHER_API_ID", ""),
		WeatherAPIKey: getEnv("WEATHER_API_KEY", ""),

		// CORS配置
		CORSAllowedOrigins: getEnvStringSlice("CORS_ALLOWED_ORIGINS", []string{
			"http://localhost", "http://localhost:3000", "http://localhost:5173",
			"http://localhost:8080", "http://127.0.0.1",
		}),

		// 服务器配置
		ServerPort:  getEnv("SERVER_PORT", "8000"),
		Environment: getEnv("ENVIRONMENT", "development"),
		DBLogSQL:    getEnvBool("DB_LOG_SQL", false), // 默认关闭SQL日志

		// 默认值配置
		DefaultSortID:  getEnvUint("DEFAULT_SORT_ID", 11),
		DefaultLabelID: getEnvUint("DEFAULT_LABEL_ID", 23),

		// 验证配置
		MaxUsernameLength:     getEnvInt("MAX_USERNAME_LENGTH", 50),
		MaxPasswordLength:     getEnvInt("MAX_PASSWORD_LENGTH", 128),
		MinPasswordLength:     getEnvInt("MIN_PASSWORD_LENGTH", 6),
		MaxArticleTitleLength: getEnvInt("MAX_ARTICLE_TITLE_LENGTH", 200),
		MaxCommentLength:      getEnvInt("MAX_COMMENT_LENGTH", 2000),
		MaxTreeHoleLength:     getEnvInt("MAX_TREEHOLE_LENGTH", 500),

		// 限流配置
		RateLimitGlobal:          getEnvInt("RATE_LIMIT_GLOBAL", 100),
		RateLimitGlobalWindow:    time.Duration(getEnvInt("RATE_LIMIT_GLOBAL_WINDOW_SECONDS", 1)) * time.Second,
		RateLimitLogin:           getEnvInt("RATE_LIMIT_LOGIN", 5),
		RateLimitLoginWindow:     time.Duration(getEnvInt("RATE_LIMIT_LOGIN_WINDOW_SECONDS", 60)) * time.Second,
		RateLimitSensitive:       getEnvInt("RATE_LIMIT_SENSITIVE", 10),
		RateLimitSensitiveWindow: time.Duration(getEnvInt("RATE_LIMIT_SENSITIVE_WINDOW_SECONDS", 60)) * time.Second,
	}

	// 生产环境安全校验：必须设置 JWT_SECRET
	if AppConfig.JWTSecret == "change-this-secret-in-production" {
		if AppConfig.Environment == "production" {
			log.Fatal("错误: 生产环境必须设置 JWT_SECRET 环境变量，禁止使用默认值")
		}
		log.Println("警告: 请在生产环境中设置 JWT_SECRET 环境变量")
	}
}

// getEnv 获取环境变量，不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt 获取整数环境变量
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

// getEnvUint 获取无符号整数环境变量
func getEnvUint(key string, defaultValue uint) uint {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.ParseUint(value, 10, 32); err == nil {
			return uint(intVal)
		}
	}
	return defaultValue
}

// getEnvStringSlice 获取逗号分隔的字符串列表环境变量
func getEnvStringSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		parts := strings.Split(value, ",")
		result := make([]string, 0, len(parts))
		for _, p := range parts {
			trimmed := strings.TrimSpace(p)
			if trimmed != "" {
				result = append(result, trimmed)
			}
		}
		if len(result) > 0 {
			return result
		}
	}
	return defaultValue
}

// getEnvBool 获取布尔环境变量
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return defaultValue
}
