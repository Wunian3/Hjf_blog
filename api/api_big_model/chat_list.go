package api_big_model

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type ChatListRequest struct {
	SessionID uint `json:"sessionID" form:"sessionID" binding:"required"`
	models.PageInf
}
type ChatListResponse struct {
	models.MODEL
	UserContent string `json:"userContent"` // 用户聊天内容
	UserAvatar  string `json:"userAvatar"`  // 用户头像
	BotContent  string `json:"botContent"`  // AI的聊天内容
	BotAvatar   string `json:"botAvatar"`   // AI的头像
	Status      bool   `json:"status"`
}

// ChatListView 对话列表
func (ApiBigModel) ChatListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr ChatListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}
	var session models.BigModelSessionModel
	err = global.DB.Take(&session, cr.SessionID).Error
	if err != nil {
		res.FailWithMessage("会话id错误", c)
		return
	}
	if claims.Role != models.AdminRole {
		// 要去验证这个会话是不是当前用户创建的
		if claims.UserID != session.UserID {
			res.FailWithMessage("会话鉴权失败", c)
			return
		}
	}

	_list, count, _ := common.ComList(models.BigModelChatModel{SessionID: cr.SessionID}, common.Option{
		PageInf: cr.PageInf,
		Preload: []string{"RoleModel", "UserModel"},
	})
	var list = make([]ChatListResponse, 0)
	for _, model := range _list {
		list = append(list, ChatListResponse{
			MODEL:       model.MODEL,
			UserContent: model.Content,
			UserAvatar:  model.UserModel.Avatar,
			BotContent:  model.BotContent,
			BotAvatar:   model.RoleModel.Icon,
			Status:      model.Status,
		})
	}

	res.OkWithList(list, count, c)
}
