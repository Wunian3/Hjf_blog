package api_advert

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// AdvertList 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInf   false  "id列表"
// @Router /api/adts [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListRes[models.AdtModel]}
func (ApiAdvert) AdvertList(c *gin.Context) {
	var cr models.PageInf
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	referer := c.GetHeader("Referer")
	//是否后台来的，管理员有权限
	isShow := true
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	list, count, _ := common.ComList(models.AdtModel{IsShow: isShow}, common.Option{
		PageInf: cr,
		Debug:   true,
	})
	res.OkWithList(list, count, c)
}
