package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// QiniuConfig 七牛云配置
type QiniuConfig struct {
	AccessKey  string
	SecretKey  string
	BucketName string
	Domain     string
}

// UploadToQiniu 上传文件到七牛云
func UploadToQiniu(file multipart.File, fileHeader *multipart.FileHeader, cfg QiniuConfig, folder string) (string, error) {
	if cfg.AccessKey == "" || cfg.SecretKey == "" || cfg.BucketName == "" {
		return "", fmt.Errorf("七牛云配置不完整")
	}

	// 生成唯一的文件名
	ext := filepath.Ext(fileHeader.Filename)
	key := fmt.Sprintf("%s/%s%s", folder, strings.ReplaceAll(uuid.New().String(), "-", ""), ext)

	// 创建凭证
	mac := qbox.NewMac(cfg.AccessKey, cfg.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: cfg.BucketName,
	}
	upToken := putPolicy.UploadToken(mac)

	// 配置上传（使用自动检测区域，无需手动指定）
	formUploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseHTTPS:      true,
		UseCdnDomains: false,
	})

	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	// 上传文件
	err := formUploader.Put(context.Background(), &ret, upToken, key, file, fileHeader.Size, &putExtra)
	if err != nil {
		return "", fmt.Errorf("上传失败: %v", err)
	}

	// 返回访问URL
	url := cfg.Domain + ret.Key
	return url, nil
}

// DeleteFromQiniu 从七牛云删除文件
func DeleteFromQiniu(imageURL string, cfg QiniuConfig) error {
	if cfg.AccessKey == "" || cfg.SecretKey == "" || cfg.BucketName == "" {
		return fmt.Errorf("七牛云配置不完整")
	}

	// 从URL中提取key
	key := strings.TrimPrefix(imageURL, cfg.Domain)
	if key == "" || key == imageURL {
		return fmt.Errorf("无效的图片URL")
	}

	mac := qbox.NewMac(cfg.AccessKey, cfg.SecretKey)
	bucketManager := storage.NewBucketManager(mac, &storage.Config{
		Zone:     &storage.ZoneHuadong,
		UseHTTPS: true,
	})

	err := bucketManager.Delete(cfg.BucketName, key)
	if err != nil {
		return fmt.Errorf("删除失败: %v", err)
	}

	return nil
}
