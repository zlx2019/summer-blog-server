package utils

// StrIsBlank 字符串是否为空
func StrIsBlank(str string) bool {
	return str == ""
}

// StrNotBlank 字符串是否不为空
func StrNotBlank(str string) bool {
	return !StrIsBlank(str)
}
