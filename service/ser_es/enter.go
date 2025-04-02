package ser_es

import (
	"blog_server/models"
	"github.com/olivere/elastic/v7"
)

type Option struct {
	models.PageInf
	Fields []string
	Tag    string
	Query  elastic.Query
}

func (o *Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit
}
