package api_settings

import (
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (ApiSettings) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hjfapi启动"})
	res.OkWithData(global.Config.SiteInf, c)
}
