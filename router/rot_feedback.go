package router

import "blog_server/api"

func (router RouterGroup) RouterFeedback() {
	app := api.ApiGroupApp.ApiFeedback
	router.POST("feedback", app.FeedBackCreate)
	router.GET("feedback", app.FeedBackList)
	router.DELETE("feedback", app.FeedBackRemove)
}
