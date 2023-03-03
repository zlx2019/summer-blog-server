/**
  @author: Zero
  @date: 2023/3/3 00:43:14
  @desc: 统一结果响应

**/

package mono

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "summer/constant"
)

// Mono 单数据响应对象
type Mono struct {
	// 响应码
	Code int `json:"code"`
	// 响应数据
	Data any `json:"data"`
	// 响应消息
	Message string `json:"message"`
}

// Just 构建响应对象
func Just(code int, data any, message string) *Mono {
	return &Mono{code, data, message}
}

// Ok 响应成功
func Ok(ctx *gin.Context) {
	OkWithMessage(Message(SUCCESS), ctx)
}

// OkMono 响应成功,并且返回Mono
func OkMono(m *Mono, ctx *gin.Context) {
	Result(m.Code, m.Data, m.Message, ctx)
}

// OkWithMessage 响应成功,返回响应消息
func OkWithMessage(message string, ctx *gin.Context) {
	Result(SUCCESS, nil, message, ctx)
}

// OkWithData 响应成功,返回数据
func OkWithData(data any, ctx *gin.Context) {
	Result(SUCCESS, data, Message(SUCCESS), ctx)
}

// Fail 响应失败
func Fail(ctx *gin.Context) {
	FailWithCode(ERROR, ctx)
}

// FailWithMessage 响应失败
func FailWithMessage(message string, ctx *gin.Context) {
	FailWithCodeMessage(ERROR, message, ctx)
}

// FailWithCode 响应失败,返回响应码
func FailWithCode(code int, ctx *gin.Context) {
	FailWithCodeMessage(code, Message(code), ctx)
}

// FailMono 响应失败,返回Mono
func FailMono(m *Mono, ctx *gin.Context) {
	FailWithCodeMessage(m.Code, m.Message, ctx)
}

// FailWithCodeMessage 响应失败,返回错误码与错误消息
func FailWithCodeMessage(code int, message string, ctx *gin.Context) {
	Result(code, nil, message, ctx)
}

// Result 将结果响应回客户端
func Result(code int, data any, message string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Just(code, data, message))
}
