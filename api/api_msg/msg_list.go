package api_msg

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
	"time"
)

type Msg struct {
	SendUserID       uint      `json:"send_user_id"` // 发送人id
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"` // 接收人id
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`       // 消息内容
	CreatedAt        time.Time `json:"created_at"`    // 最新的消息时间
	MessageCount     int       `json:"message_count"` // 消息条数
}

type MessageGroup map[uint]*Msg

func (ApiMsg) MsgList(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var messageGroup = MessageGroup{}
	var messageList []models.MsgModel
	var messages []Msg

	global.DB.Order("created_at asc").
		Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range messageList {
		// 判断是一个组的条件send_user_id 和 rev_user_id
		// 1 2  2 1
		message := Msg{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevUserNickName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}
		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	res.OkWithData(messages, c)
	return
}
