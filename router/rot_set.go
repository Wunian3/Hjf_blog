package router

import (
	"blog_server/api"
)

func (router RouterGroup) RouterSettings() {
	apiuse := api.ApiGroupApp.ApiSettings
	router.GET("settings/:name", apiuse.SettingsInfoView)
	router.PUT("settings/:name", apiuse.SettingsInfoUpdate)

}
