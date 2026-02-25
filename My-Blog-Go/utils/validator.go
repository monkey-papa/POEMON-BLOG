// Package utils 验证工具
package utils

import (
	"errors"
	"my-blog-go/config"
	"my-blog-go/models"
	"regexp"
	"strings"
	"unicode/utf8"
)

// ==================== 验证错误定义 ====================

var (
	ErrUsernameEmpty    = errors.New("用户名不能为空")
	ErrUsernameTooLong  = errors.New("用户名过长")
	ErrUsernameInvalid  = errors.New("用户名只能包含字母、数字、下划线和中文")
	ErrPasswordEmpty    = errors.New("密码不能为空")
	ErrPasswordTooShort = errors.New("密码过短")
	ErrPasswordTooLong  = errors.New("密码过长")
	ErrEmailEmpty       = errors.New("邮箱不能为空")
	ErrEmailInvalid     = errors.New("邮箱格式不正确")
	ErrCodeEmpty        = errors.New("验证码不能为空")
	ErrCodeInvalid      = errors.New("验证码格式不正确")
	ErrTitleEmpty       = errors.New("标题不能为空")
	ErrTitleTooLong     = errors.New("标题过长")
	ErrContentEmpty     = errors.New("内容不能为空")
	ErrContentTooLong   = errors.New("内容过长")
	ErrCommentTooLong   = errors.New("评论内容过长")
	ErrTreeHoleTooLong  = errors.New("留言内容过长")
	ErrURLInvalid       = errors.New("URL格式不正确")
	ErrPhoneInvalid     = errors.New("手机号格式不正确")
	ErrIDRequired       = errors.New("ID不能为空")
)

// ==================== 验证函数 ====================

// ValidateUsername 验证用户名
func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)
	if username == "" {
		return ErrUsernameEmpty
	}
	if utf8.RuneCountInString(username) > config.AppConfig.MaxUsernameLength {
		return ErrUsernameTooLong
	}
	// 只允许字母、数字、下划线、中文
	matched, _ := regexp.MatchString(`^[\w\p{Han}]+$`, username)
	if !matched {
		return ErrUsernameInvalid
	}
	return nil
}

// ValidatePassword 验证密码
func ValidatePassword(password string) error {
	if password == "" {
		return ErrPasswordEmpty
	}
	if len(password) < config.AppConfig.MinPasswordLength {
		return ErrPasswordTooShort
	}
	if len(password) > config.AppConfig.MaxPasswordLength {
		return ErrPasswordTooLong
	}
	return nil
}

// ValidateEmail 验证邮箱
func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return ErrEmailEmpty
	}
	// 邮箱正则
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrEmailInvalid
	}
	return nil
}

// ValidateCode 验证验证码
func ValidateCode(code string) error {
	code = strings.TrimSpace(code)
	if code == "" {
		return ErrCodeEmpty
	}
	// 验证码为6位字母数字
	if len(code) != 6 {
		return ErrCodeInvalid
	}
	matched, _ := regexp.MatchString(`^[A-Z0-9]{6}$`, strings.ToUpper(code))
	if !matched {
		return ErrCodeInvalid
	}
	return nil
}

// ValidateArticleTitle 验证文章标题
func ValidateArticleTitle(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return ErrTitleEmpty
	}
	if utf8.RuneCountInString(title) > config.AppConfig.MaxArticleTitleLength {
		return ErrTitleTooLong
	}
	return nil
}

// ValidateComment 验证评论内容
func ValidateComment(content string) error {
	content = strings.TrimSpace(content)
	if content == "" {
		return ErrContentEmpty
	}
	if utf8.RuneCountInString(content) > config.AppConfig.MaxCommentLength {
		return ErrCommentTooLong
	}
	// 检查禁用词
	if err := CheckProhibitedWords(content); err != nil {
		return err
	}
	return nil
}

// ValidateTreeHoleMessage 验证树洞留言
func ValidateTreeHoleMessage(message string) error {
	message = strings.TrimSpace(message)
	if message == "" {
		return ErrContentEmpty
	}
	if utf8.RuneCountInString(message) > config.AppConfig.MaxTreeHoleLength {
		return ErrTreeHoleTooLong
	}
	// 检查禁用词
	if err := CheckProhibitedWords(message); err != nil {
		return err
	}
	return nil
}

// CheckProhibitedWords 检查内容是否包含禁用词
func CheckProhibitedWords(content string) error {
	// 从数据库获取禁用词列表
	var words []models.Words
	if err := config.DB.Find(&words).Error; err != nil {
		// 数据库查询失败时不阻止提交
		return nil
	}

	// 转换为小写进行不区分大小写匹配
	lowerContent := strings.ToLower(content)

	for _, word := range words {
		if word.Message != "" && strings.Contains(lowerContent, strings.ToLower(word.Message)) {
			return errors.New("内容包含违禁词，请修改后重试")
		}
	}
	return nil
}

// ValidateURL 验证URL格式
func ValidateURL(url string) error {
	if url == "" {
		return nil // URL可选
	}
	// URL正则
	urlRegex := regexp.MustCompile(`^(https?://)?[\w.-]+(\.[\w.-]+)+[\w\-._~:/?#\[\]@!$&'()*+,;=]*$`)
	if !urlRegex.MatchString(url) {
		return ErrURLInvalid
	}
	return nil
}

// ValidatePhone 验证手机号
func ValidatePhone(phone string) error {
	if phone == "" {
		return nil // 手机号可选
	}
	// 中国大陆手机号正则
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if !phoneRegex.MatchString(phone) {
		return ErrPhoneInvalid
	}
	return nil
}

// ValidateID 验证ID
func ValidateID(id uint) error {
	if id == 0 {
		return ErrIDRequired
	}
	return nil
}

// ==================== 组合验证函数 ====================

// ValidateLoginRequest 验证登录请求
func ValidateLoginRequest(account, password string) error {
	if strings.TrimSpace(account) == "" {
		return errors.New("账号不能为空")
	}
	if password == "" {
		return ErrPasswordEmpty
	}
	return nil
}

// ValidateRegisterRequest 验证注册请求
func ValidateRegisterRequest(username, password, email, code string) error {
	if err := ValidateUsername(username); err != nil {
		return err
	}
	if err := ValidatePassword(password); err != nil {
		return err
	}
	if err := ValidateEmail(email); err != nil {
		return err
	}
	if err := ValidateCode(code); err != nil {
		return err
	}
	return nil
}

// SanitizeString 清理字符串（去除首尾空格，防止XSS）
func SanitizeString(s string) string {
	s = strings.TrimSpace(s)
	// 基本的HTML转义
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

// SanitizeHTML 清理HTML内容（保留基本格式，移除危险标签）
// 用于文章内容等需要保留格式的场景
func SanitizeHTML(html string) string {
	// 移除script标签
	scriptRegex := regexp.MustCompile(`(?i)<script[^>]*>[\s\S]*?</script>`)
	html = scriptRegex.ReplaceAllString(html, "")

	// 移除on事件属性
	onEventRegex := regexp.MustCompile(`(?i)\s+on\w+\s*=\s*["'][^"']*["']`)
	html = onEventRegex.ReplaceAllString(html, "")

	// 移除javascript:协议
	jsProtocolRegex := regexp.MustCompile(`(?i)javascript:`)
	html = jsProtocolRegex.ReplaceAllString(html, "")

	return html
}
