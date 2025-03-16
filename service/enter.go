package service

import "blog_server/service/common/ser_img"

type ServiceGroup struct {
	ServiceImage ser_img.ServiceImage
}

var ServiceGroupApp = new(ServiceGroup)
