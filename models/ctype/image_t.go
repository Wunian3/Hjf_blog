package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1
	QiNiu ImageType = 2
)

func (status ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.string())
}

func (status ImageType) string() string {
	var str string
	switch status {
	case Local:
		str = "本地"
	case QiNiu:
		str = "七牛云"
	default:
		str = "未知"
	}
	return str
}
