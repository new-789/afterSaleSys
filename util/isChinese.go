package util

import "regexp"

// IsChinese 判断是否为中文
func IsChinese(str string) bool {
	reg := regexp.MustCompile("")
	return reg.MatchString(str)
}
