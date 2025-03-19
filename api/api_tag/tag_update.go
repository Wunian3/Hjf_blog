package api_tag

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (ApiTag) TagUpdate(c *gin.Context) {

	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
		return
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error
	// 结构体转map
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("标签修改失败", c)
		return
	}

	res.OkWithMessage("标签修改成功", c)
}
