package api_msg

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
)

func (ApiMsg) MsgAllList(c *gin.Context) {
	var cr models.PageInf
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.MsgModel{}, common.Option{
		PageInf: cr,
	})
	res.OkWithList(list, count, c)
}
