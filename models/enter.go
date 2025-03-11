package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}
