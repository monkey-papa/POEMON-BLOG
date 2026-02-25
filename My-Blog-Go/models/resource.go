package models

import (
	"time"
)

// Resource 资源信息表
type Resource struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint      `gorm:"column:user_id_id" json:"userId"`
	Type       string    `gorm:"column:type;size:32" json:"type"`
	Path       string    `gorm:"column:path;size:256" json:"path"`
	Size       *int      `gorm:"column:size" json:"size,omitempty"`
	MimeType   *string   `gorm:"column:mime_type;size:256" json:"mimeType,omitempty"`
	Status     bool      `gorm:"column:status;default:true" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
}

func (Resource) TableName() string {
	return "resource"
}

// ResourcePath 资源路径表
type ResourcePath struct {
	ID           uint      `gorm:"primaryKey;column:id" json:"id"`
	Title        string    `gorm:"column:title;size:64" json:"title"`
	Classify     *string   `gorm:"column:classify;size:32" json:"classify,omitempty"`
	Cover        *string   `gorm:"column:cover;size:256" json:"cover,omitempty"`
	URL          *string   `gorm:"column:url;size:256" json:"url,omitempty"`
	Introduction *string   `gorm:"column:introduction;size:1024" json:"introduction,omitempty"`
	Type         string    `gorm:"column:type;size:32" json:"type"`
	Status       bool      `gorm:"column:status;default:true" json:"status"`
	Remark       *string   `gorm:"column:remark;type:text" json:"remark,omitempty"`
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
}

func (ResourcePath) TableName() string {
	return "resource_path"
}
