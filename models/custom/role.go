/**
  @author: Zero
  @date: 2023/3/3 16:25:24
  @desc: 用户角色

**/

package custom

import "encoding/json"

// Role 用户角色类型
type Role int

// 角色列表 1: 管理员 2:普通用户 3:游客
const (
	ADMIN = iota + 1
	USER
	VISITOR
)

// MarshalJSON 重写Role结构的JSON序列化方法
func (role *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(role.String())
}

// 重写Role结构String序列化方法
func (role *Role) String() string {
	return ParseRole(*role)
}

// ParseRole 根据角色标识,获取角色名称
func ParseRole(role Role) string {
	switch role {
	case ADMIN:
		return "管理员"
	case USER:
		return "用户"
	case VISITOR:
		return "游客"
	default:
		return "其他"
	}
}
