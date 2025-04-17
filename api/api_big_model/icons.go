package api_big_model

import (
	"blog_server/models"
	"blog_server/models/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

func (ApiBigModel) IconsView(c *gin.Context) {
	dir, err := os.ReadDir("uploads/role_icons")
	if err != nil {
		logrus.Error(err)
		res.FailWithMessage("目录不存在", c)
		return
	}
	var list []models.Options[string]
	for _, entry := range dir {
		key := "/" + path.Join("uploads/role_icons", entry.Name())
		list = append(list, models.Options[string]{
			Label: key,
			Value: key,
		})
	}

	res.OkWithData(list, c)
}
