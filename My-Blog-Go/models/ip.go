package models

import (
	"time"
)

// IP IP地址表
type IP struct {
	ID         uint       `gorm:"primaryKey;column:id" json:"id"`
	UserID     int        `gorm:"column:user_id;default:0" json:"userId"`
	IP         *string    `gorm:"column:ip;size:20" json:"ip,omitempty"`
	Province   *string    `gorm:"column:province;size:20" json:"province,omitempty"`
	City       *string    `gorm:"column:city;size:20" json:"city,omitempty"`
	CreateTime time.Time  `gorm:"column:create_time" json:"createTime"`
	KTime      *time.Time `gorm:"column:k_time" json:"kTime,omitempty"`
	MCount     int        `gorm:"column:mcount;default:0" json:"mCount"`
	DCount     int        `gorm:"column:dcount;default:0" json:"dCount"`
}

func (IP) TableName() string {
	return "ip"
}
