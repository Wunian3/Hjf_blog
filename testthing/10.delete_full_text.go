package main

import (
	"blog_server/core"
	"blog_server/global"
	"blog_server/service/ser_es"
)

func main() {
	core.IninCf()
	core.InitLog()
	global.ESClient = core.EsConnect()
	ser_es.DeleteFullTextByArticleID("dNYEvZUB9_VTBFgseSX7")
}
