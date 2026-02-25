// Package handlers 文章相关接口处理
package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"my-blog-go/config"
	"my-blog-go/middleware"
	"my-blog-go/models"
	"my-blog-go/response"
	"my-blog-go/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ==================== 请求结构体 ====================

// ArticleListRequest 文章列表请求
type ArticleListRequest struct {
	Current         int    `json:"current"`         // 当前页码
	Size            int    `json:"size"`            // 每页条数
	SearchKey       string `json:"searchKey"`       // 搜索关键词
	RecommendStatus *bool  `json:"recommendStatus"` // 推荐状态筛选
	LabelID         *uint  `json:"labelId"`         // 标签ID筛选
	SortID          *uint  `json:"sortId"`          // 分类ID筛选
	ViewAll         bool   `json:"viewAll"`         // 是否查看所有（包括隐藏），需管理员权限
}

// ArticleDetailRequest 文章详情请求
type ArticleDetailRequest struct {
	ID       uint   `json:"id" binding:"required"` // 文章ID
	Password string `json:"password"`              // 文章密码
	IsAdmin  bool   `json:"isAdmin"`               // 是否管理员视图（跳过状态检查和密码验证）
}

// ArticleSaveRequest 保存文章请求
type ArticleSaveRequest struct {
	ArticleTitle    string  `json:"articleTitle" binding:"required"` // 文章标题
	ArticleAuthor   string  `json:"articleAuthor"`                   // 作者名
	ArticleContent  string  `json:"articleContent"`                  // 文章内容
	RecommendStatus bool    `json:"recommendStatus"`                 // 推荐状态
	ViewStatus      bool    `json:"viewStatus"`                      // 可见状态
	CommentStatus   bool    `json:"commentStatus"`                   // 评论状态
	Password        *string `json:"password"`                        // 访问密码
	ArticleCover    *string `json:"articleCover"`                    // 封面图
	SortID          uint    `json:"sortId"`                          // 分类ID
	LabelID         uint    `json:"labelId"`                         // 标签ID
}

// ArticleUpdateRequest 更新文章请求
// 注意：不允许通过此接口修改 userId（文章归属），只能修改 articleAuthor（显示作者名）
type ArticleUpdateRequest struct {
	ID              uint    `json:"id" binding:"required"` // 文章ID
	ArticleTitle    *string `json:"articleTitle"`          // 文章标题
	ArticleAuthor   *string `json:"articleAuthor"`         // 作者名
	ArticleContent  *string `json:"articleContent"`        // 文章内容
	RecommendStatus *bool   `json:"recommendStatus"`       // 推荐状态
	ViewStatus      *bool   `json:"viewStatus"`            // 可见状态
	CommentStatus   *bool   `json:"commentStatus"`         // 评论状态
	Password        *string `json:"password"`              // 访问密码
	ArticleCover    *string `json:"articleCover"`          // 封面图
	SortID          *uint   `json:"sortId"`                // 分类ID
	LabelID         *uint   `json:"labelId"`               // 标签ID
}

// ArticleStatusRequest 修改文章状态请求
type ArticleStatusRequest struct {
	ArticleID       uint  `json:"articleId" binding:"required"` // 文章ID
	ViewStatus      *bool `json:"viewStatus"`                   // 可见状态
	RecommendStatus *bool `json:"recommendStatus"`              // 推荐状态
	CommentStatus   *bool `json:"commentStatus"`                // 评论状态
}

// ArticleLikeRequest 文章点赞请求
type ArticleLikeRequest struct {
	ArticleID uint `json:"articleId" binding:"required"` // 文章ID
}

// ArticleSummaryRequest 获取文章摘要请求
type ArticleSummaryRequest struct {
	Message   string `json:"message"`   // 文章内容
	ArticleID uint   `json:"articleId"` // 文章ID
}

// ==================== 接口处理函数 ====================

// ListArticles 文章列表接口
// @Summary 获取文章列表
// @Description 获取文章列表，viewAll=true查看所有文章（需管理员权限），否则只返回可见文章
func ListArticles(c *gin.Context) {
	var req ArticleListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 验证 viewAll 参数
	viewAll := isAdminRequest(c, req.ViewAll)

	query := config.DB.Model(&models.Article{})

	// 非管理员模式只查询可见文章
	if !viewAll {
		query = query.Where("view_status = ?", true)
	}

	// 应用筛选条件
	if req.LabelID != nil && *req.LabelID > 0 {
		query = query.Where("label_id = ?", *req.LabelID)
	}
	if req.SortID != nil && *req.SortID > 0 {
		query = query.Where("sort_id = ?", *req.SortID)
	}
	if req.RecommendStatus != nil {
		query = query.Where("recommend_status = ?", *req.RecommendStatus)
	}
	if req.SearchKey != "" {
		query = query.Where("article_title LIKE ?", "%"+req.SearchKey+"%")
	}

	// 查询总数
	var total int64
	query.Count(&total)

	// 规范化分页参数并查询
	var articles []models.Article
	current, size, offset := normalizePagination(req.Current, req.Size)
	query.Order("create_time DESC").Offset(offset).Limit(size).Find(&articles)

	// 验证后的 viewAll 模式不检查封面状态
	data := buildArticleList(articles, !viewAll)
	response.SuccessPage(c, data, total, current, size)
}

// GetArticleDetail 文章详情接口
// @Summary 获取文章详情
// @Description 获取文章详情，isAdmin=true跳过状态和密码验证
func GetArticleDetail(c *gin.Context) {
	var req ArticleDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 验证 isAdmin 参数
	isAdmin := isAdminRequest(c, req.IsAdmin)

	var article models.Article

	if isAdmin {
		// 管理员模式：直接获取文章（已验证 Token）
		if err := config.DB.First(&article, req.ID).Error; err != nil {
			response.Failure(c, "文章不存在")
			return
		}
		// 返回管理员视图
		data := map[string]interface{}{
			"id":              article.ID,
			"userId":          article.UserID,
			"sortId":          article.SortID,
			"labelId":         article.LabelID,
			"articleCover":    article.ArticleCover,
			"articleTitle":    article.ArticleTitle,
			"articleContent":  article.ArticleContent,
			"articleAuthor":   article.UpdateBy,
			"viewStatus":      article.ViewStatus,
			"password":        article.Password,
			"recommendStatus": article.RecommendStatus,
			"commentStatus":   article.CommentStatus,
		}
		response.Success(c, data)
		return
	}

	// 前台模式：只能查看公开文章（view_status = true）
	if err := config.DB.Where("id = ? AND view_status = ?", req.ID, true).First(&article).Error; err != nil {
		response.Failure(c, "文章不存在或未公开")
		return
	}

	// 验证文章密码
	if article.Password != nil && *article.Password != "" {
		if req.Password != *article.Password {
			response.Failure(c, "密码错误")
			return
		}
	}

	// 增加浏览量
	config.DB.Model(&article).UpdateColumn("view_count", article.ViewCount+1)

	// 检查用户是否已点赞（从 token 获取当前用户）
	likeStatus := 0
	if client, exists := c.Get("client"); exists {
		if currentUser, ok := client.(*models.Client); ok && currentUser.UserID > 0 {
			var like models.ArticleLike
			if err := config.DB.Where("user_id_id = ? AND article_id_id = ?", currentUser.UserID, req.ID).First(&like).Error; err == nil {
				likeStatus = 1
			}
		}
	}

	data := buildArticleDetail(article, likeStatus)
	response.Success(c, data)
}

// SaveArticle 保存文章
// @Summary 创建新文章
// @Description 管理员创建新文章
func SaveArticle(c *gin.Context) {
	var req ArticleSaveRequest
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

	// 验证文章标题
	if err := utils.ValidateArticleTitle(req.ArticleTitle); err != nil {
		response.Failure(c, err.Error())
		return
	}

	// 清理文章内容（防XSS）
	req.ArticleContent = utils.SanitizeHTML(req.ArticleContent)

	article := models.Article{
		ArticleTitle:    req.ArticleTitle,
		UserID:          currentUser.UserID, // 直接使用当前登录用户ID
		UpdateBy:        &req.ArticleAuthor,
		ArticleContent:  req.ArticleContent,
		RecommendStatus: req.RecommendStatus,
		ViewStatus:      req.ViewStatus,
		CommentStatus:   req.CommentStatus,
		Password:        req.Password,
		ArticleCover:    req.ArticleCover,
		SortID:          req.SortID,
		LabelID:         req.LabelID,
	}

	if err := config.DB.Create(&article).Error; err != nil {
		response.Failure(c, "保存失败: "+err.Error())
		return
	}

	response.SuccessMessage(c, "保存成功")
}

// UpdateArticle 更新文章
// @Summary 更新文章信息
// @Description 管理员更新文章内容和设置
func UpdateArticle(c *gin.Context) {
	var req ArticleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 验证文章标题
	if req.ArticleTitle != nil {
		if err := utils.ValidateArticleTitle(*req.ArticleTitle); err != nil {
			response.Failure(c, err.Error())
			return
		}
	}

	// 清理文章内容（防XSS）
	if req.ArticleContent != nil {
		cleaned := utils.SanitizeHTML(*req.ArticleContent)
		req.ArticleContent = &cleaned
	}

	// 构建更新字段
	updates := map[string]interface{}{
		"summary": "", // 更新文章时清空摘要，重新生成
	}

	if req.ArticleTitle != nil {
		updates["article_title"] = *req.ArticleTitle
	}
	if req.ArticleAuthor != nil {
		updates["update_by"] = *req.ArticleAuthor
	}
	if req.ArticleContent != nil {
		updates["article_content"] = *req.ArticleContent
	}
	if req.RecommendStatus != nil {
		updates["recommend_status"] = *req.RecommendStatus
	}
	if req.ViewStatus != nil {
		updates["view_status"] = *req.ViewStatus
	}
	if req.CommentStatus != nil {
		updates["comment_status"] = *req.CommentStatus
	}
	if req.Password != nil {
		updates["password"] = *req.Password
	}
	if req.ArticleCover != nil {
		updates["article_cover"] = *req.ArticleCover
	}
	if req.LabelID != nil {
		updates["label_id"] = *req.LabelID
	}
	if req.SortID != nil {
		updates["sort_id"] = *req.SortID
	}

	if err := config.DB.Model(&models.Article{}).Where("id = ?", req.ID).Updates(updates).Error; err != nil {
		response.Failure(c, "更新失败: "+err.Error())
		return
	}

	response.SuccessMessage(c, "更新成功")
}

// DeleteArticle 删除文章
// @Summary 删除文章
// @Description 管理员删除指定文章
func DeleteArticle(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if err := config.DB.Delete(&models.Article{}, req.ID).Error; err != nil {
		response.Failure(c, "删除失败: "+err.Error())
		return
	}

	response.SuccessMessage(c, "删除成功")
}

// ChangeArticleStatus 修改文章状态
// @Summary 修改文章状态
// @Description 批量修改文章的可见、推荐、评论状态
func ChangeArticleStatus(c *gin.Context) {
	var req ArticleStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	updates := make(map[string]interface{})

	if req.ViewStatus != nil {
		updates["view_status"] = *req.ViewStatus
	}
	if req.RecommendStatus != nil {
		updates["recommend_status"] = *req.RecommendStatus
	}
	if req.CommentStatus != nil {
		updates["comment_status"] = *req.CommentStatus
	}

	if len(updates) == 0 {
		response.Failure(c, "未指定要更新的状态")
		return
	}

	if err := config.DB.Model(&models.Article{}).Where("id = ?", req.ArticleID).Updates(updates).Error; err != nil {
		response.Failure(c, "更新失败: "+err.Error())
		return
	}

	response.SuccessMessage(c, "更新成功")
}

// ArticleLike 文章点赞/取消点赞
// @Summary 文章点赞
// @Description 用户对文章进行点赞或取消点赞操作
func ArticleLike(c *gin.Context) {
	var req ArticleLikeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 获取当前登录用户
	currentClient, exists := middleware.GetCurrentClient(c)
	if !exists {
		response.Unauthorized(c, "请先登录")
		return
	}

	articleID := req.ArticleID
	userID := currentClient.UserID

	// 使用事务保证点赞/取消点赞的原子性
	var isCancel bool
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var like models.ArticleLike
		if err := tx.Where("user_id_id = ? AND article_id_id = ?", userID, articleID).First(&like).Error; err == nil {
			// 已点赞，取消点赞
			if err := tx.Delete(&like).Error; err != nil {
				return fmt.Errorf("取消点赞失败: %w", err)
			}
			if err := tx.Model(&models.Article{}).Where("id = ?", articleID).
				UpdateColumn("like_count", tx.Raw("GREATEST(like_count - 1, 0)")).Error; err != nil {
				return fmt.Errorf("更新点赞数失败: %w", err)
			}
			isCancel = true
			return nil
		}

		// 新增点赞
		like = models.ArticleLike{
			UserID:    userID,
			ArticleID: articleID,
		}
		if err := tx.Create(&like).Error; err != nil {
			return fmt.Errorf("点赞失败: %w", err)
		}
		if err := tx.Model(&models.Article{}).Where("id = ?", articleID).
			UpdateColumn("like_count", tx.Raw("like_count + 1")).Error; err != nil {
			return fmt.Errorf("更新点赞数失败: %w", err)
		}
		return nil
	})

	if err != nil {
		response.Failure(c, err.Error())
		return
	}

	if isCancel {
		response.SuccessMessage(c, "取消点赞成功")
	} else {
		response.SuccessMessage(c, "点赞成功")
	}
}

// GetArticleSummary 获取文章摘要（SSE 流式输出）
// @Summary 获取文章AI摘要
// @Description 调用 DeepSeek API 流式生成文章摘要，已有摘要直接返回（非流式）
func GetArticleSummary(c *gin.Context) {
	var req ArticleSummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if req.Message == "" {
		response.Failure(c, "文章内容不能为空")
		return
	}

	// 检查是否已有摘要（直接返回 JSON，非流式）
	if req.ArticleID > 0 {
		var article models.Article
		if err := config.DB.First(&article, req.ArticleID).Error; err == nil && article.Summary != "" {
			response.Success(c, gin.H{"summary": article.Summary})
			return
		}
	}

	// 检查 AI API 配置
	cfg := config.AppConfig
	if cfg.AIAPIKey == "" || cfg.AIAPIKey == "your-api-key-here" {
		response.Failure(c, "AI摘要服务未配置")
		return
	}

	// 截取文章内容（避免超长，摘要不需要全文）
	content := req.Message
	if len(content) > 8000 {
		content = content[:8000]
	}

	// 构建 DeepSeek 兼容格式请求（OpenAI / 硅基流动等通用）
	apiBody := map[string]interface{}{
		"model":      cfg.AIModel,
		"max_tokens": 512,
		"stream":     true,
		"messages": []map[string]string{
			{"role": "system", "content": "请简洁概括以下文章的核心内容，要求语言精炼、通俗易懂。直接输出摘要内容，不要加任何前缀。"},
			{"role": "user", "content": content},
		},
	}
	apiBodyJSON, _ := json.Marshal(apiBody)

	apiURL := strings.TrimRight(cfg.AIAPIURL, "/") + "/chat/completions"
	httpReq, err := http.NewRequest("POST", apiURL, strings.NewReader(string(apiBodyJSON)))
	if err != nil {
		response.Failure(c, "创建请求失败")
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+cfg.AIAPIKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := SharedHTTPClient.Do(httpReq)
	if err != nil {
		response.Failure(c, "AI服务请求失败: "+err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("[Summary] AI API 错误: status=%d, body=%s", resp.StatusCode, string(body))
		response.Failure(c, "AI服务返回错误")
		return
	}

	// 设置 SSE 响应头
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("X-Accel-Buffering", "no") // 禁用 nginx 缓冲
	c.Writer.WriteHeader(http.StatusOK)

	// 逐行读取 DeepSeek SSE 响应并转发给前端
	var fullSummary strings.Builder
	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 0, 64*1024), 64*1024)

	flusher, canFlush := c.Writer.(http.Flusher)

	for scanner.Scan() {
		line := scanner.Text()

		// DeepSeek SSE 格式: "data: {...}" 或 "data: [DONE]"
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		data := strings.TrimPrefix(line, "data: ")

		// DeepSeek 流结束标记
		if data == "[DONE]" {
			break
		}

		// 解析 JSON：DeepSeek 格式 choices[0].delta.content
		var event struct {
			Choices []struct {
				Delta struct {
					Content string `json:"content"`
				} `json:"delta"`
			} `json:"choices"`
		}
		if err := json.Unmarshal([]byte(data), &event); err != nil {
			continue
		}

		if len(event.Choices) > 0 {
			text := event.Choices[0].Delta.Content
			if text != "" {
				fullSummary.WriteString(text)
				// 推送给前端（统一格式）
				chunk, _ := json.Marshal(map[string]string{"text": text})
				fmt.Fprintf(c.Writer, "data: %s\n\n", chunk)
				if canFlush {
					flusher.Flush()
				}
			}
		}
	}

	// 保存完整摘要到数据库
	summaryText := fullSummary.String()
	if req.ArticleID > 0 && summaryText != "" {
		config.DB.Model(&models.Article{}).Where("id = ?", req.ArticleID).Update("summary", summaryText)
	}

	// 发送结束信号
	fmt.Fprintf(c.Writer, "data: [DONE]\n\n")
	if canFlush {
		flusher.Flush()
	}
}

// ==================== 辅助函数 ====================

// buildArticleList 构建文章列表数据
func buildArticleList(articles []models.Article, checkCover bool) []map[string]interface{} {
	if len(articles) == 0 {
		return []map[string]interface{}{}
	}

	// 批量收集所有需要查询的用户ID
	userIDs := make([]uint, 0, len(articles))
	for _, a := range articles {
		userIDs = append(userIDs, a.UserID)
	}
	usernameMap := batchGetUsernames(userIDs)

	data := make([]map[string]interface{}, 0, len(articles))
	for _, article := range articles {
		label := getLabel(article.LabelID, article.ID)
		sort := getSort(article.SortID)
		username := usernameMap[article.UserID]
		commentCount := getCommentCount(article.ID)

		cover := article.ArticleCover
		if checkCover {
			coverStr := getValidCover(article.ArticleCover)
			cover = &coverStr
		}

		data = append(data, map[string]interface{}{
			"id":              article.ID,
			"userId":          article.UserID,
			"username":        username,
			"sortId":          article.SortID,
			"labelId":         article.LabelID,
			"articleCover":    cover,
			"articleTitle":    article.ArticleTitle,
			"articleContent":  article.ArticleContent,
			"viewCount":       article.ViewCount,
			"likeCount":       article.LikeCount,
			"viewStatus":      article.ViewStatus,
			"password":        article.Password,
			"recommendStatus": article.RecommendStatus,
			"commentStatus":   article.CommentStatus,
			"createTime":      article.CreateTime,
			"updateTime":      article.UpdateTime,
			"updateBy":        article.UpdateBy,
			"commentCount":    commentCount,
			"hasSummary":      article.Summary != "",
			"label":           label,
			"sort":            sort,
		})
	}
	return data
}

// buildArticleDetail 构建文章详情数据
func buildArticleDetail(article models.Article, likeStatus int) map[string]interface{} {
	label := getLabel(article.LabelID, article.ID)
	sort := getSort(article.SortID)
	username := getUsername(article.UserID)
	commentCount := getCommentCount(article.ID)
	cover := getValidCover(article.ArticleCover)

	return map[string]interface{}{
		"id":                article.ID,
		"summary":           article.Summary,
		"userId":            article.UserID,
		"sortId":            article.SortID,
		"labelId":           article.LabelID,
		"articleCover":      cover,
		"articleTitle":      article.ArticleTitle,
		"articleContent":    article.ArticleContent,
		"articleAuthor":     article.UpdateBy,
		"viewStatus":        article.ViewStatus,
		"password":          article.Password,
		"recommendStatus":   article.RecommendStatus,
		"commentStatus":     article.CommentStatus,
		"createTime":        article.CreateTime,
		"likeCount":         article.LikeCount,
		"updateBy":          article.UpdateBy,
		"updateTime":        article.UpdateTime,
		"username":          username,
		"commentCount":      commentCount,
		"viewCount":         article.ViewCount,
		"label":             label,
		"sort":              sort,
		"articleLikeStatus": likeStatus,
	}
}
