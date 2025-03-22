package api_digg

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/ser_redis"
	"github.com/gin-gonic/gin"
)

func (ApiDigg) DiggArticle(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// todo: 对长度校验的难度，后期修补
	ser_redis.Digg(cr.ID)
	res.OkWithMessage("文章点赞成功", c)
}
