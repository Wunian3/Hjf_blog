package api_feedback

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type IDListRequest struct {
	IDList []uint `json:"id_list" binding:"required"` // 需要删除的 ID 列表，必填
}

func (ApiFeedback) FeedBackRemove(c *gin.Context) {
	// 使用 models.RemoveRequest 作为请求结构体，与 ImageDelete 保持一致
	var cr models.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	if len(cr.IDList) == 0 {
		res.OkWithMessage("未指定任何删除项", c)
		return
	}

	var feedbacks []models.FeedbackModel
	// 先根据 IDList 查找记录
	count := global.DB.Find(&feedbacks, cr.IDList).RowsAffected
	if count == 0 {
		// 使用 "记录不存在" 以匹配 ImageDelete 的错误消息风格
		res.FailWithMessage("记录不存在", c)
		return
	}

	// 删除查找到的记录切片，与 ImageDelete 保持一致
	err := global.DB.Delete(&feedbacks).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除反馈失败", c)
		return
	}

	// 返回成功信息，格式与 ImageDelete 保持一致
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 条反馈", count), c)
}
