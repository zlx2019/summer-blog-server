/**
  @author: Zero
  @date: 2023/3/4 19:27:04
  @desc: 文章与标签关联表

**/

package models

// ArticleTag 文章与标签关系表结构
type ArticleTag struct {
	ID        uint64 `gorm:"comment: 主键ID"`
	ArticleID uint64 `gorm:"comment: 文章ID"`
	TagID     uint64 `gorm:"comment: 标签ID"`
}
