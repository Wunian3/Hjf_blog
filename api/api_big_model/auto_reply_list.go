package api_big_model

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
)

// AutoReplyListView 列表
func (ApiBigModel) AutoReplyListView(c *gin.Context) {
	var cr models.PageInf
	c.ShouldBindQuery(&cr)

	list, count, _ := common.ComList(models.AutoReplyModel{}, common.Option{
		PageInf: cr,
		Likes:   []string{"name"},
	})
	res.OkWithList(list, count, c)
	return
}
