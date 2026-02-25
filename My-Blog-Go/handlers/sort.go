// Package handlers 分类标签相关接口处理
package handlers

import (
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// SaveOrUpdateSortRequest 保存或更新分类请求
type SaveOrUpdateSortRequest struct {
	ID              *uint  `json:"id"`                                 // 分类ID（有ID则更新，无则新增）
	SortName        string `json:"sortName" binding:"required"`        // 分类名称
	SortDescription string `json:"sortDescription" binding:"required"` // 分类描述
	SortType        int    `json:"sortType"`                           // 分类类型: 0-导航栏, 1-普通
	Priority        *int   `json:"priority" binding:"required"`        // 优先级
}

// SaveOrUpdateLabelRequest 保存或更新标签请求
type SaveOrUpdateLabelRequest struct {
	ID               *uint  `json:"id"`                                  // 标签ID（有ID则更新，无则新增）
	SortID           uint   `json:"sortId" binding:"required"`           // 所属分类ID
	LabelName        string `json:"labelName" binding:"required"`        // 标签名称
	LabelDescription string `json:"labelDescription" binding:"required"` // 标签描述
}

// ==================== 接口处理函数 ====================

// GetSortInfo 获取分类信息
// @Summary 获取分类信息（公开）
// @Description 获取所有分类及其下属标签信息
func GetSortInfo(c *gin.Context) {
	var sorts []models.Sort
	config.DB.Order("id").Find(&sorts)

	// 一次性获取所有标签
	var allLabels []models.Label
	config.DB.Order("id").Find(&allLabels)

	// 一次性批量统计每个标签的文章数
	type LabelCount struct {
		LabelID uint  `gorm:"column:label_id"`
		Count   int64 `gorm:"column:count"`
	}
	var labelCounts []LabelCount
	config.DB.Model(&models.Article{}).Select("label_id, COUNT(*) as count").Group("label_id").Find(&labelCounts)
	labelCountMap := make(map[uint]int64, len(labelCounts))
	for _, lc := range labelCounts {
		labelCountMap[lc.LabelID] = lc.Count
	}

	// 一次性批量统计每个分类的文章数
	type SortCount struct {
		SortID uint  `gorm:"column:sort_id"`
		Count  int64 `gorm:"column:count"`
	}
	var sortCounts []SortCount
	config.DB.Model(&models.Article{}).Select("sort_id, COUNT(*) as count").Group("sort_id").Find(&sortCounts)
	sortCountMap := make(map[uint]int64, len(sortCounts))
	for _, sc := range sortCounts {
		sortCountMap[sc.SortID] = sc.Count
	}

	// 按 sort_id 分组标签
	labelsBySortID := make(map[uint][]models.Label)
	for _, label := range allLabels {
		labelsBySortID[label.SortID] = append(labelsBySortID[label.SortID], label)
	}

	// 构建响应
	data := make([]map[string]interface{}, 0, len(sorts))
	for _, sort := range sorts {
		labels := labelsBySortID[sort.ID]
		var labelList []map[string]interface{}
		for _, label := range labels {
			labelList = append(labelList, map[string]interface{}{
				"id":               label.ID,
				"labelName":        label.LabelName,
				"labelDescription": label.LabelDescription,
				"sortId":           label.SortID,
				"countOfLabel":     labelCountMap[label.ID],
			})
		}

		data = append(data, map[string]interface{}{
			"id":              sort.ID,
			"sortName":        sort.SortName,
			"sortDescription": sort.SortDescription,
			"sortType":        sort.SortType,
			"priority":        sort.Priority,
			"countOfSort":     sortCountMap[sort.ID],
			"lengthOfLabel":   len(labels),
			"labels":          labelList,
		})
	}

	response.Success(c, data)
}

// SaveOrUpdateSort 保存或更新分类
// @Summary 保存或更新分类
// @Description 管理员添加或更新分类，有ID则更新，无ID则新增
func SaveOrUpdateSort(c *gin.Context) {
	var req SaveOrUpdateSortRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 验证分类类型
	if req.SortType != 0 && req.SortType != 1 {
		response.Failure(c, "分类类型无效")
		return
	}

	if req.ID != nil && *req.ID > 0 {
		// 更新模式
		updates := map[string]interface{}{
			"sort_name":        req.SortName,
			"sort_description": req.SortDescription,
			"sort_type":        req.SortType,
			"priority":         req.Priority,
		}
		if err := config.DB.Model(&models.Sort{}).Where("id = ?", *req.ID).Updates(updates).Error; err != nil {
			response.Failure(c, "更新失败")
			return
		}
		response.SuccessMessage(c, "更新成功")
	} else {
		// 新增模式
		sort := models.Sort{
			SortName:        req.SortName,
			SortDescription: req.SortDescription,
			SortType:        req.SortType,
			Priority:        req.Priority,
		}
		if err := config.DB.Create(&sort).Error; err != nil {
			response.Failure(c, "保存失败")
			return
		}
		response.SuccessMessage(c, "保存成功")
	}
}

// DeleteSort 删除分类
// @Summary 删除分类
// @Description 管理员删除分类，关联的文章和标签将移至默认分类
func DeleteSort(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 将相关文章和标签的 sort_id 更新为默认值（从配置读取）
	defaultSortID := config.AppConfig.DefaultSortID
	if err := config.DB.Model(&models.Article{}).Where("sort_id = ?", req.ID).Update("sort_id", defaultSortID).Error; err != nil {
		response.Failure(c, "迁移关联文章失败: "+err.Error())
		return
	}
	if err := config.DB.Model(&models.Label{}).Where("sort_id = ?", req.ID).Update("sort_id", defaultSortID).Error; err != nil {
		response.Failure(c, "迁移关联标签失败: "+err.Error())
		return
	}

	if err := config.DB.Delete(&models.Sort{}, req.ID).Error; err != nil {
		response.Failure(c, "删除失败")
		return
	}

	response.SuccessMessage(c, "删除成功")
}

// SaveOrUpdateLabel 保存或更新标签
// @Summary 保存或更新标签
// @Description 管理员添加或更新标签，有ID则更新，无ID则新增
func SaveOrUpdateLabel(c *gin.Context) {
	var req SaveOrUpdateLabelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	if req.ID != nil && *req.ID > 0 {
		// 更新模式
		updates := map[string]interface{}{
			"sort_id":           req.SortID,
			"label_name":        req.LabelName,
			"label_description": req.LabelDescription,
		}
		if err := config.DB.Model(&models.Label{}).Where("id = ?", *req.ID).Updates(updates).Error; err != nil {
			response.Failure(c, "更新失败")
			return
		}
		response.SuccessMessage(c, "更新成功")
	} else {
		// 新增模式
		label := models.Label{
			SortID:           req.SortID,
			LabelName:        req.LabelName,
			LabelDescription: req.LabelDescription,
		}
		if err := config.DB.Create(&label).Error; err != nil {
			response.Failure(c, "保存失败")
			return
		}
		response.SuccessMessage(c, "保存成功")
	}
}

// DeleteLabel 删除标签
// @Summary 删除标签
// @Description 管理员删除标签，关联的文章将移至默认标签
func DeleteLabel(c *gin.Context) {
	var req IDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	// 将相关文章的 label_id 更新为默认值（从配置读取）
	defaultLabelID := config.AppConfig.DefaultLabelID
	if err := config.DB.Model(&models.Article{}).Where("label_id = ?", req.ID).Update("label_id", defaultLabelID).Error; err != nil {
		response.Failure(c, "迁移关联文章失败: "+err.Error())
		return
	}

	if err := config.DB.Delete(&models.Label{}, req.ID).Error; err != nil {
		response.Failure(c, "删除失败")
		return
	}

	response.SuccessMessage(c, "删除成功")
}

// ListSortAndLabel 获取所有分类和标签
// @Summary 获取分类和标签列表
// @Description 管理员获取所有分类和标签的列表
func ListSortAndLabel(c *gin.Context) {
	// 获取所有分类
	var sorts []models.Sort
	config.DB.Order("id").Find(&sorts)

	// 获取所有标签
	var labels []models.Label
	config.DB.Order("id").Find(&labels)

	// 批量统计分类文章数
	type CountResult struct {
		ID    uint  `gorm:"column:id"`
		Count int64 `gorm:"column:count"`
	}
	var sortCounts []CountResult
	config.DB.Model(&models.Article{}).Select("sort_id as id, COUNT(*) as count").Group("sort_id").Find(&sortCounts)
	sortCountMap := make(map[uint]int64, len(sortCounts))
	for _, sc := range sortCounts {
		sortCountMap[sc.ID] = sc.Count
	}

	// 批量统计标签文章数
	var labelCounts []CountResult
	config.DB.Model(&models.Article{}).Select("label_id as id, COUNT(*) as count").Group("label_id").Find(&labelCounts)
	labelCountMap := make(map[uint]int64, len(labelCounts))
	for _, lc := range labelCounts {
		labelCountMap[lc.ID] = lc.Count
	}

	var sortList []map[string]interface{}
	for _, sort := range sorts {
		sortList = append(sortList, map[string]interface{}{
			"countOfSort":     sortCountMap[sort.ID],
			"id":              sort.ID,
			"sortName":        sort.SortName,
			"sortType":        sort.SortType,
			"sortDescription": sort.SortDescription,
			"priority":        sort.Priority,
		})
	}

	var labelList []map[string]interface{}
	for _, label := range labels {
		labelList = append(labelList, map[string]interface{}{
			"countOfLabel":     labelCountMap[label.ID],
			"id":               label.ID,
			"labelDescription": label.LabelDescription,
			"labelName":        label.LabelName,
			"sortId":           label.SortID,
		})
	}

	response.Success(c, map[string]interface{}{
		"labels": labelList,
		"sorts":  sortList,
	})
}
