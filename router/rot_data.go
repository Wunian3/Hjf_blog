package router

import (
	"blog_server/api"
)

func (router RouterGroup) RouterData() {
	apiuse := api.ApiGroupApp.ApiData
	router.GET("data_login", apiuse.SevenLogin)
	router.GET("data_sum", apiuse.DataSum)
}
