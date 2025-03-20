package flag

import (
	"blog_server/core"
	"blog_server/global"
	sys_flag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string //分管理员和普通用户
	ES   string //-es create/delete
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "数据库初始化")
	user := sys_flag.String("u", "", "创建用户")
	es := sys_flag.String("es", "", "es_operation")

	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	if option.ES == "create" {
		//es_connect
		global.ESClient = core.EsConnect()
		EsCreateIndex()
	}
	//sys_flag.Usage()
}
