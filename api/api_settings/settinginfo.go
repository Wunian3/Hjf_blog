package api_settings

import (
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (ApiSettings) SettingsInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "hjfapi启动"})
	res.FailWithCode(res.SettingsError, c)
}
