package main

import (
	"blog_server/models"
	"fmt"
	"github.com/fatih/structs"
)

type AdvertRequest struct {
	models.MODEL `structs:"-"`
	Title        string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`
	Href         string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"-"` // 链接
	Images       string `json:"images" binding:"required,url" msg:"图片地址非法"`
	IsShow       bool   `json:"is_show" binding:"required" msg:"请选择是否展示" structs:"is_show"`
}

func main() {
	u1 := AdvertRequest{
		Title:  "xxx",
		Href:   "xxx",
		Images: "xxx",
		IsShow: true,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}
