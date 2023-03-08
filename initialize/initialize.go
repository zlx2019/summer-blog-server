/**
  @author: Zero
  @date: 2023/3/2 15:09:29
  @desc: 统一初始化组件配置

**/

package initialize

import _ "summer/api"

func init() {
	// 初始化配置文件
	InitLoadConfigure()
	// 初始化日志组件
	InitLogger()
	// 初始化数据访问组件
	initDataSourceConfigure()
}
