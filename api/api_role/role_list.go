package api_role

import (
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type OptionResponse struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

func (ApiRole) RoleIDList(c *gin.Context) {
	res.OkWithData([]OptionResponse{
		{"管理员", 1},
		{"普通用户", 2},
		{"游客", 3},
	}, c)
}
