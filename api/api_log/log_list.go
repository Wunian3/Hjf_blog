package api_log

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/plugin/log_stash"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	models.PageInf
	Level log_stash.Level `form:"level"`
}

func (ApiLog) LogList(c *gin.Context) {
	var cr LogRequest
	c.ShouldBindQuery(&cr)
	list, count, _ := common.ComList(log_stash.LogStashModel{Level: cr.Level}, common.Option{
		PageInf: cr.PageInf,
		Debug:   true,
		Likes:   []string{"ip", "addr"},
	})
	res.OkWithList(list, count, c)
	return
}
