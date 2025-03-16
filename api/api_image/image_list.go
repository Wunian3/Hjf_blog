package api_image

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
)

// ImageListView 图片列表
func (ApiImage) ImageListView(c *gin.Context) {
	var cr models.PageInf
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInf: cr,
		Debug:   false,
	})

	res.OkWithList(list, count, c)

	return

}
