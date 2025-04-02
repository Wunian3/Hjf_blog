package api_msg

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	_ "blog_server/models/res"
	"blog_server/service/common"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type MessageUserRecordByMeRequest struct {
	models.PageInf
	UserID uint `json:"userID" form:"userID" binding:"required"`
}

// MessageUserRecordByMe 我与某个用户的聊天列表
// @Tags 消息管理
// @Summary 我与某个用户的聊天列表
// @Description 我与某个用户的聊天列表
// @Router /api/message_users/record/me [get]
// @Param token header string  true  "token"
// @Param data query MessageUserRecordByMeRequest  true  "参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MessageModel]}
func (m ApiMsg) MessageUserRecordByMe(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr MessageUserRecordByMeRequest
	c.ShouldBindQuery(&cr)

	cr.Sort = "created_at asc"
	list, count, _ := common.ComList(models.MsgModel{}, common.Option{
		PageInf: cr.PageInf,
		Where:   global.DB.Where("(send_user_id = ? and rev_user_id = ? ) or ( rev_user_id = ? and send_user_id = ? )", claims.UserID, cr.UserID, claims.UserID, cr.UserID),
	})

	res.OkWithList(list, count, c)
}
