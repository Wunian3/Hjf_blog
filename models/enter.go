package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"  structs:"-"`
	UpdatedAt time.Time `json:"-" structs:"-"`
	CreatedAt time.Time `json:"created_at" structs:"-"`
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
