package api_chat

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ApiChat) ChatList(c *gin.Context) {
	var cr models.PageInf
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	cr.Sort = "created_at desc"

	list, count, _ := common.ComList(models.ChatModel{IsGroup: true}, common.Option{
		PageInf: cr,
	})
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ChatModel, 0)
		res.OkWithList(list, count, c)
		return
	}
	res.OkWithList(data, count, c)
}
