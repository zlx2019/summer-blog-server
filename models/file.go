/**
  @author: Zero
  @date: 2023/3/4 19:28:21
  @desc: 图片文件表

**/

package models

import "summer/models/custom"

// File 文件表结构
type File struct {
	custom.BaseModel
	Url  string `json:"url" gorm:"size:256; not null; comment: 文件url"`
	Hash string `json:"hash" gorm:"size:500; not null; comment: 文件hash值,用来判断是否重复"`
	Name string `json:"name" gorm:"size:38; not null; comment: 文件名"`
}
