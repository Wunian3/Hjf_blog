package main

import (
	"blog_server/core"
	_ "blog_server/docs"
	"blog_server/flag"
	"blog_server/global"
	"blog_server/router"
	//"flag"
	"fmt"
)

// 配置
// @title blog_server API文档
// @version 1.0
// @description blog_server API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	core.IninCf()
	global.Log = core.InitLog()
	//connect hblog_db
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
	option := flag.Parse()
	//fmt.Println(option)
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	router := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("%s ,hjfblog启动！", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
