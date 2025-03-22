package router

import "blog_server/api"

func (router RouterGroup) RouterDigg() {
	apiuse := api.ApiGroupApp.ApiDigg
	router.POST("digg/article", apiuse.DiggArticle)

}
