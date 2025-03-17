package api_image

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type ImageRes struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `gorm:"size:38" json:"name"`
}

// ImagNameList 图片名称列表
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/image_name [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]ImageRes}
func (ApiImage) ImagNameList(c *gin.Context) {

	var imageList []ImageRes
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)

	res.OkWithData(imageList, c)
}
