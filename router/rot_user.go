package router

import (
	"blog_server/api"
	"blog_server/middle"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("HJFBLOG124"))

func (router RouterGroup) RouterUser() {
	apiuse := api.ApiGroupApp.ApiUser
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", apiuse.EmailLogin)
	router.POST("login", apiuse.QQLogin)
	router.GET("users", middle.JwtAuth(), apiuse.UserList)
	router.PUT("users_role", middle.JwtAdmin(), apiuse.UserRoleUpdate)
	router.PUT("users_password", middle.JwtAuth(), apiuse.UserPasswordUpdate)
	router.POST("users_logout", middle.JwtAuth(), apiuse.UserLogout)
	router.DELETE("users", middle.JwtAdmin(), apiuse.UserDelete)
	router.POST("user_email_bind", middle.JwtAuth(), apiuse.UserEmailBind)
}
