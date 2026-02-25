package models

import (
	"time"
)

// TreeHole 树洞表
type TreeHole struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint      `gorm:"column:user_id;index;default:0" json:"userId"`    // 用户ID，记录实际发布者
	Avatar     *string   `gorm:"column:avatar;size:256" json:"avatar,omitempty"` // 显示头像（可自定义）
	Username   string    `gorm:"column:username;size:32" json:"username"`        // 显示用户名（可自定义，不唯一）
	Message    string    `gorm:"column:message;size:64" json:"message"`          // 留言内容
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
}

func (TreeHole) TableName() string {
	return "tree_hole"
}
