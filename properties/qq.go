/**
  @author: Zero
  @date: 2023/3/10 12:04:46
  @desc: qq聊天配置属性

**/

package properties

// QQ 第三方登录配置
type QQ struct {
	AppId    string `json:"appId" yaml:"app_id"`      //应用ID
	Key      string `json:"key" yaml:"key"`           //应用秘钥
	Redirect string `json:"redirect" yaml:"redirect"` //登录后跳转的地址
}
