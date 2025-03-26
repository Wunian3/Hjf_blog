package api_log

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/plugin/log_stash"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ApiLog) LogDelete(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var list []log_stash.LogStashModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("日志不存在", c)
		return
	}
	global.DB.Delete(&list)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个日志", count), c)

}
