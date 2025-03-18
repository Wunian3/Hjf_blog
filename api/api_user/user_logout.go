package api_user

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/service"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (ApiUser) UserLogout(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")

	err := service.ServiceGroupApp.ServiceUser.Logout(claims, token)

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}

	res.OkWithMessage("注销成功", c)

}
