package api_big_model

import (
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/service/common"
	"blog_server/utils/jwts"
	"github.com/gin-gonic/gin"
	"time"
)

type RoleSessionsRequest struct {
	models.PageInf
	RoleID uint `json:"roleID" form:"roleID" binding:"required"`
}

type RoleSessionResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

// RoleSessionsView 角色会话列表
func (ApiBigModel) RoleSessionsView(c *gin.Context) {
	var cr RoleSessionsRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithValidError(err, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	_list, count, _ := common.ComList(models.BigModelSessionModel{UserID: claims.UserID, RoleID: cr.RoleID}, common.Option{
		PageInf: cr.PageInf,
		Likes:   []string{"name"},
	})
	var list = make([]RoleSessionResponse, 0)
	for _, model := range _list {
		list = append(list, RoleSessionResponse{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			Name:      model.Name,
		})
	}
	res.OkWithList(list, count, c)
}
