package router

import (
	"blog_server/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.RouterSettings()
	routerGroupApp.RouterImages()
	routerGroupApp.RouterAdt()
	return router
}
