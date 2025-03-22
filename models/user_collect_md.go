package models

import "time"

// 用户收藏表 自己定义用
type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"size:32"`
	CreatedAt time.Time
}
