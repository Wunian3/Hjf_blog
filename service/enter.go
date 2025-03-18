package service

import (
	"blog_server/service/common/ser_img"
	"blog_server/service/ser_user"
)

type ServiceGroup struct {
	ServiceImage ser_img.ServiceImage
	ServiceUser  ser_user.ServiceUser
}

var ServiceGroupApp = new(ServiceGroup)
