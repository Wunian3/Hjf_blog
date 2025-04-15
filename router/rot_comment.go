package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterComment() {
	apiuse := api.ApiGroupApp.ApiComment
	router.POST("comments", middle.JwtAuth(), apiuse.CommentCreate)
	router.GET("comments/:id", apiuse.CommentList)
	router.GET("comments/digg/:id", apiuse.CommentDigg)
	router.GET("comments/articles", middle.JwtAdmin(), apiuse.CommentByArticleList)
	router.DELETE("comments/:id", middle.JwtAuth(), apiuse.CommentDelete)

}
