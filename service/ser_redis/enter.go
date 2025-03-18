package ser_redis

import (
	"blog_server/global"
	"blog_server/utils"
	"time"
)

const prefix = "logout_"

// Logout 只用于注销的redis操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(prefix+token, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
