package api_settings

import (
	"blog_server/conf"
	"blog_server/core"
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

// 修改某一项的配置信息
func (ApiSettings) SettingsInfoUpdate(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		var inf conf.SiteInf
		err = c.ShouldBindJSON(&inf)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInf = inf

	case "email":
		var inf conf.Email
		err = c.ShouldBindJSON(&inf)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = inf
	case "qq":
		var inf conf.QQ
		err = c.ShouldBindJSON(&inf)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = inf
	case "qiniu":
		var inf conf.QiNiu
		err = c.ShouldBindJSON(&inf)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = inf
	case "jwt":
		var inf conf.Jwt
		err = c.ShouldBindJSON(&inf)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = inf
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}

	core.SetYaml()
	res.OkWith(c)
}
