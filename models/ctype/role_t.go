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
	return json.Marshal(status.string())
}

func (status Role) string() string {
	var str string
	switch status {
	case PermisssionAdmin:
		str = "permisssion_admin"
	case PermisssionUser:
		str = "permisssion_user"
	case PermisssionVisitor:
		str = "permisssion_visitor"
	case PermisssionDisableUser:
		str = "permisssion_disable_user"
	default:
		str = "permisssion_unknown"
	}
	return str
}
