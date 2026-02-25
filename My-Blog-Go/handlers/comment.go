// Package handlers 评论相关接口处理
package handlers

import (
	"log"
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"
	"my-blog-go/utils"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// CommentListRequest 评论列表请求（管理员）
type CommentListRequest struct {
	Current     int    `json:"current"`     // 当前页码
	Size        int    `json:"size"`        // 每页条数
	Source      int    `json:"source"`      // 来源ID（文章ID等）
	CommentType string `json:"commentType"` // 评论类型
}

// ClientCommentListRequest 评论列表请求（前台）
type ClientCommentListRequest struct {
	Current        int    `json:"current"`        // 当前页码
	Size           int    `json:"size"`           // 每页条数
	CSize          int    `json:"csize"`          // 子评论每页条数
	Source         int    `json:"source"`         // 来源ID
	CommentType    string `json:"commentType"`    // 评论类型
	FloorCommentID *uint  `json:"floorCommentId"` // 楼层评论ID
}

// AddCommentRequest 添加评论请求
type AddCommentRequest struct {
	Source          int     `json:"source" binding:"required"`         // 来源ID
	Type            string  `json:"type"`                              // 评论类型
	ParentCommentID *uint   `json:"parentCommentId"`                   // 父评论ID
	FloorCommentID  *uint   `json:"floorCommentId"`                    // 楼层评论ID
	ParentUserID    *uint   `json:"parentUserId"`                      // 被回复用户ID
	CommentContent  string  `json:"commentContent" binding:"required"` // 评论内容
	CommentInfo     *string `json:"commentInfo"`                       // 附加信息
}

// ==================== 接口处理函数 ====================

// AdminCommentList 管理员获取评论列表
// @Summary 获取评论列表（管理员）
// @Description 管理员获取所有评论列表，支持筛选
func AdminCommentList(c *gin.Context) {
	var req CommentListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	query := config.DB.Model(&models.Comment{})

	// 应用筛选条件
	if req.Source > 0 {
		query = query.Where("source = ?", req.Source)
	}
	if req.CommentType != "" {
		query = query.Where("type = ?", req.CommentType)
	}

	var total int64
	query.Count(&total)

	var comments []models.Comment
	current, size, offset := normalizePagination(req.Current, req.Size)
	query.Order("create_time DESC").Offset(offset).Limit(size).Find(&comments)

	// 构建响应数据
	data := make([]map[string]interface{}, 0, len(comments))
	for _, comment := range comments {
		client, _ := getUserClient(comment.UserID)

		data = append(data, map[string]interface{}{
			"id":              comment.ID,
			"source":          comment.Source,
			"type":            comment.Type,
			"parentCommentId": comment.ParentCommentID,
			"username":        client.Username,
			"avatar":          client.Avatar,
			"floorCommentId":  comment.FloorCommentID,
			"parentUserId":    comment.ParentUserID,
			"likeCount":       comment.LikeCount,
			"commentContent":  comment.CommentContent,
			"commentInfo":     comment.CommentInfo,
			"createTime":      comment.CreateTime,
		})
	}

	response.SuccessPage(c, data, total, current, size)
}

// ClientCommentList 前台获取评论列表
// @Summary 获取评论列表（前台）
// @Description 获取嵌套结构的评论列表，支持楼层展开
func ClientCommentList(c *gin.Context) {
	var req ClientCommentListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 规范化分页参数（此接口使用自定义响应，不需要 current）
	_, size, offset := normalizePagination(req.Current, req.Size)

	query := config.DB.Model(&models.Comment{})
	if req.Source > 0 {
		query = query.Where("source = ?", req.Source)
	}
	if req.CommentType != "" {
		query = query.Where("type = ?", req.CommentType)
	}

	var total, parentTotal int64
	var comments []models.Comment

	// 根据是否指定楼层ID决定查询方式
	if req.FloorCommentID != nil && *req.FloorCommentID > 0 {
		// 查询指定楼层的子评论
		query = query.Where("floor_comment_id = ?", *req.FloorCommentID)
		query.Count(&total)
		query.Order("id ASC").Offset(offset).Limit(size).Find(&comments)
	} else {
		// 查询顶级评论（floor_comment_id = 0）
		floorQuery := query.Where("floor_comment_id = 0")
		floorQuery.Count(&parentTotal)
		floorQuery.Order("id ASC").Offset(offset).Limit(size).Find(&comments)
		// 总评论数
		config.DB.Model(&models.Comment{}).Where("source = ?", req.Source).Count(&total)
	}

	// 构建响应数据
	data := make([]map[string]interface{}, 0, len(comments))
	for _, comment := range comments {
		commentData := buildCommentData(comment, req.FloorCommentID == nil, req.CSize)
		data = append(data, commentData)
	}

	// 构建结果
	result := gin.H{
		"total":       total,
		"parenttotal": total - parentTotal,
		"data":        data,
	}
	if req.FloorCommentID != nil && *req.FloorCommentID > 0 {
		result["total"] = total + 1
	}

	response.Success(c, result)
}

// AddComment 添加评论
// @Summary 添加评论
// @Description 用户添加评论或回复
func AddComment(c *gin.Context) {
	var req AddCommentRequest
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

	// 验证评论内容
	if err := utils.ValidateComment(req.CommentContent); err != nil {
		response.Failure(c, err.Error())
		return
	}

	// 清理评论内容（防XSS）
	req.CommentContent = utils.SanitizeString(req.CommentContent)

	// 处理楼层评论ID交换逻辑（兼容前端逻辑）
	var comment models.Comment
	if req.FloorCommentID != nil && req.ParentCommentID != nil &&
		*req.FloorCommentID > 0 && *req.ParentCommentID > 0 {
		// 两个ID都存在时交换存储
		comment = models.Comment{
			Source:          req.Source,
			Type:            req.Type,
			ParentCommentID: req.FloorCommentID, // 交换
			UserID:          currentUser.UserID, // 直接使用当前登录用户ID
			FloorCommentID:  req.ParentCommentID, // 交换
			ParentUserID:    req.ParentUserID,
			CommentContent:  req.CommentContent,
			CommentInfo:     req.CommentInfo,
		}
	} else {
		// 正常存储
		comment = models.Comment{
			Source:          req.Source,
			Type:            req.Type,
			ParentCommentID: req.ParentCommentID,
			UserID:          currentUser.UserID, // 直接使用当前登录用户ID
			FloorCommentID:  req.FloorCommentID,
			ParentUserID:    req.ParentUserID,
			CommentContent:  req.CommentContent,
			CommentInfo:     req.CommentInfo,
		}
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		response.Failure(c, "评论失败")
		return
	}

	// 自动发送评论通知邮件（异步，不阻塞响应）
	// 1. 通知被回复的用户
	if req.ParentUserID != nil && *req.ParentUserID > 0 && *req.ParentUserID != currentUser.UserID {
		go sendCommentNotification(*req.ParentUserID, currentUser.Username, req.CommentContent)
	}
	// 2. 通知博主（Boss用户），如果评论者不是博主自己
	if currentUser.UserType != models.UserTypeBoss {
		go notifyBossUser(currentUser.Username, req.CommentContent, "comment")
	}

	response.SuccessMessage(c, "评论成功")
}

// sendCommentNotification 异步发送评论通知邮件
func sendCommentNotification(parentUserID uint, commenterName, commentContent string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[NOTIFY] 发送评论通知异常: %v", r)
		}
	}()

	parentClient, err := getUserClient(parentUserID)
	if err != nil || parentClient.Email == nil || *parentClient.Email == "" {
		return // 用户不存在或未设置邮箱，静默跳过
	}

	if err := utils.SendCommentNotification(*parentClient.Email, commentContent, commenterName); err != nil {
		log.Printf("[NOTIFY] 发送评论通知失败: userId=%d, err=%v", parentUserID, err)
	}
}

// notifyBossUser 异步通知博主（Boss 用户）
func notifyBossUser(senderName, content, notifyType string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[NOTIFY] 通知博主异常: %v", r)
		}
	}()

	// 查找 Boss 用户（userType=0）
	var bossClient models.Client
	if err := config.DB.Where("user_type = ?", models.UserTypeBoss).First(&bossClient).Error; err != nil {
		return // 没有找到 Boss 用户
	}
	if bossClient.Email == nil || *bossClient.Email == "" {
		return // Boss 未设置邮箱
	}

	if err := utils.SendCommentNotification(*bossClient.Email, content, senderName); err != nil {
		log.Printf("[NOTIFY] 通知博主失败: err=%v", err)
	}
}

// DeleteComment 删除评论
// @Summary 删除评论
// @Description 管理员删除指定评论
func DeleteComment(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if err := config.DB.Delete(&models.Comment{}, req.ID).Error; err != nil {
		response.Failure(c, "删除失败")
		return
	}

	response.SuccessMessage(c, "删除成功")
}

// ==================== 辅助函数 ====================

// collectUserAndParentIDs 从评论列表中收集所有涉及的用户ID和父评论ID
func collectUserAndParentIDs(comments []models.Comment) (userIDs []uint, parentCommentIDs []uint) {
	userIDSet := make(map[uint]bool)
	parentIDSet := make(map[uint]bool)
	for _, c := range comments {
		userIDSet[c.UserID] = true
		if c.ParentCommentID != nil && *c.ParentCommentID > 0 {
			parentIDSet[*c.ParentCommentID] = true
		}
	}
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}
	for id := range parentIDSet {
		parentCommentIDs = append(parentCommentIDs, id)
	}
	return
}

// batchLoadClients 批量加载用户信息，返回 userID -> Client 的映射
func batchLoadClients(userIDs []uint) map[uint]*models.Client {
	if len(userIDs) == 0 {
		return make(map[uint]*models.Client)
	}
	var clients []models.Client
	config.DB.Where("user_id IN ?", userIDs).Find(&clients)
	result := make(map[uint]*models.Client, len(clients))
	for i := range clients {
		result[clients[i].UserID] = &clients[i]
	}
	return result
}

// batchLoadComments 批量加载评论，返回 commentID -> Comment 的映射
func batchLoadComments(commentIDs []uint) map[uint]*models.Comment {
	if len(commentIDs) == 0 {
		return make(map[uint]*models.Comment)
	}
	var comments []models.Comment
	config.DB.Where("id IN ?", commentIDs).Find(&comments)
	result := make(map[uint]*models.Comment, len(comments))
	for i := range comments {
		result[comments[i].ID] = &comments[i]
	}
	return result
}

// buildCommentData 构建评论数据（使用预加载的缓存避免 N+1 查询）
func buildCommentData(comment models.Comment, includeChildren bool, childSize int) map[string]interface{} {
	// 收集当前评论 + 子评论的所有ID
	allComments := []models.Comment{comment}
	var childComments []models.Comment

	if includeChildren {
		childQuery := config.DB.Where("floor_comment_id = ?", comment.ID)
		if childSize > 0 {
			childQuery = childQuery.Limit(childSize)
		}
		childQuery.Order("id ASC").Find(&childComments)
		allComments = append(allComments, childComments...)
	}

	// 批量加载所有涉及的用户和父评论
	userIDs, parentCommentIDs := collectUserAndParentIDs(allComments)

	// 父评论的用户也需要加载
	parentCommentMap := batchLoadComments(parentCommentIDs)
	for _, pc := range parentCommentMap {
		userIDs = append(userIDs, pc.UserID)
	}

	clientMap := batchLoadClients(userIDs)

	// 构建当前评论数据
	data := buildSingleCommentData(comment, clientMap, parentCommentMap)

	// 构建子评论
	if includeChildren {
		var childData []map[string]interface{}
		for _, child := range childComments {
			childData = append(childData, buildSingleCommentData(child, clientMap, parentCommentMap))
		}

		var childTotal int64
		config.DB.Model(&models.Comment{}).Where("floor_comment_id = ?", comment.ID).Count(&childTotal)

		data["childComments"] = []map[string]interface{}{{
			"current": 1,
			"records": childData,
			"total":   childTotal,
		}}
	}

	return data
}

// buildSingleCommentData 使用预加载的缓存构建单条评论数据
func buildSingleCommentData(comment models.Comment, clientMap map[uint]*models.Client, parentCommentMap map[uint]*models.Comment) map[string]interface{} {
	client := clientMap[comment.UserID]
	username := ""
	avatar := ""
	userType := 2
	if client != nil {
		username = client.Username
		avatar = getValidAvatar(client.Avatar)
		userType = client.UserType
	}

	// 获取父评论用户名
	var parentUsername string
	if comment.ParentCommentID != nil {
		if parentComment, ok := parentCommentMap[*comment.ParentCommentID]; ok {
			if parentClient, ok2 := clientMap[parentComment.UserID]; ok2 {
				parentUsername = parentClient.Username
			}
		}
	}

	data := map[string]interface{}{
		"id":              comment.ID,
		"source":          comment.Source,
		"type":            comment.Type,
		"parentCommentId": comment.ParentCommentID,
		"userId":          comment.UserID,
		"username":        username,
		"avatar":          avatar,
		"userType":        userType,
		"floorCommentId":  comment.FloorCommentID,
		"parentUserId":    comment.ParentUserID,
		"likeCount":       comment.LikeCount,
		"commentContent":  comment.CommentContent,
		"commentInfo":     comment.CommentInfo,
		"createTime":      comment.CreateTime,
	}

	if parentUsername != "" {
		data["parentUsername"] = parentUsername
	}

	return data
}
