package api_big_model

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

// TagOptionsListView 标签id列表
func (ApiBigModel) TagOptionsListView(c *gin.Context) {
	var list []models.Options[uint]
	global.DB.Model(models.BigModelTagModel{}).Select("id as value", "title as label").Scan(&list)
	res.OkWithData(list, c)
}
