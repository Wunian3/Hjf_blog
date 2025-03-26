package api_user

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/plugin/log_stash"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (ApiUser) EmailLogin(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	log := log_stash.NewLogByGin(c)

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在")
		log.Warn(fmt.Sprintf("%s 用户名不存在", cr.UserName))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		log.Warn(fmt.Sprintf("用户名密码错误 %s %s ", cr.UserName, cr.Password))

		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		log.Warn(fmt.Sprintf("token生成失败 %s ", err.Error()))
		res.FailWithMessage("token生成失败", c)

		return
	}
	log = log_stash.New(c.ClientIP(), token)
	log.Info("登录成功")

	res.OkWithData(token, c)

}
