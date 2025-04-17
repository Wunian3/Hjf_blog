package api_big_model

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ChatUserRemoveView 用户删除对话
func (ApiBigModel) ChatUserRemoveView(c *gin.Context) {
	var cr models.IDRequest
	// control request
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}

	// 找会话
	var chat models.BigModelChatModel
	err = global.DB.Take(&chat, cr.ID).Error
	if err != nil {
		res.FailWithMessage("对话不存在", c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	if chat.UserID != claims.UserID {
		res.FailWithMessage("对话鉴权失败", c)
		return
	}
	// 删除会话
	err = global.DB.Delete(&chat).Error
	if err != nil {
		logrus.Error(err)
		res.FailWithMessage("对话删除失败", c)
		return
	}
	res.OkWithMessage("对话删除成功", c)
}
