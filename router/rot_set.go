package router

import (
	"blog_server/api"
)

func (router RouterGroup) RouterSettings() {
	apiSettings := api.ApiGroupApp.ApiSettings
	router.GET("settings/:name", apiSettings.SettingsInfoView)
	router.PUT("settings/:name", apiSettings.SettingsInfoUpdate)

}
