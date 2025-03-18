package api_user

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/plugin/email"
	"blog_server/utils/jwts"
	"blog_server/utils/pwd"
	"blog_server/utils/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (ApiUser) UserEmailBind(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 用户绑定邮箱后台发验证码
	// 生成4位验证码， 将生成的验证码存入session
	// 用户输入邮箱，验证码，密码校验验证码
	// 修改邮箱第一次的邮箱，和第二次的邮箱做一致性校验
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//回看session教程
	session := sessions.Default(c)
	if cr.Code == nil {
		code := random.Code(4)
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "你的验证码是 "+code)
		if err != nil {
			global.Log.Error(err)
		}
		res.OkWithMessage("验证码已发送，请查收", c)
		return
	}
	code := session.Get("valid_code")

	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMessage("密码强度太低", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}
	// 完成绑定
	res.OkWithMessage("邮箱绑定成功", c)
	return
}
