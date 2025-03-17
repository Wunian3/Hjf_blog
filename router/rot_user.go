package router

import "blog_server/api"

func (router RouterGroup) RouterUser() {
	apiuse := api.ApiGroupApp.ApiUser
	router.POST("email_login", apiuse.EmailLogin)
	router.GET("users", apiuse.UserList)
	//router.GET("images_name", apiuse.ImagNameList)
	//router.DELETE("images", apiuse.ImageDelete)
	//router.PUT("images", apiuse.ImageUpdate)

}
