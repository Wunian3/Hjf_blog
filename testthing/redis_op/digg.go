package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/service/ser_redis"
	"fmt"
)

func main() {
	// 读取配置文件
	core.IninCf()
	// 初始化日志
	global.Log = core.InitLog()

	global.Redis = core.ConnectRedis()
	global.ESClient = core.EsConnect()
	digg := ser_redis.NewDigg()
	digg.Set("grtRuZUBmwGUChGr7ByC")
	//ser_redis.Digg("grtRuZUBmwGUChGr7ByC")
	fmt.Println(digg.Get("grtRuZUBmwGUChGr7ByC"))

	fmt.Println(digg.GetInfo())
	//ser_redis.DiggClear()
}
