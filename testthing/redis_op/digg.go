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

	ser_redis.Digg("grtRuZUBmwGUChGr7ByC")
	fmt.Println(ser_redis.GetDigg("grtRuZUBmwGUChGr7ByC"))
	fmt.Println(ser_redis.GetDiggInf())
	//ser_redis.DiggClear()
}
