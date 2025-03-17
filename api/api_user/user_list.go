package api_user

import (
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/service/common"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (ApiUser) UserList(c *gin.Context) {
	// 管理员的判定
	token := c.Request.Header.Get("token")
	if token == "" {
		res.FailWithMessage("未携带token", c)
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMessage("token错误", c)
		return
	}

	var page models.PageInf
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInf: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermisssionAdmin {
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, user)
	}

	res.OkWithList(users, count, c)
}
