package router

import "blog_server/api"

func (router RouterGroup) RouterTag() {
	apiuse := api.ApiGroupApp.ApiTag
	router.POST("tags", apiuse.TagCreate)
	router.GET("tags", apiuse.TagList)
	router.GET("tag_names", apiuse.TagNameList)
	router.DELETE("tags", apiuse.TagDelete)
	router.PUT("tags/:id", apiuse.TagUpdate)
}
