package models

import (
	"time"
)

// Words 禁用词表
type Words struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	Avatar     *string   `gorm:"column:avatar;size:256" json:"avatar,omitempty"`
	Username   string    `gorm:"column:username;size:32;unique" json:"username"`
	Message    string    `gorm:"column:message;size:64" json:"message"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
}

func (Words) TableName() string {
	return "word"
}
