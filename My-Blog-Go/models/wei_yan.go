package models

import (
	"time"
)

// WeiYan 微言表
type WeiYan struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint      `gorm:"column:user_id;index" json:"userId"`
	LikeCount  int       `gorm:"column:like_count;default:0" json:"likeCount"`
	Content    string    `gorm:"column:content;size:1024" json:"content"`
	Type       string    `gorm:"column:type;size:32" json:"type"`
	Source     *int      `gorm:"column:source" json:"source,omitempty"`
	IsPublic   bool      `gorm:"column:is_public;default:false" json:"isPublic"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
}

func (WeiYan) TableName() string {
	return "wei_yan"
}
