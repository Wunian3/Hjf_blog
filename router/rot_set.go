package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterSettings() {
	apiuse := api.ApiGroupApp.ApiSettings
	router.GET("settings/site", apiuse.SiteInfo)
	router.GET("settings/:name", middle.JwtAuth(), apiuse.SettingsInfo)
	router.PUT("settings/site", apiuse.SiteInfoUpdateSite)
	router.PUT("settings/:name", middle.JwtAuth(), apiuse.SettingsInfoUpdate)

}
