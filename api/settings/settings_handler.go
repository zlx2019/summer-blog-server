/**
  @author: Zero
  @date: 2023/3/2 23:35:01
  @desc: 系统设置服务

**/

package settings

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"strconv"
	"summer/constant"
	"summer/models/mono"
	"summer/properties"
)

type HandlerSettings struct {
}

type State int

// 操作不同配置的枚举
const (
	Site State = iota + 1
	Jwt
	Email
	QN
	QQ
)

// GetConfig 根据不同的标识,获取不同的配置属性信息
// state = 1 获取站点配置信息
// state = 2 获取jwt配置信息
//...
func GetConfig(ctx *gin.Context) {
	stateQuery := ctx.Query("state")
	if stateQuery == "" {
		mono.FailWithMessage("state未填.", ctx)
		return
	}
	val, err := strconv.Atoi(stateQuery)
	if err != nil {
		mono.FailWithMessage("state参数格式错误.", ctx)
		return
	}
	switch State(val) {
	case Site:
		mono.OkWithData(constant.Config.SiteInfo, ctx)
	case Jwt:
		mono.OkWithData(constant.Config.Jwt, ctx)
	case QN:
		mono.OkWithData(constant.Config.QiNiuOss, ctx)
	case QQ:
		mono.OkWithData(constant.Config.QQ, ctx)
	case Email:
		mono.OkWithData(constant.Config.Email, ctx)
	default:
		mono.FailWithMessage("state 不存在~", ctx)
	}
}

// UpdateConfig 更改网站基本信息
func UpdateConfig(ctx *gin.Context) {
	// 要修改的配置属性对象
	var cnf any
	stateQuery := ctx.Query("state")
	if stateQuery == "" {
		mono.FailWithMessage("state未填.", ctx)
		return
	}
	val, err := strconv.Atoi(stateQuery)
	if err != nil {
		mono.FailWithMessage("state参数格式错误.", ctx)
		return
	}
	// 根据不同的state创建不同的参数绑定实体
	switch State(val) {
	case Site:
		cnf = &properties.SiteInfo{}
	case Jwt:
		cnf = &properties.Jwt{}
	case QN:
		cnf = &properties.Oss{}
	case QQ:
		cnf = &properties.QQ{}
	}
	// 解析参数
	err = ctx.ShouldBindJSON(cnf)
	if err != nil {
		mono.FailWithCode(constant.BAD_REQUEST, ctx)
		return
	}
	// 断言绑定属性的类型
	// 然后分别更新内存中的值和配置文件中的值
	switch c := cnf.(type) {
	case *properties.SiteInfo:
		constant.Config.SiteInfo = *c
	case *properties.Jwt:
		constant.Config.Jwt = *c
	case *properties.Oss:
		constant.Config.QiNiuOss = *c
	case *properties.QQ:
		constant.Config.QQ = *c
	default:
		mono.FailWithMessage("没有对应的State，无法修改", ctx)
		return
	}
	// 更改文件中的配置
	FlushConfigure()
	mono.Ok(ctx)
}

// FlushConfigure 将内存中的配置属性值,更新写回具体的配置文件中。
func FlushConfigure() {
	// 将目前内存中的配置文件属性读取为字节
	bytes, err := yaml.Marshal(constant.Config)
	if err != nil {
		constant.Log.Error("Config Marshal Error:" + err.Error())
	}
	// 写入到本地配置文件
	err = ioutil.WriteFile("settings.yml", bytes, fs.ModePerm)
	if err != nil {
		constant.Log.Error("Config Flush Write Error:" + err.Error())
	}
}
