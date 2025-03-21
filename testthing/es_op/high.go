package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/models"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	// 读取配置文件
	core.IninCf()

	// 初始化日志
	global.Log = core.InitLog()
	// 连接es
	global.ESClient = core.EsConnect()
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMultiMatchQuery("node", "title", "abstract", "content")).
		Highlight(elastic.NewHighlight().Field("title")).
		Size(100).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	for _, hit := range result.Hits.Hits {
		fmt.Println(string(hit.Source))
		fmt.Println(hit.Highlight)
	}
}

// 搜索出来的标题就是 <em>node</em>基础语法
