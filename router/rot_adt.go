package router

import "blog_server/api"

func (router RouterGroup) RouterAdt() {
	apiuse := api.ApiGroupApp.ApiAdt
	router.POST("adts", apiuse.AdvertCreateView)
	router.GET("adts", apiuse.AdvertListView)
	router.DELETE("adts", apiuse.AdvertDeleteView)
	router.PUT("adts/:id", apiuse.AdvertUpdateView)
}
