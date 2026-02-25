// Package handlers IP相关接口处理
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"my-blog-go/config"
	"my-blog-go/models"
	"my-blog-go/response"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ==================== 请求结构体 ====================

// RecordIPRequest 记录IP请求
type RecordIPRequest struct {
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
}

// ProvinceData 省份数据结构
type ProvinceData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ==================== 接口处理函数 ====================

// GetIP 获取客户端IP
// @Summary 获取客户端IP
// @Description 返回请求客户端的IP地址
func GetIP(c *gin.Context) {
	ip := GetClientIP(c)
	response.Success(c, gin.H{"ip": ip})
}

// GetIPLocation 获取IP定位和天气信息
// @Summary 获取IP定位和天气
// @Description 根据客户端IP获取地理位置和天气信息
func GetIPLocation(c *gin.Context) {
	// 统一使用 apihz.cn API 获取位置和天气
	result := getLocationAndWeatherFromApihz()
	response.Success(c, result)
}

// getLocationAndWeatherFromApihz 从 apihz.cn 一次性获取位置和天气
func getLocationAndWeatherFromApihz() gin.H {
	result := gin.H{
		"ip":       "",
		"city":     "",
		"province": "",
		"weather":  []map[string]interface{}{},
		"tip":      "",
	}

	// day=4 获取4天天气预报（今天、明天、后天、大后天）
	url := fmt.Sprintf("https://cn.apihz.cn/api/tianqi/tqybip.php?id=%s&key=%s&day=4",
		config.AppConfig.WeatherAPIID, config.AppConfig.WeatherAPIKey)

	resp, err := SharedHTTPClient.Get(url)
	if err != nil {
		log.Printf("[Weather] apihz.cn 请求失败: %v", err)
		return result
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return result
	}

	// 检查返回码
	if code, ok := data["code"].(float64); !ok || code != 200 {
		return result
	}

	// 解析位置信息
	// sheng: 省份, shi/name: 城市
	if sheng, ok := data["sheng"].(string); ok {
		result["province"] = sheng
	}
	if shi, ok := data["shi"].(string); ok {
		result["city"] = shi
	}

	// 解析天气信息
	weatherList := make([]map[string]interface{}, 0)
	weekdays := []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}

	// 今天的天气（白天）
	todayWeather := map[string]interface{}{
		"date":      time.Now().Format("2006-01-02"),
		"week":      weekdays[int(time.Now().Weekday())],
		"type":      "",
		"high":      "",
		"low":       "",
		"fengli":    "",
		"fengxiang": "",
	}

	// weather1: 白天天气, weather2: 夜间天气
	if weather1, ok := data["weather1"].(string); ok {
		todayWeather["type"] = weather1
	}
	// wd1: 最高温度, wd2: 最低温度
	if wd1, ok := data["wd1"].(string); ok {
		todayWeather["high"] = wd1 + "°C"
	}
	if wd2, ok := data["wd2"].(string); ok {
		todayWeather["low"] = wd2 + "°C"
	}
	// winddirection1: 风向, windleve1: 风力
	if windDir, ok := data["winddirection1"].(string); ok {
		todayWeather["fengxiang"] = windDir
	}
	if windLevel, ok := data["windleve1"].(string); ok {
		todayWeather["fengli"] = windLevel
	}

	weatherList = append(weatherList, todayWeather)

	// 解析未来天气（weatherday2, weatherday3...）
	for i := 2; i <= 7; i++ {
		dayKey := fmt.Sprintf("weatherday%d", i)
		dayData, ok := data[dayKey].(map[string]interface{})
		if !ok {
			continue
		}

		futureDate := time.Now().AddDate(0, 0, i-1) // weatherday2 是明天，所以 i-1
		futureWeather := map[string]interface{}{
			"date":      futureDate.Format("2006-01-02"),
			"week":      weekdays[int(futureDate.Weekday())],
			"type":      "",
			"high":      "",
			"low":       "",
			"fengli":    "",
			"fengxiang": "",
		}

		// 解析天气类型
		if weather, ok := dayData["weather1"].(string); ok {
			futureWeather["type"] = weather
		}
		// 解析温度（可能是 string 或 float64）
		if wd1, ok := dayData["wd1"].(float64); ok {
			futureWeather["high"] = fmt.Sprintf("%.0f°C", wd1)
		} else if wd1, ok := dayData["wd1"].(string); ok {
			futureWeather["high"] = wd1 + "°C"
		}
		if wd2, ok := dayData["wd2"].(float64); ok {
			futureWeather["low"] = fmt.Sprintf("%.0f°C", wd2)
		} else if wd2, ok := dayData["wd2"].(string); ok {
			futureWeather["low"] = wd2 + "°C"
		}
		// 解析风向风力
		if windDir, ok := dayData["winddirection1"].(string); ok {
			futureWeather["fengxiang"] = windDir
		}
		if windLevel, ok := dayData["windleve1"].(string); ok {
			futureWeather["fengli"] = windLevel
		}

		weatherList = append(weatherList, futureWeather)
	}

	result["weather"] = weatherList

	// 生成 tip（使用实时数据）
	if nowinfo, ok := data["nowinfo"].(map[string]interface{}); ok {
		temp := ""
		if t, ok := nowinfo["temperature"].(float64); ok {
			temp = fmt.Sprintf("%.1f°C", t)
		}
		windDir := ""
		if wd, ok := nowinfo["windDirection"].(string); ok {
			windDir = wd
		}
		windScale := ""
		if ws, ok := nowinfo["windScale"].(string); ok {
			windScale = ws
		}
		humidity := ""
		if h, ok := nowinfo["humidity"].(float64); ok {
			humidity = fmt.Sprintf("%.0f%%", h)
		}

		if weather1, ok := data["weather1"].(string); ok {
			result["tip"] = fmt.Sprintf("当前%s，温度%s，%s%s，湿度%s", weather1, temp, windDir, windScale, humidity)
		}
	}

	return result
}

// ListIP 获取IP统计
// @Summary 获取IP访问统计
// @Description 获取今日、昨日及总体的访问统计数据
func ListIP(c *gin.Context) {
	now := time.Now()

	// 今日时间范围
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayEnd := now

	// 昨日时间范围
	yesterday := now.AddDate(0, 0, -1)
	yesterdayStart := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, now.Location())
	yesterdayEnd := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 0, now.Location())

	// 统计总访问量
	var totalSum int64
	config.DB.Model(&models.IP{}).Count(&totalSum)

	// 统计今日访问量
	var todaySum int64
	config.DB.Model(&models.IP{}).Where("create_time BETWEEN ? AND ?", todayStart, todayEnd).Count(&todaySum)

	// 统计昨日访问量
	var yesterdaySum int64
	config.DB.Model(&models.IP{}).Where("create_time BETWEEN ? AND ?", yesterdayStart, yesterdayEnd).Count(&yesterdaySum)

	// 省份访问TOP10（所有时间）
	provinceAllTop := getProvinceStats(nil, nil, 10)

	// IP访问TOP10（所有时间）
	ipAllTop := getIPStats(nil, nil, 10)

	// 今日统计
	provinceTodayList := getProvinceStats(&todayStart, &todayEnd, 0)
	ipTodayList := getIPStats(&todayStart, &todayEnd, 0)
	userTodayList := getUserStats(todayStart, todayEnd)

	// 昨日统计
	provinceYesterdayList := getProvinceStats(&yesterdayStart, &yesterdayEnd, 0)
	ipYesterdayList := getIPStats(&yesterdayStart, &yesterdayEnd, 0)
	userYesterdayList := getUserStats(yesterdayStart, yesterdayEnd)

	data := map[string]interface{}{
		"total_sum":     totalSum,
		"today_sum":     todaySum,
		"yesterday_sum": yesterdaySum,
		"data": map[string]interface{}{
			"province_all_top": provinceAllTop,
			"ip_all_top":       ipAllTop,
			"today": map[string]interface{}{
				"province": provinceTodayList,
				"ip":       ipTodayList,
				"users":    userTodayList,
			},
			"yesterday": map[string]interface{}{
				"province": provinceYesterdayList,
				"ip":       ipYesterdayList,
				"users":    userYesterdayList,
			},
		},
	}

	response.Success(c, data)
}

// RecordIP 记录IP访问
// @Summary 记录IP访问
// @Description 记录用户访问的IP和地区信息
func RecordIP(c *gin.Context) {
	var req RecordIPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failure(c, "参数错误: "+err.Error())
		return
	}

	ip := GetClientIP(c)

	// 从 Token 获取 userId（已登录则取 token 中的 userId，未登录则为 0）
	userID := 0
	if client, exists := c.Get("client"); exists {
		if clientInfo, ok := client.(*models.Client); ok {
			userID = int(clientInfo.UserID)
		}
	}

	currentTime := time.Now()
	todayStart := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	todayEnd := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())

	if userID > 0 {
		handleUserIPRecord(userID, ip, req.Province, req.City, todayStart, todayEnd, currentTime)
	} else {
		handleAnonymousIPRecord(ip, req.Province, req.City, todayStart, todayEnd, currentTime)
	}

	response.SuccessMessage(c, "记录成功")
}

// MapData 获取地图数据
// @Summary 获取地图数据
// @Description 返回地图JSON数据
func MapData(c *gin.Context) {
	file, err := os.ReadFile(GetStaticFilePath("map.json"))
	if err != nil {
		response.Failure(c, "读取地图数据失败")
		return
	}

	var data interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		response.Failure(c, "解析地图数据失败")
		return
	}

	response.Success(c, data)
}

// UserMapData 获取用户地图数据
// @Summary 获取用户分布地图
// @Description 按省份统计用户分布
func UserMapData(c *gin.Context) {
	type Result struct {
		Province string
		Count    int64
	}

	var results []Result
	config.DB.Model(&models.Client{}).
		Select("province, COUNT(*) as count").
		Where("province IS NOT NULL AND province != ''").
		Group("province").
		Scan(&results)

	// 构建省份映射
	provinceMap := make(map[string]int64)
	for _, r := range results {
		provinceMap[r.Province] = r.Count
	}

	// 读取JSON模板并合并数据
	data := loadAndMergeProvinceData(GetStaticFilePath("user.json"), provinceMap)
	response.Success(c, data)
}

// ==================== 辅助函数 ====================

// getProvinceStats 获取省份统计
func getProvinceStats(start, end *time.Time, limit int) []map[string]interface{} {
	type CountResult struct {
		Province string
		City     string
		Count    int64
	}

	query := config.DB.Model(&models.IP{}).
		Select("province, city, COUNT(*) as count").
		Where("province IS NOT NULL AND province != ''")

	if start != nil && end != nil {
		query = query.Where("create_time BETWEEN ? AND ?", *start, *end)
	}

	query = query.Group("province, city").Order("count DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}

	var results []CountResult
	query.Scan(&results)

	data := make([]map[string]interface{}, 0, len(results))
	for _, r := range results {
		city := r.City
		if city == "" {
			city = "-"
		}
		data = append(data, map[string]interface{}{
			"province": r.Province,
			"city":     city,
			"count":    r.Count,
		})
	}
	return data
}

// getIPStats 获取IP统计
func getIPStats(start, end *time.Time, limit int) []map[string]interface{} {
	type CountResult struct {
		Name  string
		Count int64
	}

	query := config.DB.Model(&models.IP{}).
		Select("ip as name, COUNT(*) as count").
		Where("ip IS NOT NULL AND ip != ''")

	if start != nil && end != nil {
		query = query.Where("create_time BETWEEN ? AND ?", *start, *end)
	}

	query = query.Group("ip").Order("count DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}

	var results []CountResult
	query.Scan(&results)

	data := make([]map[string]interface{}, 0, len(results))
	for _, r := range results {
		data = append(data, map[string]interface{}{
			"ip":    r.Name,
			"count": r.Count,
		})
	}
	return data
}

// getUserStats 获取访问用户统计
func getUserStats(start, end time.Time) []map[string]interface{} {
	var userIDs []int
	config.DB.Model(&models.IP{}).
		Where("create_time BETWEEN ? AND ? AND user_id > 0", start, end).
		Distinct("user_id").
		Pluck("user_id", &userIDs)

	if len(userIDs) == 0 {
		return []map[string]interface{}{}
	}

	// 批量查询用户信息
	uintIDs := make([]uint, len(userIDs))
	for i, id := range userIDs {
		uintIDs[i] = uint(id)
	}
	clientMap := batchGetUserClients(uintIDs)

	data := make([]map[string]interface{}, 0, len(userIDs))
	for _, uid := range userIDs {
		if client, ok := clientMap[uint(uid)]; ok {
			data = append(data, map[string]interface{}{
				"userId":   client.UserID,
				"avatar":   client.Avatar,
				"userName": client.Username,
			})
		}
	}
	return data
}

// handleUserIPRecord 处理有用户ID的IP记录
func handleUserIPRecord(userID int, ip, province, city string, todayStart, todayEnd, currentTime time.Time) {
	// 检查今天这个用户用这个IP是否已经访问过
	var userIPRecord models.IP
	err := config.DB.Where("ip = ? AND user_id = ? AND create_time BETWEEN ? AND ?",
		ip, userID, todayStart, todayEnd).First(&userIPRecord).Error

	if err == nil {
		// 已有记录，更新信息
		userIPRecord.Province = &province
		userIPRecord.City = &city
		userIPRecord.CreateTime = currentTime
		config.DB.Save(&userIPRecord)
		return
	}

	// 检查是否有今天这个IP的匿名记录
	var anonymousRecord models.IP
	err = config.DB.Where("ip = ? AND user_id = 0 AND create_time BETWEEN ? AND ?",
		ip, todayStart, todayEnd).First(&anonymousRecord).Error

	if err == nil {
		// 有匿名记录，更新为该用户的记录
		anonymousRecord.UserID = userID
		anonymousRecord.Province = &province
		anonymousRecord.City = &city
		anonymousRecord.CreateTime = currentTime
		config.DB.Save(&anonymousRecord)
		return
	}

	// 检查今天这个IP是否有任何记录
	var anyRecord models.IP
	err = config.DB.Where("ip = ? AND create_time BETWEEN ? AND ?",
		ip, todayStart, todayEnd).First(&anyRecord).Error

	if err != nil {
		// 完全没有记录，创建新记录
		createIPRecord(userID, ip, province, city, currentTime)
	} else if anyRecord.UserID == 0 {
		// 有匿名记录，更新它
		anyRecord.UserID = userID
		anyRecord.Province = &province
		anyRecord.City = &city
		anyRecord.CreateTime = currentTime
		config.DB.Save(&anyRecord)
	} else {
		// 已有其他用户的记录，为当前用户创建新记录
		createIPRecord(userID, ip, province, city, currentTime)
	}
}

// handleAnonymousIPRecord 处理匿名IP记录
func handleAnonymousIPRecord(ip, province, city string, todayStart, todayEnd, currentTime time.Time) {
	var anonymousRecord models.IP
	err := config.DB.Where("ip = ? AND user_id = 0 AND create_time BETWEEN ? AND ?",
		ip, todayStart, todayEnd).First(&anonymousRecord).Error

	if err == nil {
		// 已有匿名记录，更新信息
		anonymousRecord.Province = &province
		anonymousRecord.City = &city
		anonymousRecord.CreateTime = currentTime
		config.DB.Save(&anonymousRecord)
		return
	}

	// 检查是否有该IP的其他用户记录
	var anyUserRecord models.IP
	err = config.DB.Where("ip = ? AND user_id > 0 AND create_time BETWEEN ? AND ?",
		ip, todayStart, todayEnd).First(&anyUserRecord).Error

	if err != nil {
		// 今天这个IP完全没有记录，创建匿名记录
		createIPRecord(0, ip, province, city, currentTime)
	}
	// 如果已有登录用户用这个IP访问过，不再创建匿名记录
}

// createIPRecord 创建IP记录
func createIPRecord(userID int, ip, province, city string, createTime time.Time) {
	newIP := models.IP{
		UserID:     userID,
		IP:         &ip,
		Province:   &province,
		City:       &city,
		CreateTime: createTime,
		DCount:     1,
		MCount:     1,
	}
	config.DB.Create(&newIP)
}

// loadAndMergeProvinceData 加载并合并省份数据
func loadAndMergeProvinceData(jsonPath string, provinceMap map[string]int64) []map[string]interface{} {
	file, err := os.ReadFile(jsonPath)
	if err != nil {
		// 读取失败，直接返回数据库数据
		return convertProvinceMapToSlice(provinceMap)
	}

	var provinces []ProvinceData
	if err := json.Unmarshal(file, &provinces); err != nil {
		return convertProvinceMapToSlice(provinceMap)
	}

	// 合并数据
	data := make([]map[string]interface{}, 0, len(provinces))
	for _, p := range provinces {
		value := p.Value
		for dbProvince, count := range provinceMap {
			if strings.Contains(p.Name, dbProvince) || strings.Contains(dbProvince, p.Name) {
				value = strconv.FormatInt(count, 10)
				break
			}
		}
		data = append(data, map[string]interface{}{
			"name":  p.Name,
			"value": value,
		})
	}

	return data
}

// convertProvinceMapToSlice 将省份映射转为切片
func convertProvinceMapToSlice(provinceMap map[string]int64) []map[string]interface{} {
	data := make([]map[string]interface{}, 0, len(provinceMap))
	for province, count := range provinceMap {
		data = append(data, map[string]interface{}{
			"name":  province,
			"value": count,
		})
	}
	return data
}
