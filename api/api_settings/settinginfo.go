package api_settings

import (
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// 不太适合生成文档的类型，有点小缺陷，主要是把多个接口浓缩了，接口的入参和出参不统一
func (ApiSettings) SettingsInfo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hjfapi启动"})
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInf, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("无配置信息，检查setinf！hjf滚回去上班", c)
	}
}
