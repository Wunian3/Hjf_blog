package flag

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/plugin/log_stash"
	"blog_server/plugin/log_stash_v2"
)

func Makemigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MsgModel{},
			&models.AdtModel{},
			&models.UserModel{},
			&models.CommentModel{},
			//&models.ArticleModel{},
			&models.UserCollectModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FadeBackModel{},
			&models.LogDataMd{},
			&models.ChatModel{},
			&log_stash.LogStashModel{},
			&log_stash_v2.LogModel{},
			&models.FeedbackModel{},
			&models.UserScopeModel{},
			&models.AutoReplyModel{},
			&models.BigModelRoleModel{},    //角色表
			&models.BigModelTagModel{},     //标签表
			&models.BigModelChatModel{},    //对话表
			&models.BigModelSessionModel{}, //会话表
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
