/**
  @author: Zero
  @date: 2023/3/3 11:58:03
  @desc: 系统响应码和响应消息
**/

package constant

// 各个业务块的错误码
// 10xx 为设置模块错误码
// 20xx 为用户模块错误码
const (
	// SUCCESS 请求成功 ERROR 请求失败。
	// 础响应码
	SUCCESS int = iota
	ERROR

	SettingsErr  = 1001
	UserNotFound = 2001
)

// ResultMap 响应码与响应消息的映射
var resultMap = map[int]string{
	SUCCESS:      "请求成功",
	ERROR:        "请求失败",
	SettingsErr:  "设置模块错误",
	UserNotFound: "用户不存在",
}

// Message 根据响应码获取响应消息
func Message(code int) string {
	if message, ok := resultMap[code]; ok {
		return message
	}
	return resultMap[ERROR]
}
