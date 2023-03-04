/**
  @author: Zero
  @date: 2023/3/4 21:07:10
  @desc: 广告表

**/

package models

import "summer/models/custom"

// Advert 广告表结构
type Advert struct {
	custom.BaseModel
	Title  string `json:"title" gorm:"size:32; not null; comment: 广告标题"`
	Link   string `json:"link" gorm:"size:256; comment: 广告跳转链接"`
	Banner string `json:"banner" gorm:"size:256; comment: 广告图片"`
	IsShow bool   `json:"is_show" gorm:"default:false; comment: 是否展示"`
}
