package api_menu

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
	"strings"
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

// MenuList 菜单列表
// @Tags 菜单管理
// @Summary 菜单列表
// @Description 菜单列表
// @Router /api/menus [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListRes[MenuRes]}
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
			path := banner.BannerModel.Path
			// 如果是本地路径（以 "uploads/" 开头）且没有前导斜杠，则添加 "/"
			if strings.HasPrefix(path, "uploads/") && !strings.HasPrefix(path, "/") {
				path = "/" + path
			}
			// 对于 Qiniu 路径，保持不变
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: path,
			})
		}
		menus = append(menus, MenuRes{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithList(menus, int64(len(menus)), c)
	return
}
