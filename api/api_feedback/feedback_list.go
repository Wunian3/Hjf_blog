package api_feedback

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"github.com/gin-gonic/gin"
)

type FeedbackListRequest struct {
	models.PageInf        // 嵌入分页信息结构体
	Keyword        string `form:"keyword"` // 可选：添加关键词查询参数
}

func (ApiFeedback) FeedBackList(c *gin.Context) {
	var cr models.PageInf
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 使用 common.ComList 获取反馈列表
	// 注意：common.ComList 需要能处理 models.FeedbackModel
	list, count, err := common.ComList(models.FeedbackModel{}, common.Option{
		PageInf: cr,   // 传递分页信息
		Debug:   true, // 根据需要开启 Debug
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询反馈列表失败", c)
		return
	}

	res.OkWithList(list, count, c)
}
