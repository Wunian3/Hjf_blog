package router

import "blog_server/api"

func (router RouterGroup) RouterImages() {
	img := api.ApiGroupApp.ApiImages
	router.POST("images", img.ImageUploadView)
	router.GET("images", img.ImageListView)
	router.DELETE("images", img.ImageDeleteView)
	router.PUT("images", img.ImageUpdateView)
}
