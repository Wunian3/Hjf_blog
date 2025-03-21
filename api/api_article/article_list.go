package api_article

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/ser_es"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ApiArticle) ArticleList(c *gin.Context) {
	var cr models.PageInf
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := ser_es.CommList(cr.Key, cr.Page, cr.Limit)
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}

	res.OkWithList(filter.Omit("list", list), int64(count), c)
}
