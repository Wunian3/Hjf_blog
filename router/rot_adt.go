package router

import "blog_server/api"

func (router RouterGroup) RouterAdt() {
	adt := api.ApiGroupApp.ApiAdt
	router.POST("adts", adt.AdvertCreateView)
	router.GET("adts", adt.AdvertListView)
	router.DELETE("adts", adt.AdvertDeleteView)
	router.PUT("adts/:id", adt.AdvertUpdateView)
}
