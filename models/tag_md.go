package models

// TagMd 标签表
type TagModel struct {
	MODEL
	Title string `gorm:"size:16" json:"title"`
}
