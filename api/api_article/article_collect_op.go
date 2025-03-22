package api_article

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/ser_es"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

// ArticleCollectOP 用户收藏文章，或取消收藏

func (ApiArticle) ArticleCollectOP(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	model, err := ser_es.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}

	var coll models.UserCollectModel
	err = global.DB.Take(&coll, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	var num = -1
	if err != nil {
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})
		num = 1
	}
	// 取消收藏
	global.DB.Delete(&coll)

	err = ser_es.ArticleUp(cr.ID, map[string]any{
		"collects_count": model.CollectsCount + num,
	})
	if num == 1 {
		res.OkWithMessage("收藏文章成功", c)
	} else {
		res.OkWithMessage("取消收藏成功", c)
	}
}
