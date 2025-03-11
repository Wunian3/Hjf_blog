package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ     SignStatus = 1
	SignGitee  SignStatus = 2
	SignEmail  SignStatus = 3
	SignGithub SignStatus = 4
)

func (status SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.string())
}

func (status SignStatus) string() string {
	var str string
	switch status {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "Gitee"
	case SignEmail:
		str = "Email"
	case SignGithub:
		str = "Github"
	default:
		str = "unknown"
	}
	return str
}
