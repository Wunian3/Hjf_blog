package api_user

import (
	"blog_server/global"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/service/ser_user"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 用户前台服务
type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string     `json:"password" binding:"required" msg:"请输入密码"`
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"` // 1管理员2普通用户3游客
}

// 用户前台服务
func (ApiUser) UserCreate(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err := ser_user.ServiceUser{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), c)
	return
}
