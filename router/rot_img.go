package router

import "blog_server/api"

func (router RouterGroup) RouterImages() {
	apiuse := api.ApiGroupApp.ApiImages
	router.POST("images", apiuse.ImageUploadView)
	router.GET("images", apiuse.ImageListView)
	router.GET("image_name", apiuse.ImagNameList)
	router.DELETE("images", apiuse.ImageDeleteView)
	router.PUT("images", apiuse.ImageUpdateView)

}
