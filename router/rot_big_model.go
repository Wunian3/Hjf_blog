package router

import (
	"blog_server/api"
	"blog_server/middle"
)

func (router RouterGroup) RouterBigModel() {
	apiuse := api.ApiGroupApp.ApiBigModel
	//配置管理
	{
		router.GET("big_model/usable", middle.JwtAdmin(), apiuse.ModelUsableListView)
		router.GET("big_model/setting", apiuse.ModelSettingView)
		router.PUT("big_model/setting", middle.JwtAdmin(), apiuse.ModelSettingUpdateView)
		router.GET("big_model/session_setting", middle.JwtAdmin(), apiuse.ModelSessionSettingView)
		router.PUT("big_model/session_setting", middle.JwtAdmin(), apiuse.ModelSessionSettingUpdateView)
		//自动回复
		router.PUT("big_model/auto_reply", middle.JwtAdmin(), apiuse.AutoReplyUpdateView)
		router.GET("big_model/auto_reply", middle.JwtAdmin(), apiuse.AutoReplyListView)
		router.DELETE("big_model/auto_reply", middle.JwtAdmin(), apiuse.AutoReplyRemoveView)
	}
	//角色管理
	{
		router.PUT("big_model/tags", middle.JwtAdmin(), apiuse.TagUpdateView)
		router.GET("big_model/tags/options", middle.JwtAdmin(), apiuse.TagOptionsListView) //角色标签分页
		router.GET("big_model/tags", middle.JwtAdmin(), apiuse.TagListView)
		router.DELETE("big_model/tags", middle.JwtAdmin(), apiuse.TagRemoveView)

		router.PUT("big_model/roles", middle.JwtAdmin(), apiuse.RoleUpdateView)
		router.POST("big_model/roles", middle.JwtAdmin(), apiuse.RoleCreateView)
		router.GET("big_model/roles", middle.JwtAdmin(), apiuse.RoleListView)

		router.GET("big_model/roles_history", middle.JwtAuth(), apiuse.RoleUserHistoryListView) //用户角色历史列表
		router.GET("big_model/roles/:id", apiuse.RoleDetailView)                                //角色详情
		router.DELETE("big_model/roles", middle.JwtAuth(), apiuse.RoleRemoveView)               //角色删除
		router.GET("big_model/square", apiuse.TagRoleListView)                                  //角色广场

		router.GET("big_model/role_sessions", middle.JwtAuth(), apiuse.RoleSessionsView) //角色会话列表

		router.GET("big_model/icons/options", apiuse.IconsView) //角色头像

	}
	//会话管理
	{
		router.POST("big_model/session", middle.JwtAuth(), apiuse.SessionCreateView)           //用户创建会话
		router.GET("big_model/session", middle.JwtAdmin(), apiuse.SessionListView)             //会话列表
		router.PUT("big_model/session", middle.JwtAuth(), apiuse.SessionUserUpdateNameView)    //修改会话名称
		router.DELETE("big_model/session/:id", middle.JwtAuth(), apiuse.SessionUserRemoveView) //用户删除会话名称
		router.DELETE("big_model/session", middle.JwtAdmin(), apiuse.SessionRemoveView)        //管理员删除会话名称
	}
	//对话管理
	{
		router.GET("big_model/chat_sse", apiuse.ChatCreateView)                          //用户创建对话
		router.GET("big_model/chat", middle.JwtAuth(), apiuse.ChatListView)              //对话列表
		router.DELETE("big_model/chat/:id", middle.JwtAuth(), apiuse.ChatUserRemoveView) //用户删除对话
		router.DELETE("big_model/chat", middle.JwtAdmin(), apiuse.ChatRemoveView)        //管理员删除对话
	}
	//积分管理
	router.GET("big_model/user_scope_enable", middle.JwtAuth(), apiuse.UserScopeEnableView)
	router.POST("big_model/user_scope", middle.JwtAuth(), apiuse.UserScopeView)

}
