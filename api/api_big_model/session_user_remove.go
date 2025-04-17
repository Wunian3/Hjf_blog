package api_big_model

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// SessionUserRemoveView 用户删除会话
func (ApiBigModel) SessionUserRemoveView(c *gin.Context) {
	var cr models.IDRequest
	// control request
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}

	// 找会话
	var session models.BigModelSessionModel
	err = global.DB.Preload("ChatList").Take(&session, cr.ID).Error
	if err != nil {
		res.FailWithMessage("会话不存在", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	if session.UserID != claims.UserID {
		res.FailWithMessage("会话鉴权失败", c)
		return
	}

	// 话对话记录删除
	if len(session.ChatList) > 0 {
		err = global.DB.Delete(&session.ChatList).Error
		if err != nil {
			logrus.Error(err)
		} else {
			logrus.Infof("删除关联对话 %d 个", len(session.ChatList))
		}
	}

	// 删除会话
	err = global.DB.Delete(&session).Error
	if err != nil {
		logrus.Error(err)
		res.FailWithMessage("会话删除失败", c)
		return
	}
	res.OkWithMessage("会话删除成功", c)
}
