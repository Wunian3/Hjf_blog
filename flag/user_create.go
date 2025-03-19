package flag

import (
	"blog_server/global"
	"blog_server/models/ctype"
	"blog_server/service/ser_user"
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
	if password != rePassword {
		global.Log.Error("两次密码不一致，请重新输入")
		return
	}
	role := ctype.PermisssionUser
	if permissions == "admin" {
		role = ctype.PermisssionAdmin
	}

	err := ser_user.ServiceUser{}.CreateUser(userName, nickName, password, role, email, "127.0.0.1")
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("创建用户%s", userName)

}
