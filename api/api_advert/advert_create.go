package api_advert

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法"` // 链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法"`
	IsShow bool   `json:"is_show" binding:"required" msg:"请选择是否展示"`
}

func (ApiAdvert) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advert models.AdtModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该广告已存在", c)
		return
	}

	err = global.DB.Create(&models.AdtModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败,hjf你再看看", c)
		return
	}

	res.OkWithMessage("添加广告成功，hjf你再看看", c)
}
