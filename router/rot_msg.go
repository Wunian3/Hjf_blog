package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterMsg() {
	apiuse := api.ApiGroupApp.ApiMsg
	router.POST("msgs", middle.JwtAuth(), apiuse.MsgCreate)
	router.GET("msgs_all", apiuse.MsgAllList)
	router.GET("msgs", middle.JwtAuth(), apiuse.MsgList)
	router.GET("msgs_record", middle.JwtAuth(), apiuse.MsgRecord)
	router.GET("message_users", middle.JwtAuth(), apiuse.MessageUserList)
	router.GET("message_users/me", middle.JwtAuth(), apiuse.MessageUserListByMe)
	router.GET("message_users/record", apiuse.MessageUserRecord)
	router.GET("message_users/record/me", middle.JwtAuth(), apiuse.MessageUserRecordByMe)
	router.GET("message_users/user", middle.JwtAuth(), apiuse.MessageUserListByUser)
	router.DELETE("message_users", middle.JwtAuth(), apiuse.MessageRecordDelete)

}
