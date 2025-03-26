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

// 针对菜单列表的排序，作为自定义链接表的功能补充
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
	//查表操作 设定model作为菜单，进行banner和manu连接
	var menus = make([]MenuRes, 0)
	for _, model := range menuList {
		//var banners []Banner  前端粘合的时候发现的问题，会出现nil，前端无法使用，改用make可以调整为[]
		var banners = make([]Banner, 0)
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			//循环查看表观察是否有图，如果没有就不给
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
