package api_article

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/ser_redis"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// ArticleContentByID 获取文章正文
// @Tags 文章管理
// @Summary 获取文章正文
// @Description 获取文章正文
// @Param id path int  true  "id"
// @Router /api/articles/content/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (ApiArticle) ArticleContentByID(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	ser_redis.NewArticleLook().Set(cr.ID)

	result, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(cr.ID).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage("查询失败", c)
		return
	}
	var model models.ArticleModel
	err = json.Unmarshal(result.Source, &model)
	if err != nil {
		return
	}
	res.OkWithData(model.Content, c)
}
