package api_big_model

import (
	"blog_server/conf"
	"blog_server/core"
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

// ModelSettingUpdateView 更新大模型配置
func (ApiBigModel) ModelSettingUpdateView(c *gin.Context) {
	var cr conf.Setting
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	// 验证这个name有没有乱写
	var ok bool
	for _, option := range global.Config.BigModel.ModelList {
		if option.Value == cr.Name {
			ok = true
			break
		}
	}
	if !ok {
		// 没有找到
		res.FailWithMessage("名称错误", c)
		return
	}

	global.Config.BigModel.Setting = cr
	core.SetYaml()
	res.OkWithMessage("修改成功", c)
	return
}
