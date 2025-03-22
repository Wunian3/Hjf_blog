package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/models"
	"blog_server/service/ser_redis"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	// 读取配置文件
	core.IninCf()
	// 初始化日志
	global.Log = core.InitLog()

	global.Redis = core.ConnectRedis()
	global.ESClient = core.EsConnect()
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	diggInf := ser_redis.GetDiggInf()
	lookInf := ser_redis.GetLookInf()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		digg := diggInf[hit.Id]
		look := lookInf[hit.Id]
		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		if article.DiggCount == newDigg && article.LookCount == newLook {
			logrus.Info(article.Title, "点赞数和浏览数无变化")
			continue
		}
		_, err := global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count": newDigg,
				"look_count": newLook,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Infof("%s,点赞数据同步成功， 点赞数 %d 浏览数 %d", article.Title, newDigg, newLook)
	}
	ser_redis.DiggClear()
	ser_redis.LookClear()

}
