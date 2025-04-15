package api_menu

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (ApiMenu) MenuDetailByPath(c *gin.Context) {
	path := c.Query("path") // 改为获取query参数
	if path == "" {
		res.FailWithMessage("path参数不能为空", c)
		return
	}

	var menuModel models.MenuModel
	// 根据path查询菜单
	err := global.DB.Where("path = ?", path).First(&menuModel).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 查询关联的横幅
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").
		Order("sort desc").
		Find(&menuBanners, "menu_id = ?", menuModel.ID)

	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}

	menuResponse := MenuRes{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
}
