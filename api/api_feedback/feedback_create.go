package api_feedback

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

type FeedbackCreateRequest struct {
	Email   string `json:"email" binding:"required,email" msg:"请输入正确的邮箱"`              // 邮箱，必填且需为邮箱格式
	Content string `json:"content" binding:"required,max=256" msg:"反馈内容不能为空且长度不超过256"` // 反馈内容，必填，最大长度256
}

func (ApiFeedback) FeedBackCreate(c *gin.Context) {
	var cr FeedbackCreateRequest
	// 绑定并校验请求体 JSON 数据
	if err := c.ShouldBindJSON(&cr); err != nil {
		// 如果校验失败，使用 res.FailWithError 返回详细错误信息
		res.FailWithError(err, &cr, c)
		return
	}

	// 创建 FeedbackModel 实例
	feedback := models.FeedbackModel{
		Email:   cr.Email,
		Content: cr.Content,
		// MODEL 中的 CreatedAt 和 UpdatedAt 由 GORM 自动填充
	}

	// 将反馈数据存入数据库
	err := global.DB.Create(&feedback).Error
	if err != nil {
		// 数据库操作失败，记录日志并返回错误信息
		global.Log.Error(err)
		res.FailWithMessage("反馈提交失败", c)
		return
	}

	// 返回成功信息
	res.OkWithMessage("反馈提交成功", c)
}
