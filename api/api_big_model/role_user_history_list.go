package api_big_model

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

// RoleUserHistoryListView 用户角色历史列表
func (ApiBigModel) RoleUserHistoryListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var roleIdList []uint
	global.DB.Model(models.BigModelSessionModel{}).Where("user_id = ?", claims.UserID).Group("role_id").Select("role_id").Scan(&roleIdList)
	var roleList []models.BigModelRoleModel
	global.DB.Order("created_at desc").Find(&roleList, "id in ?", roleIdList)

	var list = make([]RoleItem, 0)
	for _, model := range roleList {
		list = append(list, RoleItem{
			ID:       model.ID,
			Name:     model.Name,
			Abstract: model.Abstract,
			Icon:     model.Icon,
		})
	}
	res.OkWithData(list, c)
	return
}
