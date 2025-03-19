package router

import (
	"blog_server/api"
)

func (router RouterGroup) RouterMsg() {
	apiuse := api.ApiGroupApp.ApiMsg
	router.POST("msgs", apiuse.MsgCreate)
}
