package api_tag

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
)

func (ApiTag) TagList(c *gin.Context) {
	var cr models.PageInf
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.TagModel{}, common.Option{
		PageInf: cr,
		Debug:   true,
	})
	//todo 需要展示标签下文章的一些内容 后续补充
	res.OkWithList(list, count, c)
}
