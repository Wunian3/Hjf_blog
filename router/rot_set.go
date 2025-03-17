package router

import (
	"blog_server/api"
)

func (router RouterGroup) RouterSettings() {
	apiuse := api.ApiGroupApp.ApiSettings
	router.GET("settings/:name", apiuse.SettingsInfo)
	router.PUT("settings/:name", apiuse.SettingsInfoUpdate)

}
