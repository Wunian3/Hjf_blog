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
	router.GET("message_users", apiuse.MessageUserList)
	router.GET("message_users/me", apiuse.MessageUserListByMe)
	router.GET("message_users/record", apiuse.MessageUserRecord)
	router.GET("message_users/record/me", apiuse.MessageUserRecordByMe)
	router.GET("message_users/user", apiuse.MessageUserListByUser)
	router.DELETE("message_users", apiuse.MessageRecordDelete)

}
