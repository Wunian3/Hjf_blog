package api_menu

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type MenuNameRes struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (ApiMenu) MenuNameList(c *gin.Context) {
	var menuNameList []MenuNameRes
	global.DB.Model(models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)
	res.OkWithData(menuNameList, c)
}
