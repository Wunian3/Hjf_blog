package common

import (
	"blog_server/global"
	"blog_server/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInf
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}
	query := DB.Model(&model)

	count = DB.Select("id").Find(&list).RowsAffected

	query = DB.Where(model)
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
