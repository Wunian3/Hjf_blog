package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/models"
	"blog_server/service/ser_es"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	core.IninCf()
	core.InitLog()
	global.ESClient = core.EsConnect()
	boolSearch := elastic.NewMatchAllQuery()
	res, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())

	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)

		indexList := ser_es.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)

		bulk := global.ESClient.Bulk()
		for _, indexdata := range indexList {
			req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexdata)
			bulk.Add(req)
		}
		result, err := bulk.Do(context.Background())
		if err != nil {
			logrus.Error(err)
			continue
		}
		fmt.Println(article.Title, "添加成功", "共", len(result.Succeeded()), "条！")
		
	}
}
