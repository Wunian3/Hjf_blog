package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterUser() {
	apiuse := api.ApiGroupApp.ApiUser
	router.POST("email_login", apiuse.EmailLogin)
	router.GET("users", middle.JwtAuth(), apiuse.UserList)
	router.PUT("users_role", middle.JwtAdmin(), apiuse.UserRoleUpdate)
	router.PUT("users_password", middle.JwtAuth(), apiuse.UserPasswordUpdate)
	router.POST("users_logout", middle.JwtAuth(), apiuse.UserLogout)
	//router.DELETE("images", apiuse.ImageDelete)
	//router.PUT("images", apiuse.ImageUpdate)

}
