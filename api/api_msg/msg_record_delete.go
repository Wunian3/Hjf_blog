package api_msg

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// MessageRecordDelete 删除用户的消息记录
// @Tags 消息管理
// @Summary 删除用户的消息记录
// @Description 删除用户的消息记录
// @Router /api/message_users [delete]
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest   true  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{]}
func (ApiMsg) MessageRecordDelete(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var messageList []models.MsgModel
	global.DB.Find(&messageList, cr.IDList)

	if len(messageList) > 0 {
		err = global.DB.Delete(&messageList).Error
		if err != nil {
			res.FailWithMessage("消息记录删除失败", c)
			return
		}
	}

	res.OkWithMessage(fmt.Sprintf("共删除记录%d条", len(messageList)), c)
}
