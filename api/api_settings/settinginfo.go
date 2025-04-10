package api_settings

import (
	"blog_server/global"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

// 不太适合生成文档的类型，有点小缺陷，主要是把多个接口浓缩了，接口的入参和出参不统一
func (ApiSettings) SettingsInfo(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 校验角色，游客用户（Role == 3）不可访问
	if claims.Role == 3 {
		res.FailWithMessage("游客用户无权访问此配置", c)
		return
	}
	// 通过校验后处理其他 case
	switch cr.Name {
	case "email":
		info := global.Config.Email
		info.Password = "******"
		res.OkWithData(info, c)
	case "qq":
		info := global.Config.QQ
		info.Key = "******"
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		info := global.Config.QiNiu
		info.SecretKey = "******"
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		info := global.Config.Jwt
		info.Secret = "******"
		res.OkWithData(info, c)
	case "chat_group":
		res.OkWithData(global.Config.ChatGroup, c)
	default:
		res.FailWithMessage("无配置信息", c)

	}
}
