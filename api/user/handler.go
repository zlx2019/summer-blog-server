/**
  @author: Zero
  @date: 2023/3/2 21:02:58
  @desc: User服务API 处理函数

**/

package user

import "github.com/gin-gonic/gin"

type HandlerUser struct {
}

// Hello TestAPI
func (*HandlerUser) Hello(ctx *gin.Context) {
	ctx.JSON(200, "Hello")
}
