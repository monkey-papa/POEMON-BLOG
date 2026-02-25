package models

import (
	"time"
)

// Family 家庭信息表（用于展示情侣/家人信息）
type Family struct {
	ID             uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID         uint      `gorm:"column:user_id_id" json:"userId"`
	BgCover        string    `gorm:"column:bg_cover;size:256" json:"bgCover"`
	ManCover       string    `gorm:"column:man_cover;size:256" json:"manCover"`
	WomanCover     string    `gorm:"column:woman_cover;size:256" json:"womanCover"`
	ManName        string    `gorm:"column:man_name;size:32" json:"manName"`
	WomanName      string    `gorm:"column:woman_name;size:32" json:"womanName"`
	Timing         string    `gorm:"column:timing;size:32" json:"timing"`
	CountdownTitle *string   `gorm:"column:countdown_title;size:32" json:"countdownTitle,omitempty"`
	CountdownTime  *string   `gorm:"column:countdown_time;size:32" json:"countdownTime,omitempty"`
	Status         bool      `gorm:"column:status;default:true" json:"status"`
	FamilyInfo     *string   `gorm:"column:family_info;size:1024" json:"familyInfo,omitempty"`
	LikeCount      int       `gorm:"column:like_count;default:0" json:"likeCount"`
	CreateTime     time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

func (Family) TableName() string {
	return "family"
}
