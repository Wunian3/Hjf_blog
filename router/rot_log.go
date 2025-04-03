package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterLog() {
	apiuse := api.ApiGroupApp.ApiLogV2
	router.GET("logs/v2", apiuse.LogList)
	router.DELETE("logs/v2", middle.JwtAdmin(), apiuse.LogDelete)
	router.GET("logs/v2/read", apiuse.LogRead)
}
