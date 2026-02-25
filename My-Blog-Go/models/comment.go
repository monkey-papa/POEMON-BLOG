package models

import (
	"time"
)

// Comment 评论表
type Comment struct {
	ID              uint      `gorm:"primaryKey;column:id" json:"id"`
	Source          int       `gorm:"column:source;index:idx_comment_source" json:"source"`
	Type            string    `gorm:"column:type;size:32" json:"type"`
	ParentCommentID *uint     `gorm:"column:parent_comment_id" json:"parentCommentId,omitempty"`
	UserID          uint      `gorm:"column:user_id" json:"userId"`
	FloorCommentID  *uint     `gorm:"column:floor_comment_id" json:"floorCommentId,omitempty"`
	ParentUserID    *uint     `gorm:"column:parent_user_id" json:"parentUserId,omitempty"`
	LikeCount       int       `gorm:"column:like_count;default:0" json:"likeCount"`
	CommentContent  string    `gorm:"column:comment_content;size:1024" json:"commentContent"`
	CommentInfo     *string   `gorm:"column:comment_info;size:256" json:"commentInfo,omitempty"`
	CreateTime      time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
}

func (Comment) TableName() string {
	return "comment"
}
