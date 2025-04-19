package api_big_model

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
)

// RoleListView 列表
func (ApiBigModel) RoleListView(c *gin.Context) {
	var cr models.PageInf
	c.ShouldBindQuery(&cr)

	list, count, _ := common.ComList(models.BigModelRoleModel{}, common.Option{
		PageInf: cr,
		Likes:   []string{"name"},
		Preload: []string{"Tags"},
	})
	res.OkWithList(list, count, c)
	return
}
