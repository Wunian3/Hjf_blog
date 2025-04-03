package api_user

import (
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"blog_server/service/common"
	"blog_server/utils/desens"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserRes struct {
	models.UserModel
	RoleId int `json:"role_id"`
}
type UserListRes struct {
	models.PageInf
	Role int `json:"role" form:"role"`
}

// UserList 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Router /api/users [put]
// @Param token header string  true  "token"
// @Param data query models.PageInf  false  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListRes[models.UserModel]}
func (ApiUser) UserList(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page UserListRes
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []UserRes
	list, count, _ := common.ComList(models.UserModel{Role: ctype.Role(page.Role)}, common.Option{
		PageInf: page.PageInf,
		Likes:   []string{"nick_name"},
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermisssionAdmin {
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, UserRes{
			UserModel: user,
			RoleId:    int(user.Role),
		})
	}

	res.OkWithList(users, count, c)

	//_claims, _ := c.Get("claims")
	//claims := _claims.(*jwts.CustomClaims)
	//
	//var page models.PageInf
	//if err := c.ShouldBindQuery(&page); err != nil {
	//	res.FailWithCode(res.ArgumentError, c)
	//	return
	//}
	//var users []models.UserModel
	//list, count, _ := common.ComList(models.UserModel{}, common.Option{
	//	PageInf: page,
	//})
	//for _, user := range list {
	//	if ctype.Role(claims.Role) != ctype.PermisssionAdmin {
	//		user.UserName = ""
	//	}
	//	user.Tel = desens.DesensitizationTel(user.Tel)
	//	user.Email = desens.DesensitizationEmail(user.Email)
	//	users = append(users, user)
	//}
	//
	//res.OkWithList(users, count, c)
}
