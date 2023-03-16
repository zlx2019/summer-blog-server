/**
  @author: Zero
  @date: 2023/3/10 12:04:34
  @desc: 令牌配置属性

**/

package properties

// Jwt 令牌配置
type Jwt struct {
	// 秘钥
	Secret string `json:"secret" yaml:"secret"`
	// 有效期. 单位:小时
	Expires int `json:"expires" yaml:"expires"`
	// 颁发者
	Issuer string `json:"issuer" yaml:"issuer"`
}
