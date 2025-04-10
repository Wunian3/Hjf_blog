package ctype

import (
	"encoding/json"
)

type Role int

const (
	PermisssionAdmin       Role = 1
	PermisssionUser        Role = 2
	PermisssionVisitor     Role = 3
	PermisssionDisableUser Role = 4
)

func (status Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.String())
}

func (status Role) String() string {
	var str string
	switch status {
	case PermisssionAdmin:
		str = "管理员"
	case PermisssionUser:
		str = "普通用户"
	case PermisssionVisitor:
		str = "游客"
	case PermisssionDisableUser:
		str = "黑名单的家伙"
	default:
		str = "unknown"
	}
	return str
}
