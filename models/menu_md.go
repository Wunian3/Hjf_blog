package models

import "blog_server/models/ctype"

type MenuModel struct {
	MODEL
	Title        string        `gorm:"size:32" json:"title"`
	Path         string        `gorm:"size:32" json:"path"` //该blog下我使用/path格式偏多，好拼接前后端
	Slogan       string        `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`
	AbstractTime int           `json:"abstract_time"` // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"`
	BannerTime   int           `json:"banner_time"`         // 菜单图片的切换时间 为 0 表示不切换
	Sort         int           `gorm:"size:10" json:"sort"` // 菜单的顺序
}
