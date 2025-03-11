package main

import (
	"blog_server/core"
	"blog_server/flag"
	"blog_server/global"
	"blog_server/router"
	//"flag"
	"fmt"
)

func main() {
	//配置
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
