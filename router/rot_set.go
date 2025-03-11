package router

import (
	"blog_server/api"
)

func (router RouterGroup) RouterSettings() {
	apiSettings := api.ApiGroupApp.ApiSettings
	router.GET("settings", apiSettings.SettingsInfoView)
}
