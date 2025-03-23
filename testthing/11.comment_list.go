package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/models"
)

func main() {
	core.IninCf()
	global.Log = core.InitLog()
	global.DB = core.InitGorm()
	FindArticleCommentList("o0rawJUBYOskemlVSPnd")
}

func FindArticleCommentList(articleId string) {

	var RootCommentList []models.CommentModel
	global.DB.Find(&RootCommentList, "article_id = ? and parent_commment_id is null", articleId)

}
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
	return
}
