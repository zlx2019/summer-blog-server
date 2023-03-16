/**
  @author: Zero
  @date: 2023/3/10 12:08:53
  @desc: 云存储配置属性

**/

package properties

// Oss 七牛OSS
type Oss struct {
	AccessKey string  `json:"accessKey" yaml:"access_key"` //accessKey
	SecretKey string  `json:"secretKey" yaml:"secret_key"` //秘钥
	Bucket    string  `json:"bucket"    yaml:"bucket"`     //存储桶的名字
	Domain    string  `json:"domain"    yaml:"domain"`     //访问域名
	Zone      string  `json:"zone"      yaml:"zone"`       //存储地区
	Size      float64 `json:"size"      yaml:"size"`       //存储的文件最大限制。单位mb
}
