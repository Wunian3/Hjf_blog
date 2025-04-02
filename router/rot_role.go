package router

import "blog_server/api"

func (router RouterGroup) RouterRole() {
	app := api.ApiGroupApp.ApiRole
	router.GET("role_ids", app.RoleIDList)
}
