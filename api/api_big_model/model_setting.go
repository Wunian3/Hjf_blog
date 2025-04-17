package api_big_model

import (
	"blog_server/conf"
	"blog_server/global"
	"blog_server/models"
	"blog_server/models/res"
	"blog_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

const docsPath = "uploads/docs"

type ModelSetting struct {
	conf.Setting
	Help string `json:"help"`
}

// ModelSettingView 获取大模型配置
func (ApiBigModel) ModelSettingView(c *gin.Context) {
	token := c.GetHeader("token")
	var roleID int
	customClaims, err := jwts.ParseToken(token)
	if err == nil && customClaims != nil {
		roleID = customClaims.Role
	}
	if roleID == models.AdminRole {
		// 判断用户是不是管理员，管理员就展示所有信息
		ms := ModelSetting{
			Setting: global.Config.BigModel.Setting,
		}

		if ms.Name != "" {
			filePath := path.Join(docsPath, fmt.Sprintf("%s.md", ms.Name))
			byteData, err := os.ReadFile(filePath)
			if err == nil {
				ms.Help = string(byteData)
			}
		}

		res.OkWithData(ms, c)
		return
	}

	res.OkWithData(ModelSetting{
		Setting: conf.Setting{
			Enable: global.Config.BigModel.Setting.Enable,
			Title:  global.Config.BigModel.Setting.Title,
			Slogan: global.Config.BigModel.Setting.Slogan,
		},
	}, c)
	return
}
