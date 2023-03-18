/**
  @author: Zero
  @date: 2023/3/18 17:13:34
  @desc: 错误处理中间件

**/

package middleware

import (
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"summer/models/mono"
)
import . "summer/constant"

// Recover Gin请求全局错误处理
func Recover(ctx *gin.Context) {
	defer func() {
		// panic捕获
		if r := recover(); r != nil {
			message := errToString(r)
			//打印错误日志
			Log.Error(message)
			// 打印详细的错误栈信息
			Log.Error(string(debug.Stack()))
			// 响应错误信息
			mono.FailWithMessage(message, ctx)
			// 终止请求
			ctx.Abort()
		}
	}()
	// 放行
	ctx.Next()
}

// 获取error的错误消息
func errToString(r any) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
