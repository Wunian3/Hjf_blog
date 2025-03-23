package api_comment

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/ser_redis"
	"blog_server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (ApiComment) CommentDelete(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", c)
		return
	}
	// 统计评论下的子评论数+1,
	//流程：判断是否是子评论，若否，找父评论，减掉对应的评论数删除子评论以及当前评论，反转，然后一个一个删
	subCommentList := FindSubCommentCount(commentModel)
	count := len(subCommentList) + 1
	ser_redis.NewCommentCount().SetCount(commentModel.ArticleID, -count)

	//
	if commentModel.ParentCommentID != nil {

		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	utils.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 条评论", len(deleteCommentIDList)), c)
	return
}
