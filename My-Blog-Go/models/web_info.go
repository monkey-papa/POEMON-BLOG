package models

// WebInfo 网站信息表
type WebInfo struct {
	ID              uint    `gorm:"primaryKey;column:id" json:"id"`
	WebName         string  `gorm:"column:web_name;size:16" json:"webName"`
	WebTitle        string  `gorm:"column:web_title;size:512" json:"webTitle"`
	Notices         *string `gorm:"column:notices;size:512" json:"notices,omitempty"`
	Footer          string  `gorm:"column:footer;size:256" json:"footer"`
	BackgroundImage *string `gorm:"column:background_image;size:256" json:"backgroundImage,omitempty"`
	Avatar          string  `gorm:"column:avatar;size:256" json:"avatar"`
	RandomAvatar    *string `gorm:"column:random_avatar;type:text" json:"randomAvatar,omitempty"`
	RandomName      *string `gorm:"column:random_name;size:4096" json:"randomName,omitempty"`
	RandomCover     *string `gorm:"column:random_cover;type:text" json:"randomCover,omitempty"`
	WaifuJson       *string `gorm:"column:waifu_json;type:text" json:"waifuJson,omitempty"`
	Status          bool    `gorm:"column:status;default:true" json:"status"`
}

func (WebInfo) TableName() string {
	return "web_info"
}
