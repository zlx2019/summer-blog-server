/**
  @author: Zero
  @date: 2023/3/3 00:47:29
  @desc: 分页统一结果

**/

package page

// Page 分页响应对象
type Page[T any] struct {
	// 响应码
	Code int `json:"code"`
	// 响应序列数据
	List []T `json:"list"`
	// 响应消息
	Message string `json:"message"`
}
