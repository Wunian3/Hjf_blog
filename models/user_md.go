package models

import "blog_server/models/ctype"

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName   string           `gorm:"size:36;comment:昵称 " json:"nick_name,select(c|info)"`
	UserName   string           `gorm:"size:36;comment:用户名" json:"user_name"`
	Password   string           `gorm:"size:128;comment:密码" json:"-"`
	Avatar     string           `gorm:"size:256;comment:头像" json:"avatar,select(c)"`
	Email      string           `gorm:"size:128;comment:邮箱" json:"email,select(info)"`
	Tel        string           `gorm:"size:18;comment:手机号" json:"tel"`
	Addr       string           `gorm:"size:64;comment:地址" json:"addr,select(c|info)"` // 地址
	Token      string           `gorm:"size:64;comment:其他平台唯一id" json:"token"`         // 其他平台的唯一id
	IP         string           `gorm:"size:20;comment:ip" json:"ip,select(c)"`
	Role       ctype.Role       `gorm:"size:4;default:1;comment:权限,1管理员,2普通用户,3游客" json:"role,select(info)"`   // 权限 1 管理员 2 普通用户 3 游客
	SignStatus ctype.SignStatus `gorm:"type=smallint(6);comment:注册来源,1qq,3邮箱" json:"sign_status,select(info)"` // 注册来源
	Scope      int              `gorm:"default:0;comment:我的积分" json:"scope,select(info)"`                      //我的积分
	Integral   int              `gorm:"default:0;comment:我的积分" json:"integral,select(info)"`                   //我的积分
	Sign       string           `gorm:"size:128;comment:我的签名" json:"sign,select(info)"`
	Link       string           `gorm:"size:128;comment:我的链接地址" json:"link,select(info)"`
}
