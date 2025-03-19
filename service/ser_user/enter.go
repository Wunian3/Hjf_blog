package ser_user

import (
	"blog_server/service/ser_redis"
	"blog_server/utils/jwts"
	"time"
)

type ServiceUser struct {
}

// Logout 用户注销操作
func (ServiceUser) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return ser_redis.Logout(token, diff)

}
