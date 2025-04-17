package models

type UserBigModelRoleModel struct {
	MODEL
	UserID    uint              `json:"userID"` //用户id
	RoleID    uint              `json:"roleID"` //角色id
	RoleModel BigModelRoleModel `gorm:"foreignkey:RoleID" json:"-"`
}
