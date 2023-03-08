/**
  @author: Zero
  @date: 2023/3/4 22:31:22
  @desc: 用户登录日志表

**/

package models

import "summer/models/custom"

// LoginLog 用户登录日志表结构
type LoginLog struct {
	custom.BaseModel
	IP       string `json:"ip" gorm:"size:32; comment: 用户IP"`
	Addr     string `json:"addr" gorm:"size:50; comment: 用户地址"`
	UserID   uint64 `json:"userId" gorm:"comment: 用户ID"`
	NickName string `json:"nickName" gorm:"size:42; comment: 用户昵称"`
	Token    string `json:"token" gorm:"size:256; comment: 登录令牌"`
	Device   string `json:"device" gorm:"size:256; comment: 登录设备"`
	// 登录用户
	User User `gorm:"foreignKey:UserID"`
}
