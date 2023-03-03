/**
  @author: Zero
  @date: 2023/3/3 16:44:04
  @desc: 用户注册来源

**/

package custom

import "encoding/json"

// RegisterManner 注册来源
type RegisterManner int

// 注册来源 1: QQ注册 2: 码云注册 3:github注册 4:邮箱注册
const (
	QQ = iota + 1
	GITEE
	GITHUB
	EMAIL
)

// MarshalJSON 重写RegisterManner结构的JSON序列化方法
func (m *RegisterManner) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

// 重写RegisterManner结构String序列化方法
func (m *RegisterManner) String() string {
	return ParseRegisterManner(*m)
}

// ParseRegisterManner 根据角色标识,获取角色名称
func ParseRegisterManner(m RegisterManner) string {
	switch m {
	case QQ:
		return "QQ"
	case GITEE:
		return "Gitee"
	case GITHUB:
		return "Github"
	case EMAIL:
		return "Email"
	default:
		return "其他"
	}
}
