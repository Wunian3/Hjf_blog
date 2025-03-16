package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
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
