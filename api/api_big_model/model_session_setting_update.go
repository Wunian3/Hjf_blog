package api_big_model

import (
	"blog_server/conf"
	"blog_server/core"
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (ApiBigModel) ModelSessionSettingUpdateView(c *gin.Context) {
	var cr conf.SessionSetting
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	global.Config.BigModel.SessionSetting = cr
	core.SetYaml()
	res.OkWithMessage("修改成功", c)
	return
}
