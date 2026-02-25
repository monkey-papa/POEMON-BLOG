package models

// Sort 分类表
type Sort struct {
	ID              uint   `gorm:"primaryKey;column:id" json:"id"`
	SortName        string `gorm:"column:sort_name;size:32" json:"sortName"`
	SortDescription string `gorm:"column:sort_description;size:256" json:"sortDescription"`
	SortType        int    `gorm:"column:sort_type;default:1" json:"sortType"` // 0:导航栏分类, 1:普通分类
	Priority        *int   `gorm:"column:priority" json:"priority,omitempty"`
}

func (Sort) TableName() string {
	return "sort"
}

// Label 标签表
type Label struct {
	ID               uint   `gorm:"primaryKey;column:id" json:"id"`
	SortID           uint   `gorm:"column:sort_id" json:"sortId"`
	LabelName        string `gorm:"column:label_name;size:32" json:"labelName"`
	LabelDescription string `gorm:"column:label_description;size:256" json:"labelDescription"`
}

func (Label) TableName() string {
	return "label"
}
