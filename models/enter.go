package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"  structs:"-"`
	UpdatedAt time.Time `json:"-" structs:"-"`
	CreatedAt time.Time `json:"createdAt" structs:"-"`
}
type PageInf struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Key   string `json:"key"`
	Sort  string `json:"sort"`
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
