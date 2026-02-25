package response

import (
	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Error   bool        `json:"error"`
	Success bool        `json:"success"`
}

// PageData 分页数据结构
type PageData struct {
	TotalCount int64       `json:"totalCount"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int64       `json:"totalPage"`
	Page       int         `json:"page"`
	List       interface{} `json:"list"`
	Empty      bool        `json:"empty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    0,
		Msg:     "操作成功",
		Data:    data,
		Error:   false,
		Success: true,
	})
}

// SuccessPage 成功响应（带分页）
func SuccessPage(c *gin.Context, list interface{}, total int64, page int, pageSize int) {
	var totalPage int64 = 1
	if pageSize > 0 {
		totalPage = total / int64(pageSize)
		if total%int64(pageSize) > 0 {
			totalPage++
		}
	}

	empty := total == 0

	c.JSON(200, Response{
		Code: 0,
		Msg:  "操作成功",
		Data: PageData{
			TotalCount: total,
			PageSize:   pageSize,
			TotalPage:  totalPage,
			Page:       page,
			List:       list,
			Empty:      empty,
		},
		Error:   false,
		Success: true,
	})
}

// SuccessMessage 成功响应（只有消息）
func SuccessMessage(c *gin.Context, message string) {
	c.JSON(200, Response{
		Code:    0,
		Msg:     message,
		Data:    nil,
		Error:   false,
		Success: true,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Msg:     message,
		Data:    nil,
		Error:   true,
		Success: false,
	})
}

// Failure 失败响应 (HTTP 200，但业务失败)
func Failure(c *gin.Context, message string) {
	c.JSON(200, Response{
		Code:    -1,
		Msg:     message,
		Data:    nil,
		Error:   true,
		Success: false,
	})
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context, message string) {
	c.JSON(401, Response{
		Code:    401,
		Msg:     message,
		Data:    nil,
		Error:   true,
		Success: false,
	})
}

// Forbidden 禁止访问响应
func Forbidden(c *gin.Context, message string) {
	c.JSON(403, Response{
		Code:    403,
		Msg:     message,
		Data:    nil,
		Error:   true,
		Success: false,
	})
}
