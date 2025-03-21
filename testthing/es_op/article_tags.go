package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
}

type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

func main() {
	core.IninCf()
	global.Log = core.InitLog()
	global.ESClient = core.EsConnect()
	agg := elastic.NewTermsAggregation().Field("tags")

	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))

	query := elastic.NewBoolQuery()
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}
	var tapType TagsType
	var tagList = make([]TagsResponse, 0)
	_ = json.Unmarshal(result.Aggregations["tags"], &tapType)
	for _, bucket := range tapType.Buckets {
		var articleList []string
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}
		tagList = append(tagList, TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
	}
	fmt.Println(tagList)
}
