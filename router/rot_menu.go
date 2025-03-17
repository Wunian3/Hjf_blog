package router

import "blog_server/api"

func (router RouterGroup) RouterMenu() {
	apiuse := api.ApiGroupApp.ApiMenu
	router.POST("menus", apiuse.MenuCreate)
	router.GET("menus", apiuse.MenuList)
	router.GET("menus_name", apiuse.MenuNameList)
	//router.DELETE("menus", apiuse.ImageDeleteView)
	router.PUT("menus/:id", apiuse.MenuUpdate)

}
