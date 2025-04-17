package api_article

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/ser_es"
	"blog_server/service/ser_redis"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

type ArticleSearchRequest struct {
	models.PageInf
	Tag    string `json:"tag" form:"tag"`
	IsUser bool   `json:"is_user" form:"is_user"` // 根据这个参数判断是否显示我收藏的文章列表
	Date   string `json:"date" form:"date"`
}

func (ApiArticle) ArticleList(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	boolSearch := elastic.NewBoolQuery()

	if cr.IsUser {
		token := c.GetHeader("token")
		claims, err := jwts.ParseToken(token)
		if err == nil && !ser_redis.CheckLogout(token) {
			boolSearch.Must(elastic.NewTermsQuery("user_id", claims.UserID))
		}
	}

	if cr.Date != "" {
		date, err := time.Parse("2006-01-02", cr.Date)
		if err == nil {
			boolSearch.Must(elastic.NewRangeQuery("created_at").
				Gte(date.Format("2006-01-02") + "00:00:00").
				Lte(date.Format("2006-01-02") + "23:59:59"))
		}
	}

	list, count, err := ser_es.CommList(ser_es.Option{
		PageInf: cr.PageInf,
		Fields:  []string{"title", "content", "category"},
		Tag:     cr.Tag,
		Query:   boolSearch,
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}

	for i := range list {
		if strings.HasPrefix(list[i].BannerUrl, "uploads/") && !strings.HasPrefix(list[i].BannerUrl, "/") {
			list[i].BannerUrl = "/" + list[i].BannerUrl
		}
		if strings.HasPrefix(list[i].UserAvatar, "uploads/") && !strings.HasPrefix(list[i].UserAvatar, "/") {
			list[i].UserAvatar = "/" + list[i].UserAvatar
		}
	}

	// json-filter空值问题
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}

	res.OkWithList(data, int64(count), c)
}
