/**
  @author: Zero
  @date: 2023/3/4 21:03:06
  @desc: 用户反馈表

**/

package models

import "summer/models/custom"

// FadeBack 用户反馈表
type FadeBack struct {
	custom.BaseModel
	Email        string `json:"email" gorm:"size:64; not null; comment: 用户邮箱"`
	Content      string `json:"content" gorm:"size:256; not null; comment: 反馈内容"`
	IsReply      bool   `json:"is_reply" gorm:"default:false; comment: 是否回复"`
	ReplyContent string `json:"reply_content" gorm:"size:256; comment: 回复内容"`
}
