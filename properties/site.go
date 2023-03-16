/**
  @author: Zero
  @date: 2023/3/8 16:58:56
  @desc: 网站站点基础信息配置

**/

package properties

// SiteInfo 网站信息
type SiteInfo struct {
	Created string `json:"created" yaml:"created"` //建站日期
	BeiAn   string `json:"beiAn" yaml:"bei_an"`    //备案号
	Title   string `json:"title" yaml:"title"`     //网站标题
	Version string `json:"version" yaml:"version"` //版本号
	Email   string `json:"email" yaml:"email"`     //邮箱
	Name    string `json:"name" yaml:"name"`       //站长名称
	Work    string `json:"work" yaml:"work"`       //工作岗位
	Addr    string `json:"addr" yaml:"addr"`       //所在城市
	Slogan  string `json:"slogan" yaml:"slogan"`   //
	Web     string `json:"web" yaml:"web"`         //
	Gitee   string `json:"gitee" yaml:"gitee"`     //码云地址
	Github  string `json:"github" yaml:"github"`   //github地址
}
