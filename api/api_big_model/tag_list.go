package api_big_model

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
)

type TagListResponse struct {
	models.MODEL
	Title     string `json:"title"`     // 名称
	Color     string `json:"color"`     // 颜色
	RoleCount int    `json:"roleCount"` // 角色个数
}

// TagListView 标签新增和更新
func (ApiBigModel) TagListView(c *gin.Context) {
	var cr models.PageInf
	c.ShouldBindQuery(&cr)
	_list, count, _ := common.ComList(models.BigModelTagModel{}, common.Option{
		Likes:   []string{"title"},
		Preload: []string{"Roles"},
	})
	var list = make([]TagListResponse, 0)
	for _, model := range _list {
		list = append(list, TagListResponse{
			MODEL:     model.MODEL,
			Title:     model.Title,
			Color:     model.Color,
			RoleCount: len(model.Roles),
		})
	}
	res.OkWithList(list, count, c)
}
