package models

import (
	"time"
)

// AuthToken Token表 (兼容Django REST framework的token表)
type AuthToken struct {
	Key     string    `gorm:"primaryKey;column:key;size:40" json:"key"`
	Created time.Time `gorm:"column:created" json:"created"`
	UserID  uint      `gorm:"column:user_id;unique" json:"userId"`
}

func (AuthToken) TableName() string {
	return "authtoken_token"
}
