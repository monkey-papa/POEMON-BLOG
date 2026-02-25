// Package utils 邮件工具
package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"math/big"
	"my-blog-go/config"
	"strconv"

	"gopkg.in/gomail.v2"
)

// ==================== 邮件模板类型 ====================

// EmailType 邮件类型
type EmailType int

const (
	EmailTypeVerification EmailType = iota // 验证码邮件
	EmailTypeComment                       // 评论通知邮件
	EmailTypeApprove                       // 审核通知邮件
)

// ==================== 公共函数 ====================

// GenerateCode 生成6位随机验证码
func GenerateCode() string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, 6)
	for i := range code {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		code[i] = charset[n.Int64()]
	}
	return string(code)
}

// SendVerificationCode 发送验证码邮件
func SendVerificationCode(toAddress, code string) error {
	content := buildEmailContent(EmailTypeVerification, map[string]string{
		"code": code,
	})
	return sendEmail(toAddress, "验证码", content)
}

// SendCommentNotification 发送评论通知邮件
func SendCommentNotification(toAddress, comment, name string) error {
	content := buildEmailContent(EmailTypeComment, map[string]string{
		"name":    name,
		"comment": comment,
	})
	return sendEmail(toAddress, "评论通知", content)
}

// SendApproveNotification 发送审核通知邮件
func SendApproveNotification(toAddress, comment, name string) error {
	content := buildEmailContent(EmailTypeApprove, map[string]string{
		"name":    name,
		"comment": comment,
	})
	return sendEmail(toAddress, "文章审核", content)
}

// ==================== 内部函数 ====================

// sendEmail 发送邮件（核心函数）
func sendEmail(toAddress, subject, htmlContent string) error {
	cfg := config.AppConfig

	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EmailFrom)
	m.SetHeader("To", toAddress)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlContent)

	port, _ := strconv.Atoi(cfg.EmailSMTPPort)
	d := gomail.NewDialer(cfg.EmailSMTPHost, port, cfg.EmailFrom, cfg.EmailPassword)
	d.TLSConfig = &tls.Config{
		ServerName: cfg.EmailSMTPHost,
	}

	return d.DialAndSend(m)
}

// buildEmailContent 构建邮件HTML内容
func buildEmailContent(emailType EmailType, data map[string]string) string {
	cfg := config.AppConfig
	siteName := cfg.SiteName
	wechat := cfg.BossWechat

	// 邮件头部模板（头部图片从配置读取）
	headerImageStyle := ""
	if cfg.EmailHeaderImage != "" {
		headerImageStyle = fmt.Sprintf(` background-image: url('%s'); background-size: cover;`, cfg.EmailHeaderImage)
	}
	header := fmt.Sprintf(`<div style="font-family: serif; line-height: 22px; padding: 30px">
		  <div style="display: flex; justify-content: center; width: 100%%%%; max-width: 900px;%s border-radius: 10px"></div>
  <div style="margin-top: 20px; display: flex; flex-direction: column; align-items: center">`, headerImageStyle)

	// 邮件尾部模板
	footer := `<div style="margin-top: 20px; font-size: 12px; color: black">此邮件由系统自动发出，直接回复无效。</div>
  </div>
</div>`

	footerWithContact := fmt.Sprintf(`<div style="margin-top: 20px; font-size: 12px; color: black">此邮件由 %s 自动发出，直接回复无效，有问题请联系站长vx：%s。</div>
  </div>
</div>`, siteName, wechat)

	var body string

	switch emailType {
	case EmailTypeVerification:
		body = fmt.Sprintf(`
			<div style="margin: 10px auto 20px; text-align: center">
			  <div style="line-height: 32px; font-size: 26px; font-weight: bold; color: #000000">嘿！你在博客中收到一条新消息。</div>
			  <div style="font-size: 16px; font-weight: bold; color: rgba(0, 0, 0, 0.19); margin-top: 21px">你收到来自博客的消息</div>
			</div>
			<div style="min-width: 250px; max-width: 800px; min-height: 128px; background: #f7f7f7; border-radius: 10px; padding: 32px">
			  <div>
				<div style="font-size: 18px; font-weight: bold; color: #c5343e">博客</div>
				<div style="margin-top: 6px; font-size: 16px; color: #000000">
				  <p><abc style="color: #c5343e">%s</abc>为本次验证的验证码，请在5分钟内完成验证。为保证账号安全，请勿泄漏此验证码。</p>
				</div>
			  </div>
    </div>`, data["code"])
		return header + body + footer

	case EmailTypeComment:
		body = fmt.Sprintf(`
			<div style="margin: 10px auto 20px; text-align: center">
			  <div style="line-height: 32px; font-size: 26px; font-weight: bold; color: #000000">嘿！你在 %s 中收到一条新评论。</div>
			  <div style="font-size: 16px; font-weight: bold; color: rgba(0, 0, 0, 0.19); margin-top: 21px">你收到来自 <abc style="color: #c5343e">%s</abc> 的评论</div>
			</div>
			<div style="min-width: 250px; max-width: 800px; min-height: 128px; background: #f7f7f7; border-radius: 10px; padding: 32px">
			  <div>
				<div style="font-size: 18px; font-weight: bold; color: #c5343e">%s</div>
				<div style="margin-top: 6px; font-size: 16px; color: #000000">
				  <p>【%s】<abc style="color: #c5343e">%s</abc></p>
				</div>
			  </div>
      %s
    </div>`, siteName, data["name"], data["name"], siteName, data["comment"], buildViewButton())
		return header + body + footerWithContact

	case EmailTypeApprove:
		body = fmt.Sprintf(`
			<div style="margin: 10px auto 20px; text-align: center">
			  <div style="line-height: 32px; font-size: 26px; font-weight: bold; color: #000000">嘿！你在 %s 中收到一条文章待审核通知。</div>
			  <div style="font-size: 16px; font-weight: bold; color: rgba(0, 0, 0, 0.19); margin-top: 21px"><abc style="color: #c5343e">%s</abc> 的文章需要你审核</div>
			</div>
			<div style="min-width: 250px; max-width: 800px; min-height: 128px; background: #f7f7f7; border-radius: 10px; padding: 32px">
			  <div>
				<div style="font-size: 18px; font-weight: bold; color: #c5343e">%s</div>
				<div style="margin-top: 6px; font-size: 16px; color: #000000">
				  <p>【%s】<abc style="color: #c5343e">%s</abc></p>
				</div>
			  </div>
      %s
    </div>`, siteName, data["name"], data["name"], siteName, data["comment"], buildViewButton())
		return header + body + footerWithContact
	}

	return ""
}

// buildViewButton 构建"查看详情"按钮（链接地址从配置读取）
func buildViewButton() string {
	siteURL := config.AppConfig.SiteURL
	return fmt.Sprintf(`<a style="width: 150px; height: 38px; background: #ef859d38; border-radius: 32px; display: flex; align-items: center; justify-content: center; text-decoration: none; margin: 40px auto 0" href="%s" target="_blank" rel="noopener">
				<span style="color: #db214b">查看详情</span>
  </a>`, siteURL)
}
