package api_user

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/plugin/log_stash_v2"
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

	actionLog := log_stash_v2.NewAction(c) // 初始化操作日志
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		log_stash_v2.NewFailLogin("用户名不存在", cr.UserName, cr.Password, c)
		res.FailWithMessage("用户名或密码错误", c)
		return
	}

	if !pwd.CheckPwd(userModel.Password, cr.Password) {
		log_stash_v2.NewFailLogin("密码错误", cr.UserName, cr.Password, c)
		res.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 生成Token，确保包含UserName字段
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		UserName: userModel.UserName, // 明确传递UserName
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		actionLog.Error(fmt.Sprintf("token生成失败 %s ", err.Error()))
		res.FailWithMessage("token生成失败", c)
		return
	}

	// 关键修复：将新生成的token设置到请求头
	c.Request.Header.Set("token", token) // 确保后续解析使用最新token
	log_stash_v2.NewSuccessLogin(c)      // 此时能正确获取UserName

	// 其他记录逻辑...
	res.OkWithData(token, c)
}

//func (ApiUser) EmailLogin(c *gin.Context) {
//	var cr EmailLoginRequest
//	err := c.ShouldBindJSON(&cr)
//	if err != nil {
//		res.FailWithError(err, &cr, c)
//		return
//	}
//
//	log := log_stash.NewLogByGin(c)
//
//	var userModel models.UserModel
//	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
//	if err != nil {
//		global.Log.Warn("用户名不存在")
//		log.Warn(fmt.Sprintf("%s 用户名不存在", cr.UserName))
//		res.FailWithMessage("用户名或密码错误", c)
//		return
//	}
//	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
//	if !isCheck {
//		global.Log.Warn("用户名密码错误")
//		log.Warn(fmt.Sprintf("用户名密码错误 %s %s ", cr.UserName, cr.Password))
//		res.FailWithMessage("用户名或密码错误", c)
//		return
//	}
//	// 登录成功，生成token
//	token, err := jwts.GenToken(jwts.JwtPayLoad{
//		NickName: userModel.NickName,
//		Role:     int(userModel.Role),
//		UserID:   userModel.ID,
//		//Avatar:   userModel.Avatar,
//	})
//	if err != nil {
//		global.Log.Error(err)
//		log.Warn(fmt.Sprintf("token生成失败 %s ", err.Error()))
//		res.FailWithMessage("token生成失败", c)
//
//		return
//	}
//	ip, addr := utils.GetAddrByGin(c)
//	log = log_stash.New(c.ClientIP(), token)
//	log.Info("登录成功")
//	global.DB.Create(&models.LogDataMd{
//		UserID:    userModel.ID,
//		IP:        ip,
//		NickName:  userModel.NickName,
//		Token:     token,
//		Device:    "",
//		Addr:      addr,
//		LoginType: ctype.SignEmail,
//	})
//
//	res.OkWithData(token, c)
//
//}
