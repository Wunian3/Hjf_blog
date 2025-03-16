package api

import (
	"blog_server/api/api_advert"
	"blog_server/api/api_image"
	"blog_server/api/api_settings"
)

type ApiGroup struct {
	ApiSettings api_settings.ApiSettings
	ApiImages   api_image.ApiImage
	ApiAdt      api_advert.ApiAdvert
}

var ApiGroupApp = new(ApiGroup)
