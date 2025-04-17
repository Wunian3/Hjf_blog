package api_big_model

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type RoleItem struct {
	ID       uint   `json:"id"`       // 角色id
	Name     string `json:"name"`     // 角色名称
	Abstract string `json:"abstract"` // 角色简介
	Icon     string `json:"icon"`     //角色图标
}

type TagRoleListResponse struct {
	ID       uint       `json:"id"`       // 标签的id
	Title    string     `json:"title"`    // 名称
	RoleList []RoleItem `json:"roleList"` // 角色列表

}

// TagRoleListView 角色广场
func (ApiBigModel) TagRoleListView(c *gin.Context) {

	var _list []models.BigModelTagModel
	global.DB.Preload("Roles").Find(&_list)
	var list = make([]TagRoleListResponse, 0)
	for _, model := range _list {
		roleList := make([]RoleItem, 0)
		for _, role := range model.Roles {
			roleList = append(roleList, RoleItem{
				ID:       role.ID,
				Name:     role.Name,
				Abstract: role.Abstract,
				Icon:     role.Icon,
			})
		}
		list = append(list, TagRoleListResponse{
			ID:       model.ID,
			Title:    model.Title,
			RoleList: roleList,
		})
	}

	res.OkWithData(list, c)
}
