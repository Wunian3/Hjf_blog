package router

import (
	"blog_server/api"
)

func (router RouterGroup) RouterData() {
	apiuse := api.ApiGroupApp.ApiData
	router.GET("data_login_week", apiuse.LoginWeek)
	router.GET("data_sum", apiuse.DataSum)
}
