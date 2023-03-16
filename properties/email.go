/**
  @author: Zero
  @date: 2023/3/10 12:04:54
  @desc: 邮箱配置属性

**/

package properties

// Email 第三方邮箱配置
type Email struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Sender   string `json:"sender" yaml:"sender"`
	UseSsl   bool   `json:"use_ssl" yaml:"use_ssl"`
	UserTls  bool   `json:"user_tls" yaml:"user_tls"`
}
