package router

import "blog_server/api"

func (router RouterGroup) RouterNew() {
	apiuse := api.ApiGroupApp.ApiNew
	router.GET("news", apiuse.NewList)
}
