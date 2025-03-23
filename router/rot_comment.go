package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterComment() {
	apiuse := api.ApiGroupApp.ApiComment
	router.POST("comments", middle.JwtAuth(), apiuse.CommentCreate)
	router.GET("comments", apiuse.CommentList)
	router.GET("comments/:id", apiuse.CommentDigg)
	router.DELETE("comments/:id", apiuse.CommentDelete)

}
