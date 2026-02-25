package models

import (
	"time"
)

// AuthUser Django的auth_user表
type AuthUser struct {
	ID          uint       `gorm:"primaryKey;column:id" json:"id"`
	Password    string     `gorm:"column:password;size:128" json:"-"`
	LastLogin   *time.Time `gorm:"column:last_login" json:"lastLogin,omitempty"`
	IsSuperuser bool       `gorm:"column:is_superuser;default:false" json:"isSuperuser"`
	Username    string     `gorm:"column:username;size:150;unique" json:"username"`
	FirstName   string     `gorm:"column:first_name;size:150" json:"firstName"`
	LastName    string     `gorm:"column:last_name;size:150" json:"lastName"`
	Email       string     `gorm:"column:email;size:254" json:"email"`
	IsStaff     bool       `gorm:"column:is_staff;default:false" json:"isStaff"`
	IsActive    bool       `gorm:"column:is_active;default:true" json:"isActive"`
	DateJoined  time.Time  `gorm:"column:date_joined" json:"dateJoined"`
}

func (AuthUser) TableName() string {
	return "auth_user"
}

// Client 用户信息表
type Client struct {
	UserID          uint      `gorm:"primaryKey;column:user_id" json:"id"`
	User            *AuthUser `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Username        string    `gorm:"column:username;size:32;unique" json:"username"`
	PhoneNumber     *string   `gorm:"column:phone_number;size:16" json:"phoneNumber,omitempty"`
	Email           *string   `gorm:"column:email;size:32" json:"email,omitempty"`
	UserStatus      bool      `gorm:"column:user_status;default:true" json:"userStatus"`
	Gender          *int      `gorm:"column:gender;default:0" json:"gender,omitempty"` // 0:保密, 1:男, 2:女
	OpenID          *string   `gorm:"column:open_id;size:128" json:"openId,omitempty"`
	Avatar          *string   `gorm:"column:avatar;size:256" json:"avatar,omitempty"`
	Admire          *string   `gorm:"column:admire;size:32" json:"admire,omitempty"`
	Introduction    *string   `gorm:"column:introduction;size:4096" json:"introduction,omitempty"`
	UserType        int       `gorm:"column:user_type;default:2" json:"userType"` // 0:Boss, 1:管理员, 2:普通用户, 3:访客
	CreateTime      time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime      time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	UpdateBy        *string   `gorm:"column:update_by;size:32" json:"updateBy,omitempty"`
	Deleted         bool      `gorm:"column:deleted;default:false" json:"deleted"`
	Province        *string   `gorm:"column:province;size:20" json:"province,omitempty"`
	QiniuDomain     *string   `gorm:"column:qiniu_domain;size:128" json:"qiniuDomain,omitempty"`
	QiniuBucketName *string   `gorm:"column:qiniu_bucket_name;size:128" json:"qiniuBucketName,omitempty"`
	QiniuSecretKey  *string   `gorm:"column:qiniu_secret_key;size:128" json:"qiniuSecretKey,omitempty"`
	QiniuAccessKey  *string   `gorm:"column:qiniu_access_key;size:128" json:"qiniuAccessKey,omitempty"`
}

func (Client) TableName() string {
	return "user"
}

// UserType 用户类型常量
const (
	UserTypeBoss    = 0 // Boss
	UserTypeAdmin   = 1 // 管理员
	UserTypeNormal  = 2 // 普通用户
	UserTypeVisitor = 3 // 访客
)

// IsAdmin 检查用户是否为管理员（包括Boss）
func (c *Client) IsAdmin() bool {
	return c.UserType == UserTypeBoss || c.UserType == UserTypeAdmin
}

// IsBoss 检查用户是否为Boss
func (c *Client) IsBoss() bool {
	return c.UserType == UserTypeBoss
}
