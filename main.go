package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/router"
	"fmt"
)

func main() {
	//配置
	core.IninCf()
	global.Log = core.InitLog()
	//connect hblog_db
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
	router := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("%s ,hjfblog启动！", addr)
	router.Run(addr)
}
