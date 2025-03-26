package api_data

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type DataSumRes struct {
	UserCount      int `json:"user_count"`
	ArticleCount   int `json:"article_count"`
	MessageCount   int `json:"message_count"`
	ChatGroupCount int `json:"chat_group_count"`
	NowLoginCount  int `json:"now_login_count"`
	NowSignCount   int `json:"now_sign_count"`
}

func (ApiData) DataSum(c *gin.Context) {

	var userCount, articleCount, messageCount, ChatGroupCount int
	var nowLoginCount, nowSignCount int

	result, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())
	articleCount = int(result.Hits.TotalHits.Value) //搜索到结果总条数
	global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&userCount)
	global.DB.Model(models.MsgModel{}).Select("count(id)").Scan(&messageCount)
	global.DB.Model(models.ChatModel{IsGroup: true}).Select("count(id)").Scan(&ChatGroupCount)
	global.DB.Model(models.LogDataMd{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowLoginCount)
	global.DB.Model(models.UserModel{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowSignCount)

	res.OkWithData(DataSumRes{
		UserCount:      userCount,
		ArticleCount:   articleCount,
		MessageCount:   messageCount,
		ChatGroupCount: ChatGroupCount,
		NowLoginCount:  nowLoginCount,
		NowSignCount:   nowSignCount,
	}, c)
}
