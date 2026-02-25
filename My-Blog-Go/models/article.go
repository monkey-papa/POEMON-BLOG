package models

import (
	"time"
)

// Article 文章表
type Article struct {
	ID              uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID          uint      `gorm:"column:user_id" json:"userId"`
	SortID          uint      `gorm:"column:sort_id" json:"sortId"`
	LabelID         uint      `gorm:"column:label_id" json:"labelId"`
	ArticleCover    *string   `gorm:"column:article_cover;size:256" json:"articleCover,omitempty"`
	ArticleTitle    string    `gorm:"column:article_title;size:32" json:"articleTitle"`
	ArticleContent  string    `gorm:"column:article_content;type:text" json:"articleContent"`
	Summary         string    `gorm:"column:summary;type:text" json:"summary"`
	ViewCount       int       `gorm:"column:view_count;default:0" json:"viewCount"`
	LikeCount       int       `gorm:"column:like_count;default:0" json:"likeCount"`
	ViewStatus      bool      `gorm:"column:view_status;default:true" json:"viewStatus"`
	Password        *string   `gorm:"column:password;size:128" json:"password,omitempty"`
	RecommendStatus bool      `gorm:"column:recommend_status;default:false" json:"recommendStatus"`
	CommentStatus   bool      `gorm:"column:comment_status;default:true" json:"commentStatus"`
	CreateTime      time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime      time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	UpdateBy        *string   `gorm:"column:update_by;size:32" json:"updateBy,omitempty"`
	Deleted         bool      `gorm:"column:deleted;default:false" json:"deleted"`
}

func (Article) TableName() string {
	return "article"
}

// ArticleLike 文章点赞表
type ArticleLike struct {
	ID        uint `gorm:"primaryKey;column:id" json:"id"`
	UserID    uint `gorm:"column:user_id_id" json:"userId"`
	ArticleID uint `gorm:"column:article_id_id" json:"articleId"`
}

func (ArticleLike) TableName() string {
	return "like"
}
