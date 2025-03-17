package api_menu

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuRes struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (ApiMenu) MenuList(c *gin.Context) {
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	//查表操作
	var menus []MenuRes
	for _, model := range menuList {
		var banners []Banner
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuRes{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
	return
}
