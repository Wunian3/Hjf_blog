package api_msg

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type MsgRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询的用户id"`
}

func (ApiMsg) MsgRecord(c *gin.Context) {
	var cr MsgRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var _messageList []models.MsgModel
	var messageList = make([]models.MsgModel, 0)
	global.DB.Order("created_at asc").
		Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range _messageList {
		if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
			messageList = append(messageList, model)
		}
	}
	// todo 点开消息，里面的每一条消息，都从未读变成已读
	res.OkWithData(messageList, c)
}
