package api_advert

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AdvertDeleteView 广告删除
// @Tags 广告删除
// @Summary 广告删除
// @Description 广告删除
// @Param data body models.RemoveRequest   true  "表示多个参数"
// @Router /api/adts [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (ApiAdvert) AdvertDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var advertList []models.AdtModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个广告", count), c)

}
