/**
  @author: Zero
  @date: 2023/3/2 23:35:09
  @desc: 系统设置服务路由

**/

package settings

import (
	"github.com/gin-gonic/gin"
	"summer/router"
)

type RouterSettings struct {
}

func init() {
	// 注册Settings的路由
	router.RegisterRouter(&RouterSettings{})
}

func (*RouterSettings) Route(engine *gin.Engine) {
	group := engine.Group("/settings")
	group.GET("/", GetConfig)
	group.PUT("/", UpdateConfig)
}
