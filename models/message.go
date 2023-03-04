/**
  @author: Zero
  @date: 2023/3/4 21:11:11
  @desc: 聊天消息表

**/

package models

import "summer/models/custom"

// Message 聊天消息表结构
type Message struct {
	custom.BaseModel
	// 消息内容
	Content string `json:"content" gorm:"comment: 消息内容"`
	// 消息是否已读
	IsRead bool `json:"isRead" gorm:"default: false; comment: 是否已读"`

	// 发送者
	SendUserID         uint64 `json:"sendUserID" gorm:"index; not null; comment: 发送者ID"`
	SendUser           User   `gorm:"foreignKey:SendUserID; comment:发送人"`
	SendUserNickName   string `json:"sendUserNickName" gorm:"comment: 发送人昵称"`
	SendUserNickAvatar string `json:"sendUserNickAvatar" gorm:"comment: 发送人头像"`

	//接收者
	RevUserID         uint64 `json:"revUserID" gorm:"index; not null; comment: 接收者ID"`
	RevUser           User   `gorm:"foreignKey:RevUserID; comment:接收者"`
	RevUserNickName   string `json:"revUserNickName" gorm:"comment: 接收者昵称"`
	RevUserNickAvatar string `json:"revUserNickAvatar" gorm:"comment: 接收者头像"`
}
