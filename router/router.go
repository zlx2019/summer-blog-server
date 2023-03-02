/**
  @author: Zero
  @date: 2023/3/2 20:55:25
  @desc: Gin路由初始化
**/

package router

import "github.com/gin-gonic/gin"

// Router 路由抽象接口.
// 每个业务模块各自实现该接口进行路由绑定。
type Router interface {
	// Route 路由方法
	Route(e *gin.Engine)
}

// 要绑定的业务API路由
var routers []Router

// RegisterRouter 注册业务API路由
func RegisterRouter(route ...Router) {
	routers = append(routers, route...)
}

// InitRouters 初始化路由 将所有的业务API,统一绑定到Gin路由
func InitRouters(engine *gin.Engine) {
	for _, router := range routers {
		router.Route(engine)
	}
}
