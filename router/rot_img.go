package router

import "blog_server/api"

func (router RouterGroup) RouterImages() {
	imgup := api.ApiGroupApp.ApiImages
	router.POST("images", imgup.ImageUploadView)

}
