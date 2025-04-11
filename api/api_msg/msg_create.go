package api_msg

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type MsgRequest struct {
	//SendUserID uint   `json:"send_user_id" binding:"required"` // 发送人id
	RevUserID uint   `json:"rev_user_id" binding:"required"` // 接收人id
	Content   string `json:"content" binding:"required"`     // 消息内容
}

// MessageCreateView 发布消息
func (ApiMsg) MsgCreate(c *gin.Context) {
	var cr MsgRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var senUser, revUser models.UserModel

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	err = global.DB.Take(&senUser, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("发送人不存在", c)
		return
	}
	err = global.DB.Take(&revUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMessage("接收人不存在", c)
		return
	}

	err = global.DB.Create(&models.MsgModel{
		SendUserID:       senUser.ID,
		SendUserNickName: senUser.NickName,
		SendUserAvatar:   senUser.Avatar,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  revUser.NickName,
		RevUserAvatar:    revUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("消息发送失败", c)
		return
	}
	res.OkWithMessage("消息发送成功", c)
	return
}
