package api_menu

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (ApiMenu) MenuUpdate(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	// 选择了banner，那就添加并对第三张表进行更动
	if len(cr.ImageSortList) > 0 {

		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}

	res.OkWithMessage("修改菜单成功", c)

}
