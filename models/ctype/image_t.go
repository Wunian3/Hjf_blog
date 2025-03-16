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
		str = "local"
	case QiNiu:
		str = "qiniu"
	default:
		str = "unknown"
	}
	return str
}
