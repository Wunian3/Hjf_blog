package api_msg

import (
	_ "blog_server/models/res"
	"blog_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (m ApiMsg) MessageUserListByMe(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	c.Request.URL.RawQuery = fmt.Sprintf("userID=%d", claims.UserID)
	m.MessageUserListByUser(c)

}
