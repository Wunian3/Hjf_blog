package router

import (
	"blog_server/global"
	"blog_server/middle"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.Use(middle.LogMiddleWare())
	router.StaticFS("uploads", http.Dir("uploads"))
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

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
	routerGroupApp.RouterLog()
	routerGroupApp.RouterData()
	routerGroupApp.RouterRole()
	return router
}
