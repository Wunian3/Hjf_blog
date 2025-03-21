package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterArticle() {
	apiuse := api.ApiGroupApp.ApiArticle
	router.POST("articles", middle.JwtAuth(), apiuse.ArticleCreate)
	router.GET("articles", apiuse.ArticleList)
	router.GET("articles/detail", apiuse.ArticleDetailByTitle)
	router.GET("articles/calendar", apiuse.ArticleCalendar)
	router.GET("articles/:id", apiuse.ArticleDetail)

}
