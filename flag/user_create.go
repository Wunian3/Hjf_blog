package flag

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/utils/pwd"
	"fmt"
)

func CreateUser(permissions string) {
	//需要用户名，昵称，密码，确认密码，邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入邮箱：")
	fmt.Scan(&email)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	fmt.Printf("请再次输入密码：")
	fmt.Scan(&rePassword)
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		global.Log.Error("用户名已存在，请重新输入")
		return
	}
	if password != rePassword {
		global.Log.Error("两次密码不一致，请重新输入")
		return
	}
	hashPwd := pwd.HashPwd(password)
	role := ctype.PermisssionUser
	if permissions == "admin" {
		role = ctype.PermisssionAdmin
	}
	//默认头像 或者 随机头像
	avatar := "/uploads/avatar/default.jpg"
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("创建用户%s", userName)

}
