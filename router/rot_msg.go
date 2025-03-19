package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterMsg() {
	apiuse := api.ApiGroupApp.ApiMsg
	router.POST("msgs", apiuse.MsgCreate)
	router.GET("msgs_all", apiuse.MsgAllList)
	router.GET("msgs", middle.JwtAuth(), apiuse.MsgList)
	router.GET("msgs_record", middle.JwtAuth(), apiuse.MsgRecord)

}
