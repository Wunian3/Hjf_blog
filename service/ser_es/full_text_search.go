package ser_es

import (
	"blog_server/global"
	"blog_server/models"
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
	"strings"
)

type SearchData struct {
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

func AsyncArticleByFullText(id, title, content string) {
	indexList := GetSearchIndexDataByContent(id, title, content)

	bulk := global.ESClient.Bulk()
	for _, indexdata := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexdata)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return

	}
	logrus.Infof("%s 添加成功，共 %d 条", title, len(result.Succeeded()))
}

func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n")
	var isCode bool = false
	var headList, bodyList []string
	var body string
	headList = append(headList, getHeader(title))
	for _, s := range dataList {
		// #{1,6}
		// 判断一下是否是代码块
		if strings.HasPrefix(s, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(s, "#") && !isCode {
			headList = append(headList, getHeader(s))
			//if strings.TrimSpace(body) != "" {
			bodyList = append(bodyList, getBody(body))
			//}
			body = ""
			continue
		}
		body += s
	}
	bodyList = append(bodyList, getBody(body))
	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
		})
	}
	return searchDataList
}

func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}

func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}
