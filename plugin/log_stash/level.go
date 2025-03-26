package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel Level = 1
	InfoLevel  Level = 2
	WarnLevel  Level = 3
	ErrorLevel Level = 4
)

func (status Level) MarshalJson() ([]byte, error) {
	return json.Marshal(status.string())

}
func (status Level) string() string {
	var str string
	switch status {
	case DebugLevel:
		str = "Debug"
	case InfoLevel:
		str = "Info"
	case WarnLevel:
		str = "Warn"
	case ErrorLevel:
		str = "Error"
	default:
		str = "Else"
	}
	return str
}
