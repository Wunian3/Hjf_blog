package models

type UserScopeModel struct {
	MODEL
	UserID uint `json:"userID"`
	Scope  int  `json:"scope"`
	Status bool `json:"status"`
}
