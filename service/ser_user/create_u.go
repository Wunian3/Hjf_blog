package ser_user

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/utils/pwd"
	"errors"
)

const Avatar = "/uploads/avatar/default.png"

func (ServiceUser) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {

	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	hashPwd := pwd.HashPwd(password)
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     Avatar,
		IP:         ip,
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
