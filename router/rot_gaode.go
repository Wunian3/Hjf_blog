package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterGaode() {
	apiuse := api.ApiGroupApp.ApiGaode
	router.GET("gaode/weather", middle.JwtAuth(), apiuse.WeatherInfo)

}
