package api_image

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// ImageList 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.PageInf   false  "查询参数"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListRes[models.BannerModel]}
func (ApiImage) ImageList(c *gin.Context) {
	var cr models.PageInf
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInf: cr,
		Debug:   true,
	})
	for i := range list {
		path := list[i].Path
		if strings.HasPrefix(path, "uploads/") {
			// 如果是本地路径且没有前导斜杠，则添加 "/"
			if !strings.HasPrefix(path, "/") {
				list[i].Path = "/" + path
			}
		}
	}

	res.OkWithList(list, count, c)

	return

}
