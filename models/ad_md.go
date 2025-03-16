package models

type AdtModel struct {
	MODEL
	Title  string `gorm:"size:32" json:"title"` // 标题
	Href   string `json:"href"`                 // 链接
	Images string `json:"images"`               // 图片
	IsShow bool   `json:"is_show"`
}
