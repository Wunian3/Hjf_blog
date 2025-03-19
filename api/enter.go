package api

import (
	"blog_server/api/api_advert"
	"blog_server/api/api_image"
	"blog_server/api/api_menu"
	"blog_server/api/api_settings"
	"blog_server/api/api_tag"
	"blog_server/api/api_user"
)

type ApiGroup struct {
	ApiSettings api_settings.ApiSettings
	ApiImages   api_image.ApiImage
	ApiAdt      api_advert.ApiAdvert
	ApiMenu     api_menu.ApiMenu
	ApiUser     api_user.ApiUser
	ApiTag      api_tag.ApiTag
}

var ApiGroupApp = new(ApiGroup)
