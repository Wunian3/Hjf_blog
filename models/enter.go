package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id,select($any)"  structs:"-"`
	UpdatedAt time.Time `gorm:"comment:更新时间" json:"-" structs:"-"`
	CreatedAt time.Time `gorm:"comment:创建时间" json:"created_at,select($any)" structs:"-"`
}
type PageInf struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"`
	Sort  string `form:"sort"`
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}

type IDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type Options[T any] struct {
	Label string `json:"label"`
	Value T      `json:"value"`
}

const (
	AdminRole   = 1
	UserRole    = 2
	TouristRole = 3
)
