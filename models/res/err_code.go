package res

type ErrorCode int

const (
	//系统问题的代码编号 系统、参数
	SettingsError ErrorCode = 1001
	ArgumentError ErrorCode = 1002
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "HJF,你的BLOG系统又错误了- -，自己再研究一下吧仔细看看debug",
		ArgumentError: "HJF,你的BLOG参数又错误了- -，自己再研究一下吧仔细看看debug",
	}
)
