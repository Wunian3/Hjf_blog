package ser_cron

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/service/ser_redis"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
)

// SyncArticleData 同步文章数据到ES上
func SyncArticleData() {
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}
	diggInf := ser_redis.NewDigg().GetInfo()
	lookInf := ser_redis.NewArticleLook().GetInfo()
	commentInf := ser_redis.NewCommentCount().GetInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		digg := diggInf[hit.Id]
		look := lookInf[hit.Id]
		comment := commentInf[hit.Id]
		//计算新数据
		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment

		if digg == 0 && look == 0 && comment == 0 {
			global.Log.Infof("%s无变化", article.Title)
			continue
		}
		_, err = global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"look_count":    newLook,
				"comment_count": newComment,
				"digg_count":    newDigg,
			}).Do(context.Background())
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s更新成功 点赞数:%d 评论数:%d 浏览量:%d", article.Title, newDigg, newComment, newLook)
	}
	//clear redis
	ser_redis.NewDigg().Clear()
	ser_redis.NewArticleLook().Clear()
	ser_redis.NewCommentCount().Clear()
}
