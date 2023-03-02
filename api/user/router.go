/**
  @author: Zero
  @date: 2023/3/2 21:03:08
  @desc: User服务路由

**/

package user

import (
	"github.com/gin-gonic/gin"
	"summer/router"
)

type RouterUser struct {
}

func init() {
	// 注册User的路由
	router.RegisterRouter(&RouterUser{})
}

// Route User服务路由定义
func (u *RouterUser) Route(engine *gin.Engine) {
	handler := &HandlerUser{}
	engine.GET("/hello", handler.Hello)
}
