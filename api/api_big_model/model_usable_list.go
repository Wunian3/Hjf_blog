package api_big_model

import (
	"blog_server/global"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
)

func (ApiBigModel) ModelUsableListView(c *gin.Context) {
	res.OkWithData(global.Config.BigModel.ModelList, c)
	return
}
