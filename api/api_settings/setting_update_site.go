package api_settings

import (
	"blog_server/conf"
	"blog_server/core"
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

// 修改某一项的配置信息
func (ApiSettings) SiteInfoUpdateSite(c *gin.Context) {
	var info conf.SiteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = info
	core.SetYaml()
	res.OkWithMessage("网站信息更新成功", c)
	
}
