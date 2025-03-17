package api_advert

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
	"strings"
)

func (ApiAdvert) AdvertListView(c *gin.Context) {
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
