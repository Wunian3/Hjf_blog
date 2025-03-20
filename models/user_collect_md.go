package models

import "time"

// 用户收藏表 自己定义用
type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID uint      `gorm:"primaryKey"`
	CreatedAt time.Time
}
