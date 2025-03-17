package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/utils/jwts"
	"fmt"
)

func main() {
	core.IninCf()
	global.Log = core.InitLog()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "hjf",
		NickName: "blog",
	})
	fmt.Println(token, err)
	//claims, err := jwts.ParseToken("")
	//fmt.Println(claims, err)
}
