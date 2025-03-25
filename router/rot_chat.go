package router

import "blog_server/api"

func (router RouterGroup) RouterChat() {
	apiuse := api.ApiGroupApp.ApiChat
	router.GET("chat_groups", apiuse.ChatGroup)
	router.GET("chat_groups_records", apiuse.ChatList)

}
