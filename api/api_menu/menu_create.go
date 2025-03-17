package api_menu

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/ctype"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单别名" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`
	BannerTime    int         `json:"banner_time" structs:"banner_time"` // 切换的时间
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"`
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`
}

func (ApiMenu) MenuCreate(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var menulist []models.MenuModel
	count := global.DB.Find(&menulist, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("重复了，请重新上传", c)
		return
	}
	// banner需要数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}

	err = global.DB.Create(&menuModel).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OkWithMessage("菜单添加成功", c)
		return
	}

	var menuBannerList []models.MenuBannerModel

	for _, sort := range cr.ImageSortList {
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单图片连接失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)
}
