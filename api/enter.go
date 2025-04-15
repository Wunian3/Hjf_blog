package api

import (
	"blog_server/api/api_advert"
	"blog_server/api/api_article"
	"blog_server/api/api_chat"
	"blog_server/api/api_comment"
	"blog_server/api/api_data"
	"blog_server/api/api_digg"
	"blog_server/api/api_gaode"
	"blog_server/api/api_image"
	"blog_server/api/api_log"
	"blog_server/api/api_log_v2"
	"blog_server/api/api_menu"
	"blog_server/api/api_msg"
	"blog_server/api/api_new"
	"blog_server/api/api_role"
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
	ApiMsg      api_msg.ApiMsg
	ApiArticle  api_article.ApiArticle
	ApiDigg     api_digg.ApiDigg
	ApiComment  api_comment.ApiComment
	ApiNew      api_new.ApiNew
	ApiChat     api_chat.ApiChat
	ApiLog      api_log.ApiLog
	ApiLogV2    api_log_v2.ApiLog
	ApiData     api_data.ApiData
	ApiRole     api_role.ApiRole
	ApiGaode    api_gaode.ApiGaode
}

var ApiGroupApp = new(ApiGroup)
