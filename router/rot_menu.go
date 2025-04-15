package router

import "blog_server/api"

func (router RouterGroup) RouterMenu() {
	apiuse := api.ApiGroupApp.ApiMenu
	router.POST("menus", apiuse.MenuCreate)
	router.GET("menus", apiuse.MenuList)
	router.GET("menu_names", apiuse.MenuNameList)
	router.DELETE("menus", apiuse.MenuDelete)
	router.PUT("menus/:id", apiuse.MenuUpdate)
	router.GET("menus/:id", apiuse.MenuInf)
	router.GET("menus/detail", apiuse.MenuDetailByPath)

}
