// Package handlers 图片上传相关接口处理
package handlers

import (
	"my-blog-go/middleware"
	"my-blog-go/response"
	"my-blog-go/utils"

	"github.com/gin-gonic/gin"
)

// UploadImage 上传图片到七牛云
// @Summary 上传图片
// @Description 将图片上传到七牛云并返回访问URL
// 注意：此接口需要使用表单上传，包含 image、folder 字段
func UploadImage(c *gin.Context) {
	// 从 JWT Token 获取当前登录用户
	currentUser, exists := middleware.GetCurrentClient(c)
	if !exists {
		response.Failure(c, "未登录")
		return
	}

	// 获取上传的文件
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		response.Failure(c, "获取上传文件失败: "+err.Error())
		return
	}
	defer file.Close()

	// 获取文件夹参数
	folder := c.PostForm("folder")
	if folder == "" {
		folder = "images"
	}

	// 获取当前用户的七牛云配置
	client, err := getUserClient(currentUser.UserID)
	if err != nil {
		response.Failure(c, "用户不存在")
		return
	}

	// 检查七牛云配置
	if client.QiniuAccessKey == nil || client.QiniuSecretKey == nil ||
		client.QiniuBucketName == nil || client.QiniuDomain == nil {
		response.Failure(c, "用户未配置七牛云")
		return
	}

	// 构建七牛云配置
	qiniuCfg := utils.QiniuConfig{
		AccessKey:  *client.QiniuAccessKey,
		SecretKey:  *client.QiniuSecretKey,
		BucketName: *client.QiniuBucketName,
		Domain:     *client.QiniuDomain,
	}

	// 上传到七牛云
	imageURL, err := utils.UploadToQiniu(file, fileHeader, qiniuCfg, folder)
	if err != nil {
		response.Failure(c, "上传失败: "+err.Error())
		return
	}

	response.Success(c, gin.H{"url": imageURL})
}
