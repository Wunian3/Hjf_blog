package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayLoad jwt中预备数据
type JwtPayLoad struct {
	Username string `json:"username"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"` // 权限 管理员、普通、用户游客
	UserID   uint   `json:"user_id"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

var MySecret []byte
