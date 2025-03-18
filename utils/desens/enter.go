package desens

import "strings"

// 脱敏设置
func DesensitizationEmail(email string) string {
	// 1243624625@qq.com  1****@qq.com

	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

func DesensitizationTel(tel string) string {
	// 17305939725
	// 173****9725
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
