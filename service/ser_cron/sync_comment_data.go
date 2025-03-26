package ser_cron

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/service/ser_redis"
	"gorm.io/gorm"
)

// SyncCommentData 同步评论数据到数据库
func SyncCommentData() {
	commentDiggInf := ser_redis.NewCommentDigg().GetInfo()
	for key, count := range commentDiggInf {
		var comment models.CommentModel
		err := global.DB.Take(&comment, key).Error

		if err != nil {
			global.Log.Error(err)
			continue
		}
		err = global.DB.Model(&comment).
			Update("digg_count", gorm.Expr("digg_count + ?", count)).Error
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功 新的点赞数: %d", comment.Content, comment.DiggCount)
	}
	ser_redis.NewCommentDigg().Clear()
}
