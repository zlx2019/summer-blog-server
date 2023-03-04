/**
  @author: Zero
  @date: 2023/3/4 20:02:17
  @desc: 用户收藏文章关系表

**/

package models

import (
	"time"
)

// UserLikeArticle 自定义用户收藏文章结构
type UserLikeArticle struct {
	ID         uint64    `gorm:"comment: 主键ID"`
	UserID     uint64    `gorm:"index; not null; comment: 用户ID"`
	User       User      `gorm:"foreignKey:UserID"`
	ArticleID  uint64    `gorm:"not null; comment: 收藏的文章ID"`
	Article    Article   `gorm:"foreignKey:ArticleID"`
	CreateTime time.Time `gorm:"autoCreateTime; comment: 创建时间"`
}
