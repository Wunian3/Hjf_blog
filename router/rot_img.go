package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterImages() {
	apiuse := api.ApiGroupApp.ApiImages
	router.POST("images", apiuse.ImageUpload)
	router.GET("images", apiuse.ImageList)
	router.GET("images_name", apiuse.ImagNameList)
	router.DELETE("images", apiuse.ImageDelete)
	router.PUT("images", apiuse.ImageUpdate)
	router.POST("image", middle.JwtAdmin(), apiuse.ImageUploadData)

}
