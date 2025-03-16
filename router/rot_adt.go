package router

import "blog_server/api"

func (router RouterGroup) RouterAdt() {
	adt := api.ApiGroupApp.ApiAdt
	router.POST("adts", adt.AdvertCreateView)
	router.GET("adts", adt.AdvertCreateView)
	router.DELETE("adts", adt.AdvertCreateView)
	router.PUT("adts", adt.AdvertCreateView)
}
