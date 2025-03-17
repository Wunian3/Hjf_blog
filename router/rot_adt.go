package router

import "blog_server/api"

func (router RouterGroup) RouterAdt() {
	apiuse := api.ApiGroupApp.ApiAdt
	router.POST("adts", apiuse.AdvertCreate)
	router.GET("adts", apiuse.AdvertList)
	router.DELETE("adts", apiuse.AdvertDelete)
	router.PUT("adts/:id", apiuse.AdvertUpdate)
}
