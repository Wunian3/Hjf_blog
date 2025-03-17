package api_advert

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (ApiAdvert) AdvertUpdateView(c *gin.Context) {

	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advert models.AdtModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error
	// 结构体转map
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("广告修改失败", c)
		return
	}

	res.OkWithMessage("广告修改成功", c)
}
