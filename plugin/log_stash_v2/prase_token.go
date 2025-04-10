package log_stash_v2

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayLoad struct {
	UserName string `json:"user_name"` // 用户名
	NickName string `json:"nick_name"` // 昵称
	Role     int    `json:"role"`      // 权限  1 管理员  2 普通用户  3 游客
	UserID   uint   `json:"user_id"`   // 用户id
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

func parseToken(token string) (jwtPayload *JwtPayLoad) {
	Token, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if Token == nil || Token.Claims == nil {
		return nil
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		return nil
	}
	return &claims.JwtPayLoad
}
