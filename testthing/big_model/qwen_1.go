package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/service/big_model_ser"
	"fmt"
)

func main() {
	core.IninCf()
	global.Log = core.InitLog()
	msgChan, err := big_model_ser.Send("qwen", "可以讲一个笑话吗")
	if err != nil {
		fmt.Println(err)
		return
	}
	for msg := range msgChan {
		fmt.Println(msg)
	}
}
