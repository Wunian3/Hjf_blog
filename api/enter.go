package api

import "blog_server/api/api_settings"

type ApiGroup struct {
	ApiSettings api_settings.ApiSettings
}

var ApiGroupApp = new(ApiGroup)
