/**
  @author: Zero
  @date: 2023/3/3 19:36:27
  @desc: 文章标签

**/

package models

import "summer/models/custom"

// Tag 标签表结构
type Tag struct {
	custom.BaseModel
	Name string `json:"name" gorm:"size:30; not null; comment: 标签名称"`
}
