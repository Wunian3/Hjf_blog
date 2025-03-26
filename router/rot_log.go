package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterLog() {
	apiuse := api.ApiGroupApp.ApiLog
	router.GET("logs", apiuse.LogList)
	router.DELETE("logs", middle.JwtAdmin(), apiuse.LogDelete)
}
