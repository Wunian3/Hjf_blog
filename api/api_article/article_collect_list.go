package api_article

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"blog_server/utils/jwts"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"strings"
)

type CollectRes struct {
	models.ArticleModel
	CreatedAt string `json:"created_at"`
}

func (ApiArticle) ArticleCollList(c *gin.Context) {

	var cr models.PageInf

	c.ShouldBindQuery(&cr)

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var articleIDList []interface{}

	list, count, err := common.ComList(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInf: cr,
	})

	var collMap = map[string]string{}

	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}

	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)

	var collList = make([]CollectRes, 0)

	// 传id列表，查es
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		article.ID = hit.Id
		article.Content = ""
		collectRes := CollectRes{
			ArticleModel: article,
			CreatedAt:    collMap[hit.Id],
		}
		// Process BannerUrl: add leading / if needed
		if strings.HasPrefix(collectRes.BannerUrl, "uploads/") && !strings.HasPrefix(collectRes.BannerUrl, "/") {
			collectRes.BannerUrl = "/" + collectRes.BannerUrl
		}
		collList = append(collList, collectRes)
	}
	res.OkWithList(collList, count, c)
}
