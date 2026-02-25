package models

import (
	"time"
)

// Code 验证码表
type Code struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	Email      string    `gorm:"column:email" json:"email"`
	Code       string    `gorm:"column:code;size:10" json:"code"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
}

func (Code) TableName() string {
	return "appone_code"
}
