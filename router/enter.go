package router

import (
	"blog_server/api/api_user"
	"blog_server/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	router.GET("login", api_user.ApiUser{}.QQLogin)
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.RouterSettings()
	routerGroupApp.RouterImages()
	routerGroupApp.RouterAdt()
	routerGroupApp.RouterMenu()
	routerGroupApp.RouterUser()
	routerGroupApp.RouterTag()
	routerGroupApp.RouterMsg()
	routerGroupApp.RouterArticle()
	routerGroupApp.RouterDigg()
	routerGroupApp.RouterComment()
	routerGroupApp.RouterNew()
	routerGroupApp.RouterChat()
	return router
}
