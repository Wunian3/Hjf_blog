package api_article

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/ser_es"
	"blog_server/service/ser_redis"
	"github.com/gin-gonic/gin"
)

func (ApiArticle) ArticleDetail(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	ser_redis.Look(cr.ID)
	model, err := ser_es.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ApiArticle) ArticleDetailByTitle(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	model, err := ser_es.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
