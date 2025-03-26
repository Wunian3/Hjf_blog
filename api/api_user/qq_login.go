package api_user

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	qq "blog_server/plugin/QQ"
	"blog_server/utils"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"
	"blog_server/utils/random"
	"github.com/gin-gonic/gin"
)

func (ApiUser) QQLogin(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	openID := qqInfo.OpenID
	var user models.UserModel
	err = global.DB.Take(&user, "token = ?", openID).Error
	ip, addr := utils.GetAddrByGin(c)
	if err != nil {
		// openID是否存在,不存在，就注册
		hashPwd := pwd.HashPwd(random.RandStr(16))
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID,  // qq登录，邮箱+密码
			Password:   hashPwd, // 随机生成16位密码
			Avatar:     qqInfo.Avatar,
			Addr:       addr, // 根据ip算地址
			Token:      openID,
			IP:         ip,
			Role:       ctype.PermisssionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("注册失败", c)
			return
		}

	}
	// 登录操作
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", c)
		return
	}

	global.DB.Create(&models.LogDataMd{
		UserID:    user.ID,
		IP:        ip,
		NickName:  user.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignQQ,
	})
	res.OkWithData(token, c)
}
