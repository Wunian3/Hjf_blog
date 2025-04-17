package api_big_model

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// TagRemoveView 标签删除
func (ApiBigModel) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}
	var list []models.BigModelTagModel
	count := global.DB.Preload("Roles").Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("记录不存在", c)
		return
	}

	if len(list) > 0 {
		// 先把引用的记录删除
		for _, i2 := range list {
			global.DB.Model(&i2).Association("Roles").Delete(i2.Roles)
		}
		err = global.DB.Delete(&list).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMessage("删除标签失败", c)
			return
		}
		logrus.Infof("删除角色标签 %d 个", len(list))
	}

	res.OkWithMessage("删除角色标签成功", c)
}
